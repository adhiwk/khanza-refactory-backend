package satuanobat

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field satuan obat yang dipakai bersama oleh store & update.
type Data struct {
	Satuan string `form:"satuan" json:"satuan"`
}

func baseRules() map[string]any {
	return map[string]any{
		"satuan": "required|string|max_len:30",
	}
}

type StoreRequest struct {
	KodeSat string `form:"kode_sat" json:"kode_sat"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kode_sat"] = "required|string|max_len:4"
	return rules
}

// UpdateRequest mengganti data satuan obat (PUT); kode_sat diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
