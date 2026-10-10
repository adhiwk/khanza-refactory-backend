package rekammedis

import "time"

// CatatanObservasiInduksiPersalinan tabel `catatan_observasi_induksi_persalinan` (catatan observasi induksi persalinan, RMDataCatatanObservasiInduksiPersalinan).
type CatatanObservasiInduksiPersalinan struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat     string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	Obat         *string    `gorm:"column:obat" json:"obat"`
	Cairan       string     `gorm:"column:cairan" json:"cairan"`
	Dosis        *string    `gorm:"column:dosis" json:"dosis"`
	His          *string    `gorm:"column:his" json:"his"`
	Djj          *string    `gorm:"column:djj" json:"djj"`
	Keterangan   *string    `gorm:"column:keterangan" json:"keterangan"`
	Nip          string     `gorm:"column:nip" json:"nip"`
}

func (CatatanObservasiInduksiPersalinan) TableName() string {
	return "catatan_observasi_induksi_persalinan"
}
