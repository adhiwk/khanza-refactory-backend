package rekammedis

import "time"

// SkriningNutrisiLansia tabel `skrining_nutrisi_lansia` (skrining nutrisi lansia, RMSkriningNutrisiLansia).
type SkriningNutrisiLansia struct {
	NoRawat     string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal     *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Td          string     `gorm:"column:td" json:"td"`
	Hr          string     `gorm:"column:hr" json:"hr"`
	Rr          string     `gorm:"column:rr" json:"rr"`
	Suhu        string     `gorm:"column:suhu" json:"suhu"`
	Bb          string     `gorm:"column:bb" json:"bb"`
	Tbpb        string     `gorm:"column:tbpb" json:"tbpb"`
	Spo2        string     `gorm:"column:spo2" json:"spo2"`
	Alergi      string     `gorm:"column:alergi" json:"alergi"`
	Sg1         string     `gorm:"column:sg1" json:"sg1"`
	Nilai1      string     `gorm:"column:nilai1" json:"nilai1"`
	Sg2         string     `gorm:"column:sg2" json:"sg2"`
	Nilai2      string     `gorm:"column:nilai2" json:"nilai2"`
	Sg3         string     `gorm:"column:sg3" json:"sg3"`
	Nilai3      string     `gorm:"column:nilai3" json:"nilai3"`
	Sg4         string     `gorm:"column:sg4" json:"sg4"`
	Nilai4      string     `gorm:"column:nilai4" json:"nilai4"`
	Sg5         string     `gorm:"column:sg5" json:"sg5"`
	Nilai5      string     `gorm:"column:nilai5" json:"nilai5"`
	Sg6         string     `gorm:"column:sg6" json:"sg6"`
	Nilai6      string     `gorm:"column:nilai6" json:"nilai6"`
	TotalHasil  int        `gorm:"column:total_hasil" json:"total_hasil"`
	SkorNutrisi *string    `gorm:"column:skor_nutrisi" json:"skor_nutrisi"`
	Nip         string     `gorm:"column:nip" json:"nip"`
}

func (SkriningNutrisiLansia) TableName() string {
	return "skrining_nutrisi_lansia"
}
