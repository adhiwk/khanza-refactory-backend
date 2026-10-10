// Package obat use case master obat/alkes (DlgBarang, tabel databarang).
package obat

import (
	"strings"

	request "goravel/app/http/requests/obat"
	model "goravel/app/models/obat"
	repo "goravel/app/repository/obat"
	"goravel/app/support"
)

var (
	ErrNotFound                = support.NotFound("data obat tidak ditemukan")
	ErrAlreadyExists           = support.Conflict("kode barang sudah digunakan")
	ErrKodesatuanNotFound      = support.NotFound("data satuan tidak ditemukan")
	ErrJenisNotFound           = support.NotFound("data jenis obat tidak ditemukan")
	ErrIndustrifarmasiNotFound = support.NotFound("data industri farmasi tidak ditemukan")
	ErrKategoriBarangNotFound  = support.NotFound("data kategori obat tidak ditemukan")
	ErrGolonganBarangNotFound  = support.NotFound("data golongan obat tidak ditemukan")
)

// Filter filter tambahan list obat; Status kosong = hanya obat aktif.
type Filter struct {
	Status       string
	Kdjns        string
	KodeKategori string
	KodeGolongan string
}

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, f Filter, page, limit int) ([]model.Obat, int64, error) {
	return a.repo.PaginateWhere(search, map[string]string{
		"status":        f.Status,
		"kdjns":         f.Kdjns,
		"kode_kategori": f.KodeKategori,
		"kode_golongan": f.KodeGolongan,
	}, page, limit)
}

// Detail obat aktif maupun nonaktif.
func (a *Action) Detail(key string) (*model.Obat, error) {
	data, err := a.repo.FindAny(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.Obat, error) {
	key := strings.TrimSpace(req.KodeBrng)
	exists, err := a.repo.Exists(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}

	data := &model.Obat{KodeBrng: key, Status: "1"}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.Obat, error) {
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

// Delete menonaktifkan (status=0) sesuai form Khanza.
func (a *Action) Delete(key string) error {
	data, err := a.Detail(key)
	if err != nil {
		return err
	}
	if data.Status != "1" {
		return ErrNotFound
	}
	return a.repo.Delete(key)
}

// UpdateStatus mengaktifkan / menonaktifkan obat.
func (a *Action) UpdateStatus(key, status string) (*model.Obat, error) {
	if _, err := a.Detail(key); err != nil {
		return nil, err
	}
	if err := a.repo.SetActive(key, status); err != nil {
		return nil, err
	}
	return a.Detail(key)
}

func (a *Action) fill(m *model.Obat, d request.Data) error {
	refs := []struct {
		value string
		check func(any) (bool, error)
		err   error
	}{
		{d.KodeSatbesar, a.repo.KodesatuanExists, ErrKodesatuanNotFound},
		{d.KodeSat, a.repo.KodesatuanExists, ErrKodesatuanNotFound},
		{d.Kdjns, a.repo.JenisExists, ErrJenisNotFound},
		{d.KodeIndustri, a.repo.IndustrifarmasiExists, ErrIndustrifarmasiNotFound},
		{d.KodeKategori, a.repo.KategoriBarangExists, ErrKategoriBarangNotFound},
		{d.KodeGolongan, a.repo.GolonganBarangExists, ErrGolonganBarangNotFound},
	}
	for _, r := range refs {
		v := strings.TrimSpace(r.value)
		if v == "" {
			continue
		}
		ok, err := r.check(v)
		if err != nil {
			return err
		}
		if !ok {
			return r.err
		}
	}
	expire, err := support.ParseDate(d.Expire)
	if err != nil {
		return err
	}

	m.NamaBrng = strings.TrimSpace(d.NamaBrng)
	m.KodeSatbesar = strings.TrimSpace(d.KodeSatbesar)
	m.KodeSat = strings.TrimSpace(d.KodeSat)
	m.LetakBarang = strings.TrimSpace(d.LetakBarang)
	m.Dasar = d.Dasar
	m.HBeli = d.HBeli
	m.Ralan = d.Ralan
	m.Kelas1 = d.Kelas1
	m.Kelas2 = d.Kelas2
	m.Kelas3 = d.Kelas3
	m.Utama = d.Utama
	m.Vip = d.Vip
	m.Vvip = d.Vvip
	m.BeliLuar = d.BeliLuar
	m.JualBebas = d.JualBebas
	m.Karyawan = d.Karyawan
	m.StokMinimal = d.StokMinimal
	m.Kdjns = strings.TrimSpace(d.Kdjns)
	m.Isi = d.Isi
	m.Kapasitas = d.Kapasitas
	m.Expire = expire
	m.KodeIndustri = strings.TrimSpace(d.KodeIndustri)
	m.KodeKategori = strings.TrimSpace(d.KodeKategori)
	m.KodeGolongan = strings.TrimSpace(d.KodeGolongan)
	return nil
}
