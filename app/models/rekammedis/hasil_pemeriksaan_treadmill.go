package rekammedis

import "time"

// HasilPemeriksaanTreadmill tabel `hasil_pemeriksaan_treadmill` (hasil pemeriksaan treadmill, RMHasilPemeriksaanTreadmill).
type HasilPemeriksaanTreadmill struct {
	NoRawat               string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal               *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter              string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	KirimanDari           *string    `gorm:"column:kiriman_dari" json:"kiriman_dari"`
	DiagnosaKlinis        *string    `gorm:"column:diagnosa_klinis" json:"diagnosa_klinis"`
	Protokol              *string    `gorm:"column:protokol" json:"protokol"`
	KeteranganProtokol    *string    `gorm:"column:keterangan_protokol" json:"keterangan_protokol"`
	TdAwal                *string    `gorm:"column:td_awal" json:"td_awal"`
	NadiAwal              *string    `gorm:"column:nadi_awal" json:"nadi_awal"`
	DenyutJantungMaksimal *string    `gorm:"column:denyut_jantung_maksimal" json:"denyut_jantung_maksimal"`
	HasilPemeriksaan      *string    `gorm:"column:hasil_pemeriksaan" json:"hasil_pemeriksaan"`
	TemuanEkg             *string    `gorm:"column:temuan_ekg" json:"temuan_ekg"`
	KapasitasFungsional   *string    `gorm:"column:kapasitas_fungsional" json:"kapasitas_fungsional"`
	Interpretasi          *string    `gorm:"column:interpretasi" json:"interpretasi"`
	Kesimpulan            *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
}

func (HasilPemeriksaanTreadmill) TableName() string {
	return "hasil_pemeriksaan_treadmill"
}
