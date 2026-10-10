package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// AsuhanGiziData isian asuhan gizi.
type AsuhanGiziData struct {
	AntropometriBb      string `form:"antropometri_bb" json:"antropometri_bb"`
	AntropometriTb      string `form:"antropometri_tb" json:"antropometri_tb"`
	AntropometriImt     string `form:"antropometri_imt" json:"antropometri_imt"`
	AntropometriLla     string `form:"antropometri_lla" json:"antropometri_lla"`
	AntropometriTl      string `form:"antropometri_tl" json:"antropometri_tl"`
	AntropometriUlna    string `form:"antropometri_ulna" json:"antropometri_ulna"`
	AntropometriBbideal string `form:"antropometri_bbideal" json:"antropometri_bbideal"`
	AntropometriBbperu  string `form:"antropometri_bbperu" json:"antropometri_bbperu"`
	AntropometriTbperu  string `form:"antropometri_tbperu" json:"antropometri_tbperu"`
	AntropometriBbpertb string `form:"antropometri_bbpertb" json:"antropometri_bbpertb"`
	AntropometriLlaperu string `form:"antropometri_llaperu" json:"antropometri_llaperu"`
	Biokimia            string `form:"biokimia" json:"biokimia"`
	FisikKlinis         string `form:"fisik_klinis" json:"fisik_klinis"`
	AlergiTelur         string `form:"alergi_telur" json:"alergi_telur"`
	AlergiSusuSapi      string `form:"alergi_susu_sapi" json:"alergi_susu_sapi"`
	AlergiKacang        string `form:"alergi_kacang" json:"alergi_kacang"`
	AlergiGluten        string `form:"alergi_gluten" json:"alergi_gluten"`
	AlergiUdang         string `form:"alergi_udang" json:"alergi_udang"`
	AlergiIkan          string `form:"alergi_ikan" json:"alergi_ikan"`
	AlergiHazelnut      string `form:"alergi_hazelnut" json:"alergi_hazelnut"`
	PolaMakan           string `form:"pola_makan" json:"pola_makan"`
	RiwayatPersonal     string `form:"riwayat_personal" json:"riwayat_personal"`
	Diagnosis           string `form:"diagnosis" json:"diagnosis"`
	IntervensiGizi      string `form:"intervensi_gizi" json:"intervensi_gizi"`
	MonitoringEvaluasi  string `form:"monitoring_evaluasi" json:"monitoring_evaluasi"`
	Nip                 string `form:"nip" json:"nip"`
}

func asuhanGiziRules() map[string]any {
	rules := map[string]any{
		"antropometri_bb":      "string|max_len:5",
		"antropometri_tb":      "string|max_len:5",
		"antropometri_imt":     "string|max_len:5",
		"antropometri_lla":     "string|max_len:5",
		"antropometri_tl":      "string|max_len:5",
		"antropometri_ulna":    "string|max_len:5",
		"antropometri_bbideal": "string|max_len:5",
		"antropometri_bbperu":  "string|max_len:5",
		"antropometri_tbperu":  "string|max_len:5",
		"antropometri_bbpertb": "string|max_len:5",
		"antropometri_llaperu": "string|max_len:5",
		"biokimia":             "string|max_len:1000",
		"fisik_klinis":         "string|max_len:2000",
		"alergi_telur":         "in:Ya,Tidak",
		"alergi_susu_sapi":     "in:Ya,Tidak",
		"alergi_kacang":        "in:Ya,Tidak",
		"alergi_gluten":        "in:Ya,Tidak",
		"alergi_udang":         "in:Ya,Tidak",
		"alergi_ikan":          "in:Ya,Tidak",
		"alergi_hazelnut":      "in:Ya,Tidak",
		"pola_makan":           "string|max_len:2000",
		"riwayat_personal":     "string|max_len:1000",
		"diagnosis":            "string|max_len:2000",
		"intervensi_gizi":      "string|max_len:2000",
		"monitoring_evaluasi":  "string|max_len:1000",
		"nip":                  "required|string|max_len:20",
	}
	return rules
}

// AsuhanGiziStore simpan asuhan gizi; kolom waktu kunci kosong = sekarang.
type AsuhanGiziStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	AsuhanGiziData
}

func (r *AsuhanGiziStore) Authorize(ctx http.Context) error { return nil }

func (r *AsuhanGiziStore) Rules(ctx http.Context) map[string]any {
	rules := asuhanGiziRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *AsuhanGiziStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *AsuhanGiziStore) Payload() AsuhanGiziData { return r.AsuhanGiziData }

func (r *AsuhanGiziStore) DetailValues() map[string][]string { return nil }

// AsuhanGiziUpdate ubah asuhan gizi (PUT); kunci lewat query string.
type AsuhanGiziUpdate struct {
	AsuhanGiziData
}

func (r *AsuhanGiziUpdate) Authorize(ctx http.Context) error { return nil }

func (r *AsuhanGiziUpdate) Rules(ctx http.Context) map[string]any { return asuhanGiziRules() }

func (r *AsuhanGiziUpdate) Payload() AsuhanGiziData { return r.AsuhanGiziData }

func (r *AsuhanGiziUpdate) DetailValues() map[string][]string { return nil }
