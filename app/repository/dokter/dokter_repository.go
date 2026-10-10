package dokter

import (
	model "goravel/app/models/dokter"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Dokter, int64, error)
	Find(key string) (*model.Dokter, error)
	Exists(key string) (bool, error)
	Create(data *model.Dokter) error
	Save(data *model.Dokter) error
	Delete(key string) error
	SpesialisExists(value any) (bool, error)
	PegawaiExists(nik string) (bool, error)
}

type repository struct {
	crud.Table[model.Dokter, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Dokter, string]{
		Key:    "kd_dokter",
		Search: []string{"kd_dokter", "nm_dokter", "alumni", "no_ijn_praktek"},
		Active: "status",
		Order:  "kd_dokter",
	}}
}

func (r *repository) SpesialisExists(value any) (bool, error) {
	return crud.ExistsIn(r.Query(), "spesialis", "kd_sps", value)
}

// PegawaiExists kd_dokter wajib terdaftar sebagai pegawai.nik (FK dokter_ibfk_3).
func (r *repository) PegawaiExists(nik string) (bool, error) {
	return crud.ExistsIn(r.Query(), "pegawai", "nik", nik)
}
