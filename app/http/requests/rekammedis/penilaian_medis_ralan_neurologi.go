package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRalanNeurologiData isian penilaian awal medis ralan neurologi.
type PenilaianMedisRalanNeurologiData struct {
	Tanggal               string `form:"tanggal" json:"tanggal"`
	KdDokter              string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis             string `form:"anamnesis" json:"anamnesis"`
	Hubungan              string `form:"hubungan" json:"hubungan"`
	KeluhanUtama          string `form:"keluhan_utama" json:"keluhan_utama"`
	Rps                   string `form:"rps" json:"rps"`
	Rpd                   string `form:"rpd" json:"rpd"`
	Rpo                   string `form:"rpo" json:"rpo"`
	Alergi                string `form:"alergi" json:"alergi"`
	Kesadaran             string `form:"kesadaran" json:"kesadaran"`
	Status                string `form:"status" json:"status"`
	Td                    string `form:"td" json:"td"`
	Nadi                  string `form:"nadi" json:"nadi"`
	Suhu                  string `form:"suhu" json:"suhu"`
	Rr                    string `form:"rr" json:"rr"`
	Bb                    string `form:"bb" json:"bb"`
	Nyeri                 string `form:"nyeri" json:"nyeri"`
	Gcs                   string `form:"gcs" json:"gcs"`
	Kepala                string `form:"kepala" json:"kepala"`
	KeteranganKepala      string `form:"keterangan_kepala" json:"keterangan_kepala"`
	Thoraks               string `form:"thoraks" json:"thoraks"`
	KeteranganThoraks     string `form:"keterangan_thoraks" json:"keterangan_thoraks"`
	Abdomen               string `form:"abdomen" json:"abdomen"`
	KeteranganAbdomen     string `form:"keterangan_abdomen" json:"keterangan_abdomen"`
	Ekstremitas           string `form:"ekstremitas" json:"ekstremitas"`
	KeteranganEkstremitas string `form:"keterangan_ekstremitas" json:"keterangan_ekstremitas"`
	Columna               string `form:"columna" json:"columna"`
	KeteranganColumna     string `form:"keterangan_columna" json:"keterangan_columna"`
	Muskulos              string `form:"muskulos" json:"muskulos"`
	KeteranganMuskulos    string `form:"keterangan_muskulos" json:"keterangan_muskulos"`
	Lainnya               string `form:"lainnya" json:"lainnya"`
	Lab                   string `form:"lab" json:"lab"`
	Rad                   string `form:"rad" json:"rad"`
	Penunjanglain         string `form:"penunjanglain" json:"penunjanglain"`
	Diagnosis             string `form:"diagnosis" json:"diagnosis"`
	Diagnosis2            string `form:"diagnosis2" json:"diagnosis2"`
	Permasalahan          string `form:"permasalahan" json:"permasalahan"`
	Terapi                string `form:"terapi" json:"terapi"`
	Tindakan              string `form:"tindakan" json:"tindakan"`
	Edukasi               string `form:"edukasi" json:"edukasi"`
}

func penilaianMedisRalanNeurologiRules() map[string]any {
	rules := map[string]any{
		"tanggal":                "required|date",
		"kd_dokter":              "required|string|max_len:20",
		"anamnesis":              "required|in:Autoanamnesis,Alloanamnesis",
		"hubungan":               "string|max_len:30",
		"keluhan_utama":          "string|max_len:2000",
		"rps":                    "string|max_len:2000",
		"rpd":                    "string|max_len:1000",
		"rpo":                    "string|max_len:1000",
		"alergi":                 "string|max_len:50",
		"kesadaran":              "required|in:Compos Mentis,Apatis,Delirum",
		"status":                 "required|in:Skor < 2,Skor >= 2",
		"td":                     "string|max_len:8",
		"nadi":                   "string|max_len:5",
		"suhu":                   "string|max_len:5",
		"rr":                     "string|max_len:5",
		"bb":                     "string|max_len:5",
		"nyeri":                  "string|max_len:50",
		"gcs":                    "string|max_len:10",
		"kepala":                 "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_kepala":      "string|max_len:30",
		"thoraks":                "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_thoraks":     "string|max_len:30",
		"abdomen":                "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_abdomen":     "string|max_len:30",
		"ekstremitas":            "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_ekstremitas": "string|max_len:30",
		"columna":                "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_columna":     "string|max_len:30",
		"muskulos":               "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_muskulos":    "string|max_len:30",
		"lainnya":                "string|max_len:1000",
		"lab":                    "string|max_len:500",
		"rad":                    "string|max_len:500",
		"penunjanglain":          "string|max_len:500",
		"diagnosis":              "string|max_len:500",
		"diagnosis2":             "string|max_len:500",
		"permasalahan":           "string|max_len:500",
		"terapi":                 "string|max_len:500",
		"tindakan":               "string|max_len:500",
		"edukasi":                "string|max_len:500",
	}
	return rules
}

// PenilaianMedisRalanNeurologiStore simpan penilaian awal medis ralan neurologi; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRalanNeurologiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRalanNeurologiData
}

func (r *PenilaianMedisRalanNeurologiStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanNeurologiStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRalanNeurologiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRalanNeurologiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianMedisRalanNeurologiStore) Payload() PenilaianMedisRalanNeurologiData {
	return r.PenilaianMedisRalanNeurologiData
}

func (r *PenilaianMedisRalanNeurologiStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisRalanNeurologiUpdate ubah penilaian awal medis ralan neurologi (PUT); kunci lewat query string.
type PenilaianMedisRalanNeurologiUpdate struct {
	PenilaianMedisRalanNeurologiData
}

func (r *PenilaianMedisRalanNeurologiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanNeurologiUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRalanNeurologiRules()
}

func (r *PenilaianMedisRalanNeurologiUpdate) Payload() PenilaianMedisRalanNeurologiData {
	return r.PenilaianMedisRalanNeurologiData
}

func (r *PenilaianMedisRalanNeurologiUpdate) DetailValues() map[string][]string { return nil }
