package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningDiabetesMelitusData isian skrining diabetes melitus.
type SkriningDiabetesMelitusData struct {
	Tanggal            string `form:"tanggal" json:"tanggal"`
	Nip                string `form:"nip" json:"nip"`
	Anamnesis1         string `form:"anamnesis1" json:"anamnesis1"`
	Anamnesis2         string `form:"anamnesis2" json:"anamnesis2"`
	Anamnesis3         string `form:"anamnesis3" json:"anamnesis3"`
	Anamnesis4         string `form:"anamnesis4" json:"anamnesis4"`
	Anamnesis5         string `form:"anamnesis5" json:"anamnesis5"`
	Anamnesis6         string `form:"anamnesis6" json:"anamnesis6"`
	Anamnesis7         string `form:"anamnesis7" json:"anamnesis7"`
	Anamnesis8         string `form:"anamnesis8" json:"anamnesis8"`
	Anamnesis9         string `form:"anamnesis9" json:"anamnesis9"`
	Anamnesis10        string `form:"anamnesis10" json:"anamnesis10"`
	Anamnesis11        string `form:"anamnesis11" json:"anamnesis11"`
	Anamnesis12        string `form:"anamnesis12" json:"anamnesis12"`
	BeratBadan         string `form:"berat_badan" json:"berat_badan"`
	TinggiBadan        string `form:"tinggi_badan" json:"tinggi_badan"`
	Imt                string `form:"imt" json:"imt"`
	KasifikasiImt      string `form:"kasifikasi_imt" json:"kasifikasi_imt"`
	HasilGds           string `form:"hasil_gds" json:"hasil_gds"`
	KeteranganGds      string `form:"keterangan_gds" json:"keterangan_gds"`
	HasilGdp           string `form:"hasil_gdp" json:"hasil_gdp"`
	KeteranganGdp      string `form:"keterangan_gdp" json:"keterangan_gdp"`
	HasilSkrining      string `form:"hasil_skrining" json:"hasil_skrining"`
	KeteranganSkrining string `form:"keterangan_skrining" json:"keterangan_skrining"`
}

func skriningDiabetesMelitusRules() map[string]any {
	rules := map[string]any{
		"tanggal":             "required|date",
		"nip":                 "required|string|max_len:20",
		"anamnesis1":          "in:Ya,Tidak",
		"anamnesis2":          "in:Ya,Tidak",
		"anamnesis3":          "in:Ya,Tidak",
		"anamnesis4":          "in:Ya,Tidak",
		"anamnesis5":          "in:Ya,Tidak",
		"anamnesis6":          "in:Ya,Tidak",
		"anamnesis7":          "in:Ya,Tidak",
		"anamnesis8":          "in:Ya,Tidak",
		"anamnesis9":          "in:Ya,Tidak",
		"anamnesis10":         "in:Ya,Tidak",
		"anamnesis11":         "in:Ya,Tidak",
		"anamnesis12":         "in:Ya,Tidak",
		"berat_badan":         "string|max_len:6",
		"tinggi_badan":        "string|max_len:8",
		"imt":                 "string|max_len:6",
		"kasifikasi_imt":      "in:Berat Badan Kurang,Berat Badan Normal,Kelebihan Berat Badan,Obesitas I,Obesitas II",
		"hasil_gds":           "string|max_len:10",
		"keterangan_gds":      "string|max_len:50",
		"hasil_gdp":           "string|max_len:10",
		"keterangan_gdp":      "string|max_len:50",
		"hasil_skrining":      "in:Normal,Suspek",
		"keterangan_skrining": "string|max_len:60",
	}
	return rules
}

// SkriningDiabetesMelitusStore simpan skrining diabetes melitus; kolom waktu kunci kosong = sekarang.
type SkriningDiabetesMelitusStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningDiabetesMelitusData
}

func (r *SkriningDiabetesMelitusStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningDiabetesMelitusStore) Rules(ctx http.Context) map[string]any {
	rules := skriningDiabetesMelitusRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningDiabetesMelitusStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningDiabetesMelitusStore) Payload() SkriningDiabetesMelitusData {
	return r.SkriningDiabetesMelitusData
}

func (r *SkriningDiabetesMelitusStore) DetailValues() map[string][]string { return nil }

// SkriningDiabetesMelitusUpdate ubah skrining diabetes melitus (PUT); kunci lewat query string.
type SkriningDiabetesMelitusUpdate struct {
	SkriningDiabetesMelitusData
}

func (r *SkriningDiabetesMelitusUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningDiabetesMelitusUpdate) Rules(ctx http.Context) map[string]any {
	return skriningDiabetesMelitusRules()
}

func (r *SkriningDiabetesMelitusUpdate) Payload() SkriningDiabetesMelitusData {
	return r.SkriningDiabetesMelitusData
}

func (r *SkriningDiabetesMelitusUpdate) DetailValues() map[string][]string { return nil }
