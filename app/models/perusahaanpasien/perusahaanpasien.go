package perusahaanpasien

// PerusahaanPasien tabel `perusahaan_pasien` (instansi/perusahaan pasien).
type PerusahaanPasien struct {
	KodePerusahaan string  `gorm:"column:kode_perusahaan;primaryKey;type:varchar(8);not null" json:"kode_perusahaan"`
	NamaPerusahaan *string `gorm:"column:nama_perusahaan;type:varchar(70)" json:"nama_perusahaan"`
	Alamat         *string `gorm:"column:alamat;type:varchar(100)" json:"alamat"`
	Kota           *string `gorm:"column:kota;type:varchar(40)" json:"kota"`
	NoTelp         *string `gorm:"column:no_telp;type:varchar(27)" json:"no_telp"`
}

func (PerusahaanPasien) TableName() string {
	return "perusahaan_pasien"
}
