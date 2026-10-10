package rekammedis

import "time"

// PenilaianPasienImunitasRendah tabel `penilaian_pasien_imunitas_rendah` (penilaian pasien imunitas rendah, RMPenilaianPasienImunitasRendah).
type PenilaianPasienImunitasRendah struct {
	NoRawat                            string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                            *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter                           string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Anamnesis                          *string    `gorm:"column:anamnesis" json:"anamnesis"`
	Hubungan                           string     `gorm:"column:hubungan" json:"hubungan"`
	PasienMengetahuiKondisiPenyakitnya *string    `gorm:"column:pasien_mengetahui_kondisi_penyakitnya" json:"pasien_mengetahui_kondisi_penyakitnya"`
	KebutuhanRuangPerawatan            *string    `gorm:"column:kebutuhan_ruang_perawatan" json:"kebutuhan_ruang_perawatan"`
	RiwayatPenyakitKeluhan             *string    `gorm:"column:riwayat_penyakit_keluhan" json:"riwayat_penyakit_keluhan"`
	RiwayatPenyakitKeluarga            *string    `gorm:"column:riwayat_penyakit_keluarga" json:"riwayat_penyakit_keluarga"`
	RiwayatAlergi                      *string    `gorm:"column:riwayat_alergi" json:"riwayat_alergi"`
	RiwayatVaksinasi                   *string    `gorm:"column:riwayat_vaksinasi" json:"riwayat_vaksinasi"`
	RiwayatPengobatan                  *string    `gorm:"column:riwayat_pengobatan" json:"riwayat_pengobatan"`
	DiagnosaUtama                      *string    `gorm:"column:diagnosa_utama" json:"diagnosa_utama"`
	DiagnosaTambahan                   *string    `gorm:"column:diagnosa_tambahan" json:"diagnosa_tambahan"`
}

func (PenilaianPasienImunitasRendah) TableName() string {
	return "penilaian_pasien_imunitas_rendah"
}
