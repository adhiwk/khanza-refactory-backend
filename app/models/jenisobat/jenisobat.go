package jenisobat

// Jenis tabel `jenis` (jenis obat).
type Jenis struct {
	Kdjns      string `gorm:"column:kdjns;primaryKey;type:char(4);not null" json:"kdjns"`
	Nama       string `gorm:"column:nama;type:varchar(30);not null" json:"nama"`
	Keterangan string `gorm:"column:keterangan;type:varchar(50);not null" json:"keterangan"`
}

func (Jenis) TableName() string {
	return "jenis"
}
