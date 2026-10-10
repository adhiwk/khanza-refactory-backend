package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRalanKulitdankelaminData isian penilaian awal medis ralan kulit dan kelamin.
type PenilaianMedisRalanKulitdankelaminData struct {
	Tanggal      string `form:"tanggal" json:"tanggal"`
	KdDokter     string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis    string `form:"anamnesis" json:"anamnesis"`
	Hubungan     string `form:"hubungan" json:"hubungan"`
	KeluhanUtama string `form:"keluhan_utama" json:"keluhan_utama"`
	Rps          string `form:"rps" json:"rps"`
	Rpd          string `form:"rpd" json:"rpd"`
	Rpo          string `form:"rpo" json:"rpo"`
	Rpk          string `form:"rpk" json:"rpk"`
	Kesadaran    string `form:"kesadaran" json:"kesadaran"`
	Status       string `form:"status" json:"status"`
	Td           string `form:"td" json:"td"`
	Nadi         string `form:"nadi" json:"nadi"`
	Suhu         string `form:"suhu" json:"suhu"`
	Rr           string `form:"rr" json:"rr"`
	Bb           string `form:"bb" json:"bb"`
	Nyeri        string `form:"nyeri" json:"nyeri"`
	Gcs          string `form:"gcs" json:"gcs"`
	Statusderma  string `form:"statusderma" json:"statusderma"`
	Pemeriksaan  string `form:"pemeriksaan" json:"pemeriksaan"`
	Diagnosis    string `form:"diagnosis" json:"diagnosis"`
	Diagnosis2   string `form:"diagnosis2" json:"diagnosis2"`
	Permasalahan string `form:"permasalahan" json:"permasalahan"`
	Terapi       string `form:"terapi" json:"terapi"`
	Tindakan     string `form:"tindakan" json:"tindakan"`
	Edukasi      string `form:"edukasi" json:"edukasi"`
}

func penilaianMedisRalanKulitdankelaminRules() map[string]any {
	rules := map[string]any{
		"tanggal":       "required|date",
		"kd_dokter":     "required|string|max_len:20",
		"anamnesis":     "required|in:Autoanamnesis,Alloanamnesis",
		"hubungan":      "string|max_len:30",
		"keluhan_utama": "string|max_len:2000",
		"rps":           "string|max_len:2000",
		"rpd":           "string|max_len:1000",
		"rpo":           "string|max_len:1000",
		"rpk":           "string|max_len:50",
		"kesadaran":     "required|in:Compos Mentis,Apatis,Delirum",
		"status":        "string|max_len:15",
		"td":            "string|max_len:8",
		"nadi":          "string|max_len:5",
		"suhu":          "string|max_len:5",
		"rr":            "string|max_len:5",
		"bb":            "string|max_len:5",
		"nyeri":         "string|max_len:50",
		"gcs":           "string|max_len:10",
		"statusderma":   "string|max_len:1000",
		"pemeriksaan":   "string|max_len:100",
		"diagnosis":     "string|max_len:500",
		"diagnosis2":    "string|max_len:500",
		"permasalahan":  "string|max_len:500",
		"terapi":        "string|max_len:500",
		"tindakan":      "string|max_len:100",
		"edukasi":       "string|max_len:500",
	}
	return rules
}

// PenilaianMedisRalanKulitdankelaminStore simpan penilaian awal medis ralan kulit dan kelamin; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRalanKulitdankelaminStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRalanKulitdankelaminData
}

func (r *PenilaianMedisRalanKulitdankelaminStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanKulitdankelaminStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRalanKulitdankelaminRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRalanKulitdankelaminStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianMedisRalanKulitdankelaminStore) Payload() PenilaianMedisRalanKulitdankelaminData {
	return r.PenilaianMedisRalanKulitdankelaminData
}

func (r *PenilaianMedisRalanKulitdankelaminStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisRalanKulitdankelaminUpdate ubah penilaian awal medis ralan kulit dan kelamin (PUT); kunci lewat query string.
type PenilaianMedisRalanKulitdankelaminUpdate struct {
	PenilaianMedisRalanKulitdankelaminData
}

func (r *PenilaianMedisRalanKulitdankelaminUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanKulitdankelaminUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRalanKulitdankelaminRules()
}

func (r *PenilaianMedisRalanKulitdankelaminUpdate) Payload() PenilaianMedisRalanKulitdankelaminData {
	return r.PenilaianMedisRalanKulitdankelaminData
}

func (r *PenilaianMedisRalanKulitdankelaminUpdate) DetailValues() map[string][]string { return nil }
