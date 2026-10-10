package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ChecklistKriteriaMasukHcuData isian checklist kriteria masuk HCU.
type ChecklistKriteriaMasukHcuData struct {
	Kardiologi1 string `form:"kardiologi1" json:"kardiologi1"`
	Kardiologi2 string `form:"kardiologi2" json:"kardiologi2"`
	Kardiologi3 string `form:"kardiologi3" json:"kardiologi3"`
	Kardiologi4 string `form:"kardiologi4" json:"kardiologi4"`
	Kardiologi5 string `form:"kardiologi5" json:"kardiologi5"`
	Kardiologi6 string `form:"kardiologi6" json:"kardiologi6"`
	Pernapasan1 string `form:"pernapasan1" json:"pernapasan1"`
	Pernapasan2 string `form:"pernapasan2" json:"pernapasan2"`
	Pernapasan3 string `form:"pernapasan3" json:"pernapasan3"`
	Syaraf1     string `form:"syaraf1" json:"syaraf1"`
	Syaraf2     string `form:"syaraf2" json:"syaraf2"`
	Syaraf3     string `form:"syaraf3" json:"syaraf3"`
	Syaraf4     string `form:"syaraf4" json:"syaraf4"`
	Pencernaan1 string `form:"pencernaan1" json:"pencernaan1"`
	Pencernaan2 string `form:"pencernaan2" json:"pencernaan2"`
	Pencernaan3 string `form:"pencernaan3" json:"pencernaan3"`
	Pencernaan4 string `form:"pencernaan4" json:"pencernaan4"`
	Pembedahan1 string `form:"pembedahan1" json:"pembedahan1"`
	Pembedahan2 string `form:"pembedahan2" json:"pembedahan2"`
	Hematologi1 string `form:"hematologi1" json:"hematologi1"`
	Hematologi2 string `form:"hematologi2" json:"hematologi2"`
	Infeksi     string `form:"infeksi" json:"infeksi"`
	Nik         string `form:"nik" json:"nik"`
}

func checklistKriteriaMasukHcuRules() map[string]any {
	rules := map[string]any{
		"kardiologi1": "required|in:Ya,Tidak",
		"kardiologi2": "required|in:Ya,Tidak",
		"kardiologi3": "required|in:Ya,Tidak",
		"kardiologi4": "required|in:Ya,Tidak",
		"kardiologi5": "required|in:Ya,Tidak",
		"kardiologi6": "required|in:Ya,Tidak",
		"pernapasan1": "required|in:Ya,Tidak",
		"pernapasan2": "required|in:Ya,Tidak",
		"pernapasan3": "required|in:Ya,Tidak",
		"syaraf1":     "required|in:Ya,Tidak",
		"syaraf2":     "required|in:Ya,Tidak",
		"syaraf3":     "required|in:Ya,Tidak",
		"syaraf4":     "required|in:Ya,Tidak",
		"pencernaan1": "required|in:Ya,Tidak",
		"pencernaan2": "required|in:Ya,Tidak",
		"pencernaan3": "required|in:Ya,Tidak",
		"pencernaan4": "required|in:Ya,Tidak",
		"pembedahan1": "required|in:Ya,Tidak",
		"pembedahan2": "required|in:Ya,Tidak",
		"hematologi1": "required|in:Ya,Tidak",
		"hematologi2": "required|in:Ya,Tidak",
		"infeksi":     "required|in:Ya,Tidak",
		"nik":         "string|max_len:20",
	}
	return rules
}

// ChecklistKriteriaMasukHcuStore simpan checklist kriteria masuk HCU; kolom waktu kunci kosong = sekarang.
type ChecklistKriteriaMasukHcuStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	ChecklistKriteriaMasukHcuData
}

func (r *ChecklistKriteriaMasukHcuStore) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaMasukHcuStore) Rules(ctx http.Context) map[string]any {
	rules := checklistKriteriaMasukHcuRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *ChecklistKriteriaMasukHcuStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *ChecklistKriteriaMasukHcuStore) Payload() ChecklistKriteriaMasukHcuData {
	return r.ChecklistKriteriaMasukHcuData
}

func (r *ChecklistKriteriaMasukHcuStore) DetailValues() map[string][]string { return nil }

// ChecklistKriteriaMasukHcuUpdate ubah checklist kriteria masuk HCU (PUT); kunci lewat query string.
type ChecklistKriteriaMasukHcuUpdate struct {
	ChecklistKriteriaMasukHcuData
}

func (r *ChecklistKriteriaMasukHcuUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaMasukHcuUpdate) Rules(ctx http.Context) map[string]any {
	return checklistKriteriaMasukHcuRules()
}

func (r *ChecklistKriteriaMasukHcuUpdate) Payload() ChecklistKriteriaMasukHcuData {
	return r.ChecklistKriteriaMasukHcuData
}

func (r *ChecklistKriteriaMasukHcuUpdate) DetailValues() map[string][]string { return nil }
