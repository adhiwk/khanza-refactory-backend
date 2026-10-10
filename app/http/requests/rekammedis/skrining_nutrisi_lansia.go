package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningNutrisiLansiaData isian skrining nutrisi lansia.
type SkriningNutrisiLansiaData struct {
	Tanggal     string `form:"tanggal" json:"tanggal"`
	Td          string `form:"td" json:"td"`
	Hr          string `form:"hr" json:"hr"`
	Rr          string `form:"rr" json:"rr"`
	Suhu        string `form:"suhu" json:"suhu"`
	Bb          string `form:"bb" json:"bb"`
	Tbpb        string `form:"tbpb" json:"tbpb"`
	Spo2        string `form:"spo2" json:"spo2"`
	Alergi      string `form:"alergi" json:"alergi"`
	Sg1         string `form:"sg1" json:"sg1"`
	Nilai1      string `form:"nilai1" json:"nilai1"`
	Sg2         string `form:"sg2" json:"sg2"`
	Nilai2      string `form:"nilai2" json:"nilai2"`
	Sg3         string `form:"sg3" json:"sg3"`
	Nilai3      string `form:"nilai3" json:"nilai3"`
	Sg4         string `form:"sg4" json:"sg4"`
	Nilai4      string `form:"nilai4" json:"nilai4"`
	Sg5         string `form:"sg5" json:"sg5"`
	Nilai5      string `form:"nilai5" json:"nilai5"`
	Sg6         string `form:"sg6" json:"sg6"`
	Nilai6      string `form:"nilai6" json:"nilai6"`
	TotalHasil  int    `form:"total_hasil" json:"total_hasil"`
	SkorNutrisi string `form:"skor_nutrisi" json:"skor_nutrisi"`
	Nip         string `form:"nip" json:"nip"`
}

func skriningNutrisiLansiaRules() map[string]any {
	rules := map[string]any{
		"tanggal":      "required|date",
		"td":           "string|max_len:8",
		"hr":           "string|max_len:5",
		"rr":           "string|max_len:5",
		"suhu":         "string|max_len:5",
		"bb":           "string|max_len:5",
		"tbpb":         "string|max_len:5",
		"spo2":         "string|max_len:5",
		"alergi":       "string|max_len:100",
		"sg1":          "required|in:Asupan Makan Sangat Berkurang,Asupan Makan Agak Berkurang,Asupan Makan Tidak Berkurang",
		"nilai1":       "required|in:0,1,2",
		"sg2":          "required|in:Penurunan Berat Badan Lebih Dari 3 Kg,Tidak Tahu,Penurunan Berat Badan Antara 1 Hingga 3 Kg,Tidak Ada Penurunan Berat Badan",
		"nilai2":       "required|in:0,1,2,3",
		"sg3":          "required|in:Terbatas Dari Tempat Tidur Atau Kursi,Mampu Bangun Dari Tempat Tidur/Kursi Tetapi Tidak Bepergian Keluar Rumah,Dapat Bepergian Keluar Rumah",
		"nilai3":       "required|in:0,1,2",
		"sg4":          "required|in:Ya,Tidak",
		"nilai4":       "required|in:0,1",
		"sg5":          "required|in:Depresi Berat Atau Kepikunan Berat,Kepikunan Ringan,Tidak Ada Gangguan Psikologis",
		"nilai5":       "required|in:0,1,2",
		"sg6":          "required|in:IMT < 19,19 Hingga < 21,21 Hingga < 23,IMT >= 23,Lingkar Betis < 31,Lingkar Betis >= 31",
		"nilai6":       "required|in:0,1,2,3",
		"total_hasil":  "int",
		"skor_nutrisi": "in:Status Gizi Normal,Beresiko Malnutrisi,Malnutrisi",
		"nip":          "required|string|max_len:20",
	}
	return rules
}

// SkriningNutrisiLansiaStore simpan skrining nutrisi lansia; kolom waktu kunci kosong = sekarang.
type SkriningNutrisiLansiaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningNutrisiLansiaData
}

func (r *SkriningNutrisiLansiaStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningNutrisiLansiaStore) Rules(ctx http.Context) map[string]any {
	rules := skriningNutrisiLansiaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningNutrisiLansiaStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningNutrisiLansiaStore) Payload() SkriningNutrisiLansiaData {
	return r.SkriningNutrisiLansiaData
}

func (r *SkriningNutrisiLansiaStore) DetailValues() map[string][]string { return nil }

// SkriningNutrisiLansiaUpdate ubah skrining nutrisi lansia (PUT); kunci lewat query string.
type SkriningNutrisiLansiaUpdate struct {
	SkriningNutrisiLansiaData
}

func (r *SkriningNutrisiLansiaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningNutrisiLansiaUpdate) Rules(ctx http.Context) map[string]any {
	return skriningNutrisiLansiaRules()
}

func (r *SkriningNutrisiLansiaUpdate) Payload() SkriningNutrisiLansiaData {
	return r.SkriningNutrisiLansiaData
}

func (r *SkriningNutrisiLansiaUpdate) DetailValues() map[string][]string { return nil }
