package kamar

import (
	model "goravel/app/models/kamar"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Kamar, int64, error)
	Find(key string) (*model.Kamar, error)
	Exists(key string) (bool, error)
	Create(data *model.Kamar) error
	Save(data *model.Kamar) error
	Delete(key string) error
	BangsalExists(value any) (bool, error)
}

type repository struct {
	crud.Table[model.Kamar, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Kamar, string]{
		Key:    "kd_kamar",
		Search: []string{"kd_kamar", "kd_bangsal", "kelas", "status"},
		Active: "statusdata",
		Order:  "kd_kamar",
	}}
}

func (r *repository) BangsalExists(value any) (bool, error) {
	return crud.ExistsIn(r.Query(), "bangsal", "kd_bangsal", value)
}
