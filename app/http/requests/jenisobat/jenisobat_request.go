package jenisobat

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field jenis obat yang dipakai bersama oleh store & update.
type Data struct {
	Nama       string `form:"nama" json:"nama"`
	Keterangan string `form:"keterangan" json:"keterangan"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nama":       "required|string|max_len:30",
		"keterangan": "string|max_len:50",
	}
}

type StoreRequest struct {
	Kdjns string `form:"kdjns" json:"kdjns"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kdjns"] = "required|string|max_len:4"
	return rules
}

// UpdateRequest mengganti data jenis obat (PUT); kdjns diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
