package rekammedis

import "time"

// SkriningGizi tabel `skrining_gizi` (skrining gizi lanjut, RMDataSkriningGiziLanjut).
type SkriningGizi struct {
	NoRawat           string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal           *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	SkriningBb        *string    `gorm:"column:skrining_bb" json:"skrining_bb"`
	SkriningTb        *string    `gorm:"column:skrining_tb" json:"skrining_tb"`
	Alergi            *string    `gorm:"column:alergi" json:"alergi"`
	ParameterImt      *string    `gorm:"column:parameter_imt" json:"parameter_imt"`
	SkorImt           *string    `gorm:"column:skor_imt" json:"skor_imt"`
	ParameterBb       *string    `gorm:"column:parameter_bb" json:"parameter_bb"`
	SkorBb            *string    `gorm:"column:skor_bb" json:"skor_bb"`
	ParameterPenyakit *string    `gorm:"column:parameter_penyakit" json:"parameter_penyakit"`
	SkorPenyakit      *string    `gorm:"column:skor_penyakit" json:"skor_penyakit"`
	SkorTotal         *string    `gorm:"column:skor_total" json:"skor_total"`
	ParameterTotal    *string    `gorm:"column:parameter_total" json:"parameter_total"`
	Nip               *string    `gorm:"column:nip" json:"nip"`
}

func (SkriningGizi) TableName() string {
	return "skrining_gizi"
}
