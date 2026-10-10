package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilPemeriksaanEkgData isian hasil pemeriksaan EKG.
type HasilPemeriksaanEkgData struct {
	Tanggal        string `form:"tanggal" json:"tanggal"`
	KdDokter       string `form:"kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis string `form:"diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari    string `form:"kiriman_dari" json:"kiriman_dari"`
	Irama          string `form:"irama" json:"irama"`
	LajuJantung    string `form:"laju_jantung" json:"laju_jantung"`
	Gelombangp     string `form:"gelombangp" json:"gelombangp"`
	Intervalpr     string `form:"intervalpr" json:"intervalpr"`
	Axis           string `form:"axis" json:"axis"`
	Kompleksqrs    string `form:"kompleksqrs" json:"kompleksqrs"`
	Segmenst       string `form:"segmenst" json:"segmenst"`
	Gelombangt     string `form:"gelombangt" json:"gelombangt"`
	Kesimpulan     string `form:"kesimpulan" json:"kesimpulan"`
}

func hasilPemeriksaanEkgRules() map[string]any {
	rules := map[string]any{
		"tanggal":         "required|date",
		"kd_dokter":       "required|string|max_len:20",
		"diagnosa_klinis": "string|max_len:50",
		"kiriman_dari":    "string|max_len:50",
		"irama":           "string|max_len:40",
		"laju_jantung":    "string|max_len:40",
		"gelombangp":      "string|max_len:40",
		"intervalpr":      "string|max_len:40",
		"axis":            "string|max_len:40",
		"kompleksqrs":     "string|max_len:40",
		"segmenst":        "in:Normal,Tidak Normal",
		"gelombangt":      "in:Normal,Tidak Normal",
		"kesimpulan":      "string|max_len:200",
	}
	return rules
}

// HasilPemeriksaanEkgStore simpan hasil pemeriksaan EKG; kolom waktu kunci kosong = sekarang.
type HasilPemeriksaanEkgStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	HasilPemeriksaanEkgData
}

func (r *HasilPemeriksaanEkgStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanEkgStore) Rules(ctx http.Context) map[string]any {
	rules := hasilPemeriksaanEkgRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *HasilPemeriksaanEkgStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *HasilPemeriksaanEkgStore) Payload() HasilPemeriksaanEkgData {
	return r.HasilPemeriksaanEkgData
}

func (r *HasilPemeriksaanEkgStore) DetailValues() map[string][]string { return nil }

// HasilPemeriksaanEkgUpdate ubah hasil pemeriksaan EKG (PUT); kunci lewat query string.
type HasilPemeriksaanEkgUpdate struct {
	HasilPemeriksaanEkgData
}

func (r *HasilPemeriksaanEkgUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanEkgUpdate) Rules(ctx http.Context) map[string]any {
	return hasilPemeriksaanEkgRules()
}

func (r *HasilPemeriksaanEkgUpdate) Payload() HasilPemeriksaanEkgData {
	return r.HasilPemeriksaanEkgData
}

func (r *HasilPemeriksaanEkgUpdate) DetailValues() map[string][]string { return nil }
