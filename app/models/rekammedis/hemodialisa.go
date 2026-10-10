package rekammedis

import "time"

// Hemodialisa tabel `hemodialisa` (hemodialisa, RMHemodialisa).
type Hemodialisa struct {
	NoRawat    string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal    *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	KdDokter   *string    `gorm:"column:kd_dokter" json:"kd_dokter"`
	Lama       *string    `gorm:"column:lama" json:"lama"`
	Akses      *string    `gorm:"column:akses" json:"akses"`
	Dialist    *string    `gorm:"column:dialist" json:"dialist"`
	Transfusi  *string    `gorm:"column:transfusi" json:"transfusi"`
	Penarikan  *string    `gorm:"column:penarikan" json:"penarikan"`
	Qb         *string    `gorm:"column:qb" json:"qb"`
	Qd         *string    `gorm:"column:qd" json:"qd"`
	Ureum      *string    `gorm:"column:ureum" json:"ureum"`
	Hb         *string    `gorm:"column:hb" json:"hb"`
	Hbsag      *string    `gorm:"column:hbsag" json:"hbsag"`
	Creatinin  *string    `gorm:"column:creatinin" json:"creatinin"`
	Hiv        *string    `gorm:"column:hiv" json:"hiv"`
	Hcv        *string    `gorm:"column:hcv" json:"hcv"`
	Lain       *string    `gorm:"column:lain" json:"lain"`
	KdPenyakit *string    `gorm:"column:kd_penyakit" json:"kd_penyakit"`
}

func (Hemodialisa) TableName() string {
	return "hemodialisa"
}
