package rekammedis

import "time"

// MppEvaluasi tabel `mpp_evaluasi` (skrining MPP form a, RMSkriningMPPFormA).
type MppEvaluasi struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal      *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	KdDokter     *string    `gorm:"column:kd_dokter" json:"kd_dokter"`
	KdKonsulan   *string    `gorm:"column:kd_konsulan" json:"kd_konsulan"`
	Diagnosis    string     `gorm:"column:diagnosis" json:"diagnosis"`
	Kelompok     string     `gorm:"column:kelompok" json:"kelompok"`
	Assesmen     string     `gorm:"column:assesmen" json:"assesmen"`
	Identifikasi string     `gorm:"column:identifikasi" json:"identifikasi"`
	Rencana      string     `gorm:"column:rencana" json:"rencana"`
	Nip          string     `gorm:"column:nip" json:"nip"`
}

func (MppEvaluasi) TableName() string {
	return "mpp_evaluasi"
}
