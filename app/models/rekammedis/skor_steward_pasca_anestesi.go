package rekammedis

import "time"

// SkorStewardPascaAnestesi tabel `skor_steward_pasca_anestesi` (monitoring steward pasca anestesi, RMMonitoringStewardPascaAnestesi).
type SkorStewardPascaAnestesi struct {
	NoRawat             string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal             *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	PenilaianSkala1     *string    `gorm:"column:penilaian_skala1" json:"penilaian_skala1"`
	PenilaianNilai1     *int       `gorm:"column:penilaian_nilai1" json:"penilaian_nilai1"`
	PenilaianSkala2     *string    `gorm:"column:penilaian_skala2" json:"penilaian_skala2"`
	PenilaianNilai2     *int       `gorm:"column:penilaian_nilai2" json:"penilaian_nilai2"`
	PenilaianSkala3     *string    `gorm:"column:penilaian_skala3" json:"penilaian_skala3"`
	PenilaianNilai3     *int       `gorm:"column:penilaian_nilai3" json:"penilaian_nilai3"`
	PenilaianTotalnilai *int       `gorm:"column:penilaian_totalnilai" json:"penilaian_totalnilai"`
	Keluar              *string    `gorm:"column:keluar" json:"keluar"`
	Instruksi           *string    `gorm:"column:instruksi" json:"instruksi"`
	KdDokter            string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Nip                 string     `gorm:"column:nip" json:"nip"`
}

func (SkorStewardPascaAnestesi) TableName() string {
	return "skor_steward_pasca_anestesi"
}
