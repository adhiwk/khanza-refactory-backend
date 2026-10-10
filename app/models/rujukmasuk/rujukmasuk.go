package rujukmasuk

// RujukMasuk tabel `rujuk_masuk` (rujukan masuk).
type RujukMasuk struct {
	NoRawat       string  `gorm:"column:no_rawat;primaryKey;type:varchar(17);not null" json:"no_rawat"`
	Perujuk       *string `gorm:"column:perujuk;type:varchar(60)" json:"perujuk"`
	Alamat        string  `gorm:"column:alamat;type:varchar(70);not null" json:"alamat"`
	NoRujuk       string  `gorm:"column:no_rujuk;type:varchar(40);not null" json:"no_rujuk"`
	JmPerujuk     float64 `gorm:"column:jm_perujuk;type:double;not null" json:"jm_perujuk"`
	DokterPerujuk *string `gorm:"column:dokter_perujuk;type:varchar(50)" json:"dokter_perujuk"`
	KdPenyakit    *string `gorm:"column:kd_penyakit;type:varchar(15)" json:"kd_penyakit"`
	KategoriRujuk *string `gorm:"column:kategori_rujuk;type:enum('-','Bedah','Non-Bedah','Kebidanan','Anak')" json:"kategori_rujuk"`
	Keterangan    *string `gorm:"column:keterangan;type:varchar(200)" json:"keterangan"`
	NoBalasan     *string `gorm:"column:no_balasan;type:varchar(20)" json:"no_balasan"`
}

func (RujukMasuk) TableName() string {
	return "rujuk_masuk"
}
