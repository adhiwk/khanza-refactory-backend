package pasien

import (
	pasienmodel "goravel/app/models/pasien"

	"github.com/goravel/framework/facades"
)

type Repository interface {
	FindByNoRkmMedis(noRkmMedis string) (*pasienmodel.Pasien, error)
	ExistsByNoRkmMedis(noRkmMedis string) (bool, error)
	GetPaginated(search string, page, limit int) ([]pasienmodel.Pasien, int64, error)
	Create(pasien *pasienmodel.Pasien) error
	Update(pasien *pasienmodel.Pasien) error
	Delete(noRkmMedis string) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

// FindByNoRkmMedis mengembalikan (nil, nil) bila data tidak ditemukan
func (r *repository) FindByNoRkmMedis(noRkmMedis string) (*pasienmodel.Pasien, error) {
	var p pasienmodel.Pasien
	if err := facades.Orm().Query().Where("no_rkm_medis = ?", noRkmMedis).First(&p); err != nil {
		return nil, err
	}
	if p.NoRkmMedis == "" {
		return nil, nil
	}
	return &p, nil
}

func (r *repository) ExistsByNoRkmMedis(noRkmMedis string) (bool, error) {
	count, err := facades.Orm().Query().Model(&pasienmodel.Pasien{}).Where("no_rkm_medis = ?", noRkmMedis).Count()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// GetPaginated mencari di no_rkm_medis / nm_pasien / no_ktp / no_peserta bila search diisi
func (r *repository) GetPaginated(search string, page, limit int) ([]pasienmodel.Pasien, int64, error) {
	var list []pasienmodel.Pasien
	var total int64

	q := facades.Orm().Query().Model(&pasienmodel.Pasien{})
	if search != "" {
		like := "%" + search + "%"
		q = q.Where("(no_rkm_medis LIKE ? OR nm_pasien LIKE ? OR no_ktp LIKE ? OR no_peserta LIKE ?)", like, like, like, like)
	}

	if err := q.Order("no_rkm_medis desc").Paginate(page, limit, &list, &total); err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (r *repository) Create(pasien *pasienmodel.Pasien) error {
	return facades.Orm().Query().Create(pasien)
}

func (r *repository) Update(pasien *pasienmodel.Pasien) error {
	return facades.Orm().Query().Save(pasien)
}

func (r *repository) Delete(noRkmMedis string) error {
	_, err := facades.Orm().Query().Where("no_rkm_medis = ?", noRkmMedis).Delete(&pasienmodel.Pasien{})
	return err
}
