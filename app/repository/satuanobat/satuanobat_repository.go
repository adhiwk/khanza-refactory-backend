package satuanobat

import (
	model "goravel/app/models/satuanobat"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Kodesatuan, int64, error)
	Find(key string) (*model.Kodesatuan, error)
	Exists(key string) (bool, error)
	Create(data *model.Kodesatuan) error
	Save(data *model.Kodesatuan) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.Kodesatuan, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Kodesatuan, string]{
		Key:    "kode_sat",
		Search: []string{"kode_sat", "satuan"},
		Active: "",
		Order:  "kode_sat",
	}}
}
