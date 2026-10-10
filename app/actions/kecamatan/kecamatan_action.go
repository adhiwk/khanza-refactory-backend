package kecamatan

import (
	"strings"

	request "goravel/app/http/requests/kecamatan"
	model "goravel/app/models/kecamatan"
	repo "goravel/app/repository/kecamatan"
	"goravel/app/support"
)

var (
	ErrNotFound = support.NotFound("data kecamatan tidak ditemukan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.Kecamatan, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key int) (*model.Kecamatan, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.Kecamatan, error) {
	data := &model.Kecamatan{}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key int, d request.Data) (*model.Kecamatan, error) {
	data, err := a.Detail(key)
	if err != nil {
		return nil, err
	}
	if err := a.fill(data, d); err != nil {
		return nil, err
	}
	if err := a.repo.Save(data); err != nil {
		return nil, err
	}
	return data, nil
}

// Delete menghapus sesuai form Khanza.
func (a *Action) Delete(key int) error {
	if _, err := a.Detail(key); err != nil {
		return err
	}
	return a.repo.Delete(key)
}

func (a *Action) fill(m *model.Kecamatan, d request.Data) error {

	m.NmKec = strings.TrimSpace(d.NmKec)
	return nil
}
