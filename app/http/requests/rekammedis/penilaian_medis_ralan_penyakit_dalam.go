package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRalanPenyakitDalamData isian penilaian awal medis ralan penyakit dalam.
type PenilaianMedisRalanPenyakitDalamData struct {
	Tanggal               string `form:"tanggal" json:"tanggal"`
	KdDokter              string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis             string `form:"anamnesis" json:"anamnesis"`
	Hubungan              string `form:"hubungan" json:"hubungan"`
	KeluhanUtama          string `form:"keluhan_utama" json:"keluhan_utama"`
	Rps                   string `form:"rps" json:"rps"`
	Rpd                   string `form:"rpd" json:"rpd"`
	Rpo                   string `form:"rpo" json:"rpo"`
	Alergi                string `form:"alergi" json:"alergi"`
	Kondisi               string `form:"kondisi" json:"kondisi"`
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
	KeteranganThorak      string `form:"keterangan_thorak" json:"keterangan_thorak"`
	Abdomen               string `form:"abdomen" json:"abdomen"`
	KeteranganAbdomen     string `form:"keterangan_abdomen" json:"keterangan_abdomen"`
	Ekstremitas           string `form:"ekstremitas" json:"ekstremitas"`
	KeteranganEkstremitas string `form:"keterangan_ekstremitas" json:"keterangan_ekstremitas"`
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

func penilaianMedisRalanPenyakitDalamRules() map[string]any {
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
		"kondisi":                "string|max_len:500",
		"status":                 "string|max_len:100",
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
		"keterangan_thorak":      "string|max_len:30",
		"abdomen":                "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_abdomen":     "string|max_len:30",
		"ekstremitas":            "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_ekstremitas": "string|max_len:30",
		"lainnya":                "string|max_len:1000",
		"lab":                    "string|max_len:1000",
		"rad":                    "string|max_len:1000",
		"penunjanglain":          "string|max_len:1000",
		"diagnosis":              "string|max_len:500",
		"diagnosis2":             "string|max_len:500",
		"permasalahan":           "string|max_len:500",
		"terapi":                 "string|max_len:500",
		"tindakan":               "string|max_len:200",
		"edukasi":                "string|max_len:500",
	}
	return rules
}

// PenilaianMedisRalanPenyakitDalamStore simpan penilaian awal medis ralan penyakit dalam; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRalanPenyakitDalamStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRalanPenyakitDalamData
}

func (r *PenilaianMedisRalanPenyakitDalamStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanPenyakitDalamStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRalanPenyakitDalamRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRalanPenyakitDalamStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianMedisRalanPenyakitDalamStore) Payload() PenilaianMedisRalanPenyakitDalamData {
	return r.PenilaianMedisRalanPenyakitDalamData
}

func (r *PenilaianMedisRalanPenyakitDalamStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisRalanPenyakitDalamUpdate ubah penilaian awal medis ralan penyakit dalam (PUT); kunci lewat query string.
type PenilaianMedisRalanPenyakitDalamUpdate struct {
	PenilaianMedisRalanPenyakitDalamData
}

func (r *PenilaianMedisRalanPenyakitDalamUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanPenyakitDalamUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRalanPenyakitDalamRules()
}

func (r *PenilaianMedisRalanPenyakitDalamUpdate) Payload() PenilaianMedisRalanPenyakitDalamData {
	return r.PenilaianMedisRalanPenyakitDalamData
}

func (r *PenilaianMedisRalanPenyakitDalamUpdate) DetailValues() map[string][]string { return nil }
