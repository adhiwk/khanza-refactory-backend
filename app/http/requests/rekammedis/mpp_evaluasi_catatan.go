package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// MppEvaluasiCatatanData isian skrining MPP form b.
type MppEvaluasiCatatanData struct {
	Masalah  string `form:"masalah" json:"masalah"`
	Tinjut   string `form:"tinjut" json:"tinjut"`
	Evaluasi string `form:"evaluasi" json:"evaluasi"`
	Nip      string `form:"nip" json:"nip"`
}

func mppEvaluasiCatatanRules() map[string]any {
	rules := map[string]any{
		"masalah":  "string|max_len:500",
		"tinjut":   "string|max_len:500",
		"evaluasi": "string|max_len:500",
		"nip":      "required|string|max_len:20",
	}
	return rules
}

// MppEvaluasiCatatanStore simpan skrining MPP form b; kolom waktu kunci kosong = sekarang.
type MppEvaluasiCatatanStore struct {
	NoRawat         string `form:"no_rawat" json:"no_rawat"`
	TglImplementasi string `form:"tgl_implementasi" json:"tgl_implementasi"`
	MppEvaluasiCatatanData
}

func (r *MppEvaluasiCatatanStore) Authorize(ctx http.Context) error { return nil }

func (r *MppEvaluasiCatatanStore) Rules(ctx http.Context) map[string]any {
	rules := mppEvaluasiCatatanRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_implementasi"] = "date"
	return rules
}

func (r *MppEvaluasiCatatanStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_implementasi": r.TglImplementasi}
}

func (r *MppEvaluasiCatatanStore) Payload() MppEvaluasiCatatanData { return r.MppEvaluasiCatatanData }

func (r *MppEvaluasiCatatanStore) DetailValues() map[string][]string { return nil }

// MppEvaluasiCatatanUpdate ubah skrining MPP form b (PUT); kunci lewat query string.
type MppEvaluasiCatatanUpdate struct {
	MppEvaluasiCatatanData
}

func (r *MppEvaluasiCatatanUpdate) Authorize(ctx http.Context) error { return nil }

func (r *MppEvaluasiCatatanUpdate) Rules(ctx http.Context) map[string]any {
	return mppEvaluasiCatatanRules()
}

func (r *MppEvaluasiCatatanUpdate) Payload() MppEvaluasiCatatanData { return r.MppEvaluasiCatatanData }

func (r *MppEvaluasiCatatanUpdate) DetailValues() map[string][]string { return nil }
