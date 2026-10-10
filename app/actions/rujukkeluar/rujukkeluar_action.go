package rujukkeluar

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"

	"strings"

	request "goravel/app/http/requests/rujukkeluar"
	model "goravel/app/models/rujukkeluar"
	repo "goravel/app/repository/rujukkeluar"
	"goravel/app/support"
)

var (
	ErrNotFound           = support.NotFound("data rujukan keluar tidak ditemukan")
	ErrAlreadyExists      = support.Conflict("kode rujukan keluar sudah digunakan")
	ErrRegPeriksaNotFound = support.NotFound("data registrasi tidak ditemukan")
	ErrDokterNotFound     = support.NotFound("data dokter tidak ditemukan")
)

type Action struct {
	repo repo.Repository
}

func NewAction(repo repo.Repository) *Action {
	return &Action{repo: repo}
}

func (a *Action) List(search string, page, limit int) ([]model.Rujuk, int64, error) {
	return a.repo.Paginate(search, page, limit)
}

func (a *Action) Detail(key string) (*model.Rujuk, error) {
	data, err := a.repo.Find(key)
	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, ErrNotFound
	}
	return data, nil
}

// Create nomor rujukan dibuat otomatis dan disimpan dalam satu transaksi agar tidak bentrok.
func (a *Action) Create(req request.StoreRequest) (*model.Rujuk, error) {
	data := &model.Rujuk{}
	if err := a.fill(data, req.Data); err != nil {
		return nil, err
	}
	err := support.Transaction(func(tx contractsorm.Query) error {
		r := a.repo.WithTx(tx)
		no, err := r.NextNoRujuk()
		if err != nil {
			return err
		}
		data.NoRujuk = no
		return r.Create(data)
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (a *Action) Update(key string, d request.Data) (*model.Rujuk, error) {
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

func (a *Action) fill(m *model.Rujuk, d request.Data) error {
	if v := strings.TrimSpace(d.NoRawat); v != "" {
		ok, err := a.repo.RegPeriksaExists(v)
		if err != nil {
			return err
		}
		if !ok {
			return ErrRegPeriksaNotFound
		}
	}
	if v := strings.TrimSpace(d.KdDokter); v != "" {
		ok, err := a.repo.DokterExists(v)
		if err != nil {
			return err
		}
		if !ok {
			return ErrDokterNotFound
		}
	}
	tglrujuk, err := support.ParseDate(d.TglRujuk)
	if err != nil {
		return err
	}
	m.NoRawat = support.Nullable(d.NoRawat)
	m.RujukKe = support.Nullable(d.RujukKe)
	m.TglRujuk = tglrujuk
	m.KeteranganDiagnosa = support.Nullable(d.KeteranganDiagnosa)
	m.KdDokter = support.Nullable(d.KdDokter)
	m.KatRujuk = support.Nullable(d.KatRujuk)
	m.Ambulance = support.Nullable(d.Ambulance)
	m.Keterangan = support.Nullable(d.Keterangan)
	m.Jam = support.Nullable(d.Jam)
	return nil
}
