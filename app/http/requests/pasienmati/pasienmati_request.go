package pasienmati

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field pasien meninggal yang dipakai bersama oleh store & update.
type Data struct {
	Tanggal       string `form:"tanggal" json:"tanggal"`
	Jam           string `form:"jam" json:"jam"`
	Keterangan    string `form:"keterangan" json:"keterangan"`
	TempMeninggal string `form:"temp_meninggal" json:"temp_meninggal"`
	Icd1          string `form:"icd1" json:"icd1"`
	Icd2          string `form:"icd2" json:"icd2"`
	Icd3          string `form:"icd3" json:"icd3"`
	Icd4          string `form:"icd4" json:"icd4"`
	KdDokter      string `form:"kd_dokter" json:"kd_dokter"`
}

func baseRules() map[string]any {
	return map[string]any{
		"tanggal":        "required|date",
		"jam":            "required|string",
		"keterangan":     "required|string|max_len:100",
		"temp_meninggal": "in:-,Rumah Sakit,Puskesmas,Rumah Bersalin,Rumah Tempat Tinggal",
		"icd1":           "string|max_len:20",
		"icd2":           "string|max_len:20",
		"icd3":           "string|max_len:20",
		"icd4":           "string|max_len:20",
		"kd_dokter":      "required|string|max_len:20",
	}
}

type StoreRequest struct {
	NoRkmMedis string `form:"no_rkm_medis" json:"no_rkm_medis"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["no_rkm_medis"] = "required|string|max_len:15"
	return rules
}

// UpdateRequest mengganti data pasien meninggal (PUT); no_rkm_medis diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
