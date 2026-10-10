package kamar

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field kamar yang dipakai bersama oleh store & update.
type Data struct {
	KdBangsal string   `form:"kd_bangsal" json:"kd_bangsal"`
	TrfKamar  *float64 `form:"trf_kamar" json:"trf_kamar"`
	Status    string   `form:"status" json:"status"`
	Kelas     string   `form:"kelas" json:"kelas"`
}

func baseRules() map[string]any {
	return map[string]any{
		"kd_bangsal": "required|string|max_len:5",
		"trf_kamar":  "required|numeric",
		"status":     "required|in:ISI,KOSONG,DIBERSIHKAN,DIBOOKING,PERBAIKAN",
		"kelas":      "required|in:Kelas 1,Kelas 2,Kelas 3,Kelas Utama,Kelas VIP,Kelas VVIP",
	}
}

type StoreRequest struct {
	KdKamar string `form:"kd_kamar" json:"kd_kamar"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kd_kamar"] = "required|string|max_len:15"
	return rules
}

// UpdateRequest mengganti data kamar (PUT); kd_kamar diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
