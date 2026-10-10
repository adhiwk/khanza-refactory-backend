package rekammedis

import "time"

// PelaksanaanInformasiEdukasi tabel `pelaksanaan_informasi_edukasi` (pelaksanaan informasi edukasi, RMPelaksanaanInformasiEdukasi).
type PelaksanaanInformasiEdukasi struct {
	NoRawat                 string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                 *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Nik                     string     `gorm:"column:nik" json:"nik"`
	MateriEdukasi           *string    `gorm:"column:materi_edukasi" json:"materi_edukasi"`
	Keterangan              *string    `gorm:"column:keterangan" json:"keterangan"`
	DiberikanPada           string     `gorm:"column:diberikan_pada" json:"diberikan_pada"`
	KeteranganDiberikanPada string     `gorm:"column:keterangan_diberikan_pada" json:"keterangan_diberikan_pada"`
	LamaEdukasi             *string    `gorm:"column:lama_edukasi" json:"lama_edukasi"`
	MetodeEdukasi           *string    `gorm:"column:metode_edukasi" json:"metode_edukasi"`
	HasilVerifikasi         *string    `gorm:"column:hasil_verifikasi" json:"hasil_verifikasi"`
	Status                  *string    `gorm:"column:status" json:"status"`
}

func (PelaksanaanInformasiEdukasi) TableName() string {
	return "pelaksanaan_informasi_edukasi"
}
