package satuanobat

// Kodesatuan tabel `kodesatuan` (satuan obat).
type Kodesatuan struct {
	KodeSat string  `gorm:"column:kode_sat;primaryKey;type:char(4);not null" json:"kode_sat"`
	Satuan  *string `gorm:"column:satuan;type:varchar(30)" json:"satuan"`
}

func (Kodesatuan) TableName() string {
	return "kodesatuan"
}
