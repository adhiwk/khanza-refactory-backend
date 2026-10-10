package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRalanBedahData isian penilaian awal medis ralan bedah.
type PenilaianMedisRalanBedahData struct {
	Tanggal      string `form:"tanggal" json:"tanggal"`
	KdDokter     string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis    string `form:"anamnesis" json:"anamnesis"`
	Hubungan     string `form:"hubungan" json:"hubungan"`
	KeluhanUtama string `form:"keluhan_utama" json:"keluhan_utama"`
	Rps          string `form:"rps" json:"rps"`
	Rpd          string `form:"rpd" json:"rpd"`
	Rpo          string `form:"rpo" json:"rpo"`
	Alergi       string `form:"alergi" json:"alergi"`
	Kesadaran    string `form:"kesadaran" json:"kesadaran"`
	Status       string `form:"status" json:"status"`
	Td           string `form:"td" json:"td"`
	Nadi         string `form:"nadi" json:"nadi"`
	Suhu         string `form:"suhu" json:"suhu"`
	Rr           string `form:"rr" json:"rr"`
	Bb           string `form:"bb" json:"bb"`
	Nyeri        string `form:"nyeri" json:"nyeri"`
	Gcs          string `form:"gcs" json:"gcs"`
	Kepala       string `form:"kepala" json:"kepala"`
	Thoraks      string `form:"thoraks" json:"thoraks"`
	Abdomen      string `form:"abdomen" json:"abdomen"`
	Ekstremitas  string `form:"ekstremitas" json:"ekstremitas"`
	Genetalia    string `form:"genetalia" json:"genetalia"`
	Columna      string `form:"columna" json:"columna"`
	Muskulos     string `form:"muskulos" json:"muskulos"`
	Lainnya      string `form:"lainnya" json:"lainnya"`
	KetLokalis   string `form:"ket_lokalis" json:"ket_lokalis"`
	Lab          string `form:"lab" json:"lab"`
	Rad          string `form:"rad" json:"rad"`
	Pemeriksaan  string `form:"pemeriksaan" json:"pemeriksaan"`
	Diagnosis    string `form:"diagnosis" json:"diagnosis"`
	Diagnosis2   string `form:"diagnosis2" json:"diagnosis2"`
	Permasalahan string `form:"permasalahan" json:"permasalahan"`
	Terapi       string `form:"terapi" json:"terapi"`
	Tindakan     string `form:"tindakan" json:"tindakan"`
	Edukasi      string `form:"edukasi" json:"edukasi"`
}

func penilaianMedisRalanBedahRules() map[string]any {
	rules := map[string]any{
		"tanggal":       "required|date",
		"kd_dokter":     "required|string|max_len:20",
		"anamnesis":     "required|in:Autoanamnesis,Alloanamnesis",
		"hubungan":      "string|max_len:30",
		"keluhan_utama": "string|max_len:2000",
		"rps":           "string|max_len:2000",
		"rpd":           "string|max_len:1000",
		"rpo":           "string|max_len:1000",
		"alergi":        "string|max_len:50",
		"kesadaran":     "required|in:Compos Mentis,Apatis,Delirum",
		"status":        "string|max_len:50",
		"td":            "string|max_len:8",
		"nadi":          "string|max_len:5",
		"suhu":          "string|max_len:5",
		"rr":            "string|max_len:5",
		"bb":            "string|max_len:5",
		"nyeri":         "string|max_len:5",
		"gcs":           "string|max_len:10",
		"kepala":        "required|in:Normal,Abnormal,Tidak Diperiksa",
		"thoraks":       "required|in:Normal,Abnormal,Tidak Diperiksa",
		"abdomen":       "required|in:Normal,Abnormal,Tidak Diperiksa",
		"ekstremitas":   "required|in:Normal,Abnormal,Tidak Diperiksa",
		"genetalia":     "required|in:Normal,Abnormal,Tidak Diperiksa",
		"columna":       "required|in:Normal,Abnormal,Tidak Diperiksa",
		"muskulos":      "required|in:Normal,Abnormal,Tidak Diperiksa",
		"lainnya":       "string|max_len:1000",
		"ket_lokalis":   "string",
		"lab":           "string|max_len:500",
		"rad":           "string|max_len:500",
		"pemeriksaan":   "string|max_len:500",
		"diagnosis":     "string|max_len:500",
		"diagnosis2":    "string|max_len:500",
		"permasalahan":  "string|max_len:500",
		"terapi":        "string|max_len:500",
		"tindakan":      "string|max_len:500",
		"edukasi":       "string|max_len:500",
	}
	return rules
}

// PenilaianMedisRalanBedahStore simpan penilaian awal medis ralan bedah; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRalanBedahStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRalanBedahData
}

func (r *PenilaianMedisRalanBedahStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanBedahStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRalanBedahRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRalanBedahStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianMedisRalanBedahStore) Payload() PenilaianMedisRalanBedahData {
	return r.PenilaianMedisRalanBedahData
}

func (r *PenilaianMedisRalanBedahStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisRalanBedahUpdate ubah penilaian awal medis ralan bedah (PUT); kunci lewat query string.
type PenilaianMedisRalanBedahUpdate struct {
	PenilaianMedisRalanBedahData
}

func (r *PenilaianMedisRalanBedahUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanBedahUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRalanBedahRules()
}

func (r *PenilaianMedisRalanBedahUpdate) Payload() PenilaianMedisRalanBedahData {
	return r.PenilaianMedisRalanBedahData
}

func (r *PenilaianMedisRalanBedahUpdate) DetailValues() map[string][]string { return nil }
