package kategoriperawatan

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field kategori perawatan yang dipakai bersama oleh store & update.
type Data struct {
	NmKategori string `form:"nm_kategori" json:"nm_kategori"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nm_kategori": "required|string|max_len:30",
	}
}

type StoreRequest struct {
	KdKategori string `form:"kd_kategori" json:"kd_kategori"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kd_kategori"] = "required|string|max_len:5"
	return rules
}

// UpdateRequest mengganti data kategori perawatan (PUT); kd_kategori diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
