package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ChecklistKriteriaMasukIcuData isian checklist kriteria masuk ICU.
type ChecklistKriteriaMasukIcuData struct {
	Prioritas11                   string `form:"prioritas1_1" json:"prioritas1_1"`
	Prioritas12                   string `form:"prioritas1_2" json:"prioritas1_2"`
	Prioritas13                   string `form:"prioritas1_3" json:"prioritas1_3"`
	Prioritas14                   string `form:"prioritas1_4" json:"prioritas1_4"`
	Prioritas15                   string `form:"prioritas1_5" json:"prioritas1_5"`
	Prioritas16                   string `form:"prioritas1_6" json:"prioritas1_6"`
	Prioritas21                   string `form:"prioritas2_1" json:"prioritas2_1"`
	Prioritas22                   string `form:"prioritas2_2" json:"prioritas2_2"`
	Prioritas23                   string `form:"prioritas2_3" json:"prioritas2_3"`
	Prioritas24                   string `form:"prioritas2_4" json:"prioritas2_4"`
	Prioritas25                   string `form:"prioritas2_5" json:"prioritas2_5"`
	Prioritas26                   string `form:"prioritas2_6" json:"prioritas2_6"`
	Prioritas27                   string `form:"prioritas2_7" json:"prioritas2_7"`
	Prioritas28                   string `form:"prioritas2_8" json:"prioritas2_8"`
	Prioritas31                   string `form:"prioritas3_1" json:"prioritas3_1"`
	Prioritas32                   string `form:"prioritas3_2" json:"prioritas3_2"`
	Prioritas33                   string `form:"prioritas3_3" json:"prioritas3_3"`
	Prioritas34                   string `form:"prioritas3_4" json:"prioritas3_4"`
	KriteriaFisiologisTandaVital1 string `form:"kriteria_fisiologis_tanda_vital_1" json:"kriteria_fisiologis_tanda_vital_1"`
	KriteriaFisiologisTandaVital2 string `form:"kriteria_fisiologis_tanda_vital_2" json:"kriteria_fisiologis_tanda_vital_2"`
	KriteriaFisiologisTandaVital3 string `form:"kriteria_fisiologis_tanda_vital_3" json:"kriteria_fisiologis_tanda_vital_3"`
	KriteriaFisiologisTandaVital4 string `form:"kriteria_fisiologis_tanda_vital_4" json:"kriteria_fisiologis_tanda_vital_4"`
	KriteriaFisiologisTandaVital5 string `form:"kriteria_fisiologis_tanda_vital_5" json:"kriteria_fisiologis_tanda_vital_5"`
	KriteriaFisiologisLaborat1    string `form:"kriteria_fisiologis_laborat_1" json:"kriteria_fisiologis_laborat_1"`
	KriteriaFisiologisLaborat2    string `form:"kriteria_fisiologis_laborat_2" json:"kriteria_fisiologis_laborat_2"`
	KriteriaFisiologisLaborat3    string `form:"kriteria_fisiologis_laborat_3" json:"kriteria_fisiologis_laborat_3"`
	KriteriaFisiologisLaborat4    string `form:"kriteria_fisiologis_laborat_4" json:"kriteria_fisiologis_laborat_4"`
	KriteriaFisiologisLaborat5    string `form:"kriteria_fisiologis_laborat_5" json:"kriteria_fisiologis_laborat_5"`
	KriteriaFisiologisLaborat6    string `form:"kriteria_fisiologis_laborat_6" json:"kriteria_fisiologis_laborat_6"`
	KriteriaFisiologisRadiologi1  string `form:"kriteria_fisiologis_radiologi_1" json:"kriteria_fisiologis_radiologi_1"`
	KriteriaFisiologisRadiologi2  string `form:"kriteria_fisiologis_radiologi_2" json:"kriteria_fisiologis_radiologi_2"`
	KriteriaFisiologisKlinis1     string `form:"kriteria_fisiologis_klinis_1" json:"kriteria_fisiologis_klinis_1"`
	KriteriaFisiologisKlinis2     string `form:"kriteria_fisiologis_klinis_2" json:"kriteria_fisiologis_klinis_2"`
	KriteriaFisiologisKlinis3     string `form:"kriteria_fisiologis_klinis_3" json:"kriteria_fisiologis_klinis_3"`
	KriteriaFisiologisKlinis4     string `form:"kriteria_fisiologis_klinis_4" json:"kriteria_fisiologis_klinis_4"`
	KriteriaFisiologisKlinis5     string `form:"kriteria_fisiologis_klinis_5" json:"kriteria_fisiologis_klinis_5"`
	KriteriaFisiologisKlinis6     string `form:"kriteria_fisiologis_klinis_6" json:"kriteria_fisiologis_klinis_6"`
	KriteriaFisiologisKlinis7     string `form:"kriteria_fisiologis_klinis_7" json:"kriteria_fisiologis_klinis_7"`
	KriteriaFisiologisKlinis8     string `form:"kriteria_fisiologis_klinis_8" json:"kriteria_fisiologis_klinis_8"`
	Nik                           string `form:"nik" json:"nik"`
}

