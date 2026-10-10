package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningPneumoniaSeverityIndexData isian skrining pneumonia severity index.
type SkriningPneumoniaSeverityIndexData struct {
	Tanggal                      string `form:"tanggal" json:"tanggal"`
	NilaiUmur                    *int   `form:"nilai_umur" json:"nilai_umur"`
	TinggalDiPantiJompo          string `form:"tinggal_di_panti_jompo" json:"tinggal_di_panti_jompo"`
	NilaiTinggalDiPantiJompo     *int   `form:"nilai_tinggal_di_panti_jompo" json:"nilai_tinggal_di_panti_jompo"`
	GagalJantung                 string `form:"gagal_jantung" json:"gagal_jantung"`
	NilaiGagalJantung            *int   `form:"nilai_gagal_jantung" json:"nilai_gagal_jantung"`
	PenyakitHati                 string `form:"penyakit_hati" json:"penyakit_hati"`
	NilaiPenyakitHati            *int   `form:"nilai_penyakit_hati" json:"nilai_penyakit_hati"`
	PenyakitGinjal               string `form:"penyakit_ginjal" json:"penyakit_ginjal"`
	NilaiPenyakitGinjal          *int   `form:"nilai_penyakit_ginjal" json:"nilai_penyakit_ginjal"`
	Kanker                       string `form:"kanker" json:"kanker"`
	NilaiKanker                  *int   `form:"nilai_kanker" json:"nilai_kanker"`
	PenyakitSerebrovaskuler      string `form:"penyakit_serebrovaskuler" json:"penyakit_serebrovaskuler"`
	NilaiPenyakitSerebrovaskuler *int   `form:"nilai_penyakit_serebrovaskuler" json:"nilai_penyakit_serebrovaskuler"`
	DisorentasiMental            string `form:"disorentasi_mental" json:"disorentasi_mental"`
	NilaiDisorentasiMental       *int   `form:"nilai_disorentasi_mental" json:"nilai_disorentasi_mental"`
	FrekuensiNapas               string `form:"frekuensi_napas" json:"frekuensi_napas"`
	NilaiFrekuensiNapas          *int   `form:"nilai_frekuensi_napas" json:"nilai_frekuensi_napas"`
	TdSistolik                   string `form:"td_sistolik" json:"td_sistolik"`
	NilaiTdSistolik              *int   `form:"nilai_td_sistolik" json:"nilai_td_sistolik"`
	Suhu                         string `form:"suhu" json:"suhu"`
	NilaiSuhu                    *int   `form:"nilai_suhu" json:"nilai_suhu"`
	Nadi                         string `form:"nadi" json:"nadi"`
	NilaiNadi                    *int   `form:"nilai_nadi" json:"nilai_nadi"`
	PhDarah                      string `form:"ph_darah" json:"ph_darah"`
	NilaiPhDarah                 *int   `form:"nilai_ph_darah" json:"nilai_ph_darah"`
	Natrium                      string `form:"natrium" json:"natrium"`
	NilaiNatrium                 *int   `form:"nilai_natrium" json:"nilai_natrium"`
	Bun                          string `form:"bun" json:"bun"`
	NilaiBun                     *int   `form:"nilai_bun" json:"nilai_bun"`
	Pao                          string `form:"pao" json:"pao"`
	NilaiPao                     *int   `form:"nilai_pao" json:"nilai_pao"`
	Glukosa                      string `form:"glukosa" json:"glukosa"`
	NilaiGlukosa                 *int   `form:"nilai_glukosa" json:"nilai_glukosa"`
	EfusiPleura                  string `form:"efusi_pleura" json:"efusi_pleura"`
	NilaiEfusiPleura             *int   `form:"nilai_efusi_pleura" json:"nilai_efusi_pleura"`
	Hematokrit                   string `form:"hematokrit" json:"hematokrit"`
	NilaiHematokrit              *int   `form:"nilai_hematokrit" json:"nilai_hematokrit"`
	TotalSkor                    *int   `form:"total_skor" json:"total_skor"`
	Kelas                        string `form:"kelas" json:"kelas"`
	SkorInterpretasi             string `form:"skor_interpretasi" json:"skor_interpretasi"`
	Mortalitas                   string `form:"mortalitas" json:"mortalitas"`
	Rekomendasi                  string `form:"rekomendasi" json:"rekomendasi"`
	Nip                          string `form:"nip" json:"nip"`
}

