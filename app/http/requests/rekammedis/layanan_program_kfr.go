package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// LayananProgramKfrData isian layanan program KFR.
type LayananProgramKfrData struct {
	NoRawatLayanan string `form:"no_rawat_layanan" json:"no_rawat_layanan"`
	Tanggal        string `form:"tanggal" json:"tanggal"`
	Nip            string `form:"nip" json:"nip"`
	Program        string `form:"program" json:"program"`
}

func layananProgramKfrRules() map[string]any {
	rules := map[string]any{
		"no_rawat_layanan": "string|max_len:17",
		"tanggal":          "required|date",
		"nip":              "required|string|max_len:20",
		"program":          "string|max_len:50",
	}
	return rules
}

// LayananProgramKfrStore simpan layanan program KFR; kolom waktu kunci kosong = sekarang.
type LayananProgramKfrStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	LayananProgramKfrData
}

func (r *LayananProgramKfrStore) Authorize(ctx http.Context) error { return nil }

func (r *LayananProgramKfrStore) Rules(ctx http.Context) map[string]any {
	rules := layananProgramKfrRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *LayananProgramKfrStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *LayananProgramKfrStore) Payload() LayananProgramKfrData { return r.LayananProgramKfrData }

func (r *LayananProgramKfrStore) DetailValues() map[string][]string { return nil }

// LayananProgramKfrUpdate ubah layanan program KFR (PUT); kunci lewat query string.
type LayananProgramKfrUpdate struct {
	LayananProgramKfrData
}

func (r *LayananProgramKfrUpdate) Authorize(ctx http.Context) error { return nil }

func (r *LayananProgramKfrUpdate) Rules(ctx http.Context) map[string]any {
	return layananProgramKfrRules()
}

func (r *LayananProgramKfrUpdate) Payload() LayananProgramKfrData { return r.LayananProgramKfrData }

func (r *LayananProgramKfrUpdate) DetailValues() map[string][]string { return nil }
