package registrasi

import (
	"errors"
	"strings"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/goravel/framework/facades"

	regrequest "goravel/app/http/requests/registrasi"
	regmodel "goravel/app/models/registrasi"
	regrepo "goravel/app/repository/registrasi"
)

var (
	ErrNotFound           = errors.New("data registrasi tidak ditemukan")
	ErrPasienNotFound     = errors.New("data pasien tidak ditemukan")
	ErrDokterNotFound     = errors.New("data dokter tidak ditemukan")
	ErrPoliNotFound       = errors.New("data poliklinik tidak ditemukan")
	ErrPenjabNotFound     = errors.New("data jenis bayar tidak ditemukan")
	ErrDirawatInap        = errors.New("pasien sedang dalam masa perawatan di kamar inap")
	ErrLocked             = errors.New("data billing sudah terverifikasi / registrasi batal, data tidak boleh diubah")
	ErrExpired            = errors.New("perubahan data / penghapusan data tidak boleh lebih dari 2 x 24 jam")
	ErrInvalidDate        = errors.New("format tanggal harus YYYY-MM-DD")
	ErrInvalidTime        = errors.New("format jam harus HH:MM:SS")
	ErrNomorTidakTersedia = errors.New("gagal membuat no rawat, silakan coba lagi")
)

// maxAttempt jumlah percobaan simpan bila no_rawat bentrok (DlgReg mencoba 5 kali).
const maxAttempt = 5

// batasUbah batas waktu edit/hapus registrasi untuk non super admin (cekTanggal48jam).
const batasUbah = 48 * time.Hour

type Action struct {
	repo regrepo.Repository
}

func NewAction(repo regrepo.Repository) *Action {
	return &Action{
		repo: repo,
	}
}

func (a *Action) List(filter regrepo.Filter) ([]regmodel.RegPeriksaView, int64, error) {
	filter.Search = strings.TrimSpace(filter.Search)
	today := time.Now().Format("2006-01-02")
	if filter.TglAwal == "" {
		filter.TglAwal = today
	}
	if filter.TglAkhir == "" {
		filter.TglAkhir = today
	}
	return a.repo.Paginate(filter)
}

func (a *Action) Detail(noRawat string) (*regmodel.RegPeriksa, error) {
	reg, err := a.repo.FindByNoRawat(noRawat)
	if err != nil {
		return nil, err
	}
	if reg == nil {
		return nil, ErrNotFound
	}
	return reg, nil
}

// Create mengikuti BtnSimpan + isRegistrasi di DlgReg: no_reg & no_rawat dibuat otomatis,
// lalu dicoba ulang dengan nomor baru bila no_rawat sudah dipakai registrasi lain.
func (a *Action) Create(noRkmMedis string, data regrequest.RegistrasiData) (*regmodel.RegPeriksa, error) {
	noRkmMedis = strings.TrimSpace(noRkmMedis)
	pasien, err := a.repo.FindPasien(noRkmMedis)
	if err != nil {
		return nil, err
	}
	if pasien == nil {
		return nil, ErrPasienNotFound
	}

	dirawat, err := a.repo.IsDirawatInap(noRkmMedis)
	if err != nil {
		return nil, err
	}
	if dirawat {
		return nil, ErrDirawatInap
	}

	reg := &regmodel.RegPeriksa{
		NoRkmMedis:   &noRkmMedis,
		Stts:         strPtr("Belum"),
		StatusLanjut: "Ralan",
		StatusBayar:  "Belum Bayar",
	}
	if err := a.fill(reg, pasien, data); err != nil {
		return nil, err
	}

	pernah, err := a.repo.PernahKePoli(noRkmMedis, *reg.KdPoli)
	if err != nil {
		return nil, err
	}
	reg.StatusPoli = "Baru"
	if pernah {
		reg.StatusPoli = "Lama"
	}

	urut := facades.Config().GetString("registrasi.urut_no_reg", "dokter")
	for attempt := 1; ; attempt++ {
		noReg, err := a.repo.NextNoReg(urut, *reg.KdDokter, *reg.KdPoli, *reg.TglRegistrasi)
		if err != nil {
			return nil, err
		}
		noRawat, err := a.repo.NextNoRawat(*reg.TglRegistrasi)
		if err != nil {
			return nil, err
		}
		reg.NoReg = &noReg
		reg.NoRawat = noRawat

		err = a.repo.Create(reg)
		if err == nil {
			break
		}
		if !isDuplicate(err) {
			return nil, err
		}
		if attempt == maxAttempt {
			return nil, ErrNomorTidakTersedia
		}
	}

	if err := a.repo.UpdateUmurPasien(noRkmMedis); err != nil {
		facades.Log().Errorf("registrasi: gagal update umur pasien %s: %v", noRkmMedis, err)
	}
	return reg, nil
}

// Update mengikuti BtnEdit + ganti di DlgReg; no_rawat, no_reg & pasien tidak berubah.
func (a *Action) Update(noRawat string, data regrequest.RegistrasiData, superAdmin bool) (*regmodel.RegPeriksa, error) {
	reg, err := a.editable(noRawat, superAdmin)
	if err != nil {
		return nil, err
	}

	pasien, err := a.repo.FindPasien(deref(reg.NoRkmMedis))
	if err != nil {
		return nil, err
	}
	if pasien == nil {
		return nil, ErrPasienNotFound
	}

	if err := a.fill(reg, pasien, data); err != nil {
		return nil, err
	}

	err = a.repo.Save(reg)
	return reg, err
}

