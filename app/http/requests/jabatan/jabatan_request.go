package jabatan

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field jabatan yang dipakai bersama oleh store & update.
type Data struct {
	NmJbtn string `form:"nm_jbtn" json:"nm_jbtn"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nm_jbtn": "required|string|max_len:25",
	}
}

type StoreRequest struct {
	KdJbtn string `form:"kd_jbtn" json:"kd_jbtn"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kd_jbtn"] = "required|string|max_len:4"
	return rules
}

// UpdateRequest mengganti data jabatan (PUT); kd_jbtn diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
