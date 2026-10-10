package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianAwalKeperawatanRanapData isian penilaian awal keperawatan ranap.
type PenilaianAwalKeperawatanRanapData struct {
	Tanggal                                          string   `form:"tanggal" json:"tanggal"`
	Informasi                                        string   `form:"informasi" json:"informasi"`
	KetInformasi                                     string   `form:"ket_informasi" json:"ket_informasi"`
	TibaDiruangRawat                                 string   `form:"tiba_diruang_rawat" json:"tiba_diruang_rawat"`
	KasusTrauma                                      string   `form:"kasus_trauma" json:"kasus_trauma"`
	CaraMasuk                                        string   `form:"cara_masuk" json:"cara_masuk"`
	Rps                                              string   `form:"rps" json:"rps"`
	Rpd                                              string   `form:"rpd" json:"rpd"`
	Rpk                                              string   `form:"rpk" json:"rpk"`
	Rpo                                              string   `form:"rpo" json:"rpo"`
	RiwayatPembedahan                                string   `form:"riwayat_pembedahan" json:"riwayat_pembedahan"`
	RiwayatDirawatDirs                               string   `form:"riwayat_dirawat_dirs" json:"riwayat_dirawat_dirs"`
	AlatBantuDipakai                                 string   `form:"alat_bantu_dipakai" json:"alat_bantu_dipakai"`
	RiwayatKehamilan                                 string   `form:"riwayat_kehamilan" json:"riwayat_kehamilan"`
	RiwayatKehamilanPerkiraan                        string   `form:"riwayat_kehamilan_perkiraan" json:"riwayat_kehamilan_perkiraan"`
	RiwayatTranfusi                                  string   `form:"riwayat_tranfusi" json:"riwayat_tranfusi"`
	RiwayatAlergi                                    string   `form:"riwayat_alergi" json:"riwayat_alergi"`
	RiwayatMerokok                                   string   `form:"riwayat_merokok" json:"riwayat_merokok"`
	RiwayatMerokokJumlah                             string   `form:"riwayat_merokok_jumlah" json:"riwayat_merokok_jumlah"`
	RiwayatAlkohol                                   string   `form:"riwayat_alkohol" json:"riwayat_alkohol"`
	RiwayatAlkoholJumlah                             string   `form:"riwayat_alkohol_jumlah" json:"riwayat_alkohol_jumlah"`
	RiwayatNarkoba                                   string   `form:"riwayat_narkoba" json:"riwayat_narkoba"`
	RiwayatOlahraga                                  string   `form:"riwayat_olahraga" json:"riwayat_olahraga"`
	PemeriksaanMental                                string   `form:"pemeriksaan_mental" json:"pemeriksaan_mental"`
	PemeriksaanKeadaanUmum                           string   `form:"pemeriksaan_keadaan_umum" json:"pemeriksaan_keadaan_umum"`
	PemeriksaanGcs                                   string   `form:"pemeriksaan_gcs" json:"pemeriksaan_gcs"`
	PemeriksaanTd                                    string   `form:"pemeriksaan_td" json:"pemeriksaan_td"`
	PemeriksaanNadi                                  string   `form:"pemeriksaan_nadi" json:"pemeriksaan_nadi"`
	PemeriksaanRr                                    string   `form:"pemeriksaan_rr" json:"pemeriksaan_rr"`
	PemeriksaanSuhu                                  string   `form:"pemeriksaan_suhu" json:"pemeriksaan_suhu"`
	PemeriksaanSpo2                                  string   `form:"pemeriksaan_spo2" json:"pemeriksaan_spo2"`
	PemeriksaanBb                                    string   `form:"pemeriksaan_bb" json:"pemeriksaan_bb"`
	PemeriksaanTb                                    string   `form:"pemeriksaan_tb" json:"pemeriksaan_tb"`
	PemeriksaanSusunanKepala                         string   `form:"pemeriksaan_susunan_kepala" json:"pemeriksaan_susunan_kepala"`
	PemeriksaanSusunanKepalaKeterangan               string   `form:"pemeriksaan_susunan_kepala_keterangan" json:"pemeriksaan_susunan_kepala_keterangan"`
	PemeriksaanSusunanWajah                          string   `form:"pemeriksaan_susunan_wajah" json:"pemeriksaan_susunan_wajah"`
	PemeriksaanSusunanWajahKeterangan                string   `form:"pemeriksaan_susunan_wajah_keterangan" json:"pemeriksaan_susunan_wajah_keterangan"`
	PemeriksaanSusunanLeher                          string   `form:"pemeriksaan_susunan_leher" json:"pemeriksaan_susunan_leher"`
	PemeriksaanSusunanKejang                         string   `form:"pemeriksaan_susunan_kejang" json:"pemeriksaan_susunan_kejang"`
	PemeriksaanSusunanKejangKeterangan               string   `form:"pemeriksaan_susunan_kejang_keterangan" json:"pemeriksaan_susunan_kejang_keterangan"`
	PemeriksaanSusunanSensorik                       string   `form:"pemeriksaan_susunan_sensorik" json:"pemeriksaan_susunan_sensorik"`
	PemeriksaanKardiovaskulerDenyutNadi              string   `form:"pemeriksaan_kardiovaskuler_denyut_nadi" json:"pemeriksaan_kardiovaskuler_denyut_nadi"`
	PemeriksaanKardiovaskulerSirkulasi               string   `form:"pemeriksaan_kardiovaskuler_sirkulasi" json:"pemeriksaan_kardiovaskuler_sirkulasi"`
	PemeriksaanKardiovaskulerSirkulasiKeterangan     string   `form:"pemeriksaan_kardiovaskuler_sirkulasi_keterangan" json:"pemeriksaan_kardiovaskuler_sirkulasi_keterangan"`
	PemeriksaanKardiovaskulerPulsasi                 string   `form:"pemeriksaan_kardiovaskuler_pulsasi" json:"pemeriksaan_kardiovaskuler_pulsasi"`
	PemeriksaanRespirasiPolaNafas                    string   `form:"pemeriksaan_respirasi_pola_nafas" json:"pemeriksaan_respirasi_pola_nafas"`
	PemeriksaanRespirasiRetraksi                     string   `form:"pemeriksaan_respirasi_retraksi" json:"pemeriksaan_respirasi_retraksi"`
	PemeriksaanRespirasiSuaraNafas                   string   `form:"pemeriksaan_respirasi_suara_nafas" json:"pemeriksaan_respirasi_suara_nafas"`
	PemeriksaanRespirasiVolumePernafasan             string   `form:"pemeriksaan_respirasi_volume_pernafasan" json:"pemeriksaan_respirasi_volume_pernafasan"`
	PemeriksaanRespirasiJenisPernafasan              string   `form:"pemeriksaan_respirasi_jenis_pernafasan" json:"pemeriksaan_respirasi_jenis_pernafasan"`
	PemeriksaanRespirasiJenisPernafasanKeterangan    string   `form:"pemeriksaan_respirasi_jenis_pernafasan_keterangan" json:"pemeriksaan_respirasi_jenis_pernafasan_keterangan"`
	PemeriksaanRespirasiIramaNafas                   string   `form:"pemeriksaan_respirasi_irama_nafas" json:"pemeriksaan_respirasi_irama_nafas"`
	PemeriksaanRespirasiBatuk                        string   `form:"pemeriksaan_respirasi_batuk" json:"pemeriksaan_respirasi_batuk"`
	PemeriksaanGastrointestinalMulut                 string   `form:"pemeriksaan_gastrointestinal_mulut" json:"pemeriksaan_gastrointestinal_mulut"`
	PemeriksaanGastrointestinalMulutKeterangan       string   `form:"pemeriksaan_gastrointestinal_mulut_keterangan" json:"pemeriksaan_gastrointestinal_mulut_keterangan"`
	PemeriksaanGastrointestinalGigi                  string   `form:"pemeriksaan_gastrointestinal_gigi" json:"pemeriksaan_gastrointestinal_gigi"`
	PemeriksaanGastrointestinalGigiKeterangan        string   `form:"pemeriksaan_gastrointestinal_gigi_keterangan" json:"pemeriksaan_gastrointestinal_gigi_keterangan"`
	PemeriksaanGastrointestinalLidah                 string   `form:"pemeriksaan_gastrointestinal_lidah" json:"pemeriksaan_gastrointestinal_lidah"`
	PemeriksaanGastrointestinalLidahKeterangan       string   `form:"pemeriksaan_gastrointestinal_lidah_keterangan" json:"pemeriksaan_gastrointestinal_lidah_keterangan"`
	PemeriksaanGastrointestinalTenggorokan           string   `form:"pemeriksaan_gastrointestinal_tenggorokan" json:"pemeriksaan_gastrointestinal_tenggorokan"`
	PemeriksaanGastrointestinalTenggorokanKeterangan string   `form:"pemeriksaan_gastrointestinal_tenggorokan_keterangan" json:"pemeriksaan_gastrointestinal_tenggorokan_keterangan"`
	PemeriksaanGastrointestinalAbdomen               string   `form:"pemeriksaan_gastrointestinal_abdomen" json:"pemeriksaan_gastrointestinal_abdomen"`
	PemeriksaanGastrointestinalAbdomenKeterangan     string   `form:"pemeriksaan_gastrointestinal_abdomen_keterangan" json:"pemeriksaan_gastrointestinal_abdomen_keterangan"`
	PemeriksaanGastrointestinalPeistatikUsus         string   `form:"pemeriksaan_gastrointestinal_peistatik_usus" json:"pemeriksaan_gastrointestinal_peistatik_usus"`
	PemeriksaanGastrointestinalAnus                  string   `form:"pemeriksaan_gastrointestinal_anus" json:"pemeriksaan_gastrointestinal_anus"`
	PemeriksaanNeurologiPengelihatan                 string   `form:"pemeriksaan_neurologi_pengelihatan" json:"pemeriksaan_neurologi_pengelihatan"`
	PemeriksaanNeurologiPengelihatanKeterangan       string   `form:"pemeriksaan_neurologi_pengelihatan_keterangan" json:"pemeriksaan_neurologi_pengelihatan_keterangan"`
	PemeriksaanNeurologiAlatBantuPenglihatan         string   `form:"pemeriksaan_neurologi_alat_bantu_penglihatan" json:"pemeriksaan_neurologi_alat_bantu_penglihatan"`
	PemeriksaanNeurologiPendengaran                  string   `form:"pemeriksaan_neurologi_pendengaran" json:"pemeriksaan_neurologi_pendengaran"`
	PemeriksaanNeurologiBicara                       string   `form:"pemeriksaan_neurologi_bicara" json:"pemeriksaan_neurologi_bicara"`
	PemeriksaanNeurologiBicaraKeterangan             string   `form:"pemeriksaan_neurologi_bicara_keterangan" json:"pemeriksaan_neurologi_bicara_keterangan"`
	PemeriksaanNeurologiSensorik                     string   `form:"pemeriksaan_neurologi_sensorik" json:"pemeriksaan_neurologi_sensorik"`
	PemeriksaanNeurologiMotorik                      string   `form:"pemeriksaan_neurologi_motorik" json:"pemeriksaan_neurologi_motorik"`
	PemeriksaanNeurologiKekuatanOtot                 string   `form:"pemeriksaan_neurologi_kekuatan_otot" json:"pemeriksaan_neurologi_kekuatan_otot"`
	PemeriksaanIntegumentWarnakulit                  string   `form:"pemeriksaan_integument_warnakulit" json:"pemeriksaan_integument_warnakulit"`
	PemeriksaanIntegumentTurgor                      string   `form:"pemeriksaan_integument_turgor" json:"pemeriksaan_integument_turgor"`
	PemeriksaanIntegumentKulit                       string   `form:"pemeriksaan_integument_kulit" json:"pemeriksaan_integument_kulit"`
	PemeriksaanIntegumentDekubitas                   string   `form:"pemeriksaan_integument_dekubitas" json:"pemeriksaan_integument_dekubitas"`
	PemeriksaanMuskuloskletalPergerakanSendi         string   `form:"pemeriksaan_muskuloskletal_pergerakan_sendi" json:"pemeriksaan_muskuloskletal_pergerakan_sendi"`
	PemeriksaanMuskuloskletalKekauatanOtot           string   `form:"pemeriksaan_muskuloskletal_kekauatan_otot" json:"pemeriksaan_muskuloskletal_kekauatan_otot"`
	PemeriksaanMuskuloskletalNyeriSendi              string   `form:"pemeriksaan_muskuloskletal_nyeri_sendi" json:"pemeriksaan_muskuloskletal_nyeri_sendi"`
	PemeriksaanMuskuloskletalNyeriSendiKeterangan    string   `form:"pemeriksaan_muskuloskletal_nyeri_sendi_keterangan" json:"pemeriksaan_muskuloskletal_nyeri_sendi_keterangan"`
	PemeriksaanMuskuloskletalOedema                  string   `form:"pemeriksaan_muskuloskletal_oedema" json:"pemeriksaan_muskuloskletal_oedema"`
	PemeriksaanMuskuloskletalOedemaKeterangan        string   `form:"pemeriksaan_muskuloskletal_oedema_keterangan" json:"pemeriksaan_muskuloskletal_oedema_keterangan"`
	PemeriksaanMuskuloskletalFraktur                 string   `form:"pemeriksaan_muskuloskletal_fraktur" json:"pemeriksaan_muskuloskletal_fraktur"`
	PemeriksaanMuskuloskletalFrakturKeterangan       string   `form:"pemeriksaan_muskuloskletal_fraktur_keterangan" json:"pemeriksaan_muskuloskletal_fraktur_keterangan"`
	PemeriksaanEliminasiBabFrekuensiJumlah           string   `form:"pemeriksaan_eliminasi_bab_frekuensi_jumlah" json:"pemeriksaan_eliminasi_bab_frekuensi_jumlah"`
	PemeriksaanEliminasiBabFrekuensiDurasi           string   `form:"pemeriksaan_eliminasi_bab_frekuensi_durasi" json:"pemeriksaan_eliminasi_bab_frekuensi_durasi"`
	PemeriksaanEliminasiBabKonsistensi               string   `form:"pemeriksaan_eliminasi_bab_konsistensi" json:"pemeriksaan_eliminasi_bab_konsistensi"`
	PemeriksaanEliminasiBabWarna                     string   `form:"pemeriksaan_eliminasi_bab_warna" json:"pemeriksaan_eliminasi_bab_warna"`
	PemeriksaanEliminasiBakFrekuensiJumlah           string   `form:"pemeriksaan_eliminasi_bak_frekuensi_jumlah" json:"pemeriksaan_eliminasi_bak_frekuensi_jumlah"`
	PemeriksaanEliminasiBakFrekuensiDurasi           string   `form:"pemeriksaan_eliminasi_bak_frekuensi_durasi" json:"pemeriksaan_eliminasi_bak_frekuensi_durasi"`
	PemeriksaanEliminasiBakWarna                     string   `form:"pemeriksaan_eliminasi_bak_warna" json:"pemeriksaan_eliminasi_bak_warna"`
	PemeriksaanEliminasiBakLainlain                  string   `form:"pemeriksaan_eliminasi_bak_lainlain" json:"pemeriksaan_eliminasi_bak_lainlain"`
	PolaAktifitasMakanminum                          string   `form:"pola_aktifitas_makanminum" json:"pola_aktifitas_makanminum"`
	PolaAktifitasMandi                               string   `form:"pola_aktifitas_mandi" json:"pola_aktifitas_mandi"`
	PolaAktifitasEliminasi                           string   `form:"pola_aktifitas_eliminasi" json:"pola_aktifitas_eliminasi"`
	PolaAktifitasBerpakaian                          string   `form:"pola_aktifitas_berpakaian" json:"pola_aktifitas_berpakaian"`
	PolaAktifitasBerpindah                           string   `form:"pola_aktifitas_berpindah" json:"pola_aktifitas_berpindah"`
	PolaNutrisiFrekuesiMakan                         string   `form:"pola_nutrisi_frekuesi_makan" json:"pola_nutrisi_frekuesi_makan"`
	PolaNutrisiJenisMakanan                          string   `form:"pola_nutrisi_jenis_makanan" json:"pola_nutrisi_jenis_makanan"`
	PolaNutrisiPorsiMakan                            string   `form:"pola_nutrisi_porsi_makan" json:"pola_nutrisi_porsi_makan"`
	PolaTidurLamaTidur                               string   `form:"pola_tidur_lama_tidur" json:"pola_tidur_lama_tidur"`
	PolaTidurGangguan                                string   `form:"pola_tidur_gangguan" json:"pola_tidur_gangguan"`
	PengkajianFungsiKemampuanSehari                  string   `form:"pengkajian_fungsi_kemampuan_sehari" json:"pengkajian_fungsi_kemampuan_sehari"`
	PengkajianFungsiAktifitas                        string   `form:"pengkajian_fungsi_aktifitas" json:"pengkajian_fungsi_aktifitas"`
	PengkajianFungsiBerjalan                         string   `form:"pengkajian_fungsi_berjalan" json:"pengkajian_fungsi_berjalan"`
	PengkajianFungsiBerjalanKeterangan               string   `form:"pengkajian_fungsi_berjalan_keterangan" json:"pengkajian_fungsi_berjalan_keterangan"`
	PengkajianFungsiAmbulasi                         string   `form:"pengkajian_fungsi_ambulasi" json:"pengkajian_fungsi_ambulasi"`
	PengkajianFungsiEkstrimitasAtas                  string   `form:"pengkajian_fungsi_ekstrimitas_atas" json:"pengkajian_fungsi_ekstrimitas_atas"`
	PengkajianFungsiEkstrimitasAtasKeterangan        string   `form:"pengkajian_fungsi_ekstrimitas_atas_keterangan" json:"pengkajian_fungsi_ekstrimitas_atas_keterangan"`
	PengkajianFungsiEkstrimitasBawah                 string   `form:"pengkajian_fungsi_ekstrimitas_bawah" json:"pengkajian_fungsi_ekstrimitas_bawah"`
	PengkajianFungsiEkstrimitasBawahKeterangan       string   `form:"pengkajian_fungsi_ekstrimitas_bawah_keterangan" json:"pengkajian_fungsi_ekstrimitas_bawah_keterangan"`
	PengkajianFungsiMenggenggam                      string   `form:"pengkajian_fungsi_menggenggam" json:"pengkajian_fungsi_menggenggam"`
	PengkajianFungsiMenggenggamKeterangan            string   `form:"pengkajian_fungsi_menggenggam_keterangan" json:"pengkajian_fungsi_menggenggam_keterangan"`
	PengkajianFungsiKoordinasi                       string   `form:"pengkajian_fungsi_koordinasi" json:"pengkajian_fungsi_koordinasi"`
	PengkajianFungsiKoordinasiKeterangan             string   `form:"pengkajian_fungsi_koordinasi_keterangan" json:"pengkajian_fungsi_koordinasi_keterangan"`
	PengkajianFungsiKesimpulan                       string   `form:"pengkajian_fungsi_kesimpulan" json:"pengkajian_fungsi_kesimpulan"`
	RiwayatPsikoKondisiPsiko                         string   `form:"riwayat_psiko_kondisi_psiko" json:"riwayat_psiko_kondisi_psiko"`
	RiwayatPsikoGangguanJiwa                         string   `form:"riwayat_psiko_gangguan_jiwa" json:"riwayat_psiko_gangguan_jiwa"`
	RiwayatPsikoPerilaku                             string   `form:"riwayat_psiko_perilaku" json:"riwayat_psiko_perilaku"`
	RiwayatPsikoPerilakuKeterangan                   string   `form:"riwayat_psiko_perilaku_keterangan" json:"riwayat_psiko_perilaku_keterangan"`
	RiwayatPsikoHubunganKeluarga                     string   `form:"riwayat_psiko_hubungan_keluarga" json:"riwayat_psiko_hubungan_keluarga"`
	RiwayatPsikoTinggal                              string   `form:"riwayat_psiko_tinggal" json:"riwayat_psiko_tinggal"`
	RiwayatPsikoTinggalKeterangan                    string   `form:"riwayat_psiko_tinggal_keterangan" json:"riwayat_psiko_tinggal_keterangan"`
	RiwayatPsikoNilaiKepercayaan                     string   `form:"riwayat_psiko_nilai_kepercayaan" json:"riwayat_psiko_nilai_kepercayaan"`
	RiwayatPsikoNilaiKepercayaanKeterangan           string   `form:"riwayat_psiko_nilai_kepercayaan_keterangan" json:"riwayat_psiko_nilai_kepercayaan_keterangan"`
	RiwayatPsikoPendidikanPj                         string   `form:"riwayat_psiko_pendidikan_pj" json:"riwayat_psiko_pendidikan_pj"`
	RiwayatPsikoEdukasiDiberikan                     string   `form:"riwayat_psiko_edukasi_diberikan" json:"riwayat_psiko_edukasi_diberikan"`
	RiwayatPsikoEdukasiDiberikanKeterangan           string   `form:"riwayat_psiko_edukasi_diberikan_keterangan" json:"riwayat_psiko_edukasi_diberikan_keterangan"`
	PenilaianNyeri                                   string   `form:"penilaian_nyeri" json:"penilaian_nyeri"`
	PenilaianNyeriPenyebab                           string   `form:"penilaian_nyeri_penyebab" json:"penilaian_nyeri_penyebab"`
	PenilaianNyeriKetPenyebab                        string   `form:"penilaian_nyeri_ket_penyebab" json:"penilaian_nyeri_ket_penyebab"`
	PenilaianNyeriKualitas                           string   `form:"penilaian_nyeri_kualitas" json:"penilaian_nyeri_kualitas"`
	PenilaianNyeriKetKualitas                        string   `form:"penilaian_nyeri_ket_kualitas" json:"penilaian_nyeri_ket_kualitas"`
	PenilaianNyeriLokasi                             string   `form:"penilaian_nyeri_lokasi" json:"penilaian_nyeri_lokasi"`
	PenilaianNyeriMenyebar                           string   `form:"penilaian_nyeri_menyebar" json:"penilaian_nyeri_menyebar"`
	PenilaianNyeriSkala                              string   `form:"penilaian_nyeri_skala" json:"penilaian_nyeri_skala"`
	PenilaianNyeriWaktu                              string   `form:"penilaian_nyeri_waktu" json:"penilaian_nyeri_waktu"`
	PenilaianNyeriHilang                             string   `form:"penilaian_nyeri_hilang" json:"penilaian_nyeri_hilang"`
	PenilaianNyeriKetHilang                          string   `form:"penilaian_nyeri_ket_hilang" json:"penilaian_nyeri_ket_hilang"`
	PenilaianNyeriDiberitahukanDokter                string   `form:"penilaian_nyeri_diberitahukan_dokter" json:"penilaian_nyeri_diberitahukan_dokter"`
	PenilaianNyeriJamDiberitahukanDokter             string   `form:"penilaian_nyeri_jam_diberitahukan_dokter" json:"penilaian_nyeri_jam_diberitahukan_dokter"`
	PenilaianJatuhmorseSkala1                        string   `form:"penilaian_jatuhmorse_skala1" json:"penilaian_jatuhmorse_skala1"`
	PenilaianJatuhmorseNilai1                        *int     `form:"penilaian_jatuhmorse_nilai1" json:"penilaian_jatuhmorse_nilai1"`
	PenilaianJatuhmorseSkala2                        string   `form:"penilaian_jatuhmorse_skala2" json:"penilaian_jatuhmorse_skala2"`
	PenilaianJatuhmorseNilai2                        *int     `form:"penilaian_jatuhmorse_nilai2" json:"penilaian_jatuhmorse_nilai2"`
	PenilaianJatuhmorseSkala3                        string   `form:"penilaian_jatuhmorse_skala3" json:"penilaian_jatuhmorse_skala3"`
	PenilaianJatuhmorseNilai3                        *int     `form:"penilaian_jatuhmorse_nilai3" json:"penilaian_jatuhmorse_nilai3"`
	PenilaianJatuhmorseSkala4                        string   `form:"penilaian_jatuhmorse_skala4" json:"penilaian_jatuhmorse_skala4"`
	PenilaianJatuhmorseNilai4                        *int     `form:"penilaian_jatuhmorse_nilai4" json:"penilaian_jatuhmorse_nilai4"`
	PenilaianJatuhmorseSkala5                        string   `form:"penilaian_jatuhmorse_skala5" json:"penilaian_jatuhmorse_skala5"`
	PenilaianJatuhmorseNilai5                        *int     `form:"penilaian_jatuhmorse_nilai5" json:"penilaian_jatuhmorse_nilai5"`
	PenilaianJatuhmorseSkala6                        string   `form:"penilaian_jatuhmorse_skala6" json:"penilaian_jatuhmorse_skala6"`
	PenilaianJatuhmorseNilai6                        *int     `form:"penilaian_jatuhmorse_nilai6" json:"penilaian_jatuhmorse_nilai6"`
	PenilaianJatuhmorseTotalnilai                    *int     `form:"penilaian_jatuhmorse_totalnilai" json:"penilaian_jatuhmorse_totalnilai"`
	PenilaianJatuhsydneySkala1                       string   `form:"penilaian_jatuhsydney_skala1" json:"penilaian_jatuhsydney_skala1"`
	PenilaianJatuhsydneyNilai1                       *int     `form:"penilaian_jatuhsydney_nilai1" json:"penilaian_jatuhsydney_nilai1"`
	PenilaianJatuhsydneySkala2                       string   `form:"penilaian_jatuhsydney_skala2" json:"penilaian_jatuhsydney_skala2"`
	PenilaianJatuhsydneyNilai2                       *int     `form:"penilaian_jatuhsydney_nilai2" json:"penilaian_jatuhsydney_nilai2"`
	PenilaianJatuhsydneySkala3                       string   `form:"penilaian_jatuhsydney_skala3" json:"penilaian_jatuhsydney_skala3"`
	PenilaianJatuhsydneyNilai3                       *int     `form:"penilaian_jatuhsydney_nilai3" json:"penilaian_jatuhsydney_nilai3"`
	PenilaianJatuhsydneySkala4                       string   `form:"penilaian_jatuhsydney_skala4" json:"penilaian_jatuhsydney_skala4"`
	PenilaianJatuhsydneyNilai4                       *int     `form:"penilaian_jatuhsydney_nilai4" json:"penilaian_jatuhsydney_nilai4"`
	PenilaianJatuhsydneySkala5                       string   `form:"penilaian_jatuhsydney_skala5" json:"penilaian_jatuhsydney_skala5"`
	PenilaianJatuhsydneyNilai5                       *int     `form:"penilaian_jatuhsydney_nilai5" json:"penilaian_jatuhsydney_nilai5"`
	PenilaianJatuhsydneySkala6                       string   `form:"penilaian_jatuhsydney_skala6" json:"penilaian_jatuhsydney_skala6"`
	PenilaianJatuhsydneyNilai6                       *int     `form:"penilaian_jatuhsydney_nilai6" json:"penilaian_jatuhsydney_nilai6"`
	PenilaianJatuhsydneySkala7                       string   `form:"penilaian_jatuhsydney_skala7" json:"penilaian_jatuhsydney_skala7"`
	PenilaianJatuhsydneyNilai7                       *int     `form:"penilaian_jatuhsydney_nilai7" json:"penilaian_jatuhsydney_nilai7"`
	PenilaianJatuhsydneySkala8                       string   `form:"penilaian_jatuhsydney_skala8" json:"penilaian_jatuhsydney_skala8"`
	PenilaianJatuhsydneyNilai8                       *int     `form:"penilaian_jatuhsydney_nilai8" json:"penilaian_jatuhsydney_nilai8"`
	PenilaianJatuhsydneySkala9                       string   `form:"penilaian_jatuhsydney_skala9" json:"penilaian_jatuhsydney_skala9"`
	PenilaianJatuhsydneyNilai9                       *int     `form:"penilaian_jatuhsydney_nilai9" json:"penilaian_jatuhsydney_nilai9"`
	PenilaianJatuhsydneySkala10                      string   `form:"penilaian_jatuhsydney_skala10" json:"penilaian_jatuhsydney_skala10"`
	PenilaianJatuhsydneyNilai10                      *int     `form:"penilaian_jatuhsydney_nilai10" json:"penilaian_jatuhsydney_nilai10"`
	PenilaianJatuhsydneySkala11                      string   `form:"penilaian_jatuhsydney_skala11" json:"penilaian_jatuhsydney_skala11"`
	PenilaianJatuhsydneyNilai11                      *int     `form:"penilaian_jatuhsydney_nilai11" json:"penilaian_jatuhsydney_nilai11"`
	PenilaianJatuhsydneyTotalnilai                   *int     `form:"penilaian_jatuhsydney_totalnilai" json:"penilaian_jatuhsydney_totalnilai"`
	SkriningGizi1                                    string   `form:"skrining_gizi1" json:"skrining_gizi1"`
	NilaiGizi1                                       *int     `form:"nilai_gizi1" json:"nilai_gizi1"`
	SkriningGizi2                                    string   `form:"skrining_gizi2" json:"skrining_gizi2"`
	NilaiGizi2                                       *int     `form:"nilai_gizi2" json:"nilai_gizi2"`
	NilaiTotalGizi                                   *float64 `form:"nilai_total_gizi" json:"nilai_total_gizi"`
	SkriningGiziDiagnosaKhusus                       string   `form:"skrining_gizi_diagnosa_khusus" json:"skrining_gizi_diagnosa_khusus"`
	SkriningGiziKetDiagnosaKhusus                    string   `form:"skrining_gizi_ket_diagnosa_khusus" json:"skrining_gizi_ket_diagnosa_khusus"`
	SkriningGiziDiketahuiDietisen                    string   `form:"skrining_gizi_diketahui_dietisen" json:"skrining_gizi_diketahui_dietisen"`
	SkriningGiziJamDiketahuiDietisen                 string   `form:"skrining_gizi_jam_diketahui_dietisen" json:"skrining_gizi_jam_diketahui_dietisen"`
	Rencana                                          string   `form:"rencana" json:"rencana"`
	Nip1                                             string   `form:"nip1" json:"nip1"`
	Nip2                                             string   `form:"nip2" json:"nip2"`
	KdDokter                                         string   `form:"kd_dokter" json:"kd_dokter"`
}

