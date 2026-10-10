package kategoriobat

import (
	"strings"

	request "goravel/app/http/requests/kategoriobat"
	model "goravel/app/models/kategoriobat"
	repo "goravel/app/repository/kategoriobat"
	"goravel/app/support"
)

var (
	ErrNotFound      = support.NotFound("data kategori obat tidak ditemukan")
	ErrAlreadyExists = support.Conflict("kode kategori obat sudah digunakan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.KategoriBarang, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.KategoriBarang, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.KategoriBarang, error) {
	key := strings.TrimSpace(req.Kode)
	exists, err := a.repo.Exists(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}

	data := &model.KategoriBarang{Kode: key}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.KategoriBarang, error) {
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

func (a *Action) fill(m *model.KategoriBarang, d request.Data) error {

	m.Nama = support.Nullable(d.Nama)
	return nil
}
