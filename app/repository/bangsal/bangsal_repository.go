package bangsal

import (
	model "goravel/app/models/bangsal"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Bangsal, int64, error)
	Find(key string) (*model.Bangsal, error)
	Exists(key string) (bool, error)
	Create(data *model.Bangsal) error
	Save(data *model.Bangsal) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.Bangsal, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Bangsal, string]{
		Key:    "kd_bangsal",
		Search: []string{"kd_bangsal", "nm_bangsal"},
		Active: "status",
		Order:  "kd_bangsal",
	}}
}
