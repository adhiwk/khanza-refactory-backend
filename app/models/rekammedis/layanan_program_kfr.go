package rekammedis

import "time"

// LayananProgramKfr tabel `layanan_program_kfr` (layanan program KFR, RMLayananProgramKFR).
type LayananProgramKfr struct {
	NoRawatLayanan string     `gorm:"column:no_rawat_layanan" json:"no_rawat_layanan"`
	NoRawat        string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal        *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip            string     `gorm:"column:nip" json:"nip"`
	Program        *string    `gorm:"column:program" json:"program"`
}

func (LayananProgramKfr) TableName() string {
	return "layanan_program_kfr"
}
