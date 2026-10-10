// Package perawatan repository bersama untuk modul pelayanan per no_rawat.
package perawatan

import (
	model "goravel/app/models/perawatan"
	"goravel/app/support"
)

type KonteksRepository interface {
	Konteks(noRawat string) (*model.Konteks, error)
	PegawaiExists(nik string) (bool, error)
}

type konteksRepository struct{}

func NewKonteksRepository() KonteksRepository {
	return &konteksRepository{}
}

// Konteks registrasi + status billing + kamar inap aktif terakhir; (nil, nil) bila no_rawat tidak ada.
func (r *konteksRepository) Konteks(noRawat string) (*model.Konteks, error) {
	list := []model.Konteks{}
	err := support.DB().Raw("select reg_periksa.no_rawat,reg_periksa.no_rkm_medis,pasien.nm_pasien,reg_periksa.tgl_registrasi,"+
		"reg_periksa.jam_reg,reg_periksa.kd_poli,reg_periksa.kd_pj,ifnull(reg_periksa.stts,'') as stts,reg_periksa.status_lanjut,"+
		"exists(select 1 from billing where billing.no_rawat=reg_periksa.no_rawat) as billing,"+
		"ifnull((select kamar.kd_bangsal from kamar_inap inner join kamar on kamar_inap.kd_kamar=kamar.kd_kamar "+
		"where kamar_inap.no_rawat=reg_periksa.no_rawat order by kamar_inap.tgl_masuk desc,kamar_inap.jam_masuk desc limit 1),'') as kd_bangsal,"+
		"ifnull((select kamar.kelas from kamar_inap inner join kamar on kamar_inap.kd_kamar=kamar.kd_kamar "+
		"where kamar_inap.no_rawat=reg_periksa.no_rawat order by kamar_inap.tgl_masuk desc,kamar_inap.jam_masuk desc limit 1),'') as kelas "+
		"from reg_periksa inner join pasien on reg_periksa.no_rkm_medis=pasien.no_rkm_medis where reg_periksa.no_rawat=?", noRawat).Scan(&list)
	if err != nil || len(list) == 0 {
		return nil, err
	}
	return &list[0], nil
}

func (r *konteksRepository) PegawaiExists(nik string) (bool, error) {
	return support.DB().Table("pegawai").Where("nik = ?", nik).Exists()
}
