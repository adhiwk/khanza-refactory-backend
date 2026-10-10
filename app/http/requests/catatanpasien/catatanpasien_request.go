package catatanpasien

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field catatan pasien yang dipakai bersama oleh store & update.
type Data struct {
	Catatan string `form:"catatan" json:"catatan"`
}

func baseRules() map[string]any {
	return map[string]any{
		"catatan": "required|string",
	}
}

type StoreRequest struct {
	NoRkmMedis string `form:"no_rkm_medis" json:"no_rkm_medis"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["no_rkm_medis"] = "required|string|max_len:15"
	return rules
}

// UpdateRequest mengganti data catatan pasien (PUT); no_rkm_medis diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
