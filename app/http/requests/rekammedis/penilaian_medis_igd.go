package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisIgdData isian penilaian awal medis IGD.
type PenilaianMedisIgdData struct {
	Tanggal      string `form:"tanggal" json:"tanggal"`
	KdDokter     string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis    string `form:"anamnesis" json:"anamnesis"`
	Hubungan     string `form:"hubungan" json:"hubungan"`
	KeluhanUtama string `form:"keluhan_utama" json:"keluhan_utama"`
	Rps          string `form:"rps" json:"rps"`
	Rpd          string `form:"rpd" json:"rpd"`
	Rpk          string `form:"rpk" json:"rpk"`
	Rpo          string `form:"rpo" json:"rpo"`
	Alergi       string `form:"alergi" json:"alergi"`
	Keadaan      string `form:"keadaan" json:"keadaan"`
	Gcs          string `form:"gcs" json:"gcs"`
	Kesadaran    string `form:"kesadaran" json:"kesadaran"`
	Td           string `form:"td" json:"td"`
	Nadi         string `form:"nadi" json:"nadi"`
	Rr           string `form:"rr" json:"rr"`
	Suhu         string `form:"suhu" json:"suhu"`
	Spo          string `form:"spo" json:"spo"`
	Bb           string `form:"bb" json:"bb"`
	Tb           string `form:"tb" json:"tb"`
	Kepala       string `form:"kepala" json:"kepala"`
	Mata         string `form:"mata" json:"mata"`
	Gigi         string `form:"gigi" json:"gigi"`
	Leher        string `form:"leher" json:"leher"`
	Thoraks      string `form:"thoraks" json:"thoraks"`
	Abdomen      string `form:"abdomen" json:"abdomen"`
	Genital      string `form:"genital" json:"genital"`
	Ekstremitas  string `form:"ekstremitas" json:"ekstremitas"`
	KetFisik     string `form:"ket_fisik" json:"ket_fisik"`
	KetLokalis   string `form:"ket_lokalis" json:"ket_lokalis"`
	Ekg          string `form:"ekg" json:"ekg"`
	Rad          string `form:"rad" json:"rad"`
	Lab          string `form:"lab" json:"lab"`
	Diagnosis    string `form:"diagnosis" json:"diagnosis"`
	Tata         string `form:"tata" json:"tata"`
}

func penilaianMedisIgdRules() map[string]any {
	rules := map[string]any{
		"tanggal":       "required|date",
		"kd_dokter":     "required|string|max_len:20",
		"anamnesis":     "required|in:Autoanamnesis,Alloanamnesis",
		"hubungan":      "string|max_len:100",
		"keluhan_utama": "string|max_len:2000",
		"rps":           "string|max_len:2000",
		"rpd":           "string|max_len:1000",
		"rpk":           "string|max_len:1000",
		"rpo":           "string|max_len:1000",
		"alergi":        "string|max_len:100",
		"keadaan":       "required|in:Sehat,Sakit Ringan,Sakit Sedang,Sakit Berat",
		"gcs":           "string|max_len:10",
		"kesadaran":     "required|in:Compos Mentis,Apatis,Somnolen,Sopor,Koma",
		"td":            "string|max_len:8",
		"nadi":          "string|max_len:5",
		"rr":            "string|max_len:5",
		"suhu":          "string|max_len:5",
		"spo":           "string|max_len:5",
		"bb":            "string|max_len:5",
		"tb":            "string|max_len:5",
		"kepala":        "required|in:Normal,Abnormal,Tidak Diperiksa",
		"mata":          "required|in:Normal,Abnormal,Tidak Diperiksa",
		"gigi":          "required|in:Normal,Abnormal,Tidak Diperiksa",
		"leher":         "required|in:Normal,Abnormal,Tidak Diperiksa",
		"thoraks":       "required|in:Normal,Abnormal,Tidak Diperiksa",
		"abdomen":       "required|in:Normal,Abnormal,Tidak Diperiksa",
		"genital":       "required|in:Normal,Abnormal,Tidak Diperiksa",
		"ekstremitas":   "required|in:Normal,Abnormal,Tidak Diperiksa",
		"ket_fisik":     "string",
		"ket_lokalis":   "string",
		"ekg":           "string",
		"rad":           "string",
		"lab":           "string",
		"diagnosis":     "string|max_len:500",
		"tata":          "string",
	}
	return rules
}

// PenilaianMedisIgdStore simpan penilaian awal medis IGD; kolom waktu kunci kosong = sekarang.
type PenilaianMedisIgdStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisIgdData
}

func (r *PenilaianMedisIgdStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisIgdStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisIgdRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisIgdStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianMedisIgdStore) Payload() PenilaianMedisIgdData { return r.PenilaianMedisIgdData }

func (r *PenilaianMedisIgdStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisIgdUpdate ubah penilaian awal medis IGD (PUT); kunci lewat query string.
type PenilaianMedisIgdUpdate struct {
	PenilaianMedisIgdData
}

func (r *PenilaianMedisIgdUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisIgdUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisIgdRules()
}

func (r *PenilaianMedisIgdUpdate) Payload() PenilaianMedisIgdData { return r.PenilaianMedisIgdData }

func (r *PenilaianMedisIgdUpdate) DetailValues() map[string][]string { return nil }
