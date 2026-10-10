package pemeriksaan

import (
	"github.com/goravel/framework/contracts/http"
)

// Data isian SOAP & tanda vital; lingkar_perut hanya disimpan untuk rawat jalan.
type Data struct {
	SuhuTubuh    string `form:"suhu_tubuh" json:"suhu_tubuh"`
	Tensi        string `form:"tensi" json:"tensi"`
	Nadi         string `form:"nadi" json:"nadi"`
	Respirasi    string `form:"respirasi" json:"respirasi"`
	Tinggi       string `form:"tinggi" json:"tinggi"`
	Berat        string `form:"berat" json:"berat"`
	Spo2         string `form:"spo2" json:"spo2"`
	Gcs          string `form:"gcs" json:"gcs"`
	Kesadaran    string `form:"kesadaran" json:"kesadaran"`
	Keluhan      string `form:"keluhan" json:"keluhan"`
	Pemeriksaan  string `form:"pemeriksaan" json:"pemeriksaan"`
	Alergi       string `form:"alergi" json:"alergi"`
	LingkarPerut string `form:"lingkar_perut" json:"lingkar_perut"`
	Rtl          string `form:"rtl" json:"rtl"`
	Penilaian    string `form:"penilaian" json:"penilaian"`
	Instruksi    string `form:"instruksi" json:"instruksi"`
	Evaluasi     string `form:"evaluasi" json:"evaluasi"`
	Nip          string `form:"nip" json:"nip"`
}

func baseRules() map[string]any {
	return map[string]any{
		"suhu_tubuh":    "string|max_len:5",
		"tensi":         "string|max_len:8",
		"nadi":          "string|max_len:3",
		"respirasi":     "string|max_len:3",
		"tinggi":        "string|max_len:5",
		"berat":         "string|max_len:5",
		"spo2":          "string|max_len:3",
		"gcs":           "string|max_len:10",
		"kesadaran":     "in:Compos Mentis,Somnolence,Sopor,Coma,Alert,Confusion,Voice,Pain,Unresponsive,Apatis,Delirium,Meninggal",
		"keluhan":       "string|max_len:2000",
		"pemeriksaan":   "string|max_len:2000",
		"alergi":        "string|max_len:80",
		"lingkar_perut": "string|max_len:5",
		"rtl":           "string|max_len:2000",
		"penilaian":     "string|max_len:2000",
		"instruksi":     "string|max_len:2000",
		"evaluasi":      "string|max_len:2000",
		"nip":           "required|string|max_len:20",
	}
}

type StoreRequest struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"` // default hari ini
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`         // default jam sekarang
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

// UpdateRequest identitas baris lewat query ?no_rawat=&tgl_perawatan=&jam_rawat=.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
