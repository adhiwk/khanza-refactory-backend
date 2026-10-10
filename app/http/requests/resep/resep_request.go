package resep

import (
	"github.com/goravel/framework/contracts/http"
)

type Item struct {
	KodeBrng    string  `form:"kode_brng" json:"kode_brng"`
	Jml         float64 `form:"jml" json:"jml"`
	AturanPakai string  `form:"aturan_pakai" json:"aturan_pakai"`
}

func itemRules(rules map[string]any) map[string]any {
	rules["items"] = "required|slice"
	rules["items.*.kode_brng"] = "required|string|max_len:15"
	rules["items.*.jml"] = "required|numeric|gt:0"
	rules["items.*.aturan_pakai"] = "string|max_len:150"
	return rules
}

type StoreRequest struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	KdDokter     string `form:"kd_dokter" json:"kd_dokter"`
	TglPeresepan string `form:"tgl_peresepan" json:"tgl_peresepan"`
	JamPeresepan string `form:"jam_peresepan" json:"jam_peresepan"`
	Items        []Item `form:"items" json:"items"`
}

func (r *StoreRequest) Authorize(ctx http.Context) error { return nil }

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	return itemRules(map[string]any{
		"no_rawat":      "required|string|max_len:17",
		"kd_dokter":     "required|string|max_len:20",
		"tgl_peresepan": "date",
		"jam_peresepan": "string|len:8",
	})
}

// UpdateRequest mengganti daftar obat; no_resep lewat route.
type UpdateRequest struct {
	Items []Item `form:"items" json:"items"`
}

func (r *UpdateRequest) Authorize(ctx http.Context) error { return nil }

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return itemRules(map[string]any{})
}