// Delete mengikuti BtnHapus di DlgReg, ditambah larangan hapus bila billing sudah ada.
func (a *Action) Delete(noRawat string, superAdmin bool) error {
	if _, err := a.editable(noRawat, superAdmin); err != nil {
		return err
	}
	return a.repo.Delete(noRawat)
}

func (a *Action) editable(noRawat string, superAdmin bool) (*regmodel.RegPeriksa, error) {
	reg, err := a.Detail(noRawat)
	if err != nil {
		return nil, err
	}

	billing, err := a.repo.HasBilling(noRawat)
	if err != nil {
		return nil, err
	}
	if billing || deref(reg.Stts) == "Batal" {
		return nil, ErrLocked
	}

	if !superAdmin && reg.TglRegistrasi != nil {
		waktu, err := time.ParseInLocation("2006-01-02 15:04:05", reg.TglRegistrasi.Format("2006-01-02")+" "+deref(reg.JamReg), time.Local)
		if err == nil && time.Since(waktu) > batasUbah {
			return nil, ErrExpired
		}
	}
	return reg, nil
}

// fill mengisi kolom registrasi dari request dan data pasien (isCekPasien di DlgReg):
// penanggung jawab & jenis bayar default dari pasien, stts_daftar Baru bila pasien didaftarkan
// pada tanggal registrasi, biaya default dari tarif poli baru/lama, umur dihitung dari tgl_lahir.
func (a *Action) fill(reg *regmodel.RegPeriksa, pasien *regrepo.PasienInfo, d regrequest.RegistrasiData) error {
	now := time.Now()

	tgl := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	if s := strings.TrimSpace(d.TglRegistrasi); s != "" {
		t, err := time.ParseInLocation("2006-01-02", s, time.Local)
		if err != nil {
			return ErrInvalidDate
		}
		tgl = t
	}
	jam := now.Format("15:04:05")
	if s := strings.TrimSpace(d.JamReg); s != "" {
		if _, err := time.Parse("15:04:05", s); err != nil {
			return ErrInvalidTime
		}
		jam = s
	}

	kdDokter := strings.TrimSpace(d.KdDokter)
	ok, err := a.repo.DokterExists(kdDokter)
	if err != nil {
		return err
	}
	if !ok {
		return ErrDokterNotFound
	}

	kdPj := strings.TrimSpace(d.KdPj)
	if kdPj == "" {
		kdPj = pasien.KdPj
	}
	ok, err = a.repo.PenjabExists(kdPj)
	if err != nil {
		return err
	}
	if !ok {
		return ErrPenjabNotFound
	}

	sttsDaftar := "Lama"
	if pasien.TglDaftar != nil && pasien.TglDaftar.Format("2006-01-02") == tgl.Format("2006-01-02") {
		sttsDaftar = "Baru"
	}

	kdPoli := strings.TrimSpace(d.KdPoli)
	biaya, ok, err := a.repo.BiayaPoli(kdPoli, sttsDaftar == "Baru")
	if err != nil {
		return err
	}
	if !ok {
		return ErrPoliNotFound
	}
	if d.BiayaReg != nil {
		biaya = *d.BiayaReg
	}

	umur, sttsUmur := hitungUmur(pasien.TglLahir, now)

	reg.TglRegistrasi = &tgl
	reg.JamReg = &jam
	reg.KdDokter = &kdDokter
	reg.KdPoli = &kdPoli
	reg.KdPj = kdPj
	reg.PJawab = strPtr(orDefault(d.PJawab, pasien.NamaKeluarga))
	reg.AlmtPj = strPtr(orDefault(d.AlmtPj, pasien.Asal))
	reg.HubunganPj = strPtr(orDefault(d.HubunganPj, pasien.Keluarga))
	reg.BiayaReg = &biaya
	reg.SttsDaftar = sttsDaftar
	reg.UmurDaftar = &umur
	reg.SttsUmur = &sttsUmur
	return nil
}

// hitungUmur: tahun bila >0, lalu bulan, lalu hari (sama dengan DlgReg, dihitung terhadap hari ini).
func hitungUmur(tglLahir *time.Time, now time.Time) (int, string) {
	if tglLahir == nil {
		return 0, "Th"
	}
	lahir := time.Date(tglLahir.Year(), tglLahir.Month(), tglLahir.Day(), 0, 0, 0, 0, time.Local)
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)

	bulan := (today.Year()-lahir.Year())*12 + int(today.Month()-lahir.Month())
	if today.Day() < lahir.Day() {
		bulan--
	}
	if bulan < 0 {
		return 0, "Th"
	}
	if bulan >= 12 {
		return bulan / 12, "Th"
	}
	if bulan > 0 {
		return bulan, "Bl"
	}
	return int(today.Sub(lahir).Hours() / 24), "Hr"
}

func isDuplicate(err error) bool {
	var myErr *mysql.MySQLError
	return errors.As(err, &myErr) && myErr.Number == 1062
}

func orDefault(s, def string) string {
	if s = strings.TrimSpace(s); s != "" {
		return s
	}
	return def
}

func strPtr(s string) *string {
	return &s
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
