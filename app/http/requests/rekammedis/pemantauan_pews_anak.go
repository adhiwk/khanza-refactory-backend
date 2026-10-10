package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PemantauanPewsAnakData isian pemantauan PEWS.
type PemantauanPewsAnakData struct {
	ParameterPerilaku          string `form:"parameter_perilaku" json:"parameter_perilaku"`
	SkorPerilaku               string `form:"skor_perilaku" json:"skor_perilaku"`
	ParameterCrtAtauWarnaKulit string `form:"parameter_crt_atau_warna_kulit" json:"parameter_crt_atau_warna_kulit"`
	SkorCrtAtauWarnaKulit      string `form:"skor_crt_atau_warna_kulit" json:"skor_crt_atau_warna_kulit"`
	ParameterPerespirasi       string `form:"parameter_perespirasi" json:"parameter_perespirasi"`
	SkorPerespirasi            string `form:"skor_perespirasi" json:"skor_perespirasi"`
	SkorTotal                  string `form:"skor_total" json:"skor_total"`
	ParameterTotal             string `form:"parameter_total" json:"parameter_total"`
	Nip                        string `form:"nip" json:"nip"`
}

func pemantauanPewsAnakRules() map[string]any {
	rules := map[string]any{
		"parameter_perilaku":             "in:Sadar / Bermain,Tidur / Perubahan Perilaku,Gelisah,Tidak Merespon Terhadap Nyeri Penurunan Kesadaran",
		"skor_perilaku":                  "string|max_len:1",
		"parameter_crt_atau_warna_kulit": "in:1 - 2 dtk / Pink,3 dtk / Pucat,4 dtk / Sianosis,>=5 dtk / Mottle",
		"skor_crt_atau_warna_kulit":      "string|max_len:1",
		"parameter_perespirasi":          "in:Tidak Ada Retraksi,Cuping Hidung / O2 1-3 Lpm,Retraksi Dada / O2 4-6 Lpm,Stridor / O2 7-8 Lpm",
		"skor_perespirasi":               "string|max_len:1",
		"skor_total":                     "string|max_len:1",
		"parameter_total":                "string|max_len:250",
		"nip":                            "string|max_len:20",
	}
	return rules
}

// PemantauanPewsAnakStore simpan pemantauan PEWS; kolom waktu kunci kosong = sekarang.
type PemantauanPewsAnakStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PemantauanPewsAnakData
}

func (r *PemantauanPewsAnakStore) Authorize(ctx http.Context) error { return nil }

func (r *PemantauanPewsAnakStore) Rules(ctx http.Context) map[string]any {
	rules := pemantauanPewsAnakRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PemantauanPewsAnakStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PemantauanPewsAnakStore) Payload() PemantauanPewsAnakData { return r.PemantauanPewsAnakData }

func (r *PemantauanPewsAnakStore) DetailValues() map[string][]string { return nil }

// PemantauanPewsAnakUpdate ubah pemantauan PEWS (PUT); kunci lewat query string.
type PemantauanPewsAnakUpdate struct {
	PemantauanPewsAnakData
}

func (r *PemantauanPewsAnakUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PemantauanPewsAnakUpdate) Rules(ctx http.Context) map[string]any {
	return pemantauanPewsAnakRules()
}

func (r *PemantauanPewsAnakUpdate) Payload() PemantauanPewsAnakData { return r.PemantauanPewsAnakData }

func (r *PemantauanPewsAnakUpdate) DetailValues() map[string][]string { return nil }
