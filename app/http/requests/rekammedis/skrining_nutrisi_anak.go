package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningNutrisiAnakData isian skrining nutrisi anak.
type SkriningNutrisiAnakData struct {
	Tanggal                      string `form:"tanggal" json:"tanggal"`
	Td                           string `form:"td" json:"td"`
	Hr                           string `form:"hr" json:"hr"`
	Rr                           string `form:"rr" json:"rr"`
	Suhu                         string `form:"suhu" json:"suhu"`
	Bb                           string `form:"bb" json:"bb"`
	Tbpb                         string `form:"tbpb" json:"tbpb"`
	Spo2                         string `form:"spo2" json:"spo2"`
	Alergi                       string `form:"alergi" json:"alergi"`
	Sg1                          string `form:"sg1" json:"sg1"`
	Nilai1                       string `form:"nilai1" json:"nilai1"`
	Sg2                          string `form:"sg2" json:"sg2"`
	Nilai2                       string `form:"nilai2" json:"nilai2"`
	Sg3                          string `form:"sg3" json:"sg3"`
	Nilai3                       string `form:"nilai3" json:"nilai3"`
	Sg4                          string `form:"sg4" json:"sg4"`
	Nilai4                       string `form:"nilai4" json:"nilai4"`
	TotalHasil                   int    `form:"total_hasil" json:"total_hasil"`
	SkorNutrisi                  string `form:"skor_nutrisi" json:"skor_nutrisi"`
	DiketahuiDietisien           string `form:"diketahui_dietisien" json:"diketahui_dietisien"`
	KeteranganDiketahuiDietisien string `form:"keterangan_diketahui_dietisien" json:"keterangan_diketahui_dietisien"`
	Nip                          string `form:"nip" json:"nip"`
}

func skriningNutrisiAnakRules() map[string]any {
	rules := map[string]any{
		"tanggal":                        "required|date",
		"td":                             "string|max_len:8",
		"hr":                             "string|max_len:5",
		"rr":                             "string|max_len:5",
		"suhu":                           "string|max_len:5",
		"bb":                             "string|max_len:5",
		"tbpb":                           "string|max_len:5",
		"spo2":                           "string|max_len:5",
		"alergi":                         "string|max_len:100",
		"sg1":                            "required|in:Tidak,Ya",
		"nilai1":                         "required|in:0,1",
		"sg2":                            "required|in:Tidak,Ya",
		"nilai2":                         "required|in:0,1",
		"sg3":                            "required|in:Tidak,Ya",
		"nilai3":                         "required|in:0,1",
		"sg4":                            "required|in:Tidak,Ya",
		"nilai4":                         "required|in:0,1",
		"total_hasil":                    "int",
		"skor_nutrisi":                   "in:Risikio Berat,Risiko Sedang,Risiko Rendah",
		"diketahui_dietisien":            "required|in:Tidak,Ya",
		"keterangan_diketahui_dietisien": "string|max_len:10",
		"nip":                            "required|string|max_len:20",
	}
	return rules
}

// SkriningNutrisiAnakStore simpan skrining nutrisi anak; kolom waktu kunci kosong = sekarang.
type SkriningNutrisiAnakStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningNutrisiAnakData
}

func (r *SkriningNutrisiAnakStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningNutrisiAnakStore) Rules(ctx http.Context) map[string]any {
	rules := skriningNutrisiAnakRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningNutrisiAnakStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningNutrisiAnakStore) Payload() SkriningNutrisiAnakData {
	return r.SkriningNutrisiAnakData
}

func (r *SkriningNutrisiAnakStore) DetailValues() map[string][]string { return nil }

// SkriningNutrisiAnakUpdate ubah skrining nutrisi anak (PUT); kunci lewat query string.
type SkriningNutrisiAnakUpdate struct {
	SkriningNutrisiAnakData
}

func (r *SkriningNutrisiAnakUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningNutrisiAnakUpdate) Rules(ctx http.Context) map[string]any {
	return skriningNutrisiAnakRules()
}

func (r *SkriningNutrisiAnakUpdate) Payload() SkriningNutrisiAnakData {
	return r.SkriningNutrisiAnakData
}

func (r *SkriningNutrisiAnakUpdate) DetailValues() map[string][]string { return nil }
