package rekammedis

import "time"

// HasilPemeriksaanEkg tabel `hasil_pemeriksaan_ekg` (hasil pemeriksaan EKG, RMHasilPemeriksaanEKG).
type HasilPemeriksaanEkg struct {
	NoRawat        string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal        *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter       string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis *string    `gorm:"column:diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari    *string    `gorm:"column:kiriman_dari" json:"kiriman_dari"`
	Irama          *string    `gorm:"column:irama" json:"irama"`
	LajuJantung    *string    `gorm:"column:laju_jantung" json:"laju_jantung"`
	Gelombangp     *string    `gorm:"column:gelombangp" json:"gelombangp"`
	Intervalpr     *string    `gorm:"column:intervalpr" json:"intervalpr"`
	Axis           *string    `gorm:"column:axis" json:"axis"`
	Kompleksqrs    *string    `gorm:"column:kompleksqrs" json:"kompleksqrs"`
	Segmenst       *string    `gorm:"column:segmenst" json:"segmenst"`
	Gelombangt     *string    `gorm:"column:gelombangt" json:"gelombangt"`
	Kesimpulan     *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
}

func (HasilPemeriksaanEkg) TableName() string {
	return "hasil_pemeriksaan_ekg"
}
