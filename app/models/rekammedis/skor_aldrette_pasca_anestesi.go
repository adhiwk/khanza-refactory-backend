package rekammedis

import "time"

// SkorAldrettePascaAnestesi tabel `skor_aldrette_pasca_anestesi` (monitoring aldrette pasca anestesi, RMMonitoringAldrettePascaAnestesi).
type SkorAldrettePascaAnestesi struct {
	NoRawat             string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal             *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	PenilaianSkala1     *string    `gorm:"column:penilaian_skala1" json:"penilaian_skala1"`
	PenilaianNilai1     *int       `gorm:"column:penilaian_nilai1" json:"penilaian_nilai1"`
	PenilaianSkala2     *string    `gorm:"column:penilaian_skala2" json:"penilaian_skala2"`
	PenilaianNilai2     *int       `gorm:"column:penilaian_nilai2" json:"penilaian_nilai2"`
	PenilaianSkala3     *string    `gorm:"column:penilaian_skala3" json:"penilaian_skala3"`
	PenilaianNilai3     *int       `gorm:"column:penilaian_nilai3" json:"penilaian_nilai3"`
	PenilaianSkala4     *string    `gorm:"column:penilaian_skala4" json:"penilaian_skala4"`
	PenilaianNilai4     *int       `gorm:"column:penilaian_nilai4" json:"penilaian_nilai4"`
	PenilaianSkala5     *string    `gorm:"column:penilaian_skala5" json:"penilaian_skala5"`
	PenilaianNilai5     *int       `gorm:"column:penilaian_nilai5" json:"penilaian_nilai5"`
	PenilaianTotalnilai *int       `gorm:"column:penilaian_totalnilai" json:"penilaian_totalnilai"`
	Keluar              *string    `gorm:"column:keluar" json:"keluar"`
	Instruksi           *string    `gorm:"column:instruksi" json:"instruksi"`
	KdDokter            string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Nip                 string     `gorm:"column:nip" json:"nip"`
}

func (SkorAldrettePascaAnestesi) TableName() string {
	return "skor_aldrette_pasca_anestesi"
}
