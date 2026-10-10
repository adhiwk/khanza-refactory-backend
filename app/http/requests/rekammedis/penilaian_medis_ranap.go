package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRanapData isian penilaian awal medis ranap dewasa.
type PenilaianMedisRanapData struct {
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
	Tht          string `form:"tht" json:"tht"`
	Thoraks      string `form:"thoraks" json:"thoraks"`
	Jantung      string `form:"jantung" json:"jantung"`
	Paru         string `form:"paru" json:"paru"`
	Abdomen      string `form:"abdomen" json:"abdomen"`
	Genital      string `form:"genital" json:"genital"`
	Ekstremitas  string `form:"ekstremitas" json:"ekstremitas"`
	Kulit        string `form:"kulit" json:"kulit"`
	KetFisik     string `form:"ket_fisik" json:"ket_fisik"`
	KetLokalis   string `form:"ket_lokalis" json:"ket_lokalis"`
	Lab          string `form:"lab" json:"lab"`
	Rad          string `form:"rad" json:"rad"`
	Penunjang    string `form:"penunjang" json:"penunjang"`
	Diagnosis    string `form:"diagnosis" json:"diagnosis"`
	Tata         string `form:"tata" json:"tata"`
	Edukasi      string `form:"edukasi" json:"edukasi"`
}

func penilaianMedisRanapRules() map[string]any {
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
		"tht":           "required|in:Normal,Abnormal,Tidak Diperiksa",
		"thoraks":       "required|in:Normal,Abnormal,Tidak Diperiksa",
		"jantung":       "required|in:Normal,Abnormal,Tidak Diperiksa",
		"paru":          "required|in:Normal,Abnormal,Tidak Diperiksa",
		"abdomen":       "required|in:Normal,Abnormal,Tidak Diperiksa",
		"genital":       "required|in:Normal,Abnormal,Tidak Diperiksa",
		"ekstremitas":   "required|in:Normal,Abnormal,Tidak Diperiksa",
		"kulit":         "required|in:Normal,Abnormal,Tidak Diperiksa",
		"ket_fisik":     "string",
		"ket_lokalis":   "string",
		"lab":           "string",
		"rad":           "string",
		"penunjang":     "string",
		"diagnosis":     "string|max_len:500",
		"tata":          "string",
		"edukasi":       "string|max_len:1000",
	}
	return rules
}

// PenilaianMedisRanapStore simpan penilaian awal medis ranap dewasa; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRanapStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRanapData
}

func (r *PenilaianMedisRanapStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRanapStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRanapRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRanapStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianMedisRanapStore) Payload() PenilaianMedisRanapData {
	return r.PenilaianMedisRanapData
}

func (r *PenilaianMedisRanapStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisRanapUpdate ubah penilaian awal medis ranap dewasa (PUT); kunci lewat query string.
type PenilaianMedisRanapUpdate struct {
	PenilaianMedisRanapData
}

func (r *PenilaianMedisRanapUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRanapUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRanapRules()
}

func (r *PenilaianMedisRanapUpdate) Payload() PenilaianMedisRanapData {
	return r.PenilaianMedisRanapData
}

func (r *PenilaianMedisRanapUpdate) DetailValues() map[string][]string { return nil }
