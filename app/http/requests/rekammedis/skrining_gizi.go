package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningGiziData isian skrining gizi lanjut.
type SkriningGiziData struct {
	SkriningBb        string `form:"skrining_bb" json:"skrining_bb"`
	SkriningTb        string `form:"skrining_tb" json:"skrining_tb"`
	Alergi            string `form:"alergi" json:"alergi"`
	ParameterImt      string `form:"parameter_imt" json:"parameter_imt"`
	SkorImt           string `form:"skor_imt" json:"skor_imt"`
	ParameterBb       string `form:"parameter_bb" json:"parameter_bb"`
	SkorBb            string `form:"skor_bb" json:"skor_bb"`
	ParameterPenyakit string `form:"parameter_penyakit" json:"parameter_penyakit"`
	SkorPenyakit      string `form:"skor_penyakit" json:"skor_penyakit"`
	SkorTotal         string `form:"skor_total" json:"skor_total"`
	ParameterTotal    string `form:"parameter_total" json:"parameter_total"`
	Nip               string `form:"nip" json:"nip"`
}

func skriningGiziRules() map[string]any {
	rules := map[string]any{
		"skrining_bb":        "string|max_len:5",
		"skrining_tb":        "string|max_len:5",
		"alergi":             "string|max_len:25",
		"parameter_imt":      "string",
		"skor_imt":           "string|max_len:5",
		"parameter_bb":       "in:BB Hilang < 5%,BB Hilang 5 - 10 %,BB Hilang > 10 %",
		"skor_bb":            "string|max_len:5",
		"parameter_penyakit": "in:Ada asupan nutrisi > 5 hari,Tidak ada asupan nutrisi > 5 hari",
		"skor_penyakit":      "string|max_len:5",
		"skor_total":         "string|max_len:5",
		"parameter_total":    "string|max_len:200",
		"nip":                "string|max_len:20",
	}
	return rules
}

// SkriningGiziStore simpan skrining gizi lanjut; kolom waktu kunci kosong = sekarang.
type SkriningGiziStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	SkriningGiziData
}

func (r *SkriningGiziStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningGiziStore) Rules(ctx http.Context) map[string]any {
	rules := skriningGiziRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *SkriningGiziStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *SkriningGiziStore) Payload() SkriningGiziData { return r.SkriningGiziData }

func (r *SkriningGiziStore) DetailValues() map[string][]string { return nil }

// SkriningGiziUpdate ubah skrining gizi lanjut (PUT); kunci lewat query string.
type SkriningGiziUpdate struct {
	SkriningGiziData
}

func (r *SkriningGiziUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningGiziUpdate) Rules(ctx http.Context) map[string]any { return skriningGiziRules() }

func (r *SkriningGiziUpdate) Payload() SkriningGiziData { return r.SkriningGiziData }

func (r *SkriningGiziUpdate) DetailValues() map[string][]string { return nil }
