package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PemantauanPewsDewasaData isian pemantauan EWSD.
type PemantauanPewsDewasaData struct {
	ParameterLajuRespirasi        string `form:"parameter_laju_respirasi" json:"parameter_laju_respirasi"`
	SkorLajuRespirasi             string `form:"skor_laju_respirasi" json:"skor_laju_respirasi"`
	ParameterSaturasiOksigen      string `form:"parameter_saturasi_oksigen" json:"parameter_saturasi_oksigen"`
	SkorSaturasiOksigen           string `form:"skor_saturasi_oksigen" json:"skor_saturasi_oksigen"`
	ParameterSuplemenOksigen      string `form:"parameter_suplemen_oksigen" json:"parameter_suplemen_oksigen"`
	SkorSuplemenOksigen           string `form:"skor_suplemen_oksigen" json:"skor_suplemen_oksigen"`
	ParameterTekananDarahSistolik string `form:"parameter_tekanan_darah_sistolik" json:"parameter_tekanan_darah_sistolik"`
	SkorTekananDarahSistolik      string `form:"skor_tekanan_darah_sistolik" json:"skor_tekanan_darah_sistolik"`
	ParameterLajuJantung          string `form:"parameter_laju_jantung" json:"parameter_laju_jantung"`
	SkorLajuJantung               string `form:"skor_laju_jantung" json:"skor_laju_jantung"`
	ParameterKesadaran            string `form:"parameter_kesadaran" json:"parameter_kesadaran"`
	SkorKesadaran                 string `form:"skor_kesadaran" json:"skor_kesadaran"`
	ParameterTemperatur           string `form:"parameter_temperatur" json:"parameter_temperatur"`
	SkorTemperatur                string `form:"skor_temperatur" json:"skor_temperatur"`
	SkorTotal                     string `form:"skor_total" json:"skor_total"`
	ParameterTotal                string `form:"parameter_total" json:"parameter_total"`
	Nip                           string `form:"nip" json:"nip"`
}

func pemantauanPewsDewasaRules() map[string]any {
	rules := map[string]any{
		"parameter_laju_respirasi":         "in:<= 5,6 - 8,9 - 11,12 - 20,21 - 24,25 - 34,>= 35",
		"skor_laju_respirasi":              "string|max_len:1",
		"parameter_saturasi_oksigen":       "in:>= 95,94 - 95,92 - 93,<= 92",
		"skor_saturasi_oksigen":            "string|max_len:1",
		"parameter_suplemen_oksigen":       "in:Ya,Tidak",
		"skor_suplemen_oksigen":            "string|max_len:1",
		"parameter_tekanan_darah_sistolik": "in:>= 220,181 - 220,111 - 180,101 - 110,91 - 100,71 - 90,<= 70",
		"skor_tekanan_darah_sistolik":      "string|max_len:1",
		"parameter_laju_jantung":           "in:>= 140,131 - 140,111 - 130,91 - 110,51 - 90,41 - 50,<= 40",
		"skor_laju_jantung":                "string|max_len:1",
		"parameter_kesadaran":              "in:Sadar,Nyeri/Verbal,Unrespon",
		"skor_kesadaran":                   "string|max_len:1",
		"parameter_temperatur":             "in:<= 35,35.1 - 36,36.1 - 38,38.1 - 39,>= 39",
		"skor_temperatur":                  "string|max_len:1",
		"skor_total":                       "string|max_len:2",
		"parameter_total":                  "string|max_len:250",
		"nip":                              "string|max_len:20",
	}
	return rules
}

// PemantauanPewsDewasaStore simpan pemantauan EWSD; kolom waktu kunci kosong = sekarang.
type PemantauanPewsDewasaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PemantauanPewsDewasaData
}

func (r *PemantauanPewsDewasaStore) Authorize(ctx http.Context) error { return nil }

func (r *PemantauanPewsDewasaStore) Rules(ctx http.Context) map[string]any {
	rules := pemantauanPewsDewasaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PemantauanPewsDewasaStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PemantauanPewsDewasaStore) Payload() PemantauanPewsDewasaData {
	return r.PemantauanPewsDewasaData
}

func (r *PemantauanPewsDewasaStore) DetailValues() map[string][]string { return nil }

// PemantauanPewsDewasaUpdate ubah pemantauan EWSD (PUT); kunci lewat query string.
type PemantauanPewsDewasaUpdate struct {
	PemantauanPewsDewasaData
}

func (r *PemantauanPewsDewasaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PemantauanPewsDewasaUpdate) Rules(ctx http.Context) map[string]any {
	return pemantauanPewsDewasaRules()
}

func (r *PemantauanPewsDewasaUpdate) Payload() PemantauanPewsDewasaData {
	return r.PemantauanPewsDewasaData
}

func (r *PemantauanPewsDewasaUpdate) DetailValues() map[string][]string { return nil }
