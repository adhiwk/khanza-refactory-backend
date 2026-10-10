package penjab

// Penjab tabel `penjab` (penanggung jawab).
type Penjab struct {
	KdPj           string `gorm:"column:kd_pj;primaryKey;type:char(3);not null" json:"kd_pj"`
	PngJawab       string `gorm:"column:png_jawab;type:varchar(30);not null" json:"png_jawab"`
	NamaPerusahaan string `gorm:"column:nama_perusahaan;type:varchar(60);not null" json:"nama_perusahaan"`
	AlamatAsuransi string `gorm:"column:alamat_asuransi;type:varchar(130);not null" json:"alamat_asuransi"`
	NoTelp         string `gorm:"column:no_telp;type:varchar(40);not null" json:"no_telp"`
	Attn           string `gorm:"column:attn;type:varchar(60);not null" json:"attn"`
	Status         string `gorm:"column:status;type:enum('0','1');not null" json:"status"`
}

func (Penjab) TableName() string {
	return "penjab"
}
