package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ChecklistKriteriaMasukPicuData isian checklist kriteria masuk PICU.
type ChecklistKriteriaMasukPicuData struct {
	Kriteriaumum1 string `form:"kriteriaumum1" json:"kriteriaumum1"`
	Kriteriaumum2 string `form:"kriteriaumum2" json:"kriteriaumum2"`
	Kriteriaumum3 string `form:"kriteriaumum3" json:"kriteriaumum3"`
	Respirasi1    string `form:"respirasi1" json:"respirasi1"`
	Respirasi2    string `form:"respirasi2" json:"respirasi2"`
	Respirasi3    string `form:"respirasi3" json:"respirasi3"`
	Respirasi4    string `form:"respirasi4" json:"respirasi4"`
	Kardio1       string `form:"kardio1" json:"kardio1"`
	Kardio2       string `form:"kardio2" json:"kardio2"`
	Kardio3       string `form:"kardio3" json:"kardio3"`
	Kardio4       string `form:"kardio4" json:"kardio4"`
	Neuro1        string `form:"neuro1" json:"neuro1"`
	Neuro2        string `form:"neuro2" json:"neuro2"`
	Neuro3        string `form:"neuro3" json:"neuro3"`
	Neuro4        string `form:"neuro4" json:"neuro4"`
	Bedah1        string `form:"bedah1" json:"bedah1"`
	Bedah2        string `form:"bedah2" json:"bedah2"`
	Bedah3        string `form:"bedah3" json:"bedah3"`
	Kondisilain1  string `form:"kondisilain1" json:"kondisilain1"`
	Kondisilain2  string `form:"kondisilain2" json:"kondisilain2"`
	Kondisilain3  string `form:"kondisilain3" json:"kondisilain3"`
	Keputusan     string `form:"keputusan" json:"keputusan"`
	Keterangan    string `form:"keterangan" json:"keterangan"`
	Nik           string `form:"nik" json:"nik"`
}

func checklistKriteriaMasukPicuRules() map[string]any {
	rules := map[string]any{
		"kriteriaumum1": "required|in:Ya,Tidak",
		"kriteriaumum2": "required|in:Ya,Tidak",
		"kriteriaumum3": "required|in:Ya,Tidak",
		"respirasi1":    "required|in:Ya,Tidak",
		"respirasi2":    "required|in:Ya,Tidak",
		"respirasi3":    "required|in:Ya,Tidak",
		"respirasi4":    "required|in:Ya,Tidak",
		"kardio1":       "required|in:Ya,Tidak",
		"kardio2":       "required|in:Ya,Tidak",
		"kardio3":       "required|in:Ya,Tidak",
		"kardio4":       "required|in:Ya,Tidak",
		"neuro1":        "required|in:Ya,Tidak",
		"neuro2":        "required|in:Ya,Tidak",
		"neuro3":        "required|in:Ya,Tidak",
		"neuro4":        "required|in:Ya,Tidak",
		"bedah1":        "required|in:Ya,Tidak",
		"bedah2":        "required|in:Ya,Tidak",
		"bedah3":        "required|in:Ya,Tidak",
		"kondisilain1":  "required|in:Ya,Tidak",
		"kondisilain2":  "required|in:Ya,Tidak",
		"kondisilain3":  "required|in:Ya,Tidak",
		"keputusan":     "required|in:Diterima Di PICU,Tidak Diterima - Dirawat Di Ruang Lain",
		"keterangan":    "string|max_len:50",
		"nik":           "string|max_len:20",
	}
	return rules
}

// ChecklistKriteriaMasukPicuStore simpan checklist kriteria masuk PICU; kolom waktu kunci kosong = sekarang.
type ChecklistKriteriaMasukPicuStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	ChecklistKriteriaMasukPicuData
}

func (r *ChecklistKriteriaMasukPicuStore) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaMasukPicuStore) Rules(ctx http.Context) map[string]any {
	rules := checklistKriteriaMasukPicuRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *ChecklistKriteriaMasukPicuStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *ChecklistKriteriaMasukPicuStore) Payload() ChecklistKriteriaMasukPicuData {
	return r.ChecklistKriteriaMasukPicuData
}

func (r *ChecklistKriteriaMasukPicuStore) DetailValues() map[string][]string { return nil }

// ChecklistKriteriaMasukPicuUpdate ubah checklist kriteria masuk PICU (PUT); kunci lewat query string.
type ChecklistKriteriaMasukPicuUpdate struct {
	ChecklistKriteriaMasukPicuData
}

func (r *ChecklistKriteriaMasukPicuUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaMasukPicuUpdate) Rules(ctx http.Context) map[string]any {
	return checklistKriteriaMasukPicuRules()
}

func (r *ChecklistKriteriaMasukPicuUpdate) Payload() ChecklistKriteriaMasukPicuData {
	return r.ChecklistKriteriaMasukPicuData
}

func (r *ChecklistKriteriaMasukPicuUpdate) DetailValues() map[string][]string { return nil }
