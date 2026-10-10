package rekammedis

import "time"

// PenilaianRisikoJatuhNeonatus tabel `penilaian_risiko_jatuh_neonatus` (penilaian risiko jatuh neonatus, RMPenilaianRisikoJatuhNeonatus).
type PenilaianRisikoJatuhNeonatus struct {
	NoRawat     string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal     *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Intervensi1 *string    `gorm:"column:intervensi1" json:"intervensi1"`
	Intervensi2 *string    `gorm:"column:intervensi2" json:"intervensi2"`
	Intervensi3 *string    `gorm:"column:intervensi3" json:"intervensi3"`
	Intervensi4 *string    `gorm:"column:intervensi4" json:"intervensi4"`
	Intervensi5 *string    `gorm:"column:intervensi5" json:"intervensi5"`
	Intervensi6 *string    `gorm:"column:intervensi6" json:"intervensi6"`
	Intervensi7 *string    `gorm:"column:intervensi7" json:"intervensi7"`
	Intervensi8 *string    `gorm:"column:intervensi8" json:"intervensi8"`
	Intervensi9 *string    `gorm:"column:intervensi9" json:"intervensi9"`
	Edukasi1    *string    `gorm:"column:edukasi1" json:"edukasi1"`
	Edukasi2    *string    `gorm:"column:edukasi2" json:"edukasi2"`
	Edukasi3    *string    `gorm:"column:edukasi3" json:"edukasi3"`
	Edukasi4    *string    `gorm:"column:edukasi4" json:"edukasi4"`
	Edukasi5    *string    `gorm:"column:edukasi5" json:"edukasi5"`
	Sasaran1    *string    `gorm:"column:sasaran1" json:"sasaran1"`
	Sasaran2    *string    `gorm:"column:sasaran2" json:"sasaran2"`
	Sasaran3    *string    `gorm:"column:sasaran3" json:"sasaran3"`
	Sasaran4    *string    `gorm:"column:sasaran4" json:"sasaran4"`
	Evaluasi1   *string    `gorm:"column:evaluasi1" json:"evaluasi1"`
	Evaluasi2   *string    `gorm:"column:evaluasi2" json:"evaluasi2"`
	Evaluasi3   *string    `gorm:"column:evaluasi3" json:"evaluasi3"`
	Nip         string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianRisikoJatuhNeonatus) TableName() string {
	return "penilaian_risiko_jatuh_neonatus"
}
