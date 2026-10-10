package rekammedis

import "time"

// HasilEndoskopiTelinga tabel `hasil_endoskopi_telinga` (hasil endoskopi telinga, RMHasilEndoskopiTelinga).
type HasilEndoskopiTelinga struct {
	NoRawat                                string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                                *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter                               string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis                         string     `gorm:"column:diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari                            string     `gorm:"column:kiriman_dari" json:"kiriman_dari"`
	BentukLiangTelingaKanan                *string    `gorm:"column:bentuk_liang_telinga_kanan" json:"bentuk_liang_telinga_kanan"`
	BentukLiangTelingaKiri                 *string    `gorm:"column:bentuk_liang_telinga_kiri" json:"bentuk_liang_telinga_kiri"`
	KondisiLiangTelingaKanan               *string    `gorm:"column:kondisi_liang_telinga_kanan" json:"kondisi_liang_telinga_kanan"`
	KeteranganKondisiLiangTelingaKanan     *string    `gorm:"column:keterangan_kondisi_liang_telinga_kanan" json:"keterangan_kondisi_liang_telinga_kanan"`
	KondisiLiangTelingaKiri                *string    `gorm:"column:kondisi_liang_telinga_kiri" json:"kondisi_liang_telinga_kiri"`
	KeteranganKondisiLiangTelingaKiri      *string    `gorm:"column:keterangan_kondisi_liang_telinga_kiri" json:"keterangan_kondisi_liang_telinga_kiri"`
	MembranTimpaniIntakKanan               *string    `gorm:"column:membran_timpani_intak_kanan" json:"membran_timpani_intak_kanan"`
	MembranTimpaniIntakKiri                *string    `gorm:"column:membran_timpani_intak_kiri" json:"membran_timpani_intak_kiri"`
	MembranTimpaniPerforasiKanan           *string    `gorm:"column:membran_timpani_perforasi_kanan" json:"membran_timpani_perforasi_kanan"`
	KeteranganMembranTimpaniPerforasiKanan *string    `gorm:"column:keterangan_membran_timpani_perforasi_kanan" json:"keterangan_membran_timpani_perforasi_kanan"`
	MembranTimpaniPerforasiKiri            *string    `gorm:"column:membran_timpani_perforasi_kiri" json:"membran_timpani_perforasi_kiri"`
	KeteranganMembranTimpaniPerforasiKiri  *string    `gorm:"column:keterangan_membran_timpani_perforasi_kiri" json:"keterangan_membran_timpani_perforasi_kiri"`
	KavumTimpaniMukosaKanan                *string    `gorm:"column:kavum_timpani_mukosa_kanan" json:"kavum_timpani_mukosa_kanan"`
	KavumTimpaniMukosaKiri                 *string    `gorm:"column:kavum_timpani_mukosa_kiri" json:"kavum_timpani_mukosa_kiri"`
	KavumTimpaniOsikelKanan                *string    `gorm:"column:kavum_timpani_osikel_kanan" json:"kavum_timpani_osikel_kanan"`
	KavumTimpaniOsikelKiri                 *string    `gorm:"column:kavum_timpani_osikel_kiri" json:"kavum_timpani_osikel_kiri"`
	KavumTimpaniIsthmusKanan               *string    `gorm:"column:kavum_timpani_isthmus_kanan" json:"kavum_timpani_isthmus_kanan"`
	KavumTimpaniIsthmusKiri                *string    `gorm:"column:kavum_timpani_isthmus_kiri" json:"kavum_timpani_isthmus_kiri"`
	KavumTimpaniAnteriorKanan              *string    `gorm:"column:kavum_timpani_anterior_kanan" json:"kavum_timpani_anterior_kanan"`
	KavumTimpaniAnteriorKiri               *string    `gorm:"column:kavum_timpani_anterior_kiri" json:"kavum_timpani_anterior_kiri"`
	KavumTimpaniPosteriorKanan             *string    `gorm:"column:kavum_timpani_posterior_kanan" json:"kavum_timpani_posterior_kanan"`
	KavumTimpaniPosteriorKiri              *string    `gorm:"column:kavum_timpani_posterior_kiri" json:"kavum_timpani_posterior_kiri"`
	LainlainKanan                          *string    `gorm:"column:lainlain_kanan" json:"lainlain_kanan"`
	LainlainKiri                           *string    `gorm:"column:lainlain_kiri" json:"lainlain_kiri"`
	Kesimpulan                             *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
	Anjuran                                *string    `gorm:"column:anjuran" json:"anjuran"`
}

func (HasilEndoskopiTelinga) TableName() string {
	return "hasil_endoskopi_telinga"
}
