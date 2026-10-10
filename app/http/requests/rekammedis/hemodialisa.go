package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HemodialisaData isian hemodialisa.
type HemodialisaData struct {
	KdDokter   string `form:"kd_dokter" json:"kd_dokter"`
	Lama       string `form:"lama" json:"lama"`
	Akses      string `form:"akses" json:"akses"`
	Dialist    string `form:"dialist" json:"dialist"`
	Transfusi  string `form:"transfusi" json:"transfusi"`
	Penarikan  string `form:"penarikan" json:"penarikan"`
	Qb         string `form:"qb" json:"qb"`
	Qd         string `form:"qd" json:"qd"`
	Ureum      string `form:"ureum" json:"ureum"`
	Hb         string `form:"hb" json:"hb"`
	Hbsag      string `form:"hbsag" json:"hbsag"`
	Creatinin  string `form:"creatinin" json:"creatinin"`
	Hiv        string `form:"hiv" json:"hiv"`
	Hcv        string `form:"hcv" json:"hcv"`
	Lain       string `form:"lain" json:"lain"`
	KdPenyakit string `form:"kd_penyakit" json:"kd_penyakit"`
}

func hemodialisaRules() map[string]any {
	rules := map[string]any{
		"kd_dokter":   "string|max_len:20",
		"lama":        "string|max_len:5",
		"akses":       "string|max_len:30",
		"dialist":     "string|max_len:30",
		"transfusi":   "string|max_len:5",
		"penarikan":   "string|max_len:5",
		"qb":          "string|max_len:5",
		"qd":          "string|max_len:5",
		"ureum":       "string|max_len:10",
		"hb":          "string|max_len:10",
		"hbsag":       "string|max_len:10",
		"creatinin":   "string|max_len:10",
		"hiv":         "string|max_len:10",
		"hcv":         "string|max_len:10",
		"lain":        "string|max_len:200",
		"kd_penyakit": "string|max_len:15",
	}
	return rules
}

// HemodialisaStore simpan hemodialisa; kolom waktu kunci kosong = sekarang.
type HemodialisaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	HemodialisaData
}

func (r *HemodialisaStore) Authorize(ctx http.Context) error { return nil }

func (r *HemodialisaStore) Rules(ctx http.Context) map[string]any {
	rules := hemodialisaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *HemodialisaStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *HemodialisaStore) Payload() HemodialisaData { return r.HemodialisaData }

func (r *HemodialisaStore) DetailValues() map[string][]string { return nil }

// HemodialisaUpdate ubah hemodialisa (PUT); kunci lewat query string.
type HemodialisaUpdate struct {
	HemodialisaData
}

func (r *HemodialisaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HemodialisaUpdate) Rules(ctx http.Context) map[string]any { return hemodialisaRules() }

func (r *HemodialisaUpdate) Payload() HemodialisaData { return r.HemodialisaData }

func (r *HemodialisaUpdate) DetailValues() map[string][]string { return nil }
