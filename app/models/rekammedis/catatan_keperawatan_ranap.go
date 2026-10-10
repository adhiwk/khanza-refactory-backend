package rekammedis

import "time"

// CatatanKeperawatanRanap tabel `catatan_keperawatan_ranap` (catatan keperawatan ranap, RMDataCatatanKeperawatanRanap).
type CatatanKeperawatanRanap struct {
	Tanggal *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Jam     string     `gorm:"column:jam;primaryKey;autoIncrement:false" json:"jam"`
	NoRawat string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Uraian  *string    `gorm:"column:uraian" json:"uraian"`
	Nip     string     `gorm:"column:nip" json:"nip"`
}

func (CatatanKeperawatanRanap) TableName() string {
	return "catatan_keperawatan_ranap"
}
