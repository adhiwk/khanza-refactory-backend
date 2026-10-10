package bangsal

// Bangsal tabel `bangsal` (bangsal).
type Bangsal struct {
	KdBangsal string  `gorm:"column:kd_bangsal;primaryKey;type:char(5);not null" json:"kd_bangsal"`
	NmBangsal *string `gorm:"column:nm_bangsal;type:varchar(30)" json:"nm_bangsal"`
	Status    *string `gorm:"column:status;type:enum('0','1')" json:"status"`
}

func (Bangsal) TableName() string {
	return "bangsal"
}
