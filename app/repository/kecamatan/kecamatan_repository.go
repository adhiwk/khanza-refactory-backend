package kecamatan

import (
	model "goravel/app/models/kecamatan"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Kecamatan, int64, error)
	Find(key int) (*model.Kecamatan, error)
	Exists(key int) (bool, error)
	Create(data *model.Kecamatan) error
	Save(data *model.Kecamatan) error
	Delete(key int) error
}

type repository struct {
	crud.Table[model.Kecamatan, int]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Kecamatan, int]{
		Key:    "kd_kec",
		Search: []string{"nm_kec"},
		Active: "",
		Order:  "nm_kec",
	}}
}
