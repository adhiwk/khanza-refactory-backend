package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRalanGeriatriData isian penilaian awal medis ralan geriatri.
type PenilaianMedisRalanGeriatriData struct {
	Tanggal                 string `form:"tanggal" json:"tanggal"`
	KdDokter                string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis               string `form:"anamnesis" json:"anamnesis"`
	Hubungan                string `form:"hubungan" json:"hubungan"`
	KeluhanUtama            string `form:"keluhan_utama" json:"keluhan_utama"`
	Rps                     string `form:"rps" json:"rps"`
	Rpd                     string `form:"rpd" json:"rpd"`
	Rpo                     string `form:"rpo" json:"rpo"`
	Alergi                  string `form:"alergi" json:"alergi"`
	TulangBelakang          string `form:"tulang_belakang" json:"tulang_belakang"`
	Td                      string `form:"td" json:"td"`
	Nadi                    string `form:"nadi" json:"nadi"`
	Suhu                    string `form:"suhu" json:"suhu"`
	Rr                      string `form:"rr" json:"rr"`
	KondisiUmum             string `form:"kondisi_umum" json:"kondisi_umum"`
	StatusPsikologisGds     string `form:"status_psikologis_gds" json:"status_psikologis_gds"`
	KondisiSosial           string `form:"kondisi_sosial" json:"kondisi_sosial"`
	StatusKognitifMmse      string `form:"status_kognitif_mmse" json:"status_kognitif_mmse"`
	Kepala                  string `form:"kepala" json:"kepala"`
	KeteranganKepala        string `form:"keterangan_kepala" json:"keterangan_kepala"`
	Thoraks                 string `form:"thoraks" json:"thoraks"`
	KeteranganThoraks       string `form:"keterangan_thoraks" json:"keterangan_thoraks"`
	Abdomen                 string `form:"abdomen" json:"abdomen"`
	KeteranganAbdomen       string `form:"keterangan_abdomen" json:"keterangan_abdomen"`
	Ekstremitas             string `form:"ekstremitas" json:"ekstremitas"`
	KeteranganEkstremitas   string `form:"keterangan_ekstremitas" json:"keterangan_ekstremitas"`
	IntegumentKebersihan    string `form:"Integument_kebersihan" json:"Integument_kebersihan"`
	IntegumentWarna         string `form:"Integument_warna" json:"Integument_warna"`
	IntegumentKelembaban    string `form:"Integument_kelembaban" json:"Integument_kelembaban"`
	IntegumentGangguanKulit string `form:"Integument_gangguan_kulit" json:"Integument_gangguan_kulit"`
	StatusFungsional        string `form:"status_fungsional" json:"status_fungsional"`
	SkriningJatuh           string `form:"skrining_jatuh" json:"skrining_jatuh"`
	StatusNutrisi           string `form:"status_nutrisi" json:"status_nutrisi"`
	Lainnya                 string `form:"lainnya" json:"lainnya"`
	Lab                     string `form:"lab" json:"lab"`
	Rad                     string `form:"rad" json:"rad"`
	Pemeriksaan             string `form:"pemeriksaan" json:"pemeriksaan"`
	Diagnosis               string `form:"diagnosis" json:"diagnosis"`
	Diagnosis2              string `form:"diagnosis2" json:"diagnosis2"`
	Permasalahan            string `form:"permasalahan" json:"permasalahan"`
	Terapi                  string `form:"terapi" json:"terapi"`
	Tindakan                string `form:"tindakan" json:"tindakan"`
	Edukasi                 string `form:"edukasi" json:"edukasi"`
}

