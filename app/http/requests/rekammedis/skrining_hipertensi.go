package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningHipertensiData isian skrining hipertensi.
type SkriningHipertensiData struct {
	Tanggal               string `form:"tanggal" json:"tanggal"`
	Anamnesis1            string `form:"anamnesis1" json:"anamnesis1"`
	Anamnesis2            string `form:"anamnesis2" json:"anamnesis2"`
	Anamnesis3            string `form:"anamnesis3" json:"anamnesis3"`
	Anamnesis4            string `form:"anamnesis4" json:"anamnesis4"`
	Anamnesis5            string `form:"anamnesis5" json:"anamnesis5"`
	Anamnesis6            string `form:"anamnesis6" json:"anamnesis6"`
	Anamnesis7            string `form:"anamnesis7" json:"anamnesis7"`
	Anamnesis8            string `form:"anamnesis8" json:"anamnesis8"`
	Sistole               string `form:"sistole" json:"sistole"`
	Diastole              string `form:"diastole" json:"diastole"`
	KlasifikasiHipertensi string `form:"klasifikasi_hipertensi" json:"klasifikasi_hipertensi"`
	HasilSkrining         string `form:"hasil_skrining" json:"hasil_skrining"`
	Keterangan            string `form:"keterangan" json:"keterangan"`
	Nip                   string `form:"nip" json:"nip"`
}

func skriningHipertensiRules() map[string]any {
	rules := map[string]any{
		"tanggal":                "required|date",
		"anamnesis1":             "in:Tidak,Ya",
		"anamnesis2":             "in:Tidak,Ya",
		"anamnesis3":             "in:Tidak,Ya",
		"anamnesis4":             "in:Tidak,Ya",
		"anamnesis5":             "in:Tidak,Ya",
		"anamnesis6":             "in:Tidak,Ya",
		"anamnesis7":             "in:Tidak,Ya",
		"anamnesis8":             "in:Tidak,Ya",
		"sistole":                "string|max_len:3",
		"diastole":               "string|max_len:3",
		"klasifikasi_hipertensi": "required|in:Optimal Normal,Normal,Tinggi,Sub-group : Perbatasan,Tingkat 1 (Hipertensi Ringan),Tingkat 2 (Hipertensi Sedang),Tingkat 3 (Hipertensi Berat),Hipertensi Sistol Tensolasi,Tidak Diketahui",
		"hasil_skrining":         "string|max_len:40",
		"keterangan":             "string|max_len:100",
		"nip":                    "required|string|max_len:20",
	}
	return rules
}

// SkriningHipertensiStore simpan skrining hipertensi; kolom waktu kunci kosong = sekarang.
type SkriningHipertensiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningHipertensiData
}

func (r *SkriningHipertensiStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningHipertensiStore) Rules(ctx http.Context) map[string]any {
	rules := skriningHipertensiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningHipertensiStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningHipertensiStore) Payload() SkriningHipertensiData { return r.SkriningHipertensiData }

func (r *SkriningHipertensiStore) DetailValues() map[string][]string { return nil }

// SkriningHipertensiUpdate ubah skrining hipertensi (PUT); kunci lewat query string.
type SkriningHipertensiUpdate struct {
	SkriningHipertensiData
}

func (r *SkriningHipertensiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningHipertensiUpdate) Rules(ctx http.Context) map[string]any {
	return skriningHipertensiRules()
}

func (r *SkriningHipertensiUpdate) Payload() SkriningHipertensiData { return r.SkriningHipertensiData }

func (r *SkriningHipertensiUpdate) DetailValues() map[string][]string { return nil }
