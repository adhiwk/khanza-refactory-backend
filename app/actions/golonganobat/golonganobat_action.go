package golonganobat

import (
	"strings"

	request "goravel/app/http/requests/golonganobat"
	model "goravel/app/models/golonganobat"
	repo "goravel/app/repository/golonganobat"
	"goravel/app/support"
)

var (
	ErrNotFound      = support.NotFound("data golongan obat tidak ditemukan")
	ErrAlreadyExists = support.Conflict("kode golongan obat sudah digunakan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.GolonganBarang, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.GolonganBarang, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.GolonganBarang, error) {
	key := strings.TrimSpace(req.Kode)
	exists, err := a.repo.Exists(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}

	data := &model.GolonganBarang{Kode: key}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.GolonganBarang, error) {
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

func (a *Action) fill(m *model.GolonganBarang, d request.Data) error {

	m.Nama = support.Nullable(d.Nama)
	return nil
}