func penilaianMedisRalanGeriatriRules() map[string]any {
	rules := map[string]any{
		"tanggal":                   "required|date",
		"kd_dokter":                 "required|string|max_len:20",
		"anamnesis":                 "required|in:Autoanamnesis,Alloanamnesis",
		"hubungan":                  "string|max_len:30",
		"keluhan_utama":             "string|max_len:2000",
		"rps":                       "string|max_len:2000",
		"rpd":                       "string|max_len:1000",
		"rpo":                       "string|max_len:1000",
		"alergi":                    "string|max_len:50",
		"tulang_belakang":           "required|in:Tegap,Membungkuk,Kifosis,Skoliosis,Lordosis",
		"td":                        "string|max_len:8",
		"nadi":                      "string|max_len:5",
		"suhu":                      "string|max_len:5",
		"rr":                        "string|max_len:5",
		"kondisi_umum":              "string|max_len:1000",
		"status_psikologis_gds":     "required|in:Skor 1-4 Tidak Ada Depresi,Skor Antara 5-9 Menunjukkan Kemungkinan Besar Depresi,Skor 10 Atau Lebih Menunjukkan Depresi",
		"kondisi_sosial":            "string|max_len:500",
		"status_kognitif_mmse":      "required|in:24-30 : Tidak Ada Gangguan Kognitif,18-23 : Gangguan Kognitif Sedang,0-17 : Gangguan Kognitif Berat",
		"kepala":                    "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_kepala":         "string|max_len:100",
		"thoraks":                   "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_thoraks":        "string|max_len:100",
		"abdomen":                   "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_abdomen":        "string|max_len:100",
		"ekstremitas":               "required|in:Normal,Abnormal,Tidak Diperiksa",
		"keterangan_ekstremitas":    "string|max_len:100",
		"Integument_kebersihan":     "required|in:Normal,Abnormal",
		"Integument_warna":          "required|in:Normal,Pucat,Sianosis,Lain-lain",
		"Integument_kelembaban":     "required|in:Kering,Lembab",
		"Integument_gangguan_kulit": "required|in:Normal,Rash,Luka,Memar,Ptekie,Bula",
		"status_fungsional":         "required|in:20 : Mandiri (A),12-19 : Ketergantungan Ringan (B),9-11 : Ketergantungan Sedang (B),5-8 : Ketergantungan Berat (C),0-4 : Ketergantungan Total (C)",
		"skrining_jatuh":            "required|in:Risiko Rendah Skor 0-5,Risiko Sedang Skor 6-16,Risiko Tinggi Skor 17-30",
		"status_nutrisi":            "required|in:Skor 12-14 : Status Gizi Normal,Skor 8-11 : Berisiko Malnutrisi,Skor 0-7 : Malnutrisi",
		"lainnya":                   "string|max_len:1000",
		"lab":                       "string|max_len:500",
		"rad":                       "string|max_len:500",
		"pemeriksaan":               "string|max_len:500",
		"diagnosis":                 "string|max_len:500",
		"diagnosis2":                "string|max_len:500",
		"permasalahan":              "string|max_len:500",
		"terapi":                    "string|max_len:500",
		"tindakan":                  "string|max_len:500",
		"edukasi":                   "string|max_len:500",
	}
	return rules
}

// PenilaianMedisRalanGeriatriStore simpan penilaian awal medis ralan geriatri; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRalanGeriatriStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRalanGeriatriData
}

func (r *PenilaianMedisRalanGeriatriStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanGeriatriStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRalanGeriatriRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRalanGeriatriStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianMedisRalanGeriatriStore) Payload() PenilaianMedisRalanGeriatriData {
	return r.PenilaianMedisRalanGeriatriData
}

func (r *PenilaianMedisRalanGeriatriStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisRalanGeriatriUpdate ubah penilaian awal medis ralan geriatri (PUT); kunci lewat query string.
type PenilaianMedisRalanGeriatriUpdate struct {
	PenilaianMedisRalanGeriatriData
}

func (r *PenilaianMedisRalanGeriatriUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanGeriatriUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRalanGeriatriRules()
}

func (r *PenilaianMedisRalanGeriatriUpdate) Payload() PenilaianMedisRalanGeriatriData {
	return r.PenilaianMedisRalanGeriatriData
}

func (r *PenilaianMedisRalanGeriatriUpdate) DetailValues() map[string][]string { return nil }
