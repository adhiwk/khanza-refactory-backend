package rekammedis

import "time"

// PenilaianPasienPenyakitMenular tabel `penilaian_pasien_penyakit_menular` (penilaian pasien penyakit menular, RMPenilaianPasienPenyakitMenular).
type PenilaianPasienPenyakitMenular struct {
	NoRawat                              string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                              *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter                             string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Anamnesis                            string     `gorm:"column:anamnesis" json:"anamnesis"`
	Hubungan                             string     `gorm:"column:hubungan" json:"hubungan"`
	PasienMengetahuiKondisiPenyakitnya   *string    `gorm:"column:pasien_mengetahui_kondisi_penyakitnya" json:"pasien_mengetahui_kondisi_penyakitnya"`
	PenyakitSamaSerumah                  *string    `gorm:"column:penyakit_sama_serumah" json:"penyakit_sama_serumah"`
	RiwayatKontak                        *string    `gorm:"column:riwayat_kontak" json:"riwayat_kontak"`
	KeteranganRiwayatKontak              *string    `gorm:"column:keterangan_riwayat_kontak" json:"keterangan_riwayat_kontak"`
	TransmisiPenularanPenyakit           *string    `gorm:"column:transmisi_penularan_penyakit" json:"transmisi_penularan_penyakit"`
	KeteranganTransmisiPenularanPenyakit *string    `gorm:"column:keterangan_transmisi_penularan_penyakit" json:"keterangan_transmisi_penularan_penyakit"`
	KebutuhanRuangRawat                  *string    `gorm:"column:kebutuhan_ruang_rawat" json:"kebutuhan_ruang_rawat"`
	KeluhanYangDirasakanSaatIni          *string    `gorm:"column:keluhan_yang_dirasakan_saat_ini" json:"keluhan_yang_dirasakan_saat_ini"`
	RiwayatPenyakitKeluarga              *string    `gorm:"column:riwayat_penyakit_keluarga" json:"riwayat_penyakit_keluarga"`
	RiwayatAlergi                        *string    `gorm:"column:riwayat_alergi" json:"riwayat_alergi"`
	RiwayatVaksinasi                     *string    `gorm:"column:riwayat_vaksinasi" json:"riwayat_vaksinasi"`
	RiwayatPengobatan                    *string    `gorm:"column:riwayat_pengobatan" json:"riwayat_pengobatan"`
	DiagnosaUtama                        *string    `gorm:"column:diagnosa_utama" json:"diagnosa_utama"`
	DiagnosaTambahan                     *string    `gorm:"column:diagnosa_tambahan" json:"diagnosa_tambahan"`
}

func (PenilaianPasienPenyakitMenular) TableName() string {
	return "penilaian_pasien_penyakit_menular"
}
