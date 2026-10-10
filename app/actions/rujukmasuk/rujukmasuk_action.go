package rujukmasuk

import (
	"strings"

	request "goravel/app/http/requests/rujukmasuk"
	model "goravel/app/models/rujukmasuk"
	repo "goravel/app/repository/rujukmasuk"
	"goravel/app/support"
)

var (
	ErrNotFound           = support.NotFound("data rujukan masuk tidak ditemukan")
	ErrAlreadyExists      = support.Conflict("kode rujukan masuk sudah digunakan")
	ErrPenyakitNotFound   = support.NotFound("data penyakit tidak ditemukan")
	ErrRegPeriksaNotFound = support.NotFound("data registrasi tidak ditemukan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.RujukMasuk, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.RujukMasuk, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.RujukMasuk, error) {
	key := strings.TrimSpace(req.NoRawat)
	exists, err := a.repo.Exists(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}
	parent, err := a.repo.RegPeriksaExists(key)
	if err != nil {
		return nil, err
	}
	if !parent {
		return nil, ErrRegPeriksaNotFound
	}

	data := &model.RujukMasuk{NoRawat: key}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	noBalasan, err := a.repo.NextNoBalasan(key)
	if err != nil {
		return nil, err
	}
	data.NoBalasan = &noBalasan
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.RujukMasuk, error) {
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

func (a *Action) fill(m *model.RujukMasuk, d request.Data) error {
	if v := strings.TrimSpace(d.KdPenyakit); v != "" {
		ok, err := a.repo.PenyakitExists(v)
		if err != nil {
			return err
		}
		if !ok {
			return ErrPenyakitNotFound
		}
	}
	m.Perujuk = support.Nullable(d.Perujuk)
	m.Alamat = strings.TrimSpace(d.Alamat)
	m.NoRujuk = strings.TrimSpace(d.NoRujuk)
	m.JmPerujuk = d.JmPerujuk
	m.DokterPerujuk = support.Nullable(d.DokterPerujuk)
	m.KdPenyakit = support.Nullable(d.KdPenyakit)
	m.KategoriRujuk = support.Nullable(d.KategoriRujuk)
	m.Keterangan = support.Nullable(d.Keterangan)
	return nil
}
