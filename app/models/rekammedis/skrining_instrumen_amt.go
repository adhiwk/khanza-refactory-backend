package rekammedis

import "time"

// SkriningInstrumenAmt tabel `skrining_instrumen_amt` (skrining instrumen AMT, RMSkriningInstrumenAMT).
type SkriningInstrumenAmt struct {
	NoRawat         string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal         *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip             string     `gorm:"column:nip" json:"nip"`
	Pernyataanamt1  *string    `gorm:"column:pernyataanamt1" json:"pernyataanamt1"`
	NilaiAmt1       *int       `gorm:"column:nilai_amt1" json:"nilai_amt1"`
	Pernyataanamt2  *string    `gorm:"column:pernyataanamt2" json:"pernyataanamt2"`
	NilaiAmt2       *int       `gorm:"column:nilai_amt2" json:"nilai_amt2"`
	Pernyataanamt3  *string    `gorm:"column:pernyataanamt3" json:"pernyataanamt3"`
	NilaiAmt3       *int       `gorm:"column:nilai_amt3" json:"nilai_amt3"`
	Pernyataanamt4  *string    `gorm:"column:pernyataanamt4" json:"pernyataanamt4"`
	NilaiAmt4       *int       `gorm:"column:nilai_amt4" json:"nilai_amt4"`
	Pernyataanamt5  *string    `gorm:"column:pernyataanamt5" json:"pernyataanamt5"`
	NilaiAmt5       *int       `gorm:"column:nilai_amt5" json:"nilai_amt5"`
	Pernyataanamt6  *string    `gorm:"column:pernyataanamt6" json:"pernyataanamt6"`
	NilaiAmt6       *int       `gorm:"column:nilai_amt6" json:"nilai_amt6"`
	Pernyataanamt7  *string    `gorm:"column:pernyataanamt7" json:"pernyataanamt7"`
	NilaiAmt7       *int       `gorm:"column:nilai_amt7" json:"nilai_amt7"`
	Pernyataanamt8  *string    `gorm:"column:pernyataanamt8" json:"pernyataanamt8"`
	NilaiAmt8       *int       `gorm:"column:nilai_amt8" json:"nilai_amt8"`
	Pernyataanamt9  *string    `gorm:"column:pernyataanamt9" json:"pernyataanamt9"`
	NilaiAmt9       *int       `gorm:"column:nilai_amt9" json:"nilai_amt9"`
	Pernyataanamt10 *string    `gorm:"column:pernyataanamt10" json:"pernyataanamt10"`
	NilaiAmt10      *int       `gorm:"column:nilai_amt10" json:"nilai_amt10"`
	NilaiTotalAmt   *int       `gorm:"column:nilai_total_amt" json:"nilai_total_amt"`
	Kesimpulan      *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
}

func (SkriningInstrumenAmt) TableName() string {
	return "skrining_instrumen_amt"
}
