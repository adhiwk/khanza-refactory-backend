package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilPemeriksaanSlitLampData isian hasil pemeriksaan slit lamp.
type HasilPemeriksaanSlitLampData struct {
	Tanggal          string `form:"tanggal" json:"tanggal"`
	KdDokter         string `form:"kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis   string `form:"diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari      string `form:"kiriman_dari" json:"kiriman_dari"`
	HasilPemeriksaan string `form:"hasil_pemeriksaan" json:"hasil_pemeriksaan"`
}

func hasilPemeriksaanSlitLampRules() map[string]any {
	rules := map[string]any{
		"tanggal":           "required|date",
		"kd_dokter":         "required|string|max_len:20",
		"diagnosa_klinis":   "string|max_len:50",
		"kiriman_dari":      "string|max_len:50",
		"hasil_pemeriksaan": "string|max_len:1000",
	}
	return rules
}

// HasilPemeriksaanSlitLampStore simpan hasil pemeriksaan slit lamp; kolom waktu kunci kosong = sekarang.
type HasilPemeriksaanSlitLampStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	HasilPemeriksaanSlitLampData
}

func (r *HasilPemeriksaanSlitLampStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanSlitLampStore) Rules(ctx http.Context) map[string]any {
	rules := hasilPemeriksaanSlitLampRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *HasilPemeriksaanSlitLampStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *HasilPemeriksaanSlitLampStore) Payload() HasilPemeriksaanSlitLampData {
	return r.HasilPemeriksaanSlitLampData
}

func (r *HasilPemeriksaanSlitLampStore) DetailValues() map[string][]string { return nil }

// HasilPemeriksaanSlitLampUpdate ubah hasil pemeriksaan slit lamp (PUT); kunci lewat query string.
type HasilPemeriksaanSlitLampUpdate struct {
	HasilPemeriksaanSlitLampData
}

func (r *HasilPemeriksaanSlitLampUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanSlitLampUpdate) Rules(ctx http.Context) map[string]any {
	return hasilPemeriksaanSlitLampRules()
}

func (r *HasilPemeriksaanSlitLampUpdate) Payload() HasilPemeriksaanSlitLampData {
	return r.HasilPemeriksaanSlitLampData
}

func (r *HasilPemeriksaanSlitLampUpdate) DetailValues() map[string][]string { return nil }
