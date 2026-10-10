package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// MppEvaluasiData isian skrining MPP form a.
type MppEvaluasiData struct {
	KdDokter     string `form:"kd_dokter" json:"kd_dokter"`
	KdKonsulan   string `form:"kd_konsulan" json:"kd_konsulan"`
	Diagnosis    string `form:"diagnosis" json:"diagnosis"`
	Kelompok     string `form:"kelompok" json:"kelompok"`
	Assesmen     string `form:"assesmen" json:"assesmen"`
	Identifikasi string `form:"identifikasi" json:"identifikasi"`
	Rencana      string `form:"rencana" json:"rencana"`
	Nip          string `form:"nip" json:"nip"`
}

func mppEvaluasiRules() map[string]any {
	rules := map[string]any{
		"kd_dokter":    "string|max_len:20",
		"kd_konsulan":  "string|max_len:20",
		"diagnosis":    "string|max_len:150",
		"kelompok":     "string|max_len:150",
		"assesmen":     "string|max_len:250",
		"identifikasi": "string|max_len:250",
		"rencana":      "string|max_len:2000",
		"nip":          "required|string|max_len:20",
	}
	return rules
}

// MppEvaluasiStore simpan skrining MPP form a; kolom waktu kunci kosong = sekarang.
type MppEvaluasiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	MppEvaluasiData
}

func (r *MppEvaluasiStore) Authorize(ctx http.Context) error { return nil }

func (r *MppEvaluasiStore) Rules(ctx http.Context) map[string]any {
	rules := mppEvaluasiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *MppEvaluasiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *MppEvaluasiStore) Payload() MppEvaluasiData { return r.MppEvaluasiData }

func (r *MppEvaluasiStore) DetailValues() map[string][]string { return nil }

// MppEvaluasiUpdate ubah skrining MPP form a (PUT); kunci lewat query string.
type MppEvaluasiUpdate struct {
	MppEvaluasiData
}

func (r *MppEvaluasiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *MppEvaluasiUpdate) Rules(ctx http.Context) map[string]any { return mppEvaluasiRules() }

func (r *MppEvaluasiUpdate) Payload() MppEvaluasiData { return r.MppEvaluasiData }

func (r *MppEvaluasiUpdate) DetailValues() map[string][]string { return nil }
