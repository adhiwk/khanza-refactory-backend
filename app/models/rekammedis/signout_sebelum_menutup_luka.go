package rekammedis

import "time"

// SignoutSebelumMenutupLuka tabel `signout_sebelum_menutup_luka` (sign out sebelum menutup luka, RMSignOutSebelumMenutupLuka).
type SignoutSebelumMenutupLuka struct {
	NoRawat                          string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                          *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Sncn                             string     `gorm:"column:sncn" json:"sncn"`
	Tindakan                         string     `gorm:"column:tindakan" json:"tindakan"`
	KdDokterBedah                    string     `gorm:"column:kd_dokter_bedah" json:"kd_dokter_bedah"`
	KdDokterAnestesi                 string     `gorm:"column:kd_dokter_anestesi" json:"kd_dokter_anestesi"`
	VerbalTindakan                   *string    `gorm:"column:verbal_tindakan" json:"verbal_tindakan"`
	VerbalKelengkapanKasa            *string    `gorm:"column:verbal_kelengkapan_kasa" json:"verbal_kelengkapan_kasa"`
	VerbalInstrumen                  *string    `gorm:"column:verbal_instrumen" json:"verbal_instrumen"`
	VerbalAlatTajam                  *string    `gorm:"column:verbal_alat_tajam" json:"verbal_alat_tajam"`
	KelengkapanSpecimenLabel         *string    `gorm:"column:kelengkapan_specimen_label" json:"kelengkapan_specimen_label"`
	KelengkapanSpecimenFormulir      *string    `gorm:"column:kelengkapan_specimen_formulir" json:"kelengkapan_specimen_formulir"`
	PeninjauanKegiatanDokterBedah    *string    `gorm:"column:peninjauan_kegiatan_dokter_bedah" json:"peninjauan_kegiatan_dokter_bedah"`
	PeninjauanKegiatanDokterAnestesi *string    `gorm:"column:peninjauan_kegiatan_dokter_anestesi" json:"peninjauan_kegiatan_dokter_anestesi"`
	PeninjauanKegiatanPerawatKamarOk *string    `gorm:"column:peninjauan_kegiatan_perawat_kamar_ok" json:"peninjauan_kegiatan_perawat_kamar_ok"`
	PerhatianUtamaFasePemulihan      *string    `gorm:"column:perhatian_utama_fase_pemulihan" json:"perhatian_utama_fase_pemulihan"`
	NipPerawatOk                     *string    `gorm:"column:nip_perawat_ok" json:"nip_perawat_ok"`
}

func (SignoutSebelumMenutupLuka) TableName() string {
	return "signout_sebelum_menutup_luka"
}
