package rekammedis

import "time"

// PenilaianLanjutanSkriningFungsional tabel `penilaian_lanjutan_skrining_fungsional` (penilaian lanjutan skrining fungsional, RMPenilaianLanjutanSkriningFungsional).
type PenilaianLanjutanSkriningFungsional struct {
	NoRawat                     string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                     *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	PenilaianSkriningSkala1     *string    `gorm:"column:penilaian_skrining_skala1" json:"penilaian_skrining_skala1"`
	PenilaianSkriningNilai1     *int       `gorm:"column:penilaian_skrining_nilai1" json:"penilaian_skrining_nilai1"`
	PenilaianSkriningSkala2     *string    `gorm:"column:penilaian_skrining_skala2" json:"penilaian_skrining_skala2"`
	PenilaianSkriningNilai2     *int       `gorm:"column:penilaian_skrining_nilai2" json:"penilaian_skrining_nilai2"`
	PenilaianSkriningSkala3     *string    `gorm:"column:penilaian_skrining_skala3" json:"penilaian_skrining_skala3"`
	PenilaianSkriningNilai3     *int       `gorm:"column:penilaian_skrining_nilai3" json:"penilaian_skrining_nilai3"`
	PenilaianSkriningSkala4     *string    `gorm:"column:penilaian_skrining_skala4" json:"penilaian_skrining_skala4"`
	PenilaianSkriningNilai4     *int       `gorm:"column:penilaian_skrining_nilai4" json:"penilaian_skrining_nilai4"`
	PenilaianSkriningSkala5     *string    `gorm:"column:penilaian_skrining_skala5" json:"penilaian_skrining_skala5"`
	PenilaianSkriningNilai5     *int       `gorm:"column:penilaian_skrining_nilai5" json:"penilaian_skrining_nilai5"`
	PenilaianSkriningSkala6     *string    `gorm:"column:penilaian_skrining_skala6" json:"penilaian_skrining_skala6"`
	PenilaianSkriningNilai6     *int       `gorm:"column:penilaian_skrining_nilai6" json:"penilaian_skrining_nilai6"`
	PenilaianSkriningSkala7     *string    `gorm:"column:penilaian_skrining_skala7" json:"penilaian_skrining_skala7"`
	PenilaianSkriningNilai7     *int       `gorm:"column:penilaian_skrining_nilai7" json:"penilaian_skrining_nilai7"`
	PenilaianSkriningSkala8     *string    `gorm:"column:penilaian_skrining_skala8" json:"penilaian_skrining_skala8"`
	PenilaianSkriningNilai8     *int       `gorm:"column:penilaian_skrining_nilai8" json:"penilaian_skrining_nilai8"`
	PenilaianSkriningSkala9     *string    `gorm:"column:penilaian_skrining_skala9" json:"penilaian_skrining_skala9"`
	PenilaianSkriningNilai9     *int       `gorm:"column:penilaian_skrining_nilai9" json:"penilaian_skrining_nilai9"`
	PenilaianSkriningSkala10    *string    `gorm:"column:penilaian_skrining_skala10" json:"penilaian_skrining_skala10"`
	PenilaianSkriningNilai10    *int       `gorm:"column:penilaian_skrining_nilai10" json:"penilaian_skrining_nilai10"`
	PenilaianSkriningTotalnilai *int       `gorm:"column:penilaian_skrining_totalnilai" json:"penilaian_skrining_totalnilai"`
	Nip                         string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianLanjutanSkriningFungsional) TableName() string {
	return "penilaian_lanjutan_skrining_fungsional"
}
