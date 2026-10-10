package rekammedis

import "time"

// PenilaianLanjutanResikoJatuhPsikiatri tabel `penilaian_lanjutan_resiko_jatuh_psikiatri` (penilaian lanjutan risiko jatuh psikiatri, RMPenilaianLanjutanRisikoJatuhPsikiatri).
type PenilaianLanjutanResikoJatuhPsikiatri struct {
	NoRawat                          string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                          *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	PenilaianJatuhedmonsonSkala1     *string    `gorm:"column:penilaian_jatuhedmonson_skala1" json:"penilaian_jatuhedmonson_skala1"`
	PenilaianJatuhedmonsonNilai1     *int       `gorm:"column:penilaian_jatuhedmonson_nilai1" json:"penilaian_jatuhedmonson_nilai1"`
	PenilaianJatuhedmonsonSkala2     *string    `gorm:"column:penilaian_jatuhedmonson_skala2" json:"penilaian_jatuhedmonson_skala2"`
	PenilaianJatuhedmonsonNilai2     *int       `gorm:"column:penilaian_jatuhedmonson_nilai2" json:"penilaian_jatuhedmonson_nilai2"`
	PenilaianJatuhedmonsonSkala3     *string    `gorm:"column:penilaian_jatuhedmonson_skala3" json:"penilaian_jatuhedmonson_skala3"`
	PenilaianJatuhedmonsonNilai3     *int       `gorm:"column:penilaian_jatuhedmonson_nilai3" json:"penilaian_jatuhedmonson_nilai3"`
	PenilaianJatuhedmonsonSkala4     *string    `gorm:"column:penilaian_jatuhedmonson_skala4" json:"penilaian_jatuhedmonson_skala4"`
	PenilaianJatuhedmonsonNilai4     *int       `gorm:"column:penilaian_jatuhedmonson_nilai4" json:"penilaian_jatuhedmonson_nilai4"`
	PenilaianJatuhedmonsonSkala5     *string    `gorm:"column:penilaian_jatuhedmonson_skala5" json:"penilaian_jatuhedmonson_skala5"`
	PenilaianJatuhedmonsonNilai5     *int       `gorm:"column:penilaian_jatuhedmonson_nilai5" json:"penilaian_jatuhedmonson_nilai5"`
	PenilaianJatuhedmonsonSkala6     *string    `gorm:"column:penilaian_jatuhedmonson_skala6" json:"penilaian_jatuhedmonson_skala6"`
	PenilaianJatuhedmonsonNilai6     *int       `gorm:"column:penilaian_jatuhedmonson_nilai6" json:"penilaian_jatuhedmonson_nilai6"`
	PenilaianJatuhedmonsonTotalnilai *int       `gorm:"column:penilaian_jatuhedmonson_totalnilai" json:"penilaian_jatuhedmonson_totalnilai"`
	HasilSkrining                    *string    `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Saran                            *string    `gorm:"column:saran" json:"saran"`
	Nip                              string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianLanjutanResikoJatuhPsikiatri) TableName() string {
	return "penilaian_lanjutan_resiko_jatuh_psikiatri"
}
