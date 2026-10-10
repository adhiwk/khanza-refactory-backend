package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianAwalKeperawatanRanapBayiData isian penilaian awal keperawatan ranap bayi anak.
type PenilaianAwalKeperawatanRanapBayiData struct {
	Tanggal                                string `form:"tanggal" json:"tanggal"`
	Informasi                              string `form:"informasi" json:"informasi"`
	KetInformasi                           string `form:"ket_informasi" json:"ket_informasi"`
	TibaDiruangRawat                       string `form:"tiba_diruang_rawat" json:"tiba_diruang_rawat"`
	Rps                                    string `form:"rps" json:"rps"`
	Rpd                                    string `form:"rpd" json:"rpd"`
	Rpk                                    string `form:"rpk" json:"rpk"`
	Rpo                                    string `form:"rpo" json:"rpo"`
	Alergi                                 string `form:"alergi" json:"alergi"`
	TumbuhKembangTengkurap                 string `form:"tumbuh_kembang_tengkurap" json:"tumbuh_kembang_tengkurap"`
	TumbuhKembangDuduk                     string `form:"tumbuh_kembang_duduk" json:"tumbuh_kembang_duduk"`
	TumbuhKembangBerdiri                   string `form:"tumbuh_kembang_berdiri" json:"tumbuh_kembang_berdiri"`
	TumbuhKembangGigiPertama               string `form:"tumbuh_kembang_gigi_pertama" json:"tumbuh_kembang_gigi_pertama"`
	TumbuhKembangBerjalan                  string `form:"tumbuh_kembang_berjalan" json:"tumbuh_kembang_berjalan"`
	TumbuhKembangBicara                    string `form:"tumbuh_kembang_bicara" json:"tumbuh_kembang_bicara"`
	TumbuhKembangMembaca                   string `form:"tumbuh_kembang_membaca" json:"tumbuh_kembang_membaca"`
	TumbuhKembangMenulis                   string `form:"tumbuh_kembang_menulis" json:"tumbuh_kembang_menulis"`
	TumbuhKembangGangguanEmosi             string `form:"tumbuh_kembang_gangguan_emosi" json:"tumbuh_kembang_gangguan_emosi"`
	PersalinanAnakke                       string `form:"persalinan_anakke" json:"persalinan_anakke"`
	PersalinanDarisaudara                  string `form:"persalinan_darisaudara" json:"persalinan_darisaudara"`
	PersalinanKelahiran                    string `form:"persalinan_kelahiran" json:"persalinan_kelahiran"`
	PersalinanKelahiranKeterangan          string `form:"persalinan_kelahiran_keterangan" json:"persalinan_kelahiran_keterangan"`
	PersalinanUmurKelahiran                string `form:"persalinan_umur_kelahiran" json:"persalinan_umur_kelahiran"`
	PersalinanKelainanBawaan               string `form:"persalinan_kelainan_bawaan" json:"persalinan_kelainan_bawaan"`
	PersalinanKelainanBawaanKeterangan     string `form:"persalinan_kelainan_bawaan_keterangan" json:"persalinan_kelainan_bawaan_keterangan"`
	PersalinanBbLahir                      string `form:"persalinan_bb_lahir" json:"persalinan_bb_lahir"`
	PersalinanPbLahir                      string `form:"persalinan_pb_lahir" json:"persalinan_pb_lahir"`
	PersalinanLainnya                      string `form:"persalinan_lainnya" json:"persalinan_lainnya"`
	FisikKesadaran                         string `form:"fisik_kesadaran" json:"fisik_kesadaran"`
	FisikGcs                               string `form:"fisik_gcs" json:"fisik_gcs"`
	FisikTd                                string `form:"fisik_td" json:"fisik_td"`
	FisikRr                                string `form:"fisik_rr" json:"fisik_rr"`
	FisikSuhu                              string `form:"fisik_suhu" json:"fisik_suhu"`
	FisikNadi                              string `form:"fisik_nadi" json:"fisik_nadi"`
	FisikBb                                string `form:"fisik_bb" json:"fisik_bb"`
	FisikTb                                string `form:"fisik_tb" json:"fisik_tb"`
	FisikLp                                string `form:"fisik_lp" json:"fisik_lp"`
	FisikLk                                string `form:"fisik_lk" json:"fisik_lk"`
	FisikLd                                string `form:"fisik_ld" json:"fisik_ld"`
	SarafPusatKepala                       string `form:"saraf_pusat_kepala" json:"saraf_pusat_kepala"`
	SarafPusatKepalaKeterangan             string `form:"saraf_pusat_kepala_keterangan" json:"saraf_pusat_kepala_keterangan"`
	SarafPusatWajah                        string `form:"saraf_pusat_wajah" json:"saraf_pusat_wajah"`
	SarafPusatWajahKeterangan              string `form:"saraf_pusat_wajah_keterangan" json:"saraf_pusat_wajah_keterangan"`
	SarafPusatLeher                        string `form:"saraf_pusat_leher" json:"saraf_pusat_leher"`
	SarafPusatKejang                       string `form:"saraf_pusat_kejang" json:"saraf_pusat_kejang"`
	SarafPusatKejangKeterangan             string `form:"saraf_pusat_kejang_keterangan" json:"saraf_pusat_kejang_keterangan"`
	SarafPusatSensorik                     string `form:"saraf_pusat_sensorik" json:"saraf_pusat_sensorik"`
	KardiovaskulerPulsasi                  string `form:"kardiovaskuler_pulsasi" json:"kardiovaskuler_pulsasi"`
	KardiovaskulerSirkulasi                string `form:"kardiovaskuler_sirkulasi" json:"kardiovaskuler_sirkulasi"`
	KardiovaskulerSirkulasiKeterangan      string `form:"kardiovaskuler_sirkulasi_keterangan" json:"kardiovaskuler_sirkulasi_keterangan"`
	KardiovaskulerDenyutNadi               string `form:"kardiovaskuler_denyut_nadi" json:"kardiovaskuler_denyut_nadi"`
	RespirasiRetraksi                      string `form:"respirasi_retraksi" json:"respirasi_retraksi"`
	RespirasiPolaNafas                     string `form:"respirasi_pola_nafas" json:"respirasi_pola_nafas"`
	RespirasiSuaraNafas                    string `form:"respirasi_suara_nafas" json:"respirasi_suara_nafas"`
	RespirasiBatuk                         string `form:"respirasi_batuk" json:"respirasi_batuk"`
	RespirasiVolume                        string `form:"respirasi_volume" json:"respirasi_volume"`
	RespirasiJenisPernapasan               string `form:"respirasi_jenis_pernapasan" json:"respirasi_jenis_pernapasan"`
	RespirasiJenisPernapasanKeterangan     string `form:"respirasi_jenis_pernapasan_keterangan" json:"respirasi_jenis_pernapasan_keterangan"`
	RespirasiIrama                         string `form:"respirasi_irama" json:"respirasi_irama"`
	GastroMulut                            string `form:"gastro_mulut" json:"gastro_mulut"`
	GastroMulutKeterangan                  string `form:"gastro_mulut_keterangan" json:"gastro_mulut_keterangan"`
	GastroTenggorakan                      string `form:"gastro_tenggorakan" json:"gastro_tenggorakan"`
	GastroTenggorakanKeterangan            string `form:"gastro_tenggorakan_keterangan" json:"gastro_tenggorakan_keterangan"`
	GastroLidah                            string `form:"gastro_lidah" json:"gastro_lidah"`
	GastroLidahKeterangan                  string `form:"gastro_lidah_keterangan" json:"gastro_lidah_keterangan"`
	GastroAbdomen                          string `form:"gastro_abdomen" json:"gastro_abdomen"`
	GastroAbdomenKeterangan                string `form:"gastro_abdomen_keterangan" json:"gastro_abdomen_keterangan"`
	GastroGigi                             string `form:"gastro_gigi" json:"gastro_gigi"`
	GastroGigiKeterangan                   string `form:"gastro_gigi_keterangan" json:"gastro_gigi_keterangan"`
	GastroUsus                             string `form:"gastro_usus" json:"gastro_usus"`
	GastroAnus                             string `form:"gastro_anus" json:"gastro_anus"`
	NeurologiSensorik                      string `form:"neurologi_sensorik" json:"neurologi_sensorik"`
	NeurologiPengilihatan                  string `form:"neurologi_pengilihatan" json:"neurologi_pengilihatan"`
	NeurologiPengilihatanKeterangan        string `form:"neurologi_pengilihatan_keterangan" json:"neurologi_pengilihatan_keterangan"`
	NeurologiAlatBantuPenglihatan          string `form:"neurologi_alat_bantu_penglihatan" json:"neurologi_alat_bantu_penglihatan"`
	NeurologiMotorik                       string `form:"neurologi_motorik" json:"neurologi_motorik"`
	NeurologiPendengaran                   string `form:"neurologi_pendengaran" json:"neurologi_pendengaran"`
	NeurologiBicara                        string `form:"neurologi_bicara" json:"neurologi_bicara"`
	NeurologiBicaraKeterangan              string `form:"neurologi_bicara_keterangan" json:"neurologi_bicara_keterangan"`
	NeurologiOtot                          string `form:"neurologi_otot" json:"neurologi_otot"`
	InteKulit                              string `form:"inte_kulit" json:"inte_kulit"`
	InteWarnaKulit                         string `form:"inte_warna_kulit" json:"inte_warna_kulit"`
	InteTugor                              string `form:"inte_tugor" json:"inte_tugor"`
	InteDecubi                             string `form:"inte_decubi" json:"inte_decubi"`
	MuskuOdema                             string `form:"musku_odema" json:"musku_odema"`
	MuskuOdemaKeterangan                   string `form:"musku_odema_keterangan" json:"musku_odema_keterangan"`
	MuskuPegerakansendi                    string `form:"musku_pegerakansendi" json:"musku_pegerakansendi"`
	MuskuOtot                              string `form:"musku_otot" json:"musku_otot"`
	MuskuFraktur                           string `form:"musku_fraktur" json:"musku_fraktur"`
	MuskuFrakturKeterangan                 string `form:"musku_fraktur_keterangan" json:"musku_fraktur_keterangan"`
	MuskuNyerisendi                        string `form:"musku_nyerisendi" json:"musku_nyerisendi"`
	MuskuNyerisendiKeterangan              string `form:"musku_nyerisendi_keterangan" json:"musku_nyerisendi_keterangan"`
	EliminasiBabFrekuensi                  string `form:"eliminasi_bab_frekuensi" json:"eliminasi_bab_frekuensi"`
	EliminasiBabFrekuensiPer               string `form:"eliminasi_bab_frekuensi_per" json:"eliminasi_bab_frekuensi_per"`
	EliminasiBabKonsistesi                 string `form:"eliminasi_bab_konsistesi" json:"eliminasi_bab_konsistesi"`
	EliminasiBabWarna                      string `form:"eliminasi_bab_warna" json:"eliminasi_bab_warna"`
	EliminasiBakFrekuensi                  string `form:"eliminasi_bak_frekuensi" json:"eliminasi_bak_frekuensi"`
	EliminasiBakFrekuensiPer               string `form:"eliminasi_bak_frekuensi_per" json:"eliminasi_bak_frekuensi_per"`
	EliminasiBakWarna                      string `form:"eliminasi_bak_warna" json:"eliminasi_bak_warna"`
	EliminasiBakLainlain                   string `form:"eliminasi_bak_lainlain" json:"eliminasi_bak_lainlain"`
	PsikoKondisi                           string `form:"psiko_kondisi" json:"psiko_kondisi"`
	PsikoPerilaku                          string `form:"psiko_perilaku" json:"psiko_perilaku"`
	PsikoPerilakuKeterangan                string `form:"psiko_perilaku_keterangan" json:"psiko_perilaku_keterangan"`
	PsikoGangguanJiwa                      string `form:"psiko_gangguan_jiwa" json:"psiko_gangguan_jiwa"`
	PsikoHubunganPasien                    string `form:"psiko_hubungan_pasien" json:"psiko_hubungan_pasien"`
	PsikoTinggalDengan                     string `form:"psiko_tinggal_dengan" json:"psiko_tinggal_dengan"`
	PsikoTinggalDenganKeterangan           string `form:"psiko_tinggal_dengan_keterangan" json:"psiko_tinggal_dengan_keterangan"`
	PsikoPekerjaanPj                       string `form:"psiko_pekerjaan_pj" json:"psiko_pekerjaan_pj"`
	PsikoNilaiKepercayaan                  string `form:"psiko_nilai_kepercayaan" json:"psiko_nilai_kepercayaan"`
	PsikoNilaiKepercayaanKeterangan        string `form:"psiko_nilai_kepercayaan_keterangan" json:"psiko_nilai_kepercayaan_keterangan"`
	PsikoPendidikanPj                      string `form:"psiko_pendidikan_pj" json:"psiko_pendidikan_pj"`
	PsikoEdukasi                           string `form:"psiko_edukasi" json:"psiko_edukasi"`
	PsikoEdukasiKeterangan                 string `form:"psiko_edukasi_keterangan" json:"psiko_edukasi_keterangan"`
	EdukasiBahasa                          string `form:"edukasi_bahasa" json:"edukasi_bahasa"`
	EdukasiBacaTulis                       string `form:"edukasi_baca_tulis" json:"edukasi_baca_tulis"`
	EdukasiPenerjemah                      string `form:"edukasi_penerjemah" json:"edukasi_penerjemah"`
	EdukasiPenerjemahKeterangan            string `form:"edukasi_penerjemah_keterangan" json:"edukasi_penerjemah_keterangan"`
	EdukasiTerdapatHambatan                string `form:"edukasi_terdapat_hambatan" json:"edukasi_terdapat_hambatan"`
	EdukasiHambatanBelajar                 string `form:"edukasi_hambatan_belajar" json:"edukasi_hambatan_belajar"`
	EdukasiHambatanBelajarKeterangan       string `form:"edukasi_hambatan_belajar_keterangan" json:"edukasi_hambatan_belajar_keterangan"`
	EdukasiHambatanBicara                  string `form:"edukasi_hambatan_bicara" json:"edukasi_hambatan_bicara"`
	EdukasiBahasaIsyarat                   string `form:"edukasi_bahasa_isyarat" json:"edukasi_bahasa_isyarat"`
	EdukasiCaraBelajar                     string `form:"edukasi_cara_belajar" json:"edukasi_cara_belajar"`
	EdukasiMenerimaInformasi               string `form:"edukasi_menerima_informasi" json:"edukasi_menerima_informasi"`
	EdukasiMenerimaInformasiKeterangan     string `form:"edukasi_menerima_informasi_keterangan" json:"edukasi_menerima_informasi_keterangan"`
	EdukasiNutrisi                         string `form:"edukasi_nutrisi" json:"edukasi_nutrisi"`
	EdukasiPenyakit                        string `form:"edukasi_penyakit" json:"edukasi_penyakit"`
	EdukasiPengobatan                      string `form:"edukasi_pengobatan" json:"edukasi_pengobatan"`
	EdukasiPerawatan                       string `form:"edukasi_perawatan" json:"edukasi_perawatan"`
	SkriningGizi1                          string `form:"skrining_gizi1" json:"skrining_gizi1"`
	NilaiGizi1                             string `form:"nilai_gizi1" json:"nilai_gizi1"`
	SkriningGizi2                          string `form:"skrining_gizi2" json:"skrining_gizi2"`
	NilaiGizi2                             string `form:"nilai_gizi2" json:"nilai_gizi2"`
	SkriningGizi3                          string `form:"skrining_gizi3" json:"skrining_gizi3"`
	NilaiGizi3                             string `form:"nilai_gizi3" json:"nilai_gizi3"`
	SkriningGizi4                          string `form:"skrining_gizi4" json:"skrining_gizi4"`
	NilaiGizi4                             string `form:"nilai_gizi4" json:"nilai_gizi4"`
	TotalNilai                             string `form:"total_nilai" json:"total_nilai"`
	KeteranganSkriningGizi                 string `form:"keterangan_skrining_gizi" json:"keterangan_skrining_gizi"`
	PenilaianHumptydumptySkala1            string `form:"penilaian_humptydumpty_skala1" json:"penilaian_humptydumpty_skala1"`
	PenilaianHumptydumptyNilai1            *int   `form:"penilaian_humptydumpty_nilai1" json:"penilaian_humptydumpty_nilai1"`
	PenilaianHumptydumptySkala2            string `form:"penilaian_humptydumpty_skala2" json:"penilaian_humptydumpty_skala2"`
	PenilaianHumptydumptyNilai2            *int   `form:"penilaian_humptydumpty_nilai2" json:"penilaian_humptydumpty_nilai2"`
	PenilaianHumptydumptySkala3            string `form:"penilaian_humptydumpty_skala3" json:"penilaian_humptydumpty_skala3"`
	PenilaianHumptydumptyNilai3            *int   `form:"penilaian_humptydumpty_nilai3" json:"penilaian_humptydumpty_nilai3"`
	PenilaianHumptydumptySkala4            string `form:"penilaian_humptydumpty_skala4" json:"penilaian_humptydumpty_skala4"`
	PenilaianHumptydumptyNilai4            *int   `form:"penilaian_humptydumpty_nilai4" json:"penilaian_humptydumpty_nilai4"`
	PenilaianHumptydumptySkala5            string `form:"penilaian_humptydumpty_skala5" json:"penilaian_humptydumpty_skala5"`
	PenilaianHumptydumptyNilai5            *int   `form:"penilaian_humptydumpty_nilai5" json:"penilaian_humptydumpty_nilai5"`
	PenilaianHumptydumptySkala6            string `form:"penilaian_humptydumpty_skala6" json:"penilaian_humptydumpty_skala6"`
	PenilaianHumptydumptyNilai6            *int   `form:"penilaian_humptydumpty_nilai6" json:"penilaian_humptydumpty_nilai6"`
	PenilaianHumptydumptySkala7            string `form:"penilaian_humptydumpty_skala7" json:"penilaian_humptydumpty_skala7"`
	PenilaianHumptydumptyNilai7            *int   `form:"penilaian_humptydumpty_nilai7" json:"penilaian_humptydumpty_nilai7"`
	PenilaianHumptydumptyTotalnilai        *int   `form:"penilaian_humptydumpty_totalnilai" json:"penilaian_humptydumpty_totalnilai"`
	HasilSkriningPenilaianHumptydumpty     string `form:"hasil_skrining_penilaian_humptydumpty" json:"hasil_skrining_penilaian_humptydumpty"`
	NyeriWajah                             string `form:"nyeri_wajah" json:"nyeri_wajah"`
	NyeriNilaiWajah                        string `form:"nyeri_nilai_wajah" json:"nyeri_nilai_wajah"`
	NyeriKaki                              string `form:"nyeri_kaki" json:"nyeri_kaki"`
	NyeriNilaiKaki                         string `form:"nyeri_nilai_kaki" json:"nyeri_nilai_kaki"`
	NyeriAktifitas                         string `form:"nyeri_aktifitas" json:"nyeri_aktifitas"`
	NyeriNilaiAktifitas                    string `form:"nyeri_nilai_aktifitas" json:"nyeri_nilai_aktifitas"`
	NyeriMenangis                          string `form:"nyeri_menangis" json:"nyeri_menangis"`
	NyeriNilaiMenangis                     string `form:"nyeri_nilai_menangis" json:"nyeri_nilai_menangis"`
	NyeriBersuara                          string `form:"nyeri_bersuara" json:"nyeri_bersuara"`
	NyeriNilaiBersuara                     string `form:"nyeri_nilai_bersuara" json:"nyeri_nilai_bersuara"`
	NyeriNilaiTotal                        string `form:"nyeri_nilai_total" json:"nyeri_nilai_total"`
	NyeriKondisi                           string `form:"nyeri_kondisi" json:"nyeri_kondisi"`
	NyeriLokasi                            string `form:"nyeri_lokasi" json:"nyeri_lokasi"`
	NyeriDurasi                            string `form:"nyeri_durasi" json:"nyeri_durasi"`
	NyeriFrekuensi                         string `form:"nyeri_frekuensi" json:"nyeri_frekuensi"`
	NyeriHilang                            string `form:"nyeri_hilang" json:"nyeri_hilang"`
	NyeriHilangKeterangan                  string `form:"nyeri_hilang_keterangan" json:"nyeri_hilang_keterangan"`
	NyeriDiberitahukanPadaDokter           string `form:"nyeri_diberitahukan_pada_dokter" json:"nyeri_diberitahukan_pada_dokter"`
	NyeriDiberitahukanPadaDokterKeterangan string `form:"nyeri_diberitahukan_pada_dokter_keterangan" json:"nyeri_diberitahukan_pada_dokter_keterangan"`
	InformasiPerencanaanPulang             string `form:"informasi_perencanaan_pulang" json:"informasi_perencanaan_pulang"`
	LamaRatarata                           string `form:"lama_ratarata" json:"lama_ratarata"`
	PerencanaanPulang                      string `form:"perencanaan_pulang" json:"perencanaan_pulang"`
	KondisiKlinisPulang                    string `form:"kondisi_klinis_pulang" json:"kondisi_klinis_pulang"`
	PerawatanLanjutanDirumah               string `form:"perawatan_lanjutan_dirumah" json:"perawatan_lanjutan_dirumah"`
	CaraTransportasiPulang                 string `form:"cara_transportasi_pulang" json:"cara_transportasi_pulang"`
	TransportasiDigunakan                  string `form:"transportasi_digunakan" json:"transportasi_digunakan"`
	Rencana                                string `form:"rencana" json:"rencana"`
	Nip1                                   string `form:"nip1" json:"nip1"`
	Nip2                                   string `form:"nip2" json:"nip2"`
	KdDokter                               string `form:"kd_dokter" json:"kd_dokter"`
}

