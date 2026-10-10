package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenatalaksanaanTerapiOkupasiData isian penatalaksanaan terapi okupasi.
type PenatalaksanaanTerapiOkupasiData struct {
	Tanggal                  string `form:"tanggal" json:"tanggal"`
	Nip                      string `form:"nip" json:"nip"`
	KeluhanUtama             string `form:"keluhan_utama" json:"keluhan_utama"`
	Rpd                      string `form:"rpd" json:"rpd"`
	Rps                      string `form:"rps" json:"rps"`
	AnamnesaGeneral          string `form:"anamnesa_general" json:"anamnesa_general"`
	TandaVital               string `form:"tanda_vital" json:"tanda_vital"`
	PemeriksaanPenunjang     string `form:"pemeriksaan_penunjang" json:"pemeriksaan_penunjang"`
	Spesialisasi             string `form:"spesialisasi" json:"spesialisasi"`
	KeteranganSpesialisasi   string `form:"keterangan_spesialisasi" json:"keterangan_spesialisasi"`
	PemeriksaanOkupasiTerapi string `form:"pemeriksaan_okupasi_terapi" json:"pemeriksaan_okupasi_terapi"`
	Aset                     string `form:"aset" json:"aset"`
	Limitasi                 string `form:"limitasi" json:"limitasi"`
	DiagnosaTerapiOkupasi    string `form:"diagnosa_terapi_okupasi" json:"diagnosa_terapi_okupasi"`
	RencanaIntervensi        string `form:"rencana_intervensi" json:"rencana_intervensi"`
}

func penatalaksanaanTerapiOkupasiRules() map[string]any {
	rules := map[string]any{
		"tanggal":                    "required|date",
		"nip":                        "required|string|max_len:20",
		"keluhan_utama":              "string|max_len:400",
		"rpd":                        "string|max_len:300",
		"rps":                        "string|max_len:300",
		"anamnesa_general":           "string|max_len:300",
		"tanda_vital":                "string|max_len:150",
		"pemeriksaan_penunjang":      "string|max_len:200",
		"spesialisasi":               "required|in:Pediatri,Dewasa,Geriatri,Psikososial,-",
		"keterangan_spesialisasi":    "string|max_len:30",
		"pemeriksaan_okupasi_terapi": "string|max_len:200",
		"aset":                       "string|max_len:150",
		"limitasi":                   "string|max_len:150",
		"diagnosa_terapi_okupasi":    "string|max_len:100",
		"rencana_intervensi":         "string|max_len:400",
	}
	return rules
}

// PenatalaksanaanTerapiOkupasiStore simpan penatalaksanaan terapi okupasi; kolom waktu kunci kosong = sekarang.
type PenatalaksanaanTerapiOkupasiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenatalaksanaanTerapiOkupasiData
}

func (r *PenatalaksanaanTerapiOkupasiStore) Authorize(ctx http.Context) error { return nil }

func (r *PenatalaksanaanTerapiOkupasiStore) Rules(ctx http.Context) map[string]any {
	rules := penatalaksanaanTerapiOkupasiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenatalaksanaanTerapiOkupasiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenatalaksanaanTerapiOkupasiStore) Payload() PenatalaksanaanTerapiOkupasiData {
	return r.PenatalaksanaanTerapiOkupasiData
}

func (r *PenatalaksanaanTerapiOkupasiStore) DetailValues() map[string][]string { return nil }

// PenatalaksanaanTerapiOkupasiUpdate ubah penatalaksanaan terapi okupasi (PUT); kunci lewat query string.
type PenatalaksanaanTerapiOkupasiUpdate struct {
	PenatalaksanaanTerapiOkupasiData
}

func (r *PenatalaksanaanTerapiOkupasiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenatalaksanaanTerapiOkupasiUpdate) Rules(ctx http.Context) map[string]any {
	return penatalaksanaanTerapiOkupasiRules()
}

func (r *PenatalaksanaanTerapiOkupasiUpdate) Payload() PenatalaksanaanTerapiOkupasiData {
	return r.PenatalaksanaanTerapiOkupasiData
}

func (r *PenatalaksanaanTerapiOkupasiUpdate) DetailValues() map[string][]string { return nil }
