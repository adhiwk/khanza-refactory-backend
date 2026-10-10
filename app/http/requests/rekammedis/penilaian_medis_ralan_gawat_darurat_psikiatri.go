package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRalanGawatDaruratPsikiatriData isian penilaian awal medis IGD psikiatri.
type PenilaianMedisRalanGawatDaruratPsikiatriData struct {
	Tanggal                              string `form:"tanggal" json:"tanggal"`
	KdDokter                             string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis                            string `form:"anamnesis" json:"anamnesis"`
	Hubungan                             string `form:"hubungan" json:"hubungan"`
	KeluhanUtama                         string `form:"keluhan_utama" json:"keluhan_utama"`
	GejalaMenyertai                      string `form:"gejala_menyertai" json:"gejala_menyertai"`
	FaktorPencetus                       string `form:"faktor_pencetus" json:"faktor_pencetus"`
	RiwayatPenyakitDahulu                string `form:"riwayat_penyakit_dahulu" json:"riwayat_penyakit_dahulu"`
	KeteranganRiwayatPenyakitDahulu      string `form:"keterangan_riwayat_penyakit_dahulu" json:"keterangan_riwayat_penyakit_dahulu"`
	RiwayatKehamilan                     string `form:"riwayat_kehamilan" json:"riwayat_kehamilan"`
	RiwayatSosial                        string `form:"riwayat_sosial" json:"riwayat_sosial"`
	KeteranganRiwayatSosial              string `form:"keterangan_riwayat_sosial" json:"keterangan_riwayat_sosial"`
	RiwayatPekerjaan                     string `form:"riwayat_pekerjaan" json:"riwayat_pekerjaan"`
	KeteranganRiwayatPekerjaan           string `form:"keterangan_riwayat_pekerjaan" json:"keterangan_riwayat_pekerjaan"`
	RiwayatObatDiminum                   string `form:"riwayat_obat_diminum" json:"riwayat_obat_diminum"`
	FaktorKepribadianPremorbid           string `form:"faktor_kepribadian_premorbid" json:"faktor_kepribadian_premorbid"`
	FaktorKeturunan                      string `form:"faktor_keturunan" json:"faktor_keturunan"`
	KeteranganFaktorKeturunan            string `form:"keterangan_faktor_keturunan" json:"keterangan_faktor_keturunan"`
	FaktorOrganik                        string `form:"faktor_organik" json:"faktor_organik"`
	KeteranganFaktorOrganik              string `form:"keterangan_faktor_organik" json:"keterangan_faktor_organik"`
	RiwayatAlergi                        string `form:"riwayat_alergi" json:"riwayat_alergi"`
	FisikKesadaran                       string `form:"fisik_kesadaran" json:"fisik_kesadaran"`
	FisikTd                              string `form:"fisik_td" json:"fisik_td"`
	FisikRr                              string `form:"fisik_rr" json:"fisik_rr"`
	FisikSuhu                            string `form:"fisik_suhu" json:"fisik_suhu"`
	FisikNyeri                           string `form:"fisik_nyeri" json:"fisik_nyeri"`
	FisikNadi                            string `form:"fisik_nadi" json:"fisik_nadi"`
	FisikBb                              string `form:"fisik_bb" json:"fisik_bb"`
	FisikTb                              string `form:"fisik_tb" json:"fisik_tb"`
	FisikStatusNutrisi                   string `form:"fisik_status_nutrisi" json:"fisik_status_nutrisi"`
	FisikGcs                             string `form:"fisik_gcs" json:"fisik_gcs"`
	StatusKelainanKepala                 string `form:"status_kelainan_kepala" json:"status_kelainan_kepala"`
	KeteranganStatusKelainanKepala       string `form:"keterangan_status_kelainan_kepala" json:"keterangan_status_kelainan_kepala"`
	StatusKelainanLeher                  string `form:"status_kelainan_leher" json:"status_kelainan_leher"`
	KeteranganStatusKelainanLeher        string `form:"keterangan_status_kelainan_leher" json:"keterangan_status_kelainan_leher"`
	StatusKelainanDada                   string `form:"status_kelainan_dada" json:"status_kelainan_dada"`
	KeteranganStatusKelainanDada         string `form:"keterangan_status_kelainan_dada" json:"keterangan_status_kelainan_dada"`
	StatusKelainanPerut                  string `form:"status_kelainan_perut" json:"status_kelainan_perut"`
	KeteranganStatusKelainanPerut        string `form:"keterangan_status_kelainan_perut" json:"keterangan_status_kelainan_perut"`
	StatusKelainanAnggotaGerak           string `form:"status_kelainan_anggota_gerak" json:"status_kelainan_anggota_gerak"`
	KeteranganStatusKelainanAnggotaGerak string `form:"keterangan_status_kelainan_anggota_gerak" json:"keterangan_status_kelainan_anggota_gerak"`
	StatusLokalisata                     string `form:"status_lokalisata" json:"status_lokalisata"`
	PsikiatrikKesanUmum                  string `form:"psikiatrik_kesan_umum" json:"psikiatrik_kesan_umum"`
	PsikiatrikSikapPrilaku               string `form:"psikiatrik_sikap_prilaku" json:"psikiatrik_sikap_prilaku"`
	PsikiatrikKesadaran                  string `form:"psikiatrik_kesadaran" json:"psikiatrik_kesadaran"`
	PsikiatrikOrientasi                  string `form:"psikiatrik_orientasi" json:"psikiatrik_orientasi"`
	PsikiatrikDayaIngat                  string `form:"psikiatrik_daya_ingat" json:"psikiatrik_daya_ingat"`
	PsikiatrikPersepsi                   string `form:"psikiatrik_persepsi" json:"psikiatrik_persepsi"`
	PsikiatrikPikiran                    string `form:"psikiatrik_pikiran" json:"psikiatrik_pikiran"`
	PsikiatrikInsight                    string `form:"psikiatrik_insight" json:"psikiatrik_insight"`
	Laborat                              string `form:"laborat" json:"laborat"`
	Radiologi                            string `form:"radiologi" json:"radiologi"`
	Ekg                                  string `form:"ekg" json:"ekg"`
	Diagnosis                            string `form:"diagnosis" json:"diagnosis"`
	Permasalahan                         string `form:"permasalahan" json:"permasalahan"`
	InstruksiMedis                       string `form:"instruksi_medis" json:"instruksi_medis"`
	RencanaTarget                        string `form:"rencana_target" json:"rencana_target"`
	PulangDipulangkan                    string `form:"pulang_dipulangkan" json:"pulang_dipulangkan"`
	KeteranganPulangDipulangkan          string `form:"keterangan_pulang_dipulangkan" json:"keterangan_pulang_dipulangkan"`
	PulangDirawatDiruang                 string `form:"pulang_dirawat_diruang" json:"pulang_dirawat_diruang"`
	PulangIndikasiRanap                  string `form:"pulang_indikasi_ranap" json:"pulang_indikasi_ranap"`
	PulangDirujukKe                      string `form:"pulang_dirujuk_ke" json:"pulang_dirujuk_ke"`
	PulangAlasanDirujuk                  string `form:"pulang_alasan_dirujuk" json:"pulang_alasan_dirujuk"`
	PulangPaksa                          string `form:"pulang_paksa" json:"pulang_paksa"`
	KeteranganPulangPaksa                string `form:"keterangan_pulang_paksa" json:"keterangan_pulang_paksa"`
	PulangMeninggalIgd                   string `form:"pulang_meninggal_igd" json:"pulang_meninggal_igd"`
	PulangPenyebabKematian               string `form:"pulang_penyebab_kematian" json:"pulang_penyebab_kematian"`
	FisikPulangKesadaran                 string `form:"fisik_pulang_kesadaran" json:"fisik_pulang_kesadaran"`
	FisikPulangTd                        string `form:"fisik_pulang_td" json:"fisik_pulang_td"`
	FisikPulangNadi                      string `form:"fisik_pulang_nadi" json:"fisik_pulang_nadi"`
	FisikPulangGcs                       string `form:"fisik_pulang_gcs" json:"fisik_pulang_gcs"`
	FisikPulangSuhu                      string `form:"fisik_pulang_suhu" json:"fisik_pulang_suhu"`
	FisikPulangRr                        string `form:"fisik_pulang_rr" json:"fisik_pulang_rr"`
	Edukasi                              string `form:"edukasi" json:"edukasi"`
}

