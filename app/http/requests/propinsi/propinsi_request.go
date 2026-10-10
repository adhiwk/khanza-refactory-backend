package propinsi

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field propinsi yang dipakai bersama oleh store & update.
type Data struct {
	NmProp string `form:"nm_prop" json:"nm_prop"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nm_prop": "required|string|max_len:30",
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

// UpdateRequest mengganti data propinsi (PUT); kd_prop diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
