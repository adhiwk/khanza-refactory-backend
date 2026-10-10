package riwayatpasien

import (
	model "goravel/app/models/riwayatpasien"
	"goravel/app/repository/crud"
	"goravel/app/support"
)

type Repository interface {
	PasienExists(noRkmMedis string) (bool, error)
	ImunisasiExists(kode string) (bool, error)
	Persalinan(noRkmMedis string) ([]model.Persalinan, error)
	PersalinanExists(noRkmMedis, tglThn string) (bool, error)
	CreatePersalinan(p model.Persalinan) error
	DeletePersalinan(noRkmMedis, tglThn string) error
	Imunisasi(noRkmMedis string) ([]model.Imunisasi, error)
	ImunisasiPasienExists(noRkmMedis, kode string, ke int) (bool, error)
	CreateImunisasi(i model.Imunisasi) error
	DeleteImunisasi(noRkmMedis, kode string, ke int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) PasienExists(noRkmMedis string) (bool, error) {
	return crud.ExistsIn(support.DB(), "pasien", "no_rkm_medis", noRkmMedis)
}

func (r *repository) ImunisasiExists(kode string) (bool, error) {
	return crud.ExistsIn(support.DB(), "master_imunisasi", "kode_imunisasi", kode)
}

func (r *repository) Persalinan(noRkmMedis string) ([]model.Persalinan, error) {
	list := []model.Persalinan{}
	err := support.DB().Raw("select * from riwayat_persalinan_pasien where no_rkm_medis=? order by tgl_thn", noRkmMedis).Scan(&list)
	return list, err
}

func (r *repository) PersalinanExists(noRkmMedis, tglThn string) (bool, error) {
	return support.DB().Table("riwayat_persalinan_pasien").Where("no_rkm_medis = ? and tgl_thn = ?", noRkmMedis, tglThn).Exists()
}

func (r *repository) CreatePersalinan(p model.Persalinan) error {
	_, err := support.DB().Exec("insert into riwayat_persalinan_pasien (no_rkm_medis,tgl_thn,tempat_persalinan,usia_hamil,jenis_persalinan,"+
		"penolong,penyulit,jk,bbpb,keadaan) values (?,?,?,?,?,?,?,?,?,?)", p.NoRkmMedis, p.TglThn, p.TempatPersalinan, p.UsiaHamil,
		p.JenisPersalinan, p.Penolong, p.Penyulit, p.Jk, p.Bbpb, p.Keadaan)
	return err
}

// DeletePersalinan tabel tanpa primary key: baris dikenali dari no_rkm_medis + tgl_thn (seperti form Khanza).
func (r *repository) DeletePersalinan(noRkmMedis, tglThn string) error {
	_, err := support.DB().Exec("delete from riwayat_persalinan_pasien where no_rkm_medis=? and tgl_thn=?", noRkmMedis, tglThn)
	return err
}

func (r *repository) Imunisasi(noRkmMedis string) ([]model.Imunisasi, error) {
	list := []model.Imunisasi{}
	err := support.DB().Raw("select r.no_rkm_medis,r.kode_imunisasi,ifnull(m.nama_imunisasi,'') as nama_imunisasi,r.no_imunisasi "+
		"from riwayat_imunisasi r left join master_imunisasi m on r.kode_imunisasi=m.kode_imunisasi where r.no_rkm_medis=? "+
		"order by r.kode_imunisasi,r.no_imunisasi", noRkmMedis).Scan(&list)
	return list, err
}

func (r *repository) ImunisasiPasienExists(noRkmMedis, kode string, ke int) (bool, error) {
	return support.DB().Table("riwayat_imunisasi").Where("no_rkm_medis = ? and kode_imunisasi = ? and no_imunisasi = ?", noRkmMedis, kode, ke).Exists()
}

func (r *repository) CreateImunisasi(i model.Imunisasi) error {
	_, err := support.DB().Exec("insert into riwayat_imunisasi (no_rkm_medis,kode_imunisasi,no_imunisasi) values (?,?,?)", i.NoRkmMedis, i.KodeImunisasi, i.NoImunisasi)
	return err
}

func (r *repository) DeleteImunisasi(noRkmMedis, kode string, ke int) error {
	_, err := support.DB().Exec("delete from riwayat_imunisasi where no_rkm_medis=? and kode_imunisasi=? and no_imunisasi=?", noRkmMedis, kode, ke)
	return err
}