func penilaianMedisRalanGawatDaruratPsikiatriRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                  "required|date",
		"kd_dokter":                                "required|string|max_len:20",
		"anamnesis":                                "required|in:Autoanamnesis,Alloanamnesis",
		"hubungan":                                 "string|max_len:30",
		"keluhan_utama":                            "string|max_len:2000",
		"gejala_menyertai":                         "string|max_len:1000",
		"faktor_pencetus":                          "string|max_len:1000",
		"riwayat_penyakit_dahulu":                  "in:Tidak Ada,Ada",
		"keterangan_riwayat_penyakit_dahulu":       "string|max_len:1000",
		"riwayat_kehamilan":                        "string|max_len:1000",
		"riwayat_sosial":                           "in:Bergaul,Tidak Bergaul,Lain-lain",
		"keterangan_riwayat_sosial":                "string|max_len:50",
		"riwayat_pekerjaan":                        "in:Bekerja,Tidak Bekerja,Ganti-gantian Pekerjaan",
		"keterangan_riwayat_pekerjaan":             "string|max_len:50",
		"riwayat_obat_diminum":                     "string|max_len:1000",
		"faktor_kepribadian_premorbid":             "string|max_len:50",
		"faktor_keturunan":                         "in:Tidak Ada,Ada",
		"keterangan_faktor_keturunan":              "string|max_len:50",
		"faktor_organik":                           "in:Tidak Ada,Ada",
		"keterangan_faktor_organik":                "string|max_len:50",
		"riwayat_alergi":                           "string|max_len:50",
		"fisik_kesadaran":                          "in:Compos Mentis,Apatis,Somnolen,Sopor,Koma",
		"fisik_td":                                 "string|max_len:8",
		"fisik_rr":                                 "string|max_len:5",
		"fisik_suhu":                               "string|max_len:5",
		"fisik_nyeri":                              "in:Tidak Nyeri,Nyeri Ringan,Nyeri Sedang,Nyeri Berat,Nyeri Sangat Berat,Nyeri Tak Tertahankan",
		"fisik_nadi":                               "string|max_len:5",
		"fisik_bb":                                 "string|max_len:5",
		"fisik_tb":                                 "string|max_len:5",
		"fisik_status_nutrisi":                     "string|max_len:100",
		"fisik_gcs":                                "string|max_len:8",
		"status_kelainan_kepala":                   "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_status_kelainan_kepala":        "string|max_len:50",
		"status_kelainan_leher":                    "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_status_kelainan_leher":         "string|max_len:50",
		"status_kelainan_dada":                     "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_status_kelainan_dada":          "string|max_len:50",
		"status_kelainan_perut":                    "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_status_kelainan_perut":         "string|max_len:50",
		"status_kelainan_anggota_gerak":            "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_status_kelainan_anggota_gerak": "string|max_len:50",
		"status_lokalisata":                        "string|max_len:1000",
		"psikiatrik_kesan_umum":                    "string|max_len:50",
		"psikiatrik_sikap_prilaku":                 "string|max_len:50",
		"psikiatrik_kesadaran":                     "string|max_len:50",
		"psikiatrik_orientasi":                     "string|max_len:50",
		"psikiatrik_daya_ingat":                    "string|max_len:50",
		"psikiatrik_persepsi":                      "string|max_len:50",
		"psikiatrik_pikiran":                       "string|max_len:50",
		"psikiatrik_insight":                       "string|max_len:50",
		"laborat":                                  "string|max_len:300",
		"radiologi":                                "string|max_len:200",
		"ekg":                                      "string|max_len:200",
		"diagnosis":                                "string|max_len:1000",
		"permasalahan":                             "string|max_len:500",
		"instruksi_medis":                          "string|max_len:600",
		"rencana_target":                           "string|max_len:1000",
		"pulang_dipulangkan":                       "in:Tidak Perlu Kontrol,Kontrol/Berobat Jalan,Rawat Inap,-",
		"keterangan_pulang_dipulangkan":            "string|max_len:100",
		"pulang_dirawat_diruang":                   "string|max_len:30",
		"pulang_indikasi_ranap":                    "string|max_len:100",
		"pulang_dirujuk_ke":                        "string|max_len:70",
		"pulang_alasan_dirujuk":                    "in:-,Tempat Penuh,Perlu Fasilitas Lebih,Permintaan Pasien/Keluarga",
		"pulang_paksa":                             "in:-,Masalah Biaya,Kondisi Pasien,Masalah Lokasi Rumah,Lain-lain",
		"keterangan_pulang_paksa":                  "string|max_len:100",
		"pulang_meninggal_igd":                     "in:-,<= 2 Jam,> 2 Jam",
		"pulang_penyebab_kematian":                 "string|max_len:100",
		"fisik_pulang_kesadaran":                   "in:Compos Mentis,Apatis,Somnolen,Sopor,Koma",
		"fisik_pulang_td":                          "string|max_len:8",
		"fisik_pulang_nadi":                        "string|max_len:5",
		"fisik_pulang_gcs":                         "string|max_len:8",
		"fisik_pulang_suhu":                        "string|max_len:5",
		"fisik_pulang_rr":                          "string|max_len:5",
		"edukasi":                                  "string|max_len:1000",
	}
	return rules
}

