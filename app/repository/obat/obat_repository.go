package obat

import (
	model "goravel/app/models/obat"
	"goravel/app/repository/crud"
)

type Repository interface {
	PaginateWhere(search string, eq map[string]string, page, limit int) ([]model.Obat, int64, error)
	Find(key string) (*model.Obat, error)
	FindAny(key string) (*model.Obat, error)
	Exists(key string) (bool, error)
	Create(data *model.Obat) error
	Save(data *model.Obat) error
	Delete(key string) error
	SetActive(key, status string) error
	KodesatuanExists(value any) (bool, error)
	JenisExists(value any) (bool, error)
	IndustrifarmasiExists(value any) (bool, error)
	KategoriBarangExists(value any) (bool, error)
	GolonganBarangExists(value any) (bool, error)
}

type repository struct {
	crud.Table[model.Obat, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Obat, string]{
		Key:    "kode_brng",
		Search: []string{"kode_brng", "nama_brng"},
		Active: "status",
		Order:  "nama_brng",
	}}
}

func (r *repository) KodesatuanExists(value any) (bool, error) {
	return crud.ExistsIn(r.Query(), "kodesatuan", "kode_sat", value)
}

func (r *repository) JenisExists(value any) (bool, error) {
	return crud.ExistsIn(r.Query(), "jenis", "kdjns", value)
}

func (r *repository) IndustrifarmasiExists(value any) (bool, error) {
	return crud.ExistsIn(r.Query(), "industrifarmasi", "kode_industri", value)
}

func (r *repository) KategoriBarangExists(value any) (bool, error) {
	return crud.ExistsIn(r.Query(), "kategori_barang", "kode", value)
}

func (r *repository) GolonganBarangExists(value any) (bool, error) {
	return crud.ExistsIn(r.Query(), "golongan_barang", "kode", value)
}
