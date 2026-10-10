package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ChecklistKriteriaKeluarPicuData isian checklist kriteria keluar PICU.
type ChecklistKriteriaKeluarPicuData struct {
	Kondisiklinis1      string `form:"kondisiklinis1" json:"kondisiklinis1"`
	Kondisiklinis2      string `form:"kondisiklinis2" json:"kondisiklinis2"`
	Kondisiklinis3      string `form:"kondisiklinis3" json:"kondisiklinis3"`
	Kondisiklinis4      string `form:"kondisiklinis4" json:"kondisiklinis4"`
	Kondisiklinis5      string `form:"kondisiklinis5" json:"kondisiklinis5"`
	Kondisiklinis6      string `form:"kondisiklinis6" json:"kondisiklinis6"`
	Kebutuhanperawatan1 string `form:"kebutuhanperawatan1" json:"kebutuhanperawatan1"`
	Kebutuhanperawatan2 string `form:"kebutuhanperawatan2" json:"kebutuhanperawatan2"`
	Kebutuhanperawatan3 string `form:"kebutuhanperawatan3" json:"kebutuhanperawatan3"`
	Kebutuhanperawatan4 string `form:"kebutuhanperawatan4" json:"kebutuhanperawatan4"`
	Tindaklanjut1       string `form:"tindaklanjut1" json:"tindaklanjut1"`
	Tindaklanjut2       string `form:"tindaklanjut2" json:"tindaklanjut2"`
	Tindaklanjut3       string `form:"tindaklanjut3" json:"tindaklanjut3"`
	Tindaklanjut4       string `form:"tindaklanjut4" json:"tindaklanjut4"`
	Keputusan           string `form:"keputusan" json:"keputusan"`
	Keterangan          string `form:"keterangan" json:"keterangan"`
	Nik                 string `form:"nik" json:"nik"`
}

func checklistKriteriaKeluarPicuRules() map[string]any {
	rules := map[string]any{
		"kondisiklinis1":      "required|in:Ya,Tidak",
		"kondisiklinis2":      "required|in:Ya,Tidak",
		"kondisiklinis3":      "required|in:Ya,Tidak",
		"kondisiklinis4":      "required|in:Ya,Tidak",
		"kondisiklinis5":      "required|in:Ya,Tidak",
		"kondisiklinis6":      "required|in:Ya,Tidak",
		"kebutuhanperawatan1": "required|in:Ya,Tidak",
		"kebutuhanperawatan2": "required|in:Ya,Tidak",
		"kebutuhanperawatan3": "required|in:Ya,Tidak",
		"kebutuhanperawatan4": "required|in:Ya,Tidak",
		"tindaklanjut1":       "required|in:Ya,Tidak",
		"tindaklanjut2":       "required|in:Ya,Tidak",
		"tindaklanjut3":       "required|in:Ya,Tidak",
		"tindaklanjut4":       "required|in:Ya,Tidak",
		"keputusan":           "required|in:Layak Keluar Dari PICU/Pindah Ke Ruang Rawat Biasa,Tidak Layak Keluar/Tetap Dirawat Di PICU",
		"keterangan":          "string|max_len:50",
		"nik":                 "string|max_len:20",
	}
	return rules
}

// ChecklistKriteriaKeluarPicuStore simpan checklist kriteria keluar PICU; kolom waktu kunci kosong = sekarang.
type ChecklistKriteriaKeluarPicuStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	ChecklistKriteriaKeluarPicuData
}

func (r *ChecklistKriteriaKeluarPicuStore) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaKeluarPicuStore) Rules(ctx http.Context) map[string]any {
	rules := checklistKriteriaKeluarPicuRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *ChecklistKriteriaKeluarPicuStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *ChecklistKriteriaKeluarPicuStore) Payload() ChecklistKriteriaKeluarPicuData {
	return r.ChecklistKriteriaKeluarPicuData
}

func (r *ChecklistKriteriaKeluarPicuStore) DetailValues() map[string][]string { return nil }

// ChecklistKriteriaKeluarPicuUpdate ubah checklist kriteria keluar PICU (PUT); kunci lewat query string.
type ChecklistKriteriaKeluarPicuUpdate struct {
	ChecklistKriteriaKeluarPicuData
}

func (r *ChecklistKriteriaKeluarPicuUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaKeluarPicuUpdate) Rules(ctx http.Context) map[string]any {
	return checklistKriteriaKeluarPicuRules()
}

func (r *ChecklistKriteriaKeluarPicuUpdate) Payload() ChecklistKriteriaKeluarPicuData {
	return r.ChecklistKriteriaKeluarPicuData
}

func (r *ChecklistKriteriaKeluarPicuUpdate) DetailValues() map[string][]string { return nil }
