package icd

import (
	"github.com/goravel/framework/contracts/http"
)

type Item struct {
	Kode      string `form:"kode" json:"kode"`
	Prioritas int    `form:"prioritas" json:"prioritas"`
	Jumlah    string `form:"jumlah" json:"jumlah"`
}

// StoreRequest status Ralan/Ranap opsional (default status lanjut registrasi); prioritas 0 = urutan berikutnya.
type StoreRequest struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Status  string `form:"status" json:"status"`
	Items   []Item `form:"items" json:"items"`
}

func (r *StoreRequest) Authorize(ctx http.Context) error { return nil }

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rawat":          "required|string|max_len:17",
		"status":            "in:Ralan,Ranap",
		"items":             "required|slice",
		"items.*.kode":      "required|string|max_len:15",
		"items.*.prioritas": "int|min:0",
		"items.*.jumlah":    "string|max_len:3",
	}
}

type StatusPenyakitRequest struct {
	NoRawat        string `form:"no_rawat" json:"no_rawat"`
	Kode           string `form:"kode" json:"kode"`
	StatusPenyakit string `form:"status_penyakit" json:"status_penyakit"`
}

func (r *StatusPenyakitRequest) Authorize(ctx http.Context) error { return nil }

func (r *StatusPenyakitRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rawat":        "required|string|max_len:17",
		"kode":            "required|string|max_len:15",
		"status_penyakit": "required|in:Lama,Baru",
	}
}
