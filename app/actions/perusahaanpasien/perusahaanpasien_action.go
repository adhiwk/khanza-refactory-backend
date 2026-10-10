package perusahaanpasien

import (
	"strings"

	request "goravel/app/http/requests/perusahaanpasien"
	model "goravel/app/models/perusahaanpasien"
	repo "goravel/app/repository/perusahaanpasien"
	"goravel/app/support"
)

var (
	ErrNotFound      = support.NotFound("data instansi/perusahaan pasien tidak ditemukan")
	ErrAlreadyExists = support.Conflict("kode instansi/perusahaan pasien sudah digunakan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.PerusahaanPasien, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.PerusahaanPasien, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.PerusahaanPasien, error) {
	key := strings.TrimSpace(req.KodePerusahaan)
	exists, err := a.repo.Exists(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}

	data := &model.PerusahaanPasien{KodePerusahaan: key}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.PerusahaanPasien, error) {
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

func (a *Action) fill(m *model.PerusahaanPasien, d request.Data) error {

	m.NamaPerusahaan = support.Nullable(d.NamaPerusahaan)
	m.Alamat = support.Nullable(d.Alamat)
	m.Kota = support.Nullable(d.Kota)
	m.NoTelp = support.Nullable(d.NoTelp)
	return nil
}
