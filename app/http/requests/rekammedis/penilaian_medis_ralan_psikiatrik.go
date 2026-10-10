package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRalanPsikiatrikData isian penilaian awal medis ralan psikiatrik.
type PenilaianMedisRalanPsikiatrikData struct {
	Tanggal            string `form:"tanggal" json:"tanggal"`
	KdDokter           string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis          string `form:"anamnesis" json:"anamnesis"`
	Hubungan           string `form:"hubungan" json:"hubungan"`
	KeluhanUtama       string `form:"keluhan_utama" json:"keluhan_utama"`
	Rps                string `form:"rps" json:"rps"`
	Rpd                string `form:"rpd" json:"rpd"`
	Rpk                string `form:"rpk" json:"rpk"`
	Rpo                string `form:"rpo" json:"rpo"`
	Alergi             string `form:"alergi" json:"alergi"`
	Penampilan         string `form:"penampilan" json:"penampilan"`
	Pembicaraan        string `form:"pembicaraan" json:"pembicaraan"`
	Psikomotor         string `form:"psikomotor" json:"psikomotor"`
	Sikap              string `form:"sikap" json:"sikap"`
	Mood               string `form:"mood" json:"mood"`
	FungsiKognitif     string `form:"fungsi_kognitif" json:"fungsi_kognitif"`
	GangguanPersepsi   string `form:"gangguan_persepsi" json:"gangguan_persepsi"`
	ProsesPikir        string `form:"proses_pikir" json:"proses_pikir"`
	PengendalianImpuls string `form:"pengendalian_impuls" json:"pengendalian_impuls"`
	Tilikan            string `form:"tilikan" json:"tilikan"`
	Rta                string `form:"rta" json:"rta"`
	Keadaan            string `form:"keadaan" json:"keadaan"`
	Gcs                string `form:"gcs" json:"gcs"`
	Kesadaran          string `form:"kesadaran" json:"kesadaran"`
	Td                 string `form:"td" json:"td"`
	Nadi               string `form:"nadi" json:"nadi"`
	Rr                 string `form:"rr" json:"rr"`
	Suhu               string `form:"suhu" json:"suhu"`
	Spo                string `form:"spo" json:"spo"`
	Bb                 string `form:"bb" json:"bb"`
	Tb                 string `form:"tb" json:"tb"`
	Kepala             string `form:"kepala" json:"kepala"`
	Gigi               string `form:"gigi" json:"gigi"`
	Tht                string `form:"tht" json:"tht"`
	Thoraks            string `form:"thoraks" json:"thoraks"`
	Abdomen            string `form:"abdomen" json:"abdomen"`
	Genital            string `form:"genital" json:"genital"`
	Ekstremitas        string `form:"ekstremitas" json:"ekstremitas"`
	Kulit              string `form:"kulit" json:"kulit"`
	KetFisik           string `form:"ket_fisik" json:"ket_fisik"`
	Penunjang          string `form:"penunjang" json:"penunjang"`
	Diagnosis          string `form:"diagnosis" json:"diagnosis"`
	Tata               string `form:"tata" json:"tata"`
	Konsulrujuk        string `form:"konsulrujuk" json:"konsulrujuk"`
}

func penilaianMedisRalanPsikiatrikRules() map[string]any {
	rules := map[string]any{
		"tanggal":             "required|date",
		"kd_dokter":           "required|string|max_len:20",
		"anamnesis":           "required|in:Autoanamnesis,Alloanamnesis",
		"hubungan":            "string|max_len:30",
		"keluhan_utama":       "string|max_len:2000",
		"rps":                 "string|max_len:2000",
		"rpd":                 "string|max_len:1000",
		"rpk":                 "string|max_len:1000",
		"rpo":                 "string|max_len:1000",
		"alergi":              "string|max_len:50",
		"penampilan":          "string|max_len:200",
		"pembicaraan":         "string|max_len:200",
		"psikomotor":          "string|max_len:200",
		"sikap":               "string|max_len:200",
		"mood":                "string|max_len:200",
		"fungsi_kognitif":     "string|max_len:200",
		"gangguan_persepsi":   "string|max_len:200",
		"proses_pikir":        "string|max_len:200",
		"pengendalian_impuls": "string|max_len:200",
		"tilikan":             "string|max_len:200",
		"rta":                 "string|max_len:200",
		"keadaan":             "required|in:Sehat,Sakit Ringan,Sakit Sedang,Sakit Berat",
		"gcs":                 "string|max_len:10",
		"kesadaran":           "required|in:Compos Mentis,Apatis,Somnolen,Sopor,Koma",
		"td":                  "string|max_len:8",
		"nadi":                "string|max_len:5",
		"rr":                  "string|max_len:5",
		"suhu":                "string|max_len:5",
		"spo":                 "string|max_len:5",
		"bb":                  "string|max_len:5",
		"tb":                  "string|max_len:5",
		"kepala":              "required|in:Normal,Abnormal,Tidak Diperiksa",
		"gigi":                "required|in:Normal,Abnormal,Tidak Diperiksa",
		"tht":                 "required|in:Normal,Abnormal,Tidak Diperiksa",
		"thoraks":             "required|in:Normal,Abnormal,Tidak Diperiksa",
		"abdomen":             "required|in:Normal,Abnormal,Tidak Diperiksa",
		"genital":             "required|in:Normal,Abnormal,Tidak Diperiksa",
		"ekstremitas":         "required|in:Normal,Abnormal,Tidak Diperiksa",
		"kulit":               "required|in:Normal,Abnormal,Tidak Diperiksa",
		"ket_fisik":           "string|max_len:1000",
		"penunjang":           "string|max_len:1000",
		"diagnosis":           "string|max_len:300",
		"tata":                "string|max_len:1000",
		"konsulrujuk":         "string|max_len:500",
	}
	return rules
}

// PenilaianMedisRalanPsikiatrikStore simpan penilaian awal medis ralan psikiatrik; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRalanPsikiatrikStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRalanPsikiatrikData
}

func (r *PenilaianMedisRalanPsikiatrikStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanPsikiatrikStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRalanPsikiatrikRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRalanPsikiatrikStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianMedisRalanPsikiatrikStore) Payload() PenilaianMedisRalanPsikiatrikData {
	return r.PenilaianMedisRalanPsikiatrikData
}

func (r *PenilaianMedisRalanPsikiatrikStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisRalanPsikiatrikUpdate ubah penilaian awal medis ralan psikiatrik (PUT); kunci lewat query string.
type PenilaianMedisRalanPsikiatrikUpdate struct {
	PenilaianMedisRalanPsikiatrikData
}

func (r *PenilaianMedisRalanPsikiatrikUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanPsikiatrikUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRalanPsikiatrikRules()
}

func (r *PenilaianMedisRalanPsikiatrikUpdate) Payload() PenilaianMedisRalanPsikiatrikData {
	return r.PenilaianMedisRalanPsikiatrikData
}

func (r *PenilaianMedisRalanPsikiatrikUpdate) DetailValues() map[string][]string { return nil }
