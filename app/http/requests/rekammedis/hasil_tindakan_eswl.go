package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilTindakanEswlData isian hasil tindakan ESWL.
type HasilTindakanEswlData struct {
	Selesai             string `form:"selesai" json:"selesai"`
	KdDokter            string `form:"kd_dokter" json:"kd_dokter"`
	Nip                 string `form:"nip" json:"nip"`
	Diagnosa            string `form:"diagnosa" json:"diagnosa"`
	Tindakan            string `form:"tindakan" json:"tindakan"`
	ObatAnalgesik       string `form:"obat_analgesik" json:"obat_analgesik"`
	ObatLain            string `form:"obat_lain" json:"obat_lain"`
	UraianTindakan      string `form:"uraian_tindakan" json:"uraian_tindakan"`
	UraianTindakanFocus string `form:"uraian_tindakan_focus" json:"uraian_tindakan_focus"`
	UraianTindakanRate  string `form:"uraian_tindakan_rate" json:"uraian_tindakan_rate"`
	UraianTindakanPower string `form:"uraian_tindakan_power" json:"uraian_tindakan_power"`
	UraianTindakanShock string `form:"uraian_tindakan_shock" json:"uraian_tindakan_shock"`
	Diintegrasi         string `form:"diintegrasi" json:"diintegrasi"`
	Kekurangan          string `form:"kekurangan" json:"kekurangan"`
	Anjungan            string `form:"anjungan" json:"anjungan"`
}

func hasilTindakanEswlRules() map[string]any {
	rules := map[string]any{
		"selesai":               "required|date",
		"kd_dokter":             "required|string|max_len:20",
		"nip":                   "required|string|max_len:20",
		"diagnosa":              "string|max_len:50",
		"tindakan":              "string|max_len:50",
		"obat_analgesik":        "string|max_len:150",
		"obat_lain":             "string|max_len:150",
		"uraian_tindakan":       "string|max_len:300",
		"uraian_tindakan_focus": "string|max_len:50",
		"uraian_tindakan_rate":  "string|max_len:50",
		"uraian_tindakan_power": "string|max_len:50",
		"uraian_tindakan_shock": "string|max_len:50",
		"diintegrasi":           "string|max_len:50",
		"kekurangan":            "string|max_len:50",
		"anjungan":              "string|max_len:50",
	}
	return rules
}

// HasilTindakanEswlStore simpan hasil tindakan ESWL; kolom waktu kunci kosong = sekarang.
type HasilTindakanEswlStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Mulai   string `form:"mulai" json:"mulai"`
	HasilTindakanEswlData
}

func (r *HasilTindakanEswlStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilTindakanEswlStore) Rules(ctx http.Context) map[string]any {
	rules := hasilTindakanEswlRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["mulai"] = "date"
	return rules
}

func (r *HasilTindakanEswlStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "mulai": r.Mulai}
}

func (r *HasilTindakanEswlStore) Payload() HasilTindakanEswlData { return r.HasilTindakanEswlData }

func (r *HasilTindakanEswlStore) DetailValues() map[string][]string { return nil }

// HasilTindakanEswlUpdate ubah hasil tindakan ESWL (PUT); kunci lewat query string.
type HasilTindakanEswlUpdate struct {
	HasilTindakanEswlData
}

func (r *HasilTindakanEswlUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilTindakanEswlUpdate) Rules(ctx http.Context) map[string]any {
	return hasilTindakanEswlRules()
}

func (r *HasilTindakanEswlUpdate) Payload() HasilTindakanEswlData { return r.HasilTindakanEswlData }

func (r *HasilTindakanEswlUpdate) DetailValues() map[string][]string { return nil }
