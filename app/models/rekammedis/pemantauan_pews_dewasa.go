package rekammedis

import "time"

// PemantauanPewsDewasa tabel `pemantauan_pews_dewasa` (pemantauan EWSD, RMPemantauanEWSD).
type PemantauanPewsDewasa struct {
	NoRawat                       string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                       *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	ParameterLajuRespirasi        *string    `gorm:"column:parameter_laju_respirasi" json:"parameter_laju_respirasi"`
	SkorLajuRespirasi             *string    `gorm:"column:skor_laju_respirasi" json:"skor_laju_respirasi"`
	ParameterSaturasiOksigen      *string    `gorm:"column:parameter_saturasi_oksigen" json:"parameter_saturasi_oksigen"`
	SkorSaturasiOksigen           *string    `gorm:"column:skor_saturasi_oksigen" json:"skor_saturasi_oksigen"`
	ParameterSuplemenOksigen      *string    `gorm:"column:parameter_suplemen_oksigen" json:"parameter_suplemen_oksigen"`
	SkorSuplemenOksigen           *string    `gorm:"column:skor_suplemen_oksigen" json:"skor_suplemen_oksigen"`
	ParameterTekananDarahSistolik *string    `gorm:"column:parameter_tekanan_darah_sistolik" json:"parameter_tekanan_darah_sistolik"`
	SkorTekananDarahSistolik      *string    `gorm:"column:skor_tekanan_darah_sistolik" json:"skor_tekanan_darah_sistolik"`
	ParameterLajuJantung          *string    `gorm:"column:parameter_laju_jantung" json:"parameter_laju_jantung"`
	SkorLajuJantung               *string    `gorm:"column:skor_laju_jantung" json:"skor_laju_jantung"`
	ParameterKesadaran            *string    `gorm:"column:parameter_kesadaran" json:"parameter_kesadaran"`
	SkorKesadaran                 *string    `gorm:"column:skor_kesadaran" json:"skor_kesadaran"`
	ParameterTemperatur           *string    `gorm:"column:parameter_temperatur" json:"parameter_temperatur"`
	SkorTemperatur                *string    `gorm:"column:skor_temperatur" json:"skor_temperatur"`
	SkorTotal                     *string    `gorm:"column:skor_total" json:"skor_total"`
	ParameterTotal                *string    `gorm:"column:parameter_total" json:"parameter_total"`
	Nip                           *string    `gorm:"column:nip" json:"nip"`
}

func (PemantauanPewsDewasa) TableName() string {
	return "pemantauan_pews_dewasa"
}
