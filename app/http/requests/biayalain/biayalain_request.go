package biayalain

import (
	"github.com/goravel/framework/contracts/http"
)

type StoreRequest struct {
	NoRawat string  `form:"no_rawat" json:"no_rawat"`
	Nama    string  `form:"nama" json:"nama"`
	Besar   float64 `form:"besar" json:"besar"`
}

func (r *StoreRequest) Authorize(ctx http.Context) error { return nil }

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rawat": "required|string|max_len:17",
		"nama":     "required|string|max_len:60",
		"besar":    "required|numeric|gt:0",
	}
}
