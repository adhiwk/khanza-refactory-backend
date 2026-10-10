package cacatfisik

import (
	model "goravel/app/models/cacatfisik"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.CacatFisik, int64, error)
	Find(key int) (*model.CacatFisik, error)
	Exists(key int) (bool, error)
	Create(data *model.CacatFisik) error
	Save(data *model.CacatFisik) error
	Delete(key int) error
}

type repository struct {
	crud.Table[model.CacatFisik, int]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.CacatFisik, int]{
		Key:    "id",
		Search: []string{"nama_cacat"},
		Active: "",
		Order:  "nama_cacat",
	}}
}
