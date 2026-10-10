package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilPemeriksaanTreadmillData isian hasil pemeriksaan treadmill.
type HasilPemeriksaanTreadmillData struct {
	Tanggal               string `form:"tanggal" json:"tanggal"`
	KdDokter              string `form:"kd_dokter" json:"kd_dokter"`
	KirimanDari           string `form:"kiriman_dari" json:"kiriman_dari"`
	DiagnosaKlinis        string `form:"diagnosa_klinis" json:"diagnosa_klinis"`
	Protokol              string `form:"protokol" json:"protokol"`
	KeteranganProtokol    string `form:"keterangan_protokol" json:"keterangan_protokol"`
	TdAwal                string `form:"td_awal" json:"td_awal"`
	NadiAwal              string `form:"nadi_awal" json:"nadi_awal"`
	DenyutJantungMaksimal string `form:"denyut_jantung_maksimal" json:"denyut_jantung_maksimal"`
	HasilPemeriksaan      string `form:"hasil_pemeriksaan" json:"hasil_pemeriksaan"`
	TemuanEkg             string `form:"temuan_ekg" json:"temuan_ekg"`
	KapasitasFungsional   string `form:"kapasitas_fungsional" json:"kapasitas_fungsional"`
	Interpretasi          string `form:"interpretasi" json:"interpretasi"`
	Kesimpulan            string `form:"kesimpulan" json:"kesimpulan"`
}

func hasilPemeriksaanTreadmillRules() map[string]any {
	rules := map[string]any{
		"tanggal":                 "required|date",
		"kd_dokter":               "required|string|max_len:20",
		"kiriman_dari":            "string|max_len:50",
		"diagnosa_klinis":         "string|max_len:50",
		"protokol":                "in:Bruce,Modified Bruce,Balke,Naughton,Lainnya",
		"keterangan_protokol":     "string|max_len:30",
		"td_awal":                 "string|max_len:8",
		"nadi_awal":               "string|max_len:5",
		"denyut_jantung_maksimal": "string|max_len:5",
		"hasil_pemeriksaan":       "string|max_len:1000",
		"temuan_ekg":              "string|max_len:200",
		"kapasitas_fungsional":    "string|max_len:200",
		"interpretasi":            "string|max_len:200",
		"kesimpulan":              "string|max_len:300",
	}
	return rules
}

// HasilPemeriksaanTreadmillStore simpan hasil pemeriksaan treadmill; kolom waktu kunci kosong = sekarang.
type HasilPemeriksaanTreadmillStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	HasilPemeriksaanTreadmillData
}

func (r *HasilPemeriksaanTreadmillStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanTreadmillStore) Rules(ctx http.Context) map[string]any {
	rules := hasilPemeriksaanTreadmillRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *HasilPemeriksaanTreadmillStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *HasilPemeriksaanTreadmillStore) Payload() HasilPemeriksaanTreadmillData {
	return r.HasilPemeriksaanTreadmillData
}

func (r *HasilPemeriksaanTreadmillStore) DetailValues() map[string][]string { return nil }

// HasilPemeriksaanTreadmillUpdate ubah hasil pemeriksaan treadmill (PUT); kunci lewat query string.
type HasilPemeriksaanTreadmillUpdate struct {
	HasilPemeriksaanTreadmillData
}

func (r *HasilPemeriksaanTreadmillUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanTreadmillUpdate) Rules(ctx http.Context) map[string]any {
	return hasilPemeriksaanTreadmillRules()
}

func (r *HasilPemeriksaanTreadmillUpdate) Payload() HasilPemeriksaanTreadmillData {
	return r.HasilPemeriksaanTreadmillData
}

func (r *HasilPemeriksaanTreadmillUpdate) DetailValues() map[string][]string { return nil }
