package rekammedis

import "time"

// CatatanObservasiVentilator tabel `catatan_observasi_ventilator` (catatan observasi ventilator, RMDataCatatanObservasiVentilator).
type CatatanObservasiVentilator struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat     string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	Mode         *string    `gorm:"column:mode" json:"mode"`
	Vt           *string    `gorm:"column:vt" json:"vt"`
	Pakar        *string    `gorm:"column:pakar" json:"pakar"`
	Rr           *string    `gorm:"column:rr" json:"rr"`
	Reefps       *string    `gorm:"column:reefps" json:"reefps"`
	Ee           *string    `gorm:"column:ee" json:"ee"`
	Keterangan   string     `gorm:"column:keterangan" json:"keterangan"`
	Nip          string     `gorm:"column:nip" json:"nip"`
}

func (CatatanObservasiVentilator) TableName() string {
	return "catatan_observasi_ventilator"
}
