package rekammedis

import "time"

// SkriningNutrisiAnak tabel `skrining_nutrisi_anak` (skrining nutrisi anak, RMSkriningNutrisiAnak).
type SkriningNutrisiAnak struct {
	NoRawat                      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                      *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Td                           string     `gorm:"column:td" json:"td"`
	Hr                           string     `gorm:"column:hr" json:"hr"`
	Rr                           string     `gorm:"column:rr" json:"rr"`
	Suhu                         string     `gorm:"column:suhu" json:"suhu"`
	Bb                           string     `gorm:"column:bb" json:"bb"`
	Tbpb                         string     `gorm:"column:tbpb" json:"tbpb"`
	Spo2                         string     `gorm:"column:spo2" json:"spo2"`
	Alergi                       string     `gorm:"column:alergi" json:"alergi"`
	Sg1                          string     `gorm:"column:sg1" json:"sg1"`
	Nilai1                       string     `gorm:"column:nilai1" json:"nilai1"`
	Sg2                          string     `gorm:"column:sg2" json:"sg2"`
	Nilai2                       string     `gorm:"column:nilai2" json:"nilai2"`
	Sg3                          string     `gorm:"column:sg3" json:"sg3"`
	Nilai3                       string     `gorm:"column:nilai3" json:"nilai3"`
	Sg4                          string     `gorm:"column:sg4" json:"sg4"`
	Nilai4                       string     `gorm:"column:nilai4" json:"nilai4"`
	TotalHasil                   int        `gorm:"column:total_hasil" json:"total_hasil"`
	SkorNutrisi                  *string    `gorm:"column:skor_nutrisi" json:"skor_nutrisi"`
	DiketahuiDietisien           string     `gorm:"column:diketahui_dietisien" json:"diketahui_dietisien"`
	KeteranganDiketahuiDietisien string     `gorm:"column:keterangan_diketahui_dietisien" json:"keterangan_diketahui_dietisien"`
	Nip                          string     `gorm:"column:nip" json:"nip"`
}

func (SkriningNutrisiAnak) TableName() string {
	return "skrining_nutrisi_anak"
}
