package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianTerapiWicaraData isian penilaian terapi wicara.
type PenilaianTerapiWicaraData struct {
	Tanggal                              string `form:"tanggal" json:"tanggal"`
	DiagnosaTerapiWicara                 string `form:"diagnosa_terapi_wicara" json:"diagnosa_terapi_wicara"`
	DiagnosaMedis                        string `form:"diagnosa_medis" json:"diagnosa_medis"`
	Anamnesa                             string `form:"anamnesa" json:"anamnesa"`
	Suhu                                 string `form:"suhu" json:"suhu"`
	Rr                                   string `form:"rr" json:"rr"`
	Nadi                                 string `form:"nadi" json:"nadi"`
	Td                                   string `form:"td" json:"td"`
	PerilakuAdaptifKontakMata            string `form:"perilaku_adaptif_kontak_mata" json:"perilaku_adaptif_kontak_mata"`
	PerilakuAdaptifAtensi                string `form:"perilaku_adaptif_atensi" json:"perilaku_adaptif_atensi"`
	PerilakuAdaptifPerilaku              string `form:"perilaku_adaptif_perilaku" json:"perilaku_adaptif_perilaku"`
	KemampuanBahasaBicaraSpontan         string `form:"kemampuan_bahasa_bicara_spontan" json:"kemampuan_bahasa_bicara_spontan"`
	KemampuanBahasaPemahamanBahasa       string `form:"kemampuan_bahasa_pemahaman_bahasa" json:"kemampuan_bahasa_pemahaman_bahasa"`
	KemampuanBahasaPengujaran            string `form:"kemampuan_bahasa_pengujaran" json:"kemampuan_bahasa_pengujaran"`
	KemampuanBahasaMembaca               string `form:"kemampuan_bahasa_membaca" json:"kemampuan_bahasa_membaca"`
	KemampuanBahasaPenamaan              string `form:"kemampuan_bahasa_penamaan" json:"kemampuan_bahasa_penamaan"`
	OrganWicaraAnatomisLip               string `form:"organ_wicara_anatomis_lip" json:"organ_wicara_anatomis_lip"`
	OrganWicaraAnatomisTongue            string `form:"organ_wicara_anatomis_tongue" json:"organ_wicara_anatomis_tongue"`
	OrganWicaraAnatomisHardPalate        string `form:"organ_wicara_anatomis_hard_palate" json:"organ_wicara_anatomis_hard_palate"`
	OrganWicaraAnatomisSoftPalate        string `form:"organ_wicara_anatomis_soft_palate" json:"organ_wicara_anatomis_soft_palate"`
	OrganWicaraAnatomisUvula             string `form:"organ_wicara_anatomis_uvula" json:"organ_wicara_anatomis_uvula"`
	OrganWicaraAnatomisMandibula         string `form:"organ_wicara_anatomis_mandibula" json:"organ_wicara_anatomis_mandibula"`
	OrganWicaraAnatomisMaxila            string `form:"organ_wicara_anatomis_maxila" json:"organ_wicara_anatomis_maxila"`
	OrganWicaraAnatomisDental            string `form:"organ_wicara_anatomis_dental" json:"organ_wicara_anatomis_dental"`
	OrganWicaraAnatomisFaring            string `form:"organ_wicara_anatomis_faring" json:"organ_wicara_anatomis_faring"`
	OrganWicaraFisiologisLip             string `form:"organ_wicara_fisiologis_lip" json:"organ_wicara_fisiologis_lip"`
	OrganWicaraFisiologisTongue          string `form:"organ_wicara_fisiologis_tongue" json:"organ_wicara_fisiologis_tongue"`
	OrganWicaraFisiologisHardPalate      string `form:"organ_wicara_fisiologis_hard_palate" json:"organ_wicara_fisiologis_hard_palate"`
	OrganWicaraFisiologisSoftPalate      string `form:"organ_wicara_fisiologis_soft_palate" json:"organ_wicara_fisiologis_soft_palate"`
	OrganWicaraFisiologisUvula           string `form:"organ_wicara_fisiologis_uvula" json:"organ_wicara_fisiologis_uvula"`
	OrganWicaraFisiologisMandibula       string `form:"organ_wicara_fisiologis_mandibula" json:"organ_wicara_fisiologis_mandibula"`
	OrganWicaraFisiologisMaxilla         string `form:"organ_wicara_fisiologis_maxilla" json:"organ_wicara_fisiologis_maxilla"`
	OrganWicaraFisiologisDental          string `form:"organ_wicara_fisiologis_dental" json:"organ_wicara_fisiologis_dental"`
	OrganWicaraFisiologisFaring          string `form:"organ_wicara_fisiologis_faring" json:"organ_wicara_fisiologis_faring"`
	AktifitasOralMenghisap               string `form:"aktifitas_oral_menghisap" json:"aktifitas_oral_menghisap"`
	AktifitasOralMengunyah               string `form:"aktifitas_oral_mengunyah" json:"aktifitas_oral_mengunyah"`
	AktifitasOralMeniup                  string `form:"aktifitas_oral_meniup" json:"aktifitas_oral_meniup"`
	KemampuanArtikulasiSubtitusi         string `form:"kemampuan_artikulasi_subtitusi" json:"kemampuan_artikulasi_subtitusi"`
	KemampuanArtikulasiOmisi             string `form:"kemampuan_artikulasi_omisi" json:"kemampuan_artikulasi_omisi"`
	KemampuanArtikulasiDistorsi          string `form:"kemampuan_artikulasi_distorsi" json:"kemampuan_artikulasi_distorsi"`
	KemampuanArtikulasiAdisi             string `form:"kemampuan_artikulasi_adisi" json:"kemampuan_artikulasi_adisi"`
	Resonasi                             string `form:"resonasi" json:"resonasi"`
	KemampuanSuaraNada                   string `form:"kemampuan_suara_nada" json:"kemampuan_suara_nada"`
	KemampuanSuaraKualitas               string `form:"kemampuan_suara_kualitas" json:"kemampuan_suara_kualitas"`
	KemampuanSuaraKenyaringan            string `form:"kemampuan_suara_kenyaringan" json:"kemampuan_suara_kenyaringan"`
	KemampuanIramaKelancaran             string `form:"kemampuan_irama_kelancaran" json:"kemampuan_irama_kelancaran"`
	KemampuanMenelan                     string `form:"kemampuan_menelan" json:"kemampuan_menelan"`
	Pernafasan                           string `form:"pernafasan" json:"pernafasan"`
	TingkatKomunikasiDekodingPendengaran string `form:"tingkat_komunikasi_dekoding_pendengaran" json:"tingkat_komunikasi_dekoding_pendengaran"`
	TingkatKomunikasiDekodingPenglihatan string `form:"tingkat_komunikasi_dekoding_penglihatan" json:"tingkat_komunikasi_dekoding_penglihatan"`
	TingkatKomunikasiDekodingKinesik     string `form:"tingkat_komunikasi_dekoding_kinesik" json:"tingkat_komunikasi_dekoding_kinesik"`
	TingkatKomunikasiEnkodingBicara      string `form:"tingkat_komunikasi_enkoding_bicara" json:"tingkat_komunikasi_enkoding_bicara"`
	TingkatKomunikasiEnkodingTulisan     string `form:"tingkat_komunikasi_enkoding_tulisan" json:"tingkat_komunikasi_enkoding_tulisan"`
	TingkatKomunikasiEnkodingMimik       string `form:"tingkat_komunikasi_enkoding_mimik" json:"tingkat_komunikasi_enkoding_mimik"`
	TingkatKomunikasiEnkodingGesture     string `form:"tingkat_komunikasi_enkoding_gesture" json:"tingkat_komunikasi_enkoding_gesture"`
	PenunjangMedis                       string `form:"penunjang_medis" json:"penunjang_medis"`
	PerencanaanTerapiTujuan              string `form:"perencanaan_terapi_tujuan" json:"perencanaan_terapi_tujuan"`
	PerencanaanTerapiProgram             string `form:"perencanaan_terapi_program" json:"perencanaan_terapi_program"`
	Edukasi                              string `form:"edukasi" json:"edukasi"`
	TindakLanjut                         string `form:"tindak_lanjut" json:"tindak_lanjut"`
	Nip                                  string `form:"nip" json:"nip"`
}

func penilaianTerapiWicaraRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                 "required|date",
		"diagnosa_terapi_wicara":                  "string|max_len:100",
		"diagnosa_medis":                          "string|max_len:100",
		"anamnesa":                                "string|max_len:300",
		"suhu":                                    "string|max_len:5",
		"rr":                                      "string|max_len:5",
		"nadi":                                    "string|max_len:5",
		"td":                                      "string|max_len:8",
		"perilaku_adaptif_kontak_mata":            "string|max_len:50",
		"perilaku_adaptif_atensi":                 "string|max_len:50",
		"perilaku_adaptif_perilaku":               "string|max_len:50",
		"kemampuan_bahasa_bicara_spontan":         "string|max_len:50",
		"kemampuan_bahasa_pemahaman_bahasa":       "string|max_len:50",
		"kemampuan_bahasa_pengujaran":             "string|max_len:50",
		"kemampuan_bahasa_membaca":                "string|max_len:50",
		"kemampuan_bahasa_penamaan":               "string|max_len:50",
		"organ_wicara_anatomis_lip":               "string|max_len:30",
		"organ_wicara_anatomis_tongue":            "string|max_len:30",
		"organ_wicara_anatomis_hard_palate":       "string|max_len:30",
		"organ_wicara_anatomis_soft_palate":       "string|max_len:30",
		"organ_wicara_anatomis_uvula":             "string|max_len:30",
		"organ_wicara_anatomis_mandibula":         "string|max_len:30",
		"organ_wicara_anatomis_maxila":            "string|max_len:30",
		"organ_wicara_anatomis_dental":            "string|max_len:30",
		"organ_wicara_anatomis_faring":            "string|max_len:30",
		"organ_wicara_fisiologis_lip":             "string|max_len:30",
		"organ_wicara_fisiologis_tongue":          "string|max_len:30",
		"organ_wicara_fisiologis_hard_palate":     "string|max_len:30",
		"organ_wicara_fisiologis_soft_palate":     "string|max_len:30",
		"organ_wicara_fisiologis_uvula":           "string|max_len:30",
		"organ_wicara_fisiologis_mandibula":       "string|max_len:30",
		"organ_wicara_fisiologis_maxilla":         "string|max_len:30",
		"organ_wicara_fisiologis_dental":          "string|max_len:30",
		"organ_wicara_fisiologis_faring":          "string|max_len:30",
		"aktifitas_oral_menghisap":                "string|max_len:150",
		"aktifitas_oral_mengunyah":                "string|max_len:150",
		"aktifitas_oral_meniup":                   "string|max_len:150",
		"kemampuan_artikulasi_subtitusi":          "string|max_len:150",
		"kemampuan_artikulasi_omisi":              "string|max_len:150",
		"kemampuan_artikulasi_distorsi":           "string|max_len:150",
		"kemampuan_artikulasi_adisi":              "string|max_len:150",
		"resonasi":                                "required|in:Hiponasal,Hipernasal,Normal",
		"kemampuan_suara_nada":                    "required|in:Nada,Resonansi,Rendah,Monoton,Normal",
		"kemampuan_suara_kualitas":                "required|in:Hoarssness,Hassness,Normal",
		"kemampuan_suara_kenyaringan":             "required|in:Nyaring,Tidak Nyaring",
		"kemampuan_irama_kelancaran":              "required|in:Gagap Primer,Gagap Sekunder",
		"kemampuan_menelan":                       "string|max_len:150",
		"pernafasan":                              "string|max_len:150",
		"tingkat_komunikasi_dekoding_pendengaran": "string|max_len:30",
		"tingkat_komunikasi_dekoding_penglihatan": "string|max_len:30",
		"tingkat_komunikasi_dekoding_kinesik":     "string|max_len:30",
		"tingkat_komunikasi_enkoding_bicara":      "string|max_len:30",
		"tingkat_komunikasi_enkoding_tulisan":     "string|max_len:30",
		"tingkat_komunikasi_enkoding_mimik":       "string|max_len:30",
		"tingkat_komunikasi_enkoding_gesture":     "string|max_len:30",
		"penunjang_medis":                         "string|max_len:150",
		"perencanaan_terapi_tujuan":               "string|max_len:150",
		"perencanaan_terapi_program":              "string|max_len:150",
		"edukasi":                                 "string|max_len:150",
		"tindak_lanjut":                           "string|max_len:150",
		"nip":                                     "required|string|max_len:20",
	}
	return rules
}

// PenilaianTerapiWicaraStore simpan penilaian terapi wicara; kolom waktu kunci kosong = sekarang.
type PenilaianTerapiWicaraStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianTerapiWicaraData
}

func (r *PenilaianTerapiWicaraStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianTerapiWicaraStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianTerapiWicaraRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianTerapiWicaraStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianTerapiWicaraStore) Payload() PenilaianTerapiWicaraData {
	return r.PenilaianTerapiWicaraData
}

func (r *PenilaianTerapiWicaraStore) DetailValues() map[string][]string { return nil }

// PenilaianTerapiWicaraUpdate ubah penilaian terapi wicara (PUT); kunci lewat query string.
type PenilaianTerapiWicaraUpdate struct {
	PenilaianTerapiWicaraData
}

func (r *PenilaianTerapiWicaraUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianTerapiWicaraUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianTerapiWicaraRules()
}

func (r *PenilaianTerapiWicaraUpdate) Payload() PenilaianTerapiWicaraData {
	return r.PenilaianTerapiWicaraData
}

func (r *PenilaianTerapiWicaraUpdate) DetailValues() map[string][]string { return nil }
