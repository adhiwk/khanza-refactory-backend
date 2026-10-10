package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PemantauanMeowsObstetriData isian pemantauan MEOWS.
type PemantauanMeowsObstetriData struct {
	ParameterPernapasan           string `form:"parameter_pernapasan" json:"parameter_pernapasan"`
	SkorPernapasan                string `form:"skor_pernapasan" json:"skor_pernapasan"`
	ParameterSaturasi             string `form:"parameter_saturasi" json:"parameter_saturasi"`
	SkorSaturasi                  string `form:"skor_saturasi" json:"skor_saturasi"`
	ParameterTemperatur           string `form:"parameter_temperatur" json:"parameter_temperatur"`
	SkorTemperatur                string `form:"skor_temperatur" json:"skor_temperatur"`
	ParameterTekananDarahSistole  string `form:"parameter_tekanan_darah_sistole" json:"parameter_tekanan_darah_sistole"`
	SkorTekananDarahSistole       string `form:"skor_tekanan_darah_sistole" json:"skor_tekanan_darah_sistole"`
	ParameterTekananDarahDiastole string `form:"parameter_tekanan_darah_diastole" json:"parameter_tekanan_darah_diastole"`
	SkorTekananDarahDiastole      string `form:"skor_tekanan_darah_diastole" json:"skor_tekanan_darah_diastole"`
	ParameterDenyutJantung        string `form:"parameter_denyut_jantung" json:"parameter_denyut_jantung"`
	SkorDenyutJantung             string `form:"skor_denyut_jantung" json:"skor_denyut_jantung"`
	ParameterKesadaran            string `form:"parameter_kesadaran" json:"parameter_kesadaran"`
	SkorKesadaran                 string `form:"skor_kesadaran" json:"skor_kesadaran"`
	ParameterKetuban              string `form:"parameter_ketuban" json:"parameter_ketuban"`
	SkorKetuban                   string `form:"skor_ketuban" json:"skor_ketuban"`
	ParameterDischarge            string `form:"parameter_discharge" json:"parameter_discharge"`
	SkorDischarge                 string `form:"skor_discharge" json:"skor_discharge"`
	ParameterProteinuria          string `form:"parameter_proteinuria" json:"parameter_proteinuria"`
	SkorProteinuria               string `form:"skor_proteinuria" json:"skor_proteinuria"`
	SkorTotal                     string `form:"skor_total" json:"skor_total"`
	ParameterTotal                string `form:"parameter_total" json:"parameter_total"`
	CodeBlue                      string `form:"code_blue" json:"code_blue"`
	Nip                           string `form:"nip" json:"nip"`
}

func pemantauanMeowsObstetriRules() map[string]any {
	rules := map[string]any{
		"parameter_pernapasan":             "in:>= 30,21 - 30,11 - 20,< 12",
		"skor_pernapasan":                  "string|max_len:1",
		"parameter_saturasi":               "in:> 95,90 - 94,< 90",
		"skor_saturasi":                    "string|max_len:1",
		"parameter_temperatur":             "in:> 38,35 - 35.9,36 - 37.9,< 35",
		"skor_temperatur":                  "string|max_len:1",
		"parameter_tekanan_darah_sistole":  "in:> 160,150 - 159,100 - 140,90 - 99,< 90",
		"skor_tekanan_darah_sistole":       "string|max_len:1",
		"parameter_tekanan_darah_diastole": "in:> 110,90 - 109,< 90",
		"skor_tekanan_darah_diastole":      "string|max_len:1",
		"parameter_denyut_jantung":         "in:> 120,100 - 120,51 - 99,40 - 50,< 40",
		"skor_denyut_jantung":              "string|max_len:1",
		"parameter_kesadaran":              "in:Alert,Verbal,Pain,Unresponsive",
		"skor_kesadaran":                   "string|max_len:1",
		"parameter_ketuban":                "in:Khas,Busuk",
		"skor_ketuban":                     "string|max_len:1",
		"parameter_discharge":              "in:Normal,Banyak",
		"skor_discharge":                   "string|max_len:1",
		"parameter_proteinuria":            "in:Negatif,+,++>",
		"skor_proteinuria":                 "string|max_len:1",
		"skor_total":                       "string|max_len:2",
		"parameter_total":                  "string|max_len:250",
		"code_blue":                        "required|in:Ya,Tidak",
		"nip":                              "string|max_len:20",
	}
	return rules
}

// PemantauanMeowsObstetriStore simpan pemantauan MEOWS; kolom waktu kunci kosong = sekarang.
type PemantauanMeowsObstetriStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PemantauanMeowsObstetriData
}

func (r *PemantauanMeowsObstetriStore) Authorize(ctx http.Context) error { return nil }

func (r *PemantauanMeowsObstetriStore) Rules(ctx http.Context) map[string]any {
	rules := pemantauanMeowsObstetriRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PemantauanMeowsObstetriStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PemantauanMeowsObstetriStore) Payload() PemantauanMeowsObstetriData {
	return r.PemantauanMeowsObstetriData
}

func (r *PemantauanMeowsObstetriStore) DetailValues() map[string][]string { return nil }

// PemantauanMeowsObstetriUpdate ubah pemantauan MEOWS (PUT); kunci lewat query string.
type PemantauanMeowsObstetriUpdate struct {
	PemantauanMeowsObstetriData
}

func (r *PemantauanMeowsObstetriUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PemantauanMeowsObstetriUpdate) Rules(ctx http.Context) map[string]any {
	return pemantauanMeowsObstetriRules()
}

func (r *PemantauanMeowsObstetriUpdate) Payload() PemantauanMeowsObstetriData {
	return r.PemantauanMeowsObstetriData
}

func (r *PemantauanMeowsObstetriUpdate) DetailValues() map[string][]string { return nil }
