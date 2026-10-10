package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ChecklistKriteriaKeluarNicuData isian checklist kriteria keluar NICU.
type ChecklistKriteriaKeluarNicuData struct {
	Respirasi1 string `form:"respirasi1" json:"respirasi1"`
	Respirasi2 string `form:"respirasi2" json:"respirasi2"`
	Respirasi3 string `form:"respirasi3" json:"respirasi3"`
	Kardio1    string `form:"kardio1" json:"kardio1"`
	Kardio2    string `form:"kardio2" json:"kardio2"`
	Nutrisi1   string `form:"nutrisi1" json:"nutrisi1"`
	Nutrisi2   string `form:"nutrisi2" json:"nutrisi2"`
	Nutrisi3   string `form:"nutrisi3" json:"nutrisi3"`
	Suhutubuh1 string `form:"suhutubuh1" json:"suhutubuh1"`
	Suhutubuh2 string `form:"suhutubuh2" json:"suhutubuh2"`
	Infeksi1   string `form:"infeksi1" json:"infeksi1"`
	Infeksi2   string `form:"infeksi2" json:"infeksi2"`
	Infeksi3   string `form:"infeksi3" json:"infeksi3"`
	Keputusan  string `form:"keputusan" json:"keputusan"`
	Keterangan string `form:"keterangan" json:"keterangan"`
	Nik        string `form:"nik" json:"nik"`
}

func checklistKriteriaKeluarNicuRules() map[string]any {
	rules := map[string]any{
		"respirasi1": "required|in:Ya,Tidak",
		"respirasi2": "required|in:Ya,Tidak",
		"respirasi3": "required|in:Ya,Tidak",
		"kardio1":    "required|in:Ya,Tidak",
		"kardio2":    "required|in:Ya,Tidak",
		"nutrisi1":   "required|in:Ya,Tidak",
		"nutrisi2":   "required|in:Ya,Tidak",
		"nutrisi3":   "required|in:Ya,Tidak",
		"suhutubuh1": "required|in:Ya,Tidak",
		"suhutubuh2": "required|in:Ya,Tidak",
		"infeksi1":   "required|in:Ya,Tidak",
		"infeksi2":   "required|in:Ya,Tidak",
		"infeksi3":   "required|in:Ya,Tidak",
		"keputusan":  "required|in:Layak Dipindahkan Ke Ruang Rawat Bayi/Rawat Gabung,Layak Pulang,Tetap Di NICU",
		"keterangan": "string|max_len:50",
		"nik":        "string|max_len:20",
	}
	return rules
}

// ChecklistKriteriaKeluarNicuStore simpan checklist kriteria keluar NICU; kolom waktu kunci kosong = sekarang.
type ChecklistKriteriaKeluarNicuStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	ChecklistKriteriaKeluarNicuData
}

func (r *ChecklistKriteriaKeluarNicuStore) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaKeluarNicuStore) Rules(ctx http.Context) map[string]any {
	rules := checklistKriteriaKeluarNicuRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *ChecklistKriteriaKeluarNicuStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *ChecklistKriteriaKeluarNicuStore) Payload() ChecklistKriteriaKeluarNicuData {
	return r.ChecklistKriteriaKeluarNicuData
}

func (r *ChecklistKriteriaKeluarNicuStore) DetailValues() map[string][]string { return nil }

// ChecklistKriteriaKeluarNicuUpdate ubah checklist kriteria keluar NICU (PUT); kunci lewat query string.
type ChecklistKriteriaKeluarNicuUpdate struct {
	ChecklistKriteriaKeluarNicuData
}

func (r *ChecklistKriteriaKeluarNicuUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaKeluarNicuUpdate) Rules(ctx http.Context) map[string]any {
	return checklistKriteriaKeluarNicuRules()
}

func (r *ChecklistKriteriaKeluarNicuUpdate) Payload() ChecklistKriteriaKeluarNicuData {
	return r.ChecklistKriteriaKeluarNicuData
}

func (r *ChecklistKriteriaKeluarNicuUpdate) DetailValues() map[string][]string { return nil }
