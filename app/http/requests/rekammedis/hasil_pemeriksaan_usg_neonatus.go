package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilPemeriksaanUsgNeonatusData isian hasil pemeriksaan USG neonatus.
type HasilPemeriksaanUsgNeonatusData struct {
	Tanggal           string `form:"tanggal" json:"tanggal"`
	KdDokter          string `form:"kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis    string `form:"diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari       string `form:"kiriman_dari" json:"kiriman_dari"`
	VentrikalSinistra string `form:"ventrikal_sinistra" json:"ventrikal_sinistra"`
	VentrikalDextra   string `form:"ventrikal_dextra" json:"ventrikal_dextra"`
	Kesan             string `form:"kesan" json:"kesan"`
	Kesimpulan        string `form:"kesimpulan" json:"kesimpulan"`
	Saran             string `form:"saran" json:"saran"`
}

func hasilPemeriksaanUsgNeonatusRules() map[string]any {
	rules := map[string]any{
		"tanggal":            "required|date",
		"kd_dokter":          "required|string|max_len:20",
		"diagnosa_klinis":    "string|max_len:50",
		"kiriman_dari":       "string|max_len:50",
		"ventrikal_sinistra": "string|max_len:200",
		"ventrikal_dextra":   "string|max_len:200",
		"kesan":              "string|max_len:200",
		"kesimpulan":         "string|max_len:300",
		"saran":              "string|max_len:200",
	}
	return rules
}

// HasilPemeriksaanUsgNeonatusStore simpan hasil pemeriksaan USG neonatus; kolom waktu kunci kosong = sekarang.
type HasilPemeriksaanUsgNeonatusStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	HasilPemeriksaanUsgNeonatusData
}

func (r *HasilPemeriksaanUsgNeonatusStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanUsgNeonatusStore) Rules(ctx http.Context) map[string]any {
	rules := hasilPemeriksaanUsgNeonatusRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *HasilPemeriksaanUsgNeonatusStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *HasilPemeriksaanUsgNeonatusStore) Payload() HasilPemeriksaanUsgNeonatusData {
	return r.HasilPemeriksaanUsgNeonatusData
}

func (r *HasilPemeriksaanUsgNeonatusStore) DetailValues() map[string][]string { return nil }

// HasilPemeriksaanUsgNeonatusUpdate ubah hasil pemeriksaan USG neonatus (PUT); kunci lewat query string.
type HasilPemeriksaanUsgNeonatusUpdate struct {
	HasilPemeriksaanUsgNeonatusData
}

func (r *HasilPemeriksaanUsgNeonatusUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanUsgNeonatusUpdate) Rules(ctx http.Context) map[string]any {
	return hasilPemeriksaanUsgNeonatusRules()
}

func (r *HasilPemeriksaanUsgNeonatusUpdate) Payload() HasilPemeriksaanUsgNeonatusData {
	return r.HasilPemeriksaanUsgNeonatusData
}

func (r *HasilPemeriksaanUsgNeonatusUpdate) DetailValues() map[string][]string { return nil }
