package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianFisioterapiData isian penilaian fisioterapi.
type PenilaianFisioterapiData struct {
	Tanggal                    string `form:"tanggal" json:"tanggal"`
	Informasi                  string `form:"informasi" json:"informasi"`
	KeluhanUtama               string `form:"keluhan_utama" json:"keluhan_utama"`
	Rps                        string `form:"rps" json:"rps"`
	Rpd                        string `form:"rpd" json:"rpd"`
	Td                         string `form:"td" json:"td"`
	Hr                         string `form:"hr" json:"hr"`
	Rr                         string `form:"rr" json:"rr"`
	Suhu                       string `form:"suhu" json:"suhu"`
	NyeriTekan                 string `form:"nyeri_tekan" json:"nyeri_tekan"`
	NyeriGerak                 string `form:"nyeri_gerak" json:"nyeri_gerak"`
	NyeriDiam                  string `form:"nyeri_diam" json:"nyeri_diam"`
	Palpasi                    string `form:"palpasi" json:"palpasi"`
	LuasGerakSendi             string `form:"luas_gerak_sendi" json:"luas_gerak_sendi"`
	KekuatanOtot               string `form:"kekuatan_otot" json:"kekuatan_otot"`
	Statis                     string `form:"statis" json:"statis"`
	Dinamis                    string `form:"dinamis" json:"dinamis"`
	Kognitif                   string `form:"kognitif" json:"kognitif"`
	Auskultasi                 string `form:"auskultasi" json:"auskultasi"`
	AlatBantu                  string `form:"alat_bantu" json:"alat_bantu"`
	KetBantu                   string `form:"ket_bantu" json:"ket_bantu"`
	Prothesa                   string `form:"prothesa" json:"prothesa"`
	KetPro                     string `form:"ket_pro" json:"ket_pro"`
	Deformitas                 string `form:"deformitas" json:"deformitas"`
	KetDeformitas              string `form:"ket_deformitas" json:"ket_deformitas"`
	Resikojatuh                string `form:"resikojatuh" json:"resikojatuh"`
	KetResikojatuh             string `form:"ket_resikojatuh" json:"ket_resikojatuh"`
	Adl                        string `form:"adl" json:"adl"`
	LainlainFungsional         string `form:"lainlain_fungsional" json:"lainlain_fungsional"`
	KetFisik                   string `form:"ket_fisik" json:"ket_fisik"`
	PemeriksaanMusculoskeletal string `form:"pemeriksaan_musculoskeletal" json:"pemeriksaan_musculoskeletal"`
	PemeriksaanNeuromuscular   string `form:"pemeriksaan_neuromuscular" json:"pemeriksaan_neuromuscular"`
	PemeriksaanCardiopulmonal  string `form:"pemeriksaan_cardiopulmonal" json:"pemeriksaan_cardiopulmonal"`
	PemeriksaanIntegument      string `form:"pemeriksaan_integument" json:"pemeriksaan_integument"`
	PengukuranMusculoskeletal  string `form:"pengukuran_musculoskeletal" json:"pengukuran_musculoskeletal"`
	PengukuranNeuromuscular    string `form:"pengukuran_neuromuscular" json:"pengukuran_neuromuscular"`
	PengukuranCardiopulmonal   string `form:"pengukuran_cardiopulmonal" json:"pengukuran_cardiopulmonal"`
	PengukuranIntegument       string `form:"pengukuran_integument" json:"pengukuran_integument"`
	Penunjang                  string `form:"penunjang" json:"penunjang"`
	DiagnosisFisio             string `form:"diagnosis_fisio" json:"diagnosis_fisio"`
	RencanaTerapi              string `form:"rencana_terapi" json:"rencana_terapi"`
	Nip                        string `form:"nip" json:"nip"`
}

func penilaianFisioterapiRules() map[string]any {
	rules := map[string]any{
		"tanggal":                     "required|date",
		"informasi":                   "required|in:Autoanamnesis,Alloanamnesis",
		"keluhan_utama":               "string|max_len:150",
		"rps":                         "string|max_len:100",
		"rpd":                         "string|max_len:100",
		"td":                          "string|max_len:8",
		"hr":                          "string|max_len:5",
		"rr":                          "string|max_len:5",
		"suhu":                        "string|max_len:5",
		"nyeri_tekan":                 "string|max_len:5",
		"nyeri_gerak":                 "string|max_len:5",
		"nyeri_diam":                  "string|max_len:5",
		"palpasi":                     "string|max_len:50",
		"luas_gerak_sendi":            "string|max_len:50",
		"kekuatan_otot":               "string|max_len:50",
		"statis":                      "string|max_len:50",
		"dinamis":                     "string|max_len:50",
		"kognitif":                    "string|max_len:50",
		"auskultasi":                  "string|max_len:50",
		"alat_bantu":                  "required|in:Tidak,Ya",
		"ket_bantu":                   "string|max_len:50",
		"prothesa":                    "required|in:Tidak,Ya",
		"ket_pro":                     "string|max_len:50",
		"deformitas":                  "required|in:Tidak,Ya",
		"ket_deformitas":              "string|max_len:50",
		"resikojatuh":                 "required|in:Tidak,Ya",
		"ket_resikojatuh":             "string|max_len:50",
		"adl":                         "required|in:Mandiri,Dibantu",
		"lainlain_fungsional":         "string|max_len:70",
		"ket_fisik":                   "string",
		"pemeriksaan_musculoskeletal": "string|max_len:200",
		"pemeriksaan_neuromuscular":   "string|max_len:200",
		"pemeriksaan_cardiopulmonal":  "string|max_len:200",
		"pemeriksaan_integument":      "string|max_len:200",
		"pengukuran_musculoskeletal":  "string|max_len:200",
		"pengukuran_neuromuscular":    "string|max_len:200",
		"pengukuran_cardiopulmonal":   "string|max_len:200",
		"pengukuran_integument":       "string|max_len:200",
		"penunjang":                   "string|max_len:500",
		"diagnosis_fisio":             "string|max_len:100",
		"rencana_terapi":              "string|max_len:200",
		"nip":                         "required|string|max_len:20",
	}
	return rules
}

// PenilaianFisioterapiStore simpan penilaian fisioterapi; kolom waktu kunci kosong = sekarang.
type PenilaianFisioterapiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianFisioterapiData
}

func (r *PenilaianFisioterapiStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianFisioterapiStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianFisioterapiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianFisioterapiStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianFisioterapiStore) Payload() PenilaianFisioterapiData {
	return r.PenilaianFisioterapiData
}

func (r *PenilaianFisioterapiStore) DetailValues() map[string][]string { return nil }

// PenilaianFisioterapiUpdate ubah penilaian fisioterapi (PUT); kunci lewat query string.
type PenilaianFisioterapiUpdate struct {
	PenilaianFisioterapiData
}

func (r *PenilaianFisioterapiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianFisioterapiUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianFisioterapiRules()
}

func (r *PenilaianFisioterapiUpdate) Payload() PenilaianFisioterapiData {
	return r.PenilaianFisioterapiData
}

func (r *PenilaianFisioterapiUpdate) DetailValues() map[string][]string { return nil }
