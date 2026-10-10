package golonganobat

// GolonganBarang tabel `golongan_barang` (golongan obat).
type GolonganBarang struct {
	Kode string  `gorm:"column:kode;primaryKey;type:char(4);not null" json:"kode"`
	Nama *string `gorm:"column:nama;type:varchar(30)" json:"nama"`
}

func (GolonganBarang) TableName() string {
	return "golongan_barang"
}
