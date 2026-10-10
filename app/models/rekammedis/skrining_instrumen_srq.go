package rekammedis

import "time"

// SkriningInstrumenSrq tabel `skrining_instrumen_srq` (skrining SRQ, RMSkriningSRQ).
type SkriningInstrumenSrq struct {
	NoRawat         string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal         *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip             string     `gorm:"column:nip" json:"nip"`
	Pernyataansrq1  *string    `gorm:"column:pernyataansrq1" json:"pernyataansrq1"`
	NilaiSrq1       *int       `gorm:"column:nilai_srq1" json:"nilai_srq1"`
	Pernyataansrq2  *string    `gorm:"column:pernyataansrq2" json:"pernyataansrq2"`
	NilaiSrq2       *int       `gorm:"column:nilai_srq2" json:"nilai_srq2"`
	Pernyataansrq3  *string    `gorm:"column:pernyataansrq3" json:"pernyataansrq3"`
	NilaiSrq3       *int       `gorm:"column:nilai_srq3" json:"nilai_srq3"`
	Pernyataansrq4  *string    `gorm:"column:pernyataansrq4" json:"pernyataansrq4"`
	NilaiSrq4       *int       `gorm:"column:nilai_srq4" json:"nilai_srq4"`
	Pernyataansrq5  *string    `gorm:"column:pernyataansrq5" json:"pernyataansrq5"`
	NilaiSrq5       *int       `gorm:"column:nilai_srq5" json:"nilai_srq5"`
	Pernyataansrq6  *string    `gorm:"column:pernyataansrq6" json:"pernyataansrq6"`
	NilaiSrq6       *int       `gorm:"column:nilai_srq6" json:"nilai_srq6"`
	Pernyataansrq7  *string    `gorm:"column:pernyataansrq7" json:"pernyataansrq7"`
	NilaiSrq7       *int       `gorm:"column:nilai_srq7" json:"nilai_srq7"`
	Pernyataansrq8  *string    `gorm:"column:pernyataansrq8" json:"pernyataansrq8"`
	NilaiSrq8       *int       `gorm:"column:nilai_srq8" json:"nilai_srq8"`
	Pernyataansrq9  *string    `gorm:"column:pernyataansrq9" json:"pernyataansrq9"`
	NilaiSrq9       *int       `gorm:"column:nilai_srq9" json:"nilai_srq9"`
	Pernyataansrq10 *string    `gorm:"column:pernyataansrq10" json:"pernyataansrq10"`
	NilaiSrq10      *int       `gorm:"column:nilai_srq10" json:"nilai_srq10"`
	Pernyataansrq11 *string    `gorm:"column:pernyataansrq11" json:"pernyataansrq11"`
	NilaiSrq11      *int       `gorm:"column:nilai_srq11" json:"nilai_srq11"`
	Pernyataansrq12 *string    `gorm:"column:pernyataansrq12" json:"pernyataansrq12"`
	NilaiSrq12      *int       `gorm:"column:nilai_srq12" json:"nilai_srq12"`
	Pernyataansrq13 *string    `gorm:"column:pernyataansrq13" json:"pernyataansrq13"`
	NilaiSrq13      *int       `gorm:"column:nilai_srq13" json:"nilai_srq13"`
	Pernyataansrq14 *string    `gorm:"column:pernyataansrq14" json:"pernyataansrq14"`
	NilaiSrq14      *int       `gorm:"column:nilai_srq14" json:"nilai_srq14"`
	Pernyataansrq15 *string    `gorm:"column:pernyataansrq15" json:"pernyataansrq15"`
	NilaiSrq15      *int       `gorm:"column:nilai_srq15" json:"nilai_srq15"`
	Pernyataansrq16 *string    `gorm:"column:pernyataansrq16" json:"pernyataansrq16"`
	NilaiSrq16      *int       `gorm:"column:nilai_srq16" json:"nilai_srq16"`
	Pernyataansrq17 *string    `gorm:"column:pernyataansrq17" json:"pernyataansrq17"`
	NilaiSrq17      *int       `gorm:"column:nilai_srq17" json:"nilai_srq17"`
	Pernyataansrq18 *string    `gorm:"column:pernyataansrq18" json:"pernyataansrq18"`
	NilaiSrq18      *int       `gorm:"column:nilai_srq18" json:"nilai_srq18"`
	Pernyataansrq19 *string    `gorm:"column:pernyataansrq19" json:"pernyataansrq19"`
	NilaiSrq19      *int       `gorm:"column:nilai_srq19" json:"nilai_srq19"`
	Pernyataansrq20 *string    `gorm:"column:pernyataansrq20" json:"pernyataansrq20"`
	NilaiSrq20      *int       `gorm:"column:nilai_srq20" json:"nilai_srq20"`
	NilaiTotalSrq   *int       `gorm:"column:nilai_total_srq" json:"nilai_total_srq"`
	Kesimpulan      *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
}

func (SkriningInstrumenSrq) TableName() string {
	return "skrining_instrumen_srq"
}
