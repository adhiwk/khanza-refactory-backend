package rekammedis

import "time"

// SkriningHipertensi tabel `skrining_hipertensi` (skrining hipertensi, RMSkriningHipertensi).
type SkriningHipertensi struct {
	NoRawat               string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal               *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Anamnesis1            *string    `gorm:"column:anamnesis1" json:"anamnesis1"`
	Anamnesis2            *string    `gorm:"column:anamnesis2" json:"anamnesis2"`
	Anamnesis3            *string    `gorm:"column:anamnesis3" json:"anamnesis3"`
	Anamnesis4            *string    `gorm:"column:anamnesis4" json:"anamnesis4"`
	Anamnesis5            *string    `gorm:"column:anamnesis5" json:"anamnesis5"`
	Anamnesis6            *string    `gorm:"column:anamnesis6" json:"anamnesis6"`
	Anamnesis7            *string    `gorm:"column:anamnesis7" json:"anamnesis7"`
	Anamnesis8            *string    `gorm:"column:anamnesis8" json:"anamnesis8"`
	Sistole               string     `gorm:"column:sistole" json:"sistole"`
	Diastole              string     `gorm:"column:diastole" json:"diastole"`
	KlasifikasiHipertensi string     `gorm:"column:klasifikasi_hipertensi" json:"klasifikasi_hipertensi"`
	HasilSkrining         string     `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Keterangan            *string    `gorm:"column:keterangan" json:"keterangan"`
	Nip                   string     `gorm:"column:nip" json:"nip"`
}

func (SkriningHipertensi) TableName() string {
	return "skrining_hipertensi"
}
