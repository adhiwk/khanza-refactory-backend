package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningInstrumenSrqData isian skrining SRQ.
type SkriningInstrumenSrqData struct {
	Tanggal         string `form:"tanggal" json:"tanggal"`
	Nip             string `form:"nip" json:"nip"`
	Pernyataansrq1  string `form:"pernyataansrq1" json:"pernyataansrq1"`
	NilaiSrq1       *int   `form:"nilai_srq1" json:"nilai_srq1"`
	Pernyataansrq2  string `form:"pernyataansrq2" json:"pernyataansrq2"`
	NilaiSrq2       *int   `form:"nilai_srq2" json:"nilai_srq2"`
	Pernyataansrq3  string `form:"pernyataansrq3" json:"pernyataansrq3"`
	NilaiSrq3       *int   `form:"nilai_srq3" json:"nilai_srq3"`
	Pernyataansrq4  string `form:"pernyataansrq4" json:"pernyataansrq4"`
	NilaiSrq4       *int   `form:"nilai_srq4" json:"nilai_srq4"`
	Pernyataansrq5  string `form:"pernyataansrq5" json:"pernyataansrq5"`
	NilaiSrq5       *int   `form:"nilai_srq5" json:"nilai_srq5"`
	Pernyataansrq6  string `form:"pernyataansrq6" json:"pernyataansrq6"`
	NilaiSrq6       *int   `form:"nilai_srq6" json:"nilai_srq6"`
	Pernyataansrq7  string `form:"pernyataansrq7" json:"pernyataansrq7"`
	NilaiSrq7       *int   `form:"nilai_srq7" json:"nilai_srq7"`
	Pernyataansrq8  string `form:"pernyataansrq8" json:"pernyataansrq8"`
	NilaiSrq8       *int   `form:"nilai_srq8" json:"nilai_srq8"`
	Pernyataansrq9  string `form:"pernyataansrq9" json:"pernyataansrq9"`
	NilaiSrq9       *int   `form:"nilai_srq9" json:"nilai_srq9"`
	Pernyataansrq10 string `form:"pernyataansrq10" json:"pernyataansrq10"`
	NilaiSrq10      *int   `form:"nilai_srq10" json:"nilai_srq10"`
	Pernyataansrq11 string `form:"pernyataansrq11" json:"pernyataansrq11"`
	NilaiSrq11      *int   `form:"nilai_srq11" json:"nilai_srq11"`
	Pernyataansrq12 string `form:"pernyataansrq12" json:"pernyataansrq12"`
	NilaiSrq12      *int   `form:"nilai_srq12" json:"nilai_srq12"`
	Pernyataansrq13 string `form:"pernyataansrq13" json:"pernyataansrq13"`
	NilaiSrq13      *int   `form:"nilai_srq13" json:"nilai_srq13"`
	Pernyataansrq14 string `form:"pernyataansrq14" json:"pernyataansrq14"`
	NilaiSrq14      *int   `form:"nilai_srq14" json:"nilai_srq14"`
	Pernyataansrq15 string `form:"pernyataansrq15" json:"pernyataansrq15"`
	NilaiSrq15      *int   `form:"nilai_srq15" json:"nilai_srq15"`
	Pernyataansrq16 string `form:"pernyataansrq16" json:"pernyataansrq16"`
	NilaiSrq16      *int   `form:"nilai_srq16" json:"nilai_srq16"`
	Pernyataansrq17 string `form:"pernyataansrq17" json:"pernyataansrq17"`
	NilaiSrq17      *int   `form:"nilai_srq17" json:"nilai_srq17"`
	Pernyataansrq18 string `form:"pernyataansrq18" json:"pernyataansrq18"`
	NilaiSrq18      *int   `form:"nilai_srq18" json:"nilai_srq18"`
	Pernyataansrq19 string `form:"pernyataansrq19" json:"pernyataansrq19"`
	NilaiSrq19      *int   `form:"nilai_srq19" json:"nilai_srq19"`
	Pernyataansrq20 string `form:"pernyataansrq20" json:"pernyataansrq20"`
	NilaiSrq20      *int   `form:"nilai_srq20" json:"nilai_srq20"`
	NilaiTotalSrq   *int   `form:"nilai_total_srq" json:"nilai_total_srq"`
	Kesimpulan      string `form:"kesimpulan" json:"kesimpulan"`
}

