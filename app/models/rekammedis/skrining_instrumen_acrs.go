package rekammedis

import "time"

// SkriningInstrumenAcrs tabel `skrining_instrumen_acrs` (skrining instrumen ACRS, RMSkriningInstrumenACRS).
type SkriningInstrumenAcrs struct {
	NoRawat          string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal          *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip              string     `gorm:"column:nip" json:"nip"`
	Pernyataanacrs1  *string    `gorm:"column:pernyataanacrs1" json:"pernyataanacrs1"`
	NilaiAcrs1       *int       `gorm:"column:nilai_acrs1" json:"nilai_acrs1"`
	Pernyataanacrs2  *string    `gorm:"column:pernyataanacrs2" json:"pernyataanacrs2"`
	NilaiAcrs2       *int       `gorm:"column:nilai_acrs2" json:"nilai_acrs2"`
	Pernyataanacrs3  *string    `gorm:"column:pernyataanacrs3" json:"pernyataanacrs3"`
	NilaiAcrs3       *int       `gorm:"column:nilai_acrs3" json:"nilai_acrs3"`
	Pernyataanacrs4  *string    `gorm:"column:pernyataanacrs4" json:"pernyataanacrs4"`
	NilaiAcrs4       *int       `gorm:"column:nilai_acrs4" json:"nilai_acrs4"`
	Pernyataanacrs5  *string    `gorm:"column:pernyataanacrs5" json:"pernyataanacrs5"`
	NilaiAcrs5       *int       `gorm:"column:nilai_acrs5" json:"nilai_acrs5"`
	Pernyataanacrs6  *string    `gorm:"column:pernyataanacrs6" json:"pernyataanacrs6"`
	NilaiAcrs6       *int       `gorm:"column:nilai_acrs6" json:"nilai_acrs6"`
	Pernyataanacrs7  *string    `gorm:"column:pernyataanacrs7" json:"pernyataanacrs7"`
	NilaiAcrs7       *int       `gorm:"column:nilai_acrs7" json:"nilai_acrs7"`
	Pernyataanacrs8  *string    `gorm:"column:pernyataanacrs8" json:"pernyataanacrs8"`
	NilaiAcrs8       *int       `gorm:"column:nilai_acrs8" json:"nilai_acrs8"`
	Pernyataanacrs9  *string    `gorm:"column:pernyataanacrs9" json:"pernyataanacrs9"`
	NilaiAcrs9       *int       `gorm:"column:nilai_acrs9" json:"nilai_acrs9"`
	Pernyataanacrs10 *string    `gorm:"column:pernyataanacrs10" json:"pernyataanacrs10"`
	NilaiAcrs10      *int       `gorm:"column:nilai_acrs10" json:"nilai_acrs10"`
	NilaiTotalAcrs   *int       `gorm:"column:nilai_total_acrs" json:"nilai_total_acrs"`
	Kesimpulan       *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
}

func (SkriningInstrumenAcrs) TableName() string {
	return "skrining_instrumen_acrs"
}
