package spesialis

import (
	model "goravel/app/models/spesialis"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Spesialis, int64, error)
	Find(key string) (*model.Spesialis, error)
	Exists(key string) (bool, error)
	Create(data *model.Spesialis) error
	Save(data *model.Spesialis) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.Spesialis, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Spesialis, string]{
		Key:    "kd_sps",
		Search: []string{"kd_sps", "nm_sps"},
		Active: "",
		Order:  "kd_sps",
	}}
}
