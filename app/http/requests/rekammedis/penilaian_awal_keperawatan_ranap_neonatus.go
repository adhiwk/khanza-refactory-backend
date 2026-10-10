package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianAwalKeperawatanRanapNeonatusData isian penilaian awal keperawatan ranap neonatus.
type PenilaianAwalKeperawatanRanapNeonatusData struct {
	Tanggal                                 string `form:"tanggal" json:"tanggal"`
	AsalPasien                              string `form:"asal_pasien" json:"asal_pasien"`
	CaraMasuk                               string `form:"cara_masuk" json:"cara_masuk"`
	DiperolehDari                           string `form:"diperoleh_dari" json:"diperoleh_dari"`
	HubunganDenganPasien                    string `form:"hubungan_dengan_pasien" json:"hubungan_dengan_pasien"`
	KeluhanUtama                            string `form:"keluhan_utama" json:"keluhan_utama"`
	PrenatalG                               string `form:"prenatal_g" json:"prenatal_g"`
	PrenatalP                               string `form:"prenatal_p" json:"prenatal_p"`
	PrenatalA                               string `form:"prenatal_a" json:"prenatal_a"`
	PrenatalUk                              string `form:"prenatal_uk" json:"prenatal_uk"`
	PrenatalRiwayatPenyakitIbu              string `form:"prenatal_riwayat_penyakit_ibu" json:"prenatal_riwayat_penyakit_ibu"`
	PrenatalRiwayatPenyakitIbuKeterangan    string `form:"prenatal_riwayat_penyakit_ibu_keterangan" json:"prenatal_riwayat_penyakit_ibu_keterangan"`
	PrenatalRiwayatPengobatanIbuSelamaHamil string `form:"prenatal_riwayat_pengobatan_ibu_selama_hamil" json:"prenatal_riwayat_pengobatan_ibu_selama_hamil"`
	PrenatalPernahDirawat                   string `form:"prenatal_pernah_dirawat" json:"prenatal_pernah_dirawat"`
	PrenatalPernahDirawatKeterangan         string `form:"prenatal_pernah_dirawat_keterangan" json:"prenatal_pernah_dirawat_keterangan"`
	PrenatalStatusGiziIbu                   string `form:"prenatal_status_gizi_ibu" json:"prenatal_status_gizi_ibu"`
	IntranatalG                             string `form:"intranatal_g" json:"intranatal_g"`
	IntranatalP                             string `form:"intranatal_p" json:"intranatal_p"`
	IntranatalA                             string `form:"intranatal_a" json:"intranatal_a"`
	IntranatalKondisiLahir                  string `form:"intranatal_kondisi_lahir" json:"intranatal_kondisi_lahir"`
	IntranatalCaraPersalinan                string `form:"intranatal_cara_persalinan" json:"intranatal_cara_persalinan"`
	IntranatalCaraPersalinanKeterangan      string `form:"intranatal_cara_persalinan_keterangan" json:"intranatal_cara_persalinan_keterangan"`
	IntranatalApgar                         string `form:"intranatal_apgar" json:"intranatal_apgar"`
	IntranatalLetak                         string `form:"intranatal_letak" json:"intranatal_letak"`
	IntranatalTaliPusat                     string `form:"intranatal_tali_pusat" json:"intranatal_tali_pusat"`
	IntranatalKetuban                       string `form:"intranatal_ketuban" json:"intranatal_ketuban"`
	IntranatalBb                            string `form:"intranatal_bb" json:"intranatal_bb"`
	IntranatalPb                            string `form:"intranatal_pb" json:"intranatal_pb"`
	IntranatalLk                            string `form:"intranatal_lk" json:"intranatal_lk"`
	IntranatalLd                            string `form:"intranatal_ld" json:"intranatal_ld"`
	IntranatalLp                            string `form:"intranatal_lp" json:"intranatal_lp"`
	RisikoInfeksiMayor                      string `form:"risiko_infeksi_mayor" json:"risiko_infeksi_mayor"`
	RisikoInfeksiMayorKeterangan            string `form:"risiko_infeksi_mayor_keterangan" json:"risiko_infeksi_mayor_keterangan"`
	RisikoInfeksiMinor                      string `form:"risiko_infeksi_minor" json:"risiko_infeksi_minor"`
	RisikoInfeksiMinorKeterangan            string `form:"risiko_infeksi_minor_keterangan" json:"risiko_infeksi_minor_keterangan"`
	KebutuhanBiologisNutrisi                string `form:"kebutuhan_biologis_nutrisi" json:"kebutuhan_biologis_nutrisi"`
	KebutuhanBiologisNutrisiKeterangan      string `form:"kebutuhan_biologis_nutrisi_keterangan" json:"kebutuhan_biologis_nutrisi_keterangan"`
	KebutuhanBiologisNutrisiFrekuensi       string `form:"kebutuhan_biologis_nutrisi_frekuensi" json:"kebutuhan_biologis_nutrisi_frekuensi"`
	KebutuhanBiologisNutrisiKali            string `form:"kebutuhan_biologis_nutrisi_kali" json:"kebutuhan_biologis_nutrisi_kali"`
	KebutuhanBiologisBak                    string `form:"kebutuhan_biologis_bak" json:"kebutuhan_biologis_bak"`
	KebutuhanBiologisBakKeterangan          string `form:"kebutuhan_biologis_bak_keterangan" json:"kebutuhan_biologis_bak_keterangan"`
	KebutuhanBiologisBab                    string `form:"kebutuhan_biologis_bab" json:"kebutuhan_biologis_bab"`
	KebutuhanBiologisBabKeterangan          string `form:"kebutuhan_biologis_bab_keterangan" json:"kebutuhan_biologis_bab_keterangan"`
	AlergiObat                              string `form:"alergi_obat" json:"alergi_obat"`
	AlergiObatKeterangan                    string `form:"alergi_obat_keterangan" json:"alergi_obat_keterangan"`
	AlergiObatReaksi                        string `form:"alergi_obat_reaksi" json:"alergi_obat_reaksi"`
	AlergiMakanan                           string `form:"alergi_makanan" json:"alergi_makanan"`
	AlergiMakananKeterangan                 string `form:"alergi_makanan_keterangan" json:"alergi_makanan_keterangan"`
	AlergiMakananReaksi                     string `form:"alergi_makanan_reaksi" json:"alergi_makanan_reaksi"`
	AlergiLainnya                           string `form:"alergi_lainnya" json:"alergi_lainnya"`
	AlergiLainnyaKeterangan                 string `form:"alergi_lainnya_keterangan" json:"alergi_lainnya_keterangan"`
	AlergiLainnyaReaksi                     string `form:"alergi_lainnya_reaksi" json:"alergi_lainnya_reaksi"`
	RiwayatPenyakitKeluarga                 string `form:"riwayat_penyakit_keluarga" json:"riwayat_penyakit_keluarga"`
	RiwayatPenyakitKeluargaKeterangan       string `form:"riwayat_penyakit_keluarga_keterangan" json:"riwayat_penyakit_keluarga_keterangan"`
	RiwayatImunisasi                        string `form:"riwayat_imunisasi" json:"riwayat_imunisasi"`
	RiwayatImunisasiKeterangan              string `form:"riwayat_imunisasi_keterangan" json:"riwayat_imunisasi_keterangan"`
	RiwayatTranfusiDarah                    string `form:"riwayat_tranfusi_darah" json:"riwayat_tranfusi_darah"`
	RiwayatTranfusiDarahKeterangan          string `form:"riwayat_tranfusi_darah_keterangan" json:"riwayat_tranfusi_darah_keterangan"`
	RiwayatTranfusiDarahReaksi              string `form:"riwayat_tranfusi_darah_reaksi" json:"riwayat_tranfusi_darah_reaksi"`
	RiwayatTranfusiDarahReaksiKeterangan    string `form:"riwayat_tranfusi_darah_reaksi_keterangan" json:"riwayat_tranfusi_darah_reaksi_keterangan"`
	KebiasanIbuObatDiminum                  string `form:"kebiasan_ibu_obat_diminum" json:"kebiasan_ibu_obat_diminum"`
	KebiasanIbuObatDiminumKeterangan        string `form:"kebiasan_ibu_obat_diminum_keterangan" json:"kebiasan_ibu_obat_diminum_keterangan"`
	KebiasanIbuNarkoba                      string `form:"kebiasan_ibu_narkoba" json:"kebiasan_ibu_narkoba"`
	KebiasanIbuNarkobaKeterangan            string `form:"kebiasan_ibu_narkoba_keterangan" json:"kebiasan_ibu_narkoba_keterangan"`
	KebiasanIbuMerokok                      string `form:"kebiasan_ibu_merokok" json:"kebiasan_ibu_merokok"`
	KebiasanIbuMerokokKeterangan            string `form:"kebiasan_ibu_merokok_keterangan" json:"kebiasan_ibu_merokok_keterangan"`
	KebiasanIbuAlkohol                      string `form:"kebiasan_ibu_alkohol" json:"kebiasan_ibu_alkohol"`
	KebiasanIbuAlkoholKeterangan            string `form:"kebiasan_ibu_alkohol_keterangan" json:"kebiasan_ibu_alkohol_keterangan"`
	Kesadaran                               string `form:"kesadaran" json:"kesadaran"`
	KeadaanUmum                             string `form:"keadaan_umum" json:"keadaan_umum"`
	Gcs                                     string `form:"gcs" json:"gcs"`
	Td                                      string `form:"td" json:"td"`
	Suhu                                    string `form:"suhu" json:"suhu"`
	Hr                                      string `form:"hr" json:"hr"`
	Rr                                      string `form:"rr" json:"rr"`
	Spo2                                    string `form:"spo2" json:"spo2"`
	DownScore                               string `form:"down_score" json:"down_score"`
	Bb                                      string `form:"bb" json:"bb"`
	Tb                                      string `form:"tb" json:"tb"`
	Lk                                      string `form:"lk" json:"lk"`
	Ld                                      string `form:"ld" json:"ld"`
	Lp                                      string `form:"lp" json:"lp"`
	GdBayi                                  string `form:"gd_bayi" json:"gd_bayi"`
	GdIbu                                   string `form:"gd_ibu" json:"gd_ibu"`
	GdAyah                                  string `form:"gd_ayah" json:"gd_ayah"`
	SarafPusatGerakBayi                     string `form:"saraf_pusat_gerak_bayi" json:"saraf_pusat_gerak_bayi"`
	SarafPusatKepala                        string `form:"saraf_pusat_kepala" json:"saraf_pusat_kepala"`
	SarafPusatKepalaKeterangan              string `form:"saraf_pusat_kepala_keterangan" json:"saraf_pusat_kepala_keterangan"`
	SarafPusatUbunubun                      string `form:"saraf_pusat_ubunubun" json:"saraf_pusat_ubunubun"`
	SarafPusatUbunubunKeterangan            string `form:"saraf_pusat_ubunubun_keterangan" json:"saraf_pusat_ubunubun_keterangan"`
	SarafPusatWajah                         string `form:"saraf_pusat_wajah" json:"saraf_pusat_wajah"`
	SarafPusatWajahKeterangan               string `form:"saraf_pusat_wajah_keterangan" json:"saraf_pusat_wajah_keterangan"`
	SarafPusatKejang                        string `form:"saraf_pusat_kejang" json:"saraf_pusat_kejang"`
	SarafPusatKejangKeterangan              string `form:"saraf_pusat_kejang_keterangan" json:"saraf_pusat_kejang_keterangan"`
	SarafPusatRefleks                       string `form:"saraf_pusat_refleks" json:"saraf_pusat_refleks"`
	SarafPusatRefleksKeterangan             string `form:"saraf_pusat_refleks_keterangan" json:"saraf_pusat_refleks_keterangan"`
	SarafPusatTangisbayi                    string `form:"saraf_pusat_tangisbayi" json:"saraf_pusat_tangisbayi"`
	SarafPusatTangisbayiKeterangan          string `form:"saraf_pusat_tangisbayi_keterangan" json:"saraf_pusat_tangisbayi_keterangan"`
	KardiovaskularDenyutnadi                string `form:"kardiovaskular_denyutnadi" json:"kardiovaskular_denyutnadi"`
	KardiovaskularSirkulasi                 string `form:"kardiovaskular_sirkulasi" json:"kardiovaskular_sirkulasi"`
	KardiovaskularSirkulasiKeterangan       string `form:"kardiovaskular_sirkulasi_keterangan" json:"kardiovaskular_sirkulasi_keterangan"`
	KardiovaskularPulsasi                   string `form:"kardiovaskular_pulsasi" json:"kardiovaskular_pulsasi"`
	KardiovaskularPulsasiKeterangan         string `form:"kardiovaskular_pulsasi_keterangan" json:"kardiovaskular_pulsasi_keterangan"`
	RespirasiPolanafas                      string `form:"respirasi_polanafas" json:"respirasi_polanafas"`
	RespirasiJenispernapasan                string `form:"respirasi_jenispernapasan" json:"respirasi_jenispernapasan"`
	RespirasiJenispernapasanKeterangan      string `form:"respirasi_jenispernapasan_keterangan" json:"respirasi_jenispernapasan_keterangan"`
	RespirasiRetraksi                       string `form:"respirasi_retraksi" json:"respirasi_retraksi"`
	RespirasiAirentry                       string `form:"respirasi_airentry" json:"respirasi_airentry"`
	RespirasiMerintih                       string `form:"respirasi_merintih" json:"respirasi_merintih"`
	RespirasiSuaraNapas                     string `form:"respirasi_suara_napas" json:"respirasi_suara_napas"`
	GastrointestinalMulut                   string `form:"gastrointestinal_mulut" json:"gastrointestinal_mulut"`
	GastrointestinalMulutKeterangan         string `form:"gastrointestinal_mulut_keterangan" json:"gastrointestinal_mulut_keterangan"`
	GastrointestinalLidah                   string `form:"gastrointestinal_lidah" json:"gastrointestinal_lidah"`
	GastrointestinalLidahKeterangan         string `form:"gastrointestinal_lidah_keterangan" json:"gastrointestinal_lidah_keterangan"`
	GastrointestinalTenggorakan             string `form:"gastrointestinal_tenggorakan" json:"gastrointestinal_tenggorakan"`
	GastrointestinalTenggorakanKeterangan   string `form:"gastrointestinal_tenggorakan_keterangan" json:"gastrointestinal_tenggorakan_keterangan"`
	GastrointestinalAbdomen                 string `form:"gastrointestinal_abdomen" json:"gastrointestinal_abdomen"`
	GastrointestinalAbdomenKeterangan       string `form:"gastrointestinal_abdomen_keterangan" json:"gastrointestinal_abdomen_keterangan"`
	GastrointestinalBab                     string `form:"gastrointestinal_bab" json:"gastrointestinal_bab"`
	GastrointestinalBabKeterangan           string `form:"gastrointestinal_bab_keterangan" json:"gastrointestinal_bab_keterangan"`
	GastrointestinalWarnabab                string `form:"gastrointestinal_warnabab" json:"gastrointestinal_warnabab"`
	GastrointestinalWarnababKeterangan      string `form:"gastrointestinal_warnabab_keterangan" json:"gastrointestinal_warnabab_keterangan"`
	GastrointestinalBak                     string `form:"gastrointestinal_bak" json:"gastrointestinal_bak"`
	GastrointestinalBakKeterangan           string `form:"gastrointestinal_bak_keterangan" json:"gastrointestinal_bak_keterangan"`
	GastrointestinalBakwarna                string `form:"gastrointestinal_bakwarna" json:"gastrointestinal_bakwarna"`
	GastrointestinalBakwarnaKeterangan      string `form:"gastrointestinal_bakwarna_keterangan" json:"gastrointestinal_bakwarna_keterangan"`
	NeurologiPosisiMata                     string `form:"neurologi_posisi_mata" json:"neurologi_posisi_mata"`
	NeurologiKelopakMata                    string `form:"neurologi_kelopak_mata" json:"neurologi_kelopak_mata"`
	NeurologiKelopakMataKeterangan          string `form:"neurologi_kelopak_mata_keterangan" json:"neurologi_kelopak_mata_keterangan"`
	NeurologiBesarPupil                     string `form:"neurologi_besar_pupil" json:"neurologi_besar_pupil"`
	NeurologiKonjugtiva                     string `form:"neurologi_konjugtiva" json:"neurologi_konjugtiva"`
	NeurologiKonjugtivaKeterangan           string `form:"neurologi_konjugtiva_keterangan" json:"neurologi_konjugtiva_keterangan"`
	NeurologiSklera                         string `form:"neurologi_sklera" json:"neurologi_sklera"`
	NeurologiSkleraKeterangan               string `form:"neurologi_sklera_keterangan" json:"neurologi_sklera_keterangan"`
	NeurologiPendengaran                    string `form:"neurologi_pendengaran" json:"neurologi_pendengaran"`
	NeurologiPendengaranKeterangan          string `form:"neurologi_pendengaran_keterangan" json:"neurologi_pendengaran_keterangan"`
	NeurologiPenciuman                      string `form:"neurologi_penciuman" json:"neurologi_penciuman"`
	NeurologiPenciumanKeterangan            string `form:"neurologi_penciuman_keterangan" json:"neurologi_penciuman_keterangan"`
	IntegumentWarnaKulit                    string `form:"integument_warna_kulit" json:"integument_warna_kulit"`
	IntegumentWarnaKulitKeterangan          string `form:"integument_warna_kulit_keterangan" json:"integument_warna_kulit_keterangan"`
	IntegumentVernicKaseosa                 string `form:"integument_vernic_kaseosa" json:"integument_vernic_kaseosa"`
	IntegumentVernicKaseosaKeterangan       string `form:"integument_vernic_kaseosa_keterangan" json:"integument_vernic_kaseosa_keterangan"`
	IntegumentTurgor                        string `form:"integument_turgor" json:"integument_turgor"`
	IntegumentLanugo                        string `form:"integument_lanugo" json:"integument_lanugo"`
	IntegumentKulit                         string `form:"integument_kulit" json:"integument_kulit"`
	IntegumentRisikoDekubitas               string `form:"integument_risiko_dekubitas" json:"integument_risiko_dekubitas"`
	Reproduksi                              string `form:"reproduksi" json:"reproduksi"`
	ReproduksiKeterangan                    string `form:"reproduksi_keterangan" json:"reproduksi_keterangan"`
	MuskuloskeletalRekoilTelinga            string `form:"muskuloskeletal_rekoil_telinga" json:"muskuloskeletal_rekoil_telinga"`
	MuskuloskeletalRekoilTelingaKeterangan  string `form:"muskuloskeletal_rekoil_telinga_keterangan" json:"muskuloskeletal_rekoil_telinga_keterangan"`
	MuskuloskeletalLengan                   string `form:"muskuloskeletal_lengan" json:"muskuloskeletal_lengan"`
	MuskuloskeletalLenganKeterangan         string `form:"muskuloskeletal_lengan_keterangan" json:"muskuloskeletal_lengan_keterangan"`
	MuskuloskeletalTungkai                  string `form:"muskuloskeletal_tungkai" json:"muskuloskeletal_tungkai"`
	MuskuloskeletalTungkaiKeterangan        string `form:"muskuloskeletal_tungkai_keterangan" json:"muskuloskeletal_tungkai_keterangan"`
	MuskuloskeletalTelapakKaki              string `form:"muskuloskeletal_telapak_kaki" json:"muskuloskeletal_telapak_kaki"`
	KondisiPsikologis                       string `form:"kondisi_psikologis" json:"kondisi_psikologis"`
	GangguanJiwa                            string `form:"gangguan_jiwa" json:"gangguan_jiwa"`
	MenerimaKondisiBayi                     string `form:"menerima_kondisi_bayi" json:"menerima_kondisi_bayi"`
	StatusMenikah                           string `form:"status_menikah" json:"status_menikah"`
	MasalahPernikahan                       string `form:"masalah_pernikahan" json:"masalah_pernikahan"`
	MasalahPernikahanKeterangan             string `form:"masalah_pernikahan_keterangan" json:"masalah_pernikahan_keterangan"`
	Pekerjaan                               string `form:"pekerjaan" json:"pekerjaan"`
	Agama                                   string `form:"agama" json:"agama"`
	NilaiKepercayaan                        string `form:"nilai_kepercayaan" json:"nilai_kepercayaan"`
	NilaiKepercayaanKeterangan              string `form:"nilai_kepercayaan_keterangan" json:"nilai_kepercayaan_keterangan"`
	Suku                                    string `form:"suku" json:"suku"`
	Pendidikan                              string `form:"pendidikan" json:"pendidikan"`
	Pembayaran                              string `form:"pembayaran" json:"pembayaran"`
	TinggalBersama                          string `form:"tinggal_bersama" json:"tinggal_bersama"`
	TinggalBersamaKeterangan                string `form:"tinggal_bersama_keterangan" json:"tinggal_bersama_keterangan"`
	HubunganKeluarga                        string `form:"hubungan_keluarga" json:"hubungan_keluarga"`
	ResponEmosi                             string `form:"respon_emosi" json:"respon_emosi"`
	BahasaSehariHari                        string `form:"bahasa_sehari_hari" json:"bahasa_sehari_hari"`
	KemampuanBacatulis                      string `form:"kemampuan_bacatulis" json:"kemampuan_bacatulis"`
	ButuhPenterjemah                        string `form:"butuh_penterjemah" json:"butuh_penterjemah"`
	ButuhPenterjemahKeterangan              string `form:"butuh_penterjemah_keterangan" json:"butuh_penterjemah_keterangan"`
	TerdapatHambatanBelajar                 string `form:"terdapat_hambatan_belajar" json:"terdapat_hambatan_belajar"`
	HambatanBelajar                         string `form:"hambatan_belajar" json:"hambatan_belajar"`
	HambatanBelajarKeterangan               string `form:"hambatan_belajar_keterangan" json:"hambatan_belajar_keterangan"`
	HambatanCaraBicara                      string `form:"hambatan_cara_bicara" json:"hambatan_cara_bicara"`
	HambatanBahasaIsyarat                   string `form:"hambatan_bahasa_isyarat" json:"hambatan_bahasa_isyarat"`
	CaraBelajarDisukai                      string `form:"cara_belajar_disukai" json:"cara_belajar_disukai"`
	KesediaanMenerimaInformasi              string `form:"kesediaan_menerima_informasi" json:"kesediaan_menerima_informasi"`
	KesediaanMenerimaInformasiKeterangan    string `form:"kesediaan_menerima_informasi_keterangan" json:"kesediaan_menerima_informasi_keterangan"`
	PemahamanNutrisi                        string `form:"pemahaman_nutrisi" json:"pemahaman_nutrisi"`
	PemahamanPenyakit                       string `form:"pemahaman_penyakit" json:"pemahaman_penyakit"`
	PemahamanPengobatan                     string `form:"pemahaman_pengobatan" json:"pemahaman_pengobatan"`
	PemahamanPerawatan                      string `form:"pemahaman_perawatan" json:"pemahaman_perawatan"`
	MasalahGizi1                            string `form:"masalah_gizi1" json:"masalah_gizi1"`
	NilaiGizi1                              string `form:"nilai_gizi1" json:"nilai_gizi1"`
	MasalahGizi2                            string `form:"masalah_gizi2" json:"masalah_gizi2"`
	NilaiGizi2                              string `form:"nilai_gizi2" json:"nilai_gizi2"`
	MasalahGizi3                            string `form:"masalah_gizi3" json:"masalah_gizi3"`
	NilaiGizi3                              string `form:"nilai_gizi3" json:"nilai_gizi3"`
	Totalgizi                               int    `form:"totalgizi" json:"totalgizi"`
	KeteranganGizi                          string `form:"keterangan_gizi" json:"keterangan_gizi"`
	PenilaianHumptydumptySkala1             string `form:"penilaian_humptydumpty_skala1" json:"penilaian_humptydumpty_skala1"`
	PenilaianHumptydumptyNilai1             int    `form:"penilaian_humptydumpty_nilai1" json:"penilaian_humptydumpty_nilai1"`
	PenilaianHumptydumptySkala2             string `form:"penilaian_humptydumpty_skala2" json:"penilaian_humptydumpty_skala2"`
	PenilaianHumptydumptyNilai2             int    `form:"penilaian_humptydumpty_nilai2" json:"penilaian_humptydumpty_nilai2"`
	PenilaianHumptydumptySkala3             string `form:"penilaian_humptydumpty_skala3" json:"penilaian_humptydumpty_skala3"`
	PenilaianHumptydumptyNilai3             int    `form:"penilaian_humptydumpty_nilai3" json:"penilaian_humptydumpty_nilai3"`
	PenilaianHumptydumptySkala4             string `form:"penilaian_humptydumpty_skala4" json:"penilaian_humptydumpty_skala4"`
	PenilaianHumptydumptyNilai4             int    `form:"penilaian_humptydumpty_nilai4" json:"penilaian_humptydumpty_nilai4"`
	PenilaianHumptydumptySkala5             string `form:"penilaian_humptydumpty_skala5" json:"penilaian_humptydumpty_skala5"`
	PenilaianHumptydumptyNilai5             int    `form:"penilaian_humptydumpty_nilai5" json:"penilaian_humptydumpty_nilai5"`
	PenilaianHumptydumptySkala6             string `form:"penilaian_humptydumpty_skala6" json:"penilaian_humptydumpty_skala6"`
	PenilaianHumptydumptyNilai6             int    `form:"penilaian_humptydumpty_nilai6" json:"penilaian_humptydumpty_nilai6"`
	PenilaianHumptydumptySkala7             string `form:"penilaian_humptydumpty_skala7" json:"penilaian_humptydumpty_skala7"`
	PenilaianHumptydumptyNilai7             int    `form:"penilaian_humptydumpty_nilai7" json:"penilaian_humptydumpty_nilai7"`
	PenilaianHumptydumptyTotalnilai         int    `form:"penilaian_humptydumpty_totalnilai" json:"penilaian_humptydumpty_totalnilai"`
	PenilaianHumptydumptyHasil              string `form:"penilaian_humptydumpty_hasil" json:"penilaian_humptydumpty_hasil"`
	SkalaNips1                              string `form:"skala_nips1" json:"skala_nips1"`
	SkalaNips1Nilai                         string `form:"skala_nips1_nilai" json:"skala_nips1_nilai"`
	SkalaNips2                              string `form:"skala_nips2" json:"skala_nips2"`
	SkalaNips2Nilai                         string `form:"skala_nips2_nilai" json:"skala_nips2_nilai"`
	SkalaNips3                              string `form:"skala_nips3" json:"skala_nips3"`
	SkalaNips3Nilai                         string `form:"skala_nips3_nilai" json:"skala_nips3_nilai"`
	SkalaNips4                              string `form:"skala_nips4" json:"skala_nips4"`
	SkalaNips4Nilai                         string `form:"skala_nips4_nilai" json:"skala_nips4_nilai"`
	SkalaNips5                              string `form:"skala_nips5" json:"skala_nips5"`
	SkalaNips5Nilai                         string `form:"skala_nips5_nilai" json:"skala_nips5_nilai"`
	SkalaNipsTotal                          int    `form:"skala_nips_total" json:"skala_nips_total"`
	SkalaNipsKeterangan                     string `form:"skala_nips_keterangan" json:"skala_nips_keterangan"`
	InformasiPerencanaanPulang              string `form:"informasi_perencanaan_pulang" json:"informasi_perencanaan_pulang"`
	LamaRatarata                            string `form:"lama_ratarata" json:"lama_ratarata"`
	PerencanaanPulang                       string `form:"perencanaan_pulang" json:"perencanaan_pulang"`
	KondisiKlinisPulang                     string `form:"kondisi_klinis_pulang" json:"kondisi_klinis_pulang"`
	PerawatanLanjutanDirumah                string `form:"perawatan_lanjutan_dirumah" json:"perawatan_lanjutan_dirumah"`
	CaraTransportasiPulang                  string `form:"cara_transportasi_pulang" json:"cara_transportasi_pulang"`
	TransportasiDigunakan                   string `form:"transportasi_digunakan" json:"transportasi_digunakan"`
	Rencana                                 string `form:"rencana" json:"rencana"`
	Nip1                                    string `form:"nip1" json:"nip1"`
	Nip2                                    string `form:"nip2" json:"nip2"`
	KdDokter                                string `form:"kd_dokter" json:"kd_dokter"`
}

