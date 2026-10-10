package sukubangsa

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field suku bangsa yang dipakai bersama oleh store & update.
type Data struct {
	NamaSukuBangsa string `form:"nama_suku_bangsa" json:"nama_suku_bangsa"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nama_suku_bangsa": "required|string|max_len:30",
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

// UpdateRequest mengganti data suku bangsa (PUT); id diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
