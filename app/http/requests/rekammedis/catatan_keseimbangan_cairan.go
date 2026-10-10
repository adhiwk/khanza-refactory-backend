package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanKeseimbanganCairanData isian catatan keseimbangan cairan.
type CatatanKeseimbanganCairanData struct {
	Infus        string `form:"infus" json:"infus"`
	Tranfusi     string `form:"tranfusi" json:"tranfusi"`
	Minum        string `form:"minum" json:"minum"`
	Urine        string `form:"urine" json:"urine"`
	Drain        string `form:"drain" json:"drain"`
	Ngt          string `form:"ngt" json:"ngt"`
	Iwl          string `form:"iwl" json:"iwl"`
	Keseimbangan string `form:"keseimbangan" json:"keseimbangan"`
	Keterangan   string `form:"keterangan" json:"keterangan"`
	Nip          string `form:"nip" json:"nip"`
}

func catatanKeseimbanganCairanRules() map[string]any {
	rules := map[string]any{
		"infus":        "string|max_len:4",
		"tranfusi":     "string|max_len:4",
		"minum":        "string|max_len:4",
		"urine":        "string|max_len:4",
		"drain":        "string|max_len:4",
		"ngt":          "string|max_len:4",
		"iwl":          "string|max_len:4",
		"keseimbangan": "string|max_len:4",
		"keterangan":   "string|max_len:200",
		"nip":          "required|string|max_len:20",
	}
	return rules
}

// CatatanKeseimbanganCairanStore simpan catatan keseimbangan cairan; kolom waktu kunci kosong = sekarang.
type CatatanKeseimbanganCairanStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	CatatanKeseimbanganCairanData
}

func (r *CatatanKeseimbanganCairanStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanKeseimbanganCairanStore) Rules(ctx http.Context) map[string]any {
	rules := catatanKeseimbanganCairanRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *CatatanKeseimbanganCairanStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *CatatanKeseimbanganCairanStore) Payload() CatatanKeseimbanganCairanData {
	return r.CatatanKeseimbanganCairanData
}

func (r *CatatanKeseimbanganCairanStore) DetailValues() map[string][]string { return nil }

// CatatanKeseimbanganCairanUpdate ubah catatan keseimbangan cairan (PUT); kunci lewat query string.
type CatatanKeseimbanganCairanUpdate struct {
	CatatanKeseimbanganCairanData
}

func (r *CatatanKeseimbanganCairanUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanKeseimbanganCairanUpdate) Rules(ctx http.Context) map[string]any {
	return catatanKeseimbanganCairanRules()
}

func (r *CatatanKeseimbanganCairanUpdate) Payload() CatatanKeseimbanganCairanData {
	return r.CatatanKeseimbanganCairanData
}

func (r *CatatanKeseimbanganCairanUpdate) DetailValues() map[string][]string { return nil }
