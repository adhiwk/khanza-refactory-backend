package metoderacik

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field metode racik yang dipakai bersama oleh store & update.
type Data struct {
	NmRacik string `form:"nm_racik" json:"nm_racik"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nm_racik": "required|string|max_len:30",
	}
}

type StoreRequest struct {
	KdRacik string `form:"kd_racik" json:"kd_racik"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kd_racik"] = "required|string|max_len:3"
	return rules
}

// UpdateRequest mengganti data metode racik (PUT); kd_racik diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
