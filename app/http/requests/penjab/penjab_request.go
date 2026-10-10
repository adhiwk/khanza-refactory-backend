package penjab

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field penanggung jawab yang dipakai bersama oleh store & update.
type Data struct {
	PngJawab       string `form:"png_jawab" json:"png_jawab"`
	NamaPerusahaan string `form:"nama_perusahaan" json:"nama_perusahaan"`
	AlamatAsuransi string `form:"alamat_asuransi" json:"alamat_asuransi"`
	NoTelp         string `form:"no_telp" json:"no_telp"`
	Attn           string `form:"attn" json:"attn"`
}

func baseRules() map[string]any {
	return map[string]any{
		"png_jawab":       "required|string|max_len:30",
		"nama_perusahaan": "string|max_len:60",
		"alamat_asuransi": "string|max_len:130",
		"no_telp":         "string|max_len:40",
		"attn":            "string|max_len:60",
	}
}

type StoreRequest struct {
	KdPj string `form:"kd_pj" json:"kd_pj"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kd_pj"] = "required|string|max_len:3"
	return rules
}

// UpdateRequest mengganti data penanggung jawab (PUT); kd_pj diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
