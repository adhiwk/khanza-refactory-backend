// Package triase use case triase IGD (RMTriaseIGD): penilaian primer (zona merah, skala 1/2) atau
// sekunder (skala 3/4/5) beserta kode skala, disimpan dalam satu transaksi.
package triase

import (
	"fmt"
	"strings"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	cetakmodel "goravel/app/models/cetak"
	model "goravel/app/models/triase"
	repo "goravel/app/repository/triase"
	cetaksvc "goravel/app/services/cetak"
	"goravel/app/support"
)

var (
	ErrRegistrasiNotFound = support.NotFound("data registrasi tidak ditemukan")
	ErrNotFound           = support.NotFound("data triase tidak ditemukan")
	ErrAlreadyExists      = support.Conflict("triase untuk rawat ini sudah ada")
	ErrKasusNotFound      = support.NotFound("data macam kasus triase tidak ditemukan")
	ErrPetugasNotFound    = support.NotFound("data petugas triase tidak ditemukan")
	ErrTanpaSkala         = support.Invalid("pilih minimal satu kode skala")
	ErrTanpaPegawai       = support.Forbidden("akun belum dihubungkan ke pegawai (users.kd_pegawai)")
	ErrBukanPetugas       = support.Forbidden("hanya bisa diisi/diubah/dihapus oleh petugas triase yang bersangkutan")
)

// aturan per jenis triase.
var aturan = map[string]struct {
	skala []int
	plan  []string
	label string
}{
	model.Primer:   {[]int{1, 2}, []string{"Ruang Resusitasi", "Ruang Kritis"}, "keluhan utama"},
	model.Sekunder: {[]int{3, 4, 5}, []string{"Zona Kuning", "Zona Hijau"}, "anamnesa singkat"},
}

// Actor pengguna yang menjalankan aksi.
type Actor struct {
	KodePegawai string
	SuperAdmin  bool
}

type Action struct {
	repo  repo.Repository
	cetak *cetaksvc.Service
}

func NewAction(repo repo.Repository, cetak *cetaksvc.Service) *Action {
	return &Action{repo: repo, cetak: cetak}
}

// Cetak lembar triase IGD.
func (a *Action) Cetak(noRawat string) (*cetakmodel.Dokumen, error) {
	t, err := a.Detail(noRawat)
	if err != nil {
		return nil, err
	}
	dok, err := a.cetak.Dokumen("Triase Instalasi Gawat Darurat", "")
	if err != nil {
		return nil, err
	}
	if dok.Identitas, err = a.cetak.IdentitasRawat(t.NoRawat); err != nil {
		return nil, err
	}
	dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Judul: "Data Kedatangan & Tanda Vital",
		Baris: a.cetak.Baris(t.Utama, map[string]bool{"no_rawat": true, "kode_kasus": true}, nil)})
	if p := t.Penilaian; p != nil {
		judul, keluhan := "Triase Primer (Zona Merah)", "Keluhan Utama"
		if t.Jenis == model.Sekunder {
			judul, keluhan = "Triase Sekunder", "Anamnesa Singkat"
		}
		baris := []cetakmodel.Baris{{Label: keluhan, Nilai: cetaksvc.Format(p.Keluhan)}}
		if t.Jenis == model.Primer {
			baris = append(baris, cetakmodel.Baris{Label: "Kebutuhan Khusus", Nilai: cetaksvc.Format(p.KebutuhanKhusus)})
		}
		baris = append(baris, cetakmodel.Baris{Label: "Catatan", Nilai: cetaksvc.Format(p.Catatan)},
			cetakmodel.Baris{Label: "Keputusan", Nilai: cetaksvc.Format(p.Plan)}, cetakmodel.Baris{Label: "Tanggal Triase", Nilai: cetaksvc.Format(p.TanggalTriase)})
		dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Judul: judul, Baris: baris})
		dok.TandaTangan = []cetakmodel.TandaTangan{{Peran: "Petugas Triase", Nama: p.NmPetugas, Kode: p.Nik}}
	}
	tabel := &cetakmodel.Tabel{Kolom: []string{"Pemeriksaan", "Kode", "Pengkajian Skala " + fmt.Sprint(t.Skala.Level)}}
	for _, it := range t.Skala.Items {
		tabel.Isi = append(tabel.Isi, []string{it.NamaPemeriksaan, it.Kode, it.Pengkajian})
	}
	dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Judul: "Skala Triase", Tabel: tabel})
	return dok, nil
}

func (a *Action) List(f repo.Filter, page, limit int) ([]model.Ringkas, int64, error) {
	return a.repo.Paginate(f, page, limit)
}

func (a *Action) Detail(noRawat string) (*model.Triase, error) {
	t, err := a.repo.Find(strings.TrimSpace(noRawat))
	if err != nil {
		return nil, err
	}
	if t == nil {
		return nil, ErrNotFound
	}
	return t, nil
}

