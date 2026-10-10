package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningInstrumenEsatData isian skrining instrumen ESAT.
type SkriningInstrumenEsatData struct {
	Tanggal          string `form:"tanggal" json:"tanggal"`
	Nip              string `form:"nip" json:"nip"`
	Pernyataanesat1  string `form:"pernyataanesat1" json:"pernyataanesat1"`
	Pernyataanesat2  string `form:"pernyataanesat2" json:"pernyataanesat2"`
	Pernyataanesat3  string `form:"pernyataanesat3" json:"pernyataanesat3"`
	Pernyataanesat4  string `form:"pernyataanesat4" json:"pernyataanesat4"`
	Pernyataanesat5  string `form:"pernyataanesat5" json:"pernyataanesat5"`
	Pernyataanesat6  string `form:"pernyataanesat6" json:"pernyataanesat6"`
	Pernyataanesat7  string `form:"pernyataanesat7" json:"pernyataanesat7"`
	Pernyataanesat8  string `form:"pernyataanesat8" json:"pernyataanesat8"`
	Pernyataanesat9  string `form:"pernyataanesat9" json:"pernyataanesat9"`
	Pernyataanesat10 string `form:"pernyataanesat10" json:"pernyataanesat10"`
	Kesimpulan       string `form:"kesimpulan" json:"kesimpulan"`
	Catatan          string `form:"catatan" json:"catatan"`
}

func skriningInstrumenEsatRules() map[string]any {
	rules := map[string]any{
		"tanggal":          "required|date",
		"nip":              "required|string|max_len:20",
		"pernyataanesat1":  "in:Ya,Tidak",
		"pernyataanesat2":  "in:Ya,Tidak",
		"pernyataanesat3":  "in:Ya,Tidak",
		"pernyataanesat4":  "in:Ya,Tidak",
		"pernyataanesat5":  "in:Ya,Tidak",
		"pernyataanesat6":  "in:Ya,Tidak",
		"pernyataanesat7":  "in:Ya,Tidak",
		"pernyataanesat8":  "in:Ya,Tidak",
		"pernyataanesat9":  "in:Ya,Tidak",
		"pernyataanesat10": "in:Ya,Tidak",
		"kesimpulan":       "string|max_len:150",
		"catatan":          "string|max_len:100",
	}
	return rules
}

// SkriningInstrumenEsatStore simpan skrining instrumen ESAT; kolom waktu kunci kosong = sekarang.
type SkriningInstrumenEsatStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningInstrumenEsatData
}

func (r *SkriningInstrumenEsatStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningInstrumenEsatStore) Rules(ctx http.Context) map[string]any {
	rules := skriningInstrumenEsatRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningInstrumenEsatStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningInstrumenEsatStore) Payload() SkriningInstrumenEsatData {
	return r.SkriningInstrumenEsatData
}

func (r *SkriningInstrumenEsatStore) DetailValues() map[string][]string { return nil }

// SkriningInstrumenEsatUpdate ubah skrining instrumen ESAT (PUT); kunci lewat query string.
type SkriningInstrumenEsatUpdate struct {
	SkriningInstrumenEsatData
}

func (r *SkriningInstrumenEsatUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningInstrumenEsatUpdate) Rules(ctx http.Context) map[string]any {
	return skriningInstrumenEsatRules()
}

func (r *SkriningInstrumenEsatUpdate) Payload() SkriningInstrumenEsatData {
	return r.SkriningInstrumenEsatData
}

func (r *SkriningInstrumenEsatUpdate) DetailValues() map[string][]string { return nil }
