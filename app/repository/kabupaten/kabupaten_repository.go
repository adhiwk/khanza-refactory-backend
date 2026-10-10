package kabupaten

import (
	model "goravel/app/models/kabupaten"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Kabupaten, int64, error)
	Find(key int) (*model.Kabupaten, error)
	Exists(key int) (bool, error)
	Create(data *model.Kabupaten) error
	Save(data *model.Kabupaten) error
	Delete(key int) error
}

type repository struct {
	crud.Table[model.Kabupaten, int]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Kabupaten, int]{
		Key:    "kd_kab",
		Search: []string{"nm_kab"},
		Active: "",
		Order:  "nm_kab",
	}}
}
