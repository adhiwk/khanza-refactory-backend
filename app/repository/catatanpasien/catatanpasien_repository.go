package catatanpasien

import (
	model "goravel/app/models/catatanpasien"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.CatatanPasien, int64, error)
	Find(key string) (*model.CatatanPasien, error)
	Exists(key string) (bool, error)
	Create(data *model.CatatanPasien) error
	Save(data *model.CatatanPasien) error
	Delete(key string) error
	PasienExists(value any) (bool, error)
}

type repository struct {
	crud.Table[model.CatatanPasien, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.CatatanPasien, string]{
		Key:    "no_rkm_medis",
		Search: []string{"no_rkm_medis", "catatan"},
		Active: "",
		Order:  "no_rkm_medis",
	}}
}

func (r *repository) PasienExists(value any) (bool, error) {
	return crud.ExistsIn(r.Query(), "pasien", "no_rkm_medis", value)
}
