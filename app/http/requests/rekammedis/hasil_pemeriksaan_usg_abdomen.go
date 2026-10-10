package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilPemeriksaanUsgAbdomenData isian hasil pemeriksaan USG abdomen.
type HasilPemeriksaanUsgAbdomenData struct {
	Tanggal        string `form:"tanggal" json:"tanggal"`
	KdDokter       string `form:"kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis string `form:"diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari    string `form:"kiriman_dari" json:"kiriman_dari"`
	Esofagus       string `form:"esofagus" json:"esofagus"`
	Colon          string `form:"colon" json:"colon"`
	Gaster         string `form:"gaster" json:"gaster"`
	Hepar          string `form:"hepar" json:"hepar"`
	GallBlader     string `form:"gall_blader" json:"gall_blader"`
	Lien           string `form:"lien" json:"lien"`
	Pancreas       string `form:"pancreas" json:"pancreas"`
	GinjalDextra   string `form:"ginjal_dextra" json:"ginjal_dextra"`
	GinjalSinistra string `form:"ginjal_sinistra" json:"ginjal_sinistra"`
	Kesimpulan     string `form:"kesimpulan" json:"kesimpulan"`
}

func hasilPemeriksaanUsgAbdomenRules() map[string]any {
	rules := map[string]any{
		"tanggal":         "required|date",
		"kd_dokter":       "required|string|max_len:20",
		"diagnosa_klinis": "string|max_len:50",
		"kiriman_dari":    "string|max_len:50",
		"esofagus":        "string|max_len:200",
		"colon":           "string|max_len:200",
		"gaster":          "string|max_len:200",
		"hepar":           "string|max_len:200",
		"gall_blader":     "string|max_len:200",
		"lien":            "string|max_len:200",
		"pancreas":        "string|max_len:200",
		"ginjal_dextra":   "string|max_len:200",
		"ginjal_sinistra": "string|max_len:200",
		"kesimpulan":      "string|max_len:300",
	}
	return rules
}

// HasilPemeriksaanUsgAbdomenStore simpan hasil pemeriksaan USG abdomen; kolom waktu kunci kosong = sekarang.
type HasilPemeriksaanUsgAbdomenStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	HasilPemeriksaanUsgAbdomenData
}

func (r *HasilPemeriksaanUsgAbdomenStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanUsgAbdomenStore) Rules(ctx http.Context) map[string]any {
	rules := hasilPemeriksaanUsgAbdomenRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *HasilPemeriksaanUsgAbdomenStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *HasilPemeriksaanUsgAbdomenStore) Payload() HasilPemeriksaanUsgAbdomenData {
	return r.HasilPemeriksaanUsgAbdomenData
}

func (r *HasilPemeriksaanUsgAbdomenStore) DetailValues() map[string][]string { return nil }

// HasilPemeriksaanUsgAbdomenUpdate ubah hasil pemeriksaan USG abdomen (PUT); kunci lewat query string.
type HasilPemeriksaanUsgAbdomenUpdate struct {
	HasilPemeriksaanUsgAbdomenData
}

func (r *HasilPemeriksaanUsgAbdomenUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanUsgAbdomenUpdate) Rules(ctx http.Context) map[string]any {
	return hasilPemeriksaanUsgAbdomenRules()
}

func (r *HasilPemeriksaanUsgAbdomenUpdate) Payload() HasilPemeriksaanUsgAbdomenData {
	return r.HasilPemeriksaanUsgAbdomenData
}

func (r *HasilPemeriksaanUsgAbdomenUpdate) DetailValues() map[string][]string { return nil }
