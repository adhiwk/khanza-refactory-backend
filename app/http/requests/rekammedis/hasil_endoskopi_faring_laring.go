package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilEndoskopiFaringLaringData isian hasil endoskopi faring laring.
type HasilEndoskopiFaringLaringData struct {
	Tanggal                  string `form:"tanggal" json:"tanggal"`
	KdDokter                 string `form:"kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis           string `form:"diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari              string `form:"kiriman_dari" json:"kiriman_dari"`
	FaringUvula              string `form:"faring_uvula" json:"faring_uvula"`
	FaringArkusFaring        string `form:"faring_arkus_faring" json:"faring_arkus_faring"`
	FaringDindingPosterior   string `form:"faring_dinding_posterior" json:"faring_dinding_posterior"`
	FaringTonsil             string `form:"faring_tonsil" json:"faring_tonsil"`
	LaringTonsilLingual      string `form:"laring_tonsil_lingual" json:"laring_tonsil_lingual"`
	LaringValekula           string `form:"laring_valekula" json:"laring_valekula"`
	LaringSinusPiriformis    string `form:"laring_sinus_piriformis" json:"laring_sinus_piriformis"`
	LaringEpiglotis          string `form:"laring_epiglotis" json:"laring_epiglotis"`
	LaringArytenoid          string `form:"laring_arytenoid" json:"laring_arytenoid"`
	LaringPlikaVentrikularis string `form:"laring_plika_ventrikularis" json:"laring_plika_ventrikularis"`
	LaringPitaSuara          string `form:"laring_pita_suara" json:"laring_pita_suara"`
	LaringRimaVocalis        string `form:"laring_rima_vocalis" json:"laring_rima_vocalis"`
	LaringLainlain           string `form:"laring_lainlain" json:"laring_lainlain"`
	Kesan                    string `form:"kesan" json:"kesan"`
	Saran                    string `form:"saran" json:"saran"`
}

func hasilEndoskopiFaringLaringRules() map[string]any {
	rules := map[string]any{
		"tanggal":                    "required|date",
		"kd_dokter":                  "required|string|max_len:20",
		"diagnosa_klinis":            "string|max_len:50",
		"kiriman_dari":               "string|max_len:50",
		"faring_uvula":               "string|max_len:50",
		"faring_arkus_faring":        "string|max_len:50",
		"faring_dinding_posterior":   "string|max_len:50",
		"faring_tonsil":              "string|max_len:50",
		"laring_tonsil_lingual":      "string|max_len:50",
		"laring_valekula":            "string|max_len:50",
		"laring_sinus_piriformis":    "string|max_len:50",
		"laring_epiglotis":           "string|max_len:50",
		"laring_arytenoid":           "string|max_len:50",
		"laring_plika_ventrikularis": "string|max_len:50",
		"laring_pita_suara":          "string|max_len:50",
		"laring_rima_vocalis":        "string|max_len:50",
		"laring_lainlain":            "string|max_len:100",
		"kesan":                      "string|max_len:300",
		"saran":                      "string|max_len:300",
	}
	return rules
}

// HasilEndoskopiFaringLaringStore simpan hasil endoskopi faring laring; kolom waktu kunci kosong = sekarang.
type HasilEndoskopiFaringLaringStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	HasilEndoskopiFaringLaringData
}

func (r *HasilEndoskopiFaringLaringStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilEndoskopiFaringLaringStore) Rules(ctx http.Context) map[string]any {
	rules := hasilEndoskopiFaringLaringRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *HasilEndoskopiFaringLaringStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *HasilEndoskopiFaringLaringStore) Payload() HasilEndoskopiFaringLaringData {
	return r.HasilEndoskopiFaringLaringData
}

func (r *HasilEndoskopiFaringLaringStore) DetailValues() map[string][]string { return nil }

// HasilEndoskopiFaringLaringUpdate ubah hasil endoskopi faring laring (PUT); kunci lewat query string.
type HasilEndoskopiFaringLaringUpdate struct {
	HasilEndoskopiFaringLaringData
}

func (r *HasilEndoskopiFaringLaringUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilEndoskopiFaringLaringUpdate) Rules(ctx http.Context) map[string]any {
	return hasilEndoskopiFaringLaringRules()
}

func (r *HasilEndoskopiFaringLaringUpdate) Payload() HasilEndoskopiFaringLaringData {
	return r.HasilEndoskopiFaringLaringData
}

func (r *HasilEndoskopiFaringLaringUpdate) DetailValues() map[string][]string { return nil }
