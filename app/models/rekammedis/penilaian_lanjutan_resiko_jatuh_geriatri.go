package rekammedis

import "time"

// PenilaianLanjutanResikoJatuhGeriatri tabel `penilaian_lanjutan_resiko_jatuh_geriatri` (penilaian lanjutan risiko jatuh geriatri, RMPenilaianLanjutanRisikoJatuhGeriatri).
type PenilaianLanjutanResikoJatuhGeriatri struct {
	NoRawat                  string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                  *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	PenilaianJatuhSkala1     *string    `gorm:"column:penilaian_jatuh_skala1" json:"penilaian_jatuh_skala1"`
	PenilaianJatuhNilai1     *int       `gorm:"column:penilaian_jatuh_nilai1" json:"penilaian_jatuh_nilai1"`
	PenilaianJatuhSkala2     *string    `gorm:"column:penilaian_jatuh_skala2" json:"penilaian_jatuh_skala2"`
	PenilaianJatuhNilai2     *int       `gorm:"column:penilaian_jatuh_nilai2" json:"penilaian_jatuh_nilai2"`
	PenilaianJatuhSkala3     *string    `gorm:"column:penilaian_jatuh_skala3" json:"penilaian_jatuh_skala3"`
	PenilaianJatuhNilai3     *int       `gorm:"column:penilaian_jatuh_nilai3" json:"penilaian_jatuh_nilai3"`
	PenilaianJatuhSkala4     *string    `gorm:"column:penilaian_jatuh_skala4" json:"penilaian_jatuh_skala4"`
	PenilaianJatuhNilai4     *int       `gorm:"column:penilaian_jatuh_nilai4" json:"penilaian_jatuh_nilai4"`
	PenilaianJatuhSkala5     *string    `gorm:"column:penilaian_jatuh_skala5" json:"penilaian_jatuh_skala5"`
	PenilaianJatuhNilai5     *int       `gorm:"column:penilaian_jatuh_nilai5" json:"penilaian_jatuh_nilai5"`
	PenilaianJatuhSkala6     *string    `gorm:"column:penilaian_jatuh_skala6" json:"penilaian_jatuh_skala6"`
	PenilaianJatuhNilai6     *int       `gorm:"column:penilaian_jatuh_nilai6" json:"penilaian_jatuh_nilai6"`
	PenilaianJatuhSkala7     *string    `gorm:"column:penilaian_jatuh_skala7" json:"penilaian_jatuh_skala7"`
	PenilaianJatuhNilai7     *int       `gorm:"column:penilaian_jatuh_nilai7" json:"penilaian_jatuh_nilai7"`
	PenilaianJatuhSkala8     *string    `gorm:"column:penilaian_jatuh_skala8" json:"penilaian_jatuh_skala8"`
	PenilaianJatuhNilai8     *int       `gorm:"column:penilaian_jatuh_nilai8" json:"penilaian_jatuh_nilai8"`
	PenilaianJatuhSkala9     *string    `gorm:"column:penilaian_jatuh_skala9" json:"penilaian_jatuh_skala9"`
	PenilaianJatuhNilai9     *int       `gorm:"column:penilaian_jatuh_nilai9" json:"penilaian_jatuh_nilai9"`
	PenilaianJatuhSkala10    *string    `gorm:"column:penilaian_jatuh_skala10" json:"penilaian_jatuh_skala10"`
	PenilaianJatuhNilai10    *int       `gorm:"column:penilaian_jatuh_nilai10" json:"penilaian_jatuh_nilai10"`
	PenilaianJatuhSkala11    *string    `gorm:"column:penilaian_jatuh_skala11" json:"penilaian_jatuh_skala11"`
	PenilaianJatuhNilai11    *int       `gorm:"column:penilaian_jatuh_nilai11" json:"penilaian_jatuh_nilai11"`
	PenilaianJatuhTotalnilai *int       `gorm:"column:penilaian_jatuh_totalnilai" json:"penilaian_jatuh_totalnilai"`
	HasilSkrining            *string    `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Saran                    *string    `gorm:"column:saran" json:"saran"`
	Nip                      string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianLanjutanResikoJatuhGeriatri) TableName() string {
	return "penilaian_lanjutan_resiko_jatuh_geriatri"
}
