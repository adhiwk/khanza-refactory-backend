package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningKesehatanGigiMulutLansiaData isian skrining kesehatan gigi mulut lansia.
type SkriningKesehatanGigiMulutLansiaData struct {
	Tanggal       string `form:"tanggal" json:"tanggal"`
	KontrolGigi   string `form:"kontrol_gigi" json:"kontrol_gigi"`
	PolaMakan     string `form:"pola_makan" json:"pola_makan"`
	SikatGigi     string `form:"sikat_gigi" json:"sikat_gigi"`
	GigiPalsu     string `form:"gigi_palsu" json:"gigi_palsu"`
	GigiBerfungsi string `form:"gigi_berfungsi" json:"gigi_berfungsi"`
	MukosaMulut   string `form:"mukosa_mulut" json:"mukosa_mulut"`
	HasilSkrining string `form:"hasil_skrining" json:"hasil_skrining"`
	Keterangan    string `form:"keterangan" json:"keterangan"`
	Nip           string `form:"nip" json:"nip"`
}

func skriningKesehatanGigiMulutLansiaRules() map[string]any {
	rules := map[string]any{
		"tanggal":        "required|date",
		"kontrol_gigi":   "in:Ya,Tidak",
		"pola_makan":     "in:Ya,Tidak",
		"sikat_gigi":     "in:Ya,Tidak",
		"gigi_palsu":     "in:Ya,Tidak",
		"gigi_berfungsi": "in:Ya,Tidak",
		"mukosa_mulut":   "in:Ya,Tidak",
		"hasil_skrining": "string|max_len:50",
		"keterangan":     "string|max_len:100",
		"nip":            "required|string|max_len:20",
	}
	return rules
}

// SkriningKesehatanGigiMulutLansiaStore simpan skrining kesehatan gigi mulut lansia; kolom waktu kunci kosong = sekarang.
type SkriningKesehatanGigiMulutLansiaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningKesehatanGigiMulutLansiaData
}

func (r *SkriningKesehatanGigiMulutLansiaStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningKesehatanGigiMulutLansiaStore) Rules(ctx http.Context) map[string]any {
	rules := skriningKesehatanGigiMulutLansiaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningKesehatanGigiMulutLansiaStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *SkriningKesehatanGigiMulutLansiaStore) Payload() SkriningKesehatanGigiMulutLansiaData {
	return r.SkriningKesehatanGigiMulutLansiaData
}

func (r *SkriningKesehatanGigiMulutLansiaStore) DetailValues() map[string][]string { return nil }

// SkriningKesehatanGigiMulutLansiaUpdate ubah skrining kesehatan gigi mulut lansia (PUT); kunci lewat query string.
type SkriningKesehatanGigiMulutLansiaUpdate struct {
	SkriningKesehatanGigiMulutLansiaData
}

func (r *SkriningKesehatanGigiMulutLansiaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningKesehatanGigiMulutLansiaUpdate) Rules(ctx http.Context) map[string]any {
	return skriningKesehatanGigiMulutLansiaRules()
}

func (r *SkriningKesehatanGigiMulutLansiaUpdate) Payload() SkriningKesehatanGigiMulutLansiaData {
	return r.SkriningKesehatanGigiMulutLansiaData
}

func (r *SkriningKesehatanGigiMulutLansiaUpdate) DetailValues() map[string][]string { return nil }
