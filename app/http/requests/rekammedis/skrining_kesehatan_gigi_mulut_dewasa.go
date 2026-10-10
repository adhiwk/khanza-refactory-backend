package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningKesehatanGigiMulutDewasaData isian skrining kesehatan gigi mulut dewasa.
type SkriningKesehatanGigiMulutDewasaData struct {
	Tanggal          string `form:"tanggal" json:"tanggal"`
	KontrolGigi      string `form:"kontrol_gigi" json:"kontrol_gigi"`
	GigiBungsuTumbuh string `form:"gigi_bungsu_tumbuh" json:"gigi_bungsu_tumbuh"`
	GigiHilang       string `form:"gigi_hilang" json:"gigi_hilang"`
	GigiBerlubang    string `form:"gigi_berlubang" json:"gigi_berlubang"`
	HasilSkrining    string `form:"hasil_skrining" json:"hasil_skrining"`
	Keterangan       string `form:"keterangan" json:"keterangan"`
	Nip              string `form:"nip" json:"nip"`
}

func skriningKesehatanGigiMulutDewasaRules() map[string]any {
	rules := map[string]any{
		"tanggal":            "required|date",
		"kontrol_gigi":       "in:Ya,Tidak",
		"gigi_bungsu_tumbuh": "in:Ya,Tidak",
		"gigi_hilang":        "in:Ya,Tidak",
		"gigi_berlubang":     "in:Ya,Tidak",
		"hasil_skrining":     "string|max_len:50",
		"keterangan":         "string|max_len:100",
		"nip":                "required|string|max_len:20",
	}
	return rules
}

// SkriningKesehatanGigiMulutDewasaStore simpan skrining kesehatan gigi mulut dewasa; kolom waktu kunci kosong = sekarang.
type SkriningKesehatanGigiMulutDewasaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningKesehatanGigiMulutDewasaData
}

func (r *SkriningKesehatanGigiMulutDewasaStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningKesehatanGigiMulutDewasaStore) Rules(ctx http.Context) map[string]any {
	rules := skriningKesehatanGigiMulutDewasaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningKesehatanGigiMulutDewasaStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *SkriningKesehatanGigiMulutDewasaStore) Payload() SkriningKesehatanGigiMulutDewasaData {
	return r.SkriningKesehatanGigiMulutDewasaData
}

func (r *SkriningKesehatanGigiMulutDewasaStore) DetailValues() map[string][]string { return nil }

// SkriningKesehatanGigiMulutDewasaUpdate ubah skrining kesehatan gigi mulut dewasa (PUT); kunci lewat query string.
type SkriningKesehatanGigiMulutDewasaUpdate struct {
	SkriningKesehatanGigiMulutDewasaData
}

func (r *SkriningKesehatanGigiMulutDewasaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningKesehatanGigiMulutDewasaUpdate) Rules(ctx http.Context) map[string]any {
	return skriningKesehatanGigiMulutDewasaRules()
}

func (r *SkriningKesehatanGigiMulutDewasaUpdate) Payload() SkriningKesehatanGigiMulutDewasaData {
	return r.SkriningKesehatanGigiMulutDewasaData
}

func (r *SkriningKesehatanGigiMulutDewasaUpdate) DetailValues() map[string][]string { return nil }
