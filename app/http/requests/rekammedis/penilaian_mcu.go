package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMcuData isian MCU.
type PenilaianMcuData struct {
	Tanggal                     string `form:"tanggal" json:"tanggal"`
	KdDokter                    string `form:"kd_dokter" json:"kd_dokter"`
	Informasi                   string `form:"informasi" json:"informasi"`
	Rps                         string `form:"rps" json:"rps"`
	Rpk                         string `form:"rpk" json:"rpk"`
	Rpd                         string `form:"rpd" json:"rpd"`
	Alergi                      string `form:"alergi" json:"alergi"`
	Keadaan                     string `form:"keadaan" json:"keadaan"`
	Kesadaran                   string `form:"kesadaran" json:"kesadaran"`
	Td                          string `form:"td" json:"td"`
	Nadi                        string `form:"nadi" json:"nadi"`
	Rr                          string `form:"rr" json:"rr"`
	Tb                          string `form:"tb" json:"tb"`
	Bb                          string `form:"bb" json:"bb"`
	Suhu                        string `form:"suhu" json:"suhu"`
	Bmi                         string `form:"bmi" json:"bmi"`
	KasifikasiBmi               string `form:"kasifikasi_bmi" json:"kasifikasi_bmi"`
	LingkarPinggang             string `form:"lingkar_pinggang" json:"lingkar_pinggang"`
	RisikoLingkarPinggang       string `form:"risiko_lingkar_pinggang" json:"risiko_lingkar_pinggang"`
	Submandibula                string `form:"submandibula" json:"submandibula"`
	Axilla                      string `form:"axilla" json:"axilla"`
	Supraklavikula              string `form:"supraklavikula" json:"supraklavikula"`
	Leher                       string `form:"leher" json:"leher"`
	Inguinal                    string `form:"inguinal" json:"inguinal"`
	Oedema                      string `form:"oedema" json:"oedema"`
	SinusFrontalis              string `form:"sinus_frontalis" json:"sinus_frontalis"`
	SinusMaxilaris              string `form:"sinus_maxilaris" json:"sinus_maxilaris"`
	Rambut                      string `form:"rambut" json:"rambut"`
	Palpebra                    string `form:"palpebra" json:"palpebra"`
	Sklera                      string `form:"sklera" json:"sklera"`
	Cornea                      string `form:"cornea" json:"cornea"`
	ButaWarna                   string `form:"buta_warna" json:"buta_warna"`
	Konjungtiva                 string `form:"konjungtiva" json:"konjungtiva"`
	Lensa                       string `form:"lensa" json:"lensa"`
	Pupil                       string `form:"pupil" json:"pupil"`
	MenggunakanKacamata         string `form:"menggunakan_kacamata" json:"menggunakan_kacamata"`
	Visus                       string `form:"visus" json:"visus"`
	LuasLapangPandang           string `form:"luas_lapang_pandang" json:"luas_lapang_pandang"`
	KeteranganLuasLapangPandang string `form:"keterangan_luas_lapang_pandang" json:"keterangan_luas_lapang_pandang"`
	LubangTelinga               string `form:"lubang_telinga" json:"lubang_telinga"`
	DaunTelinga                 string `form:"daun_telinga" json:"daun_telinga"`
	SelaputPendengaran          string `form:"selaput_pendengaran" json:"selaput_pendengaran"`
	ProcMastoideus              string `form:"proc_mastoideus" json:"proc_mastoideus"`
	SeptumNasi                  string `form:"septum_nasi" json:"septum_nasi"`
	LubangHidung                string `form:"lubang_hidung" json:"lubang_hidung"`
	Sinus                       string `form:"sinus" json:"sinus"`
	Bibir                       string `form:"bibir" json:"bibir"`
	Gusi                        string `form:"gusi" json:"gusi"`
	Gigi                        string `form:"gigi" json:"gigi"`
	Caries                      string `form:"caries" json:"caries"`
	Lidah                       string `form:"lidah" json:"lidah"`
	Faring                      string `form:"faring" json:"faring"`
	Tonsil                      string `form:"tonsil" json:"tonsil"`
	KelenjarLimfe               string `form:"kelenjar_limfe" json:"kelenjar_limfe"`
	KelenjarGondok              string `form:"kelenjar_gondok" json:"kelenjar_gondok"`
	GerakanDada                 string `form:"gerakan_dada" json:"gerakan_dada"`
	VocalFemitus                string `form:"vocal_femitus" json:"vocal_femitus"`
	PerkusiDada                 string `form:"perkusi_dada" json:"perkusi_dada"`
	BunyiNapas                  string `form:"bunyi_napas" json:"bunyi_napas"`
	BunyiTambahan               string `form:"bunyi_tambahan" json:"bunyi_tambahan"`
	IctusCordis                 string `form:"ictus_cordis" json:"ictus_cordis"`
	BunyiJantung                string `form:"bunyi_jantung" json:"bunyi_jantung"`
	Batas                       string `form:"batas" json:"batas"`
	Mamae                       string `form:"mamae" json:"mamae"`
	KeteranganMamae             string `form:"keterangan_mamae" json:"keterangan_mamae"`
	Inspeksi                    string `form:"inspeksi" json:"inspeksi"`
	Palpasi                     string `form:"palpasi" json:"palpasi"`
	Hepar                       string `form:"hepar" json:"hepar"`
	PerkusiAbdomen              string `form:"perkusi_abdomen" json:"perkusi_abdomen"`
	Auskultasi                  string `form:"auskultasi" json:"auskultasi"`
	Limpa                       string `form:"limpa" json:"limpa"`
	Costovertebral              string `form:"costovertebral" json:"costovertebral"`
	Scoliosis                   string `form:"scoliosis" json:"scoliosis"`
	KondisiKulit                string `form:"kondisi_kulit" json:"kondisi_kulit"`
	PenyakitKulit               string `form:"penyakit_kulit" json:"penyakit_kulit"`
	EkstrimitasAtas             string `form:"ekstrimitas_atas" json:"ekstrimitas_atas"`
	EkstrimitasAtasKet          string `form:"ekstrimitas_atas_ket" json:"ekstrimitas_atas_ket"`
	EkstrimitasBawah            string `form:"ekstrimitas_bawah" json:"ekstrimitas_bawah"`
	EkstrimitasBawahKet         string `form:"ekstrimitas_bawah_ket" json:"ekstrimitas_bawah_ket"`
	AreaGenitalia               string `form:"area_genitalia" json:"area_genitalia"`
	KeteranganAreaGenitalia     string `form:"keterangan_area_genitalia" json:"keterangan_area_genitalia"`
	AnusPerianal                string `form:"anus_perianal" json:"anus_perianal"`
	KeteranganAnusPerianal      string `form:"keterangan_anus_perianal" json:"keterangan_anus_perianal"`
	Laborat                     string `form:"laborat" json:"laborat"`
	Radiologi                   string `form:"radiologi" json:"radiologi"`
	Ekg                         string `form:"ekg" json:"ekg"`
	Spirometri                  string `form:"spirometri" json:"spirometri"`
	Audiometri                  string `form:"audiometri" json:"audiometri"`
	Treadmill                   string `form:"treadmill" json:"treadmill"`
	RombergTest                 string `form:"romberg_test" json:"romberg_test"`
	BackStrength                string `form:"back_strength" json:"back_strength"`
	AbiTanganKanan              string `form:"abi_tangan_kanan" json:"abi_tangan_kanan"`
	AbiTanganKiri               string `form:"abi_tangan_kiri" json:"abi_tangan_kiri"`
	AbiKakiKanan                string `form:"abi_kaki_kanan" json:"abi_kaki_kanan"`
	AbiKakiKiri                 string `form:"abi_kaki_kiri" json:"abi_kaki_kiri"`
	Lainlain                    string `form:"lainlain" json:"lainlain"`
	Merokok                     string `form:"merokok" json:"merokok"`
	Alkohol                     string `form:"alkohol" json:"alkohol"`
	Kesimpulan                  string `form:"kesimpulan" json:"kesimpulan"`
	Anjuran                     string `form:"anjuran" json:"anjuran"`
}

