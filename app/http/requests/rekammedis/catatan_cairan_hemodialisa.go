package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanCairanHemodialisaData isian catatan cairan hemodialisa.
type CatatanCairanHemodialisaData struct {
	Minum       string `form:"minum" json:"minum"`
	Infus       string `form:"infus" json:"infus"`
	Tranfusi    string `form:"tranfusi" json:"tranfusi"`
	SisaPriming string `form:"sisa_priming" json:"sisa_priming"`
	WashOut     string `form:"wash_out" json:"wash_out"`
	Urine       string `form:"urine" json:"urine"`
	Pendarahan  string `form:"pendarahan" json:"pendarahan"`
	Muntah      string `form:"muntah" json:"muntah"`
	Keterangan  string `form:"keterangan" json:"keterangan"`
	Nip         string `form:"nip" json:"nip"`
}

func catatanCairanHemodialisaRules() map[string]any {
	rules := map[string]any{
		"minum":        "string|max_len:10",
		"infus":        "string|max_len:10",
		"tranfusi":     "string|max_len:10",
		"sisa_priming": "string|max_len:10",
		"wash_out":     "string|max_len:10",
		"urine":        "string|max_len:10",
		"pendarahan":   "string|max_len:10",
		"muntah":       "string|max_len:10",
		"keterangan":   "string|max_len:100",
		"nip":          "string|max_len:20",
	}
	return rules
}

// CatatanCairanHemodialisaStore simpan catatan cairan hemodialisa; kolom waktu kunci kosong = sekarang.
type CatatanCairanHemodialisaStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	CatatanCairanHemodialisaData
}

func (r *CatatanCairanHemodialisaStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanCairanHemodialisaStore) Rules(ctx http.Context) map[string]any {
	rules := catatanCairanHemodialisaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *CatatanCairanHemodialisaStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *CatatanCairanHemodialisaStore) Payload() CatatanCairanHemodialisaData {
	return r.CatatanCairanHemodialisaData
}

func (r *CatatanCairanHemodialisaStore) DetailValues() map[string][]string { return nil }

// CatatanCairanHemodialisaUpdate ubah catatan cairan hemodialisa (PUT); kunci lewat query string.
type CatatanCairanHemodialisaUpdate struct {
	CatatanCairanHemodialisaData
}

func (r *CatatanCairanHemodialisaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanCairanHemodialisaUpdate) Rules(ctx http.Context) map[string]any {
	return catatanCairanHemodialisaRules()
}

func (r *CatatanCairanHemodialisaUpdate) Payload() CatatanCairanHemodialisaData {
	return r.CatatanCairanHemodialisaData
}

func (r *CatatanCairanHemodialisaUpdate) DetailValues() map[string][]string { return nil }
