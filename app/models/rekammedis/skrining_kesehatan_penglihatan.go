package rekammedis

import "time"

// SkriningKesehatanPenglihatan tabel `skrining_kesehatan_penglihatan` (skrining kesehatan penglihatan, RMSkriningKesehatanPenglihatan).
type SkriningKesehatanPenglihatan struct {
	NoRawat        string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal        *time.Time `gorm:"column:tanggal" json:"tanggal"`
	MataLuar       *string    `gorm:"column:mata_luar" json:"mata_luar"`
	TajamKiri      *string    `gorm:"column:tajam_kiri" json:"tajam_kiri"`
	TajamKanan     *string    `gorm:"column:tajam_kanan" json:"tajam_kanan"`
	ButaWarnaKiri  *string    `gorm:"column:buta_warna_kiri" json:"buta_warna_kiri"`
	ButaWarnaKanan *string    `gorm:"column:buta_warna_kanan" json:"buta_warna_kanan"`
	Kacamata       *string    `gorm:"column:kacamata" json:"kacamata"`
	VisusKiri      *string    `gorm:"column:visus_kiri" json:"visus_kiri"`
	VisusKanan     *string    `gorm:"column:visus_kanan" json:"visus_kanan"`
	RefraksiKiri   *string    `gorm:"column:refraksi_kiri" json:"refraksi_kiri"`
	RefraksiKanan  *string    `gorm:"column:refraksi_kanan" json:"refraksi_kanan"`
	RujukRefraksi  *string    `gorm:"column:rujuk_refraksi" json:"rujuk_refraksi"`
	KatarakKiri    *string    `gorm:"column:katarak_kiri" json:"katarak_kiri"`
	KatarakKanan   *string    `gorm:"column:katarak_kanan" json:"katarak_kanan"`
	RujukKatarak   *string    `gorm:"column:rujuk_katarak" json:"rujuk_katarak"`
	HasilSkrining  string     `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Keterangan     *string    `gorm:"column:keterangan" json:"keterangan"`
	Nip            string     `gorm:"column:nip" json:"nip"`
}

func (SkriningKesehatanPenglihatan) TableName() string {
	return "skrining_kesehatan_penglihatan"
}
