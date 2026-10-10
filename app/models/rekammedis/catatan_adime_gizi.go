package rekammedis

import "time"

// CatatanAdimeGizi tabel `catatan_adime_gizi` (catatan ADIME gizi, RMCatatanADIMEGizi).
type CatatanAdimeGizi struct {
	NoRawat    string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal    *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Asesmen    *string    `gorm:"column:asesmen" json:"asesmen"`
	Diagnosis  *string    `gorm:"column:diagnosis" json:"diagnosis"`
	Intervensi *string    `gorm:"column:intervensi" json:"intervensi"`
	Monitoring *string    `gorm:"column:monitoring" json:"monitoring"`
	Evaluasi   *string    `gorm:"column:evaluasi" json:"evaluasi"`
	Instruksi  *string    `gorm:"column:instruksi" json:"instruksi"`
	Nip        *string    `gorm:"column:nip" json:"nip"`
}

func (CatatanAdimeGizi) TableName() string {
	return "catatan_adime_gizi"
}
