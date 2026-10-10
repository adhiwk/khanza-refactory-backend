package golonganobat

import (
	model "goravel/app/models/golonganobat"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.GolonganBarang, int64, error)
	Find(key string) (*model.GolonganBarang, error)
	Exists(key string) (bool, error)
	Create(data *model.GolonganBarang) error
	Save(data *model.GolonganBarang) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.GolonganBarang, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.GolonganBarang, string]{
		Key:    "kode",
		Search: []string{"kode", "nama"},
		Active: "",
		Order:  "kode",
	}}
}
