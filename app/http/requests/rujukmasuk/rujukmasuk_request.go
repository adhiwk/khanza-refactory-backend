package rujukmasuk

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field rujukan masuk yang dipakai bersama oleh store & update.
type Data struct {
	Perujuk       string  `form:"perujuk" json:"perujuk"`
	Alamat        string  `form:"alamat" json:"alamat"`
	NoRujuk       string  `form:"no_rujuk" json:"no_rujuk"`
	JmPerujuk     float64 `form:"jm_perujuk" json:"jm_perujuk"`
	DokterPerujuk string  `form:"dokter_perujuk" json:"dokter_perujuk"`
	KdPenyakit    string  `form:"kd_penyakit" json:"kd_penyakit"`
	KategoriRujuk string  `form:"kategori_rujuk" json:"kategori_rujuk"`
	Keterangan    string  `form:"keterangan" json:"keterangan"`
}

func baseRules() map[string]any {
	return map[string]any{
		"perujuk":        "required|string|max_len:60",
		"alamat":         "string|max_len:70",
		"no_rujuk":       "string|max_len:40",
		"jm_perujuk":     "required|numeric",
		"dokter_perujuk": "string|max_len:50",
		"kd_penyakit":    "string|max_len:15",
		"kategori_rujuk": "in:-,Bedah,Non-Bedah,Kebidanan,Anak",
		"keterangan":     "string|max_len:200",
	}
}

type StoreRequest struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

// UpdateRequest mengganti data rujukan masuk (PUT); no_rawat diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
