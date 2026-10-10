package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningPumaData isian skrining PUMA.
type SkriningPumaData struct {
	Tanggal                 string `form:"tanggal" json:"tanggal"`
	Jk                      string `form:"jk" json:"jk"`
	NilaiJk                 *int   `form:"nilai_jk" json:"nilai_jk"`
	Usia                    string `form:"usia" json:"usia"`
	NilaiUsia               *int   `form:"nilai_usia" json:"nilai_usia"`
	PernahMerokok           string `form:"pernah_merokok" json:"pernah_merokok"`
	NilaiPernahMerokok      *int   `form:"nilai_pernah_merokok" json:"nilai_pernah_merokok"`
	JumlahRokokPerhari      string `form:"jumlah_rokok_perhari" json:"jumlah_rokok_perhari"`
	LamaMerokok             string `form:"lama_merokok" json:"lama_merokok"`
	NapasPendek             string `form:"napas_pendek" json:"napas_pendek"`
	NilaiNapasPendek        *int   `form:"nilai_napas_pendek" json:"nilai_napas_pendek"`
	PunyaDahak              string `form:"punya_dahak" json:"punya_dahak"`
	NilaiPunyaDahak         *int   `form:"nilai_punya_dahak" json:"nilai_punya_dahak"`
	BiasaBatuk              string `form:"biasa_batuk" json:"biasa_batuk"`
	NilaiBiasaBatuk         *int   `form:"nilai_biasa_batuk" json:"nilai_biasa_batuk"`
	Spirometri              string `form:"spirometri" json:"spirometri"`
	NilaiSpirometri         *int   `form:"nilai_spirometri" json:"nilai_spirometri"`
	NilaiTotal              *int   `form:"nilai_total" json:"nilai_total"`
	KeteranganHasilSkrining string `form:"keterangan_hasil_skrining" json:"keterangan_hasil_skrining"`
	Nip                     string `form:"nip" json:"nip"`
}

func skriningPumaRules() map[string]any {
	rules := map[string]any{
		"tanggal":                   "required|date",
		"jk":                        "in:Perempuan,Laki-laki",
		"nilai_jk":                  "int",
		"usia":                      "in:40-49,50-59,>60",
		"nilai_usia":                "int",
		"pernah_merokok":            "in:Pernah,Tidak",
		"nilai_pernah_merokok":      "int",
		"jumlah_rokok_perhari":      "string|max_len:3",
		"lama_merokok":              "string|max_len:3",
		"napas_pendek":              "in:Ya,Tidak",
		"nilai_napas_pendek":        "int",
		"punya_dahak":               "in:Ya,Tidak",
		"nilai_punya_dahak":         "int",
		"biasa_batuk":               "in:Ya,Tidak",
		"nilai_biasa_batuk":         "int",
		"spirometri":                "in:Ya,Tidak",
		"nilai_spirometri":          "int",
		"nilai_total":               "int",
		"keterangan_hasil_skrining": "string|max_len:150",
		"nip":                       "required|string|max_len:20",
	}
	return rules
}

// SkriningPumaStore simpan skrining PUMA; kolom waktu kunci kosong = sekarang.
type SkriningPumaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningPumaData
}

func (r *SkriningPumaStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningPumaStore) Rules(ctx http.Context) map[string]any {
	rules := skriningPumaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningPumaStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningPumaStore) Payload() SkriningPumaData { return r.SkriningPumaData }

func (r *SkriningPumaStore) DetailValues() map[string][]string { return nil }

// SkriningPumaUpdate ubah skrining PUMA (PUT); kunci lewat query string.
type SkriningPumaUpdate struct {
	SkriningPumaData
}

func (r *SkriningPumaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningPumaUpdate) Rules(ctx http.Context) map[string]any { return skriningPumaRules() }

func (r *SkriningPumaUpdate) Payload() SkriningPumaData { return r.SkriningPumaData }

func (r *SkriningPumaUpdate) DetailValues() map[string][]string { return nil }
