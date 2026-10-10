// Package pasien use case data pasien (DlgPasien).
package pasien

import (
	"strings"

	request "goravel/app/http/requests/pasien"
	model "goravel/app/models/pasien"
	repo "goravel/app/repository/pasien"
	"goravel/app/support"
)

var (
	ErrNotFound      = support.NotFound("data pasien tidak ditemukan")
	ErrAlreadyExists = support.Conflict("no rekam medis sudah digunakan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.Pasien, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.Pasien, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.Pasien, error) {
	key := strings.TrimSpace(req.NoRkmMedis)
	exists, err := a.repo.Exists(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}

	data := &model.Pasien{NoRkmMedis: key}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if data.TglDaftar == nil {
		today := support.Today()
		data.TglDaftar = &today
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.Pasien, error) {
	data, err := a.Detail(key)
	if err != nil {
		return nil, err
	}
	if err := a.fill(data, d); err != nil {
		return nil, err
	}
	if err := a.repo.Save(data); err != nil {
		return nil, err
	}
	return data, nil
}

// Delete menghapus permanen sesuai DlgPasien.
func (a *Action) Delete(key string) error {
	if _, err := a.Detail(key); err != nil {
		return err
	}
	return a.repo.Delete(key)
}

// fill menyalin data request ke model; tgl_daftar kosong tidak menimpa nilai lama.
func (a *Action) fill(m *model.Pasien, d request.Data) error {
	refs := []struct {
		table, column string
		value         any
		label         string
	}{
		{"penjab", "kd_pj", strings.TrimSpace(d.KdPj), "jenis bayar"},
		{"kelurahan", "kd_kel", d.KdKel, "kelurahan"},
		{"kecamatan", "kd_kec", d.KdKec, "kecamatan"},
		{"kabupaten", "kd_kab", d.KdKab, "kabupaten"},
		{"propinsi", "kd_prop", d.KdProp, "propinsi"},
		{"perusahaan_pasien", "kode_perusahaan", strings.TrimSpace(d.PerusahaanPasien), "instansi/perusahaan pasien"},
		{"suku_bangsa", "id", d.SukuBangsa, "suku bangsa"},
		{"bahasa_pasien", "id", d.BahasaPasien, "bahasa pasien"},
		{"cacat_fisik", "id", d.CacatFisik, "cacat fisik"},
	}
	for _, r := range refs {
		ok, err := a.repo.RefExists(r.table, r.column, r.value)
		if err != nil {
			return err
		}
		if !ok {
			return support.NotFound("data " + r.label + " tidak ditemukan")
		}
	}
	tglLahir, err := support.ParseDate(d.TglLahir)
	if err != nil {
		return err
	}
	tglDaftar, err := support.ParseDate(d.TglDaftar)
	if err != nil {
		return err
	}

	m.NmPasien = support.Nullable(d.NmPasien)
	m.NoKtp = support.Nullable(d.NoKtp)
	m.Jk = support.Nullable(d.Jk)
	m.TmpLahir = support.Nullable(d.TmpLahir)
	m.TglLahir = tglLahir
	m.NmIbu = strings.TrimSpace(d.NmIbu)
	m.Alamat = support.Nullable(d.Alamat)
	m.GolDarah = support.Nullable(d.GolDarah)
	m.Pekerjaan = support.Nullable(d.Pekerjaan)
	m.SttsNikah = support.Nullable(d.SttsNikah)
	m.Agama = support.Nullable(d.Agama)
	if tglDaftar != nil {
		m.TglDaftar = tglDaftar
	}
	m.NoTlp = support.Nullable(d.NoTlp)
	m.Umur = strings.TrimSpace(d.Umur)
	m.Pnd = d.Pnd
	m.Keluarga = support.Nullable(d.Keluarga)
	m.NamaKeluarga = strings.TrimSpace(d.NamaKeluarga)
	m.KdPj = strings.TrimSpace(d.KdPj)
	m.NoPeserta = support.Nullable(d.NoPeserta)
	m.KdKel = d.KdKel
	m.KdKec = d.KdKec
	m.KdKab = d.KdKab
	m.PekerjaanPj = strings.TrimSpace(d.PekerjaanPj)
	m.AlamatPj = strings.TrimSpace(d.AlamatPj)
	m.KelurahanPj = strings.TrimSpace(d.KelurahanPj)
	m.KecamatanPj = strings.TrimSpace(d.KecamatanPj)
	m.KabupatenPj = strings.TrimSpace(d.KabupatenPj)
	m.PerusahaanPasien = strings.TrimSpace(d.PerusahaanPasien)
	m.SukuBangsa = d.SukuBangsa
	m.BahasaPasien = d.BahasaPasien
	m.CacatFisik = d.CacatFisik
	m.Email = strings.TrimSpace(d.Email)
	m.Nip = strings.TrimSpace(d.Nip)
	m.KdProp = d.KdProp
	m.PropinsiPj = strings.TrimSpace(d.PropinsiPj)
	return nil
}
