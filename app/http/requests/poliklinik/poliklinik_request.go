package poliklinik

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field poliklinik yang dipakai bersama oleh store & update.
type Data struct {
	NmPoli         string  `form:"nm_poli" json:"nm_poli"`
	Registrasi     float64 `form:"registrasi" json:"registrasi"`
	Registrasilama float64 `form:"registrasilama" json:"registrasilama"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nm_poli":        "required|string|max_len:50",
		"registrasi":     "required|numeric",
		"registrasilama": "required|numeric",
	}
}

type StoreRequest struct {
	KdPoli string `form:"kd_poli" json:"kd_poli"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kd_poli"] = "required|string|max_len:5"
	return rules
}

// UpdateRequest mengganti data poliklinik (PUT); kd_poli diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
