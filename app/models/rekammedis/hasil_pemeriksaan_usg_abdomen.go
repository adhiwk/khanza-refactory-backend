package rekammedis

import "time"

// HasilPemeriksaanUsgAbdomen tabel `hasil_pemeriksaan_usg_abdomen` (hasil pemeriksaan USG abdomen, RMHasilPemeriksaanUSGAbdomen).
type HasilPemeriksaanUsgAbdomen struct {
	NoRawat        string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal        *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter       string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis string     `gorm:"column:diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari    string     `gorm:"column:kiriman_dari" json:"kiriman_dari"`
	Esofagus       *string    `gorm:"column:esofagus" json:"esofagus"`
	Colon          *string    `gorm:"column:colon" json:"colon"`
	Gaster         *string    `gorm:"column:gaster" json:"gaster"`
	Hepar          *string    `gorm:"column:hepar" json:"hepar"`
	GallBlader     *string    `gorm:"column:gall_blader" json:"gall_blader"`
	Lien           *string    `gorm:"column:lien" json:"lien"`
	Pancreas       *string    `gorm:"column:pancreas" json:"pancreas"`
	GinjalDextra   *string    `gorm:"column:ginjal_dextra" json:"ginjal_dextra"`
	GinjalSinistra *string    `gorm:"column:ginjal_sinistra" json:"ginjal_sinistra"`
	Kesimpulan     *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
}

func (HasilPemeriksaanUsgAbdomen) TableName() string {
	return "hasil_pemeriksaan_usg_abdomen"
}