func (a *Action) Create(in model.Input, actor Actor) (*model.Triase, error) {
	in.NoRawat = strings.TrimSpace(in.NoRawat)
	k, err := a.repo.Konteks(in.NoRawat)
	if err != nil {
		return nil, err
	}
	if k == nil {
		return nil, ErrRegistrasiNotFound
	}
	exists, err := a.repo.Exists(in.NoRawat)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}
	if err := a.validasi(&in, actor); err != nil {
		return nil, err
	}
	if err := support.Transaction(func(tx contractsorm.Query) error { return a.repo.Insert(tx, in) }); err != nil {
		return nil, err
	}
	return a.Detail(in.NoRawat)
}

// Update mengganti isi triase; jenis boleh berubah (primer <-> sekunder).
func (a *Action) Update(noRawat string, in model.Input, actor Actor) (*model.Triase, error) {
	old, err := a.editable(noRawat, actor)
	if err != nil {
		return nil, err
	}
	in.NoRawat = old.NoRawat
	if err := a.validasi(&in, actor); err != nil {
		return nil, err
	}
	if err := support.Transaction(func(tx contractsorm.Query) error { return a.repo.Update(tx, in, old.Jenis) }); err != nil {
		return nil, err
	}
	return a.Detail(in.NoRawat)
}

func (a *Action) Delete(noRawat string, actor Actor) error {
	old, err := a.editable(noRawat, actor)
	if err != nil {
		return err
	}
	return support.Transaction(func(tx contractsorm.Query) error { return a.repo.Delete(tx, old.NoRawat) })
}

// editable hanya petugas triase (nik primer/sekunder) yang boleh mengubah/menghapus, kecuali Admin Utama.
func (a *Action) editable(noRawat string, actor Actor) (*model.Triase, error) {
	t, err := a.Detail(noRawat)
	if err != nil {
		return nil, err
	}
	if actor.SuperAdmin || t.Penilaian == nil {
		return t, nil
	}
	return t, pemilik(t.Penilaian.Nik, actor)
}

func pemilik(nik string, actor Actor) error {
	if actor.KodePegawai == "" {
		return ErrTanpaPegawai
	}
	if nik != actor.KodePegawai {
		return ErrBukanPetugas
	}
	return nil
}

func (a *Action) validasi(in *model.Input, actor Actor) error {
	r, ok := aturan[in.Jenis]
	if !ok {
		return support.Invalid("jenis triase harus primer atau sekunder")
	}
	if !containsInt(r.skala, in.SkalaLevel) {
		return support.Invalid(fmt.Sprintf("triase %s memakai skala %v", in.Jenis, r.skala))
	}
	if !containsStr(r.plan, in.Penilaian.Plan) {
		return support.Invalid("keputusan triase " + in.Jenis + " harus " + strings.Join(r.plan, " atau "))
	}
	if strings.TrimSpace(in.Penilaian.Keluhan) == "" {
		return support.Invalid(r.label + " wajib diisi")
	}
	if in.Jenis == model.Sekunder {
		in.Penilaian.KebutuhanKhusus = ""
	} else if in.Penilaian.KebutuhanKhusus == "" {
		in.Penilaian.KebutuhanKhusus = "-"
	}

	tglKunjungan, err := support.ParseDateTime(in.TglKunjungan)
	if err != nil {
		return err
	}
	tglTriase, err := support.ParseDateTime(in.Penilaian.TanggalTriase)
	if err != nil {
		return err
	}
	if tglKunjungan == nil || tglTriase == nil {
		return support.ErrInvalidDateTime
	}

	in.Penilaian.Nik = strings.TrimSpace(in.Penilaian.Nik)
	if !actor.SuperAdmin {
		if err := pemilik(in.Penilaian.Nik, actor); err != nil {
			return err
		}
	}
	if ok, err := a.repo.PegawaiExists(in.Penilaian.Nik); err != nil {
		return err
	} else if !ok {
		return ErrPetugasNotFound
	}
	in.KodeKasus = strings.TrimSpace(in.KodeKasus)
	if ok, err := a.repo.KasusExists(in.KodeKasus); err != nil {
		return err
	} else if !ok {
		return ErrKasusNotFound
	}

	seen := map[string]bool{}
	var kode []string
	for _, k := range in.SkalaKode {
		if k = strings.TrimSpace(k); k == "" || seen[k] {
			continue
		}
		seen[k] = true
		ok, err := a.repo.SkalaExists(in.SkalaLevel, k)
		if err != nil {
			return err
		}
		if !ok {
			return support.NotFound(fmt.Sprintf("kode skala %d %s tidak ditemukan", in.SkalaLevel, k))
		}
		kode = append(kode, k)
	}
	if len(kode) == 0 {
		return ErrTanpaSkala
	}
	in.SkalaKode = kode
	return nil
}

func containsInt(list []int, v int) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}

func containsStr(list []string, v string) bool {
	for _, x := range list {
		if x == v {
			return true
		}
	}
	return false
}
