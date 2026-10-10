package perusahaanpasien

import (
	model "goravel/app/models/perusahaanpasien"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.PerusahaanPasien, int64, error)
	Find(key string) (*model.PerusahaanPasien, error)
	Exists(key string) (bool, error)
	Create(data *model.PerusahaanPasien) error
	Save(data *model.PerusahaanPasien) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.PerusahaanPasien, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.PerusahaanPasien, string]{
		Key:    "kode_perusahaan",
		Search: []string{"kode_perusahaan", "nama_perusahaan", "alamat", "kota"},
		Active: "",
		Order:  "kode_perusahaan",
	}}
}
