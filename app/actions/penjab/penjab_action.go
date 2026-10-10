package penjab

import (
	"strings"

	request "goravel/app/http/requests/penjab"
	model "goravel/app/models/penjab"
	repo "goravel/app/repository/penjab"
	"goravel/app/support"
)

var (
	ErrNotFound      = support.NotFound("data penanggung jawab tidak ditemukan")
	ErrAlreadyExists = support.Conflict("kode penanggung jawab sudah digunakan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.Penjab, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.Penjab, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.Penjab, error) {
	key := strings.TrimSpace(req.KdPj)
	exists, err := a.repo.Exists(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}

	data := &model.Penjab{KdPj: key, Status: "1"}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.Penjab, error) {
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

func (a *Action) fill(m *model.Penjab, d request.Data) error {

	m.PngJawab = strings.TrimSpace(d.PngJawab)
	m.NamaPerusahaan = strings.TrimSpace(d.NamaPerusahaan)
	m.AlamatAsuransi = strings.TrimSpace(d.AlamatAsuransi)
	m.NoTelp = strings.TrimSpace(d.NoTelp)
	m.Attn = strings.TrimSpace(d.Attn)
	return nil
}
