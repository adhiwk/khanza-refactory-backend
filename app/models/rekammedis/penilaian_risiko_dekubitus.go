package rekammedis

import "time"

// PenilaianRisikoDekubitus tabel `penilaian_risiko_dekubitus` (penilaian risiko dekubitus, RMPenilaianRisikoDekubitus).
type PenilaianRisikoDekubitus struct {
	NoRawat            string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal            *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	KondisiFisik       *string    `gorm:"column:kondisi_fisik" json:"kondisi_fisik"`
	KondisiFisikNilai  *int       `gorm:"column:kondisi_fisik_nilai" json:"kondisi_fisik_nilai"`
	StatusMental       *string    `gorm:"column:status_mental" json:"status_mental"`
	StatusMentalNilai  *int       `gorm:"column:status_mental_nilai" json:"status_mental_nilai"`
	Aktifitas          *string    `gorm:"column:aktifitas" json:"aktifitas"`
	AktifitasNilai     *int       `gorm:"column:aktifitas_nilai" json:"aktifitas_nilai"`
	Mobilitas          *string    `gorm:"column:mobilitas" json:"mobilitas"`
	MobilitasNilai     *int       `gorm:"column:mobilitas_nilai" json:"mobilitas_nilai"`
	Inkontinensia      *string    `gorm:"column:inkontinensia" json:"inkontinensia"`
	InkontinensiaNilai *int       `gorm:"column:inkontinensia_nilai" json:"inkontinensia_nilai"`
	Totalnilai         *int       `gorm:"column:totalnilai" json:"totalnilai"`
	Kategorinilai      *string    `gorm:"column:kategorinilai" json:"kategorinilai"`
	Nip                string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianRisikoDekubitus) TableName() string {
	return "penilaian_risiko_dekubitus"
}
