package rekammedis

import "time"

// CatatanObservasiRanapPostpartum tabel `catatan_observasi_ranap_postpartum` (catatan observasi ranap post partum, RMDataCatatanObservasiRanapPostPartum).
type CatatanObservasiRanapPostpartum struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat     string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	Gcs          *string    `gorm:"column:gcs" json:"gcs"`
	Td           string     `gorm:"column:td" json:"td"`
	Hr           *string    `gorm:"column:hr" json:"hr"`
	Rr           *string    `gorm:"column:rr" json:"rr"`
	Suhu         *string    `gorm:"column:suhu" json:"suhu"`
	Spo2         string     `gorm:"column:spo2" json:"spo2"`
	Tfu          string     `gorm:"column:tfu" json:"tfu"`
	Kontraksi    string     `gorm:"column:kontraksi" json:"kontraksi"`
	Perdarahan   string     `gorm:"column:perdarahan" json:"perdarahan"`
	Keterangan   string     `gorm:"column:keterangan" json:"keterangan"`
	Nip          string     `gorm:"column:nip" json:"nip"`
}

func (CatatanObservasiRanapPostpartum) TableName() string {
	return "catatan_observasi_ranap_postpartum"
}
