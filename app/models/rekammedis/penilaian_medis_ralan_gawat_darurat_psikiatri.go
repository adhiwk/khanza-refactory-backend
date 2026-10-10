package rekammedis

import "time"

// PenilaianMedisRalanGawatDaruratPsikiatri tabel `penilaian_medis_ralan_gawat_darurat_psikiatri` (penilaian awal medis IGD psikiatri, RMPenilaianAwalMedisIGDPsikiatri).
type PenilaianMedisRalanGawatDaruratPsikiatri struct {
	NoRawat                              string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                              *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter                             string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Anamnesis                            string     `gorm:"column:anamnesis" json:"anamnesis"`
	Hubungan                             string     `gorm:"column:hubungan" json:"hubungan"`
	KeluhanUtama                         *string    `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	GejalaMenyertai                      *string    `gorm:"column:gejala_menyertai" json:"gejala_menyertai"`
	FaktorPencetus                       *string    `gorm:"column:faktor_pencetus" json:"faktor_pencetus"`
	RiwayatPenyakitDahulu                *string    `gorm:"column:riwayat_penyakit_dahulu" json:"riwayat_penyakit_dahulu"`
	KeteranganRiwayatPenyakitDahulu      *string    `gorm:"column:keterangan_riwayat_penyakit_dahulu" json:"keterangan_riwayat_penyakit_dahulu"`
	RiwayatKehamilan                     *string    `gorm:"column:riwayat_kehamilan" json:"riwayat_kehamilan"`
	RiwayatSosial                        *string    `gorm:"column:riwayat_sosial" json:"riwayat_sosial"`
	KeteranganRiwayatSosial              *string    `gorm:"column:keterangan_riwayat_sosial" json:"keterangan_riwayat_sosial"`
	RiwayatPekerjaan                     *string    `gorm:"column:riwayat_pekerjaan" json:"riwayat_pekerjaan"`
	KeteranganRiwayatPekerjaan           *string    `gorm:"column:keterangan_riwayat_pekerjaan" json:"keterangan_riwayat_pekerjaan"`
	RiwayatObatDiminum                   *string    `gorm:"column:riwayat_obat_diminum" json:"riwayat_obat_diminum"`
	FaktorKepribadianPremorbid           *string    `gorm:"column:faktor_kepribadian_premorbid" json:"faktor_kepribadian_premorbid"`
	FaktorKeturunan                      *string    `gorm:"column:faktor_keturunan" json:"faktor_keturunan"`
	KeteranganFaktorKeturunan            *string    `gorm:"column:keterangan_faktor_keturunan" json:"keterangan_faktor_keturunan"`
	FaktorOrganik                        *string    `gorm:"column:faktor_organik" json:"faktor_organik"`
	KeteranganFaktorOrganik              *string    `gorm:"column:keterangan_faktor_organik" json:"keterangan_faktor_organik"`
	RiwayatAlergi                        *string    `gorm:"column:riwayat_alergi" json:"riwayat_alergi"`
	FisikKesadaran                       *string    `gorm:"column:fisik_kesadaran" json:"fisik_kesadaran"`
	FisikTd                              *string    `gorm:"column:fisik_td" json:"fisik_td"`
	FisikRr                              *string    `gorm:"column:fisik_rr" json:"fisik_rr"`
	FisikSuhu                            *string    `gorm:"column:fisik_suhu" json:"fisik_suhu"`
	FisikNyeri                           *string    `gorm:"column:fisik_nyeri" json:"fisik_nyeri"`
	FisikNadi                            *string    `gorm:"column:fisik_nadi" json:"fisik_nadi"`
	FisikBb                              *string    `gorm:"column:fisik_bb" json:"fisik_bb"`
	FisikTb                              *string    `gorm:"column:fisik_tb" json:"fisik_tb"`
	FisikStatusNutrisi                   *string    `gorm:"column:fisik_status_nutrisi" json:"fisik_status_nutrisi"`
	FisikGcs                             *string    `gorm:"column:fisik_gcs" json:"fisik_gcs"`
	StatusKelainanKepala                 *string    `gorm:"column:status_kelainan_kepala" json:"status_kelainan_kepala"`
	KeteranganStatusKelainanKepala       *string    `gorm:"column:keterangan_status_kelainan_kepala" json:"keterangan_status_kelainan_kepala"`
	StatusKelainanLeher                  *string    `gorm:"column:status_kelainan_leher" json:"status_kelainan_leher"`
	KeteranganStatusKelainanLeher        *string    `gorm:"column:keterangan_status_kelainan_leher" json:"keterangan_status_kelainan_leher"`
	StatusKelainanDada                   *string    `gorm:"column:status_kelainan_dada" json:"status_kelainan_dada"`
	KeteranganStatusKelainanDada         *string    `gorm:"column:keterangan_status_kelainan_dada" json:"keterangan_status_kelainan_dada"`
	StatusKelainanPerut                  *string    `gorm:"column:status_kelainan_perut" json:"status_kelainan_perut"`
	KeteranganStatusKelainanPerut        *string    `gorm:"column:keterangan_status_kelainan_perut" json:"keterangan_status_kelainan_perut"`
	StatusKelainanAnggotaGerak           *string    `gorm:"column:status_kelainan_anggota_gerak" json:"status_kelainan_anggota_gerak"`
	KeteranganStatusKelainanAnggotaGerak *string    `gorm:"column:keterangan_status_kelainan_anggota_gerak" json:"keterangan_status_kelainan_anggota_gerak"`
	StatusLokalisata                     *string    `gorm:"column:status_lokalisata" json:"status_lokalisata"`
	PsikiatrikKesanUmum                  *string    `gorm:"column:psikiatrik_kesan_umum" json:"psikiatrik_kesan_umum"`
	PsikiatrikSikapPrilaku               *string    `gorm:"column:psikiatrik_sikap_prilaku" json:"psikiatrik_sikap_prilaku"`
	PsikiatrikKesadaran                  *string    `gorm:"column:psikiatrik_kesadaran" json:"psikiatrik_kesadaran"`
	PsikiatrikOrientasi                  *string    `gorm:"column:psikiatrik_orientasi" json:"psikiatrik_orientasi"`
	PsikiatrikDayaIngat                  *string    `gorm:"column:psikiatrik_daya_ingat" json:"psikiatrik_daya_ingat"`
	PsikiatrikPersepsi                   *string    `gorm:"column:psikiatrik_persepsi" json:"psikiatrik_persepsi"`
	PsikiatrikPikiran                    *string    `gorm:"column:psikiatrik_pikiran" json:"psikiatrik_pikiran"`
	PsikiatrikInsight                    *string    `gorm:"column:psikiatrik_insight" json:"psikiatrik_insight"`
	Laborat                              *string    `gorm:"column:laborat" json:"laborat"`
	Radiologi                            *string    `gorm:"column:radiologi" json:"radiologi"`
	Ekg                                  *string    `gorm:"column:ekg" json:"ekg"`
	Diagnosis                            *string    `gorm:"column:diagnosis" json:"diagnosis"`
	Permasalahan                         *string    `gorm:"column:permasalahan" json:"permasalahan"`
	InstruksiMedis                       *string    `gorm:"column:instruksi_medis" json:"instruksi_medis"`
	RencanaTarget                        *string    `gorm:"column:rencana_target" json:"rencana_target"`
	PulangDipulangkan                    *string    `gorm:"column:pulang_dipulangkan" json:"pulang_dipulangkan"`
	KeteranganPulangDipulangkan          *string    `gorm:"column:keterangan_pulang_dipulangkan" json:"keterangan_pulang_dipulangkan"`
	PulangDirawatDiruang                 *string    `gorm:"column:pulang_dirawat_diruang" json:"pulang_dirawat_diruang"`
	PulangIndikasiRanap                  *string    `gorm:"column:pulang_indikasi_ranap" json:"pulang_indikasi_ranap"`
	PulangDirujukKe                      *string    `gorm:"column:pulang_dirujuk_ke" json:"pulang_dirujuk_ke"`
	PulangAlasanDirujuk                  *string    `gorm:"column:pulang_alasan_dirujuk" json:"pulang_alasan_dirujuk"`
	PulangPaksa                          *string    `gorm:"column:pulang_paksa" json:"pulang_paksa"`
	KeteranganPulangPaksa                *string    `gorm:"column:keterangan_pulang_paksa" json:"keterangan_pulang_paksa"`
	PulangMeninggalIgd                   *string    `gorm:"column:pulang_meninggal_igd" json:"pulang_meninggal_igd"`
	PulangPenyebabKematian               *string    `gorm:"column:pulang_penyebab_kematian" json:"pulang_penyebab_kematian"`
	FisikPulangKesadaran                 *string    `gorm:"column:fisik_pulang_kesadaran" json:"fisik_pulang_kesadaran"`
	FisikPulangTd                        *string    `gorm:"column:fisik_pulang_td" json:"fisik_pulang_td"`
	FisikPulangNadi                      *string    `gorm:"column:fisik_pulang_nadi" json:"fisik_pulang_nadi"`
	FisikPulangGcs                       *string    `gorm:"column:fisik_pulang_gcs" json:"fisik_pulang_gcs"`
	FisikPulangSuhu                      *string    `gorm:"column:fisik_pulang_suhu" json:"fisik_pulang_suhu"`
	FisikPulangRr                        *string    `gorm:"column:fisik_pulang_rr" json:"fisik_pulang_rr"`
	Edukasi                              *string    `gorm:"column:edukasi" json:"edukasi"`
}

func (PenilaianMedisRalanGawatDaruratPsikiatri) TableName() string {
	return "penilaian_medis_ralan_gawat_darurat_psikiatri"
}
