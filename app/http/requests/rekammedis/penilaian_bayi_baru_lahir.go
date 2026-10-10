package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianBayiBaruLahirData isian penilaian bayi baru lahir.
type PenilaianBayiBaruLahirData struct {
	Tanggal                              string `form:"tanggal" json:"tanggal"`
	KdDokter                             string `form:"kd_dokter" json:"kd_dokter"`
	NoRkmMedisIbu                        string `form:"no_rkm_medis_ibu" json:"no_rkm_medis_ibu"`
	PenyakitDideritaIbu                  string `form:"penyakit_diderita_ibu" json:"penyakit_diderita_ibu"`
	KeteranganPenyakitDideritaIbu        string `form:"keterangan_penyakit_diderita_ibu" json:"keterangan_penyakit_diderita_ibu"`
	ObatDikonsumsiSelamaKehamilan        string `form:"obat_dikonsumsi_selama_kehamilan" json:"obat_dikonsumsi_selama_kehamilan"`
	PerawatanAntenatal                   string `form:"perawatan_antenatal" json:"perawatan_antenatal"`
	KeteranganPerawatanAntenatal         string `form:"keterangan_perawatan_antenatal" json:"keterangan_perawatan_antenatal"`
	TerdaftarEkohort                     string `form:"terdaftar_ekohort" json:"terdaftar_ekohort"`
	KeteranganTerdaftarEkohort           string `form:"keterangan_terdaftar_ekohort" json:"keterangan_terdaftar_ekohort"`
	PenyulitKehamilan                    string `form:"penyulit_kehamilan" json:"penyulit_kehamilan"`
	KeteranganPenyulitKehamilan          string `form:"keterangan_penyulit_kehamilan" json:"keterangan_penyulit_kehamilan"`
	Alergi                               string `form:"alergi" json:"alergi"`
	KeteranganLainnyaRiwayatMaternal     string `form:"keterangan_lainnya_riwayat_maternal" json:"keterangan_lainnya_riwayat_maternal"`
	UmurKehamilan                        string `form:"umur_kehamilan" json:"umur_kehamilan"`
	Kehamilan                            string `form:"kehamilan" json:"kehamilan"`
	KeteranganKehamilan                  string `form:"keterangan_kehamilan" json:"keterangan_kehamilan"`
	UrutanKehamilan                      string `form:"urutan_kehamilan" json:"urutan_kehamilan"`
	JamKetubanPecah                      string `form:"jam_ketuban_pecah" json:"jam_ketuban_pecah"`
	MenitKetubanPecah                    string `form:"menit_ketuban_pecah" json:"menit_ketuban_pecah"`
	JumlahAirKetuban                     string `form:"jumlah_air_ketuban" json:"jumlah_air_ketuban"`
	WarnaAirKetuban                      string `form:"warna_air_ketuban" json:"warna_air_ketuban"`
	BauAirKetuban                        string `form:"bau_air_ketuban" json:"bau_air_ketuban"`
	LetakBayi                            string `form:"letak_bayi" json:"letak_bayi"`
	MacamPersalinan                      string `form:"macam_persalinan" json:"macam_persalinan"`
	KeteranganMacamPersalinan            string `form:"keterangan_macam_persalinan" json:"keterangan_macam_persalinan"`
	IndikasiPersalinanOperatif           string `form:"indikasi_persalinan_operatif" json:"indikasi_persalinan_operatif"`
	KeteranganIndikasiPersalinanOperatif string `form:"keterangan_indikasi_persalinan_operatif" json:"keterangan_indikasi_persalinan_operatif"`
	LamaGawatJanin                       string `form:"lama_gawat_janin" json:"lama_gawat_janin"`
	ObatSelamaPersalinan                 string `form:"obat_selama_persalinan" json:"obat_selama_persalinan"`
	BeratPlacenta                        string `form:"berat_placenta" json:"berat_placenta"`
	KelainanPlacenta                     string `form:"kelainan_placenta" json:"kelainan_placenta"`
	KeteranganLainnyaRiwayatPersalinan   string `form:"keterangan_lainnya_riwayat_persalinan" json:"keterangan_lainnya_riwayat_persalinan"`
	F1                                   string `form:"f1" json:"f1"`
	U1                                   string `form:"u1" json:"u1"`
	T1                                   string `form:"t1" json:"t1"`
	R1                                   string `form:"r1" json:"r1"`
	W1                                   string `form:"w1" json:"w1"`
	N1                                   string `form:"n1" json:"n1"`
	F5                                   string `form:"f5" json:"f5"`
	U5                                   string `form:"u5" json:"u5"`
	T5                                   string `form:"t5" json:"t5"`
	R5                                   string `form:"r5" json:"r5"`
	W5                                   string `form:"w5" json:"w5"`
	N5                                   string `form:"n5" json:"n5"`
	F10                                  string `form:"f10" json:"f10"`
	U10                                  string `form:"u10" json:"u10"`
	T10                                  string `form:"t10" json:"t10"`
	R10                                  string `form:"r10" json:"r10"`
	W10                                  string `form:"w10" json:"w10"`
	N10                                  string `form:"n10" json:"n10"`
	Bblahir                              string `form:"bblahir" json:"bblahir"`
	PanjangBadan                         string `form:"panjang_badan" json:"panjang_badan"`
	LingkarKepala                        string `form:"lingkar_kepala" json:"lingkar_kepala"`
	LingkarDada                          string `form:"lingkar_dada" json:"lingkar_dada"`
	ResusitasiSaatLahir                  string `form:"resusitasi_saat_lahir" json:"resusitasi_saat_lahir"`
	KeteranganResusitasiSaatLahir        string `form:"keterangan_resusitasi_saat_lahir" json:"keterangan_resusitasi_saat_lahir"`
	ObatDiberikanSaatLahir               string `form:"obat_diberikan_saat_lahir" json:"obat_diberikan_saat_lahir"`
	KeteranganLainnyaKeadaanBayi         string `form:"keterangan_lainnya_keadaan_bayi" json:"keterangan_lainnya_keadaan_bayi"`
	KondisiUmum                          string `form:"kondisi_umum" json:"kondisi_umum"`
	KeteranganKondisiUmum                string `form:"keterangan_kondisi_umum" json:"keterangan_kondisi_umum"`
	Kulit                                string `form:"kulit" json:"kulit"`
	KeteranganKulit                      string `form:"keterangan_kulit" json:"keterangan_kulit"`
	Kepala                               string `form:"kepala" json:"kepala"`
	KeteranganKepala                     string `form:"keterangan_kepala" json:"keterangan_kepala"`
	Leher                                string `form:"leher" json:"leher"`
	KeteranganLeher                      string `form:"keterangan_leher" json:"keterangan_leher"`
	Mata                                 string `form:"mata" json:"mata"`
	KeteranganMata                       string `form:"keterangan_mata" json:"keterangan_mata"`
	Hidung                               string `form:"hidung" json:"hidung"`
	KeteranganHidung                     string `form:"keterangan_hidung" json:"keterangan_hidung"`
	Telinga                              string `form:"telinga" json:"telinga"`
	KeteranganTelinga                    string `form:"keterangan_telinga" json:"keterangan_telinga"`
	Dada                                 string `form:"dada" json:"dada"`
	KeteranganDada                       string `form:"keterangan_dada" json:"keterangan_dada"`
	Paru                                 string `form:"paru" json:"paru"`
	KeteranganParu                       string `form:"keterangan_paru" json:"keterangan_paru"`
	Jantung                              string `form:"jantung" json:"jantung"`
	KeteranganJantung                    string `form:"keterangan_jantung" json:"keterangan_jantung"`
	Perut                                string `form:"perut" json:"perut"`
	KeteranganPerut                      string `form:"keterangan_perut" json:"keterangan_perut"`
	TaliPusat                            string `form:"tali_pusat" json:"tali_pusat"`
	KeteranganTaliPusat                  string `form:"keterangan_tali_pusat" json:"keterangan_tali_pusat"`
	AlatKelamin                          string `form:"alat_kelamin" json:"alat_kelamin"`
	KeteranganAlatKelamin                string `form:"keterangan_alat_kelamin" json:"keterangan_alat_kelamin"`
	RuasTulangBelakang                   string `form:"ruas_tulang_belakang" json:"ruas_tulang_belakang"`
	KeteranganRuasTulangBelakang         string `form:"keterangan_ruas_tulang_belakang" json:"keterangan_ruas_tulang_belakang"`
	Extrimitas                           string `form:"extrimitas" json:"extrimitas"`
	KeteranganExtrimitas                 string `form:"keterangan_extrimitas" json:"keterangan_extrimitas"`
	Anus                                 string `form:"anus" json:"anus"`
	KeteranganAnus                       string `form:"keterangan_anus" json:"keterangan_anus"`
	Refleks                              string `form:"refleks" json:"refleks"`
	KeteranganRefleks                    string `form:"keterangan_refleks" json:"keterangan_refleks"`
	DenyutFemoral                        string `form:"denyut_femoral" json:"denyut_femoral"`
	KeteranganDenyutFemoral              string `form:"keterangan_denyut_femoral" json:"keterangan_denyut_femoral"`
	PemeriksaanFisikLainnya              string `form:"pemeriksaan_fisik_lainnya" json:"pemeriksaan_fisik_lainnya"`
	PemeriksaanPenunjang                 string `form:"pemeriksaan_penunjang" json:"pemeriksaan_penunjang"`
	Diagnosa                             string `form:"diagnosa" json:"diagnosa"`
	Tatalaksana                          string `form:"tatalaksana" json:"tatalaksana"`
}

func penilaianBayiBaruLahirRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                 "required|date",
		"kd_dokter":                               "required|string|max_len:20",
		"no_rkm_medis_ibu":                        "string|max_len:15",
		"penyakit_diderita_ibu":                   "in:Tidak Ada,Ada",
		"keterangan_penyakit_diderita_ibu":        "string|max_len:70",
		"obat_dikonsumsi_selama_kehamilan":        "string|max_len:150",
		"perawatan_antenatal":                     "in:Ya,Tidak",
		"keterangan_perawatan_antenatal":          "string|max_len:40",
		"terdaftar_ekohort":                       "in:Ya,Tidak",
		"keterangan_terdaftar_ekohort":            "string|max_len:40",
		"penyulit_kehamilan":                      "in:Tidak Ada,Hiperemesis,CPD,Kelainan Letak,Solutio Placenta,Placenta Previa,KPD,Oligohydramnion,Polyhydramnion,Prolaps Tali Pusat,IUGR,Pre-eklamsi,Lainnya",
		"keterangan_penyulit_kehamilan":           "string|max_len:60",
		"alergi":                                  "string|max_len:60",
		"keterangan_lainnya_riwayat_maternal":     "string|max_len:150",
		"umur_kehamilan":                          "string|max_len:30",
		"kehamilan":                               "in:Tunggal,Kembar",
		"keterangan_kehamilan":                    "string|max_len:30",
		"urutan_kehamilan":                        "string|max_len:4",
		"jam_ketuban_pecah":                       "string|max_len:4",
		"menit_ketuban_pecah":                     "string|max_len:4",
		"jumlah_air_ketuban":                      "string|max_len:20",
		"warna_air_ketuban":                       "string|max_len:20",
		"bau_air_ketuban":                         "string|max_len:20",
		"letak_bayi":                              "string|max_len:70",
		"macam_persalinan":                        "in:Spontan,Porceps,Vacum,Sectio Caesarea,Lainnya",
		"keterangan_macam_persalinan":             "string|max_len:40",
		"indikasi_persalinan_operatif":            "in:Tidak Ada,Gawat Janin,SC Sebelumnya,Kala II Memanjang,Komplikasi Tali Pusat,Malposisi,Gawat Ibu,CPD,Lainnya",
		"keterangan_indikasi_persalinan_operatif": "string|max_len:50",
		"lama_gawat_janin":                        "string|max_len:4",
		"obat_selama_persalinan":                  "string|max_len:150",
		"berat_placenta":                          "string|max_len:4",
		"kelainan_placenta":                       "string|max_len:70",
		"keterangan_lainnya_riwayat_persalinan":   "string|max_len:150",
		"f1":                                      "string|max_len:1",
		"u1":                                      "string|max_len:1",
		"t1":                                      "string|max_len:1",
		"r1":                                      "string|max_len:1",
		"w1":                                      "string|max_len:1",
		"n1":                                      "string|max_len:2",
		"f5":                                      "string|max_len:1",
		"u5":                                      "string|max_len:1",
		"t5":                                      "string|max_len:1",
		"r5":                                      "string|max_len:1",
		"w5":                                      "string|max_len:1",
		"n5":                                      "string|max_len:2",
		"f10":                                     "string|max_len:1",
		"u10":                                     "string|max_len:1",
		"t10":                                     "string|max_len:1",
		"r10":                                     "string|max_len:1",
		"w10":                                     "string|max_len:1",
		"n10":                                     "string|max_len:2",
		"bblahir":                                 "string|max_len:5",
		"panjang_badan":                           "string|max_len:5",
		"lingkar_kepala":                          "string|max_len:5",
		"lingkar_dada":                            "string|max_len:5",
		"resusitasi_saat_lahir":                   "in:Tidak,Rangsang Taktil,O2,Ventilasi Dengan Maskre,Ventilasi Dengan EET,Lainnya",
		"keterangan_resusitasi_saat_lahir":        "string|max_len:70",
		"obat_diberikan_saat_lahir":               "string|max_len:150",
		"keterangan_lainnya_keadaan_bayi":         "string|max_len:150",
		"kondisi_umum":                            "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_kondisi_umum":                 "string|max_len:40",
		"kulit":                                   "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_kulit":                        "string|max_len:40",
		"kepala":                                  "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_kepala":                       "string|max_len:40",
		"leher":                                   "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_leher":                        "string|max_len:40",
		"mata":                                    "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_mata":                         "string|max_len:40",
		"hidung":                                  "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_hidung":                       "string|max_len:40",
		"telinga":                                 "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_telinga":                      "string|max_len:40",
		"dada":                                    "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_dada":                         "string|max_len:40",
		"paru":                                    "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_paru":                         "string|max_len:40",
		"jantung":                                 "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_jantung":                      "string|max_len:40",
		"perut":                                   "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_perut":                        "string|max_len:40",
		"tali_pusat":                              "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_tali_pusat":                   "string|max_len:40",
		"alat_kelamin":                            "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_alat_kelamin":                 "string|max_len:40",
		"ruas_tulang_belakang":                    "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_ruas_tulang_belakang":         "string|max_len:40",
		"extrimitas":                              "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_extrimitas":                   "string|max_len:40",
		"anus":                                    "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_anus":                         "string|max_len:40",
		"refleks":                                 "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_refleks":                      "string|max_len:40",
		"denyut_femoral":                          "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_denyut_femoral":               "string|max_len:40",
		"pemeriksaan_fisik_lainnya":               "string|max_len:300",
		"pemeriksaan_penunjang":                   "string|max_len:500",
		"diagnosa":                                "string|max_len:300",
		"tatalaksana":                             "string|max_len:1000",
	}
	return rules
}

// PenilaianBayiBaruLahirStore simpan penilaian bayi baru lahir; kolom waktu kunci kosong = sekarang.
type PenilaianBayiBaruLahirStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianBayiBaruLahirData
}

func (r *PenilaianBayiBaruLahirStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianBayiBaruLahirStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianBayiBaruLahirRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianBayiBaruLahirStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianBayiBaruLahirStore) Payload() PenilaianBayiBaruLahirData {
	return r.PenilaianBayiBaruLahirData
}

func (r *PenilaianBayiBaruLahirStore) DetailValues() map[string][]string { return nil }

// PenilaianBayiBaruLahirUpdate ubah penilaian bayi baru lahir (PUT); kunci lewat query string.
type PenilaianBayiBaruLahirUpdate struct {
	PenilaianBayiBaruLahirData
}

func (r *PenilaianBayiBaruLahirUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianBayiBaruLahirUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianBayiBaruLahirRules()
}

func (r *PenilaianBayiBaruLahirUpdate) Payload() PenilaianBayiBaruLahirData {
	return r.PenilaianBayiBaruLahirData
}

func (r *PenilaianBayiBaruLahirUpdate) DetailValues() map[string][]string { return nil }
