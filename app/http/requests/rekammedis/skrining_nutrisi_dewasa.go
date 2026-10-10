package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningNutrisiDewasaData isian skrining nutrisi dewasa.
type SkriningNutrisiDewasaData struct {
	Tanggal    string `form:"tanggal" json:"tanggal"`
	Td         string `form:"td" json:"td"`
	Hr         string `form:"hr" json:"hr"`
	Rr         string `form:"rr" json:"rr"`
	Suhu       string `form:"suhu" json:"suhu"`
	Bb         string `form:"bb" json:"bb"`
	Tbpb       string `form:"tbpb" json:"tbpb"`
	Spo2       string `form:"spo2" json:"spo2"`
	Alergi     string `form:"alergi" json:"alergi"`
	Sg1        string `form:"sg1" json:"sg1"`
	Nilai1     string `form:"nilai1" json:"nilai1"`
	Sg2        string `form:"sg2" json:"sg2"`
	Nilai2     string `form:"nilai2" json:"nilai2"`
	TotalHasil int    `form:"total_hasil" json:"total_hasil"`
	Nip        string `form:"nip" json:"nip"`
}

func skriningNutrisiDewasaRules() map[string]any {
	rules := map[string]any{
		"tanggal":     "required|date",
		"td":          "string|max_len:8",
		"hr":          "string|max_len:5",
		"rr":          "string|max_len:5",
		"suhu":        "string|max_len:5",
		"bb":          "string|max_len:5",
		"tbpb":        "string|max_len:5",
		"spo2":        "string|max_len:5",
		"alergi":      "string|max_len:100",
		"sg1":         "required|string",
		"nilai1":      "required|in:0,1,2,3,4",
		"sg2":         "required|in:Ya,Tidak",
		"nilai2":      "required|in:0,1",
		"total_hasil": "int",
		"nip":         "required|string|max_len:20",
	}
	return rules
}

// SkriningNutrisiDewasaStore simpan skrining nutrisi dewasa; kolom waktu kunci kosong = sekarang.
type SkriningNutrisiDewasaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningNutrisiDewasaData
}

func (r *SkriningNutrisiDewasaStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningNutrisiDewasaStore) Rules(ctx http.Context) map[string]any {
	rules := skriningNutrisiDewasaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningNutrisiDewasaStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningNutrisiDewasaStore) Payload() SkriningNutrisiDewasaData {
	return r.SkriningNutrisiDewasaData
}

func (r *SkriningNutrisiDewasaStore) DetailValues() map[string][]string { return nil }

// SkriningNutrisiDewasaUpdate ubah skrining nutrisi dewasa (PUT); kunci lewat query string.
type SkriningNutrisiDewasaUpdate struct {
	SkriningNutrisiDewasaData
}

func (r *SkriningNutrisiDewasaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningNutrisiDewasaUpdate) Rules(ctx http.Context) map[string]any {
	return skriningNutrisiDewasaRules()
}

func (r *SkriningNutrisiDewasaUpdate) Payload() SkriningNutrisiDewasaData {
	return r.SkriningNutrisiDewasaData
}

func (r *SkriningNutrisiDewasaUpdate) DetailValues() map[string][]string { return nil }
