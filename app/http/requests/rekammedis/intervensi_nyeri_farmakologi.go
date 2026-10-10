package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// IntervensiNyeriFarmakologiData isian intervensi nyeri farmakologi.
type IntervensiNyeriFarmakologiData struct {
	NamaObat  string `form:"nama_obat" json:"nama_obat"`
	DosisEfek string `form:"dosis_efek" json:"dosis_efek"`
	Rute      string `form:"rute" json:"rute"`
	Nip       string `form:"nip" json:"nip"`
}

func intervensiNyeriFarmakologiRules() map[string]any {
	rules := map[string]any{
		"nama_obat":  "string|max_len:100",
		"dosis_efek": "string|max_len:100",
		"rute":       "string|max_len:30",
		"nip":        "required|string|max_len:20",
	}
	return rules
}

// IntervensiNyeriFarmakologiStore simpan intervensi nyeri farmakologi; kolom waktu kunci kosong = sekarang.
type IntervensiNyeriFarmakologiStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	IntervensiNyeriFarmakologiData
}

func (r *IntervensiNyeriFarmakologiStore) Authorize(ctx http.Context) error { return nil }

func (r *IntervensiNyeriFarmakologiStore) Rules(ctx http.Context) map[string]any {
	rules := intervensiNyeriFarmakologiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *IntervensiNyeriFarmakologiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *IntervensiNyeriFarmakologiStore) Payload() IntervensiNyeriFarmakologiData {
	return r.IntervensiNyeriFarmakologiData
}

func (r *IntervensiNyeriFarmakologiStore) DetailValues() map[string][]string { return nil }

// IntervensiNyeriFarmakologiUpdate ubah intervensi nyeri farmakologi (PUT); kunci lewat query string.
type IntervensiNyeriFarmakologiUpdate struct {
	IntervensiNyeriFarmakologiData
}

func (r *IntervensiNyeriFarmakologiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *IntervensiNyeriFarmakologiUpdate) Rules(ctx http.Context) map[string]any {
	return intervensiNyeriFarmakologiRules()
}

func (r *IntervensiNyeriFarmakologiUpdate) Payload() IntervensiNyeriFarmakologiData {
	return r.IntervensiNyeriFarmakologiData
}

func (r *IntervensiNyeriFarmakologiUpdate) DetailValues() map[string][]string { return nil }
