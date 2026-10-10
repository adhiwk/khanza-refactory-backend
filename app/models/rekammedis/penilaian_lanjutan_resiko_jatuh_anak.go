package rekammedis

import "time"

// PenilaianLanjutanResikoJatuhAnak tabel `penilaian_lanjutan_resiko_jatuh_anak` (penilaian lanjutan risiko jatuh anak, RMPenilaianLanjutanRisikoJatuhAnak).
type PenilaianLanjutanResikoJatuhAnak struct {
	NoRawat                         string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                         *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	PenilaianHumptydumptySkala1     *string    `gorm:"column:penilaian_humptydumpty_skala1" json:"penilaian_humptydumpty_skala1"`
	PenilaianHumptydumptyNilai1     *int       `gorm:"column:penilaian_humptydumpty_nilai1" json:"penilaian_humptydumpty_nilai1"`
	PenilaianHumptydumptySkala2     *string    `gorm:"column:penilaian_humptydumpty_skala2" json:"penilaian_humptydumpty_skala2"`
	PenilaianHumptydumptyNilai2     *int       `gorm:"column:penilaian_humptydumpty_nilai2" json:"penilaian_humptydumpty_nilai2"`
	PenilaianHumptydumptySkala3     *string    `gorm:"column:penilaian_humptydumpty_skala3" json:"penilaian_humptydumpty_skala3"`
	PenilaianHumptydumptyNilai3     *int       `gorm:"column:penilaian_humptydumpty_nilai3" json:"penilaian_humptydumpty_nilai3"`
	PenilaianHumptydumptySkala4     *string    `gorm:"column:penilaian_humptydumpty_skala4" json:"penilaian_humptydumpty_skala4"`
	PenilaianHumptydumptyNilai4     *int       `gorm:"column:penilaian_humptydumpty_nilai4" json:"penilaian_humptydumpty_nilai4"`
	PenilaianHumptydumptySkala5     *string    `gorm:"column:penilaian_humptydumpty_skala5" json:"penilaian_humptydumpty_skala5"`
	PenilaianHumptydumptyNilai5     *int       `gorm:"column:penilaian_humptydumpty_nilai5" json:"penilaian_humptydumpty_nilai5"`
	PenilaianHumptydumptySkala6     *string    `gorm:"column:penilaian_humptydumpty_skala6" json:"penilaian_humptydumpty_skala6"`
	PenilaianHumptydumptyNilai6     *int       `gorm:"column:penilaian_humptydumpty_nilai6" json:"penilaian_humptydumpty_nilai6"`
	PenilaianHumptydumptySkala7     *string    `gorm:"column:penilaian_humptydumpty_skala7" json:"penilaian_humptydumpty_skala7"`
	PenilaianHumptydumptyNilai7     *int       `gorm:"column:penilaian_humptydumpty_nilai7" json:"penilaian_humptydumpty_nilai7"`
	PenilaianHumptydumptyTotalnilai *int       `gorm:"column:penilaian_humptydumpty_totalnilai" json:"penilaian_humptydumpty_totalnilai"`
	HasilSkrining                   *string    `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Saran                           *string    `gorm:"column:saran" json:"saran"`
	Nip                             string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianLanjutanResikoJatuhAnak) TableName() string {
	return "penilaian_lanjutan_resiko_jatuh_anak"
}
