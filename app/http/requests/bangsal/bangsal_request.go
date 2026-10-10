package bangsal

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field bangsal yang dipakai bersama oleh store & update.
type Data struct {
	NmBangsal string `form:"nm_bangsal" json:"nm_bangsal"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nm_bangsal": "required|string|max_len:30",
	}
}

type StoreRequest struct {
	KdBangsal string `form:"kd_bangsal" json:"kd_bangsal"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kd_bangsal"] = "required|string|max_len:5"
	return rules
}

// UpdateRequest mengganti data bangsal (PUT); kd_bangsal diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
