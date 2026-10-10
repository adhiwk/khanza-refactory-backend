package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ChecklistKriteriaKeluarHcuData isian checklist kriteria keluar HCU.
type ChecklistKriteriaKeluarHcuData struct {
	Kriteria1  string `form:"kriteria1" json:"kriteria1"`
	Kriteria2  string `form:"kriteria2" json:"kriteria2"`
	Kriteria3  string `form:"kriteria3" json:"kriteria3"`
	Kriteria4  string `form:"kriteria4" json:"kriteria4"`
	Kriteria5  string `form:"kriteria5" json:"kriteria5"`
	Kriteria6  string `form:"kriteria6" json:"kriteria6"`
	Kriteria7  string `form:"kriteria7" json:"kriteria7"`
	Kriteria8  string `form:"kriteria8" json:"kriteria8"`
	Kriteria9  string `form:"kriteria9" json:"kriteria9"`
	Kriteria10 string `form:"kriteria10" json:"kriteria10"`
	Kriteria11 string `form:"kriteria11" json:"kriteria11"`
	Kriteria12 string `form:"kriteria12" json:"kriteria12"`
	Nik        string `form:"nik" json:"nik"`
}

func checklistKriteriaKeluarHcuRules() map[string]any {
	rules := map[string]any{
		"kriteria1":  "required|in:Ya,Tidak",
		"kriteria2":  "required|in:Ya,Tidak",
		"kriteria3":  "required|in:Ya,Tidak",
		"kriteria4":  "required|in:Ya,Tidak",
		"kriteria5":  "required|in:Ya,Tidak",
		"kriteria6":  "required|in:Ya,Tidak",
		"kriteria7":  "required|in:Ya,Tidak",
		"kriteria8":  "required|in:Ya,Tidak",
		"kriteria9":  "required|in:Ya,Tidak",
		"kriteria10": "required|in:Ya,Tidak",
		"kriteria11": "required|in:Ya,Tidak",
		"kriteria12": "required|in:Ya,Tidak",
		"nik":        "string|max_len:20",
	}
	return rules
}

// ChecklistKriteriaKeluarHcuStore simpan checklist kriteria keluar HCU; kolom waktu kunci kosong = sekarang.
type ChecklistKriteriaKeluarHcuStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	ChecklistKriteriaKeluarHcuData
}

func (r *ChecklistKriteriaKeluarHcuStore) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaKeluarHcuStore) Rules(ctx http.Context) map[string]any {
	rules := checklistKriteriaKeluarHcuRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *ChecklistKriteriaKeluarHcuStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *ChecklistKriteriaKeluarHcuStore) Payload() ChecklistKriteriaKeluarHcuData {
	return r.ChecklistKriteriaKeluarHcuData
}

func (r *ChecklistKriteriaKeluarHcuStore) DetailValues() map[string][]string { return nil }

// ChecklistKriteriaKeluarHcuUpdate ubah checklist kriteria keluar HCU (PUT); kunci lewat query string.
type ChecklistKriteriaKeluarHcuUpdate struct {
	ChecklistKriteriaKeluarHcuData
}

func (r *ChecklistKriteriaKeluarHcuUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaKeluarHcuUpdate) Rules(ctx http.Context) map[string]any {
	return checklistKriteriaKeluarHcuRules()
}

func (r *ChecklistKriteriaKeluarHcuUpdate) Payload() ChecklistKriteriaKeluarHcuData {
	return r.ChecklistKriteriaKeluarHcuData
}

func (r *ChecklistKriteriaKeluarHcuUpdate) DetailValues() map[string][]string { return nil }
