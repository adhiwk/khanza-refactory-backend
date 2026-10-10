package rekammedis

import "time"

// PenilaianDehidrasi tabel `penilaian_dehidrasi` (penilaian derajat dehidrasi, RMPenilaianDerajatDehidrasi).
type PenilaianDehidrasi struct {
	NoRawat             string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal             *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Penilaian1          *string    `gorm:"column:penilaian1" json:"penilaian1"`
	PenilaianNilai1     *int       `gorm:"column:penilaian_nilai1" json:"penilaian_nilai1"`
	Penilaian2          *string    `gorm:"column:penilaian2" json:"penilaian2"`
	PenilaianNilai2     *int       `gorm:"column:penilaian_nilai2" json:"penilaian_nilai2"`
	Penilaian3          *string    `gorm:"column:penilaian3" json:"penilaian3"`
	PenilaianNilai3     *int       `gorm:"column:penilaian_nilai3" json:"penilaian_nilai3"`
	Penilaian4          *string    `gorm:"column:penilaian4" json:"penilaian4"`
	PenilaianNilai4     *int       `gorm:"column:penilaian_nilai4" json:"penilaian_nilai4"`
	Penilaian5          *string    `gorm:"column:penilaian5" json:"penilaian5"`
	PenilaianNilai5     *int       `gorm:"column:penilaian_nilai5" json:"penilaian_nilai5"`
	Penilaian6          *string    `gorm:"column:penilaian6" json:"penilaian6"`
	PenilaianNilai6     *int       `gorm:"column:penilaian_nilai6" json:"penilaian_nilai6"`
	PenilaianTotalnilai *int       `gorm:"column:penilaian_totalnilai" json:"penilaian_totalnilai"`
	HasilPenilaian      *string    `gorm:"column:hasil_penilaian" json:"hasil_penilaian"`
	KdDokter            string     `gorm:"column:kd_dokter" json:"kd_dokter"`
}

func (PenilaianDehidrasi) TableName() string {
	return "penilaian_dehidrasi"
}
