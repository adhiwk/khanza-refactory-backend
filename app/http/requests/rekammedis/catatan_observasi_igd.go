package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanObservasiIgdData isian catatan observasi IGD.
type CatatanObservasiIgdData struct {
	Gcs  string `form:"gcs" json:"gcs"`
	Td   string `form:"td" json:"td"`
	Hr   string `form:"hr" json:"hr"`
	Rr   string `form:"rr" json:"rr"`
	Suhu string `form:"suhu" json:"suhu"`
	Spo2 string `form:"spo2" json:"spo2"`
	Nip  string `form:"nip" json:"nip"`
}

func catatanObservasiIgdRules() map[string]any {
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

// CatatanObservasiIgdStore simpan catatan observasi IGD; kolom waktu kunci kosong = sekarang.
type CatatanObservasiIgdStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	CatatanObservasiIgdData
}

func (r *CatatanObservasiIgdStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiIgdStore) Rules(ctx http.Context) map[string]any {
	rules := catatanObservasiIgdRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *CatatanObservasiIgdStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *CatatanObservasiIgdStore) Payload() CatatanObservasiIgdData {
	return r.CatatanObservasiIgdData
}

func (r *CatatanObservasiIgdStore) DetailValues() map[string][]string { return nil }

// CatatanObservasiIgdUpdate ubah catatan observasi IGD (PUT); kunci lewat query string.
type CatatanObservasiIgdUpdate struct {
	CatatanObservasiIgdData
}

func (r *CatatanObservasiIgdUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiIgdUpdate) Rules(ctx http.Context) map[string]any {
	return catatanObservasiIgdRules()
}

func (r *CatatanObservasiIgdUpdate) Payload() CatatanObservasiIgdData {
	return r.CatatanObservasiIgdData
}

func (r *CatatanObservasiIgdUpdate) DetailValues() map[string][]string { return nil }
