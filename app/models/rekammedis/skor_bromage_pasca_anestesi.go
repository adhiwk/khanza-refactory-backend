package rekammedis

import "time"

// SkorBromagePascaAnestesi tabel `skor_bromage_pasca_anestesi` (monitoring bromage pasca anestesi, RMMonitoringBromagePascaAnestesi).
type SkorBromagePascaAnestesi struct {
	NoRawat         string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal         *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	PenilaianSkala1 *string    `gorm:"column:penilaian_skala1" json:"penilaian_skala1"`
	PenilaianNilai1 *int       `gorm:"column:penilaian_nilai1" json:"penilaian_nilai1"`
	Keluar          *string    `gorm:"column:keluar" json:"keluar"`
	Instruksi       *string    `gorm:"column:instruksi" json:"instruksi"`
	KdDokter        string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Nip             string     `gorm:"column:nip" json:"nip"`
}

func (SkorBromagePascaAnestesi) TableName() string {
	return "skor_bromage_pasca_anestesi"
}
