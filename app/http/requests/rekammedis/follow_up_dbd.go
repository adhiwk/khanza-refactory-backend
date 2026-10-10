package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// FollowUpDbdData isian follow up DBD.
type FollowUpDbdData struct {
	Hemoglobin   string `form:"hemoglobin" json:"hemoglobin"`
	Hematokrit   string `form:"hematokrit" json:"hematokrit"`
	Leokosit     string `form:"leokosit" json:"leokosit"`
	Trombosit    string `form:"trombosit" json:"trombosit"`
	TerapiCairan string `form:"terapi_cairan" json:"terapi_cairan"`
	Nip          string `form:"nip" json:"nip"`
}

func followUpDbdRules() map[string]any {
	rules := map[string]any{
		"hemoglobin":    "string|max_len:5",
		"hematokrit":    "string|max_len:5",
		"leokosit":      "string|max_len:7",
		"trombosit":     "string|max_len:10",
		"terapi_cairan": "string|max_len:100",
		"nip":           "required|string|max_len:20",
	}
	return rules
}

// FollowUpDbdStore simpan follow up DBD; kolom waktu kunci kosong = sekarang.
type FollowUpDbdStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	FollowUpDbdData
}

func (r *FollowUpDbdStore) Authorize(ctx http.Context) error { return nil }

func (r *FollowUpDbdStore) Rules(ctx http.Context) map[string]any {
	rules := followUpDbdRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *FollowUpDbdStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *FollowUpDbdStore) Payload() FollowUpDbdData { return r.FollowUpDbdData }

func (r *FollowUpDbdStore) DetailValues() map[string][]string { return nil }

// FollowUpDbdUpdate ubah follow up DBD (PUT); kunci lewat query string.
type FollowUpDbdUpdate struct {
	FollowUpDbdData
}

func (r *FollowUpDbdUpdate) Authorize(ctx http.Context) error { return nil }

func (r *FollowUpDbdUpdate) Rules(ctx http.Context) map[string]any { return followUpDbdRules() }

func (r *FollowUpDbdUpdate) Payload() FollowUpDbdData { return r.FollowUpDbdData }

func (r *FollowUpDbdUpdate) DetailValues() map[string][]string { return nil }
