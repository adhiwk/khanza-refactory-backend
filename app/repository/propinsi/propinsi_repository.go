package propinsi

import (
	model "goravel/app/models/propinsi"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Propinsi, int64, error)
	Find(key int) (*model.Propinsi, error)
	Exists(key int) (bool, error)
	Create(data *model.Propinsi) error
	Save(data *model.Propinsi) error
	Delete(key int) error
}

type repository struct {
	crud.Table[model.Propinsi, int]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Propinsi, int]{
		Key:    "kd_prop",
		Search: []string{"nm_prop"},
		Active: "",
		Order:  "nm_prop",
	}}
}
