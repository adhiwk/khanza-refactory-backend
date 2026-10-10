package kategoriobat

import (
	model "goravel/app/models/kategoriobat"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.KategoriBarang, int64, error)
	Find(key string) (*model.KategoriBarang, error)
	Exists(key string) (bool, error)
	Create(data *model.KategoriBarang) error
	Save(data *model.KategoriBarang) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.KategoriBarang, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.KategoriBarang, string]{
		Key:    "kode",
		Search: []string{"kode", "nama"},
		Active: "",
		Order:  "kode",
	}}
}
