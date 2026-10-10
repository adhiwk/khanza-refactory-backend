package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRanapJantungData isian penilaian awal medis ranap jantung.
type PenilaianMedisRanapJantungData struct {
	Tanggal               string `form:"tanggal" json:"tanggal"`
	KdDokter              string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis             string `form:"anamnesis" json:"anamnesis"`
	Hubungan              string `form:"hubungan" json:"hubungan"`
	KeluhanUtama          string `form:"keluhan_utama" json:"keluhan_utama"`
	Rps                   string `form:"rps" json:"rps"`
	Rpk                   string `form:"rpk" json:"rpk"`
	Rpd                   string `form:"rpd" json:"rpd"`
	Rpo                   string `form:"rpo" json:"rpo"`
	Alergi                string `form:"alergi" json:"alergi"`
	Td                    string `form:"td" json:"td"`
	Bb                    string `form:"bb" json:"bb"`
	Tb                    string `form:"tb" json:"tb"`
	Suhu                  string `form:"suhu" json:"suhu"`
	Nadi                  string `form:"nadi" json:"nadi"`
	Rr                    string `form:"rr" json:"rr"`
	KeadaanUmum           string `form:"keadaan_umum" json:"keadaan_umum"`
	Nyeri                 string `form:"nyeri" json:"nyeri"`
	StatusNutrisi         string `form:"status_nutrisi" json:"status_nutrisi"`
	Jantung               string `form:"jantung" json:"jantung"`
	KeteranganJantung     string `form:"keterangan_jantung" json:"keterangan_jantung"`
	Paru                  string `form:"paru" json:"paru"`
	KeteranganParu        string `form:"keterangan_paru" json:"keterangan_paru"`
	Ekstrimitas           string `form:"ekstrimitas" json:"ekstrimitas"`
	KeteranganEkstrimitas string `form:"keterangan_ekstrimitas" json:"keterangan_ekstrimitas"`
	Lainnya               string `form:"lainnya" json:"lainnya"`
	Lab                   string `form:"lab" json:"lab"`
	Ekg                   string `form:"ekg" json:"ekg"`
	PenunjangLain         string `form:"penunjang_lain" json:"penunjang_lain"`
	Diagnosis             string `form:"diagnosis" json:"diagnosis"`
	Diagnosis2            string `form:"diagnosis2" json:"diagnosis2"`
	Permasalahan          string `form:"permasalahan" json:"permasalahan"`
	Terapi                string `form:"terapi" json:"terapi"`
	Tindakan              string `form:"tindakan" json:"tindakan"`
	Edukasi               string `form:"edukasi" json:"edukasi"`
}

func penilaianMedisRanapJantungRules() map[string]any {
	rules := map[string]any{
		"tanggal":                "required|date",
		"kd_dokter":              "required|string|max_len:20",
		"anamnesis":              "required|in:Autoanamnesis,Alloanamnesis",
		"hubungan":               "string|max_len:30",
		"keluhan_utama":          "string|max_len:2000",
		"rps":                    "string|max_len:2000",
		"rpk":                    "string|max_len:1000",
		"rpd":                    "string|max_len:1000",
		"rpo":                    "string|max_len:1000",
		"alergi":                 "string|max_len:50",
		"td":                     "string|max_len:8",
		"bb":                     "string|max_len:5",
		"tb":                     "string|max_len:5",
		"suhu":                   "string|max_len:5",
		"nadi":                   "string|max_len:5",
		"rr":                     "string|max_len:5",
		"keadaan_umum":           "in:Sehat,Sakit Ringan,Sakit Sedang,Sakit Berat",
		"nyeri":                  "string|max_len:50",
		"status_nutrisi":         "string|max_len:50",
		"jantung":                "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_jantung":     "string|max_len:50",
		"paru":                   "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_paru":        "string|max_len:50",
		"ekstrimitas":            "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_ekstrimitas": "string|max_len:50",
		"lainnya":                "string|max_len:1000",
		"lab":                    "string|max_len:500",
		"ekg":                    "string|max_len:500",
		"penunjang_lain":         "string|max_len:500",
		"diagnosis":              "string|max_len:500",
		"diagnosis2":             "string|max_len:500",
		"permasalahan":           "string|max_len:500",
		"terapi":                 "string|max_len:500",
		"tindakan":               "string|max_len:500",
		"edukasi":                "string|max_len:500",
	}
	return rules
}

// PenilaianMedisRanapJantungStore simpan penilaian awal medis ranap jantung; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRanapJantungStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRanapJantungData
}

func (r *PenilaianMedisRanapJantungStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRanapJantungStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRanapJantungRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRanapJantungStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianMedisRanapJantungStore) Payload() PenilaianMedisRanapJantungData {
	return r.PenilaianMedisRanapJantungData
}

func (r *PenilaianMedisRanapJantungStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisRanapJantungUpdate ubah penilaian awal medis ranap jantung (PUT); kunci lewat query string.
type PenilaianMedisRanapJantungUpdate struct {
	PenilaianMedisRanapJantungData
}

func (r *PenilaianMedisRanapJantungUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRanapJantungUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRanapJantungRules()
}

func (r *PenilaianMedisRanapJantungUpdate) Payload() PenilaianMedisRanapJantungData {
	return r.PenilaianMedisRanapJantungData
}

func (r *PenilaianMedisRanapJantungUpdate) DetailValues() map[string][]string { return nil }
