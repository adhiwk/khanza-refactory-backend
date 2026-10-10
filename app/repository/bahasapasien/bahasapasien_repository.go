package bahasapasien

import (
	model "goravel/app/models/bahasapasien"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.BahasaPasien, int64, error)
	Find(key int) (*model.BahasaPasien, error)
	Exists(key int) (bool, error)
	Create(data *model.BahasaPasien) error
	Save(data *model.BahasaPasien) error
	Delete(key int) error
}

type repository struct {
	crud.Table[model.BahasaPasien, int]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.BahasaPasien, int]{
		Key:    "id",
		Search: []string{"nama_bahasa"},
		Active: "",
		Order:  "nama_bahasa",
	}}
}