func skriningInstrumenSrqRules() map[string]any {
	rules := map[string]any{
		"tanggal":         "required|date",
		"nip":             "required|string|max_len:20",
		"pernyataansrq1":  "in:Tidak,Ya",
		"nilai_srq1":      "int",
		"pernyataansrq2":  "in:Tidak,Ya",
		"nilai_srq2":      "int",
		"pernyataansrq3":  "in:Tidak,Ya",
		"nilai_srq3":      "int",
		"pernyataansrq4":  "in:Tidak,Ya",
		"nilai_srq4":      "int",
		"pernyataansrq5":  "in:Tidak,Ya",
		"nilai_srq5":      "int",
		"pernyataansrq6":  "in:Tidak,Ya",
		"nilai_srq6":      "int",
		"pernyataansrq7":  "in:Tidak,Ya",
		"nilai_srq7":      "int",
		"pernyataansrq8":  "in:Tidak,Ya",
		"nilai_srq8":      "int",
		"pernyataansrq9":  "in:Tidak,Ya",
		"nilai_srq9":      "int",
		"pernyataansrq10": "in:Tidak,Ya",
		"nilai_srq10":     "int",
		"pernyataansrq11": "in:Tidak,Ya",
		"nilai_srq11":     "int",
		"pernyataansrq12": "in:Tidak,Ya",
		"nilai_srq12":     "int",
		"pernyataansrq13": "in:Tidak,Ya",
		"nilai_srq13":     "int",
		"pernyataansrq14": "in:Tidak,Ya",
		"nilai_srq14":     "int",
		"pernyataansrq15": "in:Tidak,Ya",
		"nilai_srq15":     "int",
		"pernyataansrq16": "in:Tidak,Ya",
		"nilai_srq16":     "int",
		"pernyataansrq17": "in:Tidak,Ya",
		"nilai_srq17":     "int",
		"pernyataansrq18": "in:Tidak,Ya",
		"nilai_srq18":     "int",
		"pernyataansrq19": "in:Tidak,Ya",
		"nilai_srq19":     "int",
		"pernyataansrq20": "in:Tidak,Ya",
		"nilai_srq20":     "int",
		"nilai_total_srq": "int",
		"kesimpulan":      "string|max_len:100",
	}
	return rules
}

// SkriningInstrumenSrqStore simpan skrining SRQ; kolom waktu kunci kosong = sekarang.
type SkriningInstrumenSrqStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningInstrumenSrqData
}

func (r *SkriningInstrumenSrqStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningInstrumenSrqStore) Rules(ctx http.Context) map[string]any {
	rules := skriningInstrumenSrqRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningInstrumenSrqStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningInstrumenSrqStore) Payload() SkriningInstrumenSrqData {
	return r.SkriningInstrumenSrqData
}

func (r *SkriningInstrumenSrqStore) DetailValues() map[string][]string { return nil }

// SkriningInstrumenSrqUpdate ubah skrining SRQ (PUT); kunci lewat query string.
type SkriningInstrumenSrqUpdate struct {
	SkriningInstrumenSrqData
}

func (r *SkriningInstrumenSrqUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningInstrumenSrqUpdate) Rules(ctx http.Context) map[string]any {
	return skriningInstrumenSrqRules()
}

func (r *SkriningInstrumenSrqUpdate) Payload() SkriningInstrumenSrqData {
	return r.SkriningInstrumenSrqData
}

func (r *SkriningInstrumenSrqUpdate) DetailValues() map[string][]string { return nil }
