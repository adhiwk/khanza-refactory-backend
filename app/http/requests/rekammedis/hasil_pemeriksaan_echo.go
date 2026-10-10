package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilPemeriksaanEchoData isian hasil pemeriksaan echo.
type HasilPemeriksaanEchoData struct {
	Tanggal          string `form:"tanggal" json:"tanggal"`
	KdDokter         string `form:"kd_dokter" json:"kd_dokter"`
	Sistolik         string `form:"sistolik" json:"sistolik"`
	Diastolic        string `form:"diastolic" json:"diastolic"`
	Kontraktilitas   string `form:"kontraktilitas" json:"kontraktilitas"`
	DimensiRuang     string `form:"dimensi_ruang" json:"dimensi_ruang"`
	Katup            string `form:"katup" json:"katup"`
	AnalisaSegmental string `form:"analisa_segmental" json:"analisa_segmental"`
	Erap             string `form:"erap" json:"erap"`
	LainLain         string `form:"lain_lain" json:"lain_lain"`
	Kesimpulan       string `form:"kesimpulan" json:"kesimpulan"`
}

func hasilPemeriksaanEchoRules() map[string]any {
	rules := map[string]any{
		"tanggal":           "required|date",
		"kd_dokter":         "required|string|max_len:20",
		"sistolik":          "string|max_len:30",
		"diastolic":         "string|max_len:30",
		"kontraktilitas":    "string|max_len:30",
		"dimensi_ruang":     "string|max_len:50",
		"katup":             "string|max_len:50",
		"analisa_segmental": "string|max_len:100",
		"erap":              "string|max_len:15",
		"lain_lain":         "string|max_len:100",
		"kesimpulan":        "string|max_len:200",
	}
	return rules
}

// HasilPemeriksaanEchoStore simpan hasil pemeriksaan echo; kolom waktu kunci kosong = sekarang.
type HasilPemeriksaanEchoStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	HasilPemeriksaanEchoData
}

func (r *HasilPemeriksaanEchoStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanEchoStore) Rules(ctx http.Context) map[string]any {
	rules := hasilPemeriksaanEchoRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *HasilPemeriksaanEchoStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *HasilPemeriksaanEchoStore) Payload() HasilPemeriksaanEchoData {
	return r.HasilPemeriksaanEchoData
}

func (r *HasilPemeriksaanEchoStore) DetailValues() map[string][]string { return nil }

// HasilPemeriksaanEchoUpdate ubah hasil pemeriksaan echo (PUT); kunci lewat query string.
type HasilPemeriksaanEchoUpdate struct {
	HasilPemeriksaanEchoData
}

func (r *HasilPemeriksaanEchoUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanEchoUpdate) Rules(ctx http.Context) map[string]any {
	return hasilPemeriksaanEchoRules()
}

func (r *HasilPemeriksaanEchoUpdate) Payload() HasilPemeriksaanEchoData {
	return r.HasilPemeriksaanEchoData
}

func (r *HasilPemeriksaanEchoUpdate) DetailValues() map[string][]string { return nil }
