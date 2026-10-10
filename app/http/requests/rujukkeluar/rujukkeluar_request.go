package rujukkeluar

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field rujukan keluar yang dipakai bersama oleh store & update.
type Data struct {
	NoRawat            string `form:"no_rawat" json:"no_rawat"`
	RujukKe            string `form:"rujuk_ke" json:"rujuk_ke"`
	TglRujuk           string `form:"tgl_rujuk" json:"tgl_rujuk"`
	KeteranganDiagnosa string `form:"keterangan_diagnosa" json:"keterangan_diagnosa"`
	KdDokter           string `form:"kd_dokter" json:"kd_dokter"`
	KatRujuk           string `form:"kat_rujuk" json:"kat_rujuk"`
	Ambulance          string `form:"ambulance" json:"ambulance"`
	Keterangan         string `form:"keterangan" json:"keterangan"`
	Jam                string `form:"jam" json:"jam"`
}

func baseRules() map[string]any {
	return map[string]any{
		"no_rawat":            "required|string|max_len:17",
		"rujuk_ke":            "required|string|max_len:150",
		"tgl_rujuk":           "required|date",
		"keterangan_diagnosa": "string",
		"kd_dokter":           "required|string|max_len:20",
		"kat_rujuk":           "in:-,Bedah,Non Bedah,Kebidanan,Anak",
		"ambulance":           "in:-,AGD,SENDIRI,SWASTA",
		"keterangan":          "string",
		"jam":                 "string",
	}
}

// StoreRequest no_rujuk dibuat otomatis oleh server.
type StoreRequest struct {
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}

// UpdateRequest mengganti data rujukan keluar (PUT); no_rujuk diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
