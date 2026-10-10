package poliklinik

import (
	"strings"

	request "goravel/app/http/requests/poliklinik"
	model "goravel/app/models/poliklinik"
	repo "goravel/app/repository/poliklinik"
	"goravel/app/support"
)

var (
	ErrNotFound      = support.NotFound("data poliklinik tidak ditemukan")
	ErrAlreadyExists = support.Conflict("kode poliklinik sudah digunakan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.Poliklinik, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.Poliklinik, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.Poliklinik, error) {
	key := strings.TrimSpace(req.KdPoli)
	exists, err := a.repo.Exists(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}

	data := &model.Poliklinik{KdPoli: key, Status: "1"}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.Poliklinik, error) {
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

// Delete menonaktifkan (status=0) sesuai form Khanza.
func (a *Action) Delete(key string) error {
	if _, err := a.Detail(key); err != nil {
		return err
	}
	return a.repo.Delete(key)
}

func (a *Action) fill(m *model.Poliklinik, d request.Data) error {

	m.NmPoli = support.Nullable(d.NmPoli)
	m.Registrasi = d.Registrasi
	m.Registrasilama = d.Registrasilama
	return nil
}
