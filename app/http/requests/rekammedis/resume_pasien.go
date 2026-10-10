package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ResumePasienData isian resume pasien.
type ResumePasienData struct {
	KdDokter             string `form:"kd_dokter" json:"kd_dokter"`
	KeluhanUtama         string `form:"keluhan_utama" json:"keluhan_utama"`
	JalannyaPenyakit     string `form:"jalannya_penyakit" json:"jalannya_penyakit"`
	PemeriksaanPenunjang string `form:"pemeriksaan_penunjang" json:"pemeriksaan_penunjang"`
	HasilLaborat         string `form:"hasil_laborat" json:"hasil_laborat"`
	DiagnosaUtama        string `form:"diagnosa_utama" json:"diagnosa_utama"`
	KdDiagnosaUtama      string `form:"kd_diagnosa_utama" json:"kd_diagnosa_utama"`
	DiagnosaSekunder     string `form:"diagnosa_sekunder" json:"diagnosa_sekunder"`
	KdDiagnosaSekunder   string `form:"kd_diagnosa_sekunder" json:"kd_diagnosa_sekunder"`
	DiagnosaSekunder2    string `form:"diagnosa_sekunder2" json:"diagnosa_sekunder2"`
	KdDiagnosaSekunder2  string `form:"kd_diagnosa_sekunder2" json:"kd_diagnosa_sekunder2"`
	DiagnosaSekunder3    string `form:"diagnosa_sekunder3" json:"diagnosa_sekunder3"`
	KdDiagnosaSekunder3  string `form:"kd_diagnosa_sekunder3" json:"kd_diagnosa_sekunder3"`
	DiagnosaSekunder4    string `form:"diagnosa_sekunder4" json:"diagnosa_sekunder4"`
	KdDiagnosaSekunder4  string `form:"kd_diagnosa_sekunder4" json:"kd_diagnosa_sekunder4"`
	ProsedurUtama        string `form:"prosedur_utama" json:"prosedur_utama"`
	KdProsedurUtama      string `form:"kd_prosedur_utama" json:"kd_prosedur_utama"`
	ProsedurSekunder     string `form:"prosedur_sekunder" json:"prosedur_sekunder"`
	KdProsedurSekunder   string `form:"kd_prosedur_sekunder" json:"kd_prosedur_sekunder"`
	ProsedurSekunder2    string `form:"prosedur_sekunder2" json:"prosedur_sekunder2"`
	KdProsedurSekunder2  string `form:"kd_prosedur_sekunder2" json:"kd_prosedur_sekunder2"`
	ProsedurSekunder3    string `form:"prosedur_sekunder3" json:"prosedur_sekunder3"`
	KdProsedurSekunder3  string `form:"kd_prosedur_sekunder3" json:"kd_prosedur_sekunder3"`
	KondisiPulang        string `form:"kondisi_pulang" json:"kondisi_pulang"`
	ObatPulang           string `form:"obat_pulang" json:"obat_pulang"`
}

func resumePasienRules() map[string]any {
	rules := map[string]any{
		"kd_dokter":             "required|string|max_len:20",
		"keluhan_utama":         "string",
		"jalannya_penyakit":     "string",
		"pemeriksaan_penunjang": "string",
		"hasil_laborat":         "string",
		"diagnosa_utama":        "string|max_len:80",
		"kd_diagnosa_utama":     "string|max_len:10",
		"diagnosa_sekunder":     "string|max_len:80",
		"kd_diagnosa_sekunder":  "string|max_len:10",
		"diagnosa_sekunder2":    "string|max_len:80",
		"kd_diagnosa_sekunder2": "string|max_len:10",
		"diagnosa_sekunder3":    "string|max_len:80",
		"kd_diagnosa_sekunder3": "string|max_len:10",
		"diagnosa_sekunder4":    "string|max_len:80",
		"kd_diagnosa_sekunder4": "string|max_len:10",
		"prosedur_utama":        "string|max_len:80",
		"kd_prosedur_utama":     "string|max_len:8",
		"prosedur_sekunder":     "string|max_len:80",
		"kd_prosedur_sekunder":  "string|max_len:8",
		"prosedur_sekunder2":    "string|max_len:80",
		"kd_prosedur_sekunder2": "string|max_len:8",
		"prosedur_sekunder3":    "string|max_len:80",
		"kd_prosedur_sekunder3": "string|max_len:8",
		"kondisi_pulang":        "required|in:Hidup,Meninggal",
		"obat_pulang":           "string",
	}
	return rules
}

// ResumePasienStore simpan resume pasien; kolom waktu kunci kosong = sekarang.
type ResumePasienStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	ResumePasienData
}

func (r *ResumePasienStore) Authorize(ctx http.Context) error { return nil }

func (r *ResumePasienStore) Rules(ctx http.Context) map[string]any {
	rules := resumePasienRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *ResumePasienStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *ResumePasienStore) Payload() ResumePasienData { return r.ResumePasienData }

func (r *ResumePasienStore) DetailValues() map[string][]string { return nil }

// ResumePasienUpdate ubah resume pasien (PUT); kunci lewat query string.
type ResumePasienUpdate struct {
	ResumePasienData
}

func (r *ResumePasienUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ResumePasienUpdate) Rules(ctx http.Context) map[string]any { return resumePasienRules() }

func (r *ResumePasienUpdate) Payload() ResumePasienData { return r.ResumePasienData }

func (r *ResumePasienUpdate) DetailValues() map[string][]string { return nil }
