package kecamatan

// Kecamatan tabel `kecamatan` (kecamatan).
type Kecamatan struct {
	KdKec int    `gorm:"column:kd_kec;primaryKey;autoIncrement;type:int(11);not null" json:"kd_kec"`
	NmKec string `gorm:"column:nm_kec;type:varchar(60);not null" json:"nm_kec"`
}

func (Kecamatan) TableName() string {
	return "kecamatan"
}
