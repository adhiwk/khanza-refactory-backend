package golonganobat

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field golongan obat yang dipakai bersama oleh store & update.
type Data struct {
	Nama string `form:"nama" json:"nama"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nama": "required|string|max_len:30",
	}
}

type StoreRequest struct {
	Kode string `form:"kode" json:"kode"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kode"] = "required|string|max_len:4"
	return rules
}

// UpdateRequest mengganti data golongan obat (PUT); kode diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
