package rekammedis

import "time"

// HasilPemeriksaanUsgUrologi tabel `hasil_pemeriksaan_usg_urologi` (hasil pemeriksaan USG urologi, RMHasilPemeriksaanUSGUrologi).
type HasilPemeriksaanUsgUrologi struct {
	NoRawat        string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal        *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter       string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis string     `gorm:"column:diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari    string     `gorm:"column:kiriman_dari" json:"kiriman_dari"`
	GinjalKanan    *string    `gorm:"column:ginjal_kanan" json:"ginjal_kanan"`
	GinjalKiri     *string    `gorm:"column:ginjal_kiri" json:"ginjal_kiri"`
	VesicaUrinaria *string    `gorm:"column:vesica_urinaria" json:"vesica_urinaria"`
	Tambahan       *string    `gorm:"column:tambahan" json:"tambahan"`
}

func (HasilPemeriksaanUsgUrologi) TableName() string {
	return "hasil_pemeriksaan_usg_urologi"
}
