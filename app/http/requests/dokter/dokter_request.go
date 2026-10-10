package dokter

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field dokter yang dipakai bersama oleh store & update.
type Data struct {
	NmDokter     string `form:"nm_dokter" json:"nm_dokter"`
	Jk           string `form:"jk" json:"jk"`
	TmpLahir     string `form:"tmp_lahir" json:"tmp_lahir"`
	TglLahir     string `form:"tgl_lahir" json:"tgl_lahir"`
	GolDrh       string `form:"gol_drh" json:"gol_drh"`
	Agama        string `form:"agama" json:"agama"`
	AlmtTgl      string `form:"almt_tgl" json:"almt_tgl"`
	NoTelp       string `form:"no_telp" json:"no_telp"`
	Email        string `form:"email" json:"email"`
	SttsNikah    string `form:"stts_nikah" json:"stts_nikah"`
	KdSps        string `form:"kd_sps" json:"kd_sps"`
	Alumni       string `form:"alumni" json:"alumni"`
	NoIjnPraktek string `form:"no_ijn_praktek" json:"no_ijn_praktek"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nm_dokter":      "required|string|max_len:50",
		"jk":             "required|in:L,P",
		"tmp_lahir":      "string|max_len:20",
		"tgl_lahir":      "date",
		"gol_drh":        "in:A,B,O,AB,-",
		"agama":          "string|max_len:12",
		"almt_tgl":       "string|max_len:60",
		"no_telp":        "string|max_len:13",
		"email":          "string|max_len:70",
		"stts_nikah":     "in:BELUM MENIKAH,MENIKAH,JANDA,DUDHA,JOMBLO",
		"kd_sps":         "required|string|max_len:5",
		"alumni":         "string|max_len:60",
		"no_ijn_praktek": "string|max_len:120",
	}
}

type StoreRequest struct {
	KdDokter string `form:"kd_dokter" json:"kd_dokter"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kd_dokter"] = "required|string|max_len:20"
	return rules
}

// UpdateRequest mengganti data dokter (PUT); kd_dokter diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}
