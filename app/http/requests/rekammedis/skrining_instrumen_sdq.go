package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningInstrumenSdqData isian skrining instrumen SDQ.
type SkriningInstrumenSdqData struct {
	Tanggal              string `form:"tanggal" json:"tanggal"`
	Nip                  string `form:"nip" json:"nip"`
	Pernyataansdq1       string `form:"pernyataansdq1" json:"pernyataansdq1"`
	NilaiSdq1            *int   `form:"nilai_sdq1" json:"nilai_sdq1"`
	Pernyataansdq2       string `form:"pernyataansdq2" json:"pernyataansdq2"`
	NilaiSdq2            *int   `form:"nilai_sdq2" json:"nilai_sdq2"`
	Pernyataansdq3       string `form:"pernyataansdq3" json:"pernyataansdq3"`
	NilaiSdq3            *int   `form:"nilai_sdq3" json:"nilai_sdq3"`
	Pernyataansdq4       string `form:"pernyataansdq4" json:"pernyataansdq4"`
	NilaiSdq4            *int   `form:"nilai_sdq4" json:"nilai_sdq4"`
	Pernyataansdq5       string `form:"pernyataansdq5" json:"pernyataansdq5"`
	NilaiSdq5            *int   `form:"nilai_sdq5" json:"nilai_sdq5"`
	Pernyataansdq6       string `form:"pernyataansdq6" json:"pernyataansdq6"`
	NilaiSdq6            *int   `form:"nilai_sdq6" json:"nilai_sdq6"`
	Pernyataansdq7       string `form:"pernyataansdq7" json:"pernyataansdq7"`
	NilaiSdq7            *int   `form:"nilai_sdq7" json:"nilai_sdq7"`
	Pernyataansdq8       string `form:"pernyataansdq8" json:"pernyataansdq8"`
	NilaiSdq8            *int   `form:"nilai_sdq8" json:"nilai_sdq8"`
	Pernyataansdq9       string `form:"pernyataansdq9" json:"pernyataansdq9"`
	NilaiSdq9            *int   `form:"nilai_sdq9" json:"nilai_sdq9"`
	Pernyataansdq10      string `form:"pernyataansdq10" json:"pernyataansdq10"`
	NilaiSdq10           *int   `form:"nilai_sdq10" json:"nilai_sdq10"`
	Pernyataansdq11      string `form:"pernyataansdq11" json:"pernyataansdq11"`
	NilaiSdq11           *int   `form:"nilai_sdq11" json:"nilai_sdq11"`
	Pernyataansdq12      string `form:"pernyataansdq12" json:"pernyataansdq12"`
	NilaiSdq12           *int   `form:"nilai_sdq12" json:"nilai_sdq12"`
	Pernyataansdq13      string `form:"pernyataansdq13" json:"pernyataansdq13"`
	NilaiSdq13           *int   `form:"nilai_sdq13" json:"nilai_sdq13"`
	Pernyataansdq14      string `form:"pernyataansdq14" json:"pernyataansdq14"`
	NilaiSdq14           *int   `form:"nilai_sdq14" json:"nilai_sdq14"`
	Pernyataansdq15      string `form:"pernyataansdq15" json:"pernyataansdq15"`
	NilaiSdq15           *int   `form:"nilai_sdq15" json:"nilai_sdq15"`
	Pernyataansdq16      string `form:"pernyataansdq16" json:"pernyataansdq16"`
	NilaiSdq16           *int   `form:"nilai_sdq16" json:"nilai_sdq16"`
	Pernyataansdq17      string `form:"pernyataansdq17" json:"pernyataansdq17"`
	NilaiSdq17           *int   `form:"nilai_sdq17" json:"nilai_sdq17"`
	Pernyataansdq18      string `form:"pernyataansdq18" json:"pernyataansdq18"`
	NilaiSdq18           *int   `form:"nilai_sdq18" json:"nilai_sdq18"`
	Pernyataansdq19      string `form:"pernyataansdq19" json:"pernyataansdq19"`
	NilaiSdq19           *int   `form:"nilai_sdq19" json:"nilai_sdq19"`
	Pernyataansdq20      string `form:"pernyataansdq20" json:"pernyataansdq20"`
	NilaiSdq20           *int   `form:"nilai_sdq20" json:"nilai_sdq20"`
	Pernyataansdq21      string `form:"pernyataansdq21" json:"pernyataansdq21"`
	NilaiSdq21           *int   `form:"nilai_sdq21" json:"nilai_sdq21"`
	Pernyataansdq22      string `form:"pernyataansdq22" json:"pernyataansdq22"`
	NilaiSdq22           *int   `form:"nilai_sdq22" json:"nilai_sdq22"`
	Pernyataansdq23      string `form:"pernyataansdq23" json:"pernyataansdq23"`
	NilaiSdq23           *int   `form:"nilai_sdq23" json:"nilai_sdq23"`
	Pernyataansdq24      string `form:"pernyataansdq24" json:"pernyataansdq24"`
	NilaiSdq24           *int   `form:"nilai_sdq24" json:"nilai_sdq24"`
	Pernyataansdq25      string `form:"pernyataansdq25" json:"pernyataansdq25"`
	NilaiSdq25           *int   `form:"nilai_sdq25" json:"nilai_sdq25"`
	NilaiTotalSdq        *int   `form:"nilai_total_sdq" json:"nilai_total_sdq"`
	GejalaEmosional      string `form:"gejala_emosional" json:"gejala_emosional"`
	NilaiGejalaEmosional *int   `form:"nilai_gejala_emosional" json:"nilai_gejala_emosional"`
	MasalahPerilaku      string `form:"masalah_perilaku" json:"masalah_perilaku"`
	NilaiMasalahPerilaku *int   `form:"nilai_masalah_perilaku" json:"nilai_masalah_perilaku"`
	Hiperaktivitas       string `form:"hiperaktivitas" json:"hiperaktivitas"`
	NilaiHiperaktivitas  *int   `form:"nilai_hiperaktivitas" json:"nilai_hiperaktivitas"`
	TemanSebaya          string `form:"teman_sebaya" json:"teman_sebaya"`
	NilaiTemanSebaya     *int   `form:"nilai_teman_sebaya" json:"nilai_teman_sebaya"`
	Kekuatan             string `form:"kekuatan" json:"kekuatan"`
	NilaiKekuatan        *int   `form:"nilai_kekuatan" json:"nilai_kekuatan"`
	Kesulitan            string `form:"kesulitan" json:"kesulitan"`
	NilaiKesulitan       *int   `form:"nilai_kesulitan" json:"nilai_kesulitan"`
	Keterangan           string `form:"keterangan" json:"keterangan"`
}

