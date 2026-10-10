package rekammedis

import "time"

// SkriningKesehatanGigiMulutDewasa tabel `skrining_kesehatan_gigi_mulut_dewasa` (skrining kesehatan gigi mulut dewasa, RMSkriningKesehatanGigiMulutDewasa).
type SkriningKesehatanGigiMulutDewasa struct {
	NoRawat          string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal          *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KontrolGigi      *string    `gorm:"column:kontrol_gigi" json:"kontrol_gigi"`
	GigiBungsuTumbuh *string    `gorm:"column:gigi_bungsu_tumbuh" json:"gigi_bungsu_tumbuh"`
	GigiHilang       *string    `gorm:"column:gigi_hilang" json:"gigi_hilang"`
	GigiBerlubang    *string    `gorm:"column:gigi_berlubang" json:"gigi_berlubang"`
	HasilSkrining    *string    `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Keterangan       string     `gorm:"column:keterangan" json:"keterangan"`
	Nip              string     `gorm:"column:nip" json:"nip"`
}

func (SkriningKesehatanGigiMulutDewasa) TableName() string {
	return "skrining_kesehatan_gigi_mulut_dewasa"
}
