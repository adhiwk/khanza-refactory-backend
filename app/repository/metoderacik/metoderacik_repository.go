package metoderacik

import (
	model "goravel/app/models/metoderacik"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.MetodeRacik, int64, error)
	Find(key string) (*model.MetodeRacik, error)
	Exists(key string) (bool, error)
	Create(data *model.MetodeRacik) error
	Save(data *model.MetodeRacik) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.MetodeRacik, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.MetodeRacik, string]{
		Key:    "kd_racik",
		Search: []string{"kd_racik", "nm_racik"},
		Active: "",
		Order:  "kd_racik",
	}}
}
