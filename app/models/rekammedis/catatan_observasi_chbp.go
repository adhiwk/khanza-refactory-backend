package rekammedis

import "time"

// CatatanObservasiChbp tabel `catatan_observasi_chbp` (catatan observasi CHBP, RMDataCatatanObservasiCHBP).
type CatatanObservasiChbp struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat     string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	Td           string     `gorm:"column:td" json:"td"`
	Hr           *string    `gorm:"column:hr" json:"hr"`
	Suhu         *string    `gorm:"column:suhu" json:"suhu"`
	Djj          string     `gorm:"column:djj" json:"djj"`
	His          string     `gorm:"column:his" json:"his"`
	Ppv          string     `gorm:"column:ppv" json:"ppv"`
	Keterangan   string     `gorm:"column:keterangan" json:"keterangan"`
	Nip          string     `gorm:"column:nip" json:"nip"`
}

func (CatatanObservasiChbp) TableName() string {
	return "catatan_observasi_chbp"
}