// PenilaianAwalKeperawatanRanapBayiDetail daftar kode detail (diganti seluruhnya saat simpan/ubah).
type PenilaianAwalKeperawatanRanapBayiDetail struct {
	KodeMasalah []string `form:"kode_masalah" json:"kode_masalah"`
	KodeRencana []string `form:"kode_rencana" json:"kode_rencana"`
}

func penilaianAwalKeperawatanRanapBayiRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                    "required|date",
		"informasi":                                  "required|in:Autoanamnesis,Alloanamnesis",
		"ket_informasi":                              "string|max_len:30",
		"tiba_diruang_rawat":                         "required|in:Jalan Tanpa Bantuan,Kursi Roda,Brankar",
		"rps":                                        "string|max_len:300",
		"rpd":                                        "string|max_len:150",
		"rpk":                                        "string|max_len:150",
		"rpo":                                        "string|max_len:150",
		"alergi":                                     "string|max_len:40",
		"tumbuh_kembang_tengkurap":                   "string|max_len:15",
		"tumbuh_kembang_duduk":                       "string|max_len:15",
		"tumbuh_kembang_berdiri":                     "string|max_len:15",
		"tumbuh_kembang_gigi_pertama":                "string|max_len:15",
		"tumbuh_kembang_berjalan":                    "string|max_len:15",
		"tumbuh_kembang_bicara":                      "string|max_len:15",
		"tumbuh_kembang_membaca":                     "string|max_len:15",
		"tumbuh_kembang_menulis":                     "string|max_len:15",
		"tumbuh_kembang_gangguan_emosi":              "string|max_len:30",
		"persalinan_anakke":                          "string|max_len:2",
		"persalinan_darisaudara":                     "string|max_len:2",
		"persalinan_kelahiran":                       "required|in:Spontan,Sectio Caesaria,Lain-Lain",
		"persalinan_kelahiran_keterangan":            "string|max_len:30",
		"persalinan_umur_kelahiran":                  "required|in:Cukup Bulan,Kurang Bulan",
		"persalinan_kelainan_bawaan":                 "required|in:Tidak Ada,Ada",
		"persalinan_kelainan_bawaan_keterangan":      "string|max_len:30",
		"persalinan_bb_lahir":                        "string|max_len:5",
		"persalinan_pb_lahir":                        "string|max_len:5",
		"persalinan_lainnya":                         "string|max_len:100",
		"fisik_kesadaran":                            "required|in:Compos Mentis,Apatis,Somnolen,Sopor,Soporcoma,Coma",
		"fisik_gcs":                                  "string|max_len:6",
		"fisik_td":                                   "string|max_len:8",
		"fisik_rr":                                   "string|max_len:5",
		"fisik_suhu":                                 "string|max_len:5",
		"fisik_nadi":                                 "string|max_len:5",
		"fisik_bb":                                   "string|max_len:5",
		"fisik_tb":                                   "string|max_len:5",
		"fisik_lp":                                   "string|max_len:5",
		"fisik_lk":                                   "string|max_len:5",
		"fisik_ld":                                   "string|max_len:5",
		"saraf_pusat_kepala":                         "required|in:TAK,Hydrocephalus,Hematoma,Lain-lain",
		"saraf_pusat_kepala_keterangan":              "string|max_len:30",
		"saraf_pusat_wajah":                          "required|in:TAK,Asimetris,Kelainan Kongenital",
		"saraf_pusat_wajah_keterangan":               "string|max_len:30",
		"saraf_pusat_leher":                          "required|in:TAK,Kaku Kuduk,Pembesaran Thyroid,Pembesaran KGB",
		"saraf_pusat_kejang":                         "required|in:TAK,Kuat,Ada",
		"saraf_pusat_kejang_keterangan":              "string|max_len:30",
		"saraf_pusat_sensorik":                       "required|in:TAK,Sakit Nyeri,Rasa Kebas",
		"kardiovaskuler_pulsasi":                     "required|in:Kuat,Lemah,Lain-lain",
		"kardiovaskuler_sirkulasi":                   "required|in:Akral Hangat,Akral Dingin,Edema",
		"kardiovaskuler_sirkulasi_keterangan":        "string|max_len:30",
		"kardiovaskuler_denyut_nadi":                 "required|in:Teratur,Tidak Teratur",
		"respirasi_retraksi":                         "required|in:Tidak Ada,Ringan,Berat",
		"respirasi_pola_nafas":                       "required|in:Normal,Bradipnea,Tachipnea",
		"respirasi_suara_nafas":                      "required|in:Vesikuler,Wheezing,Rhonki",
		"respirasi_batuk":                            "required|in:Tidak,Ya : Produktif,Ya : Non Produktif",
		"respirasi_volume":                           "required|in:Normal,Hiperventilasi,Hipoventilasi",
		"respirasi_jenis_pernapasan":                 "required|in:Pernafasan Dada,Alat Bantu Pernafasaan",
		"respirasi_jenis_pernapasan_keterangan":      "string|max_len:30",
		"respirasi_irama":                            "required|in:Teratur,Tidak Teratur",
		"gastro_mulut":                               "required|in:TAK,Stomatitis,Mukosa Kering,Bibir Pucat,Lain-lain",
		"gastro_mulut_keterangan":                    "string|max_len:30",
		"gastro_tenggorakan":                         "required|in:TAK,Gangguan Menelan,Sakit Menelan,Lain-lain",
		"gastro_tenggorakan_keterangan":              "string|max_len:30",
		"gastro_lidah":                               "required|in:TAK,Kotor,Gerak Asimetris,Lain-lain",
		"gastro_lidah_keterangan":                    "string|max_len:30",
		"gastro_abdomen":                             "required|in:Supel,Asictes,Tegang,Nyeri Tekan/Lepas,Lain-lain",
		"gastro_abdomen_keterangan":                  "string|max_len:30",
		"gastro_gigi":                                "required|in:TAK,Karies,Goyang,Lain-lain",
		"gastro_gigi_keterangan":                     "string|max_len:30",
		"gastro_usus":                                "required|in:TAK,Tidak Ada Bising Usus,Hiperistaltik",
		"gastro_anus":                                "required|in:TAK,Atresia Ani",
		"neurologi_sensorik":                         "required|in:TAK,Sakit Nyeri,Rasa Kebas,Lain-lain",
		"neurologi_pengilihatan":                     "required|in:TAK,Ada Kelainan",
		"neurologi_pengilihatan_keterangan":          "string|max_len:30",
		"neurologi_alat_bantu_penglihatan":           "required|in:Tidak,Kacamata,Lensa Kontak",
		"neurologi_motorik":                          "required|in:TAK,Hemiparese,Tetraparese,Tremor,Lain-lain",
		"neurologi_pendengaran":                      "required|in:TAK,Berdengung,Nyeri,Tuli,Keluar Cairan,Lain-lain",
		"neurologi_bicara":                           "required|in:Jelas,Tidak Jelas",
		"neurologi_bicara_keterangan":                "string|max_len:30",
		"neurologi_otot":                             "required|in:Kuat,Lemah",
		"inte_kulit":                                 "required|in:Normal,Rash/Kemerahan,Luka,Memar,Ptekie,Bula",
		"inte_warna_kulit":                           "required|in:Normal,Pucat,Sianosis,Lain-lain",
		"inte_tugor":                                 "required|in:Baik,Sedang,Buruk",
		"inte_decubi":                                "required|string",
		"musku_odema":                                "required|in:Tidak Ada,Ada",
		"musku_odema_keterangan":                     "string|max_len:30",
		"musku_pegerakansendi":                       "required|in:Bebas,Terbatas",
		"musku_otot":                                 "required|in:Baik,Lemah,Tremor",
		"musku_fraktur":                              "required|in:Tidak Ada,Ada",
		"musku_fraktur_keterangan":                   "string|max_len:30",
		"musku_nyerisendi":                           "required|in:Tidak Ada,Ada",
		"musku_nyerisendi_keterangan":                "string|max_len:30",
		"eliminasi_bab_frekuensi":                    "string|max_len:3",
		"eliminasi_bab_frekuensi_per":                "string|max_len:10",
		"eliminasi_bab_konsistesi":                   "string|max_len:20",
		"eliminasi_bab_warna":                        "string|max_len:20",
		"eliminasi_bak_frekuensi":                    "string|max_len:3",
		"eliminasi_bak_frekuensi_per":                "string|max_len:10",
		"eliminasi_bak_warna":                        "string|max_len:20",
		"eliminasi_bak_lainlain":                     "string|max_len:20",
		"psiko_kondisi":                              "required|in:Tidak Ada Masalah,Marah,Takut,Depresi,Cepat Lelah,Cemas,Gelisah,Sulit Tidur,Lain-lain",
		"psiko_perilaku":                             "required|in:Tidak Ada Masalah,Perilaku Kekerasan,Gangguan Efek,Gangguan Memori,Halusinasi,Kecenderungan Percobaan Bunuh Diri,Lain-lain",
		"psiko_perilaku_keterangan":                  "string|max_len:30",
		"psiko_gangguan_jiwa":                        "required|in:Tidak,Ya",
		"psiko_hubungan_pasien":                      "required|in:Harmonis,Kurang Harmonis,Tidak Harmonis,Konflik Besar",
		"psiko_tinggal_dengan":                       "required|in:Sendiri,Orang Tua,Panti Asuhan,Keluarga,Lain-lain",
		"psiko_tinggal_dengan_keterangan":            "string|max_len:30",
		"psiko_pekerjaan_pj":                         "string|max_len:30",
		"psiko_nilai_kepercayaan":                    "required|in:Tidak Ada,Ada",
		"psiko_nilai_kepercayaan_keterangan":         "string|max_len:30",
		"psiko_pendidikan_pj":                        "required|in:-,TS,TK,SD,SMP,SMA,SLTA/SEDERAJAT,D1,D2,D3,D4,S1,S2,S3",
		"psiko_edukasi":                              "required|in:Pasien,Keluarga,Lainnya",
		"psiko_edukasi_keterangan":                   "string|max_len:30",
		"edukasi_bahasa":                             "string|max_len:30",
		"edukasi_baca_tulis":                         "required|in:Baik,Kurang,Tidak Bisa",
		"edukasi_penerjemah":                         "required|in:Tidak,Ya",
		"edukasi_penerjemah_keterangan":              "string|max_len:30",
		"edukasi_terdapat_hambatan":                  "required|in:Tidak,Ya",
		"edukasi_hambatan_belajar":                   "required|in:-,Gangguan Pendengaran,Gangguan Penglihatan,Gangguan Kognitif,Gangguan Fisik,Gangguan Emosi,Keterbatasan Bahasa,Keterbatasan Budaya,Keterbatasan Spiritual,Agama,Lainnya",
		"edukasi_hambatan_belajar_keterangan":        "string|max_len:30",
		"edukasi_hambatan_bicara":                    "required|in:Normal,Gangguan Bicara",
		"edukasi_bahasa_isyarat":                     "required|in:Tidak,Ya",
		"edukasi_cara_belajar":                       "required|in:Audio,Lisan,Visual,Demonstrasi,Tulisan",
		"edukasi_menerima_informasi":                 "required|in:Ya,Tidak",
		"edukasi_menerima_informasi_keterangan":      "string|max_len:30",
		"edukasi_nutrisi":                            "required|in:Ya,Tidak",
		"edukasi_penyakit":                           "required|in:Ya,Tidak",
		"edukasi_pengobatan":                         "required|in:Ya,Tidak",
		"edukasi_perawatan":                          "required|in:Ya,Tidak",
		"skrining_gizi1":                             "required|in:Tidak,Ya",
		"nilai_gizi1":                                "string|max_len:1",
		"skrining_gizi2":                             "required|in:Tidak,Ya",
		"nilai_gizi2":                                "string|max_len:1",
		"skrining_gizi3":                             "required|in:Tidak,Ya",
		"nilai_gizi3":                                "string|max_len:1",
		"skrining_gizi4":                             "required|in:Tidak,Ya",
		"nilai_gizi4":                                "string|max_len:1",
		"total_nilai":                                "string|max_len:2",
		"keterangan_skrining_gizi":                   "string|max_len:50",
		"penilaian_humptydumpty_skala1":              "in:0 - 3 Tahun,3 - 7 Tahun,7 - 13 Tahun,> 13 Tahun",
		"penilaian_humptydumpty_nilai1":              "int",
		"penilaian_humptydumpty_skala2":              "in:Laki-laki,Perempuan",
		"penilaian_humptydumpty_nilai2":              "int",
		"penilaian_humptydumpty_skala3":              "string",
		"penilaian_humptydumpty_nilai3":              "int",
		"penilaian_humptydumpty_skala4":              "in:Tidak Sadar Terhadap Keterbatasan,Lupa Keterbatasan,Mengetahui Kemampuan Diri",
		"penilaian_humptydumpty_nilai4":              "int",
		"penilaian_humptydumpty_skala5":              "in:Riwayat Jatuh Dari Tempat Tidur Saat Bayi/Anak,Pasien Menggunakan Alat Bantu/Box/Mebel,Pasien Berada Di Tempat Tidur,Di Luar Ruang Rawat",
		"penilaian_humptydumpty_nilai5":              "int",
		"penilaian_humptydumpty_skala6":              "in:Dalam 24 Jam,Dalam 48 Jam,> 48 Jam",
		"penilaian_humptydumpty_nilai6":              "int",
		"penilaian_humptydumpty_skala7":              "string",
		"penilaian_humptydumpty_nilai7":              "int",
		"penilaian_humptydumpty_totalnilai":          "int",
		"hasil_skrining_penilaian_humptydumpty":      "string|max_len:50",
		"nyeri_wajah":                                "required|in:Tersenyum/tidak ada ekspresi khusus,Terkadang meringis/menarik diri,Sering menggetarkan dagu dan mengatupkan rahang",
		"nyeri_nilai_wajah":                          "string|max_len:1",
		"nyeri_kaki":                                 "required|in:Gerakan normal/relaksasi,Tidak tenang/tegang,Kaki dibuat menendang/menarik",
		"nyeri_nilai_kaki":                           "string|max_len:1",
		"nyeri_aktifitas":                            "required|string",
		"nyeri_nilai_aktifitas":                      "string|max_len:1",
		"nyeri_menangis":                             "required|string",
		"nyeri_nilai_menangis":                       "string|max_len:1",
		"nyeri_bersuara":                             "required|string",
		"nyeri_nilai_bersuara":                       "string|max_len:1",
		"nyeri_nilai_total":                          "string|max_len:2",
		"nyeri_kondisi":                              "required|in:Tidak Ada Nyeri,Nyeri Akut,Nyeri Kronis",
		"nyeri_lokasi":                               "string|max_len:50",
		"nyeri_durasi":                               "string|max_len:25",
		"nyeri_frekuensi":                            "string|max_len:25",
		"nyeri_hilang":                               "required|in:Minum Obat,Istirahat,Mendengar Music,Berubah Posisi/Tidur,Lain-lain",
		"nyeri_hilang_keterangan":                    "string|max_len:40",
		"nyeri_diberitahukan_pada_dokter":            "required|in:Tidak,Ya",
		"nyeri_diberitahukan_pada_dokter_keterangan": "string|max_len:15",
		"informasi_perencanaan_pulang":               "required|in:Ya,Tidak",
		"lama_ratarata":                              "string|max_len:3",
		"perencanaan_pulang":                         "required|date",
		"kondisi_klinis_pulang":                      "string|max_len:100",
		"perawatan_lanjutan_dirumah":                 "string|max_len:300",
		"cara_transportasi_pulang":                   "required|in:Mandiri,Dibantu Sebagian,Dibantu Keseluruhan,Menggunakan Rostul,Menggunakan Brankar,Berjalan",
		"transportasi_digunakan":                     "required|in:Kendaraan Pribadi,Mobil Ambulance,Kendaraan Umum",
		"rencana":                                    "string|max_len:200",
		"nip1":                                       "required|string|max_len:20",
		"nip2":                                       "required|string|max_len:20",
		"kd_dokter":                                  "required|string|max_len:20",
	}
	rules["kode_masalah"] = "slice"
	rules["kode_masalah.*"] = "string|max_len:20"
	rules["kode_rencana"] = "slice"
	rules["kode_rencana.*"] = "string|max_len:20"
	return rules
}

