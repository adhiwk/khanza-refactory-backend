package industrifarmasi

import (
	"strings"

	request "goravel/app/http/requests/industrifarmasi"
	model "goravel/app/models/industrifarmasi"
	repo "goravel/app/repository/industrifarmasi"
	"goravel/app/support"
)

var (
	ErrNotFound      = support.NotFound("data industri farmasi tidak ditemukan")
	ErrAlreadyExists = support.Conflict("kode industri farmasi sudah digunakan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.Industrifarmasi, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.Industrifarmasi, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.Industrifarmasi, error) {
	key := strings.TrimSpace(req.KodeIndustri)
	exists, err := a.repo.Exists(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}

	data := &model.Industrifarmasi{KodeIndustri: key}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.Industrifarmasi, error) {
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

func (a *Action) fill(m *model.Industrifarmasi, d request.Data) error {

	m.NamaIndustri = support.Nullable(d.NamaIndustri)
	m.Alamat = support.Nullable(d.Alamat)
	m.Kota = support.Nullable(d.Kota)
	m.NoTelp = support.Nullable(d.NoTelp)
	return nil
}
