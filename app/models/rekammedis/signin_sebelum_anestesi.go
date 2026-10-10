package rekammedis

import "time"

// SigninSebelumAnestesi tabel `signin_sebelum_anestesi` (sign in sebelum anastesi, RMSignInSebelumAnastesi).
type SigninSebelumAnestesi struct {
	NoRawat                                   string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                                   *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Sncn                                      string     `gorm:"column:sncn" json:"sncn"`
	Tindakan                                  string     `gorm:"column:tindakan" json:"tindakan"`
	KdDokterBedah                             string     `gorm:"column:kd_dokter_bedah" json:"kd_dokter_bedah"`
	KdDokterAnestesi                          string     `gorm:"column:kd_dokter_anestesi" json:"kd_dokter_anestesi"`
	Identitas                                 *string    `gorm:"column:identitas" json:"identitas"`
	PenandaanAreaOperasi                      *string    `gorm:"column:penandaan_area_operasi" json:"penandaan_area_operasi"`
	Alergi                                    *string    `gorm:"column:alergi" json:"alergi"`
	ResikoAspirasi                            *string    `gorm:"column:resiko_aspirasi" json:"resiko_aspirasi"`
	ResikoAspirasiRencanaAntisipasi           *string    `gorm:"column:resiko_aspirasi_rencana_antisipasi" json:"resiko_aspirasi_rencana_antisipasi"`
	ResikoKehilanganDarah                     *string    `gorm:"column:resiko_kehilangan_darah" json:"resiko_kehilangan_darah"`
	ResikoKehilanganDarahLine                 *string    `gorm:"column:resiko_kehilangan_darah_line" json:"resiko_kehilangan_darah_line"`
	ResikoKehilanganDarahRencanaAntisipasi    *string    `gorm:"column:resiko_kehilangan_darah_rencana_antisipasi" json:"resiko_kehilangan_darah_rencana_antisipasi"`
	KesiapanAlatObatAnestesi                  *string    `gorm:"column:kesiapan_alat_obat_anestesi" json:"kesiapan_alat_obat_anestesi"`
	KesiapanAlatObatAnestesiRencanaAntisipasi *string    `gorm:"column:kesiapan_alat_obat_anestesi_rencana_antisipasi" json:"kesiapan_alat_obat_anestesi_rencana_antisipasi"`
	NipPerawatOk                              *string    `gorm:"column:nip_perawat_ok" json:"nip_perawat_ok"`
}

func (SigninSebelumAnestesi) TableName() string {
	return "signin_sebelum_anestesi"
}
