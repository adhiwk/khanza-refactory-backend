package templateradiologi

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/templateradiologi"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.TemplateHasilRadiologi, int64, error)
	Find(key string) (*model.TemplateHasilRadiologi, error)
	NextKode(tx contractsorm.Query) (string, error)
	Create(tx contractsorm.Query, data *model.TemplateHasilRadiologi) error
	Save(data *model.TemplateHasilRadiologi) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.TemplateHasilRadiologi, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.TemplateHasilRadiologi, string]{
		Key:    "no_template",
		Search: []string{"no_template", "nama_pemeriksaan"},
		Order:  "no_template",
	}}
}

// NextKode "R" + 4 digit (Valid.autoNomer MasterTemplateHasilRadiologi).
func (r *repository) NextKode(tx contractsorm.Query) (string, error) {
	return crud.NextCode(tx, "template_hasil_radiologi", "no_template", "R", 4)
}

func (r *repository) Create(tx contractsorm.Query, data *model.TemplateHasilRadiologi) error {
	return r.Table.WithTx(tx).Create(data)
}
