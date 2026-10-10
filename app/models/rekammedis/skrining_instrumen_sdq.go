package rekammedis

import "time"

// SkriningInstrumenSdq tabel `skrining_instrumen_sdq` (skrining instrumen SDQ, RMSkriningInstrumenSDQ).
type SkriningInstrumenSdq struct {
	NoRawat              string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal              *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip                  string     `gorm:"column:nip" json:"nip"`
	Pernyataansdq1       *string    `gorm:"column:pernyataansdq1" json:"pernyataansdq1"`
	NilaiSdq1            *int       `gorm:"column:nilai_sdq1" json:"nilai_sdq1"`
	Pernyataansdq2       *string    `gorm:"column:pernyataansdq2" json:"pernyataansdq2"`
	NilaiSdq2            *int       `gorm:"column:nilai_sdq2" json:"nilai_sdq2"`
	Pernyataansdq3       *string    `gorm:"column:pernyataansdq3" json:"pernyataansdq3"`
	NilaiSdq3            *int       `gorm:"column:nilai_sdq3" json:"nilai_sdq3"`
	Pernyataansdq4       *string    `gorm:"column:pernyataansdq4" json:"pernyataansdq4"`
	NilaiSdq4            *int       `gorm:"column:nilai_sdq4" json:"nilai_sdq4"`
	Pernyataansdq5       *string    `gorm:"column:pernyataansdq5" json:"pernyataansdq5"`
	NilaiSdq5            *int       `gorm:"column:nilai_sdq5" json:"nilai_sdq5"`
	Pernyataansdq6       *string    `gorm:"column:pernyataansdq6" json:"pernyataansdq6"`
	NilaiSdq6            *int       `gorm:"column:nilai_sdq6" json:"nilai_sdq6"`
	Pernyataansdq7       *string    `gorm:"column:pernyataansdq7" json:"pernyataansdq7"`
	NilaiSdq7            *int       `gorm:"column:nilai_sdq7" json:"nilai_sdq7"`
	Pernyataansdq8       *string    `gorm:"column:pernyataansdq8" json:"pernyataansdq8"`
	NilaiSdq8            *int       `gorm:"column:nilai_sdq8" json:"nilai_sdq8"`
	Pernyataansdq9       *string    `gorm:"column:pernyataansdq9" json:"pernyataansdq9"`
	NilaiSdq9            *int       `gorm:"column:nilai_sdq9" json:"nilai_sdq9"`
	Pernyataansdq10      *string    `gorm:"column:pernyataansdq10" json:"pernyataansdq10"`
	NilaiSdq10           *int       `gorm:"column:nilai_sdq10" json:"nilai_sdq10"`
	Pernyataansdq11      *string    `gorm:"column:pernyataansdq11" json:"pernyataansdq11"`
	NilaiSdq11           *int       `gorm:"column:nilai_sdq11" json:"nilai_sdq11"`
	Pernyataansdq12      *string    `gorm:"column:pernyataansdq12" json:"pernyataansdq12"`
	NilaiSdq12           *int       `gorm:"column:nilai_sdq12" json:"nilai_sdq12"`
	Pernyataansdq13      *string    `gorm:"column:pernyataansdq13" json:"pernyataansdq13"`
	NilaiSdq13           *int       `gorm:"column:nilai_sdq13" json:"nilai_sdq13"`
	Pernyataansdq14      *string    `gorm:"column:pernyataansdq14" json:"pernyataansdq14"`
	NilaiSdq14           *int       `gorm:"column:nilai_sdq14" json:"nilai_sdq14"`
	Pernyataansdq15      *string    `gorm:"column:pernyataansdq15" json:"pernyataansdq15"`
	NilaiSdq15           *int       `gorm:"column:nilai_sdq15" json:"nilai_sdq15"`
	Pernyataansdq16      *string    `gorm:"column:pernyataansdq16" json:"pernyataansdq16"`
	NilaiSdq16           *int       `gorm:"column:nilai_sdq16" json:"nilai_sdq16"`
	Pernyataansdq17      *string    `gorm:"column:pernyataansdq17" json:"pernyataansdq17"`
	NilaiSdq17           *int       `gorm:"column:nilai_sdq17" json:"nilai_sdq17"`
	Pernyataansdq18      *string    `gorm:"column:pernyataansdq18" json:"pernyataansdq18"`
	NilaiSdq18           *int       `gorm:"column:nilai_sdq18" json:"nilai_sdq18"`
	Pernyataansdq19      *string    `gorm:"column:pernyataansdq19" json:"pernyataansdq19"`
	NilaiSdq19           *int       `gorm:"column:nilai_sdq19" json:"nilai_sdq19"`
	Pernyataansdq20      *string    `gorm:"column:pernyataansdq20" json:"pernyataansdq20"`
	NilaiSdq20           *int       `gorm:"column:nilai_sdq20" json:"nilai_sdq20"`
	Pernyataansdq21      *string    `gorm:"column:pernyataansdq21" json:"pernyataansdq21"`
	NilaiSdq21           *int       `gorm:"column:nilai_sdq21" json:"nilai_sdq21"`
	Pernyataansdq22      *string    `gorm:"column:pernyataansdq22" json:"pernyataansdq22"`
	NilaiSdq22           *int       `gorm:"column:nilai_sdq22" json:"nilai_sdq22"`
	Pernyataansdq23      *string    `gorm:"column:pernyataansdq23" json:"pernyataansdq23"`
	NilaiSdq23           *int       `gorm:"column:nilai_sdq23" json:"nilai_sdq23"`
	Pernyataansdq24      *string    `gorm:"column:pernyataansdq24" json:"pernyataansdq24"`
	NilaiSdq24           *int       `gorm:"column:nilai_sdq24" json:"nilai_sdq24"`
	Pernyataansdq25      *string    `gorm:"column:pernyataansdq25" json:"pernyataansdq25"`
	NilaiSdq25           *int       `gorm:"column:nilai_sdq25" json:"nilai_sdq25"`
	NilaiTotalSdq        *int       `gorm:"column:nilai_total_sdq" json:"nilai_total_sdq"`
	GejalaEmosional      *string    `gorm:"column:gejala_emosional" json:"gejala_emosional"`
	NilaiGejalaEmosional *int       `gorm:"column:nilai_gejala_emosional" json:"nilai_gejala_emosional"`
	MasalahPerilaku      *string    `gorm:"column:masalah_perilaku" json:"masalah_perilaku"`
	NilaiMasalahPerilaku *int       `gorm:"column:nilai_masalah_perilaku" json:"nilai_masalah_perilaku"`
	Hiperaktivitas       *string    `gorm:"column:hiperaktivitas" json:"hiperaktivitas"`
	NilaiHiperaktivitas  *int       `gorm:"column:nilai_hiperaktivitas" json:"nilai_hiperaktivitas"`
	TemanSebaya          *string    `gorm:"column:teman_sebaya" json:"teman_sebaya"`
	NilaiTemanSebaya     *int       `gorm:"column:nilai_teman_sebaya" json:"nilai_teman_sebaya"`
	Kekuatan             *string    `gorm:"column:kekuatan" json:"kekuatan"`
	NilaiKekuatan        *int       `gorm:"column:nilai_kekuatan" json:"nilai_kekuatan"`
	Kesulitan            *string    `gorm:"column:kesulitan" json:"kesulitan"`
	NilaiKesulitan       *int       `gorm:"column:nilai_kesulitan" json:"nilai_kesulitan"`
	Keterangan           *string    `gorm:"column:keterangan" json:"keterangan"`
}

func (SkriningInstrumenSdq) TableName() string {
	return "skrining_instrumen_sdq"
}
