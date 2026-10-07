package obat

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"
	"github.com/goravel/framework/facades"

	obatModel "goravel/app/models/obat"
)

const connection = "mysql_kedua"

// Filter untuk list / pencarian obat
type Filter struct {
	Search       string // cari di kode_brng / nama_brng
	Status       string // "0" / "1" / "" (semua)
	Kdjns        string
	KodeKategori string
	KodeGolongan string
	Page         int
	Limit        int
}

type Repository interface {
	Paginate(filter Filter) ([]obatModel.Obat, int64, error)
	FindByKode(kode string) (*obatModel.Obat, error)
	ExistsByKode(kode string) (bool, error)
	Create(data *obatModel.Obat) error
	Save(data *obatModel.Obat) error
	Delete(kode string) (int64, error)
	UpdateStatus(kode string, status string) (int64, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) query() contractsorm.Query {
	return facades.Orm().Connection(connection).Query()
}

func (r *repository) Paginate(filter Filter) ([]obatModel.Obat, int64, error) {
	var (
		list  []obatModel.Obat
		total int64
	)

	q := r.query().Model(&obatModel.Obat{})

	if filter.Search != "" {
		like := "%" + filter.Search + "%"
		q = q.Where("(kode_brng LIKE ? OR nama_brng LIKE ?)", like, like)
	}
	if filter.Status != "" {
		q = q.Where("status = ?", filter.Status)
	}
	if filter.Kdjns != "" {
		q = q.Where("kdjns = ?", filter.Kdjns)
	}
	if filter.KodeKategori != "" {
		q = q.Where("kode_kategori = ?", filter.KodeKategori)
	}
	if filter.KodeGolongan != "" {
		q = q.Where("kode_golongan = ?", filter.KodeGolongan)
	}

	if err := q.Order("nama_brng asc").Paginate(filter.Page, filter.Limit, &list, &total); err != nil {
		return nil, 0, err
	}

	return list, total, nil
}

// FindByKode mengembalikan (nil, nil) bila data tidak ditemukan
func (r *repository) FindByKode(kode string) (*obatModel.Obat, error) {
	var data obatModel.Obat
	if err := r.query().Where("kode_brng = ?", kode).First(&data); err != nil {
		return nil, err
	}
	if data.KodeBrng == "" {
		return nil, nil
	}
	return &data, nil
}

func (r *repository) ExistsByKode(kode string) (bool, error) {
	count, err := r.query().Model(&obatModel.Obat{}).Where("kode_brng = ?", kode).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *repository) Create(data *obatModel.Obat) error {
	return r.query().Create(data)
}

// Save meng-update semua kolom (termasuk nilai 0 / string kosong)
func (r *repository) Save(data *obatModel.Obat) error {
	return r.query().Save(data)
}

func (r *repository) Delete(kode string) (int64, error) {
	res, err := r.query().Where("kode_brng = ?", kode).Delete(&obatModel.Obat{})
	if err != nil {
		return 0, err
	}
	return res.RowsAffected, nil
}

func (r *repository) UpdateStatus(kode string, status string) (int64, error) {
	res, err := r.query().Model(&obatModel.Obat{}).
		Where("kode_brng = ?", kode).
		Update("status", status)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected, nil
}
