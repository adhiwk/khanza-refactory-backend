package rekammedis

import "time"

// HasilPemeriksaanUsgGynecologi tabel `hasil_pemeriksaan_usg_gynecologi` (hasil pemeriksaan USG gynecologi, RMHasilPemeriksaanUSGGynecologi).
type HasilPemeriksaanUsgGynecologi struct {
	NoRawat        string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal        *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter       string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis string     `gorm:"column:diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari    string     `gorm:"column:kiriman_dari" json:"kiriman_dari"`
	Uterus         *string    `gorm:"column:uterus" json:"uterus"`
	Parametrium    *string    `gorm:"column:parametrium" json:"parametrium"`
	Ovarium        *string    `gorm:"column:ovarium" json:"ovarium"`
	Doppler        *string    `gorm:"column:doppler" json:"doppler"`
	Kesimpulan     *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
}

func (HasilPemeriksaanUsgGynecologi) TableName() string {
	return "hasil_pemeriksaan_usg_gynecologi"
}
