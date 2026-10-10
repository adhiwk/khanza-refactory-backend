package catatanpasien

import (
	"strings"

	request "goravel/app/http/requests/catatanpasien"
	model "goravel/app/models/catatanpasien"
	repo "goravel/app/repository/catatanpasien"
	"goravel/app/support"
)

var (
	ErrNotFound       = support.NotFound("data catatan pasien tidak ditemukan")
	ErrAlreadyExists  = support.Conflict("kode catatan pasien sudah digunakan")
	ErrPasienNotFound = support.NotFound("data pasien tidak ditemukan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.CatatanPasien, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.CatatanPasien, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.CatatanPasien, error) {
	key := strings.TrimSpace(req.NoRkmMedis)
	exists, err := a.repo.Exists(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}
	parent, err := a.repo.PasienExists(key)
	if err != nil {
		return nil, err
	}
	if !parent {
		return nil, ErrPasienNotFound
	}

	data := &model.CatatanPasien{NoRkmMedis: key}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.CatatanPasien, error) {
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
func (a *Action) Delete(key string) error {
	if _, err := a.Detail(key); err != nil {
		return err
	}
	return a.repo.Delete(key)
}

func (a *Action) fill(m *model.CatatanPasien, d request.Data) error {

	m.Catatan = support.Nullable(d.Catatan)
	return nil
}
