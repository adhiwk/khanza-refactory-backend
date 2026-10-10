package rekammedis

import "time"

// CatatanPengkajianPaskaOperasi tabel `catatan_pengkajian_paska_operasi` (catatan pengkajian paska operasi, RMCatatanPengkajianPaskaOperasi).
type CatatanPengkajianPaskaOperasi struct {
	NoRawat            string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal            *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	KdDokter           string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	RawatPaskaOperasi  *string    `gorm:"column:rawat_paska_operasi" json:"rawat_paska_operasi"`
	Cairan             *string    `gorm:"column:cairan" json:"cairan"`
	Antibiotika        *string    `gorm:"column:antibiotika" json:"antibiotika"`
	Analgetika         *string    `gorm:"column:analgetika" json:"analgetika"`
	MedikamentosaLain  *string    `gorm:"column:medikamentosa_lain" json:"medikamentosa_lain"`
	Diet               *string    `gorm:"column:diet" json:"diet"`
	PemeriksaanLaborat *string    `gorm:"column:pemeriksaan_laborat" json:"pemeriksaan_laborat"`
	Tranfusi           *string    `gorm:"column:tranfusi" json:"tranfusi"`
	Lainlain           *string    `gorm:"column:lainlain" json:"lainlain"`
}

func (CatatanPengkajianPaskaOperasi) TableName() string {
	return "catatan_pengkajian_paska_operasi"
}