// PenilaianAwalKeperawatanRanapDetail daftar kode detail (diganti seluruhnya saat simpan/ubah).
type PenilaianAwalKeperawatanRanapDetail struct {
	KodeMasalah []string `form:"kode_masalah" json:"kode_masalah"`
	KodeRencana []string `form:"kode_rencana" json:"kode_rencana"`
}

func penilaianAwalKeperawatanRanapRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                "required|date",
		"informasi":                              "required|in:Autoanamnesis,Alloanamnesis",
		"ket_informasi":                          "string|max_len:30",
		"tiba_diruang_rawat":                     "required|in:Jalan Tanpa Bantuan,Kursi Roda,Brankar",
		"kasus_trauma":                           "in:Trauma,Non Trauma",
		"cara_masuk":                             "required|in:Poli,IGD,Lain-lain",
		"rps":                                    "string|max_len:300",
		"rpd":                                    "string|max_len:100",
		"rpk":                                    "string|max_len:100",
		"rpo":                                    "string|max_len:100",
		"riwayat_pembedahan":                     "string|max_len:40",
		"riwayat_dirawat_dirs":                   "string|max_len:40",
		"alat_bantu_dipakai":                     "required|in:Tidak Ada,Kacamata,Prothesa,Alat Bantu Dengar,Lain-lain",
		"riwayat_kehamilan":                      "required|in:Tidak,Ya",
		"riwayat_kehamilan_perkiraan":            "string|max_len:30",
		"riwayat_tranfusi":                       "string|max_len:40",
		"riwayat_alergi":                         "string|max_len:40",
		"riwayat_merokok":                        "required|in:Tidak,Ya",
		"riwayat_merokok_jumlah":                 "string|max_len:5",
		"riwayat_alkohol":                        "required|in:Tidak,Ya",
		"riwayat_alkohol_jumlah":                 "string|max_len:5",
		"riwayat_narkoba":                        "required|in:Tidak,Ya",
		"riwayat_olahraga":                       "required|in:Tidak,Ya",
		"pemeriksaan_mental":                     "string|max_len:40",
		"pemeriksaan_keadaan_umum":               "required|in:Baik,Sedang,Buruk",
		"pemeriksaan_gcs":                        "string|max_len:10",
		"pemeriksaan_td":                         "string|max_len:8",
		"pemeriksaan_nadi":                       "string|max_len:5",
		"pemeriksaan_rr":                         "string|max_len:5",
		"pemeriksaan_suhu":                       "string|max_len:5",
		"pemeriksaan_spo2":                       "string|max_len:5",
		"pemeriksaan_bb":                         "string|max_len:5",
		"pemeriksaan_tb":                         "string|max_len:5",
		"pemeriksaan_susunan_kepala":             "required|in:TAK,Hydrocephalus,Hematoma,Lain-lain",
		"pemeriksaan_susunan_kepala_keterangan":  "string|max_len:50",
		"pemeriksaan_susunan_wajah":              "required|in:TAK,Asimetris,Kelainan Kongenital",
		"pemeriksaan_susunan_wajah_keterangan":   "string|max_len:50",
		"pemeriksaan_susunan_leher":              "required|in:TAK,Kaku Kuduk,Pembesaran Thyroid,Pembesaran KGB",
		"pemeriksaan_susunan_kejang":             "required|in:TAK,Kuat,Ada",
		"pemeriksaan_susunan_kejang_keterangan":  "string|max_len:50",
		"pemeriksaan_susunan_sensorik":           "required|in:TAK,Sakit Nyeri,Rasa kebas",
		"pemeriksaan_kardiovaskuler_denyut_nadi": "required|in:Teratur,Tidak Teratur",
		"pemeriksaan_kardiovaskuler_sirkulasi":   "required|in:Akral Hangat,Akral Dingin,Edema",
		"pemeriksaan_kardiovaskuler_sirkulasi_keterangan":     "string|max_len:50",
		"pemeriksaan_kardiovaskuler_pulsasi":                  "required|in:Kuat,Lemah,Lain-lain",
		"pemeriksaan_respirasi_pola_nafas":                    "required|in:Normal,Bradipnea,Tachipnea",
		"pemeriksaan_respirasi_retraksi":                      "required|in:Tidak Ada,Ringan,Berat",
		"pemeriksaan_respirasi_suara_nafas":                   "required|in:Vesikuler,Wheezing,Rhonki",
		"pemeriksaan_respirasi_volume_pernafasan":             "required|in:Normal,Hiperventilasi,Hipoventilasi",
		"pemeriksaan_respirasi_jenis_pernafasan":              "required|in:Pernafasan Dada,Alat Bantu Pernafasaan",
		"pemeriksaan_respirasi_jenis_pernafasan_keterangan":   "string|max_len:50",
		"pemeriksaan_respirasi_irama_nafas":                   "required|in:Teratur,Tidak Teratur",
		"pemeriksaan_respirasi_batuk":                         "required|in:Tidak,Ya : Produktif,Ya : Non Produktif",
		"pemeriksaan_gastrointestinal_mulut":                  "required|in:TAK,Stomatitis,Mukosa Kering,Bibir Pucat,Lain-lain",
		"pemeriksaan_gastrointestinal_mulut_keterangan":       "string|max_len:50",
		"pemeriksaan_gastrointestinal_gigi":                   "required|in:TAK,Karies,Goyang,Lain-lain",
		"pemeriksaan_gastrointestinal_gigi_keterangan":        "string|max_len:50",
		"pemeriksaan_gastrointestinal_lidah":                  "required|in:TAK,Kotor,Gerak Asimetris,Lain-lain",
		"pemeriksaan_gastrointestinal_lidah_keterangan":       "string|max_len:50",
		"pemeriksaan_gastrointestinal_tenggorokan":            "required|in:TAK,Gangguan Menelan,Sakit Menelan,Lain-lain",
		"pemeriksaan_gastrointestinal_tenggorokan_keterangan": "string|max_len:50",
		"pemeriksaan_gastrointestinal_abdomen":                "required|in:Supel,Asictes, Tegang,Nyeri Tekan/Lepas,Lain-lain",
		"pemeriksaan_gastrointestinal_abdomen_keterangan":     "string|max_len:50",
		"pemeriksaan_gastrointestinal_peistatik_usus":         "required|in:TAK,Tidak Ada Bising Usus,Hiperistaltik",
		"pemeriksaan_gastrointestinal_anus":                   "required|in:TAK,Atresia Ani",
		"pemeriksaan_neurologi_pengelihatan":                  "required|in:TAK,Ada Kelainan",
		"pemeriksaan_neurologi_pengelihatan_keterangan":       "string|max_len:50",
		"pemeriksaan_neurologi_alat_bantu_penglihatan":        "required|in:Tidak,Kacamata,Lensa Kontak",
		"pemeriksaan_neurologi_pendengaran":                   "required|in:TAK,Berdengung,Nyeri,Tuli,Keluar Cairan,Lain-lain",
		"pemeriksaan_neurologi_bicara":                        "required|in:Jelas,Tidak Jelas",
		"pemeriksaan_neurologi_bicara_keterangan":             "string|max_len:50",
		"pemeriksaan_neurologi_sensorik":                      "required|in:TAK,Sakit Nyeri,Rasa Kebas,Lain-lain",
		"pemeriksaan_neurologi_motorik":                       "required|in:TAK,Hemiparese,Tetraparese,Tremor,Lain-lain",
		"pemeriksaan_neurologi_kekuatan_otot":                 "required|in:Kuat,Lemah",
		"pemeriksaan_integument_warnakulit":                   "required|in:Pucat,Sianosis,Normal,Lain-lain",
		"pemeriksaan_integument_turgor":                       "required|in:Baik,Sedang,Buruk",
		"pemeriksaan_integument_kulit":                        "required|in:Normal,Rash/Kemerahan,Luka,Memar,Ptekie,Bula",
		"pemeriksaan_integument_dekubitas":                    "required|string",
		"pemeriksaan_muskuloskletal_pergerakan_sendi":         "required|in:Bebas,Terbatas",
		"pemeriksaan_muskuloskletal_kekauatan_otot":           "required|in:Baik,Lemah,Tremor",
		"pemeriksaan_muskuloskletal_nyeri_sendi":              "required|in:Tidak Ada,Ada",
		"pemeriksaan_muskuloskletal_nyeri_sendi_keterangan":   "string|max_len:50",
		"pemeriksaan_muskuloskletal_oedema":                   "required|in:Tidak Ada,Ada",
		"pemeriksaan_muskuloskletal_oedema_keterangan":        "string|max_len:50",
		"pemeriksaan_muskuloskletal_fraktur":                  "required|in:Tidak Ada,Ada",
		"pemeriksaan_muskuloskletal_fraktur_keterangan":       "string|max_len:50",
		"pemeriksaan_eliminasi_bab_frekuensi_jumlah":          "string|max_len:5",
		"pemeriksaan_eliminasi_bab_frekuensi_durasi":          "string|max_len:10",
		"pemeriksaan_eliminasi_bab_konsistensi":               "string|max_len:30",
		"pemeriksaan_eliminasi_bab_warna":                     "string|max_len:30",
		"pemeriksaan_eliminasi_bak_frekuensi_jumlah":          "string|max_len:5",
		"pemeriksaan_eliminasi_bak_frekuensi_durasi":          "string|max_len:10",
		"pemeriksaan_eliminasi_bak_warna":                     "string|max_len:30",
		"pemeriksaan_eliminasi_bak_lainlain":                  "string|max_len:30",
		"pola_aktifitas_makanminum":                           "required|in:Mandiri,Bantuan Orang Lain",
		"pola_aktifitas_mandi":                                "required|in:Mandiri,Bantuan Orang Lain",
		"pola_aktifitas_eliminasi":                            "required|in:Mandiri,Bantuan Orang Lain",
		"pola_aktifitas_berpakaian":                           "required|in:Mandiri,Bantuan Orang Lain",
		"pola_aktifitas_berpindah":                            "required|in:Mandiri,Bantuan Orang Lain",
		"pola_nutrisi_frekuesi_makan":                         "string|max_len:3",
		"pola_nutrisi_jenis_makanan":                          "string|max_len:20",
		"pola_nutrisi_porsi_makan":                            "string|max_len:3",
		"pola_tidur_lama_tidur":                               "string|max_len:3",
		"pola_tidur_gangguan":                                 "required|in:Tidak Ada Gangguan,Insomnia",
		"pengkajian_fungsi_kemampuan_sehari":                  "required|in:Mandiri,Bantuan Minimal,Bantuan Sebagian,Ketergantungan Total",
		"pengkajian_fungsi_aktifitas":                         "required|in:Tirah Baring,Duduk,Berjalan",
		"pengkajian_fungsi_berjalan":                          "required|in:TAK,Penurunan Kekuatan/ROM,Paralisis,Sering Jatuh,Deformitas,Hilang Keseimbangan,Riwayat Patah Tulang,Lain-lain",
		"pengkajian_fungsi_berjalan_keterangan":               "string|max_len:40",
		"pengkajian_fungsi_ambulasi":                          "required|in:Walker,Tongkat,Kursi Roda,Tidak Menggunakan",
		"pengkajian_fungsi_ekstrimitas_atas":                  "required|in:TAK,Lemah,Oedema,Tidak Simetris,Lain-lain",
		"pengkajian_fungsi_ekstrimitas_atas_keterangan":       "string|max_len:40",
		"pengkajian_fungsi_ekstrimitas_bawah":                 "required|in:TAK,Varises,Oedema,Tidak Simetris,Lain-lain",
		"pengkajian_fungsi_ekstrimitas_bawah_keterangan":      "string|max_len:40",
		"pengkajian_fungsi_menggenggam":                       "required|in:Tidak Ada Kesulitan,Terakhir,Lain-lain",
		"pengkajian_fungsi_menggenggam_keterangan":            "string|max_len:40",
		"pengkajian_fungsi_koordinasi":                        "required|in:Tidak Ada Kesulitan,Ada Masalah",
		"pengkajian_fungsi_koordinasi_keterangan":             "string|max_len:40",
		"pengkajian_fungsi_kesimpulan":                        "required|in:Ya (Co DPJP),Tidak (Tidak Perlu Co DPJP)",
		"riwayat_psiko_kondisi_psiko":                         "required|in:Tidak Ada Masalah,Marah,Takut,Depresi,Cepat Lelah,Cemas,Gelisah,Sulit Tidur,Lain-lain",
		"riwayat_psiko_gangguan_jiwa":                         "required|in:Ya,Tidak",
		"riwayat_psiko_perilaku":                              "required|in:Tidak Ada Masalah,Perilaku Kekerasan,Gangguan Efek,Gangguan Memori,Halusinasi,Kecenderungan Percobaan Bunuh Diri,Lain-lain",
		"riwayat_psiko_perilaku_keterangan":                   "string|max_len:40",
		"riwayat_psiko_hubungan_keluarga":                     "required|in:Harmonis,Kurang Harmonis,Tidak Harmonis,Konflik Besar",
		"riwayat_psiko_tinggal":                               "required|in:Sendiri,Orang Tua,Suami/Istri,Keluarga,Lain-lain",
		"riwayat_psiko_tinggal_keterangan":                    "string|max_len:40",
		"riwayat_psiko_nilai_kepercayaan":                     "required|in:Tidak Ada,Ada",
		"riwayat_psiko_nilai_kepercayaan_keterangan":          "string|max_len:40",
		"riwayat_psiko_pendidikan_pj":                         "required|in:-,TS,TK,SD,SMP,SMA,SLTA/SEDERAJAT,D1,D2,D3,D4,S1,S2,S3",
		"riwayat_psiko_edukasi_diberikan":                     "required|in:Pasien,Keluarga",
		"riwayat_psiko_edukasi_diberikan_keterangan":          "string|max_len:40",
		"penilaian_nyeri":                                     "required|in:Tidak Ada Nyeri,Nyeri Akut,Nyeri Kronis",
		"penilaian_nyeri_penyebab":                            "required|in:Proses Penyakit,Benturan,Lain-lain",
		"penilaian_nyeri_ket_penyebab":                        "string|max_len:50",
		"penilaian_nyeri_kualitas":                            "required|in:Seperti Tertusuk,Berdenyut,Teriris,Tertindih,Tertiban,Lain-lain",
		"penilaian_nyeri_ket_kualitas":                        "string|max_len:50",
		"penilaian_nyeri_lokasi":                              "string|max_len:50",
		"penilaian_nyeri_menyebar":                            "required|in:Tidak,Ya",
		"penilaian_nyeri_skala":                               "required|in:0,1,2,3,4,5,6,7,8,9,10",
		"penilaian_nyeri_waktu":                               "string|max_len:5",
		"penilaian_nyeri_hilang":                              "required|in:Istirahat,Medengar Musik,Minum Obat",
		"penilaian_nyeri_ket_hilang":                          "string|max_len:50",
		"penilaian_nyeri_diberitahukan_dokter":                "required|in:Tidak,Ya",
		"penilaian_nyeri_jam_diberitahukan_dokter":            "string|max_len:10",
		"penilaian_jatuhmorse_skala1":                         "in:Tidak,Ya",
		"penilaian_jatuhmorse_nilai1":                         "int",
		"penilaian_jatuhmorse_skala2":                         "in:Tidak,Ya",
		"penilaian_jatuhmorse_nilai2":                         "int",
		"penilaian_jatuhmorse_skala3":                         "in:Tidak Ada/Kursi Roda/Perawat/Tirah Baring,Tongkat/Alat Penopang,Berpegangan Pada Perabot",
		"penilaian_jatuhmorse_nilai3":                         "int",
		"penilaian_jatuhmorse_skala4":                         "in:Tidak,Ya",
		"penilaian_jatuhmorse_nilai4":                         "int",
		"penilaian_jatuhmorse_skala5":                         "in:Normal/Tirah Baring/Imobilisasi,Lemah,Terganggu",
		"penilaian_jatuhmorse_nilai5":                         "int",
		"penilaian_jatuhmorse_skala6":                         "in:Sadar Akan Kemampuan Diri Sendiri,Sering Lupa Akan Keterbatasan Yang Dimiliki",
		"penilaian_jatuhmorse_nilai6":                         "int",
		"penilaian_jatuhmorse_totalnilai":                     "int",
		"penilaian_jatuhsydney_skala1":                        "in:Tidak,Ya",
		"penilaian_jatuhsydney_nilai1":                        "int",
		"penilaian_jatuhsydney_skala2":                        "in:Tidak,Ya",
		"penilaian_jatuhsydney_nilai2":                        "int",
		"penilaian_jatuhsydney_skala3":                        "in:Tidak,Ya",
		"penilaian_jatuhsydney_nilai3":                        "int",
		"penilaian_jatuhsydney_skala4":                        "in:Tidak,Ya",
		"penilaian_jatuhsydney_nilai4":                        "int",
		"penilaian_jatuhsydney_skala5":                        "in:Tidak,Ya",
		"penilaian_jatuhsydney_nilai5":                        "int",
		"penilaian_jatuhsydney_skala6":                        "in:Tidak,Ya",
		"penilaian_jatuhsydney_nilai6":                        "int",
		"penilaian_jatuhsydney_skala7":                        "in:Tidak,Ya",
		"penilaian_jatuhsydney_nilai7":                        "int",
		"penilaian_jatuhsydney_skala8":                        "in:Tidak,Ya",
		"penilaian_jatuhsydney_nilai8":                        "int",
		"penilaian_jatuhsydney_skala9":                        "in:Tidak,Ya",
		"penilaian_jatuhsydney_nilai9":                        "int",
		"penilaian_jatuhsydney_skala10":                       "in:Tidak,Ya",
		"penilaian_jatuhsydney_nilai10":                       "int",
		"penilaian_jatuhsydney_skala11":                       "in:Tidak,Ya",
		"penilaian_jatuhsydney_nilai11":                       "int",
		"penilaian_jatuhsydney_totalnilai":                    "int",
		"skrining_gizi1":                                      "in:Tidak ada penurunan berat badan,Tidak yakin/ tidak tahu/ terasa baju lebih longgar,Ya 1-5 kg,Ya 6-10 kg,Ya 11-15 kg,Ya > 15 kg",
		"nilai_gizi1":                                         "int",
		"skrining_gizi2":                                      "in:Tidak,Ya",
		"nilai_gizi2":                                         "int",
		"nilai_total_gizi":                                    "numeric",
		"skrining_gizi_diagnosa_khusus":                       "in:Tidak,Ya",
		"skrining_gizi_ket_diagnosa_khusus":                   "string|max_len:50",
		"skrining_gizi_diketahui_dietisen":                    "in:Tidak,Ya",
		"skrining_gizi_jam_diketahui_dietisen":                "string|max_len:10",
		"rencana":                                             "string|max_len:200",
		"nip1":                                                "required|string|max_len:20",
		"nip2":                                                "required|string|max_len:20",
		"kd_dokter":                                           "required|string|max_len:20",
	}
	rules["kode_masalah"] = "slice"
	rules["kode_masalah.*"] = "string|max_len:20"
	rules["kode_rencana"] = "slice"
	rules["kode_rencana.*"] = "string|max_len:20"
	return rules
}

