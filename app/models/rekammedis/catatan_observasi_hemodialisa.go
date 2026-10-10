package rekammedis

import "time"

// CatatanObservasiHemodialisa tabel `catatan_observasi_hemodialisa` (catatan observasi hemodialisa, RMDataCatatanObservasiHemodialisa).
type CatatanObservasiHemodialisa struct {
	NoRawat       string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan  *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat      string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	Qb            *string    `gorm:"column:qb" json:"qb"`
	Qd            string     `gorm:"column:qd" json:"qd"`
	TekananArteri *string    `gorm:"column:tekanan_arteri" json:"tekanan_arteri"`
	TekananVena   *string    `gorm:"column:tekanan_vena" json:"tekanan_vena"`
	Tmp           *string    `gorm:"column:tmp" json:"tmp"`
	Ufr           *string    `gorm:"column:ufr" json:"ufr"`
	Tensi         *string    `gorm:"column:tensi" json:"tensi"`
	Nadi          *string    `gorm:"column:nadi" json:"nadi"`
	Suhu          *string    `gorm:"column:suhu" json:"suhu"`
	Spo2          *string    `gorm:"column:spo2" json:"spo2"`
	Tindakan      *string    `gorm:"column:tindakan" json:"tindakan"`
	Ufg           *string    `gorm:"column:ufg" json:"ufg"`
	BarcodeHf     string     `gorm:"column:barcode_hf" json:"barcode_hf"`
	Nip           string     `gorm:"column:nip" json:"nip"`
}

func (CatatanObservasiHemodialisa) TableName() string {
	return "catatan_observasi_hemodialisa"
}
