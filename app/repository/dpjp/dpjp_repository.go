package dpjp

import (
	contractsorm "github.com/goravel/framework/contracts/database/orm"

	model "goravel/app/models/dpjp"
	"goravel/app/repository/crud"
	"goravel/app/support"
)

type Repository interface {
	List(noRawat string) ([]model.Dpjp, error)
	RegistrasiExists(noRawat string) (bool, error)
	DokterExists(kdDokter string) (bool, error)
	Exists(noRawat, kdDokter string) (bool, error)
	Create(tx contractsorm.Query, noRawat, kdDokter string) error
	Delete(noRawat, kdDokter string) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) List(noRawat string) ([]model.Dpjp, error) {
	list := []model.Dpjp{}
	err := support.DB().Table("dpjp_ranap").Select("dpjp_ranap.no_rawat", "dpjp_ranap.kd_dokter", "ifnull(dokter.nm_dokter,'') as nm_dokter").
		Join("left join dokter on dpjp_ranap.kd_dokter=dokter.kd_dokter").Where("dpjp_ranap.no_rawat = ?", noRawat).Scan(&list)
	return list, err
}

func (r *repository) RegistrasiExists(noRawat string) (bool, error) {
	return crud.ExistsIn(support.DB(), "reg_periksa", "no_rawat", noRawat)
}

func (r *repository) DokterExists(kdDokter string) (bool, error) {
	return support.DB().Table("dokter").Where("kd_dokter = ? and status = '1'", kdDokter).Exists()
}

func (r *repository) Exists(noRawat, kdDokter string) (bool, error) {
	return support.DB().Table("dpjp_ranap").Where("no_rawat = ? and kd_dokter = ?", noRawat, kdDokter).Exists()
}

func (r *repository) Create(tx contractsorm.Query, noRawat, kdDokter string) error {
	_, err := tx.Exec("insert into dpjp_ranap (no_rawat,kd_dokter) values (?,?)", noRawat, kdDokter)
	return err
}

func (r *repository) Delete(noRawat, kdDokter string) error {
	_, err := support.DB().Exec("delete from dpjp_ranap where no_rawat=? and kd_dokter=?", noRawat, kdDokter)
	return err
}
