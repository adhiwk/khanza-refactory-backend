package rekammedis

import "time"

// CatatanCekGds tabel `catatan_cek_gds` (catatan cek GDS, RMDataCatatanCekGDS).
type CatatanCekGds struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat     string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	Gdp          *string    `gorm:"column:gdp" json:"gdp"`
	Insulin      *string    `gorm:"column:insulin" json:"insulin"`
	ObatGula     *string    `gorm:"column:obat_gula" json:"obat_gula"`
	Nip          string     `gorm:"column:nip" json:"nip"`
}

func (CatatanCekGds) TableName() string {
	return "catatan_cek_gds"
}
