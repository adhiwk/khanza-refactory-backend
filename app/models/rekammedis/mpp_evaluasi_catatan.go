package rekammedis

import "time"

// MppEvaluasiCatatan tabel `mpp_evaluasi_catatan` (skrining MPP form b, RMSkriningMPPFormB).
type MppEvaluasiCatatan struct {
	NoRawat         string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglImplementasi *time.Time `gorm:"column:tgl_implementasi;primaryKey;autoIncrement:false" json:"tgl_implementasi"`
	Masalah         *string    `gorm:"column:masalah" json:"masalah"`
	Tinjut          *string    `gorm:"column:tinjut" json:"tinjut"`
	Evaluasi        *string    `gorm:"column:evaluasi" json:"evaluasi"`
	Nip             string     `gorm:"column:nip" json:"nip"`
}

func (MppEvaluasiCatatan) TableName() string {
	return "mpp_evaluasi_catatan"
}
