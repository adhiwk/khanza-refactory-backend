package pasienmati

import (
	"strings"

	request "goravel/app/http/requests/pasienmati"
	model "goravel/app/models/pasienmati"
	repo "goravel/app/repository/pasienmati"
	"goravel/app/support"
)

var (
	ErrNotFound       = support.NotFound("data pasien meninggal tidak ditemukan")
	ErrAlreadyExists  = support.Conflict("kode pasien meninggal sudah digunakan")
	ErrDokterNotFound = support.NotFound("data dokter tidak ditemukan")
	ErrPasienNotFound = support.NotFound("data pasien tidak ditemukan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.PasienMati, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.PasienMati, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.PasienMati, error) {
	key := strings.TrimSpace(req.NoRkmMedis)
	exists, err := a.repo.Exists(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}
	parent, err := a.repo.PasienExists(key)
	if err != nil {
		return nil, err
	}
	if !parent {
		return nil, ErrPasienNotFound
	}

	data := &model.PasienMati{NoRkmMedis: key}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.PasienMati, error) {
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

func (a *Action) fill(m *model.PasienMati, d request.Data) error {
	if v := strings.TrimSpace(d.KdDokter); v != "" {
		ok, err := a.repo.DokterExists(v)
		if err != nil {
			return err
		}
		if !ok {
			return ErrDokterNotFound
		}
	}
	tanggal, err := support.ParseDate(d.Tanggal)
	if err != nil {
		return err
	}
	m.Tanggal = tanggal
	m.Jam = support.Nullable(d.Jam)
	m.Keterangan = support.Nullable(d.Keterangan)
	m.TempMeninggal = support.Nullable(d.TempMeninggal)
	m.Icd1 = support.Nullable(d.Icd1)
	m.Icd2 = support.Nullable(d.Icd2)
	m.Icd3 = support.Nullable(d.Icd3)
	m.Icd4 = support.Nullable(d.Icd4)
	m.KdDokter = strings.TrimSpace(d.KdDokter)
	return nil
}
