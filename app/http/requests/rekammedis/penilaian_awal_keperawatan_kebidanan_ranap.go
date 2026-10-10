package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianAwalKeperawatanKebidananRanapData isian laporan pemantauan anastesi.
type PenilaianAwalKeperawatanKebidananRanapData struct {
	Tanggal                                 string  `form:"tanggal" json:"tanggal"`
	Informasi                               string  `form:"informasi" json:"informasi"`
	TibaDiruangRawat                        string  `form:"tiba_diruang_rawat" json:"tiba_diruang_rawat"`
	CaraMasuk                               string  `form:"cara_masuk" json:"cara_masuk"`
	Keluhan                                 string  `form:"keluhan" json:"keluhan"`
	Rpk                                     string  `form:"rpk" json:"rpk"`
	Psk                                     string  `form:"psk" json:"psk"`
	Rp                                      string  `form:"rp" json:"rp"`
	Alergi                                  string  `form:"alergi" json:"alergi"`
	KomplikasiSebelumnya                    string  `form:"komplikasi_sebelumnya" json:"komplikasi_sebelumnya"`
	KeteranganKomplikasiSebelumnya          string  `form:"keterangan_komplikasi_sebelumnya" json:"keterangan_komplikasi_sebelumnya"`
	RiwayatMensUmur                         string  `form:"riwayat_mens_umur" json:"riwayat_mens_umur"`
	RiwayatMensLamanya                      string  `form:"riwayat_mens_lamanya" json:"riwayat_mens_lamanya"`
	RiwayatMensBanyaknya                    string  `form:"riwayat_mens_banyaknya" json:"riwayat_mens_banyaknya"`
	RiwayatMensSiklus                       string  `form:"riwayat_mens_siklus" json:"riwayat_mens_siklus"`
	RiwayatMensKetSiklus                    string  `form:"riwayat_mens_ket_siklus" json:"riwayat_mens_ket_siklus"`
	RiwayatMensDirasakan                    string  `form:"riwayat_mens_dirasakan" json:"riwayat_mens_dirasakan"`
	RiwayatPerkawinanStatus                 string  `form:"riwayat_perkawinan_status" json:"riwayat_perkawinan_status"`
	RiwayatPerkawinanKetStatus              string  `form:"riwayat_perkawinan_ket_status" json:"riwayat_perkawinan_ket_status"`
	RiwayatPerkawinanUsia1                  string  `form:"riwayat_perkawinan_usia1" json:"riwayat_perkawinan_usia1"`
	RiwayatPerkawinanKetUsia1               string  `form:"riwayat_perkawinan_ket_usia1" json:"riwayat_perkawinan_ket_usia1"`
	RiwayatPerkawinanUsia2                  string  `form:"riwayat_perkawinan_usia2" json:"riwayat_perkawinan_usia2"`
	RiwayatPerkawinanKetUsia2               string  `form:"riwayat_perkawinan_ket_usia2" json:"riwayat_perkawinan_ket_usia2"`
	RiwayatPerkawinanUsia3                  string  `form:"riwayat_perkawinan_usia3" json:"riwayat_perkawinan_usia3"`
	RiwayatPerkawinanKetUsia3               string  `form:"riwayat_perkawinan_ket_usia3" json:"riwayat_perkawinan_ket_usia3"`
	RiwayatPersalinanG                      string  `form:"riwayat_persalinan_g" json:"riwayat_persalinan_g"`
	RiwayatPersalinanP                      string  `form:"riwayat_persalinan_p" json:"riwayat_persalinan_p"`
	RiwayatPersalinanA                      string  `form:"riwayat_persalinan_a" json:"riwayat_persalinan_a"`
	RiwayatPersalinanHidup                  string  `form:"riwayat_persalinan_hidup" json:"riwayat_persalinan_hidup"`
	RiwayatHamilHpht                        string  `form:"riwayat_hamil_hpht" json:"riwayat_hamil_hpht"`
	RiwayatHamilUsiahamil                   string  `form:"riwayat_hamil_usiahamil" json:"riwayat_hamil_usiahamil"`
	RiwayatHamilTp                          string  `form:"riwayat_hamil_tp" json:"riwayat_hamil_tp"`
	RiwayatHamilImunisasi                   string  `form:"riwayat_hamil_imunisasi" json:"riwayat_hamil_imunisasi"`
	RiwayatHamilAnc                         string  `form:"riwayat_hamil_anc" json:"riwayat_hamil_anc"`
	RiwayatHamilAncke                       string  `form:"riwayat_hamil_ancke" json:"riwayat_hamil_ancke"`
	RiwayatHamilKetAncke                    string  `form:"riwayat_hamil_ket_ancke" json:"riwayat_hamil_ket_ancke"`
	RiwayatHamilKeluhanHamilMuda            string  `form:"riwayat_hamil_keluhan_hamil_muda" json:"riwayat_hamil_keluhan_hamil_muda"`
	RiwayatHamilKeluhanHamilTua             string  `form:"riwayat_hamil_keluhan_hamil_tua" json:"riwayat_hamil_keluhan_hamil_tua"`
	RiwayatKb                               string  `form:"riwayat_kb" json:"riwayat_kb"`
	RiwayatKbLamanya                        string  `form:"riwayat_kb_lamanya" json:"riwayat_kb_lamanya"`
	RiwayatKbKomplikasi                     string  `form:"riwayat_kb_komplikasi" json:"riwayat_kb_komplikasi"`
	RiwayatKbKetKomplikasi                  string  `form:"riwayat_kb_ket_komplikasi" json:"riwayat_kb_ket_komplikasi"`
	RiwayatKbKapaberhenti                   string  `form:"riwayat_kb_kapaberhenti" json:"riwayat_kb_kapaberhenti"`
	RiwayatKbAlasanberhenti                 string  `form:"riwayat_kb_alasanberhenti" json:"riwayat_kb_alasanberhenti"`
	RiwayatGenekologi                       string  `form:"riwayat_genekologi" json:"riwayat_genekologi"`
	RiwayatKebiasaanObat                    string  `form:"riwayat_kebiasaan_obat" json:"riwayat_kebiasaan_obat"`
	RiwayatKebiasaanKetObat                 string  `form:"riwayat_kebiasaan_ket_obat" json:"riwayat_kebiasaan_ket_obat"`
	RiwayatKebiasaanMerokok                 string  `form:"riwayat_kebiasaan_merokok" json:"riwayat_kebiasaan_merokok"`
	RiwayatKebiasaanKetMerokok              string  `form:"riwayat_kebiasaan_ket_merokok" json:"riwayat_kebiasaan_ket_merokok"`
	RiwayatKebiasaanAlkohol                 string  `form:"riwayat_kebiasaan_alkohol" json:"riwayat_kebiasaan_alkohol"`
	RiwayatKebiasaanKetAlkohol              string  `form:"riwayat_kebiasaan_ket_alkohol" json:"riwayat_kebiasaan_ket_alkohol"`
	RiwayatKebiasaanNarkoba                 string  `form:"riwayat_kebiasaan_narkoba" json:"riwayat_kebiasaan_narkoba"`
	PemeriksaanKebidananMental              string  `form:"pemeriksaan_kebidanan_mental" json:"pemeriksaan_kebidanan_mental"`
	PemeriksaanKebidananKeadaanUmum         string  `form:"pemeriksaan_kebidanan_keadaan_umum" json:"pemeriksaan_kebidanan_keadaan_umum"`
	PemeriksaanKebidananGcs                 string  `form:"pemeriksaan_kebidanan_gcs" json:"pemeriksaan_kebidanan_gcs"`
	PemeriksaanKebidananTd                  string  `form:"pemeriksaan_kebidanan_td" json:"pemeriksaan_kebidanan_td"`
	PemeriksaanKebidananNadi                string  `form:"pemeriksaan_kebidanan_nadi" json:"pemeriksaan_kebidanan_nadi"`
	PemeriksaanKebidananRr                  string  `form:"pemeriksaan_kebidanan_rr" json:"pemeriksaan_kebidanan_rr"`
	PemeriksaanKebidananSuhu                string  `form:"pemeriksaan_kebidanan_suhu" json:"pemeriksaan_kebidanan_suhu"`
	PemeriksaanKebidananSpo2                string  `form:"pemeriksaan_kebidanan_spo2" json:"pemeriksaan_kebidanan_spo2"`
	PemeriksaanKebidananBb                  string  `form:"pemeriksaan_kebidanan_bb" json:"pemeriksaan_kebidanan_bb"`
	PemeriksaanKebidananTb                  string  `form:"pemeriksaan_kebidanan_tb" json:"pemeriksaan_kebidanan_tb"`
	PemeriksaanKebidananLila                string  `form:"pemeriksaan_kebidanan_lila" json:"pemeriksaan_kebidanan_lila"`
	PemeriksaanKebidananTfu                 string  `form:"pemeriksaan_kebidanan_tfu" json:"pemeriksaan_kebidanan_tfu"`
	PemeriksaanKebidananTbj                 string  `form:"pemeriksaan_kebidanan_tbj" json:"pemeriksaan_kebidanan_tbj"`
	PemeriksaanKebidananLetak               string  `form:"pemeriksaan_kebidanan_letak" json:"pemeriksaan_kebidanan_letak"`
	PemeriksaanKebidananPresentasi          string  `form:"pemeriksaan_kebidanan_presentasi" json:"pemeriksaan_kebidanan_presentasi"`
	PemeriksaanKebidananPenurunan           string  `form:"pemeriksaan_kebidanan_penurunan" json:"pemeriksaan_kebidanan_penurunan"`
	PemeriksaanKebidananHis                 string  `form:"pemeriksaan_kebidanan_his" json:"pemeriksaan_kebidanan_his"`
	PemeriksaanKebidananKekuatan            string  `form:"pemeriksaan_kebidanan_kekuatan" json:"pemeriksaan_kebidanan_kekuatan"`
	PemeriksaanKebidananLamanya             string  `form:"pemeriksaan_kebidanan_lamanya" json:"pemeriksaan_kebidanan_lamanya"`
	PemeriksaanKebidananDjj                 string  `form:"pemeriksaan_kebidanan_djj" json:"pemeriksaan_kebidanan_djj"`
	PemeriksaanKebidananKetDjj              string  `form:"pemeriksaan_kebidanan_ket_djj" json:"pemeriksaan_kebidanan_ket_djj"`
	PemeriksaanKebidananPortio              string  `form:"pemeriksaan_kebidanan_portio" json:"pemeriksaan_kebidanan_portio"`
	PemeriksaanKebidananPembukaan           string  `form:"pemeriksaan_kebidanan_pembukaan" json:"pemeriksaan_kebidanan_pembukaan"`
	PemeriksaanKebidananKetuban             string  `form:"pemeriksaan_kebidanan_ketuban" json:"pemeriksaan_kebidanan_ketuban"`
	PemeriksaanKebidananHodge               string  `form:"pemeriksaan_kebidanan_hodge" json:"pemeriksaan_kebidanan_hodge"`
	PemeriksaanKebidananPanggul             string  `form:"pemeriksaan_kebidanan_panggul" json:"pemeriksaan_kebidanan_panggul"`
	PemeriksaanKebidananInspekulo           string  `form:"pemeriksaan_kebidanan_inspekulo" json:"pemeriksaan_kebidanan_inspekulo"`
	PemeriksaanKebidananKetInspekulo        string  `form:"pemeriksaan_kebidanan_ket_inspekulo" json:"pemeriksaan_kebidanan_ket_inspekulo"`
	PemeriksaanKebidananLakmus              string  `form:"pemeriksaan_kebidanan_lakmus" json:"pemeriksaan_kebidanan_lakmus"`
	PemeriksaanKebidananKetLakmus           string  `form:"pemeriksaan_kebidanan_ket_lakmus" json:"pemeriksaan_kebidanan_ket_lakmus"`
	PemeriksaanKebidananCtg                 string  `form:"pemeriksaan_kebidanan_ctg" json:"pemeriksaan_kebidanan_ctg"`
	PemeriksaanKebidananKetCtg              string  `form:"pemeriksaan_kebidanan_ket_ctg" json:"pemeriksaan_kebidanan_ket_ctg"`
	PemeriksaanUmumKepala                   string  `form:"pemeriksaan_umum_kepala" json:"pemeriksaan_umum_kepala"`
	PemeriksaanUmumMuka                     string  `form:"pemeriksaan_umum_muka" json:"pemeriksaan_umum_muka"`
	PemeriksaanUmumMata                     string  `form:"pemeriksaan_umum_mata" json:"pemeriksaan_umum_mata"`
	PemeriksaanUmumHidung                   string  `form:"pemeriksaan_umum_hidung" json:"pemeriksaan_umum_hidung"`
	PemeriksaanUmumTelinga                  string  `form:"pemeriksaan_umum_telinga" json:"pemeriksaan_umum_telinga"`
	PemeriksaanUmumMulut                    string  `form:"pemeriksaan_umum_mulut" json:"pemeriksaan_umum_mulut"`
	PemeriksaanUmumLeher                    string  `form:"pemeriksaan_umum_leher" json:"pemeriksaan_umum_leher"`
	PemeriksaanUmumDada                     string  `form:"pemeriksaan_umum_dada" json:"pemeriksaan_umum_dada"`
	PemeriksaanUmumPerut                    string  `form:"pemeriksaan_umum_perut" json:"pemeriksaan_umum_perut"`
	PemeriksaanUmumGenitalia                string  `form:"pemeriksaan_umum_genitalia" json:"pemeriksaan_umum_genitalia"`
	PemeriksaanUmumEkstrimitas              string  `form:"pemeriksaan_umum_ekstrimitas" json:"pemeriksaan_umum_ekstrimitas"`
	PengkajianFungsiKemampuanAktifitas      string  `form:"pengkajian_fungsi_kemampuan_aktifitas" json:"pengkajian_fungsi_kemampuan_aktifitas"`
	PengkajianFungsiBerjalan                string  `form:"pengkajian_fungsi_berjalan" json:"pengkajian_fungsi_berjalan"`
	PengkajianFungsiKetBerjalan             string  `form:"pengkajian_fungsi_ket_berjalan" json:"pengkajian_fungsi_ket_berjalan"`
	PengkajianFungsiAktivitas               string  `form:"pengkajian_fungsi_aktivitas" json:"pengkajian_fungsi_aktivitas"`
	PengkajianFungsiAmbulasi                string  `form:"pengkajian_fungsi_ambulasi" json:"pengkajian_fungsi_ambulasi"`
	PengkajianFungsiEkstrimitasAtas         string  `form:"pengkajian_fungsi_ekstrimitas_atas" json:"pengkajian_fungsi_ekstrimitas_atas"`
	PengkajianFungsiKetEkstrimitasAtas      string  `form:"pengkajian_fungsi_ket_ekstrimitas_atas" json:"pengkajian_fungsi_ket_ekstrimitas_atas"`
	PengkajianFungsiEkstrimitasBawah        string  `form:"pengkajian_fungsi_ekstrimitas_bawah" json:"pengkajian_fungsi_ekstrimitas_bawah"`
	PengkajianFungsiKetEkstrimitasBawah     string  `form:"pengkajian_fungsi_ket_ekstrimitas_bawah" json:"pengkajian_fungsi_ket_ekstrimitas_bawah"`
	PengkajianFungsiKemampuanMenggenggam    string  `form:"pengkajian_fungsi_kemampuan_menggenggam" json:"pengkajian_fungsi_kemampuan_menggenggam"`
	PengkajianFungsiKetKemampuanMenggenggam string  `form:"pengkajian_fungsi_ket_kemampuan_menggenggam" json:"pengkajian_fungsi_ket_kemampuan_menggenggam"`
	PengkajianFungsiKoordinasi              string  `form:"pengkajian_fungsi_koordinasi" json:"pengkajian_fungsi_koordinasi"`
	PengkajianFungsiKetKoordinasi           string  `form:"pengkajian_fungsi_ket_koordinasi" json:"pengkajian_fungsi_ket_koordinasi"`
	PengkajianFungsiGangguanFungsi          string  `form:"pengkajian_fungsi_gangguan_fungsi" json:"pengkajian_fungsi_gangguan_fungsi"`
	RiwayatPsikoKondisipsiko                string  `form:"riwayat_psiko_kondisipsiko" json:"riwayat_psiko_kondisipsiko"`
	RiwayatPsikoAdakahPrilaku               string  `form:"riwayat_psiko_adakah_prilaku" json:"riwayat_psiko_adakah_prilaku"`
	RiwayatPsikoKetAdakahPrilaku            string  `form:"riwayat_psiko_ket_adakah_prilaku" json:"riwayat_psiko_ket_adakah_prilaku"`
	RiwayatPsikoGangguanJiwa                string  `form:"riwayat_psiko_gangguan_jiwa" json:"riwayat_psiko_gangguan_jiwa"`
	RiwayatPsikoHubunganPasien              string  `form:"riwayat_psiko_hubungan_pasien" json:"riwayat_psiko_hubungan_pasien"`
	RiwayatPsikoTinggalDengan               string  `form:"riwayat_psiko_tinggal_dengan" json:"riwayat_psiko_tinggal_dengan"`
	RiwayatPsikoKetTinggalDengan            string  `form:"riwayat_psiko_ket_tinggal_dengan" json:"riwayat_psiko_ket_tinggal_dengan"`
	RiwayatPsikoBudaya                      string  `form:"riwayat_psiko_budaya" json:"riwayat_psiko_budaya"`
	RiwayatPsikoKetBudaya                   string  `form:"riwayat_psiko_ket_budaya" json:"riwayat_psiko_ket_budaya"`
	RiwayatPsikoPendPj                      string  `form:"riwayat_psiko_pend_pj" json:"riwayat_psiko_pend_pj"`
	RiwayatPsikoEdukasiPada                 string  `form:"riwayat_psiko_edukasi_pada" json:"riwayat_psiko_edukasi_pada"`
	RiwayatPsikoKetEdukasiPada              string  `form:"riwayat_psiko_ket_edukasi_pada" json:"riwayat_psiko_ket_edukasi_pada"`
	PenilaianNyeri                          string  `form:"penilaian_nyeri" json:"penilaian_nyeri"`
	PenilaianNyeriPenyebab                  string  `form:"penilaian_nyeri_penyebab" json:"penilaian_nyeri_penyebab"`
	PenilaianNyeriKetPenyebab               string  `form:"penilaian_nyeri_ket_penyebab" json:"penilaian_nyeri_ket_penyebab"`
	PenilaianNyeriKualitas                  string  `form:"penilaian_nyeri_kualitas" json:"penilaian_nyeri_kualitas"`
	PenilaianNyeriKetKualitas               string  `form:"penilaian_nyeri_ket_kualitas" json:"penilaian_nyeri_ket_kualitas"`
	PenilaianNyeriLokasi                    string  `form:"penilaian_nyeri_lokasi" json:"penilaian_nyeri_lokasi"`
	PenilaianNyeriMenyebar                  string  `form:"penilaian_nyeri_menyebar" json:"penilaian_nyeri_menyebar"`
	PenilaianNyeriSkala                     string  `form:"penilaian_nyeri_skala" json:"penilaian_nyeri_skala"`
	PenilaianNyeriWaktu                     string  `form:"penilaian_nyeri_waktu" json:"penilaian_nyeri_waktu"`
	PenilaianNyeriHilang                    string  `form:"penilaian_nyeri_hilang" json:"penilaian_nyeri_hilang"`
	PenilaianNyeriKetHilang                 string  `form:"penilaian_nyeri_ket_hilang" json:"penilaian_nyeri_ket_hilang"`
	PenilaianNyeriDiberitahukanDokter       string  `form:"penilaian_nyeri_diberitahukan_dokter" json:"penilaian_nyeri_diberitahukan_dokter"`
	PenilaianNyeriJamDiberitahukanDokter    string  `form:"penilaian_nyeri_jam_diberitahukan_dokter" json:"penilaian_nyeri_jam_diberitahukan_dokter"`
	PenilaianJatuhSkala1                    string  `form:"penilaian_jatuh_skala1" json:"penilaian_jatuh_skala1"`
	PenilaianJatuhNilai1                    int     `form:"penilaian_jatuh_nilai1" json:"penilaian_jatuh_nilai1"`
	PenilaianJatuhSkala2                    string  `form:"penilaian_jatuh_skala2" json:"penilaian_jatuh_skala2"`
	PenilaianJatuhNilai2                    int     `form:"penilaian_jatuh_nilai2" json:"penilaian_jatuh_nilai2"`
	PenilaianJatuhSkala3                    string  `form:"penilaian_jatuh_skala3" json:"penilaian_jatuh_skala3"`
	PenilaianJatuhNilai3                    int     `form:"penilaian_jatuh_nilai3" json:"penilaian_jatuh_nilai3"`
	PenilaianJatuhSkala4                    string  `form:"penilaian_jatuh_skala4" json:"penilaian_jatuh_skala4"`
	PenilaianJatuhNilai4                    int     `form:"penilaian_jatuh_nilai4" json:"penilaian_jatuh_nilai4"`
	PenilaianJatuhSkala5                    string  `form:"penilaian_jatuh_skala5" json:"penilaian_jatuh_skala5"`
	PenilaianJatuhNilai5                    int     `form:"penilaian_jatuh_nilai5" json:"penilaian_jatuh_nilai5"`
	PenilaianJatuhSkala6                    string  `form:"penilaian_jatuh_skala6" json:"penilaian_jatuh_skala6"`
	PenilaianJatuhNilai6                    int     `form:"penilaian_jatuh_nilai6" json:"penilaian_jatuh_nilai6"`
	PenilaianJatuhTotalnilai                float64 `form:"penilaian_jatuh_totalnilai" json:"penilaian_jatuh_totalnilai"`
	SkriningGizi1                           string  `form:"skrining_gizi1" json:"skrining_gizi1"`
	NilaiGizi1                              int     `form:"nilai_gizi1" json:"nilai_gizi1"`
	SkriningGizi2                           string  `form:"skrining_gizi2" json:"skrining_gizi2"`
	NilaiGizi2                              int     `form:"nilai_gizi2" json:"nilai_gizi2"`
	NilaiTotalGizi                          float64 `form:"nilai_total_gizi" json:"nilai_total_gizi"`
	SkriningGiziDiagnosaKhusus              string  `form:"skrining_gizi_diagnosa_khusus" json:"skrining_gizi_diagnosa_khusus"`
	SkriningGiziKetDiagnosaKhusus           string  `form:"skrining_gizi_ket_diagnosa_khusus" json:"skrining_gizi_ket_diagnosa_khusus"`
	SkriningGiziDiketahuiDietisen           string  `form:"skrining_gizi_diketahui_dietisen" json:"skrining_gizi_diketahui_dietisen"`
	SkriningGiziJamDiketahuiDietisen        string  `form:"skrining_gizi_jam_diketahui_dietisen" json:"skrining_gizi_jam_diketahui_dietisen"`
	Masalah                                 string  `form:"masalah" json:"masalah"`
	Rencana                                 string  `form:"rencana" json:"rencana"`
	Nip1                                    string  `form:"nip1" json:"nip1"`
	Nip2                                    string  `form:"nip2" json:"nip2"`
	KdDokter                                string  `form:"kd_dokter" json:"kd_dokter"`
}

