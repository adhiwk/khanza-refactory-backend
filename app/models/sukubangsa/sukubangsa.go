package sukubangsa

// SukuBangsa tabel `suku_bangsa` (suku bangsa).
type SukuBangsa struct {
	Id             int     `gorm:"column:id;primaryKey;autoIncrement;type:int(11);not null" json:"id"`
	NamaSukuBangsa *string `gorm:"column:nama_suku_bangsa;type:varchar(30)" json:"nama_suku_bangsa"`
}

func (SukuBangsa) TableName() string {
	return "suku_bangsa"
}
