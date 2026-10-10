package pasienmati

import (
	model "goravel/app/models/pasienmati"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.PasienMati, int64, error)
	Find(key string) (*model.PasienMati, error)
	Exists(key string) (bool, error)
	Create(data *model.PasienMati) error
	Save(data *model.PasienMati) error
	Delete(key string) error
	DokterExists(value any) (bool, error)
	PasienExists(value any) (bool, error)
}

type repository struct {
	crud.Table[model.PasienMati, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.PasienMati, string]{
		Key:    "no_rkm_medis",
		Search: []string{"no_rkm_medis", "keterangan", "icd1"},
		Active: "",
		Order:  "tanggal desc, jam desc",
	}}
}

func (r *repository) DokterExists(value any) (bool, error) {
	return crud.ExistsIn(r.Query(), "dokter", "kd_dokter", value)
}

func (r *repository) PasienExists(value any) (bool, error) {
	return crud.ExistsIn(r.Query(), "pasien", "no_rkm_medis", value)
}
