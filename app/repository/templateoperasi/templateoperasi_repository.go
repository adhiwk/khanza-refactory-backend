package templateoperasi

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/templateoperasi"
	"goravel/app/repository/crud"
)

type Repository interface {
	Paginate(search string, page, limit int) ([]model.TemplateLaporanOperasi, int64, error)
	Find(key string) (*model.TemplateLaporanOperasi, error)
	NextKode(tx contractsorm.Query) (string, error)
	Create(tx contractsorm.Query, data *model.TemplateLaporanOperasi) error
	Save(data *model.TemplateLaporanOperasi) error
	Delete(key string) error
}

type repository struct {
	crud.Table[model.TemplateLaporanOperasi, string]
}

func NewRepository() Repository {
	return &repository{crud.Table[model.TemplateLaporanOperasi, string]{
		Key:    "no_template",
		Search: []string{"no_template", "nama_operasi", "diagnosa_preop"},
		Order:  "no_template",
	}}
}

// NextKode "O" + 4 digit (Valid.autoNomer MasterTemplateLaporanOperasi).
func (r *repository) NextKode(tx contractsorm.Query) (string, error) {
	return crud.NextCode(tx, "template_laporan_operasi", "no_template", "O", 4)
}

func (r *repository) Create(tx contractsorm.Query, data *model.TemplateLaporanOperasi) error {
	return r.Table.WithTx(tx).Create(data)
}
