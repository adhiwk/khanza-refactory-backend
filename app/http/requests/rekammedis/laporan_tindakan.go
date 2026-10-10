package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// LaporanTindakanData isian laporan tindakan.
type LaporanTindakanData struct {
	KdDokter              string `form:"kd_dokter" json:"kd_dokter"`
	Nip                   string `form:"nip" json:"nip"`
	DiagnosaPraTindakan   string `form:"diagnosa_pra_tindakan" json:"diagnosa_pra_tindakan"`
	DiagnosaPascaTindakan string `form:"diagnosa_pasca_tindakan" json:"diagnosa_pasca_tindakan"`
	TindakanMedik         string `form:"tindakan_medik" json:"tindakan_medik"`
	Uraian                string `form:"uraian" json:"uraian"`
	Hasil                 string `form:"hasil" json:"hasil"`
	Kesimpulan            string `form:"kesimpulan" json:"kesimpulan"`
}

func laporanTindakanRules() map[string]any {
	rules := map[string]any{
		"kd_dokter":               "required|string|max_len:20",
		"nip":                     "string|max_len:20",
		"diagnosa_pra_tindakan":   "string|max_len:50",
		"diagnosa_pasca_tindakan": "string|max_len:50",
		"tindakan_medik":          "string|max_len:300",
		"uraian":                  "string|max_len:3000",
		"hasil":                   "string|max_len:1000",
		"kesimpulan":              "string|max_len:500",
	}
	return rules
}

// LaporanTindakanStore simpan laporan tindakan; kolom waktu kunci kosong = sekarang.
type LaporanTindakanStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	LaporanTindakanData
}

func (r *LaporanTindakanStore) Authorize(ctx http.Context) error { return nil }

func (r *LaporanTindakanStore) Rules(ctx http.Context) map[string]any {
	rules := laporanTindakanRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *LaporanTindakanStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *LaporanTindakanStore) Payload() LaporanTindakanData { return r.LaporanTindakanData }

func (r *LaporanTindakanStore) DetailValues() map[string][]string { return nil }

// LaporanTindakanUpdate ubah laporan tindakan (PUT); kunci lewat query string.
type LaporanTindakanUpdate struct {
	LaporanTindakanData
}

func (r *LaporanTindakanUpdate) Authorize(ctx http.Context) error { return nil }

func (r *LaporanTindakanUpdate) Rules(ctx http.Context) map[string]any { return laporanTindakanRules() }

func (r *LaporanTindakanUpdate) Payload() LaporanTindakanData { return r.LaporanTindakanData }

func (r *LaporanTindakanUpdate) DetailValues() map[string][]string { return nil }
