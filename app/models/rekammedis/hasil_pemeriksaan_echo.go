package rekammedis

import "time"

// HasilPemeriksaanEcho tabel `hasil_pemeriksaan_echo` (hasil pemeriksaan echo, RMHasilPemeriksaanEcho).
type HasilPemeriksaanEcho struct {
	NoRawat          string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal          *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter         string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Sistolik         *string    `gorm:"column:sistolik" json:"sistolik"`
	Diastolic        *string    `gorm:"column:diastolic" json:"diastolic"`
	Kontraktilitas   *string    `gorm:"column:kontraktilitas" json:"kontraktilitas"`
	DimensiRuang     *string    `gorm:"column:dimensi_ruang" json:"dimensi_ruang"`
	Katup            *string    `gorm:"column:katup" json:"katup"`
	AnalisaSegmental *string    `gorm:"column:analisa_segmental" json:"analisa_segmental"`
	Erap             *string    `gorm:"column:erap" json:"erap"`
	LainLain         *string    `gorm:"column:lain_lain" json:"lain_lain"`
	Kesimpulan       *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
}

func (HasilPemeriksaanEcho) TableName() string {
	return "hasil_pemeriksaan_echo"
}
