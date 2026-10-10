// Package rekonsiliasi use case rekonsiliasi obat (RMRekonsiliasiObat) dan konfirmasi farmasi (RMCariRekonsiliasiObat).
package rekonsiliasi

import (
	"fmt"
	"strings"
	"time"

	contractsorm "github.com/goravel/framework/contracts/database/orm"

	cetakmodel "goravel/app/models/cetak"
	model "goravel/app/models/rekonsiliasi"
	repo "goravel/app/repository/rekonsiliasi"
	cetaksvc "goravel/app/services/cetak"
	"goravel/app/support"
)

var (
	ErrRegistrasiNotFound = support.NotFound("data registrasi tidak ditemukan")
	ErrNotFound           = support.NotFound("data rekonsiliasi obat tidak ditemukan")
	ErrPetugasNotFound    = support.NotFound("data petugas tidak ditemukan")
	ErrTanpaObat          = support.Invalid("isi minimal satu obat rekonsiliasi")
	ErrSebelumRegistrasi  = support.Invalid("waktu wawancara tidak boleh sebelum waktu registrasi")
	ErrTanpaPegawai       = support.Forbidden("akun belum dihubungkan ke pegawai (users.kd_pegawai)")
	ErrBukanPetugas       = support.Forbidden("petugas rekonsiliasi harus akun sendiri")
	ErrDikonfirmasi       = support.Forbidden("rekonsiliasi sudah dikonfirmasi oleh apoteker")
)

// Actor pengguna; BolehKonfirmasi setara akses.getkonfirmasi_rekonsiliasi_obat().
type Actor struct {
	KodePegawai     string
	SuperAdmin      bool
	BolehKonfirmasi bool
}

type Action struct {
	repo  repo.Repository
	cetak *cetaksvc.Service
}

func NewAction(repo repo.Repository, cetak *cetaksvc.Service) *Action {
	return &Action{repo: repo, cetak: cetak}
}

// Cetak lembar rekonsiliasi obat (rptDataRekonsiliasiObat).
func (a *Action) Cetak(no string) (*cetakmodel.Dokumen, error) {
	r, err := a.Detail(no)
	if err != nil {
		return nil, err
	}
	dok, err := a.cetak.Dokumen("Rekonsiliasi Obat", r.NoRekonsiliasi)
	if err != nil {
		return nil, err
	}
	if dok.Identitas, err = a.cetak.IdentitasRawat(r.NoRawat); err != nil {
		return nil, err
	}
	f := cetaksvc.Format
	dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Judul: "Data Rekonsiliasi", Baris: []cetakmodel.Baris{
		{Label: "Tanggal Wawancara", Nilai: f(r.TanggalWawancara)}, {Label: "Rekonsiliasi Saat", Nilai: f(r.RekonsiliasiObatSaat)},
		{Label: "Alergi Obat", Nilai: f(r.AlergiObat)}, {Label: "Manifestasi Alergi", Nilai: f(r.ManifestasiAlergi)},
		{Label: "Dampak Alergi", Nilai: f(r.DampakAlergi)},
	}})
	tabel := &cetakmodel.Tabel{Kolom: []string{"No", "Nama Obat", "Dosis", "Frekuensi", "Cara Pemberian", "Pemberian Terakhir", "Tindak Lanjut", "Perubahan Aturan Pakai"}}
	for i, o := range r.Obat {
		tabel.Isi = append(tabel.Isi, []string{fmt.Sprint(i + 1), o.NamaObat, f(o.DosisObat), f(o.Frekuensi), f(o.CaraPemberian),
			f(o.WaktuPemberianTerakhir), f(o.TindakLanjut), f(o.PerubahanAturanPakai)})
	}
	dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Judul: "Daftar Obat", Tabel: tabel})
	dok.TandaTangan = []cetakmodel.TandaTangan{{Peran: "Petugas", Nama: r.NmPetugas, Kode: r.Nip}}
	if k := r.Konfirmasi; k != nil {
		dok.Bagian = append(dok.Bagian, cetakmodel.Bagian{Judul: "Konfirmasi Farmasi", Baris: []cetakmodel.Baris{
			{Label: "Diterima Farmasi", Nilai: f(k.DiterimaFarmasi)}, {Label: "Dikonfirmasi Apoteker", Nilai: f(k.DikonfirmasiApoteker)},
			{Label: "Diserahkan ke Pasien", Nilai: f(k.DiserahkanPasien)},
		}})
		dok.TandaTangan = append(dok.TandaTangan, cetakmodel.TandaTangan{Peran: "Apoteker", Nama: k.NmPetugas, Kode: k.Nip})
	}
	return dok, nil
}

func (a *Action) List(f repo.Filter, page, limit int) ([]model.Rekonsiliasi, int64, error) {
	return a.repo.Paginate(f, page, limit)
}

func (a *Action) Detail(no string) (*model.Rekonsiliasi, error) {
	r, err := a.repo.Find(strings.TrimSpace(no))
	if err != nil {
		return nil, err
	}
	if r == nil {
		return nil, ErrNotFound
	}
	return r, nil
}

