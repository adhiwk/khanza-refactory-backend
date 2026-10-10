// Package kamarinap use case penempatan rawat inap (DlgKamarInap): masuk, pindah kamar, pulang, batal pulang.
package kamarinap

import (
	"math"
	"strings"
	"time"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/kamarinap"
	repo "goravel/app/repository/kamarinap"
	"goravel/app/support"
)

var (
	ErrRegistrasiNotFound = support.NotFound("data registrasi tidak ditemukan")
	ErrKamarNotFound      = support.NotFound("data kamar tidak ditemukan")
	ErrTidakDirawat       = support.NotFound("pasien tidak sedang menempati kamar inap")
	ErrBelumPulang        = support.Conflict("pasien belum dipulangkan")
	ErrSudahDirawat       = support.Conflict("pasien sedang dalam masa perawatan di kamar inap")
	ErrKamarTerisi        = support.Conflict("status kamar sudah terisi, silakan pilih kamar kosong")
	ErrKamarSama          = support.Invalid("kamar tujuan sama dengan kamar saat ini")
	ErrTerkunci           = support.Conflict("data billing sudah terverifikasi / registrasi batal, data tidak boleh diubah")
	ErrSebelumMasuk       = support.Invalid("waktu keluar tidak boleh sebelum waktu masuk kamar")
	ErrPindahTerlaluAwal  = support.Invalid("waktu pindah harus setelah waktu masuk kamar saat ini")
	ErrSebelumRegistrasi  = support.Invalid("waktu masuk tidak boleh sebelum waktu registrasi")
	ErrStatusPulang       = support.Invalid("status pulang tidak valid")
	ErrModePindah         = support.Invalid("mode pindah harus 1, 2, 3, atau 4")
)

// StatusPulang pilihan stts_pulang saat checkout ('-' dan 'Pindah Kamar' diatur sistem).
var StatusPulang = []string{"Sehat", "Rujuk", "APS", "+", "Meninggal", "Sembuh", "Membaik", "Pulang Paksa",
	"Status Belum Lengkap", "Atas Persetujuan Dokter", "Atas Permintaan Sendiri", "Isoman", "Lain-lain"}

// ModePindah opsi pindah kamar DlgKamarInap (Rganti1..Rganti4).
type ModePindah int

const (
	// GantiHapus kamar lama dihapus, pasien dianggap masuk kamar baru sejak waktu pindah.
	GantiHapus ModePindah = 1
	// GantiUbah kamar & tarif pada baris yang sama diganti (waktu masuk tetap).
	GantiUbah ModePindah = 2
	// Pindah baris lama ditutup "Pindah Kamar" dengan tarif lama, lalu baris baru dibuat.
	Pindah ModePindah = 3
	// PindahTarifTertinggi seperti Pindah, tetapi baris lama ditagih dengan tarif tertinggi (lama/baru).
	PindahTarifTertinggi ModePindah = 4
)

type MasukInput struct {
	NoRawat      string
	KdKamar      string
	DiagnosaAwal string
	Tgl, Jam     string
}

type PindahInput struct {
	NoRawat  string
	KdKamar  string
	Mode     ModePindah
	Tgl, Jam string
}

type PulangInput struct {
	NoRawat       string
	SttsPulang    string
	DiagnosaAkhir string
	Tgl, Jam      string
}

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(f repo.Filter) ([]model.KamarInap, int64, error) {
	today := time.Now().Format(support.DateLayout)
	if f.TglAwal == "" {
		f.TglAwal = today
	}
	if f.TglAkhir == "" {
		f.TglAkhir = today
	}
	return a.repo.Paginate(f)
}

func (a *Action) Riwayat(noRawat string) ([]model.KamarInap, error) {
	return a.repo.Riwayat(strings.TrimSpace(noRawat))
}

