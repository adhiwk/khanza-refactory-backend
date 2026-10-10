package kelurahan

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field kelurahan yang dipakai bersama oleh store & update.
type Data struct {
	NmKel string `form:"nm_kel" json:"nm_kel"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nm_kel": "required|string|max_len:60",
	}
}

type StoreRequest struct {
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	return rules
}

// UpdateRequest mengganti data kelurahan (PUT); kd_kel diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
