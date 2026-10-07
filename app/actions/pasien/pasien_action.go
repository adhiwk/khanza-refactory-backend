package pasien

import (
	"errors"
	"strings"
	"time"

	pasienrequest "goravel/app/http/requests/pasien"
	pasienmodel "goravel/app/models/pasien"
	pasienrepo "goravel/app/repository/pasien"
)

var (
	ErrNotFound      = errors.New("data pasien tidak ditemukan")
	ErrAlreadyExists = errors.New("no rekam medis sudah digunakan")
	ErrInvalidDate   = errors.New("format tanggal harus YYYY-MM-DD")
)

type Action struct {
	repo pasienrepo.Repository
}

func NewAction(repo pasienrepo.Repository) *Action {
	return &Action{
		repo: repo,
	}
}

func (a *Action) List(search string, page, limit int) ([]pasienmodel.Pasien, int64, error) {
	return a.repo.GetPaginated(strings.TrimSpace(search), page, limit)
}

func (a *Action) Detail(noRkmMedis string) (*pasienmodel.Pasien, error) {
	pasien, err := a.repo.FindByNoRkmMedis(noRkmMedis)
	if err != nil {
		return nil, err
	}
	if pasien == nil {
		return nil, ErrNotFound
	}
	return pasien, nil
}

func (a *Action) Create(noRkmMedis string, data pasienrequest.PasienData) (*pasienmodel.Pasien, error) {
	noRkmMedis = strings.TrimSpace(noRkmMedis)

	exists, err := a.repo.ExistsByNoRkmMedis(noRkmMedis)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}

	pasien := &pasienmodel.Pasien{NoRkmMedis: noRkmMedis}
	if err := fill(pasien, data); err != nil {
		return nil, err
	}
	if pasien.TglDaftar == nil {
		today := time.Now().Truncate(24 * time.Hour)
		pasien.TglDaftar = &today
	}

	err = a.repo.Create(pasien)
	return pasien, err
}

func (a *Action) Update(noRkmMedis string, data pasienrequest.PasienData) (*pasienmodel.Pasien, error) {
	pasien, err := a.Detail(noRkmMedis)
	if err != nil {
		return nil, err
	}

	if err := fill(pasien, data); err != nil {
		return nil, err
	}

	err = a.repo.Update(pasien)
	return pasien, err
}

func (a *Action) Delete(noRkmMedis string) error {
	if _, err := a.Detail(noRkmMedis); err != nil {
		return err
	}
	return a.repo.Delete(noRkmMedis)
}

// fill menyalin data request ke model; tgl_daftar yang kosong tidak menimpa nilai lama.
func fill(p *pasienmodel.Pasien, d pasienrequest.PasienData) error {
	tglLahir, err := parseDate(d.TglLahir)
	if err != nil {
		return err
	}
	tglDaftar, err := parseDate(d.TglDaftar)
	if err != nil {
		return err
	}

	p.NmPasien = nullable(d.NmPasien)
	p.NoKtp = nullable(d.NoKtp)
	p.Jk = nullable(d.Jk)
	p.TmpLahir = nullable(d.TmpLahir)
	p.TglLahir = tglLahir
	p.NmIbu = d.NmIbu
	p.Alamat = nullable(d.Alamat)
	p.GolDarah = nullable(d.GolDarah)
	p.Pekerjaan = nullable(d.Pekerjaan)
	p.SttsNikah = nullable(d.SttsNikah)
	p.Agama = nullable(d.Agama)
	if tglDaftar != nil {
		p.TglDaftar = tglDaftar
	}
	p.NoTlp = nullable(d.NoTlp)
	p.Umur = d.Umur
	p.Pnd = d.Pnd
	p.Keluarga = nullable(d.Keluarga)
	p.NamaKeluarga = d.NamaKeluarga
	p.KdPj = d.KdPj
	p.NoPeserta = nullable(d.NoPeserta)
	p.KdKel = d.KdKel
	p.KdKec = d.KdKec
	p.KdKab = d.KdKab
	p.PekerjaanPj = d.PekerjaanPj
	p.AlamatPj = d.AlamatPj
	p.KelurahanPj = d.KelurahanPj
	p.KecamatanPj = d.KecamatanPj
	p.KabupatenPj = d.KabupatenPj
	p.PerusahaanPasien = d.PerusahaanPasien
	p.SukuBangsa = d.SukuBangsa
	p.BahasaPasien = d.BahasaPasien
	p.CacatFisik = d.CacatFisik
	p.Email = d.Email
	p.Nip = d.Nip
	p.KdProp = d.KdProp
	p.PropinsiPj = d.PropinsiPj
	return nil
}

func nullable(s string) *string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	return &s
}

func parseDate(s string) (*time.Time, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, nil
	}
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return nil, ErrInvalidDate
	}
	return &t, nil
}
