package kategoriobat

// KategoriBarang tabel `kategori_barang` (kategori obat).
type KategoriBarang struct {
	Kode string  `gorm:"column:kode;primaryKey;type:char(4);not null" json:"kode"`
	Nama *string `gorm:"column:nama;type:varchar(30)" json:"nama"`
}

func (KategoriBarang) TableName() string {
	return "kategori_barang"
}
