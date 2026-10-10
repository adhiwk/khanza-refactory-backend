package kelurahan

import (
	model "goravel/app/models/kelurahan"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Kelurahan, int64, error)
	Find(key int) (*model.Kelurahan, error)
	Exists(key int) (bool, error)
	Create(data *model.Kelurahan) error
	Save(data *model.Kelurahan) error
	Delete(key int) error
}

type repository struct {
	crud.Table[model.Kelurahan, int]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Kelurahan, int]{
		Key:    "kd_kel",
		Search: []string{"nm_kel"},
		Active: "",
		Order:  "nm_kel",
	}}
}
