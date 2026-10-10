package spesialis

// Spesialis tabel `spesialis` (spesialis).
type Spesialis struct {
	KdSps string  `gorm:"column:kd_sps;primaryKey;type:char(5);not null" json:"kd_sps"`
	NmSps *string `gorm:"column:nm_sps;type:varchar(30)" json:"nm_sps"`
}

func (Spesialis) TableName() string {
	return "spesialis"
}
