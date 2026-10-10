package kelurahan

// Kelurahan tabel `kelurahan` (kelurahan).
type Kelurahan struct {
	KdKel int    `gorm:"column:kd_kel;primaryKey;autoIncrement;type:int(11);not null" json:"kd_kel"`
	NmKel string `gorm:"column:nm_kel;type:varchar(60);not null" json:"nm_kel"`
}

func (Kelurahan) TableName() string {
	return "kelurahan"
}