// Create nomor RO+tanggal dibuat otomatis; waktu wawancara kosong = sekarang.
func (a *Action) Create(in model.Rekonsiliasi, actor Actor) (*model.Rekonsiliasi, error) {
	in.NoRawat = strings.TrimSpace(in.NoRawat)
	k, err := a.repo.Konteks(in.NoRawat)
	if err != nil {
		return nil, err
	}
	if k == nil {
		return nil, ErrRegistrasiNotFound
	}
	waktu, err := a.validasi(&in, actor)
	if err != nil {
		return nil, err
	}
	if !actor.SuperAdmin && k.SebelumRegistrasi(waktu) {
		return nil, ErrSebelumRegistrasi
	}
	err = support.Transaction(func(tx contractsorm.Query) error {
		no, err := a.repo.NextNomor(tx, waktu.Format("20060102"))
		if err != nil {
			return err
		}
		in.NoRekonsiliasi = no
		if err := a.repo.Insert(tx, in); err != nil {
			return err
		}
		return a.repo.ReplaceObat(tx, no, in.Obat)
	})
	if err != nil {
		return nil, err
	}
	return a.Detail(in.NoRekonsiliasi)
}

// Update mengganti isi & daftar obat; rekonsiliasi terkonfirmasi hanya oleh pemegang hak konfirmasi.
func (a *Action) Update(no string, in model.Rekonsiliasi, actor Actor) (*model.Rekonsiliasi, error) {
	old, err := a.editable(no, actor)
	if err != nil {
		return nil, err
	}
	in.NoRekonsiliasi, in.NoRawat = old.NoRekonsiliasi, old.NoRawat
	if _, err := a.validasi(&in, actor); err != nil {
		return nil, err
	}
	err = support.Transaction(func(tx contractsorm.Query) error {
		if err := a.repo.Update(tx, in); err != nil {
			return err
		}
		return a.repo.ReplaceObat(tx, in.NoRekonsiliasi, in.Obat)
	})
	if err != nil {
		return nil, err
	}
	return a.Detail(in.NoRekonsiliasi)
}

func (a *Action) Delete(no string, actor Actor) error {
	old, err := a.editable(no, actor)
	if err != nil {
		return err
	}
	return support.Transaction(func(tx contractsorm.Query) error { return a.repo.Delete(tx, old.NoRekonsiliasi) })
}

// Konfirmasi mencatat diterima farmasi / dikonfirmasi apoteker / diserahkan ke pasien (mengganti konfirmasi lama).
func (a *Action) Konfirmasi(no string, k model.Konfirmasi) (*model.Rekonsiliasi, error) {
	r, err := a.Detail(no)
	if err != nil {
		return nil, err
	}
	k.Nip = strings.TrimSpace(k.Nip)
	if ok, err := a.repo.PetugasExists(k.Nip); err != nil {
		return nil, err
	} else if !ok {
		return nil, ErrPetugasNotFound
	}
	for _, s := range []*string{&k.DiterimaFarmasi, &k.DikonfirmasiApoteker, &k.DiserahkanPasien} {
		t, err := support.ParseDateTime(*s)
		if err != nil {
			return nil, err
		}
		if t == nil {
			return nil, support.ErrInvalidDateTime
		}
	}
	if err := support.Transaction(func(tx contractsorm.Query) error { return a.repo.SimpanKonfirmasi(tx, r.NoRekonsiliasi, k) }); err != nil {
		return nil, err
	}
	return a.Detail(r.NoRekonsiliasi)
}

func (a *Action) editable(no string, actor Actor) (*model.Rekonsiliasi, error) {
	r, err := a.Detail(no)
	if err != nil {
		return nil, err
	}
	if r.Dikonfirmasi && !actor.SuperAdmin && !actor.BolehKonfirmasi {
		return nil, ErrDikonfirmasi
	}
	return r, nil
}

func (a *Action) validasi(in *model.Rekonsiliasi, actor Actor) (time.Time, error) {
	waktu, err := support.ParseDateTime(in.TanggalWawancara)
	if err != nil {
		return time.Time{}, err
	}
	if waktu == nil {
		now := time.Now().Truncate(time.Second)
		waktu = &now
	}
	in.TanggalWawancara = waktu.Format(support.DateTimeLayout)
	in.Nip = strings.TrimSpace(in.Nip)
	if !actor.SuperAdmin {
		if actor.KodePegawai == "" {
			return time.Time{}, ErrTanpaPegawai
		}
		if in.Nip != actor.KodePegawai {
			return time.Time{}, ErrBukanPetugas
		}
	}
	if ok, err := a.repo.PetugasExists(in.Nip); err != nil {
		return time.Time{}, err
	} else if !ok {
		return time.Time{}, ErrPetugasNotFound
	}
	in.DampakAlergi = support.OrDefault(in.DampakAlergi, "-")
	var obat []model.Obat
	for _, o := range in.Obat {
		if o.NamaObat = strings.TrimSpace(o.NamaObat); o.NamaObat == "" {
			continue
		}
		o.TindakLanjut = support.OrDefault(o.TindakLanjut, "Lanjut")
		obat = append(obat, o)
	}
	if len(obat) == 0 {
		return time.Time{}, ErrTanpaObat
	}
	in.Obat = obat
	return *waktu, nil
}
