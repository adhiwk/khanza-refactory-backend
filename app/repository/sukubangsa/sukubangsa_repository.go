package sukubangsa

import (
	model "goravel/app/models/sukubangsa"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.SukuBangsa, int64, error)
	Find(key int) (*model.SukuBangsa, error)
	Exists(key int) (bool, error)
	Create(data *model.SukuBangsa) error
	Save(data *model.SukuBangsa) error
	Delete(key int) error
}

type repository struct {
	crud.Table[model.SukuBangsa, int]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.SukuBangsa, int]{
		Key:    "id",
		Search: []string{"nama_suku_bangsa"},
		Active: "",
		Order:  "nama_suku_bangsa",
	}}
}
