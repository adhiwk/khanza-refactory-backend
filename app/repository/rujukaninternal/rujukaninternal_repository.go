package rujukaninternal

import (
	model "goravel/app/models/rujukaninternal"
	"goravel/app/repository/crud"
	"goravel/app/support"
)

type Repository interface {
	List(noRawat string) ([]model.RujukanInternal, error)
	RegistrasiExists(noRawat string) (bool, error)
	DokterExists(kdDokter string) (bool, error)
	PoliExists(kdPoli string) (bool, error)
	Exists(noRawat, kdDokter string) (bool, error)
	Create(data model.RujukanInternal) error
	Delete(noRawat, kdDokter string) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) List(noRawat string) ([]model.RujukanInternal, error) {
	list := []model.RujukanInternal{}
	err := support.DB().Table("rujukan_internal_poli").
		Select("rujukan_internal_poli.no_rawat", "rujukan_internal_poli.kd_dokter", "ifnull(dokter.nm_dokter,'') as nm_dokter",
			"ifnull(rujukan_internal_poli.kd_poli,'') as kd_poli", "ifnull(poliklinik.nm_poli,'') as nm_poli").
		Join("left join dokter on rujukan_internal_poli.kd_dokter=dokter.kd_dokter").
		Join("left join poliklinik on rujukan_internal_poli.kd_poli=poliklinik.kd_poli").
		Where("rujukan_internal_poli.no_rawat = ?", noRawat).Scan(&list)
	return list, err
}

func (r *repository) RegistrasiExists(noRawat string) (bool, error) {
	return crud.ExistsIn(support.DB(), "reg_periksa", "no_rawat", noRawat)
}

func (r *repository) DokterExists(kdDokter string) (bool, error) {
	return support.DB().Table("dokter").Where("kd_dokter = ? and status = '1'", kdDokter).Exists()
}

func (r *repository) PoliExists(kdPoli string) (bool, error) {
	return support.DB().Table("poliklinik").Where("kd_poli = ? and status = '1'", kdPoli).Exists()
}

func (r *repository) Exists(noRawat, kdDokter string) (bool, error) {
	return support.DB().Table("rujukan_internal_poli").Where("no_rawat = ? and kd_dokter = ?", noRawat, kdDokter).Exists()
}

func (r *repository) Create(d model.RujukanInternal) error {
	_, err := support.DB().Exec("insert into rujukan_internal_poli (no_rawat,kd_dokter,kd_poli) values (?,?,?)", d.NoRawat, d.KdDokter, d.KdPoli)
	return err
}

func (r *repository) Delete(noRawat, kdDokter string) error {
	_, err := support.DB().Exec("delete from rujukan_internal_poli where no_rawat=? and kd_dokter=?", noRawat, kdDokter)
	return err
}
