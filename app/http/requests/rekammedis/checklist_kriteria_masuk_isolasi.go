package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ChecklistKriteriaMasukIsolasiData isian checklist kriteria masuk isolasi.
type ChecklistKriteriaMasukIsolasiData struct {
	AirborneTb               string `form:"airborne_tb" json:"airborne_tb"`
	AirborneCampak           string `form:"airborne_campak" json:"airborne_campak"`
	AirborneVarisela         string `form:"airborne_varisela" json:"airborne_varisela"`
	AirborneZosterDiseminata string `form:"airborne_zoster_diseminata" json:"airborne_zoster_diseminata"`
	AirborneLainnya          string `form:"airborne_lainnya" json:"airborne_lainnya"`
	DropletCovid19           string `form:"droplet_covid19" json:"droplet_covid19"`
	DropletInfluenza         string `form:"droplet_influenza" json:"droplet_influenza"`
	DropletDifteri           string `form:"droplet_difteri" json:"droplet_difteri"`
	DropletPertusis          string `form:"droplet_pertusis" json:"droplet_pertusis"`
	DropletMeningitis        string `form:"droplet_meningitis" json:"droplet_meningitis"`
	DropletLainnya           string `form:"droplet_lainnya" json:"droplet_lainnya"`
	KontakMdro               string `form:"kontak_mdro" json:"kontak_mdro"`
	KontakClostridium        string `form:"kontak_clostridium" json:"kontak_clostridium"`
	KontakScabies            string `form:"kontak_scabies" json:"kontak_scabies"`
	KontakLukaDrainase       string `form:"kontak_luka_drainase" json:"kontak_luka_drainase"`
	KontakDiareInfeksi       string `form:"kontak_diare_infeksi" json:"kontak_diare_infeksi"`
	KontakLainnya            string `form:"kontak_lainnya" json:"kontak_lainnya"`
	KontakErat               string `form:"kontak_erat" json:"kontak_erat"`
	RiwayatPerjalananWabah   string `form:"riwayat_perjalanan_wabah" json:"riwayat_perjalanan_wabah"`
	RiwayatMdro              string `form:"riwayat_mdro" json:"riwayat_mdro"`
	HasilLabRadiologiPositif string `form:"hasil_lab_radiologi_positif" json:"hasil_lab_radiologi_positif"`
	GejalaKlinisMenular      string `form:"gejala_klinis_menular" json:"gejala_klinis_menular"`
	PasienImunokompromis     string `form:"pasien_imunokompromis" json:"pasien_imunokompromis"`
	InstruksiDpjp            string `form:"instruksi_dpjp" json:"instruksi_dpjp"`
	PersetujuanIsolasi       string `form:"persetujuan_isolasi" json:"persetujuan_isolasi"`
	KelengkapanJaminan       string `form:"kelengkapan_jaminan" json:"kelengkapan_jaminan"`
	KetersediaanApd          string `form:"ketersediaan_apd" json:"ketersediaan_apd"`
	FasilitasCuciTangan      string `form:"fasilitas_cuci_tangan" json:"fasilitas_cuci_tangan"`
	TekananNegatifBerfungsi  string `form:"tekanan_negatif_berfungsi" json:"tekanan_negatif_berfungsi"`
	PintuOtomatisBerfungsi   string `form:"pintu_otomatis_berfungsi" json:"pintu_otomatis_berfungsi"`
	IndikasiIsolasi          string `form:"indikasi_isolasi" json:"indikasi_isolasi"`
	JenisIsolasi             string `form:"jenis_isolasi" json:"jenis_isolasi"`
	DiagnosaIsolasi          string `form:"diagnosa_isolasi" json:"diagnosa_isolasi"`
	Keterangan               string `form:"keterangan" json:"keterangan"`
	Nik                      string `form:"nik" json:"nik"`
}

