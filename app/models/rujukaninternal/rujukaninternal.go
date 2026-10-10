package rujukaninternal

// RujukanInternal rujukan ke poli/dokter lain dalam satu kunjungan (tabel rujukan_internal_poli).
type RujukanInternal struct {
	NoRawat  string `gorm:"column:no_rawat" json:"no_rawat"`
	KdDokter string `gorm:"column:kd_dokter" json:"kd_dokter"`
	NmDokter string `gorm:"column:nm_dokter;->" json:"nm_dokter"`
	KdPoli   string `gorm:"column:kd_poli" json:"kd_poli"`
	NmPoli   string `gorm:"column:nm_poli;->" json:"nm_poli"`
}