// Masuk check-in: kamar harus KOSONG & aktif; status kamar menjadi ISI dan registrasi menjadi Ranap.
func (a *Action) Masuk(in MasukInput) (*model.KamarInap, error) {
	k, err := a.repo.Konteks(strings.TrimSpace(in.NoRawat))
	if err != nil {
		return nil, err
	}
	if k == nil {
		return nil, ErrRegistrasiNotFound
	}
	if k.Terkunci() {
		return nil, ErrTerkunci
	}
	dirawat, err := a.repo.PasienDirawat(k.NoRkmMedis)
	if err != nil {
		return nil, err
	}
	if dirawat {
		return nil, ErrSudahDirawat
	}
	waktu, err := waktuInput(in.Tgl, in.Jam)
	if err != nil {
		return nil, err
	}
	if k.SebelumRegistrasi(waktu) {
		return nil, ErrSebelumRegistrasi
	}

	err = support.Transaction(func(tx contractsorm.Query) error {
		kamar, err := a.kamarKosong(tx, in.KdKamar)
		if err != nil {
			return err
		}
		if err := a.repo.Insert(tx, k.NoRawat, kamar.KdKamar, kamar.TrfKamar, strings.TrimSpace(in.DiagnosaAwal), "",
			waktu.Format(support.DateLayout), waktu.Format(support.TimeLayout)); err != nil {
			return err
		}
		if err := a.repo.SetStatusLanjut(tx, k.NoRawat, "Ranap"); err != nil {
			return err
		}
		return a.repo.SetStatusKamar(tx, kamar.KdKamar, "ISI")
	})
	if err != nil {
		return nil, err
	}
	return a.repo.Aktif(support.DB(), k.NoRawat)
}

// Pindah pindah/ganti kamar sesuai mode DlgKamarInap.
func (a *Action) Pindah(in PindahInput) (*model.KamarInap, error) {
	if in.Mode < GantiHapus || in.Mode > PindahTarifTertinggi {
		return nil, ErrModePindah
	}
	if err := a.cekTerkunci(in.NoRawat); err != nil {
		return nil, err
	}
	waktu, err := waktuInput(in.Tgl, in.Jam)
	if err != nil {
		return nil, err
	}
	set, err := a.repo.Pengaturan()
	if err != nil {
		return nil, err
	}
	noRawat := strings.TrimSpace(in.NoRawat)

	err = support.Transaction(func(tx contractsorm.Query) error {
		lama, err := a.repo.Aktif(tx, noRawat)
		if err != nil {
			return err
		}
		if lama == nil {
			return ErrTidakDirawat
		}
		if strings.TrimSpace(in.KdKamar) == lama.KdKamar {
			return ErrKamarSama
		}
		baru, err := a.kamarKosong(tx, in.KdKamar)
		if err != nil {
			return err
		}
		masuk, err := time.ParseInLocation(support.DateTimeLayout, lama.TglMasuk+" "+lama.JamMasuk, time.Local)
		if err != nil {
			return err
		}
		tgl, jam := waktu.Format(support.DateLayout), waktu.Format(support.TimeLayout)
		// Mode yang membuat baris baru memakai waktu pindah sebagai bagian primary key (no_rawat, tgl_masuk, jam_masuk).
		if in.Mode != GantiUbah && !waktu.After(masuk) {
			return ErrPindahTerlaluAwal
		}

		switch in.Mode {
		case GantiHapus:
			if err := a.repo.Insert(tx, noRawat, baru.KdKamar, baru.TrfKamar, lama.DiagnosaAwal, lama.DiagnosaAkhir, tgl, jam); err != nil {
				return err
			}
			if err := a.repo.Hapus(tx, lama.Key()); err != nil {
				return err
			}
		case GantiUbah:
			if err := a.repo.GantiKamar(tx, lama.Key(), baru.KdKamar, baru.TrfKamar); err != nil {
				return err
			}
		case Pindah, PindahTarifTertinggi:
			hari := LamaInap(masuk, waktu, set)
			trf := lama.TrfKamar
			if in.Mode == PindahTarifTertinggi {
				trf = math.Max(lama.TrfKamar, baru.TrfKamar)
			}
			if err := a.repo.Tutup(tx, lama.Key(), trf, tgl, jam, hari, hari*trf, "Pindah Kamar", lama.DiagnosaAkhir); err != nil {
				return err
			}
			if err := a.repo.Insert(tx, noRawat, baru.KdKamar, baru.TrfKamar, lama.DiagnosaAwal, lama.DiagnosaAkhir, tgl, jam); err != nil {
				return err
			}
		}
		if err := a.repo.SetStatusKamar(tx, lama.KdKamar, "KOSONG"); err != nil {
			return err
		}
		return a.repo.SetStatusKamar(tx, baru.KdKamar, "ISI")
	})
	if err != nil {
		return nil, err
	}
	return a.repo.Aktif(support.DB(), noRawat)
}

