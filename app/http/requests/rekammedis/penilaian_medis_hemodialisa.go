package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisHemodialisaData isian penilaian awal medis hemodialisa.
type PenilaianMedisHemodialisaData struct {
	Tanggal                       string `form:"tanggal" json:"tanggal"`
	KdDokter                      string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis                     string `form:"anamnesis" json:"anamnesis"`
	Hubungan                      string `form:"hubungan" json:"hubungan"`
	Ruangan                       string `form:"ruangan" json:"ruangan"`
	Alergi                        string `form:"alergi" json:"alergi"`
	Nyeri                         string `form:"nyeri" json:"nyeri"`
	StatusNutrisi                 string `form:"status_nutrisi" json:"status_nutrisi"`
	Hipertensi                    string `form:"hipertensi" json:"hipertensi"`
	KeteranganHipertensi          string `form:"keterangan_hipertensi" json:"keterangan_hipertensi"`
	Diabetes                      string `form:"diabetes" json:"diabetes"`
	KeteranganDiabetes            string `form:"keterangan_diabetes" json:"keterangan_diabetes"`
	BatuSaluranKemih              string `form:"batu_saluran_kemih" json:"batu_saluran_kemih"`
	KeteranganBatuSaluranKemih    string `form:"keterangan_batu_saluran_kemih" json:"keterangan_batu_saluran_kemih"`
	OperasiSaluranKemih           string `form:"operasi_saluran_kemih" json:"operasi_saluran_kemih"`
	KeteranganOperasiSaluranKemih string `form:"keterangan_operasi_saluran_kemih" json:"keterangan_operasi_saluran_kemih"`
	InfeksiSaluranKemih           string `form:"infeksi_saluran_kemih" json:"infeksi_saluran_kemih"`
	KeteranganInfeksiSaluranKemih string `form:"keterangan_infeksi_saluran_kemih" json:"keterangan_infeksi_saluran_kemih"`
	BengkakSeluruhTubuh           string `form:"bengkak_seluruh_tubuh" json:"bengkak_seluruh_tubuh"`
	KeteranganBengkakSeluruhTubuh string `form:"keterangan_bengkak_seluruh_tubuh" json:"keterangan_bengkak_seluruh_tubuh"`
	UrinBerdarah                  string `form:"urin_berdarah" json:"urin_berdarah"`
	KeteranganUrinBerdarah        string `form:"keterangan_urin_berdarah" json:"keterangan_urin_berdarah"`
	PenyakitGinjalLaom            string `form:"penyakit_ginjal_laom" json:"penyakit_ginjal_laom"`
	KeteranganPenyakitGinjalLaom  string `form:"keterangan_penyakit_ginjal_laom" json:"keterangan_penyakit_ginjal_laom"`
	PenyakitLain                  string `form:"penyakit_lain" json:"penyakit_lain"`
	KeteranganPenyakitLain        string `form:"keterangan_penyakit_lain" json:"keterangan_penyakit_lain"`
	KonsumsiObatNefro             string `form:"konsumsi_obat_nefro" json:"konsumsi_obat_nefro"`
	KeteranganKonsumsiObatNefro   string `form:"keterangan_konsumsi_obat_nefro" json:"keterangan_konsumsi_obat_nefro"`
	DialisisPertama               string `form:"dialisis_pertama" json:"dialisis_pertama"`
	PernahCpad                    string `form:"pernah_cpad" json:"pernah_cpad"`
	TanggalCpad                   string `form:"tanggal_cpad" json:"tanggal_cpad"`
	PernahTransplantasi           string `form:"pernah_transplantasi" json:"pernah_transplantasi"`
	TanggalTransplantasi          string `form:"tanggal_transplantasi" json:"tanggal_transplantasi"`
	KeadaanUmum                   string `form:"keadaan_umum" json:"keadaan_umum"`
	Kesadaran                     string `form:"kesadaran" json:"kesadaran"`
	Nadi                          string `form:"nadi" json:"nadi"`
	Bb                            string `form:"bb" json:"bb"`
	Td                            string `form:"td" json:"td"`
	Suhu                          string `form:"suhu" json:"suhu"`
	Napas                         string `form:"napas" json:"napas"`
	Tb                            string `form:"tb" json:"tb"`
	Hepatomegali                  string `form:"hepatomegali" json:"hepatomegali"`
	Splenomegali                  string `form:"splenomegali" json:"splenomegali"`
	Ascites                       string `form:"ascites" json:"ascites"`
	Edema                         string `form:"edema" json:"edema"`
	Whezzing                      string `form:"whezzing" json:"whezzing"`
	Ronchi                        string `form:"ronchi" json:"ronchi"`
	Ikterik                       string `form:"ikterik" json:"ikterik"`
	TekananVena                   string `form:"tekanan_vena" json:"tekanan_vena"`
	Anemia                        string `form:"anemia" json:"anemia"`
	Kardiomegali                  string `form:"kardiomegali" json:"kardiomegali"`
	Bising                        string `form:"bising" json:"bising"`
	Thorax                        string `form:"thorax" json:"thorax"`
	TanggalThorax                 string `form:"tanggal_thorax" json:"tanggal_thorax"`
	Ekg                           string `form:"ekg" json:"ekg"`
	TanggalEkg                    string `form:"tanggal_ekg" json:"tanggal_ekg"`
	Bno                           string `form:"bno" json:"bno"`
	TanggalBno                    string `form:"tanggal_bno" json:"tanggal_bno"`
	Usg                           string `form:"usg" json:"usg"`
	TanggalUsg                    string `form:"tanggal_usg" json:"tanggal_usg"`
	Renogram                      string `form:"renogram" json:"renogram"`
	TanggalRenogram               string `form:"tanggal_renogram" json:"tanggal_renogram"`
	Biopsi                        string `form:"biopsi" json:"biopsi"`
	TanggalBiopsi                 string `form:"tanggal_biopsi" json:"tanggal_biopsi"`
	Ctscan                        string `form:"ctscan" json:"ctscan"`
	TanggalCtscan                 string `form:"tanggal_ctscan" json:"tanggal_ctscan"`
	Arteriografi                  string `form:"arteriografi" json:"arteriografi"`
	TanggalArteriografi           string `form:"tanggal_arteriografi" json:"tanggal_arteriografi"`
	KulturUrin                    string `form:"kultur_urin" json:"kultur_urin"`
	TanggalKulturUrin             string `form:"tanggal_kultur_urin" json:"tanggal_kultur_urin"`
	Laborat                       string `form:"laborat" json:"laborat"`
	TanggalLaborat                string `form:"tanggal_laborat" json:"tanggal_laborat"`
	Hematokrit                    string `form:"hematokrit" json:"hematokrit"`
	Hemoglobin                    string `form:"hemoglobin" json:"hemoglobin"`
	Leukosit                      string `form:"leukosit" json:"leukosit"`
	Trombosit                     string `form:"trombosit" json:"trombosit"`
	HitungJenis                   string `form:"hitung_jenis" json:"hitung_jenis"`
	Ureum                         string `form:"ureum" json:"ureum"`
	UrinLengkap                   string `form:"urin_lengkap" json:"urin_lengkap"`
	Kreatinin                     string `form:"kreatinin" json:"kreatinin"`
	Cct                           string `form:"cct" json:"cct"`
	Sgot                          string `form:"sgot" json:"sgot"`
	Sgpt                          string `form:"sgpt" json:"sgpt"`
	Ct                            string `form:"ct" json:"ct"`
	AsamUrat                      string `form:"asam_urat" json:"asam_urat"`
	Hbsag                         string `form:"hbsag" json:"hbsag"`
	AntiHcv                       string `form:"anti_hcv" json:"anti_hcv"`
	Edukasi                       string `form:"edukasi" json:"edukasi"`
}

