package rekonsiliasi

import (
	"github.com/goravel/framework/contracts/http"
)

type Obat struct {
	NamaObat               string `form:"nama_obat" json:"nama_obat"`
	DosisObat              string `form:"dosis_obat" json:"dosis_obat"`
	Frekuensi              string `form:"frekuensi" json:"frekuensi"`
	CaraPemberian          string `form:"cara_pemberian" json:"cara_pemberian"`
	WaktuPemberianTerakhir string `form:"waktu_pemberian_terakhir" json:"waktu_pemberian_terakhir"`
	TindakLanjut           string `form:"tindak_lanjut" json:"tindak_lanjut"`
	PerubahanAturanPakai   string `form:"perubahan_aturan_pakai" json:"perubahan_aturan_pakai"`
}

type Data struct {
	TanggalWawancara     string `form:"tanggal_wawancara" json:"tanggal_wawancara"` // kosong = sekarang
	RekonsiliasiObatSaat string `form:"rekonsiliasi_obat_saat" json:"rekonsiliasi_obat_saat"`
	AlergiObat           string `form:"alergi_obat" json:"alergi_obat"`
	ManifestasiAlergi    string `form:"manifestasi_alergi" json:"manifestasi_alergi"`
	DampakAlergi         string `form:"dampak_alergi" json:"dampak_alergi"`
	Nip                  string `form:"nip" json:"nip"`
	Obat                 []Obat `form:"obat" json:"obat"`
}

func baseRules() map[string]any {
	return map[string]any{
		"tanggal_wawancara":               "date",
		"rekonsiliasi_obat_saat":          "required|in:Admisi,Transfer Antar Ruang,Pindah Faskes Lain,Pulang",
		"alergi_obat":                     "string|max_len:70",
		"manifestasi_alergi":              "string|max_len:70",
		"dampak_alergi":                   "in:-,Ringan,Sedang,Berat",
		"nip":                             "required|string|max_len:20",
		"obat":                            "required|slice",
		"obat.*.nama_obat":                "required|string|max_len:100",
		"obat.*.dosis_obat":               "string|max_len:20",
		"obat.*.frekuensi":                "string|max_len:10",
		"obat.*.cara_pemberian":           "string|max_len:150",
		"obat.*.waktu_pemberian_terakhir": "string|max_len:20",
		"obat.*.tindak_lanjut":            "in:Lanjut,Stop",
		"obat.*.perubahan_aturan_pakai":   "string|max_len:150",
	}
}

type StoreRequest struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error { return nil }

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error { return nil }

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any { return baseRules() }

type KonfirmasiRequest struct {
	DiterimaFarmasi      string `form:"diterima_farmasi" json:"diterima_farmasi"`
	DikonfirmasiApoteker string `form:"dikonfirmasi_apoteker" json:"dikonfirmasi_apoteker"`
	DiserahkanPasien     string `form:"diserahkan_pasien" json:"diserahkan_pasien"`
	Nip                  string `form:"nip" json:"nip"`
}

func (r *KonfirmasiRequest) Authorize(ctx http.Context) error { return nil }

func (r *KonfirmasiRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"diterima_farmasi":      "required|date",
		"dikonfirmasi_apoteker": "required|date",
		"diserahkan_pasien":     "required|date",
		"nip":                   "required|string|max_len:20",
	}
}
