package industrifarmasi

// Industrifarmasi tabel `industrifarmasi` (industri farmasi).
type Industrifarmasi struct {
	KodeIndustri string  `gorm:"column:kode_industri;primaryKey;type:char(5);not null" json:"kode_industri"`
	NamaIndustri *string `gorm:"column:nama_industri;type:varchar(50)" json:"nama_industri"`
	Alamat       *string `gorm:"column:alamat;type:varchar(50)" json:"alamat"`
	Kota         *string `gorm:"column:kota;type:varchar(20)" json:"kota"`
	NoTelp       *string `gorm:"column:no_telp;type:varchar(20)" json:"no_telp"`
}

func (Industrifarmasi) TableName() string {
	return "industrifarmasi"
}
