package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanObservasiRuangOkData isian catatan observasi ruang operasi.
type CatatanObservasiRuangOkData struct {
	Gcs        string `form:"gcs" json:"gcs"`
	Td         string `form:"td" json:"td"`
	Hr         string `form:"hr" json:"hr"`
	Rr         string `form:"rr" json:"rr"`
	Suhu       string `form:"suhu" json:"suhu"`
	Spo2       string `form:"spo2" json:"spo2"`
	Keterangan string `form:"keterangan" json:"keterangan"`
	Nip        string `form:"nip" json:"nip"`
}

func catatanObservasiRuangOkRules() map[string]any {
	rules := map[string]any{
		"gcs":        "string|max_len:10",
		"td":         "string|max_len:8",
		"hr":         "string|max_len:5",
		"rr":         "string|max_len:5",
		"suhu":       "string|max_len:5",
		"spo2":       "string|max_len:3",
		"keterangan": "string|max_len:100",
		"nip":        "required|string|max_len:20",
	}
	return rules
}

// CatatanObservasiRuangOkStore simpan catatan observasi ruang operasi; kolom waktu kunci kosong = sekarang.
type CatatanObservasiRuangOkStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	CatatanObservasiRuangOkData
}

func (r *CatatanObservasiRuangOkStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiRuangOkStore) Rules(ctx http.Context) map[string]any {
	rules := catatanObservasiRuangOkRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *CatatanObservasiRuangOkStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *CatatanObservasiRuangOkStore) Payload() CatatanObservasiRuangOkData {
	return r.CatatanObservasiRuangOkData
}

func (r *CatatanObservasiRuangOkStore) DetailValues() map[string][]string { return nil }

// CatatanObservasiRuangOkUpdate ubah catatan observasi ruang operasi (PUT); kunci lewat query string.
type CatatanObservasiRuangOkUpdate struct {
	CatatanObservasiRuangOkData
}

func (r *CatatanObservasiRuangOkUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiRuangOkUpdate) Rules(ctx http.Context) map[string]any {
	return catatanObservasiRuangOkRules()
}

func (r *CatatanObservasiRuangOkUpdate) Payload() CatatanObservasiRuangOkData {
	return r.CatatanObservasiRuangOkData
}

func (r *CatatanObservasiRuangOkUpdate) DetailValues() map[string][]string { return nil }