// PenilaianMedisRalanGawatDaruratPsikiatriStore simpan penilaian awal medis IGD psikiatri; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRalanGawatDaruratPsikiatriStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRalanGawatDaruratPsikiatriData
}

func (r *PenilaianMedisRalanGawatDaruratPsikiatriStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanGawatDaruratPsikiatriStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRalanGawatDaruratPsikiatriRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRalanGawatDaruratPsikiatriStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianMedisRalanGawatDaruratPsikiatriStore) Payload() PenilaianMedisRalanGawatDaruratPsikiatriData {
	return r.PenilaianMedisRalanGawatDaruratPsikiatriData
}

func (r *PenilaianMedisRalanGawatDaruratPsikiatriStore) DetailValues() map[string][]string {
	return nil
}

// PenilaianMedisRalanGawatDaruratPsikiatriUpdate ubah penilaian awal medis IGD psikiatri (PUT); kunci lewat query string.
type PenilaianMedisRalanGawatDaruratPsikiatriUpdate struct {
	PenilaianMedisRalanGawatDaruratPsikiatriData
}

func (r *PenilaianMedisRalanGawatDaruratPsikiatriUpdate) Authorize(ctx http.Context) error {
	return nil
}

func (r *PenilaianMedisRalanGawatDaruratPsikiatriUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRalanGawatDaruratPsikiatriRules()
}

func (r *PenilaianMedisRalanGawatDaruratPsikiatriUpdate) Payload() PenilaianMedisRalanGawatDaruratPsikiatriData {
	return r.PenilaianMedisRalanGawatDaruratPsikiatriData
}

func (r *PenilaianMedisRalanGawatDaruratPsikiatriUpdate) DetailValues() map[string][]string {
	return nil
}
