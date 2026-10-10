package kategoriperawatan

import (
	model "goravel/app/models/kategoriperawatan"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.KategoriPerawatan, int64, error)
	Find(key string) (*model.KategoriPerawatan, error)
	Exists(key string) (bool, error)
	Create(data *model.KategoriPerawatan) error
	Save(data *model.KategoriPerawatan) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.KategoriPerawatan, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.KategoriPerawatan, string]{
		Key:    "kd_kategori",
		Search: []string{"kd_kategori", "nm_kategori"},
		Active: "",
		Order:  "kd_kategori",
	}}
}
