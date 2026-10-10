package rekammedis

import "time"

// PemantauanMeowsObstetri tabel `pemantauan_meows_obstetri` (pemantauan MEOWS, RMPemantauanMEOWS).
type PemantauanMeowsObstetri struct {
	NoRawat                       string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                       *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	ParameterPernapasan           *string    `gorm:"column:parameter_pernapasan" json:"parameter_pernapasan"`
	SkorPernapasan                *string    `gorm:"column:skor_pernapasan" json:"skor_pernapasan"`
	ParameterSaturasi             *string    `gorm:"column:parameter_saturasi" json:"parameter_saturasi"`
	SkorSaturasi                  *string    `gorm:"column:skor_saturasi" json:"skor_saturasi"`
	ParameterTemperatur           *string    `gorm:"column:parameter_temperatur" json:"parameter_temperatur"`
	SkorTemperatur                *string    `gorm:"column:skor_temperatur" json:"skor_temperatur"`
	ParameterTekananDarahSistole  *string    `gorm:"column:parameter_tekanan_darah_sistole" json:"parameter_tekanan_darah_sistole"`
	SkorTekananDarahSistole       *string    `gorm:"column:skor_tekanan_darah_sistole" json:"skor_tekanan_darah_sistole"`
	ParameterTekananDarahDiastole *string    `gorm:"column:parameter_tekanan_darah_diastole" json:"parameter_tekanan_darah_diastole"`
	SkorTekananDarahDiastole      *string    `gorm:"column:skor_tekanan_darah_diastole" json:"skor_tekanan_darah_diastole"`
	ParameterDenyutJantung        *string    `gorm:"column:parameter_denyut_jantung" json:"parameter_denyut_jantung"`
	SkorDenyutJantung             *string    `gorm:"column:skor_denyut_jantung" json:"skor_denyut_jantung"`
	ParameterKesadaran            *string    `gorm:"column:parameter_kesadaran" json:"parameter_kesadaran"`
	SkorKesadaran                 *string    `gorm:"column:skor_kesadaran" json:"skor_kesadaran"`
	ParameterKetuban              *string    `gorm:"column:parameter_ketuban" json:"parameter_ketuban"`
	SkorKetuban                   *string    `gorm:"column:skor_ketuban" json:"skor_ketuban"`
	ParameterDischarge            *string    `gorm:"column:parameter_discharge" json:"parameter_discharge"`
	SkorDischarge                 *string    `gorm:"column:skor_discharge" json:"skor_discharge"`
	ParameterProteinuria          *string    `gorm:"column:parameter_proteinuria" json:"parameter_proteinuria"`
	SkorProteinuria               *string    `gorm:"column:skor_proteinuria" json:"skor_proteinuria"`
	SkorTotal                     string     `gorm:"column:skor_total" json:"skor_total"`
	ParameterTotal                *string    `gorm:"column:parameter_total" json:"parameter_total"`
	CodeBlue                      string     `gorm:"column:code_blue" json:"code_blue"`
	Nip                           *string    `gorm:"column:nip" json:"nip"`
}

func (PemantauanMeowsObstetri) TableName() string {
	return "pemantauan_meows_obstetri"
}
