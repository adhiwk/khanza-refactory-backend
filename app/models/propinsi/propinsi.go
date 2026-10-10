package propinsi

// Propinsi tabel `propinsi` (propinsi).
type Propinsi struct {
	KdProp int    `gorm:"column:kd_prop;primaryKey;autoIncrement;type:int(11);not null" json:"kd_prop"`
	NmProp string `gorm:"column:nm_prop;type:varchar(30);not null" json:"nm_prop"`
}

func (Propinsi) TableName() string {
	return "propinsi"
}
