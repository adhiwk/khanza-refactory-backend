package industrifarmasi

import (
	model "goravel/app/models/industrifarmasi"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Industrifarmasi, int64, error)
	Find(key string) (*model.Industrifarmasi, error)
	Exists(key string) (bool, error)
	Create(data *model.Industrifarmasi) error
	Save(data *model.Industrifarmasi) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.Industrifarmasi, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Industrifarmasi, string]{
		Key:    "kode_industri",
		Search: []string{"kode_industri", "nama_industri", "kota"},
		Active: "",
		Order:  "kode_industri",
	}}
}