func penilaianMedisHemodialisaRules() map[string]any {
	rules := map[string]any{
		"tanggal":                          "date",
		"kd_dokter":                        "string|max_len:20",
		"anamnesis":                        "in:Autoanamnesis,Alloanamnesis",
		"hubungan":                         "string|max_len:30",
		"ruangan":                          "string|max_len:50",
		"alergi":                           "string|max_len:100",
		"nyeri":                            "in:Tidak Nyeri,Nyeri Ringan,Nyeri Sedang,Nyeri Berat,Nyeri Sangat Berat,Nyeri Tak Tertahankan",
		"status_nutrisi":                   "string|max_len:100",
		"hipertensi":                       "in:Tidak,Ya",
		"keterangan_hipertensi":            "string|max_len:30",
		"diabetes":                         "in:Tidak,Ya",
		"keterangan_diabetes":              "string|max_len:30",
		"batu_saluran_kemih":               "in:Tidak,Ya",
		"keterangan_batu_saluran_kemih":    "string|max_len:30",
		"operasi_saluran_kemih":            "in:Tidak,Ya",
		"keterangan_operasi_saluran_kemih": "string|max_len:30",
		"infeksi_saluran_kemih":            "in:Tidak,Ya",
		"keterangan_infeksi_saluran_kemih": "string|max_len:30",
		"bengkak_seluruh_tubuh":            "in:Tidak,Ya",
		"keterangan_bengkak_seluruh_tubuh": "string|max_len:30",
		"urin_berdarah":                    "in:Tidak,Ya",
		"keterangan_urin_berdarah":         "string|max_len:30",
		"penyakit_ginjal_laom":             "in:Tidak,Ya",
		"keterangan_penyakit_ginjal_laom":  "string|max_len:30",
		"penyakit_lain":                    "in:Tidak,Ya",
		"keterangan_penyakit_lain":         "string|max_len:30",
		"konsumsi_obat_nefro":              "in:Tidak,Ya",
		"keterangan_konsumsi_obat_nefro":   "string|max_len:30",
		"dialisis_pertama":                 "required|date",
		"pernah_cpad":                      "required|in:Tidak,Ya",
		"tanggal_cpad":                     "required|date",
		"pernah_transplantasi":             "required|in:Tidak,Ya",
		"tanggal_transplantasi":            "required|date",
		"keadaan_umum":                     "required|in:Sehat,Sakit Ringan,Sakit Sedang,Sakit Berat",
		"kesadaran":                        "required|in:Compos Mentis,Apatis,Somnolen,Sopor,Koma",
		"nadi":                             "string|max_len:5",
		"bb":                               "string|max_len:5",
		"td":                               "string|max_len:8",
		"suhu":                             "string|max_len:5",
		"napas":                            "string|max_len:5",
		"tb":                               "string|max_len:5",
		"hepatomegali":                     "required|in:Tidak,Ya",
		"splenomegali":                     "required|in:Tidak,Ya",
		"ascites":                          "required|in:Tidak,Ya",
		"edema":                            "required|in:Tidak,Ya",
		"whezzing":                         "required|in:Tidak,Ya",
		"ronchi":                           "required|in:Tidak,Ya",
		"ikterik":                          "required|in:Tidak,Ya",
		"tekanan_vena":                     "required|in:Normal,Meningkat",
		"anemia":                           "required|in:Tidak,Ya",
		"kardiomegali":                     "required|in:Tidak,Ya",
		"bising":                           "required|in:Tidak,Ya",
		"thorax":                           "required|in:Tidak,Ya",
		"tanggal_thorax":                   "required|date",
		"ekg":                              "required|in:Tidak,Ya",
		"tanggal_ekg":                      "required|date",
		"bno":                              "required|in:Tidak,Ya",
		"tanggal_bno":                      "required|date",
		"usg":                              "required|in:Tidak,Ya",
		"tanggal_usg":                      "required|date",
		"renogram":                         "required|in:Tidak,Ya",
		"tanggal_renogram":                 "required|date",
		"biopsi":                           "required|in:Tidak,Ya",
		"tanggal_biopsi":                   "required|date",
		"ctscan":                           "required|in:Tidak,Ya",
		"tanggal_ctscan":                   "required|date",
		"arteriografi":                     "required|in:Tidak,Ya",
		"tanggal_arteriografi":             "required|date",
		"kultur_urin":                      "required|in:Tidak,Ya",
		"tanggal_kultur_urin":              "required|date",
		"laborat":                          "required|in:Tidak,Ya",
		"tanggal_laborat":                  "required|date",
		"hematokrit":                       "string|max_len:30",
		"hemoglobin":                       "string|max_len:30",
		"leukosit":                         "string|max_len:30",
		"trombosit":                        "string|max_len:30",
		"hitung_jenis":                     "string|max_len:30",
		"ureum":                            "string|max_len:30",
		"urin_lengkap":                     "string|max_len:30",
		"kreatinin":                        "string|max_len:30",
		"cct":                              "string|max_len:30",
		"sgot":                             "string|max_len:30",
		"sgpt":                             "string|max_len:30",
		"ct":                               "string|max_len:30",
		"asam_urat":                        "string|max_len:30",
		"hbsag":                            "required|in:Non Reaktif,Reaktif",
		"anti_hcv":                         "required|in:Non Reaktif,Reaktif",
		"edukasi":                          "string|max_len:1000",
	}
	return rules
}

// PenilaianMedisHemodialisaStore simpan penilaian awal medis hemodialisa; kolom waktu kunci kosong = sekarang.
type PenilaianMedisHemodialisaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisHemodialisaData
}

func (r *PenilaianMedisHemodialisaStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisHemodialisaStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisHemodialisaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisHemodialisaStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianMedisHemodialisaStore) Payload() PenilaianMedisHemodialisaData {
	return r.PenilaianMedisHemodialisaData
}

func (r *PenilaianMedisHemodialisaStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisHemodialisaUpdate ubah penilaian awal medis hemodialisa (PUT); kunci lewat query string.
type PenilaianMedisHemodialisaUpdate struct {
	PenilaianMedisHemodialisaData
}

func (r *PenilaianMedisHemodialisaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisHemodialisaUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisHemodialisaRules()
}

func (r *PenilaianMedisHemodialisaUpdate) Payload() PenilaianMedisHemodialisaData {
	return r.PenilaianMedisHemodialisaData
}

func (r *PenilaianMedisHemodialisaUpdate) DetailValues() map[string][]string { return nil }