func penilaianAwalKeperawatanKebidananRanapRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                     "required|date",
		"informasi":                                   "required|in:Autoanamnesis,Alloanamnesis",
		"tiba_diruang_rawat":                          "required|in:Jalan Tanpa Bantuan,Kursi Roda,Brankar",
		"cara_masuk":                                  "required|in:Poli,IGD,VK,OK,Lain-lain",
		"keluhan":                                     "string|max_len:500",
		"rpk":                                         "string|max_len:100",
		"psk":                                         "string|max_len:100",
		"rp":                                          "string|max_len:100",
		"alergi":                                      "string|max_len:25",
		"komplikasi_sebelumnya":                       "required|in:Tidak,HAP,HPP,PEB/PER/Eklamsi,Lain-lain",
		"keterangan_komplikasi_sebelumnya":            "string|max_len:30",
		"riwayat_mens_umur":                           "string|max_len:10",
		"riwayat_mens_lamanya":                        "string|max_len:10",
		"riwayat_mens_banyaknya":                      "string|max_len:10",
		"riwayat_mens_siklus":                         "string|max_len:10",
		"riwayat_mens_ket_siklus":                     "required|in:Teratur,Tidak Teratur",
		"riwayat_mens_dirasakan":                      "required|in:Tidak Ada Masalah,Dismenorhea,Spotting,Menorhagia,PMS",
		"riwayat_perkawinan_status":                   "required|in:Menikah,Tidak / Belum Menikah",
		"riwayat_perkawinan_ket_status":               "string|max_len:5",
		"riwayat_perkawinan_usia1":                    "string|max_len:5",
		"riwayat_perkawinan_ket_usia1":                "required|in:-,Masih Menikah,Cerai,Meninggal",
		"riwayat_perkawinan_usia2":                    "string|max_len:5",
		"riwayat_perkawinan_ket_usia2":                "required|in:-,Masih Menikah,Cerai,Meninggal",
		"riwayat_perkawinan_usia3":                    "string|max_len:5",
		"riwayat_perkawinan_ket_usia3":                "required|in:-,Masih Menikah,Cerai,Meninggal",
		"riwayat_persalinan_g":                        "string|max_len:10",
		"riwayat_persalinan_p":                        "string|max_len:10",
		"riwayat_persalinan_a":                        "string|max_len:10",
		"riwayat_persalinan_hidup":                    "string|max_len:10",
		"riwayat_hamil_hpht":                          "required|date",
		"riwayat_hamil_usiahamil":                     "string|max_len:10",
		"riwayat_hamil_tp":                            "required|date",
		"riwayat_hamil_imunisasi":                     "required|in:Tidak Pernah,T I,TT II,TT III,TT IV",
		"riwayat_hamil_anc":                           "string|max_len:5",
		"riwayat_hamil_ancke":                         "string|max_len:5",
		"riwayat_hamil_ket_ancke":                     "required|in:Teratur,Tidak Teratur",
		"riwayat_hamil_keluhan_hamil_muda":            "required|in:Tidak Ada,Mual,Muntah,Perdarahan,Lainâlain",
		"riwayat_hamil_keluhan_hamil_tua":             "required|in:Tidak Ada,Mual,Muntah,Perdarahan,Lainâlain",
		"riwayat_kb":                                  "required|in:Belum Pernah,Suntik,Pil,AKDR,MOW,Implan,Kondom,Kalender,MAL,Coitus Interuptus",
		"riwayat_kb_lamanya":                          "string|max_len:10",
		"riwayat_kb_komplikasi":                       "required|in:Tidak Ada,Ada",
		"riwayat_kb_ket_komplikasi":                   "string|max_len:50",
		"riwayat_kb_kapaberhenti":                     "string|max_len:20",
		"riwayat_kb_alasanberhenti":                   "string|max_len:50",
		"riwayat_genekologi":                          "required|in:Tidak Ada,Infertilitas,Infeksi Virus,PMS,Cervisitis Kronis,Endometriosis,Mioma,Polip Cervix,Kanker Kandungan,Operasi Kandungan",
		"riwayat_kebiasaan_obat":                      "required|in:-,Obat Obatan,Vitamin,Jamu Jamuan",
		"riwayat_kebiasaan_ket_obat":                  "string|max_len:100",
		"riwayat_kebiasaan_merokok":                   "required|in:Tidak,Ya",
		"riwayat_kebiasaan_ket_merokok":               "string|max_len:5",
		"riwayat_kebiasaan_alkohol":                   "required|in:Tidak,Ya",
		"riwayat_kebiasaan_ket_alkohol":               "string|max_len:5",
		"riwayat_kebiasaan_narkoba":                   "required|in:Tidak,Ya",
		"pemeriksaan_kebidanan_mental":                "string|max_len:40",
		"pemeriksaan_kebidanan_keadaan_umum":          "required|in:Baik,Sedang,Buruk",
		"pemeriksaan_kebidanan_gcs":                   "string|max_len:10",
		"pemeriksaan_kebidanan_td":                    "string|max_len:8",
		"pemeriksaan_kebidanan_nadi":                  "string|max_len:5",
		"pemeriksaan_kebidanan_rr":                    "string|max_len:5",
		"pemeriksaan_kebidanan_suhu":                  "string|max_len:5",
		"pemeriksaan_kebidanan_spo2":                  "string|max_len:5",
		"pemeriksaan_kebidanan_bb":                    "string|max_len:5",
		"pemeriksaan_kebidanan_tb":                    "string|max_len:5",
		"pemeriksaan_kebidanan_lila":                  "string|max_len:5",
		"pemeriksaan_kebidanan_tfu":                   "string|max_len:10",
		"pemeriksaan_kebidanan_tbj":                   "string|max_len:10",
		"pemeriksaan_kebidanan_letak":                 "string|max_len:10",
		"pemeriksaan_kebidanan_presentasi":            "string|max_len:10",
		"pemeriksaan_kebidanan_penurunan":             "string|max_len:10",
		"pemeriksaan_kebidanan_his":                   "string|max_len:10",
		"pemeriksaan_kebidanan_kekuatan":              "string|max_len:10",
		"pemeriksaan_kebidanan_lamanya":               "string|max_len:5",
		"pemeriksaan_kebidanan_djj":                   "string|max_len:5",
		"pemeriksaan_kebidanan_ket_djj":               "required|in:Teratur,Tidak Teratur",
		"pemeriksaan_kebidanan_portio":                "string|max_len:10",
		"pemeriksaan_kebidanan_pembukaan":             "string|max_len:5",
		"pemeriksaan_kebidanan_ketuban":               "string|max_len:10",
		"pemeriksaan_kebidanan_hodge":                 "string|max_len:10",
		"pemeriksaan_kebidanan_panggul":               "required|in:Luas,Sedang,Sempit,Tidak Dilakukan Pemeriksaan",
		"pemeriksaan_kebidanan_inspekulo":             "required|in:Dilakukan,Tidak",
		"pemeriksaan_kebidanan_ket_inspekulo":         "string|max_len:50",
		"pemeriksaan_kebidanan_lakmus":                "required|in:Dilakukan,Tidak",
		"pemeriksaan_kebidanan_ket_lakmus":            "string|max_len:50",
		"pemeriksaan_kebidanan_ctg":                   "required|in:Dilakukan,Tidak",
		"pemeriksaan_kebidanan_ket_ctg":               "string|max_len:50",
		"pemeriksaan_umum_kepala":                     "required|in:Normocephale,Hydrocephalus,Lain-lain",
		"pemeriksaan_umum_muka":                       "required|in:Normal,Pucat,Oedem,Lain-lain",
		"pemeriksaan_umum_mata":                       "required|in:Conjungtiva Merah Muda,Conjungtiva Pucat,Sklera Ikterik,Pandangan Kabur,Lain-lain",
		"pemeriksaan_umum_hidung":                     "required|in:Normal,Sekret,Polip,Lain-lain",
		"pemeriksaan_umum_telinga":                    "required|in:Bersih,Serumen,Polip,Lain-lain",
		"pemeriksaan_umum_mulut":                      "required|in:Bersih,Kotor,Lain-lain",
		"pemeriksaan_umum_leher":                      "required|in:Normal,Pembesaran KGB,Pembesaran Kelenjar Tiroid,Lain-lain",
		"pemeriksaan_umum_dada":                       "required|in:Mamae Simetris,Mamae Asimetris,Aerola Hiperpigmentasi,Kolustrum (+),Tumor,Puting Susu Menonjol",
		"pemeriksaan_umum_perut":                      "required|in:Luka Bekas Operasi,Nyeri Tekan (Ya),Nyeri Tekan (Tidak),Lain-lain",
		"pemeriksaan_umum_genitalia":                  "required|in:Bersih,Kotor,Varises,Oedem,Hematoma,Hemoroid,Lain-lain",
		"pemeriksaan_umum_ekstrimitas":                "required|in:Normal,Oedem,Refleks Patella Ada,Lain-lain",
		"pengkajian_fungsi_kemampuan_aktifitas":       "required|in:Mandiri,Bantuan minimal,Bantuan Sebagian,Ketergantungan Total",
		"pengkajian_fungsi_berjalan":                  "required|in:TAK,Penurunan Kekuatan/ROM,Paralisis,Sering Jatuh,Deformitas,Hilang keseimbangan,Riwayat Patah Tulang,Lain-lain",
		"pengkajian_fungsi_ket_berjalan":              "string|max_len:50",
		"pengkajian_fungsi_aktivitas":                 "required|in:Tirah Baring,Duduk,Berjalan",
		"pengkajian_fungsi_ambulasi":                  "required|in:Walker,Tongkat,Kursi Roda,Tidak Menggunakan",
		"pengkajian_fungsi_ekstrimitas_atas":          "required|in:TAK,Lemah,Oedema,Tidak Simetris,Lain-lain",
		"pengkajian_fungsi_ket_ekstrimitas_atas":      "string|max_len:50",
		"pengkajian_fungsi_ekstrimitas_bawah":         "required|in:TAK,Varises,Oedema,Tidak Simetris,Lain-lain",
		"pengkajian_fungsi_ket_ekstrimitas_bawah":     "string|max_len:50",
		"pengkajian_fungsi_kemampuan_menggenggam":     "required|in:Tidak Ada Kesulitan,Terakhir,Lain-lain",
		"pengkajian_fungsi_ket_kemampuan_menggenggam": "string|max_len:50",
		"pengkajian_fungsi_koordinasi":                "required|in:Tidak Ada Kesulitan,Ada Masalah",
		"pengkajian_fungsi_ket_koordinasi":            "string|max_len:50",
		"pengkajian_fungsi_gangguan_fungsi":           "required|in:Tidak (Tidak Perlu Co DPJP),Ya (Co DPJP)",
		"riwayat_psiko_kondisipsiko":                  "required|in:Tidak Ada Masalah,Marah,Takut,Depresi,Cepat Lelah,Cemas,Gelisah,Sulit Tidur,Lain-lain",
		"riwayat_psiko_adakah_prilaku":                "required|in:Tidak Ada Masalah,Perilaku Kekerasan,Gangguan Efek,Gangguan Memori,Halusinasi,Kecenderungan Percobaan Bunuh Diri,Lain-lain",
		"riwayat_psiko_ket_adakah_prilaku":            "string|max_len:50",
		"riwayat_psiko_gangguan_jiwa":                 "required|in:Tidak,Ya",
		"riwayat_psiko_hubungan_pasien":               "required|in:Harmonis,Kurang Harmonis,Tidak Harmonis,Konflik Besar",
		"riwayat_psiko_tinggal_dengan":                "required|in:Sendiri,Orang Tua,Suami/Istri,Keluarga,Lain-lain",
		"riwayat_psiko_ket_tinggal_dengan":            "string|max_len:50",
		"riwayat_psiko_budaya":                        "required|in:Tidak Ada,Ada",
		"riwayat_psiko_ket_budaya":                    "string|max_len:50",
		"riwayat_psiko_pend_pj":                       "required|in:-,TS,TK,SD,SMP,SMA,SLTA/SEDERAJAT,D1,D2,D3,D4,S1,S2,S3",
		"riwayat_psiko_edukasi_pada":                  "required|in:Pasien,Keluarga",
		"riwayat_psiko_ket_edukasi_pada":              "string|max_len:50",
		"penilaian_nyeri":                             "required|in:Tidak Ada Nyeri,Nyeri Akut,Nyeri Kronis",
		"penilaian_nyeri_penyebab":                    "required|in:Proses Penyakit,Benturan,Lain-lain",
		"penilaian_nyeri_ket_penyebab":                "string|max_len:50",
		"penilaian_nyeri_kualitas":                    "required|in:Seperti Tertusuk,Berdenyut,Teriris,Tertindih,Tertiban,Lain-lain",
		"penilaian_nyeri_ket_kualitas":                "string|max_len:50",
		"penilaian_nyeri_lokasi":                      "string|max_len:50",
		"penilaian_nyeri_menyebar":                    "required|in:Tidak,Ya",
		"penilaian_nyeri_skala":                       "required|in:0,1,2,3,4,5,6,7,8,9,10",
		"penilaian_nyeri_waktu":                       "string|max_len:5",
		"penilaian_nyeri_hilang":                      "required|in:Istirahat,Medengar Musik,Minum Obat",
		"penilaian_nyeri_ket_hilang":                  "string|max_len:50",
		"penilaian_nyeri_diberitahukan_dokter":        "required|in:Tidak,Ya",
		"penilaian_nyeri_jam_diberitahukan_dokter":    "string|max_len:10",
		"penilaian_jatuh_skala1":                      "required|in:Tidak,Ya",
		"penilaian_jatuh_nilai1":                      "int",
		"penilaian_jatuh_skala2":                      "required|in:Tidak,Ya",
		"penilaian_jatuh_nilai2":                      "int",
		"penilaian_jatuh_skala3":                      "required|in:Tidak Ada/Kursi Roda/Perawat/Tirah Baring,Tongkat/Alat Penopang,Berpegangan Pada Perabot",
		"penilaian_jatuh_nilai3":                      "int",
		"penilaian_jatuh_skala4":                      "required|in:Tidak,Ya",
		"penilaian_jatuh_nilai4":                      "int",
		"penilaian_jatuh_skala5":                      "required|in:Normal/Tirah Baring/Imobilisasi,Lemah,Terganggu",
		"penilaian_jatuh_nilai5":                      "int",
		"penilaian_jatuh_skala6":                      "required|in:Sadar Akan Kemampuan Diri Sendiri,Sering Lupa Akan Keterbatasan Yang Dimiliki",
		"penilaian_jatuh_nilai6":                      "int",
		"penilaian_jatuh_totalnilai":                  "numeric",
		"skrining_gizi1":                              "required|in:Tidak ada penurunan berat badan,Tidak yakin/ tidak tahu/ terasa baju lebih longgar,Ya 1-5 kg,Ya 6-10 kg,Ya 11-15 kg,Ya > 15 kg",
		"nilai_gizi1":                                 "int",
		"skrining_gizi2":                              "required|in:Tidak,Ya",
		"nilai_gizi2":                                 "int",
		"nilai_total_gizi":                            "numeric",
		"skrining_gizi_diagnosa_khusus":               "required|in:Tidak,Ya",
		"skrining_gizi_ket_diagnosa_khusus":           "string|max_len:50",
		"skrining_gizi_diketahui_dietisen":            "required|in:Tidak,Ya",
		"skrining_gizi_jam_diketahui_dietisen":        "string|max_len:10",
		"masalah":                                     "string|max_len:1000",
		"rencana":                                     "string|max_len:1000",
		"nip1":                                        "required|string|max_len:20",
		"nip2":                                        "required|string|max_len:20",
		"kd_dokter":                                   "required|string|max_len:20",
	}
	return rules
}

