package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanObservasiRanapData isian catatan observasi ranap.
type CatatanObservasiRanapData struct {
	Gcs  string `form:"gcs" json:"gcs"`
	Td   string `form:"td" json:"td"`
	Hr   string `form:"hr" json:"hr"`
	Rr   string `form:"rr" json:"rr"`
	Suhu string `form:"suhu" json:"suhu"`
	Spo2 string `form:"spo2" json:"spo2"`
	Nip  string `form:"nip" json:"nip"`
}

func catatanObservasiRanapRules() map[string]any {
	rules := map[string]any{
		"gcs":  "string|max_len:10",
		"td":   "string|max_len:8",
		"hr":   "string|max_len:5",
		"rr":   "string|max_len:5",
		"suhu": "string|max_len:5",
		"spo2": "string|max_len:3",
		"nip":  "required|string|max_len:20",
	}
	return rules
}

// CatatanObservasiRanapStore simpan catatan observasi ranap; kolom waktu kunci kosong = sekarang.
type CatatanObservasiRanapStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	CatatanObservasiRanapData
}

func (r *CatatanObservasiRanapStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiRanapStore) Rules(ctx http.Context) map[string]any {
	rules := catatanObservasiRanapRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *CatatanObservasiRanapStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *CatatanObservasiRanapStore) Payload() CatatanObservasiRanapData {
	return r.CatatanObservasiRanapData
}

func (r *CatatanObservasiRanapStore) DetailValues() map[string][]string { return nil }

// CatatanObservasiRanapUpdate ubah catatan observasi ranap (PUT); kunci lewat query string.
type CatatanObservasiRanapUpdate struct {
	CatatanObservasiRanapData
}

func (r *CatatanObservasiRanapUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiRanapUpdate) Rules(ctx http.Context) map[string]any {
	return catatanObservasiRanapRules()
}

func (r *CatatanObservasiRanapUpdate) Payload() CatatanObservasiRanapData {
	return r.CatatanObservasiRanapData
}

func (r *CatatanObservasiRanapUpdate) DetailValues() map[string][]string { return nil }
