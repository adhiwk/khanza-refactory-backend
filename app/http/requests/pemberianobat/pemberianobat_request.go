package pemberianobat

import (
	"github.com/goravel/framework/contracts/http"
)

type Item struct {
	KodeBrng    string  `form:"kode_brng" json:"kode_brng"`
	Jml         float64 `form:"jml" json:"jml"`
	Embalase    float64 `form:"embalase" json:"embalase"`
	Tuslah      float64 `form:"tuslah" json:"tuslah"`
	AturanPakai string  `form:"aturan_pakai" json:"aturan_pakai"`
	NoBatch     string  `form:"no_batch" json:"no_batch"`
	NoFaktur    string  `form:"no_faktur" json:"no_faktur"`
}

// StoreRequest harga jual & HPP dihitung server; kd_bangsal (depo) & jenis_harga opsional.
type StoreRequest struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	KdBangsal    string `form:"kd_bangsal" json:"kd_bangsal"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	Jam          string `form:"jam" json:"jam"`
	JenisHarga   string `form:"jenis_harga" json:"jenis_harga"`
	Items        []Item `form:"items" json:"items"`
}

func (r *StoreRequest) Authorize(ctx http.Context) error { return nil }

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rawat":             "required|string|max_len:17",
		"kd_bangsal":           "string|max_len:5",
		"tgl_perawatan":        "date",
		"jam":                  "string|len:8",
		"jenis_harga":          "string",
		"items":                "required|slice",
		"items.*.kode_brng":    "required|string|max_len:15",
		"items.*.jml":          "required|numeric|gt:0",
		"items.*.embalase":     "numeric|min:0",
		"items.*.tuslah":       "numeric|min:0",
		"items.*.aturan_pakai": "string|max_len:150",
		"items.*.no_batch":     "string|max_len:20",
		"items.*.no_faktur":    "string|max_len:20",
	}
}

type DeleteRequest struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	Jam          string `form:"jam" json:"jam"`
	KodeBrng     string `form:"kode_brng" json:"kode_brng"`
	NoBatch      string `form:"no_batch" json:"no_batch"`
	NoFaktur     string `form:"no_faktur" json:"no_faktur"`
}

func (r *DeleteRequest) Authorize(ctx http.Context) error { return nil }

func (r *DeleteRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rawat":      "required|string|max_len:17",
		"tgl_perawatan": "required|date",
		"jam":           "required|string|len:8",
		"kode_brng":     "required|string|max_len:15",
		"no_batch":      "string|max_len:20",
		"no_faktur":     "string|max_len:20",
	}
}
