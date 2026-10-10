package rekammedis

import "time"

// HasilEndoskopiHidung tabel `hasil_endoskopi_hidung` (hasil endoskopi hidung, RMHasilEndoskopiHidung).
type HasilEndoskopiHidung struct {
	NoRawat            string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal            *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter           string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis     string     `gorm:"column:diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari        string     `gorm:"column:kiriman_dari" json:"kiriman_dari"`
	KondisiHidungKanan *string    `gorm:"column:kondisi_hidung_kanan" json:"kondisi_hidung_kanan"`
	KondisiHidungKiri  *string    `gorm:"column:kondisi_hidung_kiri" json:"kondisi_hidung_kiri"`
	KavumNasiKanan     *string    `gorm:"column:kavum_nasi_kanan" json:"kavum_nasi_kanan"`
	KavumNasiKiri      *string    `gorm:"column:kavum_nasi_kiri" json:"kavum_nasi_kiri"`
	KonkaInferiorKanan *string    `gorm:"column:konka_inferior_kanan" json:"konka_inferior_kanan"`
	KonkaInferiorKiri  *string    `gorm:"column:konka_inferior_kiri" json:"konka_inferior_kiri"`
	MeatusMediusKanan  *string    `gorm:"column:meatus_medius_kanan" json:"meatus_medius_kanan"`
	MeatusMediusKiri   *string    `gorm:"column:meatus_medius_kiri" json:"meatus_medius_kiri"`
	SeptumKanan        *string    `gorm:"column:septum_kanan" json:"septum_kanan"`
	SeptumKiri         *string    `gorm:"column:septum_kiri" json:"septum_kiri"`
	NasofaringKanan    *string    `gorm:"column:nasofaring_kanan" json:"nasofaring_kanan"`
	NasofaringKiri     *string    `gorm:"column:nasofaring_kiri" json:"nasofaring_kiri"`
	LainlainKanan      *string    `gorm:"column:lainlain_kanan" json:"lainlain_kanan"`
	LainlainKiri       *string    `gorm:"column:lainlain_kiri" json:"lainlain_kiri"`
	Kesimpulan         *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
}

func (HasilEndoskopiHidung) TableName() string {
	return "hasil_endoskopi_hidung"
}
