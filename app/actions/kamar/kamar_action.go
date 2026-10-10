package kamar

import (
	"strings"

	request "goravel/app/http/requests/kamar"
	model "goravel/app/models/kamar"
	repo "goravel/app/repository/kamar"
	"goravel/app/support"
)

var (
	ErrNotFound        = support.NotFound("data kamar tidak ditemukan")
	ErrAlreadyExists   = support.Conflict("kode kamar sudah digunakan")
	ErrBangsalNotFound = support.NotFound("data bangsal tidak ditemukan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.Kamar, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.Kamar, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.Kamar, error) {
	key := strings.TrimSpace(req.KdKamar)
	exists, err := a.repo.Exists(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}

	data := &model.Kamar{KdKamar: key, Statusdata: support.StrPtr("1")}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.Kamar, error) {
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

func (a *Action) fill(m *model.Kamar, d request.Data) error {
	if v := strings.TrimSpace(d.KdBangsal); v != "" {
		ok, err := a.repo.BangsalExists(v)
		if err != nil {
			return err
		}
		if !ok {
			return ErrBangsalNotFound
		}
	}
	m.KdBangsal = support.Nullable(d.KdBangsal)
	m.TrfKamar = d.TrfKamar
	m.Status = support.Nullable(d.Status)
	m.Kelas = support.Nullable(d.Kelas)
	return nil
}
