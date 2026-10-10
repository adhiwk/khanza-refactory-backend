package rekammedis

import "time"

// KonselingFarmasi tabel `konseling_farmasi` (konseling farmasi, RMKonselingFarmasi).
type KonselingFarmasi struct {
	NoRawat       string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal       *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Diagnosa      *string    `gorm:"column:diagnosa" json:"diagnosa"`
	ObatPemakaian *string    `gorm:"column:obat_pemakaian" json:"obat_pemakaian"`
	RiwayatAlergi *string    `gorm:"column:riwayat_alergi" json:"riwayat_alergi"`
	Keluhan       *string    `gorm:"column:keluhan" json:"keluhan"`
	PernahDatang  *string    `gorm:"column:pernah_datang" json:"pernah_datang"`
	TindakLanjut  *string    `gorm:"column:tindak_lanjut" json:"tindak_lanjut"`
	Nip           string     `gorm:"column:nip" json:"nip"`
}

func (KonselingFarmasi) TableName() string {
	return "konseling_farmasi"
}
