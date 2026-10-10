package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningInstrumenAmtData isian skrining instrumen AMT.
type SkriningInstrumenAmtData struct {
	Tanggal         string `form:"tanggal" json:"tanggal"`
	Nip             string `form:"nip" json:"nip"`
	Pernyataanamt1  string `form:"pernyataanamt1" json:"pernyataanamt1"`
	NilaiAmt1       *int   `form:"nilai_amt1" json:"nilai_amt1"`
	Pernyataanamt2  string `form:"pernyataanamt2" json:"pernyataanamt2"`
	NilaiAmt2       *int   `form:"nilai_amt2" json:"nilai_amt2"`
	Pernyataanamt3  string `form:"pernyataanamt3" json:"pernyataanamt3"`
	NilaiAmt3       *int   `form:"nilai_amt3" json:"nilai_amt3"`
	Pernyataanamt4  string `form:"pernyataanamt4" json:"pernyataanamt4"`
	NilaiAmt4       *int   `form:"nilai_amt4" json:"nilai_amt4"`
	Pernyataanamt5  string `form:"pernyataanamt5" json:"pernyataanamt5"`
	NilaiAmt5       *int   `form:"nilai_amt5" json:"nilai_amt5"`
	Pernyataanamt6  string `form:"pernyataanamt6" json:"pernyataanamt6"`
	NilaiAmt6       *int   `form:"nilai_amt6" json:"nilai_amt6"`
	Pernyataanamt7  string `form:"pernyataanamt7" json:"pernyataanamt7"`
	NilaiAmt7       *int   `form:"nilai_amt7" json:"nilai_amt7"`
	Pernyataanamt8  string `form:"pernyataanamt8" json:"pernyataanamt8"`
	NilaiAmt8       *int   `form:"nilai_amt8" json:"nilai_amt8"`
	Pernyataanamt9  string `form:"pernyataanamt9" json:"pernyataanamt9"`
	NilaiAmt9       *int   `form:"nilai_amt9" json:"nilai_amt9"`
	Pernyataanamt10 string `form:"pernyataanamt10" json:"pernyataanamt10"`
	NilaiAmt10      *int   `form:"nilai_amt10" json:"nilai_amt10"`
	NilaiTotalAmt   *int   `form:"nilai_total_amt" json:"nilai_total_amt"`
	Kesimpulan      string `form:"kesimpulan" json:"kesimpulan"`
}

func skriningInstrumenAmtRules() map[string]any {
	rules := map[string]any{
		"tanggal":         "required|date",
		"nip":             "required|string|max_len:20",
		"pernyataanamt1":  "in:Salah,Benar",
		"nilai_amt1":      "int",
		"pernyataanamt2":  "in:Salah,Benar",
		"nilai_amt2":      "int",
		"pernyataanamt3":  "in:Salah,Benar",
		"nilai_amt3":      "int",
		"pernyataanamt4":  "in:Salah,Benar",
		"nilai_amt4":      "int",
		"pernyataanamt5":  "in:Salah,Benar",
		"nilai_amt5":      "int",
		"pernyataanamt6":  "in:Salah,Benar",
		"nilai_amt6":      "int",
		"pernyataanamt7":  "in:Salah,Benar",
		"nilai_amt7":      "int",
		"pernyataanamt8":  "in:Salah,Benar",
		"nilai_amt8":      "int",
		"pernyataanamt9":  "in:Salah,Benar",
		"nilai_amt9":      "int",
		"pernyataanamt10": "in:Salah,Benar",
		"nilai_amt10":     "int",
		"nilai_total_amt": "int",
		"kesimpulan":      "string|max_len:100",
	}
	return rules
}

// SkriningInstrumenAmtStore simpan skrining instrumen AMT; kolom waktu kunci kosong = sekarang.
type SkriningInstrumenAmtStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningInstrumenAmtData
}

func (r *SkriningInstrumenAmtStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningInstrumenAmtStore) Rules(ctx http.Context) map[string]any {
	rules := skriningInstrumenAmtRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningInstrumenAmtStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningInstrumenAmtStore) Payload() SkriningInstrumenAmtData {
	return r.SkriningInstrumenAmtData
}

func (r *SkriningInstrumenAmtStore) DetailValues() map[string][]string { return nil }

// SkriningInstrumenAmtUpdate ubah skrining instrumen AMT (PUT); kunci lewat query string.
type SkriningInstrumenAmtUpdate struct {
	SkriningInstrumenAmtData
}

func (r *SkriningInstrumenAmtUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningInstrumenAmtUpdate) Rules(ctx http.Context) map[string]any {
	return skriningInstrumenAmtRules()
}

func (r *SkriningInstrumenAmtUpdate) Payload() SkriningInstrumenAmtData {
	return r.SkriningInstrumenAmtData
}

func (r *SkriningInstrumenAmtUpdate) DetailValues() map[string][]string { return nil }
