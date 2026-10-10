package rekammedis

import "time"

// CatatanKeperawatanRalan tabel `catatan_keperawatan_ralan` (catatan keperawatan ralan, RMDataCatatanKeperawatanRalan).
type CatatanKeperawatanRalan struct {
	Tanggal *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Jam     string     `gorm:"column:jam;primaryKey;autoIncrement:false" json:"jam"`
	NoRawat string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Uraian  *string    `gorm:"column:uraian" json:"uraian"`
	Nip     string     `gorm:"column:nip" json:"nip"`
}

func (CatatanKeperawatanRalan) TableName() string {
	return "catatan_keperawatan_ralan"
}
