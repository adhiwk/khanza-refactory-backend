package pasien

import (
	model "goravel/app/models/pasien"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.Pasien, int64, error)
	Find(key string) (*model.Pasien, error)
	Exists(key string) (bool, error)
	Create(data *model.Pasien) error
	Save(data *model.Pasien) error
	Delete(key string) error
	// RefExists cek referensi master (FK pasien_ibfk_*); table & column hanya dari konstanta Action.
	RefExists(table, column string, value any) (bool, error)
}

type repository struct {
	crud.Table[model.Pasien, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.Pasien, string]{
		Key:    "no_rkm_medis",
		Search: []string{"no_rkm_medis", "nm_pasien", "no_ktp", "no_peserta"},
		Order:  "no_rkm_medis desc",
	}}
}

func (r *repository) RefExists(table, column string, value any) (bool, error) {
	return crud.ExistsIn(r.Query(), table, column, value)
}
