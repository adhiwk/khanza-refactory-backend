package igd

import (
	"github.com/goravel/framework/contracts/http"

	regrequest "goravel/app/http/requests/registrasi"
)

// StoreRequest registrasi IGD; kd_poli diabaikan (selalu IGDK). Asal rujukan opsional.
type StoreRequest struct {
	NoRkmMedis    string `form:"no_rkm_medis" json:"no_rkm_medis"`
	Perujuk       string `form:"perujuk" json:"perujuk"`
	AlamatPerujuk string `form:"alamat_perujuk" json:"alamat_perujuk"`
	NoRujuk       string `form:"no_rujuk" json:"no_rujuk"`
	regrequest.RegistrasiData
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rkm_medis":   "required|string|max_len:15",
		"tgl_registrasi": "date",
		"jam_reg":        "string|len:8",
		"kd_dokter":      "required|string|max_len:20",
		"kd_pj":          "string|max_len:3",
		"p_jawab":        "string|max_len:100",
		"almt_pj":        "string|max_len:200",
		"hubunganpj":     "string|max_len:20",
		"biaya_reg":      "numeric|min:0",
		"perujuk":        "string|max_len:60",
		"alamat_perujuk": "string|max_len:70",
		"no_rujuk":       "string|max_len:40",
	}
}