func skriningPneumoniaSeverityIndexRules() map[string]any {
	rules := map[string]any{
		"tanggal":                        "required|date",
		"nilai_umur":                     "int",
		"tinggal_di_panti_jompo":         "in:Tidak,Ya",
		"nilai_tinggal_di_panti_jompo":   "int",
		"gagal_jantung":                  "in:Tidak Ada,Ada",
		"nilai_gagal_jantung":            "int",
		"penyakit_hati":                  "in:Tidak Ada,Ada",
		"nilai_penyakit_hati":            "int",
		"penyakit_ginjal":                "in:Tidak Ada,Ada",
		"nilai_penyakit_ginjal":          "int",
		"kanker":                         "in:Tidak Ada,Ada",
		"nilai_kanker":                   "int",
		"penyakit_serebrovaskuler":       "in:Tidak Ada,Ada",
		"nilai_penyakit_serebrovaskuler": "int",
		"disorentasi_mental":             "in:Tidak,Ya",
		"nilai_disorentasi_mental":       "int",
		"frekuensi_napas":                "in:Tidak,Ya",
		"nilai_frekuensi_napas":          "int",
		"td_sistolik":                    "in:Tidak,Ya",
		"nilai_td_sistolik":              "int",
		"suhu":                           "in:Tidak,Ya",
		"nilai_suhu":                     "int",
		"nadi":                           "in:Tidak,Ya",
		"nilai_nadi":                     "int",
		"ph_darah":                       "in:Tidak,Ya",
		"nilai_ph_darah":                 "int",
		"natrium":                        "in:Tidak,Ya",
		"nilai_natrium":                  "int",
		"bun":                            "in:Tidak,Ya",
		"nilai_bun":                      "int",
		"pao":                            "in:Tidak,Ya",
		"nilai_pao":                      "int",
		"glukosa":                        "in:Tidak,Ya",
		"nilai_glukosa":                  "int",
		"efusi_pleura":                   "in:Tidak,Ya",
		"nilai_efusi_pleura":             "int",
		"hematokrit":                     "in:Tidak,Ya",
		"nilai_hematokrit":               "int",
		"total_skor":                     "int",
		"kelas":                          "in:I,II,III,IV,V",
		"skor_interpretasi":              "in:Ditentukan Algoritma Awal,?70,71â90,91â130,>130",
		"mortalitas":                     "in:<0.4%,~0.6%,~2.8%,~8â9%,~27â30%",
		"rekomendasi":                    "string",
		"nip":                            "required|string|max_len:20",
	}
	return rules
}

// SkriningPneumoniaSeverityIndexStore simpan skrining pneumonia severity index; kolom waktu kunci kosong = sekarang.
type SkriningPneumoniaSeverityIndexStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningPneumoniaSeverityIndexData
}

func (r *SkriningPneumoniaSeverityIndexStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningPneumoniaSeverityIndexStore) Rules(ctx http.Context) map[string]any {
	rules := skriningPneumoniaSeverityIndexRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningPneumoniaSeverityIndexStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *SkriningPneumoniaSeverityIndexStore) Payload() SkriningPneumoniaSeverityIndexData {
	return r.SkriningPneumoniaSeverityIndexData
}

func (r *SkriningPneumoniaSeverityIndexStore) DetailValues() map[string][]string { return nil }

// SkriningPneumoniaSeverityIndexUpdate ubah skrining pneumonia severity index (PUT); kunci lewat query string.
type SkriningPneumoniaSeverityIndexUpdate struct {
	SkriningPneumoniaSeverityIndexData
}

func (r *SkriningPneumoniaSeverityIndexUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningPneumoniaSeverityIndexUpdate) Rules(ctx http.Context) map[string]any {
	return skriningPneumoniaSeverityIndexRules()
}

func (r *SkriningPneumoniaSeverityIndexUpdate) Payload() SkriningPneumoniaSeverityIndexData {
	return r.SkriningPneumoniaSeverityIndexData
}

func (r *SkriningPneumoniaSeverityIndexUpdate) DetailValues() map[string][]string { return nil }
