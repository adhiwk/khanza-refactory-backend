package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanObservasiChbpData isian catatan observasi CHBP.
type CatatanObservasiChbpData struct {
	Td         string `form:"td" json:"td"`
	Hr         string `form:"hr" json:"hr"`
	Suhu       string `form:"suhu" json:"suhu"`
	Djj        string `form:"djj" json:"djj"`
	His        string `form:"his" json:"his"`
	Ppv        string `form:"ppv" json:"ppv"`
	Keterangan string `form:"keterangan" json:"keterangan"`
	Nip        string `form:"nip" json:"nip"`
}

func catatanObservasiChbpRules() map[string]any {
	rules := map[string]any{
		"td":         "string|max_len:8",
		"hr":         "string|max_len:5",
		"suhu":       "string|max_len:5",
		"djj":        "string|max_len:5",
		"his":        "string|max_len:20",
		"ppv":        "string|max_len:10",
		"keterangan": "string|max_len:50",
		"nip":        "required|string|max_len:20",
	}
	return rules
}

// CatatanObservasiChbpStore simpan catatan observasi CHBP; kolom waktu kunci kosong = sekarang.
type CatatanObservasiChbpStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	CatatanObservasiChbpData
}

func (r *CatatanObservasiChbpStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiChbpStore) Rules(ctx http.Context) map[string]any {
	rules := catatanObservasiChbpRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *CatatanObservasiChbpStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *CatatanObservasiChbpStore) Payload() CatatanObservasiChbpData {
	return r.CatatanObservasiChbpData
}

func (r *CatatanObservasiChbpStore) DetailValues() map[string][]string { return nil }

// CatatanObservasiChbpUpdate ubah catatan observasi CHBP (PUT); kunci lewat query string.
type CatatanObservasiChbpUpdate struct {
	CatatanObservasiChbpData
}

func (r *CatatanObservasiChbpUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiChbpUpdate) Rules(ctx http.Context) map[string]any {
	return catatanObservasiChbpRules()
}

func (r *CatatanObservasiChbpUpdate) Payload() CatatanObservasiChbpData {
	return r.CatatanObservasiChbpData
}

func (r *CatatanObservasiChbpUpdate) DetailValues() map[string][]string { return nil }
