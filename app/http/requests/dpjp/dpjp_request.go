package dpjp

import (
	"github.com/goravel/framework/contracts/http"
)

type StoreRequest struct {
	NoRawat  string   `form:"no_rawat" json:"no_rawat"`
	KdDokter []string `form:"kd_dokter" json:"kd_dokter"`
}

func (r *StoreRequest) Authorize(ctx http.Context) error { return nil }

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rawat":    "required|string|max_len:17",
		"kd_dokter":   "required|slice",
		"kd_dokter.*": "required|string|max_len:20",
	}
}
