package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRalanRehabMedikData isian penilaian awal medis ralan rehab medik.
type PenilaianMedisRalanRehabMedikData struct {
	Tanggal               string `form:"tanggal" json:"tanggal"`
	KdDokter              string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis             string `form:"anamnesis" json:"anamnesis"`
	Hubungan              string `form:"hubungan" json:"hubungan"`
	KeluhanUtama          string `form:"keluhan_utama" json:"keluhan_utama"`
	Rps                   string `form:"rps" json:"rps"`
	Rpd                   string `form:"rpd" json:"rpd"`
	Alergi                string `form:"alergi" json:"alergi"`
	Kesadaran             string `form:"kesadaran" json:"kesadaran"`
	Nyeri                 string `form:"nyeri" json:"nyeri"`
	SkalaNyeri            string `form:"skala_nyeri" json:"skala_nyeri"`
	Td                    string `form:"td" json:"td"`
	Nadi                  string `form:"nadi" json:"nadi"`
	Suhu                  string `form:"suhu" json:"suhu"`
	Rr                    string `form:"rr" json:"rr"`
	Bb                    string `form:"bb" json:"bb"`
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
	ResikoJatuh           string `form:"resiko_jatuh" json:"resiko_jatuh"`
	ResikoNutrisional     string `form:"resiko_nutrisional" json:"resiko_nutrisional"`
	KebutuhanFungsional   string `form:"kebutuhan_fungsional" json:"kebutuhan_fungsional"`
	DiagnosaMedis         string `form:"diagnosa_medis" json:"diagnosa_medis"`
	DiagnosaFungsi        string `form:"diagnosa_fungsi" json:"diagnosa_fungsi"`
	PenunjangLain         string `form:"penunjang_lain" json:"penunjang_lain"`
	Fisio                 string `form:"fisio" json:"fisio"`
	Okupasi               string `form:"okupasi" json:"okupasi"`
	Wicara                string `form:"wicara" json:"wicara"`
	Akupuntur             string `form:"akupuntur" json:"akupuntur"`
	Tatalain              string `form:"tatalain" json:"tatalain"`
	FrekuensiTerapi       string `form:"frekuensi_terapi" json:"frekuensi_terapi"`
	Fisioterapi           string `form:"fisioterapi" json:"fisioterapi"`
	TerapiOkupasi         string `form:"terapi_okupasi" json:"terapi_okupasi"`
	TerapiWicara          string `form:"terapi_wicara" json:"terapi_wicara"`
	TerapiAkupuntur       string `form:"terapi_akupuntur" json:"terapi_akupuntur"`
	TerapiLainnya         string `form:"terapi_lainnya" json:"terapi_lainnya"`
	Edukasi               string `form:"edukasi" json:"edukasi"`
}

func penilaianMedisRalanRehabMedikRules() map[string]any {
	rules := map[string]any{
		"tanggal":                "required|date",
		"kd_dokter":              "string|max_len:20",
		"anamnesis":              "in:Autoanamnesis,Alloanamnesis",
		"hubungan":               "string|max_len:30",
		"keluhan_utama":          "string|max_len:2000",
		"rps":                    "string|max_len:2000",
		"rpd":                    "string|max_len:1000",
		"alergi":                 "string|max_len:50",
		"kesadaran":              "in:Compos Mentis,Apatis,Delirum",
		"nyeri":                  "in:Tidak Nyeri,Nyeri Sedang,Nyeri Sangat Hebat",
		"skala_nyeri":            "in:0,1,2,3,4,5,6,7,8,9,10",
		"td":                     "string|max_len:8",
		"nadi":                   "string|max_len:5",
		"suhu":                   "string|max_len:5",
		"rr":                     "string|max_len:5",
		"bb":                     "string|max_len:5",
		"kepala":                 "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_kepala":      "string|max_len:30",
		"thoraks":                "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_thoraks":     "string|max_len:30",
		"abdomen":                "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_abdomen":     "string|max_len:30",
		"ekstremitas":            "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_ekstremitas": "string|max_len:30",
		"columna":                "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_columna":     "string|max_len:30",
		"muskulos":               "in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_muskulos":    "string|max_len:30",
		"lainnya":                "string|max_len:1000",
		"resiko_jatuh":           "in:Tidak Berisiko,Berisiko Sedang,Berisiko Tinggi",
		"resiko_nutrisional":     "in:Tidak Berisiko Malnutrisi,Berisiko Malnutrisi,Malnutrisi",
		"kebutuhan_fungsional":   "in:Tidak Perlu Bantuan,Perlu Bantuan,Perlu Bantuan Total",
		"diagnosa_medis":         "string|max_len:500",
		"diagnosa_fungsi":        "string|max_len:500",
		"penunjang_lain":         "string|max_len:500",
		"fisio":                  "string|max_len:100",
		"okupasi":                "string|max_len:100",
		"wicara":                 "string|max_len:100",
		"akupuntur":              "string|max_len:100",
		"tatalain":               "string|max_len:100",
		"frekuensi_terapi":       "string|max_len:40",
		"fisioterapi":            "required|date",
		"terapi_okupasi":         "required|date",
		"terapi_wicara":          "required|date",
		"terapi_akupuntur":       "required|date",
		"terapi_lainnya":         "required|date",
		"edukasi":                "string|max_len:500",
	}
	return rules
}

// PenilaianMedisRalanRehabMedikStore simpan penilaian awal medis ralan rehab medik; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRalanRehabMedikStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRalanRehabMedikData
}

func (r *PenilaianMedisRalanRehabMedikStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanRehabMedikStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRalanRehabMedikRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRalanRehabMedikStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianMedisRalanRehabMedikStore) Payload() PenilaianMedisRalanRehabMedikData {
	return r.PenilaianMedisRalanRehabMedikData
}

func (r *PenilaianMedisRalanRehabMedikStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisRalanRehabMedikUpdate ubah penilaian awal medis ralan rehab medik (PUT); kunci lewat query string.
type PenilaianMedisRalanRehabMedikUpdate struct {
	PenilaianMedisRalanRehabMedikData
}

func (r *PenilaianMedisRalanRehabMedikUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanRehabMedikUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRalanRehabMedikRules()
}

func (r *PenilaianMedisRalanRehabMedikUpdate) Payload() PenilaianMedisRalanRehabMedikData {
	return r.PenilaianMedisRalanRehabMedikData
}

func (r *PenilaianMedisRalanRehabMedikUpdate) DetailValues() map[string][]string { return nil }
