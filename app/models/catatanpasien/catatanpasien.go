package catatanpasien

// CatatanPasien tabel `catatan_pasien` (catatan pasien).
type CatatanPasien struct {
	NoRkmMedis string  `gorm:"column:no_rkm_medis;primaryKey;type:varchar(15);not null" json:"no_rkm_medis"`
	Catatan    *string `gorm:"column:catatan;type:text" json:"catatan"`
}

func (CatatanPasien) TableName() string {
	return "catatan_pasien"
}
