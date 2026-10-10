package industrifarmasi

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field industri farmasi yang dipakai bersama oleh store & update.
type Data struct {
	NamaIndustri string `form:"nama_industri" json:"nama_industri"`
	Alamat       string `form:"alamat" json:"alamat"`
	Kota         string `form:"kota" json:"kota"`
	NoTelp       string `form:"no_telp" json:"no_telp"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nama_industri": "required|string|max_len:50",
		"alamat":        "string|max_len:50",
		"kota":          "string|max_len:20",
		"no_telp":       "string|max_len:20",
	}
}

type StoreRequest struct {
	KodeIndustri string `form:"kode_industri" json:"kode_industri"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kode_industri"] = "required|string|max_len:5"
	return rules
}

// UpdateRequest mengganti data industri farmasi (PUT); kode_industri diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
