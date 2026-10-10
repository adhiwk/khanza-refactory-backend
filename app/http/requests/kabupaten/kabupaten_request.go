package kabupaten

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field kabupaten yang dipakai bersama oleh store & update.
type Data struct {
	NmKab string `form:"nm_kab" json:"nm_kab"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nm_kab": "required|string|max_len:60",
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

// UpdateRequest mengganti data kabupaten (PUT); kd_kab diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