// PenilaianAwalKeperawatanRanapStore simpan penilaian awal keperawatan ranap; kolom waktu kunci kosong = sekarang.
type PenilaianAwalKeperawatanRanapStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianAwalKeperawatanRanapData
	PenilaianAwalKeperawatanRanapDetail
}

func (r *PenilaianAwalKeperawatanRanapStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanRanapStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianAwalKeperawatanRanapRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianAwalKeperawatanRanapStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianAwalKeperawatanRanapStore) Payload() PenilaianAwalKeperawatanRanapData {
	return r.PenilaianAwalKeperawatanRanapData
}

func (r *PenilaianAwalKeperawatanRanapStore) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}

// PenilaianAwalKeperawatanRanapUpdate ubah penilaian awal keperawatan ranap (PUT); kunci lewat query string.
type PenilaianAwalKeperawatanRanapUpdate struct {
	PenilaianAwalKeperawatanRanapData
	PenilaianAwalKeperawatanRanapDetail
}

func (r *PenilaianAwalKeperawatanRanapUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanRanapUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianAwalKeperawatanRanapRules()
}

func (r *PenilaianAwalKeperawatanRanapUpdate) Payload() PenilaianAwalKeperawatanRanapData {
	return r.PenilaianAwalKeperawatanRanapData
}

func (r *PenilaianAwalKeperawatanRanapUpdate) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}
