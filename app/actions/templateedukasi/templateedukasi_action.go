// Package templateedukasi use case template informasi edukasi (MasterTemplateInformasiEdukasi).
package templateedukasi

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"

	request "goravel/app/http/requests/templateedukasi"
	model "goravel/app/models/templateedukasi"
	repo "goravel/app/repository/templateedukasi"
	"goravel/app/support"
)

var ErrNotFound = support.NotFound("data template informasi edukasi tidak ditemukan")

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.TemplatePelaksanaanInformasiEdukasi, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.TemplatePelaksanaanInformasiEdukasi, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

// Create nomor template dibuat otomatis di dalam transaksi.
func (a *Action) Create(d request.Data) (*model.TemplatePelaksanaanInformasiEdukasi, error) {
	data := &model.TemplatePelaksanaanInformasiEdukasi{}
	fill(data, d)
	err := support.Transaction(func(tx contractsorm.Query) error {
		no, err := a.repo.NextKode(tx)
		if err != nil {
			return err
		}
		data.NoTemplate = no
		return a.repo.Create(tx, data)
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.TemplatePelaksanaanInformasiEdukasi, error) {
	data, err := a.Detail(key)
	if err != nil {
		return nil, err
	}
	fill(data, d)
	if err := a.repo.Save(data); err != nil {
		return nil, err
	}
	return data, nil
}

// Delete hapus permanen (Valid.hapusTabletf).
func (a *Action) Delete(key string) error {
	if _, err := a.Detail(key); err != nil {
		return err
	}
	return a.repo.Delete(key)
}

func fill(m *model.TemplatePelaksanaanInformasiEdukasi, d request.Data) {
	m.MateriEdukasi = support.Nullable(d.MateriEdukasi)
	m.LamaEdukasi = support.Nullable(d.LamaEdukasi)
	m.MetodeEdukasi = support.Nullable(d.MetodeEdukasi)
}
