package penjab

import (
	model "goravel/app/models/penjab"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Penjab, int64, error)
	Find(key string) (*model.Penjab, error)
	Exists(key string) (bool, error)
	Create(data *model.Penjab) error
	Save(data *model.Penjab) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.Penjab, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Penjab, string]{
		Key:    "kd_pj",
		Search: []string{"kd_pj", "png_jawab", "nama_perusahaan"},
		Active: "status",
		Order:  "png_jawab",
	}}
}
