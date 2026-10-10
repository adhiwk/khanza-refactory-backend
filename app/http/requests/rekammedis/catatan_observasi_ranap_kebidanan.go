package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanObservasiRanapKebidananData isian catatan observasi ranap kebidanan.
type CatatanObservasiRanapKebidananData struct {
	Gcs       string `form:"gcs" json:"gcs"`
	Td        string `form:"td" json:"td"`
	Hr        string `form:"hr" json:"hr"`
	Rr        string `form:"rr" json:"rr"`
	Suhu      string `form:"suhu" json:"suhu"`
	Spo2      string `form:"spo2" json:"spo2"`
	Kontraksi string `form:"kontraksi" json:"kontraksi"`
	Bjj       string `form:"bjj" json:"bjj"`
	Ppv       string `form:"ppv" json:"ppv"`
	Vt        string `form:"vt" json:"vt"`
	Nip       string `form:"nip" json:"nip"`
}

func catatanObservasiRanapKebidananRules() map[string]any {
	rules := map[string]any{
		"gcs":       "string|max_len:10",
		"td":        "string|max_len:8",
		"hr":        "string|max_len:5",
		"rr":        "string|max_len:5",
		"suhu":      "string|max_len:5",
		"spo2":      "string|max_len:3",
		"kontraksi": "string|max_len:15",
		"bjj":       "string|max_len:5",
		"ppv":       "string|max_len:15",
		"vt":        "string|max_len:30",
		"nip":       "required|string|max_len:20",
	}
	return rules
}

// CatatanObservasiRanapKebidananStore simpan catatan observasi ranap kebidanan; kolom waktu kunci kosong = sekarang.
type CatatanObservasiRanapKebidananStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	CatatanObservasiRanapKebidananData
}

func (r *CatatanObservasiRanapKebidananStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiRanapKebidananStore) Rules(ctx http.Context) map[string]any {
	rules := catatanObservasiRanapKebidananRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *CatatanObservasiRanapKebidananStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *CatatanObservasiRanapKebidananStore) Payload() CatatanObservasiRanapKebidananData {
	return r.CatatanObservasiRanapKebidananData
}

func (r *CatatanObservasiRanapKebidananStore) DetailValues() map[string][]string { return nil }

// CatatanObservasiRanapKebidananUpdate ubah catatan observasi ranap kebidanan (PUT); kunci lewat query string.
type CatatanObservasiRanapKebidananUpdate struct {
	CatatanObservasiRanapKebidananData
}

func (r *CatatanObservasiRanapKebidananUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiRanapKebidananUpdate) Rules(ctx http.Context) map[string]any {
	return catatanObservasiRanapKebidananRules()
}

func (r *CatatanObservasiRanapKebidananUpdate) Payload() CatatanObservasiRanapKebidananData {
	return r.CatatanObservasiRanapKebidananData
}

func (r *CatatanObservasiRanapKebidananUpdate) DetailValues() map[string][]string { return nil }