// PenilaianAwalKeperawatanRanapBayiStore simpan penilaian awal keperawatan ranap bayi anak; kolom waktu kunci kosong = sekarang.
type PenilaianAwalKeperawatanRanapBayiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianAwalKeperawatanRanapBayiData
	PenilaianAwalKeperawatanRanapBayiDetail
}

func (r *PenilaianAwalKeperawatanRanapBayiStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanRanapBayiStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianAwalKeperawatanRanapBayiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianAwalKeperawatanRanapBayiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianAwalKeperawatanRanapBayiStore) Payload() PenilaianAwalKeperawatanRanapBayiData {
	return r.PenilaianAwalKeperawatanRanapBayiData
}

func (r *PenilaianAwalKeperawatanRanapBayiStore) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}

// PenilaianAwalKeperawatanRanapBayiUpdate ubah penilaian awal keperawatan ranap bayi anak (PUT); kunci lewat query string.
type PenilaianAwalKeperawatanRanapBayiUpdate struct {
	PenilaianAwalKeperawatanRanapBayiData
	PenilaianAwalKeperawatanRanapBayiDetail
}

func (r *PenilaianAwalKeperawatanRanapBayiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanRanapBayiUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianAwalKeperawatanRanapBayiRules()
}

func (r *PenilaianAwalKeperawatanRanapBayiUpdate) Payload() PenilaianAwalKeperawatanRanapBayiData {
	return r.PenilaianAwalKeperawatanRanapBayiData
}

func (r *PenilaianAwalKeperawatanRanapBayiUpdate) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}
