package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRalanBedahMulutData isian penilaian awal medis ralan bedah mulut.
type PenilaianMedisRalanBedahMulutData struct {
	Tanggal               string `form:"tanggal" json:"tanggal"`
	KdDokter              string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis             string `form:"anamnesis" json:"anamnesis"`
	Hubungan              string `form:"hubungan" json:"hubungan"`
	KeluhanUtama          string `form:"keluhan_utama" json:"keluhan_utama"`
	Rps                   string `form:"rps" json:"rps"`
	Rpk                   string `form:"rpk" json:"rpk"`
	Alergi                string `form:"alergi" json:"alergi"`
	Keadaan               string `form:"keadaan" json:"keadaan"`
	Kesadaran             string `form:"kesadaran" json:"kesadaran"`
	Nyeri                 string `form:"nyeri" json:"nyeri"`
	Td                    string `form:"td" json:"td"`
	Nadi                  string `form:"nadi" json:"nadi"`
	Suhu                  string `form:"suhu" json:"suhu"`
	Rr                    string `form:"rr" json:"rr"`
	Bb                    string `form:"bb" json:"bb"`
	Tb                    string `form:"tb" json:"tb"`
	StatusNutrisi         string `form:"status_nutrisi" json:"status_nutrisi"`
	Kulit                 string `form:"kulit" json:"kulit"`
	KeteranganKulit       string `form:"keterangan_kulit" json:"keterangan_kulit"`
	Kepala                string `form:"kepala" json:"kepala"`
	KeteranganKepala      string `form:"keterangan_kepala" json:"keterangan_kepala"`
	Mata                  string `form:"mata" json:"mata"`
	KeteranganMata        string `form:"keterangan_mata" json:"keterangan_mata"`
	Leher                 string `form:"leher" json:"leher"`
	KeteranganLeher       string `form:"keterangan_leher" json:"keterangan_leher"`
	Kelenjar              string `form:"kelenjar" json:"kelenjar"`
	KeteranganKelenjar    string `form:"keterangan_kelenjar" json:"keterangan_kelenjar"`
	Dada                  string `form:"dada" json:"dada"`
	KeteranganDada        string `form:"keterangan_dada" json:"keterangan_dada"`
	Perut                 string `form:"perut" json:"perut"`
	KeteranganPerut       string `form:"keterangan_perut" json:"keterangan_perut"`
	Ekstremitas           string `form:"ekstremitas" json:"ekstremitas"`
	KeteranganEkstremitas string `form:"keterangan_ekstremitas" json:"keterangan_ekstremitas"`
	Wajah                 string `form:"wajah" json:"wajah"`
	Intra                 string `form:"intra" json:"intra"`
	Gigigeligi            string `form:"gigigeligi" json:"gigigeligi"`
	Lab                   string `form:"lab" json:"lab"`
	Rad                   string `form:"rad" json:"rad"`
	Penunjang             string `form:"penunjang" json:"penunjang"`
	Diagnosis             string `form:"diagnosis" json:"diagnosis"`
	Diagnosis2            string `form:"diagnosis2" json:"diagnosis2"`
	Permasalahan          string `form:"permasalahan" json:"permasalahan"`
	Terapi                string `form:"terapi" json:"terapi"`
	Tindakan              string `form:"tindakan" json:"tindakan"`
	Edukasi               string `form:"edukasi" json:"edukasi"`
}

func penilaianMedisRalanBedahMulutRules() map[string]any {
	rules := map[string]any{
		"tanggal":                "required|date",
		"kd_dokter":              "required|string|max_len:20",
		"anamnesis":              "required|in:Autoanamnesis,Alloanamnesis",
		"hubungan":               "string|max_len:30",
		"keluhan_utama":          "string|max_len:2000",
		"rps":                    "string|max_len:2000",
		"rpk":                    "string|max_len:1000",
		"alergi":                 "string|max_len:50",
		"keadaan":                "required|in:Baik,Sedang,Lemah,Buruk",
		"kesadaran":              "required|in:Compos Mentis,Apatis,Somnolen,Sopor,Koma",
		"nyeri":                  "required|in:Tidak Nyeri,Nyeri Ringan,Nyeri Sedang,Nyeri Berat,Nyeri Sangat Berat,Nyeri Tak Tertahankan",
		"td":                     "string|max_len:8",
		"nadi":                   "string|max_len:5",
		"suhu":                   "string|max_len:5",
		"rr":                     "string|max_len:5",
		"bb":                     "string|max_len:5",
		"tb":                     "string|max_len:5",
		"status_nutrisi":         "string|max_len:50",
		"kulit":                  "required|in:Tidak,Ya",
		"keterangan_kulit":       "string|max_len:30",
		"kepala":                 "required|in:Tidak,Ya",
		"keterangan_kepala":      "string|max_len:30",
		"mata":                   "required|in:Tidak,Ya",
		"keterangan_mata":        "string|max_len:30",
		"leher":                  "required|in:Tidak,Ya",
		"keterangan_leher":       "string|max_len:30",
		"kelenjar":               "required|in:Tidak,Ya",
		"keterangan_kelenjar":    "string|max_len:30",
		"dada":                   "required|in:Tidak,Ya",
		"keterangan_dada":        "string|max_len:30",
		"perut":                  "required|in:Tidak,Ya",
		"keterangan_perut":       "string|max_len:30",
		"ekstremitas":            "required|in:Tidak,Ya",
		"keterangan_ekstremitas": "string|max_len:30",
		"wajah":                  "string|max_len:1000",
		"intra":                  "string|max_len:1000",
		"gigigeligi":             "string|max_len:1000",
		"lab":                    "string|max_len:300",
		"rad":                    "string|max_len:300",
		"penunjang":              "string|max_len:300",
		"diagnosis":              "string|max_len:500",
		"diagnosis2":             "string|max_len:500",
		"permasalahan":           "string|max_len:1000",
		"terapi":                 "string|max_len:1000",
		"tindakan":               "string|max_len:1000",
		"edukasi":                "string|max_len:1000",
	}
	return rules
}

// PenilaianMedisRalanBedahMulutStore simpan penilaian awal medis ralan bedah mulut; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRalanBedahMulutStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRalanBedahMulutData
}

func (r *PenilaianMedisRalanBedahMulutStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanBedahMulutStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRalanBedahMulutRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRalanBedahMulutStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianMedisRalanBedahMulutStore) Payload() PenilaianMedisRalanBedahMulutData {
	return r.PenilaianMedisRalanBedahMulutData
}

func (r *PenilaianMedisRalanBedahMulutStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisRalanBedahMulutUpdate ubah penilaian awal medis ralan bedah mulut (PUT); kunci lewat query string.
type PenilaianMedisRalanBedahMulutUpdate struct {
	PenilaianMedisRalanBedahMulutData
}

func (r *PenilaianMedisRalanBedahMulutUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanBedahMulutUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRalanBedahMulutRules()
}

func (r *PenilaianMedisRalanBedahMulutUpdate) Payload() PenilaianMedisRalanBedahMulutData {
	return r.PenilaianMedisRalanBedahMulutData
}

func (r *PenilaianMedisRalanBedahMulutUpdate) DetailValues() map[string][]string { return nil }