// Pulang checkout: hitung lama & biaya kamar, tutup baris aktif, kamar menjadi KOSONG.
func (a *Action) Pulang(in PulangInput) (*model.KamarInap, error) {
	if !validStatusPulang(in.SttsPulang) {
		return nil, ErrStatusPulang
	}
	if err := a.cekTerkunci(in.NoRawat); err != nil {
		return nil, err
	}
	waktu, err := waktuInput(in.Tgl, in.Jam)
	if err != nil {
		return nil, err
	}
	set, err := a.repo.Pengaturan()
	if err != nil {
		return nil, err
	}
	noRawat := strings.TrimSpace(in.NoRawat)

	var hasil *model.KamarInap
	err = support.Transaction(func(tx contractsorm.Query) error {
		row, err := a.repo.Aktif(tx, noRawat)
		if err != nil {
			return err
		}
		if row == nil {
			return ErrTidakDirawat
		}
		if _, err := a.repo.LockKamar(tx, row.KdKamar); err != nil {
			return err
		}
		masuk, err := time.ParseInLocation(support.DateTimeLayout, row.TglMasuk+" "+row.JamMasuk, time.Local)
		if err != nil {
			return err
		}
		if waktu.Before(masuk) {
			return ErrSebelumMasuk
		}
		hari := LamaInap(masuk, waktu, set)
		if err := a.repo.Tutup(tx, row.Key(), row.TrfKamar, waktu.Format(support.DateLayout), waktu.Format(support.TimeLayout),
			hari, hari*row.TrfKamar, in.SttsPulang, strings.TrimSpace(in.DiagnosaAkhir)); err != nil {
			return err
		}
		hasil = row
		return a.repo.SetStatusKamar(tx, row.KdKamar, "KOSONG")
	})
	if err != nil {
		return nil, err
	}
	return a.repo.Terakhir(support.DB(), hasil.NoRawat)
}

// BatalPulang membuka kembali baris kamar terakhir; kamar harus masih kosong.
func (a *Action) BatalPulang(noRawat string) (*model.KamarInap, error) {
	if err := a.cekTerkunci(noRawat); err != nil {
		return nil, err
	}
	noRawat = strings.TrimSpace(noRawat)
	err := support.Transaction(func(tx contractsorm.Query) error {
		row, err := a.repo.Terakhir(tx, noRawat)
		if err != nil {
			return err
		}
		if row == nil {
			return ErrTidakDirawat
		}
		if row.SttsPulang == "-" || row.SttsPulang == "Pindah Kamar" {
			return ErrBelumPulang
		}
		if _, err := a.kamarKosong(tx, row.KdKamar); err != nil {
			return err
		}
		if err := a.repo.BukaKembali(tx, row.Key()); err != nil {
			return err
		}
		return a.repo.SetStatusKamar(tx, row.KdKamar, "ISI")
	})
	if err != nil {
		return nil, err
	}
	return a.repo.Aktif(support.DB(), noRawat)
}

func (a *Action) cekTerkunci(noRawat string) error {
	k, err := a.repo.Konteks(strings.TrimSpace(noRawat))
	if err != nil {
		return err
	}
	if k == nil {
		return ErrRegistrasiNotFound
	}
	if k.Terkunci() {
		return ErrTerkunci
	}
	return nil
}

func (a *Action) kamarKosong(tx contractsorm.Query, kdKamar string) (*model.Kamar, error) {
	kamar, err := a.repo.LockKamar(tx, strings.TrimSpace(kdKamar))
	if err != nil {
		return nil, err
	}
	if kamar == nil || kamar.Statusdata != "1" {
		return nil, ErrKamarNotFound
	}
	if kamar.Status != "KOSONG" {
		return nil, ErrKamarTerisi
	}
	return kamar, nil
}

// LamaInap jumlah hari kamar (rumus DlgKamarInap): selisih hari kalender; hari yang sama dihitung 1
// bila lebih dari jam minimal; ditambah 1 bila "hitung hari awal" aktif.
func LamaInap(masuk, keluar time.Time, set model.Pengaturan) float64 {
	d1 := time.Date(masuk.Year(), masuk.Month(), masuk.Day(), 0, 0, 0, 0, time.UTC)
	d2 := time.Date(keluar.Year(), keluar.Month(), keluar.Day(), 0, 0, 0, 0, time.UTC)
	hari := math.Round(d2.Sub(d1).Hours() / 24)
	if hari == 0 {
		if keluar.Sub(masuk).Seconds() > 3600*set.JamMinimal {
			hari = 1
		}
	}
	if set.HitungHariAwal {
		hari++
	}
	return hari
}

func waktuInput(tgl, jam string) (time.Time, error) {
	d, err := support.DateOrToday(tgl)
	if err != nil {
		return time.Time{}, err
	}
	j, err := support.TimeOrNow(jam)
	if err != nil {
		return time.Time{}, err
	}
	return time.ParseInLocation(support.DateTimeLayout, d.Format(support.DateLayout)+" "+j, time.Local)
}

func validStatusPulang(s string) bool {
	for _, v := range StatusPulang {
		if v == s {
			return true
		}
	}
	return false
}
