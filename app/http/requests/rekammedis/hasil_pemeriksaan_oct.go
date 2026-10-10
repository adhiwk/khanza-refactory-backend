package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilPemeriksaanOctData isian hasil pemeriksaan OCT.
type HasilPemeriksaanOctData struct {
	Tanggal          string `form:"tanggal" json:"tanggal"`
	KdDokter         string `form:"kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis   string `form:"diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari      string `form:"kiriman_dari" json:"kiriman_dari"`
	HasilPemeriksaan string `form:"hasil_pemeriksaan" json:"hasil_pemeriksaan"`
}

func hasilPemeriksaanOctRules() map[string]any {
	rules := map[string]any{
		"tanggal":           "required|date",
		"kd_dokter":         "required|string|max_len:20",
		"diagnosa_klinis":   "string|max_len:50",
		"kiriman_dari":      "string|max_len:50",
		"hasil_pemeriksaan": "string|max_len:1000",
	}
	return rules
}

// HasilPemeriksaanOctStore simpan hasil pemeriksaan OCT; kolom waktu kunci kosong = sekarang.
type HasilPemeriksaanOctStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	HasilPemeriksaanOctData
}

func (r *HasilPemeriksaanOctStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanOctStore) Rules(ctx http.Context) map[string]any {
	rules := hasilPemeriksaanOctRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *HasilPemeriksaanOctStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *HasilPemeriksaanOctStore) Payload() HasilPemeriksaanOctData {
	return r.HasilPemeriksaanOctData
}

func (r *HasilPemeriksaanOctStore) DetailValues() map[string][]string { return nil }

// HasilPemeriksaanOctUpdate ubah hasil pemeriksaan OCT (PUT); kunci lewat query string.
type HasilPemeriksaanOctUpdate struct {
	HasilPemeriksaanOctData
}

func (r *HasilPemeriksaanOctUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanOctUpdate) Rules(ctx http.Context) map[string]any {
	return hasilPemeriksaanOctRules()
}

func (r *HasilPemeriksaanOctUpdate) Payload() HasilPemeriksaanOctData {
	return r.HasilPemeriksaanOctData
}

func (r *HasilPemeriksaanOctUpdate) DetailValues() map[string][]string { return nil }
