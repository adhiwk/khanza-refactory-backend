package dpjp

// Dpjp dokter penanggung jawab pelayanan rawat inap (tabel dpjp_ranap).
type Dpjp struct {
	NoRawat  string `gorm:"column:no_rawat" json:"no_rawat"`
	KdDokter string `gorm:"column:kd_dokter" json:"kd_dokter"`
	NmDokter string `gorm:"column:nm_dokter;->" json:"nm_dokter"`
}
