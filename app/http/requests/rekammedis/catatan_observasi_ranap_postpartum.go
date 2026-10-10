package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanObservasiRanapPostpartumData isian catatan observasi ranap post partum.
type CatatanObservasiRanapPostpartumData struct {
	Gcs        string `form:"gcs" json:"gcs"`
	Td         string `form:"td" json:"td"`
	Hr         string `form:"hr" json:"hr"`
	Rr         string `form:"rr" json:"rr"`
	Suhu       string `form:"suhu" json:"suhu"`
	Spo2       string `form:"spo2" json:"spo2"`
	Tfu        string `form:"tfu" json:"tfu"`
	Kontraksi  string `form:"kontraksi" json:"kontraksi"`
	Perdarahan string `form:"perdarahan" json:"perdarahan"`
	Keterangan string `form:"keterangan" json:"keterangan"`
	Nip        string `form:"nip" json:"nip"`
}

func catatanObservasiRanapPostpartumRules() map[string]any {
	rules := map[string]any{
		"gcs":        "string|max_len:10",
		"td":         "string|max_len:8",
		"hr":         "string|max_len:5",
		"rr":         "string|max_len:5",
		"suhu":       "string|max_len:5",
		"spo2":       "string|max_len:3",
		"tfu":        "string|max_len:15",
		"kontraksi":  "string|max_len:15",
		"perdarahan": "string|max_len:15",
		"keterangan": "string|max_len:30",
		"nip":        "required|string|max_len:20",
	}
	return rules
}

// CatatanObservasiRanapPostpartumStore simpan catatan observasi ranap post partum; kolom waktu kunci kosong = sekarang.
type CatatanObservasiRanapPostpartumStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	CatatanObservasiRanapPostpartumData
}

func (r *CatatanObservasiRanapPostpartumStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiRanapPostpartumStore) Rules(ctx http.Context) map[string]any {
	rules := catatanObservasiRanapPostpartumRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *CatatanObservasiRanapPostpartumStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *CatatanObservasiRanapPostpartumStore) Payload() CatatanObservasiRanapPostpartumData {
	return r.CatatanObservasiRanapPostpartumData
}

func (r *CatatanObservasiRanapPostpartumStore) DetailValues() map[string][]string { return nil }

// CatatanObservasiRanapPostpartumUpdate ubah catatan observasi ranap post partum (PUT); kunci lewat query string.
type CatatanObservasiRanapPostpartumUpdate struct {
	CatatanObservasiRanapPostpartumData
}

func (r *CatatanObservasiRanapPostpartumUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiRanapPostpartumUpdate) Rules(ctx http.Context) map[string]any {
	return catatanObservasiRanapPostpartumRules()
}

func (r *CatatanObservasiRanapPostpartumUpdate) Payload() CatatanObservasiRanapPostpartumData {
	return r.CatatanObservasiRanapPostpartumData
}

func (r *CatatanObservasiRanapPostpartumUpdate) DetailValues() map[string][]string { return nil }