// PenilaianAwalKeperawatanRanapNeonatusDetail daftar kode detail (diganti seluruhnya saat simpan/ubah).
type PenilaianAwalKeperawatanRanapNeonatusDetail struct {
	KodeMasalah []string `form:"kode_masalah" json:"kode_masalah"`
	KodeRencana []string `form:"kode_rencana" json:"kode_rencana"`
}

func penilaianAwalKeperawatanRanapNeonatusRules() map[string]any {
	rules := map[string]any{
		"tanggal":                       "required|date",
		"asal_pasien":                   "in:IGD,Poliklinik,Kamar Bersalin,Kamar Operasi",
		"cara_masuk":                    "in:Baby Box,Inkubator,Kursi Roda",
		"diperoleh_dari":                "string|max_len:30",
		"hubungan_dengan_pasien":        "string|max_len:30",
		"keluhan_utama":                 "string|max_len:300",
		"prenatal_g":                    "string|max_len:10",
		"prenatal_p":                    "string|max_len:10",
		"prenatal_a":                    "string|max_len:10",
		"prenatal_uk":                   "string|max_len:30",
		"prenatal_riwayat_penyakit_ibu": "in:Tidak Ada,DM,Hipertensi,Jantung,TBC,Hep B,Asma,PMS,Lainnya",
		"prenatal_riwayat_penyakit_ibu_keterangan":     "string|max_len:30",
		"prenatal_riwayat_pengobatan_ibu_selama_hamil": "string|max_len:100",
		"prenatal_pernah_dirawat":                      "in:Ya,Tidak",
		"prenatal_pernah_dirawat_keterangan":           "string|max_len:50",
		"prenatal_status_gizi_ibu":                     "in:Baik,Buruk",
		"intranatal_g":                                 "string|max_len:10",
		"intranatal_p":                                 "string|max_len:10",
		"intranatal_a":                                 "string|max_len:10",
		"intranatal_kondisi_lahir":                     "string|max_len:30",
		"intranatal_cara_persalinan":                   "in:Spontan,Vacum Ekstraksi,Forcep Ekstraksi,Sectio Caesarea,Lainnya",
		"intranatal_cara_persalinan_keterangan":        "string|max_len:30",
		"intranatal_apgar":                             "string|max_len:10",
		"intranatal_letak":                             "string|max_len:30",
		"intranatal_tali_pusat":                        "in:Segar,Layu,Simpul",
		"intranatal_ketuban":                           "required|in:Jernih,Darah,Putih Keruh,Hijau,Meconium",
		"intranatal_bb":                                "string|max_len:5",
		"intranatal_pb":                                "string|max_len:5",
		"intranatal_lk":                                "string|max_len:5",
		"intranatal_ld":                                "string|max_len:5",
		"intranatal_lp":                                "string|max_len:5",
		"risiko_infeksi_mayor":                         "required|in:-,Ibu Demam >= 38Â° C,KPD > 24 Jam,Ketuban Hijau,Korioamniotis,Fetal Distress",
		"risiko_infeksi_mayor_keterangan":              "string|max_len:30",
		"risiko_infeksi_minor":                         "required|in:-,KPD > 12 Jam,Asfiksia,BBLR,ISK,UK < 37 Minggu,Gemeli,Keputihan,Ibu Temperatur > 37Â° C",
		"risiko_infeksi_minor_keterangan":              "string|max_len:30",
		"kebutuhan_biologis_nutrisi":                   "required|in:Air Susu Ibu,Lainnya",
		"kebutuhan_biologis_nutrisi_keterangan":        "string|max_len:50",
		"kebutuhan_biologis_nutrisi_frekuensi":         "string|max_len:5",
		"kebutuhan_biologis_nutrisi_kali":              "string|max_len:3",
		"kebutuhan_biologis_bak":                       "required|in:Ada Keluhan,Tidak Ada Keluhan",
		"kebutuhan_biologis_bak_keterangan":            "string|max_len:30",
		"kebutuhan_biologis_bab":                       "required|in:Ada Keluhan,Tidak Ada Keluhan",
		"kebutuhan_biologis_bab_keterangan":            "string|max_len:30",
		"alergi_obat":                                  "required|in:Tidak Ada,Tidak Diketahui,Ada",
		"alergi_obat_keterangan":                       "string|max_len:40",
		"alergi_obat_reaksi":                           "string|max_len:40",
		"alergi_makanan":                               "required|in:Tidak Ada,Tidak Diketahui,Ada",
		"alergi_makanan_keterangan":                    "string|max_len:40",
		"alergi_makanan_reaksi":                        "string|max_len:40",
		"alergi_lainnya":                               "required|in:Tidak Ada,Tidak Diketahui,Ada",
		"alergi_lainnya_keterangan":                    "string|max_len:40",
		"alergi_lainnya_reaksi":                        "string|max_len:40",
		"riwayat_penyakit_keluarga":                    "required|in:Tidak Ada,Diabetes,Kanker,Asma,Hipertensi,Jantung,Lainnya",
		"riwayat_penyakit_keluarga_keterangan":         "string|max_len:70",
		"riwayat_imunisasi":                            "required|in:Tidak Ada,Ada",
		"riwayat_imunisasi_keterangan":                 "string|max_len:70",
		"riwayat_tranfusi_darah":                       "required|in:Tidak Ada,Ada",
		"riwayat_tranfusi_darah_keterangan":            "string|max_len:30",
		"riwayat_tranfusi_darah_reaksi":                "required|in:Tidak Ada,Ada",
		"riwayat_tranfusi_darah_reaksi_keterangan":     "string|max_len:30",
		"kebiasan_ibu_obat_diminum":                    "required|in:Tidak Ada,Vitamin,Jamu-jamuan,Lainnya",
		"kebiasan_ibu_obat_diminum_keterangan":         "string|max_len:40",
		"kebiasan_ibu_narkoba":                         "required|in:Tidak Ada,Ada",
		"kebiasan_ibu_narkoba_keterangan":              "string|max_len:40",
		"kebiasan_ibu_merokok":                         "required|in:Tidak,Ya",
		"kebiasan_ibu_merokok_keterangan":              "string|max_len:3",
		"kebiasan_ibu_alkohol":                         "required|in:Tidak,Ya",
		"kebiasan_ibu_alkohol_keterangan":              "string|max_len:3",
		"kesadaran":                                    "required|in:Compos Mentis,Apatis,Somnolen,Sopor,Soporcoma,Coma",
		"keadaan_umum":                                 "required|in:Tampak Tidak Sakit,Sakit Ringan,Sakit Sedang,Sakit Berat",
		"gcs":                                          "string|max_len:10",
		"td":                                           "string|max_len:8",
		"suhu":                                         "string|max_len:5",
		"hr":                                           "string|max_len:5",
		"rr":                                           "string|max_len:5",
		"spo2":                                         "string|max_len:5",
		"down_score":                                   "string|max_len:5",
		"bb":                                           "string|max_len:5",
		"tb":                                           "string|max_len:5",
		"lk":                                           "string|max_len:5",
		"ld":                                           "string|max_len:5",
		"lp":                                           "string|max_len:5",
		"gd_bayi":                                      "required|in:A +,A -,B +,B -,AB +,AB -,O +,O -,-",
		"gd_ibu":                                       "required|in:A +,A -,B +,B -,AB +,AB -,O +,O -,-",
		"gd_ayah":                                      "required|in:A +,A -,B +,B -,AB +,AB -,O +,O -,-",
		"saraf_pusat_gerak_bayi":                       "required|in:Aktif,Tidak Aktif",
		"saraf_pusat_kepala":                           "required|in:TAK,Hydrocephalus,Hematoma,Lainnya",
		"saraf_pusat_kepala_keterangan":                "string|max_len:30",
		"saraf_pusat_ubunubun":                         "required|in:Datar,Cekung,Menonjol,Lainnya",
		"saraf_pusat_ubunubun_keterangan":              "string|max_len:30",
		"saraf_pusat_wajah":                            "required|in:TAK,Asimetris,Kelainan Kongenital,Lainnya",
		"saraf_pusat_wajah_keterangan":                 "string|max_len:30",
		"saraf_pusat_kejang":                           "required|in:Tidak Ada,Ada",
		"saraf_pusat_kejang_keterangan":                "string|max_len:30",
		"saraf_pusat_refleks":                          "required|in:Moro,Menelan,Hisap,Babinski,Rooting,Lainnya",
		"saraf_pusat_refleks_keterangan":               "string|max_len:30",
		"saraf_pusat_tangisbayi":                       "required|in:Kuat,Melengking,Lainnya",
		"saraf_pusat_tangisbayi_keterangan":            "string|max_len:30",
		"kardiovaskular_denyutnadi":                    "required|in:Teratur,Tidak Teratur",
		"kardiovaskular_sirkulasi":                     "required|in:Akral Hangat,Akral Dingin,CRT,Edema",
		"kardiovaskular_sirkulasi_keterangan":          "string|max_len:30",
		"kardiovaskular_pulsasi":                       "required|in:Kuat,Lemah,Lainnya",
		"kardiovaskular_pulsasi_keterangan":            "string|max_len:30",
		"respirasi_polanafas":                          "required|in:Normal,Bradipnea,Tachipnea,Hiperventilasi,Othopneu,Kusmaul,Apneu",
		"respirasi_jenispernapasan":                    "required|in:Pernapasan Dada,Pernapasan Perut,Bahu Diangkat,Cuping Hidung,Alat Bantu Napas",
		"respirasi_jenispernapasan_keterangan":         "string|max_len:30",
		"respirasi_retraksi":                           "required|in:Tidak Ada,Ringan,Berat",
		"respirasi_airentry":                           "required|in:Udara Masuk,Penurunan Udara Masuk,Tidak Ada Udara Masuk",
		"respirasi_merintih":                           "required|in:Tidak Ada,Terdengar Dengan Stetoskop,Terdengar Tanpa Stetoskop",
		"respirasi_suara_napas":                        "required|in:Vesikuler,Wheezing,Ronkhi,Stridor",
		"gastrointestinal_mulut":                       "required|in:TAK,Simetris,Asimetri,Bibir Pucat,Lainnya",
		"gastrointestinal_mulut_keterangan":            "string|max_len:30",
		"gastrointestinal_lidah":                       "required|in:TAK,Kotor,Gerak Asimetris,Lainnya",
		"gastrointestinal_lidah_keterangan":            "string|max_len:30",
		"gastrointestinal_tenggorakan":                 "required|in:TAK,Ada Kelainan",
		"gastrointestinal_tenggorakan_keterangan":      "string|max_len:30",
		"gastrointestinal_abdomen":                     "required|in:Supel,Asites,Tegang,BU,Lainnya",
		"gastrointestinal_abdomen_keterangan":          "string|max_len:30",
		"gastrointestinal_bab":                         "required|in:Normal,Konstipasi,Melena,Colostomy,Diare,Meco Pertama",
		"gastrointestinal_bab_keterangan":              "string|max_len:30",
		"gastrointestinal_warnabab":                    "required|in:Kuning,Dempul,Coklat,Hijau,Lainnya",
		"gastrointestinal_warnabab_keterangan":         "string|max_len:30",
		"gastrointestinal_bak":                         "required|in:Normal,Hematuri,Urin Menetes,Oliguri,BAK Pertama",
		"gastrointestinal_bak_keterangan":              "string|max_len:30",
		"gastrointestinal_bakwarna":                    "required|in:Jernih,Kuning,Kuning Pekat,Lainnya",
		"gastrointestinal_bakwarna_keterangan":         "string|max_len:30",
		"neurologi_posisi_mata":                        "required|in:Simetris,Asimetris",
		"neurologi_kelopak_mata":                       "required|in:TAK,Edema,Cekung,Lainnya",
		"neurologi_kelopak_mata_keterangan":            "string|max_len:30",
		"neurologi_besar_pupil":                        "required|in:Isokor,Anisokor",
		"neurologi_konjugtiva":                         "required|in:TAK,Anemis,Konjungtivitis,Lainnya",
		"neurologi_konjugtiva_keterangan":              "string|max_len:30",
		"neurologi_sklera":                             "required|in:TAK,Ikterik,Perdarahan,Lainnya",
		"neurologi_sklera_keterangan":                  "string|max_len:30",
		"neurologi_pendengaran":                        "required|in:TAK,Asimetris,Keluar Cairan,Lainnya",
		"neurologi_pendengaran_keterangan":             "string|max_len:30",
		"neurologi_penciuman":                          "required|in:TAK,Asimetris,Keluar Cairan,Lainnya",
		"neurologi_penciuman_keterangan":               "string|max_len:30",
		"integument_warna_kulit":                       "required|in:Pucat,Sianosis,Normal,Lainnya",
		"integument_warna_kulit_keterangan":            "string|max_len:30",
		"integument_vernic_kaseosa":                    "required|in:Ada,Tidak Ada,Lainnya",
		"integument_vernic_kaseosa_keterangan":         "string|max_len:30",
		"integument_turgor":                            "required|in:Baik,Sedang,Buruk",
		"integument_lanugo":                            "required|in:Tidak Ada,Banyak,Tipis,Bercak-bercak Tanpa Lanugo,Sebagian Besar Tanpa Lanugo",
		"integument_kulit":                             "required|in:Normal,Rash/Kemerahan,Luka,Memar,Ptekie,Bula",
		"integument_risiko_dekubitas":                  "required|in:Jaringan/Elastisitas Kurang,Imobilisasi",
		"reproduksi":                                   "required|in:Normal,Hipospadia,Epispadia,Fimosis,Hidrokel,Keputihan,Vagina Skintag,Lainnya",
		"reproduksi_keterangan":                        "string|max_len:70",
		"muskuloskeletal_rekoil_telinga":               "required|in:Rekoil Lambat,Rekoil Cepat,Rekoil Segera,Lainnya",
		"muskuloskeletal_rekoil_telinga_keterangan":    "string|max_len:30",
		"muskuloskeletal_lengan":                       "required|in:Fleksi,Ekstensi,Pergerakan Aktif,Pergerakan Tidak Aktif,Lainnya",
		"muskuloskeletal_lengan_keterangan":            "string|max_len:30",
		"muskuloskeletal_tungkai":                      "required|in:Fleksi,Ekstensi,Pergerakan Aktif,Pergerakan Tidak Aktif,Lainnya",
		"muskuloskeletal_tungkai_keterangan":           "string|max_len:30",
		"muskuloskeletal_telapak_kaki":                 "required|in:Tipis,Garis Transversal Anterior,Garis 2/3 Anterior,Seluruh Telapak Kaki",
		"kondisi_psikologis":                           "required|in:Tidak Ada Masalah,Marah,Takut,Depresi,Cepat Lelah,Cemas,Gelisah,Sulit Tidur,Lainnya",
		"gangguan_jiwa":                                "required|in:Tidak,Ya",
		"menerima_kondisi_bayi":                        "required|in:Menerima,Tidak Menerima",
		"status_menikah":                               "required|in:Menikah,Belum Menikah,Janda,Duda",
		"masalah_pernikahan":                           "required|in:Tidak Ada,Cerai/Istri Baru,Simpanan,Lainnya",
		"masalah_pernikahan_keterangan":                "string|max_len:30",
		"pekerjaan":                                    "string|max_len:20",
		"agama":                                        "string|max_len:20",
		"nilai_kepercayaan":                            "required|in:Tidak Ada,Ada",
		"nilai_kepercayaan_keterangan":                 "string|max_len:40",
		"suku":                                         "string|max_len:20",
		"pendidikan":                                   "string|max_len:20",
		"pembayaran":                                   "string|max_len:40",
		"tinggal_bersama":                              "required|in:Suami,Orang Tua, Teman,Anak,Sendiri,Lainnya",
		"tinggal_bersama_keterangan":                   "string|max_len:30",
		"hubungan_keluarga":                            "required|in:Harmonis,Kurang Harmonis,Tidak Harmonis,Konflik Besar",
		"respon_emosi":                                 "required|in:Takut Terhadap Terapi / Pembedahan / Lingkungan RS,Marah / Tegang,Sedih,Menangis,Senang,Takut,Mampu Menahan Diri,Cemas,Rendah Diri,Gelisah,Tenang,Mudah Tersinggung",
		"bahasa_sehari_hari":                           "string|max_len:20",
		"kemampuan_bacatulis":                          "required|in:Baik, Kurang,Tidak Bisa",
		"butuh_penterjemah":                            "required|in:Tidak,Ya",
		"butuh_penterjemah_keterangan":                 "string|max_len:20",
		"terdapat_hambatan_belajar":                    "required|in:Tidak,Ya",
		"hambatan_belajar":                             "required|in:-,Gangguan Pendengaran,Gangguan Penglihatan,Gangguan Kognitif,Gangguan Fisik,Gangguan Emosi,Keterbatasan Bahasa,Keterbatasan Budaya,Keterbatasan Spiritual,Agama,Lainnya",
		"hambatan_belajar_keterangan":                  "string|max_len:40",
		"hambatan_cara_bicara":                         "required|in:Normal,Gangguan Bicara",
		"hambatan_bahasa_isyarat":                      "required|in:Tidak,Ya",
		"cara_belajar_disukai":                         "required|in:Audio,Lisan,Visual,Demonstrasi,Tulisan",
		"kesediaan_menerima_informasi":                 "required|in:Ya,Tidak",
		"kesediaan_menerima_informasi_keterangan":      "string|max_len:40",
		"pemahaman_nutrisi":                            "required|in:Ya,Tidak",
		"pemahaman_penyakit":                           "required|in:Ya,Tidak",
		"pemahaman_pengobatan":                         "required|in:Ya,Tidak",
		"pemahaman_perawatan":                          "required|in:Ya,Tidak",
		"masalah_gizi1":                                "required|in:Tidak,Ya",
		"nilai_gizi1":                                  "required|in:0,1",
		"masalah_gizi2":                                "required|in:Tidak,Ya",
		"nilai_gizi2":                                  "required|in:0,1",
		"masalah_gizi3":                                "required|in:Tidak,Ya",
		"nilai_gizi3":                                  "required|in:0,1",
		"totalgizi":                                    "int",
		"keterangan_gizi":                              "string|max_len:50",
		"penilaian_humptydumpty_skala1":                "required|in:0 - 3 Tahun,3 - 7 Tahun,7 - 13 Tahun,> 13 Tahun",
		"penilaian_humptydumpty_nilai1":                "int",
		"penilaian_humptydumpty_skala2":                "required|in:Laki-laki,Perempuan",
		"penilaian_humptydumpty_nilai2":                "int",
		"penilaian_humptydumpty_skala3":                "required|string",
		"penilaian_humptydumpty_nilai3":                "int",
		"penilaian_humptydumpty_skala4":                "required|in:Tidak Sadar Terhadap Keterbatasan,Lupa Keterbatasan,Mengetahui Kemampuan Diri",
		"penilaian_humptydumpty_nilai4":                "int",
		"penilaian_humptydumpty_skala5":                "required|in:Riwayat Jatuh Dari Tempat Tidur Saat Bayi/Anak,Pasien Menggunakan Alat Bantu/Box/Mebel,Pasien Berada Di Tempat Tidur,Di Luar Ruang Rawat",
		"penilaian_humptydumpty_nilai5":                "int",
		"penilaian_humptydumpty_skala6":                "required|in:Dalam 24 Jam,Dalam 48 Jam,> 48 Jam",
		"penilaian_humptydumpty_nilai6":                "int",
		"penilaian_humptydumpty_skala7":                "required|string",
		"penilaian_humptydumpty_nilai7":                "int",
		"penilaian_humptydumpty_totalnilai":            "int",
		"penilaian_humptydumpty_hasil":                 "required|in:Risiko Rendah 7 - 11,Risiko Tinggi >=12",
		"skala_nips1":                                  "required|in:Otot Relaks / Wajah Tenang / Ekspresi Netral,Meringis / Otot Wajah Tegang / Alis Berkerut (Ekspresi Wajah Negatif)",
		"skala_nips1_nilai":                            "required|in:0,1",
		"skala_nips2":                                  "required|in:Tenang / Tidak Menangis,Merengek / Mengerang Lemah Intermiten,Menangis Keras / Melengking Terus Menerus / Menangis Tanpa Suara Bila Bayi Intubasi",
		"skala_nips2_nilai":                            "required|in:0,1,2",
		"skala_nips3":                                  "required|in:Relaks / Bernapas Biasa,Perubahan Napas / Tarikan Irreguler / Lebih Cepat Dibandingkan Biasa / Menahan Napas / Tersedak",
		"skala_nips3_nilai":                            "required|in:0,1",
		"skala_nips4":                                  "required|in:Relaks / Tidak Ada Kekakuan Otot / Gerakan Tungkai Biasa,Fleksi / Ekstensi Tegang Kaku",
		"skala_nips4_nilai":                            "required|in:0,1",
		"skala_nips5":                                  "required|in:Tenang / Tidur Lelap / Bangun,Gelisah",
		"skala_nips5_nilai":                            "required|in:0,1",
		"skala_nips_total":                             "int",
		"skala_nips_keterangan":                        "required|in:0 : Tidak Nyeri,1-2 : Nyeri Ringan,3-4 : Nyeri Sedang,> 4 : Nyeri Hebat",
		"informasi_perencanaan_pulang":                 "required|in:Ya,Tidak",
		"lama_ratarata":                                "string|max_len:3",
		"perencanaan_pulang":                           "required|date",
		"kondisi_klinis_pulang":                        "string|max_len:100",
		"perawatan_lanjutan_dirumah":                   "string|max_len:300",
		"cara_transportasi_pulang":                     "required|in:Mandiri,Dibantu Sebagian,Dibantu Keseluruhan,Menggunakan Rostul,Menggunakan Brankar,Berjalan",
		"transportasi_digunakan":                       "required|in:Kendaraan Pribadi,Mobil Ambulance,Kendaraan Umum",
		"rencana":                                      "string|max_len:200",
		"nip1":                                         "required|string|max_len:20",
		"nip2":                                         "required|string|max_len:20",
		"kd_dokter":                                    "required|string|max_len:20",
	}
	rules["kode_masalah"] = "slice"
	rules["kode_masalah.*"] = "string|max_len:20"
	rules["kode_rencana"] = "slice"
	rules["kode_rencana.*"] = "string|max_len:20"
	return rules
}

