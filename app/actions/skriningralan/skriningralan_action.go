// Package skriningralan use case skrining rawat jalan (RMSKriningRawatJalan): dilakukan sebelum registrasi
// sehingga dikunci per pasien + waktu, bukan per no_rawat.
package skriningralan

import (
	"strings"

	cetakmodel "goravel/app/models/cetak"
	model "goravel/app/models/skriningralan"
	repo "goravel/app/repository/skriningralan"
	cetaksvc "goravel/app/services/cetak"
	"goravel/app/support"
)

var (
	ErrNotFound        = support.NotFound("data skrining rawat jalan tidak ditemukan")
	ErrPasienNotFound  = support.NotFound("data pasien tidak ditemukan")
	ErrPetugasNotFound = support.NotFound("data petugas tidak ditemukan")
	ErrAlreadyExists   = support.Conflict("skrining pasien pada tanggal & jam tersebut sudah ada")
	ErrTanpaPegawai    = support.Forbidden("akun belum dihubungkan ke pegawai (users.kd_pegawai)")
	ErrBukanPetugas    = support.Forbidden("hanya bisa diisi/diubah/dihapus oleh petugas yang bersangkutan")
)

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

// Cetak lembar skrining rawat jalan.
func (a *Action) Cetak(key model.Key) (*cetakmodel.Dokumen, error) {
	s, err := a.Detail(key)
	if err != nil {
		return nil, err
	}
	dok, err := a.cetak.Dokumen("Skrining Pasien Rawat Jalan", "")
	if err != nil {
		return nil, err
	}
	if dok.Identitas, err = a.cetak.IdentitasPasien(s.NoRkmMedis); err != nil {
		return nil, err
	}
	dok.Bagian = []cetakmodel.Bagian{{Judul: "Hasil Skrining",
		Baris: a.cetak.Baris(s, map[string]bool{"no_rkm_medis": true, "nm_pasien": true, "nip": true, "nm_petugas": true}, nil)}}
	dok.TandaTangan = []cetakmodel.TandaTangan{{Peran: "Petugas", Nama: s.NmPetugas, Kode: s.Nip}}
	return dok, nil
}

func (a *Action) List(f repo.Filter, page, limit int) ([]model.Skrining, int64, error) {
	return a.repo.Paginate(f, page, limit)
}

func (a *Action) Detail(key model.Key) (*model.Skrining, error) {
	s, err := a.repo.Find(model.Key{Tanggal: strings.TrimSpace(key.Tanggal), Jam: strings.TrimSpace(key.Jam), NoRkmMedis: strings.TrimSpace(key.NoRkmMedis)})
	if err != nil {
		return nil, err
	}
	if s == nil {
		return nil, ErrNotFound
	}
	return s, nil
}

// Create tanggal/jam kosong = sekarang.
func (a *Action) Create(in model.Skrining, actor Actor) (*model.Skrining, error) {
	tgl, err := support.DateOrToday(in.Tanggal)
	if err != nil {
		return nil, err
	}
	if in.Jam, err = support.TimeOrNow(in.Jam); err != nil {
		return nil, err
	}
	in.Tanggal = tgl.Format(support.DateLayout)
	in.NoRkmMedis = strings.TrimSpace(in.NoRkmMedis)
	if ok, err := a.repo.PasienExists(in.NoRkmMedis); err != nil {
		return nil, err
	} else if !ok {
		return nil, ErrPasienNotFound
	}
	if err := a.petugas(&in, actor); err != nil {
		return nil, err
	}
	key := model.Key{Tanggal: in.Tanggal, Jam: in.Jam, NoRkmMedis: in.NoRkmMedis}
	if old, err := a.repo.Find(key); err != nil {
		return nil, err
	} else if old != nil {
		return nil, ErrAlreadyExists
	}
	if err := a.repo.Create(in); err != nil {
		return nil, err
	}
	return a.Detail(key)
}

func (a *Action) Update(key model.Key, in model.Skrining, actor Actor) (*model.Skrining, error) {
	old, err := a.editable(key, actor)
	if err != nil {
		return nil, err
	}
	in.Tanggal, in.Jam, in.NoRkmMedis = old.Tanggal, old.Jam, old.NoRkmMedis
	if err := a.petugas(&in, actor); err != nil {
		return nil, err
	}
	if err := a.repo.Update(in); err != nil {
		return nil, err
	}
	return a.Detail(model.Key{Tanggal: old.Tanggal, Jam: old.Jam, NoRkmMedis: old.NoRkmMedis})
}

// Delete hanya petugas pengisi kecuali Admin Utama (BtnHapus RMSKriningRawatJalan).
func (a *Action) Delete(key model.Key, actor Actor) error {
	old, err := a.editable(key, actor)
	if err != nil {
		return err
	}
	return a.repo.Delete(model.Key{Tanggal: old.Tanggal, Jam: old.Jam, NoRkmMedis: old.NoRkmMedis})
}

func (a *Action) editable(key model.Key, actor Actor) (*model.Skrining, error) {
	s, err := a.Detail(key)
	if err != nil {
		return nil, err
	}
	if actor.SuperAdmin {
		return s, nil
	}
	return s, pemilik(s.Nip, actor)
}

func pemilik(nip string, actor Actor) error {
	if actor.KodePegawai == "" {
		return ErrTanpaPegawai
	}
	if nip != actor.KodePegawai {
		return ErrBukanPetugas
	}
	return nil
}

func (a *Action) petugas(in *model.Skrining, actor Actor) error {
	in.Nip = strings.TrimSpace(in.Nip)
	if !actor.SuperAdmin {
		if err := pemilik(in.Nip, actor); err != nil {
			return err
		}
	}
	ok, err := a.repo.PetugasExists(in.Nip)
	if err != nil {
		return err
	}
	if !ok {
		return ErrPetugasNotFound
	}
	return nil
}
