// Package templateradiologi use case template hasil radiologi (MasterTemplateHasilRadiologi).
package templateradiologi

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"

	request "goravel/app/http/requests/templateradiologi"
	model "goravel/app/models/templateradiologi"
	repo "goravel/app/repository/templateradiologi"
	"goravel/app/support"
)

var ErrNotFound = support.NotFound("data template hasil radiologi tidak ditemukan")

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.TemplateHasilRadiologi, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.TemplateHasilRadiologi, error) {
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
func (a *Action) Create(d request.Data) (*model.TemplateHasilRadiologi, error) {
	data := &model.TemplateHasilRadiologi{}
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

func (a *Action) Update(key string, d request.Data) (*model.TemplateHasilRadiologi, error) {
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

func fill(m *model.TemplateHasilRadiologi, d request.Data) {
	m.NamaPemeriksaan = support.Nullable(d.NamaPemeriksaan)
	m.TemplateHasilRadiologi = support.Nullable(d.TemplateHasilRadiologi)
}
