package perusahaanpasien

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field instansi/perusahaan pasien yang dipakai bersama oleh store & update.
type Data struct {
	NamaPerusahaan string `form:"nama_perusahaan" json:"nama_perusahaan"`
	Alamat         string `form:"alamat" json:"alamat"`
	Kota           string `form:"kota" json:"kota"`
	NoTelp         string `form:"no_telp" json:"no_telp"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nama_perusahaan": "required|string|max_len:70",
		"alamat":          "string|max_len:100",
		"kota":            "string|max_len:40",
		"no_telp":         "string|max_len:27",
	}
}

type StoreRequest struct {
	KodePerusahaan string `form:"kode_perusahaan" json:"kode_perusahaan"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kode_perusahaan"] = "required|string|max_len:8"
	return rules
}

// UpdateRequest mengganti data instansi/perusahaan pasien (PUT); kode_perusahaan diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