// PenilaianAwalKeperawatanKebidananRanapStore simpan laporan pemantauan anastesi; kolom waktu kunci kosong = sekarang.
type PenilaianAwalKeperawatanKebidananRanapStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianAwalKeperawatanKebidananRanapData
}

func (r *PenilaianAwalKeperawatanKebidananRanapStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanKebidananRanapStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianAwalKeperawatanKebidananRanapRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianAwalKeperawatanKebidananRanapStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianAwalKeperawatanKebidananRanapStore) Payload() PenilaianAwalKeperawatanKebidananRanapData {
	return r.PenilaianAwalKeperawatanKebidananRanapData
}

func (r *PenilaianAwalKeperawatanKebidananRanapStore) DetailValues() map[string][]string { return nil }

// PenilaianAwalKeperawatanKebidananRanapUpdate ubah laporan pemantauan anastesi (PUT); kunci lewat query string.
type PenilaianAwalKeperawatanKebidananRanapUpdate struct {
	PenilaianAwalKeperawatanKebidananRanapData
}

func (r *PenilaianAwalKeperawatanKebidananRanapUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanKebidananRanapUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianAwalKeperawatanKebidananRanapRules()
}

func (r *PenilaianAwalKeperawatanKebidananRanapUpdate) Payload() PenilaianAwalKeperawatanKebidananRanapData {
	return r.PenilaianAwalKeperawatanKebidananRanapData
}

func (r *PenilaianAwalKeperawatanKebidananRanapUpdate) DetailValues() map[string][]string { return nil }
