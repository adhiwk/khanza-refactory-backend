package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningThalassemiaData isian skrining talasemia.
type SkriningThalassemiaData struct {
	Tanggal                string `form:"tanggal" json:"tanggal"`
	Nip                    string `form:"nip" json:"nip"`
	TransfusiDarah         string `form:"transfusi_darah" json:"transfusi_darah"`
	RutinTransfusi         string `form:"rutin_transfusi" json:"rutin_transfusi"`
	SaudaraThalassemia     string `form:"saudara_thalassemia" json:"saudara_thalassemia"`
	TumbuhKembangTerlambat string `form:"tumbuh_kembang_terlambat" json:"tumbuh_kembang_terlambat"`
	Anemia                 string `form:"anemia" json:"anemia"`
	Ikterus                string `form:"ikterus" json:"ikterus"`
	PerutBuncit            string `form:"perut_buncit" json:"perut_buncit"`
	GiziKurang             string `form:"gizi_kurang" json:"gizi_kurang"`
	FaciesCooley           string `form:"facies_cooley" json:"facies_cooley"`
	PerawakanPendek        string `form:"perawakan_pendek" json:"perawakan_pendek"`
	HiperpigmentasiKulit   string `form:"hiperpigmentasi_kulit" json:"hiperpigmentasi_kulit"`
	Hemoglobin             string `form:"hemoglobin" json:"hemoglobin"`
	Mvc                    string `form:"mvc" json:"mvc"`
	Mchc                   string `form:"mchc" json:"mchc"`
	DarahTepi              string `form:"darah_tepi" json:"darah_tepi"`
	TindakLanjut           string `form:"tindak_lanjut" json:"tindak_lanjut"`
}

func skriningThalassemiaRules() map[string]any {
	rules := map[string]any{
		"tanggal":                  "required|date",
		"nip":                      "required|string|max_len:20",
		"transfusi_darah":          "in:Tidak,Ya",
		"rutin_transfusi":          "in:Tidak,Ya",
		"saudara_thalassemia":      "in:Tidak,Ya",
		"tumbuh_kembang_terlambat": "in:Tidak,Ya",
		"anemia":                   "in:Tidak,Ya",
		"ikterus":                  "in:Tidak,Ya",
		"perut_buncit":             "in:Tidak,Ya",
		"gizi_kurang":              "in:Tidak,Ya",
		"facies_cooley":            "in:Tidak,Ya",
		"perawakan_pendek":         "in:Tidak,Ya",
		"hiperpigmentasi_kulit":    "in:Tidak,Ya",
		"hemoglobin":               "in:Normal,< 11 mg/dl",
		"mvc":                      "in:Normal,< 80 fl",
		"mchc":                     "in:Normal,< 27 pq",
		"darah_tepi":               "in:Normal,< 100.000 mm3",
		"tindak_lanjut":            "string|max_len:300",
	}
	return rules
}

// SkriningThalassemiaStore simpan skrining talasemia; kolom waktu kunci kosong = sekarang.
type SkriningThalassemiaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningThalassemiaData
}

func (r *SkriningThalassemiaStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningThalassemiaStore) Rules(ctx http.Context) map[string]any {
	rules := skriningThalassemiaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningThalassemiaStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningThalassemiaStore) Payload() SkriningThalassemiaData {
	return r.SkriningThalassemiaData
}

func (r *SkriningThalassemiaStore) DetailValues() map[string][]string { return nil }

// SkriningThalassemiaUpdate ubah skrining talasemia (PUT); kunci lewat query string.
type SkriningThalassemiaUpdate struct {
	SkriningThalassemiaData
}

func (r *SkriningThalassemiaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningThalassemiaUpdate) Rules(ctx http.Context) map[string]any {
	return skriningThalassemiaRules()
}

func (r *SkriningThalassemiaUpdate) Payload() SkriningThalassemiaData {
	return r.SkriningThalassemiaData
}

func (r *SkriningThalassemiaUpdate) DetailValues() map[string][]string { return nil }
