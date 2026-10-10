package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ChecklistPemberianFibrinolitikData isian checklist pemberian fibrinolitik.
type ChecklistPemberianFibrinolitikData struct {
	Tanggal                     string `form:"tanggal" json:"tanggal"`
	Nip                         string `form:"nip" json:"nip"`
	KontraIndikasi1             string `form:"kontra_indikasi1" json:"kontra_indikasi1"`
	KeteranganKontraIndikasi1   string `form:"keterangan_kontra_indikasi1" json:"keterangan_kontra_indikasi1"`
	KontraIndikasi2             string `form:"kontra_indikasi2" json:"kontra_indikasi2"`
	KeteranganKontraIndikasi2   string `form:"keterangan_kontra_indikasi2" json:"keterangan_kontra_indikasi2"`
	KontraIndikasi3             string `form:"kontra_indikasi3" json:"kontra_indikasi3"`
	KeteranganKontraIndikasi3   string `form:"keterangan_kontra_indikasi3" json:"keterangan_kontra_indikasi3"`
	KontraIndikasi4             string `form:"kontra_indikasi4" json:"kontra_indikasi4"`
	KeteranganKontraIndikasi4   string `form:"keterangan_kontra_indikasi4" json:"keterangan_kontra_indikasi4"`
	KontraIndikasi5             string `form:"kontra_indikasi5" json:"kontra_indikasi5"`
	KeteranganKontraIndikasi5   string `form:"keterangan_kontra_indikasi5" json:"keterangan_kontra_indikasi5"`
	KontraIndikasi6             string `form:"kontra_indikasi6" json:"kontra_indikasi6"`
	KeteranganKontraIndikasi6   string `form:"keterangan_kontra_indikasi6" json:"keterangan_kontra_indikasi6"`
	KontraIndikasi7             string `form:"kontra_indikasi7" json:"kontra_indikasi7"`
	KeteranganKontraIndikasi7   string `form:"keterangan_kontra_indikasi7" json:"keterangan_kontra_indikasi7"`
	KontraIndikasi8             string `form:"kontra_indikasi8" json:"kontra_indikasi8"`
	KeteranganKontraIndikasi8   string `form:"keterangan_kontra_indikasi8" json:"keterangan_kontra_indikasi8"`
	KontraIndikasi9             string `form:"kontra_indikasi9" json:"kontra_indikasi9"`
	KeteranganKontraIndikasi9   string `form:"keterangan_kontra_indikasi9" json:"keterangan_kontra_indikasi9"`
	KontraIndikasi10            string `form:"kontra_indikasi10" json:"kontra_indikasi10"`
	KeteranganKontraIndikasi10  string `form:"keterangan_kontra_indikasi10" json:"keterangan_kontra_indikasi10"`
	RisikoTinggi1               string `form:"risiko_tinggi1" json:"risiko_tinggi1"`
	KeteranganRisikoTinggi1     string `form:"keterangan_risiko_tinggi1" json:"keterangan_risiko_tinggi1"`
	RisikoTinggi2               string `form:"risiko_tinggi2" json:"risiko_tinggi2"`
	KeteranganRisikoTinggi2     string `form:"keterangan_risiko_tinggi2" json:"keterangan_risiko_tinggi2"`
	RisikoTinggi3               string `form:"risiko_tinggi3" json:"risiko_tinggi3"`
	KeteranganRisikoTinggi3     string `form:"keterangan_risiko_tinggi3" json:"keterangan_risiko_tinggi3"`
	RisikoTinggi4               string `form:"risiko_tinggi4" json:"risiko_tinggi4"`
	KeteranganRisikoTinggi4     string `form:"keterangan_risiko_tinggi4" json:"keterangan_risiko_tinggi4"`
	RisikoTinggi5               string `form:"risiko_tinggi5" json:"risiko_tinggi5"`
	KeteranganRisikoTinggi5     string `form:"keterangan_risiko_tinggi5" json:"keterangan_risiko_tinggi5"`
	Kesimpulan                  string `form:"kesimpulan" json:"kesimpulan"`
	PersyaratanEkgPreStreptase  string `form:"persyaratan_ekg_pre_streptase" json:"persyaratan_ekg_pre_streptase"`
	PersyaratanEkgPostStreptase string `form:"persyaratan_ekg_post_streptase" json:"persyaratan_ekg_post_streptase"`
	CekTroponin                 string `form:"cek_troponin" json:"cek_troponin"`
}

