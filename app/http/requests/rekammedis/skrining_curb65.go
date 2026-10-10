package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningCurb65Data isian skrining 65.
type SkriningCurb65Data struct {
	Tanggal           string `form:"tanggal" json:"tanggal"`
	Nip               string `form:"nip" json:"nip"`
	Pernyataancurb651 string `form:"pernyataancurb651" json:"pernyataancurb651"`
	NilaiCurb651      *int   `form:"nilai_curb651" json:"nilai_curb651"`
	Pernyataancurb652 string `form:"pernyataancurb652" json:"pernyataancurb652"`
	NilaiCurb652      *int   `form:"nilai_curb652" json:"nilai_curb652"`
	Pernyataancurb653 string `form:"pernyataancurb653" json:"pernyataancurb653"`
	NilaiCurb653      *int   `form:"nilai_curb653" json:"nilai_curb653"`
	Pernyataancurb654 string `form:"pernyataancurb654" json:"pernyataancurb654"`
	NilaiCurb654      *int   `form:"nilai_curb654" json:"nilai_curb654"`
	Pernyataancurb655 string `form:"pernyataancurb655" json:"pernyataancurb655"`
	NilaiCurb655      *int   `form:"nilai_curb655" json:"nilai_curb655"`
	NilaiTotalCurb65  *int   `form:"nilai_total_curb65" json:"nilai_total_curb65"`
	Kesimpulan        string `form:"kesimpulan" json:"kesimpulan"`
}

func skriningCurb65Rules() map[string]any {
	rules := map[string]any{
		"tanggal":            "required|date",
		"nip":                "required|string|max_len:20",
		"pernyataancurb651":  "in:Ya,Tidak",
		"nilai_curb651":      "int",
		"pernyataancurb652":  "in:Ya,Tidak",
		"nilai_curb652":      "int",
		"pernyataancurb653":  "in:Ya,Tidak",
		"nilai_curb653":      "int",
		"pernyataancurb654":  "in:Ya,Tidak",
		"nilai_curb654":      "int",
		"pernyataancurb655":  "in:Ya,Tidak",
		"nilai_curb655":      "int",
		"nilai_total_curb65": "int",
		"kesimpulan":         "string|max_len:100",
	}
	return rules
}

// SkriningCurb65Store simpan skrining 65; kolom waktu kunci kosong = sekarang.
type SkriningCurb65Store struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningCurb65Data
}

func (r *SkriningCurb65Store) Authorize(ctx http.Context) error { return nil }

func (r *SkriningCurb65Store) Rules(ctx http.Context) map[string]any {
	rules := skriningCurb65Rules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningCurb65Store) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningCurb65Store) Payload() SkriningCurb65Data { return r.SkriningCurb65Data }

func (r *SkriningCurb65Store) DetailValues() map[string][]string { return nil }

// SkriningCurb65Update ubah skrining 65 (PUT); kunci lewat query string.
type SkriningCurb65Update struct {
	SkriningCurb65Data
}

func (r *SkriningCurb65Update) Authorize(ctx http.Context) error { return nil }

func (r *SkriningCurb65Update) Rules(ctx http.Context) map[string]any { return skriningCurb65Rules() }

func (r *SkriningCurb65Update) Payload() SkriningCurb65Data { return r.SkriningCurb65Data }

func (r *SkriningCurb65Update) DetailValues() map[string][]string { return nil }