func checklistKriteriaMasukIsolasiRules() map[string]any {
	rules := map[string]any{
		"airborne_tb":                 "required|in:Ya,Tidak",
		"airborne_campak":             "required|in:Ya,Tidak",
		"airborne_varisela":           "required|in:Ya,Tidak",
		"airborne_zoster_diseminata":  "required|in:Ya,Tidak",
		"airborne_lainnya":            "required|in:Ya,Tidak",
		"droplet_covid19":             "required|in:Ya,Tidak",
		"droplet_influenza":           "required|in:Ya,Tidak",
		"droplet_difteri":             "required|in:Ya,Tidak",
		"droplet_pertusis":            "required|in:Ya,Tidak",
		"droplet_meningitis":          "required|in:Ya,Tidak",
		"droplet_lainnya":             "required|in:Ya,Tidak",
		"kontak_mdro":                 "required|in:Ya,Tidak",
		"kontak_clostridium":          "required|in:Ya,Tidak",
		"kontak_scabies":              "required|in:Ya,Tidak",
		"kontak_luka_drainase":        "required|in:Ya,Tidak",
		"kontak_diare_infeksi":        "required|in:Ya,Tidak",
		"kontak_lainnya":              "required|in:Ya,Tidak",
		"kontak_erat":                 "required|in:Ya,Tidak",
		"riwayat_perjalanan_wabah":    "required|in:Ya,Tidak",
		"riwayat_mdro":                "required|in:Ya,Tidak",
		"hasil_lab_radiologi_positif": "required|in:Ya,Tidak",
		"gejala_klinis_menular":       "required|in:Ya,Tidak",
		"pasien_imunokompromis":       "required|in:Ya,Tidak",
		"instruksi_dpjp":              "required|in:Ya,Tidak",
		"persetujuan_isolasi":         "required|in:Ya,Tidak",
		"kelengkapan_jaminan":         "required|in:Ya,Tidak",
		"ketersediaan_apd":            "required|in:Ya,Tidak",
		"fasilitas_cuci_tangan":       "required|in:Ya,Tidak",
		"tekanan_negatif_berfungsi":   "required|in:Ya,Tidak",
		"pintu_otomatis_berfungsi":    "required|in:Ya,Tidak",
		"indikasi_isolasi":            "required|in:Masuk Di Ruangan Isolasi,Tetap Di Ruangan Biasa",
		"jenis_isolasi":               "required|in:Standar,Kontak,Droplet,Airborne,Kontak+Droplet,Airborne+Kontak,Protective",
		"diagnosa_isolasi":            "string|max_len:100",
		"keterangan":                  "string|max_len:100",
		"nik":                         "string|max_len:20",
	}
	return rules
}

// ChecklistKriteriaMasukIsolasiStore simpan checklist kriteria masuk isolasi; kolom waktu kunci kosong = sekarang.
type ChecklistKriteriaMasukIsolasiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	ChecklistKriteriaMasukIsolasiData
}

func (r *ChecklistKriteriaMasukIsolasiStore) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaMasukIsolasiStore) Rules(ctx http.Context) map[string]any {
	rules := checklistKriteriaMasukIsolasiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *ChecklistKriteriaMasukIsolasiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *ChecklistKriteriaMasukIsolasiStore) Payload() ChecklistKriteriaMasukIsolasiData {
	return r.ChecklistKriteriaMasukIsolasiData
}

func (r *ChecklistKriteriaMasukIsolasiStore) DetailValues() map[string][]string { return nil }

// ChecklistKriteriaMasukIsolasiUpdate ubah checklist kriteria masuk isolasi (PUT); kunci lewat query string.
type ChecklistKriteriaMasukIsolasiUpdate struct {
	ChecklistKriteriaMasukIsolasiData
}

func (r *ChecklistKriteriaMasukIsolasiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistKriteriaMasukIsolasiUpdate) Rules(ctx http.Context) map[string]any {
	return checklistKriteriaMasukIsolasiRules()
}

func (r *ChecklistKriteriaMasukIsolasiUpdate) Payload() ChecklistKriteriaMasukIsolasiData {
	return r.ChecklistKriteriaMasukIsolasiData
}

func (r *ChecklistKriteriaMasukIsolasiUpdate) DetailValues() map[string][]string { return nil }
