package registrasi

import (
	"github.com/goravel/framework/contracts/http"
)

// RegistrasiData field yang dipakai bersama oleh store & update.
// kd_pj, p_jawab, almt_pj, hubunganpj kosong diisi dari data pasien; biaya_reg kosong diambil dari tarif poliklinik.
type RegistrasiData struct {
	TglRegistrasi string   `form:"tgl_registrasi" json:"tgl_registrasi"` // YYYY-MM-DD, default hari ini
	JamReg        string   `form:"jam_reg" json:"jam_reg"`               // HH:MM:SS, default jam sekarang
	KdDokter      string   `form:"kd_dokter" json:"kd_dokter"`
	KdPoli        string   `form:"kd_poli" json:"kd_poli"`
	KdPj          string   `form:"kd_pj" json:"kd_pj"`
	PJawab        string   `form:"p_jawab" json:"p_jawab"`
	AlmtPj        string   `form:"almt_pj" json:"almt_pj"`
	HubunganPj    string   `form:"hubunganpj" json:"hubunganpj"`
	BiayaReg      *float64 `form:"biaya_reg" json:"biaya_reg"`
}

func baseRules() map[string]any {
	return map[string]any{
		"tgl_registrasi": "date",
		"jam_reg":        "string|len:8",
		"kd_dokter":      "required|string|max_len:20",
		"kd_poli":        "required|string|max_len:5",
		"kd_pj":          "string|max_len:3",
		"p_jawab":        "string|max_len:100",
		"almt_pj":        "string|max_len:200",
		"hubunganpj":     "string|max_len:20",
		"biaya_reg":      "numeric|min:0",
	}
}

type StoreRegistrasiRequest struct {
	NoRkmMedis string `form:"no_rkm_medis" json:"no_rkm_medis"`
	RegistrasiData
}

func (r *StoreRegistrasiRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRegistrasiRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["no_rkm_medis"] = "required|string|max_len:15"
	return rules
}

func (r *StoreRegistrasiRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

// UpdateRegistrasiRequest mengganti data registrasi (PUT); no_rawat diambil dari route, pasien tidak bisa diganti.
type UpdateRegistrasiRequest struct {
	RegistrasiData
}

func (r *UpdateRegistrasiRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRegistrasiRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}

func (r *UpdateRegistrasiRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}
