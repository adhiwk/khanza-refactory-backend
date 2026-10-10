package kamar

// Kamar tabel `kamar` (kamar).
type Kamar struct {
	KdKamar    string   `gorm:"column:kd_kamar;primaryKey;type:varchar(15);not null" json:"kd_kamar"`
	KdBangsal  *string  `gorm:"column:kd_bangsal;type:char(5)" json:"kd_bangsal"`
	TrfKamar   *float64 `gorm:"column:trf_kamar;type:double" json:"trf_kamar"`
	Status     *string  `gorm:"column:status;type:enum('ISI','KOSONG','DIBERSIHKAN','DIBOOKING','PERBAIKAN')" json:"status"`
	Kelas      *string  `gorm:"column:kelas;type:enum('Kelas 1','Kelas 2','Kelas 3','Kelas Utama','Kelas VIP','Kelas VVIP')" json:"kelas"`
	Statusdata *string  `gorm:"column:statusdata;type:enum('0','1')" json:"statusdata"`
}

func (Kamar) TableName() string {
	return "kamar"
}
