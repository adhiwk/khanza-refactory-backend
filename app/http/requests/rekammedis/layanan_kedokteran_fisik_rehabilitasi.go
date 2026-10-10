package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// LayananKedokteranFisikRehabilitasiData isian layanan kedokteran fisik rehabilitasi.
type LayananKedokteranFisikRehabilitasiData struct {
	Tanggal                       string `form:"tanggal" json:"tanggal"`
	KdDokter                      string `form:"kd_dokter" json:"kd_dokter"`
	Pendamping                    string `form:"pendamping" json:"pendamping"`
	KeteranganPendamping          string `form:"keterangan_pendamping" json:"keterangan_pendamping"`
	Anamnesa                      string `form:"anamnesa" json:"anamnesa"`
	PemeriksaanFisik              string `form:"pemeriksaan_fisik" json:"pemeriksaan_fisik"`
	DiagnosaMedis                 string `form:"diagnosa_medis" json:"diagnosa_medis"`
	DiagnosaFungsi                string `form:"diagnosa_fungsi" json:"diagnosa_fungsi"`
	Tatalaksana                   string `form:"tatalaksana" json:"tatalaksana"`
	Anjuran                       string `form:"anjuran" json:"anjuran"`
	Evaluasi                      string `form:"evaluasi" json:"evaluasi"`
	SuspekPenyakitKerja           string `form:"suspek_penyakit_kerja" json:"suspek_penyakit_kerja"`
	KeteranganSuspekPenyakitKerja string `form:"keterangan_suspek_penyakit_kerja" json:"keterangan_suspek_penyakit_kerja"`
	StatusProgram                 string `form:"status_program" json:"status_program"`
}

func layananKedokteranFisikRehabilitasiRules() map[string]any {
	rules := map[string]any{
		"tanggal":                          "required|date",
		"kd_dokter":                        "required|string|max_len:20",
		"pendamping":                       "in:Tidak,Suami,Istri,Anak,Keluarga,Lainnya",
		"keterangan_pendamping":            "string|max_len:30",
		"anamnesa":                         "string|max_len:500",
		"pemeriksaan_fisik":                "string|max_len:1500",
		"diagnosa_medis":                   "string|max_len:200",
		"diagnosa_fungsi":                  "string|max_len:200",
		"tatalaksana":                      "string|max_len:2000",
		"anjuran":                          "string|max_len:500",
		"evaluasi":                         "string|max_len:500",
		"suspek_penyakit_kerja":            "in:Tidak,Ya",
		"keterangan_suspek_penyakit_kerja": "string|max_len:70",
		"status_program":                   "required|in:Belum Selesai,Sudah Selesai",
	}
	return rules
}

// LayananKedokteranFisikRehabilitasiStore simpan layanan kedokteran fisik rehabilitasi; kolom waktu kunci kosong = sekarang.
type LayananKedokteranFisikRehabilitasiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	LayananKedokteranFisikRehabilitasiData
}

func (r *LayananKedokteranFisikRehabilitasiStore) Authorize(ctx http.Context) error { return nil }

func (r *LayananKedokteranFisikRehabilitasiStore) Rules(ctx http.Context) map[string]any {
	rules := layananKedokteranFisikRehabilitasiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *LayananKedokteranFisikRehabilitasiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *LayananKedokteranFisikRehabilitasiStore) Payload() LayananKedokteranFisikRehabilitasiData {
	return r.LayananKedokteranFisikRehabilitasiData
}

func (r *LayananKedokteranFisikRehabilitasiStore) DetailValues() map[string][]string { return nil }

// LayananKedokteranFisikRehabilitasiUpdate ubah layanan kedokteran fisik rehabilitasi (PUT); kunci lewat query string.
type LayananKedokteranFisikRehabilitasiUpdate struct {
	LayananKedokteranFisikRehabilitasiData
}

func (r *LayananKedokteranFisikRehabilitasiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *LayananKedokteranFisikRehabilitasiUpdate) Rules(ctx http.Context) map[string]any {
	return layananKedokteranFisikRehabilitasiRules()
}

func (r *LayananKedokteranFisikRehabilitasiUpdate) Payload() LayananKedokteranFisikRehabilitasiData {
	return r.LayananKedokteranFisikRehabilitasiData
}

func (r *LayananKedokteranFisikRehabilitasiUpdate) DetailValues() map[string][]string { return nil }
