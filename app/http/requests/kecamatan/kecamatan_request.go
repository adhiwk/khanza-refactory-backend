package kecamatan

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field kecamatan yang dipakai bersama oleh store & update.
type Data struct {
	NmKec string `form:"nm_kec" json:"nm_kec"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nm_kec": "required|string|max_len:60",
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

// UpdateRequest mengganti data kecamatan (PUT); kd_kec diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
