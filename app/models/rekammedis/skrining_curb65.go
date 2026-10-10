package rekammedis

import "time"

// SkriningCurb65 tabel `skrining_curb65` (skrining 65, RMSkriningCURB65).
type SkriningCurb65 struct {
	NoRawat           string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal           *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip               string     `gorm:"column:nip" json:"nip"`
	Pernyataancurb651 *string    `gorm:"column:pernyataancurb651" json:"pernyataancurb651"`
	NilaiCurb651      *int       `gorm:"column:nilai_curb651" json:"nilai_curb651"`
	Pernyataancurb652 *string    `gorm:"column:pernyataancurb652" json:"pernyataancurb652"`
	NilaiCurb652      *int       `gorm:"column:nilai_curb652" json:"nilai_curb652"`
	Pernyataancurb653 *string    `gorm:"column:pernyataancurb653" json:"pernyataancurb653"`
	NilaiCurb653      *int       `gorm:"column:nilai_curb653" json:"nilai_curb653"`
	Pernyataancurb654 *string    `gorm:"column:pernyataancurb654" json:"pernyataancurb654"`
	NilaiCurb654      *int       `gorm:"column:nilai_curb654" json:"nilai_curb654"`
	Pernyataancurb655 *string    `gorm:"column:pernyataancurb655" json:"pernyataancurb655"`
	NilaiCurb655      *int       `gorm:"column:nilai_curb655" json:"nilai_curb655"`
	NilaiTotalCurb65  *int       `gorm:"column:nilai_total_curb65" json:"nilai_total_curb65"`
	Kesimpulan        *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
}

func (SkriningCurb65) TableName() string {
	return "skrining_curb65"
}
