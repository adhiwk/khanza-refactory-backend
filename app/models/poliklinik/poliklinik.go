package poliklinik

// Poliklinik tabel `poliklinik` (poliklinik).
type Poliklinik struct {
	KdPoli         string  `gorm:"column:kd_poli;primaryKey;type:char(5);not null" json:"kd_poli"`
	NmPoli         *string `gorm:"column:nm_poli;type:varchar(50)" json:"nm_poli"`
	Registrasi     float64 `gorm:"column:registrasi;type:double;not null" json:"registrasi"`
	Registrasilama float64 `gorm:"column:registrasilama;type:double;not null" json:"registrasilama"`
	Status         string  `gorm:"column:status;type:enum('0','1');not null" json:"status"`
}

func (Poliklinik) TableName() string {
	return "poliklinik"
}
