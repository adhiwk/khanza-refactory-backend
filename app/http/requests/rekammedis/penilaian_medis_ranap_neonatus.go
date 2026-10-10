package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRanapNeonatusData isian penilaian awal medis ranap neonatus.
type PenilaianMedisRanapNeonatusData struct {
	Tanggal                        string `form:"tanggal" json:"tanggal"`
	KdDokter                       string `form:"kd_dokter" json:"kd_dokter"`
	NoRkmMedisIbu                  string `form:"no_rkm_medis_ibu" json:"no_rkm_medis_ibu"`
	G                              string `form:"g" json:"g"`
	P                              string `form:"p" json:"p"`
	A                              string `form:"a" json:"a"`
	Hidup                          string `form:"hidup" json:"hidup"`
	Usiahamil                      string `form:"usiahamil" json:"usiahamil"`
	Hbsag                          string `form:"hbsag" json:"hbsag"`
	Hiv                            string `form:"hiv" json:"hiv"`
	Syphilis                       string `form:"syphilis" json:"syphilis"`
	RiwayatObstetriIbu             string `form:"riwayat_obstetri_ibu" json:"riwayat_obstetri_ibu"`
	KeteranganRiwayatObstetriIbu   string `form:"keterangan_riwayat_obstetri_ibu" json:"keterangan_riwayat_obstetri_ibu"`
	FaktorRisikoNeonatal           string `form:"faktor_risiko_neonatal" json:"faktor_risiko_neonatal"`
	KeteranganFaktorRisikoNeonatal string `form:"keterangan_faktor_risiko_neonatal" json:"keterangan_faktor_risiko_neonatal"`
	TanggalPersalinan              string `form:"tanggal_persalinan" json:"tanggal_persalinan"`
	BersalinDi                     string `form:"bersalin_di" json:"bersalin_di"`
	InisiasiMenyusui               string `form:"inisiasi_menyusui" json:"inisiasi_menyusui"`
	JenisPersalinan                string `form:"jenis_persalinan" json:"jenis_persalinan"`
	Indikasi                       string `form:"indikasi" json:"indikasi"`
	Aterm                          string `form:"aterm" json:"aterm"`
	Bernafas                       string `form:"bernafas" json:"bernafas"`
	TanusOtot                      string `form:"tanus_otot" json:"tanus_otot"`
	CairanAmnion                   string `form:"cairan_amnion" json:"cairan_amnion"`
	F1                             string `form:"f1" json:"f1"`
	U1                             string `form:"u1" json:"u1"`
	T1                             string `form:"t1" json:"t1"`
	R1                             string `form:"r1" json:"r1"`
	W1                             string `form:"w1" json:"w1"`
	N1                             string `form:"n1" json:"n1"`
	F5                             string `form:"f5" json:"f5"`
	U5                             string `form:"u5" json:"u5"`
	T5                             string `form:"t5" json:"t5"`
	R5                             string `form:"r5" json:"r5"`
	W5                             string `form:"w5" json:"w5"`
	N5                             string `form:"n5" json:"n5"`
	F10                            string `form:"f10" json:"f10"`
	U10                            string `form:"u10" json:"u10"`
	T10                            string `form:"t10" json:"t10"`
	R10                            string `form:"r10" json:"r10"`
	W10                            string `form:"w10" json:"w10"`
	N10                            string `form:"n10" json:"n10"`
	FrekuensiNapas                 string `form:"frekuensi_napas" json:"frekuensi_napas"`
	NilaiFrekuensiNapas            *int   `form:"nilai_frekuensi_napas" json:"nilai_frekuensi_napas"`
	Retraksi                       string `form:"retraksi" json:"retraksi"`
	NilaiRetraksi                  *int   `form:"nilai_retraksi" json:"nilai_retraksi"`
	Sianosis                       string `form:"sianosis" json:"sianosis"`
	NilaiSianosis                  *int   `form:"nilai_sianosis" json:"nilai_sianosis"`
	JalanMasukUdara                string `form:"jalan_masuk_udara" json:"jalan_masuk_udara"`
	NilaiJalanMasukUdara           *int   `form:"nilai_jalan_masuk_udara" json:"nilai_jalan_masuk_udara"`
	Grunting                       string `form:"grunting" json:"grunting"`
	NilaiGrunting                  *int   `form:"nilai_grunting" json:"nilai_grunting"`
	TotalDownScore                 *int   `form:"total_down_score" json:"total_down_score"`
	KeteranganDownScore            string `form:"keterangan_down_Score" json:"keterangan_down_Score"`
	Nadi                           string `form:"nadi" json:"nadi"`
	Rr                             string `form:"rr" json:"rr"`
	Suhu                           string `form:"suhu" json:"suhu"`
	Saturasi                       string `form:"saturasi" json:"saturasi"`
	Bb                             string `form:"bb" json:"bb"`
	Pb                             string `form:"pb" json:"pb"`
	Lk                             string `form:"lk" json:"lk"`
	Ld                             string `form:"ld" json:"ld"`
	KeadaanUmum                    string `form:"keadaan_umum" json:"keadaan_umum"`
	KeteranganKeadaanUmum          string `form:"keterangan_keadaan_umum" json:"keterangan_keadaan_umum"`
	Kulit                          string `form:"kulit" json:"kulit"`
	KeteranganKulit                string `form:"keterangan_kulit" json:"keterangan_kulit"`
	Kepala                         string `form:"kepala" json:"kepala"`
	KeteranganKepala               string `form:"keterangan_kepala" json:"keterangan_kepala"`
	Mata                           string `form:"mata" json:"mata"`
	KeteranganMata                 string `form:"keterangan_mata" json:"keterangan_mata"`
	Telinga                        string `form:"telinga" json:"telinga"`
	KeteranganTelinga              string `form:"keterangan_telinga" json:"keterangan_telinga"`
	Hidung                         string `form:"hidung" json:"hidung"`
	KeteranganHidung               string `form:"keterangan_hidung" json:"keterangan_hidung"`
	Mulut                          string `form:"mulut" json:"mulut"`
	KeteranganMulut                string `form:"keterangan_mulut" json:"keterangan_mulut"`
	Tenggorokan                    string `form:"tenggorokan" json:"tenggorokan"`
	KeteranganTenggorokan          string `form:"keterangan_tenggorokan" json:"keterangan_tenggorokan"`
	Leher                          string `form:"leher" json:"leher"`
	KeteranganLeher                string `form:"keterangan_leher" json:"keterangan_leher"`
	Thorax                         string `form:"thorax" json:"thorax"`
	KeteranganThorax               string `form:"keterangan_thorax" json:"keterangan_thorax"`
	Abdomen                        string `form:"abdomen" json:"abdomen"`
	KeteranganAbdomen              string `form:"keterangan_abdomen" json:"keterangan_abdomen"`
	Genitalia                      string `form:"genitalia" json:"genitalia"`
	KeteranganGenitalia            string `form:"keterangan_genitalia" json:"keterangan_genitalia"`
	Anus                           string `form:"anus" json:"anus"`
	KeteranganAnus                 string `form:"keterangan_anus" json:"keterangan_anus"`
	Muskulos                       string `form:"muskulos" json:"muskulos"`
	KeteranganMuskulos             string `form:"keterangan_muskulos" json:"keterangan_muskulos"`
	Ekstrimitas                    string `form:"ekstrimitas" json:"ekstrimitas"`
	KeteranganEkstrimitas          string `form:"keterangan_ekstrimitas" json:"keterangan_ekstrimitas"`
	Paru                           string `form:"paru" json:"paru"`
	KeteranganParu                 string `form:"keterangan_paru" json:"keterangan_paru"`
	Refleks                        string `form:"refleks" json:"refleks"`
	KeteranganRefleks              string `form:"keterangan_refleks" json:"keterangan_refleks"`
	KelainanLainnya                string `form:"kelainan_lainnya" json:"kelainan_lainnya"`
	PemeriksaanRegional            string `form:"pemeriksaan_regional" json:"pemeriksaan_regional"`
	Lab                            string `form:"lab" json:"lab"`
	Radiologi                      string `form:"radiologi" json:"radiologi"`
	Penunjanglainnya               string `form:"penunjanglainnya" json:"penunjanglainnya"`
	Diagnosis                      string `form:"diagnosis" json:"diagnosis"`
	Tata                           string `form:"tata" json:"tata"`
	Edukasi                        string `form:"edukasi" json:"edukasi"`
}

