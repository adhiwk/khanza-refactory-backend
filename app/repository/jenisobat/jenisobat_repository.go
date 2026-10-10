package jenisobat

import (
	model "goravel/app/models/jenisobat"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Jenis, int64, error)
	Find(key string) (*model.Jenis, error)
	Exists(key string) (bool, error)
	Create(data *model.Jenis) error
	Save(data *model.Jenis) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.Jenis, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Jenis, string]{
		Key:    "kdjns",
		Search: []string{"kdjns", "nama"},
		Active: "",
		Order:  "kdjns",
	}}
}
