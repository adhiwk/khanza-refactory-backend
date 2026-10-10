package poliklinik

import (
	model "goravel/app/models/poliklinik"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Poliklinik, int64, error)
	Find(key string) (*model.Poliklinik, error)
	Exists(key string) (bool, error)
	Create(data *model.Poliklinik) error
	Save(data *model.Poliklinik) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.Poliklinik, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Poliklinik, string]{
		Key:    "kd_poli",
		Search: []string{"kd_poli", "nm_poli"},
		Active: "status",
		Order:  "kd_poli",
	}}
}
