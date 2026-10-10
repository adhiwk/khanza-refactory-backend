package jabatan

import (
	model "goravel/app/models/jabatan"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Jabatan, int64, error)
	Find(key string) (*model.Jabatan, error)
	Exists(key string) (bool, error)
	Create(data *model.Jabatan) error
	Save(data *model.Jabatan) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.Jabatan, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Jabatan, string]{
		Key:    "kd_jbtn",
		Search: []string{"kd_jbtn", "nm_jbtn"},
		Active: "",
		Order:  "kd_jbtn",
	}}
}
