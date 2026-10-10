package templateedukasi

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/templateedukasi"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.TemplatePelaksanaanInformasiEdukasi, int64, error)
	Find(key string) (*model.TemplatePelaksanaanInformasiEdukasi, error)
	NextKode(tx contractsorm.Query) (string, error)
	Create(tx contractsorm.Query, data *model.TemplatePelaksanaanInformasiEdukasi) error
	Save(data *model.TemplatePelaksanaanInformasiEdukasi) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.TemplatePelaksanaanInformasiEdukasi, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.TemplatePelaksanaanInformasiEdukasi, string]{
		Key:    "no_template",
		Search: []string{"no_template", "materi_edukasi"},
		Order:  "no_template",
	}}
}

// NextKode "E" + 3 digit (Valid.autoNomer MasterTemplateInformasiEdukasi).
func (r *repository) NextKode(tx contractsorm.Query) (string, error) {
	return crud.NextCode(tx, "template_pelaksanaan_informasi_edukasi", "no_template", "E", 3)
}

func (r *repository) Create(tx contractsorm.Query, data *model.TemplatePelaksanaanInformasiEdukasi) error {
	return r.Table.WithTx(tx).Create(data)
}