func checklistPemberianFibrinolitikRules() map[string]any {
	rules := map[string]any{
		"tanggal":                        "required|date",
		"nip":                            "required|string|max_len:20",
		"kontra_indikasi1":               "in:Tidak,Ya",
		"keterangan_kontra_indikasi1":    "string|max_len:30",
		"kontra_indikasi2":               "in:Tidak,Ya",
		"keterangan_kontra_indikasi2":    "string|max_len:30",
		"kontra_indikasi3":               "in:Tidak,Ya",
		"keterangan_kontra_indikasi3":    "string|max_len:30",
		"kontra_indikasi4":               "in:Tidak,Ya",
		"keterangan_kontra_indikasi4":    "string|max_len:30",
		"kontra_indikasi5":               "in:Tidak,Ya",
		"keterangan_kontra_indikasi5":    "string|max_len:30",
		"kontra_indikasi6":               "in:Tidak,Ya",
		"keterangan_kontra_indikasi6":    "string|max_len:30",
		"kontra_indikasi7":               "in:Tidak,Ya",
		"keterangan_kontra_indikasi7":    "string|max_len:30",
		"kontra_indikasi8":               "in:Tidak,Ya",
		"keterangan_kontra_indikasi8":    "string|max_len:30",
		"kontra_indikasi9":               "in:Tidak,Ya",
		"keterangan_kontra_indikasi9":    "string|max_len:30",
		"kontra_indikasi10":              "in:Tidak,Ya",
		"keterangan_kontra_indikasi10":   "string|max_len:30",
		"risiko_tinggi1":                 "in:Tidak,Ya",
		"keterangan_risiko_tinggi1":      "string|max_len:30",
		"risiko_tinggi2":                 "in:Tidak,Ya",
		"keterangan_risiko_tinggi2":      "string|max_len:30",
		"risiko_tinggi3":                 "in:Tidak,Ya",
		"keterangan_risiko_tinggi3":      "string|max_len:30",
		"risiko_tinggi4":                 "in:Tidak,Ya",
		"keterangan_risiko_tinggi4":      "string|max_len:30",
		"risiko_tinggi5":                 "in:Tidak,Ya",
		"keterangan_risiko_tinggi5":      "string|max_len:30",
		"kesimpulan":                     "string|max_len:150",
		"persyaratan_ekg_pre_streptase":  "string|max_len:80",
		"persyaratan_ekg_post_streptase": "string|max_len:80",
		"cek_troponin":                   "string|max_len:80",
	}
	return rules
}

// ChecklistPemberianFibrinolitikStore simpan checklist pemberian fibrinolitik; kolom waktu kunci kosong = sekarang.
type ChecklistPemberianFibrinolitikStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	ChecklistPemberianFibrinolitikData
}

func (r *ChecklistPemberianFibrinolitikStore) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistPemberianFibrinolitikStore) Rules(ctx http.Context) map[string]any {
	rules := checklistPemberianFibrinolitikRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *ChecklistPemberianFibrinolitikStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *ChecklistPemberianFibrinolitikStore) Payload() ChecklistPemberianFibrinolitikData {
	return r.ChecklistPemberianFibrinolitikData
}

func (r *ChecklistPemberianFibrinolitikStore) DetailValues() map[string][]string { return nil }

// ChecklistPemberianFibrinolitikUpdate ubah checklist pemberian fibrinolitik (PUT); kunci lewat query string.
type ChecklistPemberianFibrinolitikUpdate struct {
	ChecklistPemberianFibrinolitikData
}

func (r *ChecklistPemberianFibrinolitikUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistPemberianFibrinolitikUpdate) Rules(ctx http.Context) map[string]any {
	return checklistPemberianFibrinolitikRules()
}

func (r *ChecklistPemberianFibrinolitikUpdate) Payload() ChecklistPemberianFibrinolitikData {
	return r.ChecklistPemberianFibrinolitikData
}

func (r *ChecklistPemberianFibrinolitikUpdate) DetailValues() map[string][]string { return nil }