// PenilaianAwalKeperawatanRanapNeonatusStore simpan penilaian awal keperawatan ranap neonatus; kolom waktu kunci kosong = sekarang.
type PenilaianAwalKeperawatanRanapNeonatusStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianAwalKeperawatanRanapNeonatusData
	PenilaianAwalKeperawatanRanapNeonatusDetail
}

func (r *PenilaianAwalKeperawatanRanapNeonatusStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanRanapNeonatusStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianAwalKeperawatanRanapNeonatusRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianAwalKeperawatanRanapNeonatusStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianAwalKeperawatanRanapNeonatusStore) Payload() PenilaianAwalKeperawatanRanapNeonatusData {
	return r.PenilaianAwalKeperawatanRanapNeonatusData
}

func (r *PenilaianAwalKeperawatanRanapNeonatusStore) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}

// PenilaianAwalKeperawatanRanapNeonatusUpdate ubah penilaian awal keperawatan ranap neonatus (PUT); kunci lewat query string.
type PenilaianAwalKeperawatanRanapNeonatusUpdate struct {
	PenilaianAwalKeperawatanRanapNeonatusData
	PenilaianAwalKeperawatanRanapNeonatusDetail
}

func (r *PenilaianAwalKeperawatanRanapNeonatusUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanRanapNeonatusUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianAwalKeperawatanRanapNeonatusRules()
}

func (r *PenilaianAwalKeperawatanRanapNeonatusUpdate) Payload() PenilaianAwalKeperawatanRanapNeonatusData {
	return r.PenilaianAwalKeperawatanRanapNeonatusData
}

func (r *PenilaianAwalKeperawatanRanapNeonatusUpdate) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}
