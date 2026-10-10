package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningInstrumenAcrsData isian skrining instrumen ACRS.
type SkriningInstrumenAcrsData struct {
	Tanggal          string `form:"tanggal" json:"tanggal"`
	Nip              string `form:"nip" json:"nip"`
	Pernyataanacrs1  string `form:"pernyataanacrs1" json:"pernyataanacrs1"`
	NilaiAcrs1       *int   `form:"nilai_acrs1" json:"nilai_acrs1"`
	Pernyataanacrs2  string `form:"pernyataanacrs2" json:"pernyataanacrs2"`
	NilaiAcrs2       *int   `form:"nilai_acrs2" json:"nilai_acrs2"`
	Pernyataanacrs3  string `form:"pernyataanacrs3" json:"pernyataanacrs3"`
	NilaiAcrs3       *int   `form:"nilai_acrs3" json:"nilai_acrs3"`
	Pernyataanacrs4  string `form:"pernyataanacrs4" json:"pernyataanacrs4"`
	NilaiAcrs4       *int   `form:"nilai_acrs4" json:"nilai_acrs4"`
	Pernyataanacrs5  string `form:"pernyataanacrs5" json:"pernyataanacrs5"`
	NilaiAcrs5       *int   `form:"nilai_acrs5" json:"nilai_acrs5"`
	Pernyataanacrs6  string `form:"pernyataanacrs6" json:"pernyataanacrs6"`
	NilaiAcrs6       *int   `form:"nilai_acrs6" json:"nilai_acrs6"`
	Pernyataanacrs7  string `form:"pernyataanacrs7" json:"pernyataanacrs7"`
	NilaiAcrs7       *int   `form:"nilai_acrs7" json:"nilai_acrs7"`
	Pernyataanacrs8  string `form:"pernyataanacrs8" json:"pernyataanacrs8"`
	NilaiAcrs8       *int   `form:"nilai_acrs8" json:"nilai_acrs8"`
	Pernyataanacrs9  string `form:"pernyataanacrs9" json:"pernyataanacrs9"`
	NilaiAcrs9       *int   `form:"nilai_acrs9" json:"nilai_acrs9"`
	Pernyataanacrs10 string `form:"pernyataanacrs10" json:"pernyataanacrs10"`
	NilaiAcrs10      *int   `form:"nilai_acrs10" json:"nilai_acrs10"`
	NilaiTotalAcrs   *int   `form:"nilai_total_acrs" json:"nilai_total_acrs"`
	Kesimpulan       string `form:"kesimpulan" json:"kesimpulan"`
}

func skriningInstrumenAcrsRules() map[string]any {
	rules := map[string]any{
		"tanggal":          "required|date",
		"nip":              "required|string|max_len:20",
		"pernyataanacrs1":  "in:Tidak Pernah,Kadang-kadang,Sering,Selalu",
		"nilai_acrs1":      "int",
		"pernyataanacrs2":  "in:Tidak Pernah,Kadang-kadang,Sering,Selalu",
		"nilai_acrs2":      "int",
		"pernyataanacrs3":  "in:Tidak Pernah,Kadang-kadang,Sering,Selalu",
		"nilai_acrs3":      "int",
		"pernyataanacrs4":  "in:Tidak Pernah,Kadang-kadang,Sering,Selalu",
		"nilai_acrs4":      "int",
		"pernyataanacrs5":  "in:Tidak Pernah,Kadang-kadang,Sering,Selalu",
		"nilai_acrs5":      "int",
		"pernyataanacrs6":  "in:Tidak Pernah,Kadang-kadang,Sering,Selalu",
		"nilai_acrs6":      "int",
		"pernyataanacrs7":  "in:Tidak Pernah,Kadang-kadang,Sering,Selalu",
		"nilai_acrs7":      "int",
		"pernyataanacrs8":  "in:Tidak Pernah,Kadang-kadang,Sering,Selalu",
		"nilai_acrs8":      "int",
		"pernyataanacrs9":  "in:Tidak Pernah,Kadang-kadang,Sering,Selalu",
		"nilai_acrs9":      "int",
		"pernyataanacrs10": "in:Tidak Pernah,Kadang-kadang,Sering,Selalu",
		"nilai_acrs10":     "int",
		"nilai_total_acrs": "int",
		"kesimpulan":       "string|max_len:100",
	}
	return rules
}

// SkriningInstrumenAcrsStore simpan skrining instrumen ACRS; kolom waktu kunci kosong = sekarang.
type SkriningInstrumenAcrsStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningInstrumenAcrsData
}

func (r *SkriningInstrumenAcrsStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningInstrumenAcrsStore) Rules(ctx http.Context) map[string]any {
	rules := skriningInstrumenAcrsRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningInstrumenAcrsStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningInstrumenAcrsStore) Payload() SkriningInstrumenAcrsData {
	return r.SkriningInstrumenAcrsData
}

func (r *SkriningInstrumenAcrsStore) DetailValues() map[string][]string { return nil }

// SkriningInstrumenAcrsUpdate ubah skrining instrumen ACRS (PUT); kunci lewat query string.
type SkriningInstrumenAcrsUpdate struct {
	SkriningInstrumenAcrsData
}

func (r *SkriningInstrumenAcrsUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningInstrumenAcrsUpdate) Rules(ctx http.Context) map[string]any {
	return skriningInstrumenAcrsRules()
}

func (r *SkriningInstrumenAcrsUpdate) Payload() SkriningInstrumenAcrsData {
	return r.SkriningInstrumenAcrsData
}

func (r *SkriningInstrumenAcrsUpdate) DetailValues() map[string][]string { return nil }
