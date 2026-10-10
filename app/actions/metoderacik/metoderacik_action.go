package metoderacik

import (
	"strings"

	request "goravel/app/http/requests/metoderacik"
	model "goravel/app/models/metoderacik"
	repo "goravel/app/repository/metoderacik"
	"goravel/app/support"
)

var (
	ErrNotFound      = support.NotFound("data metode racik tidak ditemukan")
	ErrAlreadyExists = support.Conflict("kode metode racik sudah digunakan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.MetodeRacik, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.MetodeRacik, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.MetodeRacik, error) {
	key := strings.TrimSpace(req.KdRacik)
	exists, err := a.repo.Exists(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}

	data := &model.MetodeRacik{KdRacik: key}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.MetodeRacik, error) {
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

func (a *Action) fill(m *model.MetodeRacik, d request.Data) error {

	m.NmRacik = strings.TrimSpace(d.NmRacik)
	return nil
}
