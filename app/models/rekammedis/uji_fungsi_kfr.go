package rekammedis

import "time"

// UjiFungsiKfr tabel `uji_fungsi_kfr` (uji fungsi KFR, RMUjiFungsiKFR).
type UjiFungsiKfr struct {
	NoRawat             string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal             *time.Time `gorm:"column:tanggal" json:"tanggal"`
	DiagnosisFungsional *string    `gorm:"column:diagnosis_fungsional" json:"diagnosis_fungsional"`
	DiagnosisMedis      *string    `gorm:"column:diagnosis_medis" json:"diagnosis_medis"`
	HasilDidapat        *string    `gorm:"column:hasil_didapat" json:"hasil_didapat"`
	Kesimpulan          *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
	Rekomedasi          *string    `gorm:"column:rekomedasi" json:"rekomedasi"`
	KdDokter            *string    `gorm:"column:kd_dokter" json:"kd_dokter"`
}

func (UjiFungsiKfr) TableName() string {
	return "uji_fungsi_kfr"
}
