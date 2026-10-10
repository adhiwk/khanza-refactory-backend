package rekammedis

import "time"

// CatatanKeseimbanganCairan tabel `catatan_keseimbangan_cairan` (catatan keseimbangan cairan, RMDataCatatanKeseimbanganCairan).
type CatatanKeseimbanganCairan struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat     string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	Infus        *string    `gorm:"column:infus" json:"infus"`
	Tranfusi     string     `gorm:"column:tranfusi" json:"tranfusi"`
	Minum        *string    `gorm:"column:minum" json:"minum"`
	Urine        *string    `gorm:"column:urine" json:"urine"`
	Drain        *string    `gorm:"column:drain" json:"drain"`
	Ngt          string     `gorm:"column:ngt" json:"ngt"`
	Iwl          string     `gorm:"column:iwl" json:"iwl"`
	Keseimbangan string     `gorm:"column:keseimbangan" json:"keseimbangan"`
	Keterangan   string     `gorm:"column:keterangan" json:"keterangan"`
	Nip          string     `gorm:"column:nip" json:"nip"`
}

func (CatatanKeseimbanganCairan) TableName() string {
	return "catatan_keseimbangan_cairan"
}
