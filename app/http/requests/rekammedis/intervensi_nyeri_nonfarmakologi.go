package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// IntervensiNyeriNonfarmakologiData isian intervensi nyeri non farmakologi.
type IntervensiNyeriNonfarmakologiData struct {
	Intervensi string `form:"intervensi" json:"intervensi"`
	Nip        string `form:"nip" json:"nip"`
}

func intervensiNyeriNonfarmakologiRules() map[string]any {
	rules := map[string]any{
		"intervensi": "string|max_len:200",
		"nip":        "required|string|max_len:20",
	}
	return rules
}

// IntervensiNyeriNonfarmakologiStore simpan intervensi nyeri non farmakologi; kolom waktu kunci kosong = sekarang.
type IntervensiNyeriNonfarmakologiStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	IntervensiNyeriNonfarmakologiData
}

func (r *IntervensiNyeriNonfarmakologiStore) Authorize(ctx http.Context) error { return nil }

func (r *IntervensiNyeriNonfarmakologiStore) Rules(ctx http.Context) map[string]any {
	rules := intervensiNyeriNonfarmakologiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *IntervensiNyeriNonfarmakologiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *IntervensiNyeriNonfarmakologiStore) Payload() IntervensiNyeriNonfarmakologiData {
	return r.IntervensiNyeriNonfarmakologiData
}

func (r *IntervensiNyeriNonfarmakologiStore) DetailValues() map[string][]string { return nil }

// IntervensiNyeriNonfarmakologiUpdate ubah intervensi nyeri non farmakologi (PUT); kunci lewat query string.
type IntervensiNyeriNonfarmakologiUpdate struct {
	IntervensiNyeriNonfarmakologiData
}

func (r *IntervensiNyeriNonfarmakologiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *IntervensiNyeriNonfarmakologiUpdate) Rules(ctx http.Context) map[string]any {
	return intervensiNyeriNonfarmakologiRules()
}

func (r *IntervensiNyeriNonfarmakologiUpdate) Payload() IntervensiNyeriNonfarmakologiData {
	return r.IntervensiNyeriNonfarmakologiData
}

func (r *IntervensiNyeriNonfarmakologiUpdate) DetailValues() map[string][]string { return nil }