func checklistKriteriaMasukIcuRules() map[string]any {
	rules := map[string]any{
		"prioritas1_1":                      "required|in:Ya,Tidak",
		"prioritas1_2":                      "required|in:Ya,Tidak",
		"prioritas1_3":                      "required|in:Ya,Tidak",
		"prioritas1_4":                      "required|in:Ya,Tidak",
		"prioritas1_5":                      "required|in:Ya,Tidak",
		"prioritas1_6":                      "required|in:Ya,Tidak",
		"prioritas2_1":                      "required|in:Ya,Tidak",
		"prioritas2_2":                      "required|in:Ya,Tidak",
		"prioritas2_3":                      "required|in:Ya,Tidak",
		"prioritas2_4":                      "required|in:Ya,Tidak",
		"prioritas2_5":                      "required|in:Ya,Tidak",
		"prioritas2_6":                      "required|in:Ya,Tidak",
		"prioritas2_7":                      "required|in:Ya,Tidak",
		"prioritas2_8":                      "required|in:Ya,Tidak",
		"prioritas3_1":                      "required|in:Ya,Tidak",
		"prioritas3_2":                      "required|in:Ya,Tidak",
		"prioritas3_3":                      "required|in:Ya,Tidak",
		"prioritas3_4":                      "required|in:Ya,Tidak",
		"kriteria_fisiologis_tanda_vital_1": "required|in:Ya,Tidak",
		"kriteria_fisiologis_tanda_vital_2": "required|in:Ya,Tidak",
		"kriteria_fisiologis_tanda_vital_3": "required|in:Ya,Tidak",
		"kriteria_fisiologis_tanda_vital_4": "required|in:Ya,Tidak",
		"kriteria_fisiologis_tanda_vital_5": "required|in:Ya,Tidak",
		"kriteria_fisiologis_laborat_1":     "required|in:Ya,Tidak",
		"kriteria_fisiologis_laborat_2":     "required|in:Ya,Tidak",
		"kriteria_fisiologis_laborat_3":     "required|in:Ya,Tidak",
		"kriteria_fisiologis_laborat_4":     "required|in:Ya,Tidak",
		"kriteria_fisiologis_laborat_5":     "required|in:Ya,Tidak",
		"kriteria_fisiologis_laborat_6":     "required|in:Ya,Tidak",
		"kriteria_fisiologis_radiologi_1":   "required|in:Ya,Tidak",
		"kriteria_fisiologis_radiologi_2":   "required|in:Ya,Tidak",
		"kriteria_fisiologis_klinis_1":      "required|in:Ya,Tidak",
		"kriteria_fisiologis_klinis_2":      "required|in:Ya,Tidak",
		"kriteria_fisiologis_klinis_3":      "required|in:Ya,Tidak",
		"kriteria_fisiologis_klinis_4":      "required|in:Ya,Tidak",
		"kriteria_fisiologis_klinis_5":      "required|in:Ya,Tidak",
		"kriteria_fisiologis_klinis_6":      "required|in:Ya,Tidak",
		"kriteria_fisiologis_klinis_7":      "required|in:Ya,Tidak",
		"kriteria_fisiologis_klinis_8":      "required|in:Ya,Tidak",
		"nik":                               "string|max_len:20",
	}
	return rules
}

// ChecklistKriteriaMasukIcuStore simpan checklist kriteria masuk ICU; kolom waktu kunci kosong = sekarang.
type ChecklistKriteriaMasukIcuStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	ChecklistKriteriaMasukIcuData
}

func (r *ChecklistKriteriaMasukIcuStore) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaMasukIcuStore) Rules(ctx http.Context) map[string]any {
	rules := checklistKriteriaMasukIcuRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *ChecklistKriteriaMasukIcuStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *ChecklistKriteriaMasukIcuStore) Payload() ChecklistKriteriaMasukIcuData {
	return r.ChecklistKriteriaMasukIcuData
}

func (r *ChecklistKriteriaMasukIcuStore) DetailValues() map[string][]string { return nil }

// ChecklistKriteriaMasukIcuUpdate ubah checklist kriteria masuk ICU (PUT); kunci lewat query string.
type ChecklistKriteriaMasukIcuUpdate struct {
	ChecklistKriteriaMasukIcuData
}

func (r *ChecklistKriteriaMasukIcuUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaMasukIcuUpdate) Rules(ctx http.Context) map[string]any {
	return checklistKriteriaMasukIcuRules()
}

func (r *ChecklistKriteriaMasukIcuUpdate) Payload() ChecklistKriteriaMasukIcuData {
	return r.ChecklistKriteriaMasukIcuData
}

func (r *ChecklistKriteriaMasukIcuUpdate) DetailValues() map[string][]string { return nil }