func penilaianMedisRanapNeonatusRules() map[string]any {
	rules := map[string]any{
		"tanggal":                           "required|date",
		"kd_dokter":                         "required|string|max_len:20",
		"no_rkm_medis_ibu":                  "string|max_len:15",
		"g":                                 "string|max_len:10",
		"p":                                 "string|max_len:10",
		"a":                                 "string|max_len:10",
		"hidup":                             "string|max_len:10",
		"usiahamil":                         "string|max_len:10",
		"hbsag":                             "in:Negatif (-),Positif (+),Tidak Ada Keterangan",
		"hiv":                               "in:Negatif (-),Positif (+),Tidak Ada Keterangan",
		"syphilis":                          "in:Negatif (-),Positif (+),Tidak Ada Keterangan",
		"riwayat_obstetri_ibu":              "in:Tidak Ada,Demam Pada Ibu > 38Â°C,Ketuban Pecah Dini,Ada Ketuban Berbau/Keruh,Nyeri Berkemih/ISK,Ibu DM,Ibu Hipertensi,Ibu Perdarahan,Ibu Eklamsia/Pre Eklamsia,Lainnya",
		"keterangan_riwayat_obstetri_ibu":   "string|max_len:70",
		"faktor_risiko_neonatal":            "in:Tidak Ada,Kelahiran Preterm,Kelahiran Post Date,DJJ Abnormal,Lainnya",
		"keterangan_faktor_risiko_neonatal": "string|max_len:70",
		"tanggal_persalinan":                "required|date",
		"bersalin_di":                       "string|max_len:70",
		"inisiasi_menyusui":                 "in:Ya,Tidak",
		"jenis_persalinan":                  "in:Spontan/Normal,Induksi,Sectio Caesaria,Vacum Ekstraksi",
		"indikasi":                          "string|max_len:70",
		"aterm":                             "in:Ya,Tidak",
		"bernafas":                          "in:Ya,Tidak",
		"tanus_otot":                        "in:Ya,Tidak",
		"cairan_amnion":                     "in:Ya,Tidak",
		"f1":                                "string|max_len:1",
		"u1":                                "string|max_len:1",
		"t1":                                "string|max_len:1",
		"r1":                                "string|max_len:1",
		"w1":                                "string|max_len:1",
		"n1":                                "string|max_len:2",
		"f5":                                "string|max_len:1",
		"u5":                                "string|max_len:1",
		"t5":                                "string|max_len:1",
		"r5":                                "string|max_len:1",
		"w5":                                "string|max_len:1",
		"n5":                                "string|max_len:2",
		"f10":                               "string|max_len:1",
		"u10":                               "string|max_len:1",
		"t10":                               "string|max_len:1",
		"r10":                               "string|max_len:1",
		"w10":                               "string|max_len:1",
		"n10":                               "string|max_len:2",
		"frekuensi_napas":                   "in:< 60,60 - 80,> 80",
		"nilai_frekuensi_napas":             "int",
		"retraksi":                          "in:Tidak Ada,Retraksi Ringan,Retraksi Berat",
		"nilai_retraksi":                    "int",
		"sianosis":                          "in:Tidak Ada,Hilang Dengan O2,Tidak Hilang Dengan O2",
		"nilai_sianosis":                    "int",
		"jalan_masuk_udara":                 "in:Baik,Penurunan Ringan Udara Masuk,Tidak Ada Udara Masuk",
		"nilai_jalan_masuk_udara":           "int",
		"grunting":                          "in:Tidak Ada,Dapat Didengar Dengan Stetoskop,Dapat Didengar Tanpa Stetoskop",
		"nilai_grunting":                    "int",
		"total_down_score":                  "int",
		"keterangan_down_Score":             "string|max_len:40",
		"nadi":                              "string|max_len:5",
		"rr":                                "string|max_len:5",
		"suhu":                              "string|max_len:5",
		"saturasi":                          "string|max_len:5",
		"bb":                                "string|max_len:5",
		"pb":                                "string|max_len:5",
		"lk":                                "string|max_len:5",
		"ld":                                "string|max_len:5",
		"keadaan_umum":                      "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_keadaan_umum":           "string|max_len:50",
		"kulit":                             "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_kulit":                  "string|max_len:50",
		"kepala":                            "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_kepala":                 "string|max_len:50",
		"mata":                              "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_mata":                   "string|max_len:50",
		"telinga":                           "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_telinga":                "string|max_len:50",
		"hidung":                            "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_hidung":                 "string|max_len:50",
		"mulut":                             "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_mulut":                  "string|max_len:50",
		"tenggorokan":                       "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_tenggorokan":            "string|max_len:50",
		"leher":                             "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_leher":                  "string|max_len:50",
		"thorax":                            "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_thorax":                 "string|max_len:50",
		"abdomen":                           "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_abdomen":                "string|max_len:50",
		"genitalia":                         "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_genitalia":              "string|max_len:50",
		"anus":                              "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_anus":                   "string|max_len:50",
		"muskulos":                          "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_muskulos":               "string|max_len:50",
		"ekstrimitas":                       "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_ekstrimitas":            "string|max_len:50",
		"paru":                              "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_paru":                   "string|max_len:50",
		"refleks":                           "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_refleks":                "string|max_len:50",
		"kelainan_lainnya":                  "string|max_len:80",
		"pemeriksaan_regional":              "string|max_len:500",
		"lab":                               "string|max_len:500",
		"radiologi":                         "string|max_len:500",
		"penunjanglainnya":                  "string|max_len:500",
		"diagnosis":                         "string|max_len:500",
		"tata":                              "string|max_len:2000",
		"edukasi":                           "string|max_len:1000",
	}
	return rules
}

// PenilaianMedisRanapNeonatusStore simpan penilaian awal medis ranap neonatus; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRanapNeonatusStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRanapNeonatusData
}

func (r *PenilaianMedisRanapNeonatusStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRanapNeonatusStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRanapNeonatusRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRanapNeonatusStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianMedisRanapNeonatusStore) Payload() PenilaianMedisRanapNeonatusData {
	return r.PenilaianMedisRanapNeonatusData
}

func (r *PenilaianMedisRanapNeonatusStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisRanapNeonatusUpdate ubah penilaian awal medis ranap neonatus (PUT); kunci lewat query string.
type PenilaianMedisRanapNeonatusUpdate struct {
	PenilaianMedisRanapNeonatusData
}

func (r *PenilaianMedisRanapNeonatusUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRanapNeonatusUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRanapNeonatusRules()
}

func (r *PenilaianMedisRanapNeonatusUpdate) Payload() PenilaianMedisRanapNeonatusData {
	return r.PenilaianMedisRanapNeonatusData
}

func (r *PenilaianMedisRanapNeonatusUpdate) DetailValues() map[string][]string { return nil }
