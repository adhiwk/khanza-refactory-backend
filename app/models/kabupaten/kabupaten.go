package kabupaten

// Kabupaten tabel `kabupaten` (kabupaten).
type Kabupaten struct {
	KdKab int    `gorm:"column:kd_kab;primaryKey;autoIncrement;type:int(11);not null" json:"kd_kab"`
	NmKab string `gorm:"column:nm_kab;type:varchar(60);not null" json:"nm_kab"`
}

func (Kabupaten) TableName() string {
	return "kabupaten"
}
