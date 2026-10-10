package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PemantauanEwsNeonatusData isian pemantauan EWS neonatus.
type PemantauanEwsNeonatusData struct {
	Parameter1     string `form:"parameter1" json:"parameter1"`
	Skor1          string `form:"skor1" json:"skor1"`
	Parameter2     string `form:"parameter2" json:"parameter2"`
	Skor2          string `form:"skor2" json:"skor2"`
	Parameter3     string `form:"parameter3" json:"parameter3"`
	Skor3          string `form:"skor3" json:"skor3"`
	Parameter4     string `form:"parameter4" json:"parameter4"`
	Skor4          string `form:"skor4" json:"skor4"`
	Parameter5     string `form:"parameter5" json:"parameter5"`
	Skor5          string `form:"skor5" json:"skor5"`
	Parameter6     string `form:"parameter6" json:"parameter6"`
	Skor6          string `form:"skor6" json:"skor6"`
	Parameter7     string `form:"parameter7" json:"parameter7"`
	Skor7          string `form:"skor7" json:"skor7"`
	Parameter8     string `form:"parameter8" json:"parameter8"`
	Skor8          string `form:"skor8" json:"skor8"`
	SkorTotal      string `form:"skor_total" json:"skor_total"`
	ParameterTotal string `form:"parameter_total" json:"parameter_total"`
	CodeBlue       string `form:"code_blue" json:"code_blue"`
	Nip            string `form:"nip" json:"nip"`
}

func pemantauanEwsNeonatusRules() map[string]any {
	rules := map[string]any{
		"parameter1":      "in:<= 29,30 - 39,40 - 60,>= 61",
		"skor1":           "string|max_len:1",
		"parameter2":      "in:<= 90,90 - 93,>= 94",
		"skor2":           "string|max_len:1",
		"parameter3":      "in:Tidak,Ya",
		"skor3":           "string|max_len:1",
		"parameter4":      "in:<= 80,81 - 119,120 - 160,161 - 180,>= 181",
		"skor4":           "string|max_len:1",
		"parameter5":      "in:Berat,Ringan,Tidak",
		"skor5":           "string|max_len:1",
		"parameter6":      "in:>= 3 Detik,<= 3 Detik",
		"skor6":           "string|max_len:1",
		"parameter7":      "string",
		"skor7":           "string|max_len:1",
		"parameter8":      "in:Pink,Pucat",
		"skor8":           "string|max_len:1",
		"skor_total":      "string|max_len:2",
		"parameter_total": "string|max_len:250",
		"code_blue":       "required|in:Ya,Tidak",
		"nip":             "string|max_len:20",
	}
	return rules
}

// PemantauanEwsNeonatusStore simpan pemantauan EWS neonatus; kolom waktu kunci kosong = sekarang.
type PemantauanEwsNeonatusStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PemantauanEwsNeonatusData
}

func (r *PemantauanEwsNeonatusStore) Authorize(ctx http.Context) error { return nil }

func (r *PemantauanEwsNeonatusStore) Rules(ctx http.Context) map[string]any {
	rules := pemantauanEwsNeonatusRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PemantauanEwsNeonatusStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PemantauanEwsNeonatusStore) Payload() PemantauanEwsNeonatusData {
	return r.PemantauanEwsNeonatusData
}

func (r *PemantauanEwsNeonatusStore) DetailValues() map[string][]string { return nil }

// PemantauanEwsNeonatusUpdate ubah pemantauan EWS neonatus (PUT); kunci lewat query string.
type PemantauanEwsNeonatusUpdate struct {
	PemantauanEwsNeonatusData
}

func (r *PemantauanEwsNeonatusUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PemantauanEwsNeonatusUpdate) Rules(ctx http.Context) map[string]any {
	return pemantauanEwsNeonatusRules()
}

func (r *PemantauanEwsNeonatusUpdate) Payload() PemantauanEwsNeonatusData {
	return r.PemantauanEwsNeonatusData
}

func (r *PemantauanEwsNeonatusUpdate) DetailValues() map[string][]string { return nil }
