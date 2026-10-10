package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilPemeriksaanUsgUrologiData isian hasil pemeriksaan USG urologi.
type HasilPemeriksaanUsgUrologiData struct {
	Tanggal        string `form:"tanggal" json:"tanggal"`
	KdDokter       string `form:"kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis string `form:"diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari    string `form:"kiriman_dari" json:"kiriman_dari"`
	GinjalKanan    string `form:"ginjal_kanan" json:"ginjal_kanan"`
	GinjalKiri     string `form:"ginjal_kiri" json:"ginjal_kiri"`
	VesicaUrinaria string `form:"vesica_urinaria" json:"vesica_urinaria"`
	Tambahan       string `form:"tambahan" json:"tambahan"`
}

func hasilPemeriksaanUsgUrologiRules() map[string]any {
	rules := map[string]any{
		"tanggal":         "required|date",
		"kd_dokter":       "required|string|max_len:20",
		"diagnosa_klinis": "string|max_len:50",
		"kiriman_dari":    "string|max_len:50",
		"ginjal_kanan":    "string|max_len:200",
		"ginjal_kiri":     "string|max_len:200",
		"vesica_urinaria": "string|max_len:200",
		"tambahan":        "string|max_len:300",
	}
	return rules
}

// HasilPemeriksaanUsgUrologiStore simpan hasil pemeriksaan USG urologi; kolom waktu kunci kosong = sekarang.
type HasilPemeriksaanUsgUrologiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	HasilPemeriksaanUsgUrologiData
}

func (r *HasilPemeriksaanUsgUrologiStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanUsgUrologiStore) Rules(ctx http.Context) map[string]any {
	rules := hasilPemeriksaanUsgUrologiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *HasilPemeriksaanUsgUrologiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *HasilPemeriksaanUsgUrologiStore) Payload() HasilPemeriksaanUsgUrologiData {
	return r.HasilPemeriksaanUsgUrologiData
}

func (r *HasilPemeriksaanUsgUrologiStore) DetailValues() map[string][]string { return nil }

// HasilPemeriksaanUsgUrologiUpdate ubah hasil pemeriksaan USG urologi (PUT); kunci lewat query string.
type HasilPemeriksaanUsgUrologiUpdate struct {
	HasilPemeriksaanUsgUrologiData
}

func (r *HasilPemeriksaanUsgUrologiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanUsgUrologiUpdate) Rules(ctx http.Context) map[string]any {
	return hasilPemeriksaanUsgUrologiRules()
}

func (r *HasilPemeriksaanUsgUrologiUpdate) Payload() HasilPemeriksaanUsgUrologiData {
	return r.HasilPemeriksaanUsgUrologiData
}

func (r *HasilPemeriksaanUsgUrologiUpdate) DetailValues() map[string][]string { return nil }
