package dokter

import (
	"strings"

	request "goravel/app/http/requests/dokter"
	model "goravel/app/models/dokter"
	repo "goravel/app/repository/dokter"
	"goravel/app/support"
)

var (
	ErrNotFound          = support.NotFound("data dokter tidak ditemukan")
	ErrAlreadyExists     = support.Conflict("kode dokter sudah digunakan")
	ErrSpesialisNotFound = support.NotFound("data spesialis tidak ditemukan")
	ErrPegawaiNotFound   = support.NotFound("kode dokter belum terdaftar sebagai pegawai")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.Dokter, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.Dokter, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

func (a *Action) Create(req request.StoreRequest) (*model.Dokter, error) {
	key := strings.TrimSpace(req.KdDokter)
	exists, err := a.repo.Exists(key)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, ErrAlreadyExists
	}
	pegawai, err := a.repo.PegawaiExists(key)
	if err != nil {
		return nil, err
	}
	if !pegawai {
		return nil, ErrPegawaiNotFound
	}

	data := &model.Dokter{KdDokter: key, Status: "1"}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	if err := a.repo.Create(data); err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.Dokter, error) {
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

func (a *Action) fill(m *model.Dokter, d request.Data) error {
	if v := strings.TrimSpace(d.KdSps); v != "" {
		ok, err := a.repo.SpesialisExists(v)
		if err != nil {
			return err
		}
		if !ok {
			return ErrSpesialisNotFound
		}
	}
	tgllahir, err := support.ParseDate(d.TglLahir)
	if err != nil {
		return err
	}
	m.NmDokter = support.Nullable(d.NmDokter)
	m.Jk = support.Nullable(d.Jk)
	m.TmpLahir = support.Nullable(d.TmpLahir)
	m.TglLahir = tgllahir
	m.GolDrh = support.Nullable(d.GolDrh)
	m.Agama = support.Nullable(d.Agama)
	m.AlmtTgl = support.Nullable(d.AlmtTgl)
	m.NoTelp = support.Nullable(d.NoTelp)
	m.Email = strings.TrimSpace(d.Email)
	m.SttsNikah = support.Nullable(d.SttsNikah)
	m.KdSps = support.Nullable(d.KdSps)
	m.Alumni = support.Nullable(d.Alumni)
	m.NoIjnPraktek = support.Nullable(d.NoIjnPraktek)
	return nil
}