func penilaianMcuRules() map[string]any {
	rules := map[string]any{
		"tanggal":                        "required|date",
		"kd_dokter":                      "required|string|max_len:20",
		"informasi":                      "required|in:Autoanamnesis,Alloanamnesis",
		"rps":                            "string|max_len:2000",
		"rpk":                            "string|max_len:1000",
		"rpd":                            "string|max_len:1000",
		"alergi":                         "string|max_len:150",
		"keadaan":                        "required|in:Baik,Tidak Baik",
		"kesadaran":                      "required|in:Composmentis,Apatis,Somnolen",
		"td":                             "string|max_len:8",
		"nadi":                           "string|max_len:5",
		"rr":                             "string|max_len:5",
		"tb":                             "string|max_len:5",
		"bb":                             "string|max_len:5",
		"suhu":                           "string|max_len:5",
		"bmi":                            "string|max_len:6",
		"kasifikasi_bmi":                 "required|in:Berat Badan Kurang,Berat Badan Normal,Kelebihan Berat Badan,Obesitas I,Obesitas II",
		"lingkar_pinggang":               "string|max_len:6",
		"risiko_lingkar_pinggang":        "required|in:Rendah,Cukup,Meningkat,Moderat,Berat,Sangat",
		"submandibula":                   "required|in:Tidak Membesar,Membesar,-",
		"axilla":                         "required|in:Tidak Membesar,Membesar,-",
		"supraklavikula":                 "required|in:Tidak Membesar,Membesar,-",
		"leher":                          "required|in:Tidak Membesar,Membesar,-",
		"inguinal":                       "required|in:Tidak Membesar,Membesar,-",
		"oedema":                         "required|in:Tidak Ada,Ada,-",
		"sinus_frontalis":                "required|in:Tidak Ada,Ada,=",
		"sinus_maxilaris":                "required|in:Tidak Ada,Ada,-",
		"rambut":                         "string|max_len:100",
		"palpebra":                       "required|in:Normal,Oedem,Ptosis,-",
		"sklera":                         "required|in:Normal,Ikterik,-",
		"cornea":                         "required|in:Normal,Tidak Normal,-",
		"buta_warna":                     "required|in:Normal,Buta Warna Partial,Buta Warna Total,-",
		"konjungtiva":                    "required|in:Normal,Anemis,Hiperemis,-",
		"lensa":                          "required|in:Jernih,Keruh,Kacamata,-",
		"pupil":                          "required|in:Isokor,Anisokor,-",
		"menggunakan_kacamata":           "required|in:Tidak,Ya,-",
		"visus":                          "string|max_len:50",
		"luas_lapang_pandang":            "required|in:Normal,Tidak Normal,-",
		"keterangan_luas_lapang_pandang": "string|max_len:50",
		"lubang_telinga":                 "required|in:Normal,Tidak Normal,Lapang,Sempit,Serumen Prop,-",
		"daun_telinga":                   "required|in:Normal,Tidak Normal,-",
		"selaput_pendengaran":            "required|in:Intak,Tidak Intak,-",
		"proc_mastoideus":                "required|in:Normal,Tidak Normal,-",
		"septum_nasi":                    "required|in:Normal,Deviasi,-",
		"lubang_hidung":                  "required|in:Lapang,Rhinore,Epistaksis,-",
		"sinus":                          "required|in:Normal,Tidak Normal,-",
		"bibir":                          "required|in:Lembab,Kering",
		"gusi":                           "required|in:Normal,Tidak Normal,-",
		"gigi":                           "required|in:Normal,Tidak Normal,-",
		"caries":                         "required|in:Tidak Ada,Ada,-",
		"lidah":                          "required|in:Bersih,Kotor,Tremor,-",
		"faring":                         "required|in:Normal,Hiperemis,-",
		"tonsil":                         "required|in:T1-T1,T2-T2,T3-T3,T4-T4,T0-T0,-",
		"kelenjar_limfe":                 "required|in:Tidak Membesar,Membesar,-",
		"kelenjar_gondok":                "required|in:Tidak Membesar,Membesar,-",
		"gerakan_dada":                   "required|in:Simetris,Tidak Simetris,-",
		"vocal_femitus":                  "required|in:Sama,Tidak Sama,-",
		"perkusi_dada":                   "required|in:Sonor,Pekak,-",
		"bunyi_napas":                    "required|in:Vesikuler,Bronkhial,Trakeal,-",
		"bunyi_tambahan":                 "required|in:Tidak Ada,Wheezing,Ronchi,-",
		"ictus_cordis":                   "required|in:Tidak Terlihat,Terlihat, Teraba,Tidak Teraba,-",
		"bunyi_jantung":                  "required|in:Reguler,Irreguler,Gallop,Lain-lain,-",
		"batas":                          "required|in:Normal,Melebar,-",
		"mamae":                          "required|in:Normal,Tidak Normal,-",
		"keterangan_mamae":               "string|max_len:100",
		"inspeksi":                       "required|in:Datar,Cembung,-",
		"palpasi":                        "required|in:Supel,Tegang (Defans Muscular),Nyeri Tekan Epigastrium,Nyeri Tekan Suprapubik,Nyeri Tekan Right Lower Quadrant,Nyeri Tekan Left Lower Quadrant,-",
		"hepar":                          "required|in:Tidak Membesar,Membesar,-",
		"perkusi_abdomen":                "required|in:Timpani,Hipertimpani,Dull,-",
		"auskultasi":                     "required|in:Normal,Bising Usus Meningkat,Bising Usus Menurun,-",
		"limpa":                          "required|in:Tidak Membesar,Membesar,-",
		"costovertebral":                 "required|in:Tidak Ada,Ada Di Kiri,Ada Di Kanan,-",
		"scoliosis":                      "required|in:Tidak Ada,Ada,-",
		"kondisi_kulit":                  "required|in:Normal,Tato,Penyakit Kulit,-",
		"penyakit_kulit":                 "string|max_len:100",
		"ekstrimitas_atas":               "required|in:Normal,Tidak Normal,-",
		"ekstrimitas_atas_ket":           "string|max_len:50",
		"ekstrimitas_bawah":              "required|in:Normal,Tidak Normal,-",
		"ekstrimitas_bawah_ket":          "string|max_len:50",
		"area_genitalia":                 "required|in:Tidak Ada Kelainan,Ada Kelainan,-",
		"keterangan_area_genitalia":      "string|max_len:100",
		"anus_perianal":                  "required|in:Normal,Tidak Normal,-",
		"keterangan_anus_perianal":       "string|max_len:100",
		"laborat":                        "string",
		"radiologi":                      "string",
		"ekg":                            "string",
		"spirometri":                     "string",
		"audiometri":                     "string",
		"treadmill":                      "string",
		"romberg_test":                   "string",
		"back_strength":                  "string",
		"abi_tangan_kanan":               "string",
		"abi_tangan_kiri":                "string",
		"abi_kaki_kanan":                 "string",
		"abi_kaki_kiri":                  "string",
		"lainlain":                       "string",
		"merokok":                        "string|max_len:100",
		"alkohol":                        "string|max_len:100",
		"kesimpulan":                     "string",
		"anjuran":                        "string",
	}
	return rules
}

// PenilaianMcuStore simpan MCU; kolom waktu kunci kosong = sekarang.
type PenilaianMcuStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMcuData
}

func (r *PenilaianMcuStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMcuStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMcuRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMcuStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianMcuStore) Payload() PenilaianMcuData { return r.PenilaianMcuData }

func (r *PenilaianMcuStore) DetailValues() map[string][]string { return nil }

// PenilaianMcuUpdate ubah MCU (PUT); kunci lewat query string.
type PenilaianMcuUpdate struct {
	PenilaianMcuData
}

func (r *PenilaianMcuUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMcuUpdate) Rules(ctx http.Context) map[string]any { return penilaianMcuRules() }

func (r *PenilaianMcuUpdate) Payload() PenilaianMcuData { return r.PenilaianMcuData }

func (r *PenilaianMcuUpdate) DetailValues() map[string][]string { return nil }
