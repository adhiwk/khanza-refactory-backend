package rekammedis

import "time"

// SkriningGiziKehamilan tabel `skrining_gizi_kehamilan` (skrining gizi kehamilan, RMDataSkriningGiziKehamilan).
type SkriningGiziKehamilan struct {
	NoRawat    string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal    *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Parameter1 *string    `gorm:"column:parameter1" json:"parameter1"`
	Skor1      string     `gorm:"column:skor1" json:"skor1"`
	Parameter2 *string    `gorm:"column:parameter2" json:"parameter2"`
	Skor2      string     `gorm:"column:skor2" json:"skor2"`
	Parameter3 *string    `gorm:"column:parameter3" json:"parameter3"`
	Skor3      string     `gorm:"column:skor3" json:"skor3"`
	Parameter4 *string    `gorm:"column:parameter4" json:"parameter4"`
	Skor4      string     `gorm:"column:skor4" json:"skor4"`
	NilaiSkor  *string    `gorm:"column:nilai_skor" json:"nilai_skor"`
	Keterangan *string    `gorm:"column:keterangan" json:"keterangan"`
	Nip        *string    `gorm:"column:nip" json:"nip"`
}

func (SkriningGiziKehamilan) TableName() string {
	return "skrining_gizi_kehamilan"
}
