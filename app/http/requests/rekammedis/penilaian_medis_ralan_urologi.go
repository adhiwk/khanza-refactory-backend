package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRalanUrologiData isian penilaian awal medis ralan urologi.
type PenilaianMedisRalanUrologiData struct {
	Tanggal               string `form:"tanggal" json:"tanggal"`
	KdDokter              string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis             string `form:"anamnesis" json:"anamnesis"`
	Hubungan              string `form:"hubungan" json:"hubungan"`
	KeluhanUtama          string `form:"keluhan_utama" json:"keluhan_utama"`
	Rps                   string `form:"rps" json:"rps"`
	Rpk                   string `form:"rpk" json:"rpk"`
	Rpd                   string `form:"rpd" json:"rpd"`
	Rpo                   string `form:"rpo" json:"rpo"`
	RiwayatKebiasaan      string `form:"riwayat_kebiasaan" json:"riwayat_kebiasaan"`
	RiwayatOperasiUrologi string `form:"riwayat_operasi_urologi" json:"riwayat_operasi_urologi"`
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
	Thoraks               string `form:"thoraks" json:"thoraks"`
	KeteranganThoraks     string `form:"keterangan_thoraks" json:"keterangan_thoraks"`
	Abdomen               string `form:"abdomen" json:"abdomen"`
	KeteranganAbdomen     string `form:"keterangan_abdomen" json:"keterangan_abdomen"`
	Ekstrimitas           string `form:"ekstrimitas" json:"ekstrimitas"`
	KeteranganEkstrimitas string `form:"keterangan_ekstrimitas" json:"keterangan_ekstrimitas"`
	NyeriKetokCva         string `form:"nyeri_ketok_cva" json:"nyeri_ketok_cva"`
	GenitaliaEksternal    string `form:"genitalia_eksternal" json:"genitalia_eksternal"`
	ColokDubur            string `form:"colok_dubur" json:"colok_dubur"`
	Lainnya               string `form:"lainnya" json:"lainnya"`
	Urinalisis            string `form:"urinalisis" json:"urinalisis"`
	Darah                 string `form:"darah" json:"darah"`
	UsgUrologi            string `form:"usg_urologi" json:"usg_urologi"`
	Radiologi             string `form:"radiologi" json:"radiologi"`
	PenunjangLain         string `form:"penunjang_lain" json:"penunjang_lain"`
	Diagnosis             string `form:"diagnosis" json:"diagnosis"`
	Diagnosis2            string `form:"diagnosis2" json:"diagnosis2"`
	Permasalahan          string `form:"permasalahan" json:"permasalahan"`
	Terapi                string `form:"terapi" json:"terapi"`
	Tindakan              string `form:"tindakan" json:"tindakan"`
	Edukasi               string `form:"edukasi" json:"edukasi"`
}

func penilaianMedisRalanUrologiRules() map[string]any {
	rules := map[string]any{
		"tanggal":                 "required|date",
		"kd_dokter":               "required|string|max_len:20",
		"anamnesis":               "required|in:Autoanamnesis,Alloanamnesis",
		"hubungan":                "string|max_len:30",
		"keluhan_utama":           "string|max_len:2000",
		"rps":                     "string|max_len:2000",
		"rpk":                     "string|max_len:1000",
		"rpd":                     "string|max_len:1000",
		"rpo":                     "string|max_len:1000",
		"riwayat_kebiasaan":       "string|max_len:1000",
		"riwayat_operasi_urologi": "string|max_len:1000",
		"alergi":                  "string|max_len:50",
		"td":                      "string|max_len:8",
		"bb":                      "string|max_len:5",
		"tb":                      "string|max_len:5",
		"suhu":                    "string|max_len:5",
		"nadi":                    "string|max_len:5",
		"rr":                      "string|max_len:5",
		"keadaan_umum":            "in:Sehat,Sakit Ringan,Sakit Sedang,Sakit Berat",
		"nyeri":                   "string|max_len:50",
		"status_nutrisi":          "string|max_len:50",
		"thoraks":                 "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_thoraks":      "string|max_len:50",
		"abdomen":                 "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_abdomen":      "string|max_len:50",
		"ekstrimitas":             "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_ekstrimitas":  "string|max_len:50",
		"nyeri_ketok_cva":         "string|max_len:100",
		"genitalia_eksternal":     "string|max_len:100",
		"colok_dubur":             "string|max_len:100",
		"lainnya":                 "string|max_len:1000",
		"urinalisis":              "string|max_len:500",
		"darah":                   "string|max_len:500",
		"usg_urologi":             "string|max_len:500",
		"radiologi":               "string|max_len:500",
		"penunjang_lain":          "string|max_len:500",
		"diagnosis":               "string|max_len:500",
		"diagnosis2":              "string|max_len:500",
		"permasalahan":            "string|max_len:500",
		"terapi":                  "string|max_len:500",
		"tindakan":                "string|max_len:500",
		"edukasi":                 "string|max_len:500",
	}
	return rules
}

// PenilaianMedisRalanUrologiStore simpan penilaian awal medis ralan urologi; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRalanUrologiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRalanUrologiData
}

func (r *PenilaianMedisRalanUrologiStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanUrologiStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRalanUrologiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRalanUrologiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianMedisRalanUrologiStore) Payload() PenilaianMedisRalanUrologiData {
	return r.PenilaianMedisRalanUrologiData
}

func (r *PenilaianMedisRalanUrologiStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisRalanUrologiUpdate ubah penilaian awal medis ralan urologi (PUT); kunci lewat query string.
type PenilaianMedisRalanUrologiUpdate struct {
	PenilaianMedisRalanUrologiData
}

func (r *PenilaianMedisRalanUrologiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanUrologiUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRalanUrologiRules()
}

func (r *PenilaianMedisRalanUrologiUpdate) Payload() PenilaianMedisRalanUrologiData {
	return r.PenilaianMedisRalanUrologiData
}

func (r *PenilaianMedisRalanUrologiUpdate) DetailValues() map[string][]string { return nil }
