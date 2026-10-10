package skriningralan

import (
	"github.com/goravel/framework/contracts/http"
)

type Data struct {
	Geriatri    string `form:"geriatri" json:"geriatri"`
	Kesadaran   string `form:"kesadaran" json:"kesadaran"`
	Pernapasan  string `form:"pernapasan" json:"pernapasan"`
	NyeriDada   string `form:"nyeri_dada" json:"nyeri_dada"`
	SkalaNyeri  string `form:"skala_nyeri" json:"skala_nyeri"`
	Batuk       string `form:"batuk" json:"batuk"`
	RisikoJatuh string `form:"risiko_jatuh" json:"risiko_jatuh"`
	Keputusan   string `form:"keputusan" json:"keputusan"`
	Nip         string `form:"nip" json:"nip"`
}

func baseRules() map[string]any {
	return map[string]any{
		"geriatri":     "required|in:Ya,Tidak",
		"kesadaran":    "required|in:Sadar penuh,Tampak mengantuk/gelisah bicara tidak jelas,Tidak sadar,Batuk > 2 minggu",
		"pernapasan":   "required|in:Nafas normal,Tampak sesak,Tidak bernafas",
		"nyeri_dada":   "required|in:Tidak ada,Ada (Tingkat sedang),Nyeri dada kiri tembus punggung",
		"skala_nyeri":  "required|in:Tidak sakit,Sedikit sakit,Agak mengganggu,Mengganggu aktivitas,Sangat mengganggu,Tak tertahankan",
		"batuk":        "required|in:Tidak,Ya < 2 minggu,Ya > 2 minggu",
		"risiko_jatuh": "required|in:Tidak,Ya",
		"keputusan":    "required|in:Sesuai antrian,IGD",
		"nip":          "required|string|max_len:20",
	}
}

// StoreRequest tanggal/jam kosong = sekarang.
type StoreRequest struct {
	NoRkmMedis string `form:"no_rkm_medis" json:"no_rkm_medis"`
	Tanggal    string `form:"tanggal" json:"tanggal"`
	Jam        string `form:"jam" json:"jam"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error { return nil }

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["no_rkm_medis"] = "required|string|max_len:15"
	rules["tanggal"] = "date"
	rules["jam"] = "string|len:8"
	return rules
}

// UpdateRequest kunci lewat query string ?tanggal=&jam=&no_rkm_medis=.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error { return nil }

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any { return baseRules() }
