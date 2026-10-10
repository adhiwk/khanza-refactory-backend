package rekammedis

import "time"

// HasilPemeriksaanSlitLamp tabel `hasil_pemeriksaan_slit_lamp` (hasil pemeriksaan slit lamp, RMHasilPemeriksaanSlitLamp).
type HasilPemeriksaanSlitLamp struct {
	NoRawat          string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal          *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter         string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis   *string    `gorm:"column:diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari      *string    `gorm:"column:kiriman_dari" json:"kiriman_dari"`
	HasilPemeriksaan *string    `gorm:"column:hasil_pemeriksaan" json:"hasil_pemeriksaan"`
}

func (HasilPemeriksaanSlitLamp) TableName() string {
	return "hasil_pemeriksaan_slit_lamp"
}
