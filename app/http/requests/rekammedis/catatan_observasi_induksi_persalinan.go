package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanObservasiInduksiPersalinanData isian catatan observasi induksi persalinan.
type CatatanObservasiInduksiPersalinanData struct {
	Obat       string `form:"obat" json:"obat"`
	Cairan     string `form:"cairan" json:"cairan"`
	Dosis      string `form:"dosis" json:"dosis"`
	His        string `form:"his" json:"his"`
	Djj        string `form:"djj" json:"djj"`
	Keterangan string `form:"keterangan" json:"keterangan"`
	Nip        string `form:"nip" json:"nip"`
}

func catatanObservasiInduksiPersalinanRules() map[string]any {
	rules := map[string]any{
		"obat":       "string|max_len:50",
		"cairan":     "string|max_len:50",
		"dosis":      "string|max_len:10",
		"his":        "string|max_len:50",
		"djj":        "string|max_len:5",
		"keterangan": "string|max_len:50",
		"nip":        "required|string|max_len:20",
	}
	return rules
}

// CatatanObservasiInduksiPersalinanStore simpan catatan observasi induksi persalinan; kolom waktu kunci kosong = sekarang.
type CatatanObservasiInduksiPersalinanStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	CatatanObservasiInduksiPersalinanData
}

func (r *CatatanObservasiInduksiPersalinanStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiInduksiPersalinanStore) Rules(ctx http.Context) map[string]any {
	rules := catatanObservasiInduksiPersalinanRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *CatatanObservasiInduksiPersalinanStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *CatatanObservasiInduksiPersalinanStore) Payload() CatatanObservasiInduksiPersalinanData {
	return r.CatatanObservasiInduksiPersalinanData
}

func (r *CatatanObservasiInduksiPersalinanStore) DetailValues() map[string][]string { return nil }

// CatatanObservasiInduksiPersalinanUpdate ubah catatan observasi induksi persalinan (PUT); kunci lewat query string.
type CatatanObservasiInduksiPersalinanUpdate struct {
	CatatanObservasiInduksiPersalinanData
}

func (r *CatatanObservasiInduksiPersalinanUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiInduksiPersalinanUpdate) Rules(ctx http.Context) map[string]any {
	return catatanObservasiInduksiPersalinanRules()
}

func (r *CatatanObservasiInduksiPersalinanUpdate) Payload() CatatanObservasiInduksiPersalinanData {
	return r.CatatanObservasiInduksiPersalinanData
}

func (r *CatatanObservasiInduksiPersalinanUpdate) DetailValues() map[string][]string { return nil }
