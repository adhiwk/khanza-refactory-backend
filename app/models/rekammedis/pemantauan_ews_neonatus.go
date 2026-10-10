package rekammedis

import "time"

// PemantauanEwsNeonatus tabel `pemantauan_ews_neonatus` (pemantauan EWS neonatus, RMPemantauanEWSNeonatus).
type PemantauanEwsNeonatus struct {
	NoRawat        string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal        *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Parameter1     *string    `gorm:"column:parameter1" json:"parameter1"`
	Skor1          *string    `gorm:"column:skor1" json:"skor1"`
	Parameter2     *string    `gorm:"column:parameter2" json:"parameter2"`
	Skor2          *string    `gorm:"column:skor2" json:"skor2"`
	Parameter3     *string    `gorm:"column:parameter3" json:"parameter3"`
	Skor3          *string    `gorm:"column:skor3" json:"skor3"`
	Parameter4     *string    `gorm:"column:parameter4" json:"parameter4"`
	Skor4          *string    `gorm:"column:skor4" json:"skor4"`
	Parameter5     *string    `gorm:"column:parameter5" json:"parameter5"`
	Skor5          *string    `gorm:"column:skor5" json:"skor5"`
	Parameter6     *string    `gorm:"column:parameter6" json:"parameter6"`
	Skor6          *string    `gorm:"column:skor6" json:"skor6"`
	Parameter7     *string    `gorm:"column:parameter7" json:"parameter7"`
	Skor7          *string    `gorm:"column:skor7" json:"skor7"`
	Parameter8     *string    `gorm:"column:parameter8" json:"parameter8"`
	Skor8          *string    `gorm:"column:skor8" json:"skor8"`
	SkorTotal      *string    `gorm:"column:skor_total" json:"skor_total"`
	ParameterTotal *string    `gorm:"column:parameter_total" json:"parameter_total"`
	CodeBlue       string     `gorm:"column:code_blue" json:"code_blue"`
	Nip            *string    `gorm:"column:nip" json:"nip"`
}

func (PemantauanEwsNeonatus) TableName() string {
	return "pemantauan_ews_neonatus"
}
