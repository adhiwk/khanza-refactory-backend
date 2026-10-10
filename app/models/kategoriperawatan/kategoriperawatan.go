package kategoriperawatan

// KategoriPerawatan tabel `kategori_perawatan` (kategori perawatan).
type KategoriPerawatan struct {
	KdKategori string  `gorm:"column:kd_kategori;primaryKey;type:char(5);not null" json:"kd_kategori"`
	NmKategori *string `gorm:"column:nm_kategori;type:varchar(30)" json:"nm_kategori"`
}

func (KategoriPerawatan) TableName() string {
	return "kategori_perawatan"
}
