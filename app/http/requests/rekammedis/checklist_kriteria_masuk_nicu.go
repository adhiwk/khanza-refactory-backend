package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ChecklistKriteriaMasukNicuData isian checklist kriteria masuk NICU.
type ChecklistKriteriaMasukNicuData struct {
	Respirasi1   string `form:"respirasi1" json:"respirasi1"`
	Respirasi2   string `form:"respirasi2" json:"respirasi2"`
	Respirasi3   string `form:"respirasi3" json:"respirasi3"`
	Respirasi4   string `form:"respirasi4" json:"respirasi4"`
	Prematur1    string `form:"prematur1" json:"prematur1"`
	Prematur2    string `form:"prematur2" json:"prematur2"`
	Prematur3    string `form:"prematur3" json:"prematur3"`
	Kardio1      string `form:"kardio1" json:"kardio1"`
	Kardio2      string `form:"kardio2" json:"kardio2"`
	Kardio3      string `form:"kardio3" json:"kardio3"`
	Neuro1       string `form:"neuro1" json:"neuro1"`
	Neuro2       string `form:"neuro2" json:"neuro2"`
	Neuro3       string `form:"neuro3" json:"neuro3"`
	Metabolik1   string `form:"metabolik1" json:"metabolik1"`
	Metabolik2   string `form:"metabolik2" json:"metabolik2"`
	Metabolik3   string `form:"metabolik3" json:"metabolik3"`
	Kondisilain1 string `form:"kondisilain1" json:"kondisilain1"`
	Kondisilain2 string `form:"kondisilain2" json:"kondisilain2"`
	Kondisilain3 string `form:"kondisilain3" json:"kondisilain3"`
	Kondisilain4 string `form:"kondisilain4" json:"kondisilain4"`
	Keputusan    string `form:"keputusan" json:"keputusan"`
	Keterangan   string `form:"keterangan" json:"keterangan"`
	Nik          string `form:"nik" json:"nik"`
}

func checklistKriteriaMasukNicuRules() map[string]any {
	rules := map[string]any{
		"respirasi1":   "required|in:Ya,Tidak",
		"respirasi2":   "required|in:Ya,Tidak",
		"respirasi3":   "required|in:Ya,Tidak",
		"respirasi4":   "required|in:Ya,Tidak",
		"prematur1":    "required|in:Ya,Tidak",
		"prematur2":    "required|in:Ya,Tidak",
		"prematur3":    "required|in:Ya,Tidak",
		"kardio1":      "required|in:Ya,Tidak",
		"kardio2":      "required|in:Ya,Tidak",
		"kardio3":      "required|in:Ya,Tidak",
		"neuro1":       "required|in:Ya,Tidak",
		"neuro2":       "required|in:Ya,Tidak",
		"neuro3":       "required|in:Ya,Tidak",
		"metabolik1":   "required|in:Ya,Tidak",
		"metabolik2":   "required|in:Ya,Tidak",
		"metabolik3":   "required|in:Ya,Tidak",
		"kondisilain1": "required|in:Ya,Tidak",
		"kondisilain2": "required|in:Ya,Tidak",
		"kondisilain3": "required|in:Ya,Tidak",
		"kondisilain4": "required|in:Ya,Tidak",
		"keputusan":    "required|in:Diterima Di NICU,Tidak Diterima - Dirawat Di Ruang Perawatan Bayi Lain",
		"keterangan":   "string|max_len:50",
		"nik":          "string|max_len:20",
	}
	return rules
}

// ChecklistKriteriaMasukNicuStore simpan checklist kriteria masuk NICU; kolom waktu kunci kosong = sekarang.
type ChecklistKriteriaMasukNicuStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	ChecklistKriteriaMasukNicuData
}

func (r *ChecklistKriteriaMasukNicuStore) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaMasukNicuStore) Rules(ctx http.Context) map[string]any {
	rules := checklistKriteriaMasukNicuRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *ChecklistKriteriaMasukNicuStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *ChecklistKriteriaMasukNicuStore) Payload() ChecklistKriteriaMasukNicuData {
	return r.ChecklistKriteriaMasukNicuData
}

func (r *ChecklistKriteriaMasukNicuStore) DetailValues() map[string][]string { return nil }

// ChecklistKriteriaMasukNicuUpdate ubah checklist kriteria masuk NICU (PUT); kunci lewat query string.
type ChecklistKriteriaMasukNicuUpdate struct {
	ChecklistKriteriaMasukNicuData
}

func (r *ChecklistKriteriaMasukNicuUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaMasukNicuUpdate) Rules(ctx http.Context) map[string]any {
	return checklistKriteriaMasukNicuRules()
}

func (r *ChecklistKriteriaMasukNicuUpdate) Payload() ChecklistKriteriaMasukNicuData {
	return r.ChecklistKriteriaMasukNicuData
}

func (r *ChecklistKriteriaMasukNicuUpdate) DetailValues() map[string][]string { return nil }
