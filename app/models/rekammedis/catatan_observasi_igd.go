package rekammedis

import "time"

// CatatanObservasiIgd tabel `catatan_observasi_igd` (catatan observasi IGD, RMDataCatatanObservasiIGD).
type CatatanObservasiIgd struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat     string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	Gcs          *string    `gorm:"column:gcs" json:"gcs"`
	Td           string     `gorm:"column:td" json:"td"`
	Hr           *string    `gorm:"column:hr" json:"hr"`
	Rr           *string    `gorm:"column:rr" json:"rr"`
	Suhu         *string    `gorm:"column:suhu" json:"suhu"`
	Spo2         string     `gorm:"column:spo2" json:"spo2"`
	Nip          string     `gorm:"column:nip" json:"nip"`
}

func (CatatanObservasiIgd) TableName() string {
	return "catatan_observasi_igd"
}
