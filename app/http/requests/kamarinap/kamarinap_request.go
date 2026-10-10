package kamarinap

import (
	"github.com/goravel/framework/contracts/http"
)

type MasukRequest struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	KdKamar      string `form:"kd_kamar" json:"kd_kamar"`
	DiagnosaAwal string `form:"diagnosa_awal" json:"diagnosa_awal"`
	TglMasuk     string `form:"tgl_masuk" json:"tgl_masuk"`
	JamMasuk     string `form:"jam_masuk" json:"jam_masuk"`
}

func (r *MasukRequest) Authorize(ctx http.Context) error { return nil }

func (r *MasukRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rawat":      "required|string|max_len:17",
		"kd_kamar":      "required|string|max_len:15",
		"diagnosa_awal": "required|string|max_len:100",
		"tgl_masuk":     "date",
		"jam_masuk":     "string|len:8",
	}
}

// PindahRequest mode: 1 ganti (hapus kamar lama), 2 ganti (ubah kamar), 3 pindah, 4 pindah dengan tarif tertinggi.
type PindahRequest struct {
	NoRawat   string `form:"no_rawat" json:"no_rawat"`
	KdKamar   string `form:"kd_kamar" json:"kd_kamar"`
	Mode      int    `form:"mode" json:"mode"`
	TglPindah string `form:"tgl_pindah" json:"tgl_pindah"`
	JamPindah string `form:"jam_pindah" json:"jam_pindah"`
}

func (r *PindahRequest) Authorize(ctx http.Context) error { return nil }

func (r *PindahRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rawat":   "required|string|max_len:17",
		"kd_kamar":   "required|string|max_len:15",
		"mode":       "required|in:1,2,3,4",
		"tgl_pindah": "date",
		"jam_pindah": "string|len:8",
	}
}

type PulangRequest struct {
	NoRawat       string `form:"no_rawat" json:"no_rawat"`
	SttsPulang    string `form:"stts_pulang" json:"stts_pulang"`
	DiagnosaAkhir string `form:"diagnosa_akhir" json:"diagnosa_akhir"`
	TglKeluar     string `form:"tgl_keluar" json:"tgl_keluar"`
	JamKeluar     string `form:"jam_keluar" json:"jam_keluar"`
}

func (r *PulangRequest) Authorize(ctx http.Context) error { return nil }

func (r *PulangRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rawat":       "required|string|max_len:17",
		"stts_pulang":    "required|string",
		"diagnosa_akhir": "required|string|max_len:100",
		"tgl_keluar":     "date",
		"jam_keluar":     "string|len:8",
	}
}

type BatalPulangRequest struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
}

func (r *BatalPulangRequest) Authorize(ctx http.Context) error { return nil }

func (r *BatalPulangRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{"no_rawat": "required|string|max_len:17"}
}
