package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRalanThtData isian penilaian awal medis ralan THT.
type PenilaianMedisRalanThtData struct {
	Tanggal          string `form:"tanggal" json:"tanggal"`
	KdDokter         string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis        string `form:"anamnesis" json:"anamnesis"`
	Hubungan         string `form:"hubungan" json:"hubungan"`
	KeluhanUtama     string `form:"keluhan_utama" json:"keluhan_utama"`
	Rps              string `form:"rps" json:"rps"`
	Rpd              string `form:"rpd" json:"rpd"`
	Rpo              string `form:"rpo" json:"rpo"`
	Alergi           string `form:"alergi" json:"alergi"`
	Td               string `form:"td" json:"td"`
	Nadi             string `form:"nadi" json:"nadi"`
	Rr               string `form:"rr" json:"rr"`
	Suhu             string `form:"suhu" json:"suhu"`
	Bb               string `form:"bb" json:"bb"`
	Tb               string `form:"tb" json:"tb"`
	Nyeri            string `form:"nyeri" json:"nyeri"`
	StatusNutrisi    string `form:"status_nutrisi" json:"status_nutrisi"`
	Kondisi          string `form:"kondisi" json:"kondisi"`
	KetLokalis       string `form:"ket_lokalis" json:"ket_lokalis"`
	Lab              string `form:"lab" json:"lab"`
	Rad              string `form:"rad" json:"rad"`
	TesPendengaran   string `form:"tes_pendengaran" json:"tes_pendengaran"`
	Penunjang        string `form:"penunjang" json:"penunjang"`
	Diagnosis        string `form:"diagnosis" json:"diagnosis"`
	Diagnosisbanding string `form:"diagnosisbanding" json:"diagnosisbanding"`
	Permasalahan     string `form:"permasalahan" json:"permasalahan"`
	Terapi           string `form:"terapi" json:"terapi"`
	Tindakan         string `form:"tindakan" json:"tindakan"`
	Tatalaksana      string `form:"tatalaksana" json:"tatalaksana"`
	Edukasi          string `form:"edukasi" json:"edukasi"`
}

func penilaianMedisRalanThtRules() map[string]any {
	rules := map[string]any{
		"tanggal":          "required|date",
		"kd_dokter":        "required|string|max_len:20",
		"anamnesis":        "required|in:Autoanamnesis,Alloanamnesis",
		"hubungan":         "string|max_len:30",
		"keluhan_utama":    "string|max_len:2000",
		"rps":              "string|max_len:2000",
		"rpd":              "string|max_len:1000",
		"rpo":              "string|max_len:1000",
		"alergi":           "string|max_len:50",
		"td":               "string|max_len:8",
		"nadi":             "string|max_len:5",
		"rr":               "string|max_len:5",
		"suhu":             "string|max_len:5",
		"bb":               "string|max_len:5",
		"tb":               "string|max_len:5",
		"nyeri":            "string|max_len:50",
		"status_nutrisi":   "string|max_len:50",
		"kondisi":          "string",
		"ket_lokalis":      "string",
		"lab":              "string",
		"rad":              "string",
		"tes_pendengaran":  "string",
		"penunjang":        "string",
		"diagnosis":        "string|max_len:500",
		"diagnosisbanding": "string|max_len:500",
		"permasalahan":     "string",
		"terapi":           "string",
		"tindakan":         "string",
		"tatalaksana":      "string",
		"edukasi":          "string|max_len:1000",
	}
	return rules
}

// PenilaianMedisRalanThtStore simpan penilaian awal medis ralan THT; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRalanThtStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRalanThtData
}

func (r *PenilaianMedisRalanThtStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanThtStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRalanThtRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRalanThtStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianMedisRalanThtStore) Payload() PenilaianMedisRalanThtData {
	return r.PenilaianMedisRalanThtData
}

func (r *PenilaianMedisRalanThtStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisRalanThtUpdate ubah penilaian awal medis ralan THT (PUT); kunci lewat query string.
type PenilaianMedisRalanThtUpdate struct {
	PenilaianMedisRalanThtData
}

func (r *PenilaianMedisRalanThtUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanThtUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRalanThtRules()
}

func (r *PenilaianMedisRalanThtUpdate) Payload() PenilaianMedisRalanThtData {
	return r.PenilaianMedisRalanThtData
}

func (r *PenilaianMedisRalanThtUpdate) DetailValues() map[string][]string { return nil }
