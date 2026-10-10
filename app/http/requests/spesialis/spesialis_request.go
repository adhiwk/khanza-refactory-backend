package spesialis

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field spesialis yang dipakai bersama oleh store & update.
type Data struct {
	NmSps string `form:"nm_sps" json:"nm_sps"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nm_sps": "required|string|max_len:30",
	}
}

type StoreRequest struct {
	KdSps string `form:"kd_sps" json:"kd_sps"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kd_sps"] = "required|string|max_len:5"
	return rules
}

// UpdateRequest mengganti data spesialis (PUT); kd_sps diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
