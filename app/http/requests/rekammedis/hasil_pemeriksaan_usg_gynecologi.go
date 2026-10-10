package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilPemeriksaanUsgGynecologiData isian hasil pemeriksaan USG gynecologi.
type HasilPemeriksaanUsgGynecologiData struct {
	Tanggal        string `form:"tanggal" json:"tanggal"`
	KdDokter       string `form:"kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis string `form:"diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari    string `form:"kiriman_dari" json:"kiriman_dari"`
	Uterus         string `form:"uterus" json:"uterus"`
	Parametrium    string `form:"parametrium" json:"parametrium"`
	Ovarium        string `form:"ovarium" json:"ovarium"`
	Doppler        string `form:"doppler" json:"doppler"`
	Kesimpulan     string `form:"kesimpulan" json:"kesimpulan"`
}

func hasilPemeriksaanUsgGynecologiRules() map[string]any {
	rules := map[string]any{
		"tanggal":         "required|date",
		"kd_dokter":       "required|string|max_len:20",
		"diagnosa_klinis": "string|max_len:50",
		"kiriman_dari":    "string|max_len:50",
		"uterus":          "string|max_len:200",
		"parametrium":     "string|max_len:200",
		"ovarium":         "string|max_len:200",
		"doppler":         "string|max_len:200",
		"kesimpulan":      "string|max_len:300",
	}
	return rules
}

// HasilPemeriksaanUsgGynecologiStore simpan hasil pemeriksaan USG gynecologi; kolom waktu kunci kosong = sekarang.
type HasilPemeriksaanUsgGynecologiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	HasilPemeriksaanUsgGynecologiData
}

func (r *HasilPemeriksaanUsgGynecologiStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanUsgGynecologiStore) Rules(ctx http.Context) map[string]any {
	rules := hasilPemeriksaanUsgGynecologiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *HasilPemeriksaanUsgGynecologiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *HasilPemeriksaanUsgGynecologiStore) Payload() HasilPemeriksaanUsgGynecologiData {
	return r.HasilPemeriksaanUsgGynecologiData
}

func (r *HasilPemeriksaanUsgGynecologiStore) DetailValues() map[string][]string { return nil }

// HasilPemeriksaanUsgGynecologiUpdate ubah hasil pemeriksaan USG gynecologi (PUT); kunci lewat query string.
type HasilPemeriksaanUsgGynecologiUpdate struct {
	HasilPemeriksaanUsgGynecologiData
}

func (r *HasilPemeriksaanUsgGynecologiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanUsgGynecologiUpdate) Rules(ctx http.Context) map[string]any {
	return hasilPemeriksaanUsgGynecologiRules()
}

func (r *HasilPemeriksaanUsgGynecologiUpdate) Payload() HasilPemeriksaanUsgGynecologiData {
	return r.HasilPemeriksaanUsgGynecologiData
}

func (r *HasilPemeriksaanUsgGynecologiUpdate) DetailValues() map[string][]string { return nil }