func skriningInstrumenSdqRules() map[string]any {
	rules := map[string]any{
		"tanggal":                "required|date",
		"nip":                    "required|string|max_len:20",
		"pernyataansdq1":         "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq1":             "int",
		"pernyataansdq2":         "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq2":             "int",
		"pernyataansdq3":         "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq3":             "int",
		"pernyataansdq4":         "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq4":             "int",
		"pernyataansdq5":         "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq5":             "int",
		"pernyataansdq6":         "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq6":             "int",
		"pernyataansdq7":         "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq7":             "int",
		"pernyataansdq8":         "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq8":             "int",
		"pernyataansdq9":         "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq9":             "int",
		"pernyataansdq10":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq10":            "int",
		"pernyataansdq11":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq11":            "int",
		"pernyataansdq12":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq12":            "int",
		"pernyataansdq13":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq13":            "int",
		"pernyataansdq14":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq14":            "int",
		"pernyataansdq15":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq15":            "int",
		"pernyataansdq16":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq16":            "int",
		"pernyataansdq17":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq17":            "int",
		"pernyataansdq18":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq18":            "int",
		"pernyataansdq19":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq19":            "int",
		"pernyataansdq20":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq20":            "int",
		"pernyataansdq21":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq21":            "int",
		"pernyataansdq22":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq22":            "int",
		"pernyataansdq23":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq23":            "int",
		"pernyataansdq24":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq24":            "int",
		"pernyataansdq25":        "in:Tidak Benar,Agak Benar,Selalu Benar",
		"nilai_sdq25":            "int",
		"nilai_total_sdq":        "int",
		"gejala_emosional":       "in:Abnormal,Ambang/Boderline,Normal",
		"nilai_gejala_emosional": "int",
		"masalah_perilaku":       "in:Abnormal,Ambang/Boderline,Normal",
		"nilai_masalah_perilaku": "int",
		"hiperaktivitas":         "in:Abnormal,Ambang/Boderline,Normal",
		"nilai_hiperaktivitas":   "int",
		"teman_sebaya":           "in:Abnormal,Ambang/Boderline,Normal",
		"nilai_teman_sebaya":     "int",
		"kekuatan":               "in:Abnormal,Ambang/Boderline,Normal",
		"nilai_kekuatan":         "int",
		"kesulitan":              "in:Abnormal,Ambang/Boderline,Normal",
		"nilai_kesulitan":        "int",
		"keterangan":             "string|max_len:200",
	}
	return rules
}

// SkriningInstrumenSdqStore simpan skrining instrumen SDQ; kolom waktu kunci kosong = sekarang.
type SkriningInstrumenSdqStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningInstrumenSdqData
}

func (r *SkriningInstrumenSdqStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningInstrumenSdqStore) Rules(ctx http.Context) map[string]any {
	rules := skriningInstrumenSdqRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningInstrumenSdqStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningInstrumenSdqStore) Payload() SkriningInstrumenSdqData {
	return r.SkriningInstrumenSdqData
}

func (r *SkriningInstrumenSdqStore) DetailValues() map[string][]string { return nil }

// SkriningInstrumenSdqUpdate ubah skrining instrumen SDQ (PUT); kunci lewat query string.
type SkriningInstrumenSdqUpdate struct {
	SkriningInstrumenSdqData
}

func (r *SkriningInstrumenSdqUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningInstrumenSdqUpdate) Rules(ctx http.Context) map[string]any {
	return skriningInstrumenSdqRules()
}

func (r *SkriningInstrumenSdqUpdate) Payload() SkriningInstrumenSdqData {
	return r.SkriningInstrumenSdqData
}

func (r *SkriningInstrumenSdqUpdate) DetailValues() map[string][]string { return nil }
