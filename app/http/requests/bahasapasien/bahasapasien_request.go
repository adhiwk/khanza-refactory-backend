package bahasapasien

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field bahasa pasien yang dipakai bersama oleh store & update.
type Data struct {
	NamaBahasa string `form:"nama_bahasa" json:"nama_bahasa"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nama_bahasa": "required|string|max_len:30",
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

// UpdateRequest mengganti data bahasa pasien (PUT); id diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
