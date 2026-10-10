package rekammedis

import "time"

// SkriningKesehatanGigiMulutLansia tabel `skrining_kesehatan_gigi_mulut_lansia` (skrining kesehatan gigi mulut lansia, RMSkriningKesehatanGigiMulutLansia).
type SkriningKesehatanGigiMulutLansia struct {
	NoRawat       string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal       *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KontrolGigi   *string    `gorm:"column:kontrol_gigi" json:"kontrol_gigi"`
	PolaMakan     *string    `gorm:"column:pola_makan" json:"pola_makan"`
	SikatGigi     *string    `gorm:"column:sikat_gigi" json:"sikat_gigi"`
	GigiPalsu     *string    `gorm:"column:gigi_palsu" json:"gigi_palsu"`
	GigiBerfungsi *string    `gorm:"column:gigi_berfungsi" json:"gigi_berfungsi"`
	MukosaMulut   *string    `gorm:"column:mukosa_mulut" json:"mukosa_mulut"`
	HasilSkrining *string    `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Keterangan    string     `gorm:"column:keterangan" json:"keterangan"`
	Nip           string     `gorm:"column:nip" json:"nip"`
}

func (SkriningKesehatanGigiMulutLansia) TableName() string {
	return "skrining_kesehatan_gigi_mulut_lansia"
}
