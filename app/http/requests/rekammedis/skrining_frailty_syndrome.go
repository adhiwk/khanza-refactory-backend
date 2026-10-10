package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningFrailtySyndromeData isian skrining frailty syndrome.
type SkriningFrailtySyndromeData struct {
	Tanggal                 string `form:"tanggal" json:"tanggal"`
	Resistensi              string `form:"resistensi" json:"resistensi"`
	NilaiResistensi         *int   `form:"nilai_resistensi" json:"nilai_resistensi"`
	Aktivitas               string `form:"aktivitas" json:"aktivitas"`
	NilaiAktivitas          *int   `form:"nilai_aktivitas" json:"nilai_aktivitas"`
	PenyakitTidakPernah     string `form:"penyakit_tidak_pernah" json:"penyakit_tidak_pernah"`
	PenyakitKanker          string `form:"penyakit_kanker" json:"penyakit_kanker"`
	PenyakitGagalJantung    string `form:"penyakit_gagal_jantung" json:"penyakit_gagal_jantung"`
	PenyakitGinjal          string `form:"penyakit_ginjal" json:"penyakit_ginjal"`
	PenyakitNyeriDada       string `form:"penyakit_nyeri_dada" json:"penyakit_nyeri_dada"`
	PenyakitSeranganJantung string `form:"penyakit_serangan_jantung" json:"penyakit_serangan_jantung"`
	PenyakitStroke          string `form:"penyakit_stroke" json:"penyakit_stroke"`
	PenyakitAsma            string `form:"penyakit_asma" json:"penyakit_asma"`
	PenyakitNyeriSendi      string `form:"penyakit_nyeri_sendi" json:"penyakit_nyeri_sendi"`
	PenyakitParuKronis      string `form:"penyakit_paru_kronis" json:"penyakit_paru_kronis"`
	PenyakitHipertensi      string `form:"penyakit_hipertensi" json:"penyakit_hipertensi"`
	PenyakitDiabetes        string `form:"penyakit_diabetes" json:"penyakit_diabetes"`
	NilaiPenyakit           *int   `form:"nilai_penyakit" json:"nilai_penyakit"`
	UsahaBerjalan           string `form:"usaha_berjalan" json:"usaha_berjalan"`
	NilaiUsahaBerjalan      *int   `form:"nilai_usaha_berjalan" json:"nilai_usaha_berjalan"`
	BeratBadan              string `form:"berat_badan" json:"berat_badan"`
	NilaiBeratBadan         *int   `form:"nilai_berat_badan" json:"nilai_berat_badan"`
	NilaiTotal              *int   `form:"nilai_total" json:"nilai_total"`
	HasilSkrining           string `form:"hasil_skrining" json:"hasil_skrining"`
	Keterangan              string `form:"keterangan" json:"keterangan"`
	Nip                     string `form:"nip" json:"nip"`
}

func skriningFrailtySyndromeRules() map[string]any {
	rules := map[string]any{
		"tanggal":                   "required|date",
		"resistensi":                "in:Tidak,Ya",
		"nilai_resistensi":          "int",
		"aktivitas":                 "in:Jarang,Kadang-kadang,Sebagian Besar Waktu,Sepanjang Waktu",
		"nilai_aktivitas":           "int",
		"penyakit_tidak_pernah":     "in:Ya,Tidak",
		"penyakit_kanker":           "in:Ya,Tidak",
		"penyakit_gagal_jantung":    "in:Ya,Tidak",
		"penyakit_ginjal":           "in:Ya,Tidak",
		"penyakit_nyeri_dada":       "in:Ya,Tidak",
		"penyakit_serangan_jantung": "in:Ya,Tidak",
		"penyakit_stroke":           "in:Ya,Tidak",
		"penyakit_asma":             "in:Ya,Tidak",
		"penyakit_nyeri_sendi":      "in:Ya,Tidak",
		"penyakit_paru_kronis":      "in:Ya,Tidak",
		"penyakit_hipertensi":       "in:Ya,Tidak",
		"penyakit_diabetes":         "in:Ya,Tidak",
		"nilai_penyakit":            "int",
		"usaha_berjalan":            "in:Tidak,Ya",
		"nilai_usaha_berjalan":      "int",
		"berat_badan":               "in:< 5%,>= 5%",
		"nilai_berat_badan":         "int",
		"nilai_total":               "int",
		"hasil_skrining":            "in:Rapuh/Renta,Pra-Kerapuhan,Segar & Tidak Rapuh",
		"keterangan":                "string|max_len:40",
		"nip":                       "required|string|max_len:20",
	}
	return rules
}

// SkriningFrailtySyndromeStore simpan skrining frailty syndrome; kolom waktu kunci kosong = sekarang.
type SkriningFrailtySyndromeStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningFrailtySyndromeData
}

func (r *SkriningFrailtySyndromeStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningFrailtySyndromeStore) Rules(ctx http.Context) map[string]any {
	rules := skriningFrailtySyndromeRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningFrailtySyndromeStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningFrailtySyndromeStore) Payload() SkriningFrailtySyndromeData {
	return r.SkriningFrailtySyndromeData
}

func (r *SkriningFrailtySyndromeStore) DetailValues() map[string][]string { return nil }

// SkriningFrailtySyndromeUpdate ubah skrining frailty syndrome (PUT); kunci lewat query string.
type SkriningFrailtySyndromeUpdate struct {
	SkriningFrailtySyndromeData
}

func (r *SkriningFrailtySyndromeUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningFrailtySyndromeUpdate) Rules(ctx http.Context) map[string]any {
	return skriningFrailtySyndromeRules()
}

func (r *SkriningFrailtySyndromeUpdate) Payload() SkriningFrailtySyndromeData {
	return r.SkriningFrailtySyndromeData
}

func (r *SkriningFrailtySyndromeUpdate) DetailValues() map[string][]string { return nil }
