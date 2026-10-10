package rekammedis

import "time"

// LaporanTindakan tabel `laporan_tindakan` (laporan tindakan, RMLaporanTindakan).
type LaporanTindakan struct {
	NoRawat               string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal               *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	KdDokter              string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Nip                   *string    `gorm:"column:nip" json:"nip"`
	DiagnosaPraTindakan   string     `gorm:"column:diagnosa_pra_tindakan" json:"diagnosa_pra_tindakan"`
	DiagnosaPascaTindakan string     `gorm:"column:diagnosa_pasca_tindakan" json:"diagnosa_pasca_tindakan"`
	TindakanMedik         string     `gorm:"column:tindakan_medik" json:"tindakan_medik"`
	Uraian                string     `gorm:"column:uraian" json:"uraian"`
	Hasil                 string     `gorm:"column:hasil" json:"hasil"`
	Kesimpulan            string     `gorm:"column:kesimpulan" json:"kesimpulan"`
}

func (LaporanTindakan) TableName() string {
	return "laporan_tindakan"
}
