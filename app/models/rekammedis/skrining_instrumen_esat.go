package rekammedis

import "time"

// SkriningInstrumenEsat tabel `skrining_instrumen_esat` (skrining instrumen ESAT, RMSkriningInstrumenESAT).
type SkriningInstrumenEsat struct {
	NoRawat          string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal          *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip              string     `gorm:"column:nip" json:"nip"`
	Pernyataanesat1  *string    `gorm:"column:pernyataanesat1" json:"pernyataanesat1"`
	Pernyataanesat2  *string    `gorm:"column:pernyataanesat2" json:"pernyataanesat2"`
	Pernyataanesat3  *string    `gorm:"column:pernyataanesat3" json:"pernyataanesat3"`
	Pernyataanesat4  *string    `gorm:"column:pernyataanesat4" json:"pernyataanesat4"`
	Pernyataanesat5  *string    `gorm:"column:pernyataanesat5" json:"pernyataanesat5"`
	Pernyataanesat6  *string    `gorm:"column:pernyataanesat6" json:"pernyataanesat6"`
	Pernyataanesat7  *string    `gorm:"column:pernyataanesat7" json:"pernyataanesat7"`
	Pernyataanesat8  *string    `gorm:"column:pernyataanesat8" json:"pernyataanesat8"`
	Pernyataanesat9  *string    `gorm:"column:pernyataanesat9" json:"pernyataanesat9"`
	Pernyataanesat10 *string    `gorm:"column:pernyataanesat10" json:"pernyataanesat10"`
	Kesimpulan       *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
	Catatan          *string    `gorm:"column:catatan" json:"catatan"`
}

func (SkriningInstrumenEsat) TableName() string {
	return "skrining_instrumen_esat"
}
