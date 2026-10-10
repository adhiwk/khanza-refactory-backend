package kabupaten

import (
	"strings"

	request "goravel/app/http/requests/kabupaten"
	model "goravel/app/models/kabupaten"
	repo "goravel/app/repository/kabupaten"
	"goravel/app/support"
)

var (
	ErrNotFound = support.NotFound("data kabupaten tidak ditemukan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.Kabupaten, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key int) (*model.Kabupaten, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.Kabupaten, error) {
	data := &model.Kabupaten{}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key int, d request.Data) (*model.Kabupaten, error) {
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

func (a *Action) fill(m *model.Kabupaten, d request.Data) error {

	m.NmKab = strings.TrimSpace(d.NmKab)
	return nil
}
