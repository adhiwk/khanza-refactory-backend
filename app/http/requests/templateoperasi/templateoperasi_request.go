package templateoperasi

import (
	"github.com/goravel/framework/contracts/http"
)

// Data isian template laporan operasi; no_template dibuat otomatis ("O" + 4 digit).
type Data struct {
	NamaOperasi      string `form:"nama_operasi" json:"nama_operasi"`
	DiagnosaPreop    string `form:"diagnosa_preop" json:"diagnosa_preop"`
	DiagnosaPostop   string `form:"diagnosa_postop" json:"diagnosa_postop"`
	JaringanDieksisi string `form:"jaringan_dieksisi" json:"jaringan_dieksisi"`
	PermintaanPa     string `form:"permintaan_pa" json:"permintaan_pa"`
	LaporanOperasi   string `form:"laporan_operasi" json:"laporan_operasi"`
}

type Request struct {
	Data
}

func (r *Request) Authorize(ctx http.Context) error { return nil }

func (r *Request) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"nama_operasi":      "required|string|max_len:100",
		"diagnosa_preop":    "string|max_len:100",
		"diagnosa_postop":   "string|max_len:100",
		"jaringan_dieksisi": "string|max_len:100",
		"permintaan_pa":     "required|in:Ya,Tidak",
		"laporan_operasi":   "required|string",
	}
}
