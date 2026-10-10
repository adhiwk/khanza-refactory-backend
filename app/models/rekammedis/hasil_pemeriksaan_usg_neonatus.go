package rekammedis

import "time"

// HasilPemeriksaanUsgNeonatus tabel `hasil_pemeriksaan_usg_neonatus` (hasil pemeriksaan USG neonatus, RMHasilPemeriksaanUSGNeonatus).
type HasilPemeriksaanUsgNeonatus struct {
	NoRawat           string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal           *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter          string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis    *string    `gorm:"column:diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari       string     `gorm:"column:kiriman_dari" json:"kiriman_dari"`
	VentrikalSinistra *string    `gorm:"column:ventrikal_sinistra" json:"ventrikal_sinistra"`
	VentrikalDextra   *string    `gorm:"column:ventrikal_dextra" json:"ventrikal_dextra"`
	Kesan             *string    `gorm:"column:kesan" json:"kesan"`
	Kesimpulan        *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
	Saran             *string    `gorm:"column:saran" json:"saran"`
}

func (HasilPemeriksaanUsgNeonatus) TableName() string {
	return "hasil_pemeriksaan_usg_neonatus"
}
