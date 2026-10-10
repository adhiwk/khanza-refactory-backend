package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// UjiFungsiKfrData isian uji fungsi KFR.
type UjiFungsiKfrData struct {
	Tanggal             string `form:"tanggal" json:"tanggal"`
	DiagnosisFungsional string `form:"diagnosis_fungsional" json:"diagnosis_fungsional"`
	DiagnosisMedis      string `form:"diagnosis_medis" json:"diagnosis_medis"`
	HasilDidapat        string `form:"hasil_didapat" json:"hasil_didapat"`
	Kesimpulan          string `form:"kesimpulan" json:"kesimpulan"`
	Rekomedasi          string `form:"rekomedasi" json:"rekomedasi"`
	KdDokter            string `form:"kd_dokter" json:"kd_dokter"`
}

func ujiFungsiKfrRules() map[string]any {
	rules := map[string]any{
		"tanggal":              "date",
		"diagnosis_fungsional": "string|max_len:50",
		"diagnosis_medis":      "string|max_len:50",
		"hasil_didapat":        "string|max_len:100",
		"kesimpulan":           "string|max_len:100",
		"rekomedasi":           "string|max_len:100",
		"kd_dokter":            "string|max_len:20",
	}
	return rules
}

// UjiFungsiKfrStore simpan uji fungsi KFR; kolom waktu kunci kosong = sekarang.
type UjiFungsiKfrStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	UjiFungsiKfrData
}

func (r *UjiFungsiKfrStore) Authorize(ctx http.Context) error { return nil }

func (r *UjiFungsiKfrStore) Rules(ctx http.Context) map[string]any {
	rules := ujiFungsiKfrRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *UjiFungsiKfrStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *UjiFungsiKfrStore) Payload() UjiFungsiKfrData { return r.UjiFungsiKfrData }

func (r *UjiFungsiKfrStore) DetailValues() map[string][]string { return nil }

// UjiFungsiKfrUpdate ubah uji fungsi KFR (PUT); kunci lewat query string.
type UjiFungsiKfrUpdate struct {
	UjiFungsiKfrData
}

func (r *UjiFungsiKfrUpdate) Authorize(ctx http.Context) error { return nil }

func (r *UjiFungsiKfrUpdate) Rules(ctx http.Context) map[string]any { return ujiFungsiKfrRules() }

func (r *UjiFungsiKfrUpdate) Payload() UjiFungsiKfrData { return r.UjiFungsiKfrData }

func (r *UjiFungsiKfrUpdate) DetailValues() map[string][]string { return nil }
