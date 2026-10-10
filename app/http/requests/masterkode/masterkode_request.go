package masterkode

import (
	"github.com/goravel/framework/contracts/http"
)

// Data isian master kode; kode kosong saat simpan = dibuat otomatis, kode_induk wajib untuk rencana/skala.
type Data struct {
	Nama      string `form:"nama" json:"nama"`
	KodeInduk string `form:"kode_induk" json:"kode_induk"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nama":       "required|string|max_len:1000",
		"kode_induk": "string|max_len:3",
	}
}

type StoreRequest struct {
	Kode string `form:"kode" json:"kode"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error { return nil }

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kode"] = "string|max_len:3"
	return rules
}

// UpdateRequest kode diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error { return nil }

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any { return baseRules() }
