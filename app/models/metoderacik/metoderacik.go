package metoderacik

// MetodeRacik tabel `metode_racik` (metode racik).
type MetodeRacik struct {
	KdRacik string `gorm:"column:kd_racik;primaryKey;type:varchar(3);not null" json:"kd_racik"`
	NmRacik string `gorm:"column:nm_racik;type:varchar(30);not null" json:"nm_racik"`
}

func (MetodeRacik) TableName() string {
	return "metode_racik"
}
