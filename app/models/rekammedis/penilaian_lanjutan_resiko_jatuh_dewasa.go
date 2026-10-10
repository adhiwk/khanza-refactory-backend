package rekammedis

import "time"

// PenilaianLanjutanResikoJatuhDewasa tabel `penilaian_lanjutan_resiko_jatuh_dewasa` (penilaian lanjutan risiko jatuh dewasa, RMPenilaianLanjutanRisikoJatuhDewasa).
type PenilaianLanjutanResikoJatuhDewasa struct {
	NoRawat                       string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                       *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	PenilaianJatuhmorseSkala1     *string    `gorm:"column:penilaian_jatuhmorse_skala1" json:"penilaian_jatuhmorse_skala1"`
	PenilaianJatuhmorseNilai1     *int       `gorm:"column:penilaian_jatuhmorse_nilai1" json:"penilaian_jatuhmorse_nilai1"`
	PenilaianJatuhmorseSkala2     *string    `gorm:"column:penilaian_jatuhmorse_skala2" json:"penilaian_jatuhmorse_skala2"`
	PenilaianJatuhmorseNilai2     *int       `gorm:"column:penilaian_jatuhmorse_nilai2" json:"penilaian_jatuhmorse_nilai2"`
	PenilaianJatuhmorseSkala3     *string    `gorm:"column:penilaian_jatuhmorse_skala3" json:"penilaian_jatuhmorse_skala3"`
	PenilaianJatuhmorseNilai3     *int       `gorm:"column:penilaian_jatuhmorse_nilai3" json:"penilaian_jatuhmorse_nilai3"`
	PenilaianJatuhmorseSkala4     *string    `gorm:"column:penilaian_jatuhmorse_skala4" json:"penilaian_jatuhmorse_skala4"`
	PenilaianJatuhmorseNilai4     *int       `gorm:"column:penilaian_jatuhmorse_nilai4" json:"penilaian_jatuhmorse_nilai4"`
	PenilaianJatuhmorseSkala5     *string    `gorm:"column:penilaian_jatuhmorse_skala5" json:"penilaian_jatuhmorse_skala5"`
	PenilaianJatuhmorseNilai5     *int       `gorm:"column:penilaian_jatuhmorse_nilai5" json:"penilaian_jatuhmorse_nilai5"`
	PenilaianJatuhmorseSkala6     *string    `gorm:"column:penilaian_jatuhmorse_skala6" json:"penilaian_jatuhmorse_skala6"`
	PenilaianJatuhmorseNilai6     *int       `gorm:"column:penilaian_jatuhmorse_nilai6" json:"penilaian_jatuhmorse_nilai6"`
	PenilaianJatuhmorseTotalnilai *int       `gorm:"column:penilaian_jatuhmorse_totalnilai" json:"penilaian_jatuhmorse_totalnilai"`
	HasilSkrining                 *string    `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Saran                         *string    `gorm:"column:saran" json:"saran"`
	Nip                           string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianLanjutanResikoJatuhDewasa) TableName() string {
	return "penilaian_lanjutan_resiko_jatuh_dewasa"
}
