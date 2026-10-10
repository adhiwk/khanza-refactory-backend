package bahasapasien

// BahasaPasien tabel `bahasa_pasien` (bahasa pasien).
type BahasaPasien struct {
	Id         int     `gorm:"column:id;primaryKey;autoIncrement;type:int(11);not null" json:"id"`
	NamaBahasa *string `gorm:"column:nama_bahasa;type:varchar(30)" json:"nama_bahasa"`
}

func (BahasaPasien) TableName() string {
	return "bahasa_pasien"
}
