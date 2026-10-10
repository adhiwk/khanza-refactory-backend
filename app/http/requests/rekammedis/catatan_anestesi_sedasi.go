package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanAnestesiSedasiData isian catatan anastesi sedasi.
type CatatanAnestesiSedasiData struct {
	KdDokterBedah                           string `form:"kd_dokter_bedah" json:"kd_dokter_bedah"`
	KdDokterAnestesi                        string `form:"kd_dokter_anestesi" json:"kd_dokter_anestesi"`
	DiagnosaPreBedah                        string `form:"diagnosa_pre_bedah" json:"diagnosa_pre_bedah"`
	TindakanJenisPembedahan                 string `form:"tindakan_jenis_pembedahan" json:"tindakan_jenis_pembedahan"`
	DiagnosaPascaBedah                      string `form:"diagnosa_pasca_bedah" json:"diagnosa_pasca_bedah"`
	PreInduksiJam                           string `form:"pre_induksi_jam" json:"pre_induksi_jam"`
	PreInduksiKesadaran                     string `form:"pre_induksi_kesadaran" json:"pre_induksi_kesadaran"`
	PreInduksiTd                            string `form:"pre_induksi_td" json:"pre_induksi_td"`
	PreInduksiNadi                          string `form:"pre_induksi_nadi" json:"pre_induksi_nadi"`
	PreInduksiRr                            string `form:"pre_induksi_rr" json:"pre_induksi_rr"`
	PreInduksiSuhu                          string `form:"pre_induksi_suhu" json:"pre_induksi_suhu"`
	PreInduksiO2                            string `form:"pre_induksi_o2" json:"pre_induksi_o2"`
	PreInduksiTb                            string `form:"pre_induksi_tb" json:"pre_induksi_tb"`
	PreInduksiBb                            string `form:"pre_induksi_bb" json:"pre_induksi_bb"`
	PreInduksiRhesus                        string `form:"pre_induksi_rhesus" json:"pre_induksi_rhesus"`
	PreInduksiHb                            string `form:"pre_induksi_hb" json:"pre_induksi_hb"`
	PreInduksiHt                            string `form:"pre_induksi_ht" json:"pre_induksi_ht"`
	PreInduksiLeko                          string `form:"pre_induksi_leko" json:"pre_induksi_leko"`
	PreInduksiTrombo                        string `form:"pre_induksi_trombo" json:"pre_induksi_trombo"`
	PreInduksiBtct                          string `form:"pre_induksi_btct" json:"pre_induksi_btct"`
	PreInduksiGds                           string `form:"pre_induksi_gds" json:"pre_induksi_gds"`
	PreInduksiLainlain                      string `form:"pre_induksi_lainlain" json:"pre_induksi_lainlain"`
	TeknikAlatHiopotensi                    string `form:"teknik_alat_hiopotensi" json:"teknik_alat_hiopotensi"`
	TeknikAlatTci                           string `form:"teknik_alat_tci" json:"teknik_alat_tci"`
	TeknikAlatCpb                           string `form:"teknik_alat_cpb" json:"teknik_alat_cpb"`
	TeknikAlatVentilasi                     string `form:"teknik_alat_ventilasi" json:"teknik_alat_ventilasi"`
	TeknikAlatBroncoskopy                   string `form:"teknik_alat_broncoskopy" json:"teknik_alat_broncoskopy"`
	TeknikAlatGlidescopi                    string `form:"teknik_alat_glidescopi" json:"teknik_alat_glidescopi"`
	TeknikAlatUsg                           string `form:"teknik_alat_usg" json:"teknik_alat_usg"`
	TeknikAlatStimulatorSaraf               string `form:"teknik_alat_stimulator_saraf" json:"teknik_alat_stimulator_saraf"`
	TeknikAlatLainlain                      string `form:"teknik_alat_lainlain" json:"teknik_alat_lainlain"`
	MonitoringEkg                           string `form:"monitoring_ekg" json:"monitoring_ekg"`
	MonitoringEkgKeterangan                 string `form:"monitoring_ekg_keterangan" json:"monitoring_ekg_keterangan"`
	MonitoringArteri                        string `form:"monitoring_arteri" json:"monitoring_arteri"`
	MonitoringArteriKeterangan              string `form:"monitoring_arteri_keterangan" json:"monitoring_arteri_keterangan"`
	MonitoringCvp                           string `form:"monitoring_cvp" json:"monitoring_cvp"`
	MonitoringCvpKeterangan                 string `form:"monitoring_cvp_keterangan" json:"monitoring_cvp_keterangan"`
	MonitoringEtco                          string `form:"monitoring_etco" json:"monitoring_etco"`
	MonitoringStetoskop                     string `form:"monitoring_stetoskop" json:"monitoring_stetoskop"`
	MonitoringNibp                          string `form:"monitoring_nibp" json:"monitoring_nibp"`
	MonitoringNgt                           string `form:"monitoring_ngt" json:"monitoring_ngt"`
	MonitoringBis                           string `form:"monitoring_bis" json:"monitoring_bis"`
	MonitoringCathAPulmo                    string `form:"monitoring_cath_a_pulmo" json:"monitoring_cath_a_pulmo"`
	MonitoringSpo2                          string `form:"monitoring_spo2" json:"monitoring_spo2"`
	MonitoringKateter                       string `form:"monitoring_kateter" json:"monitoring_kateter"`
	MonitoringTemp                          string `form:"monitoring_temp" json:"monitoring_temp"`
	MonitoringLainlain                      string `form:"monitoring_lainlain" json:"monitoring_lainlain"`
	StatusFisikAsa                          string `form:"status_fisik_asa" json:"status_fisik_asa"`
	StatusFisikAlergi                       string `form:"status_fisik_alergi" json:"status_fisik_alergi"`
	StatusFisikAlergiKeterangan             string `form:"status_fisik_alergi_keterangan" json:"status_fisik_alergi_keterangan"`
	StatusFisikPenyulitSedasi               string `form:"status_fisik_penyulit_sedasi" json:"status_fisik_penyulit_sedasi"`
	PerencanaanLanjut                       string `form:"perencanaan_lanjut" json:"perencanaan_lanjut"`
	PerencanaanLanjutSedasi                 string `form:"perencanaan_lanjut_sedasi" json:"perencanaan_lanjut_sedasi"`
	PerencanaanLanjutSedasiKeterangan       string `form:"perencanaan_lanjut_sedasi_keterangan" json:"perencanaan_lanjut_sedasi_keterangan"`
	PerencanaanLanjutSpinal                 string `form:"perencanaan_lanjut_spinal" json:"perencanaan_lanjut_spinal"`
	PerencanaanLanjutAnestesiUmum           string `form:"perencanaan_lanjut_anestesi_umum" json:"perencanaan_lanjut_anestesi_umum"`
	PerencanaanLanjutAnestesiUmumKeterangan string `form:"perencanaan_lanjut_anestesi_umum_keterangan" json:"perencanaan_lanjut_anestesi_umum_keterangan"`
	PerencanaanLanjutBlokPerifer            string `form:"perencanaan_lanjut_blok_perifer" json:"perencanaan_lanjut_blok_perifer"`
	PerencanaanLanjutBlokPeriferKeterangan  string `form:"perencanaan_lanjut_blok_perifer_keterangan" json:"perencanaan_lanjut_blok_perifer_keterangan"`
	PerencanaanLanjutEpidural               string `form:"perencanaan_lanjut_epidural" json:"perencanaan_lanjut_epidural"`
	PerencanaanBatal                        string `form:"perencanaan_batal" json:"perencanaan_batal"`
	PerencanaanBatalAlasan                  string `form:"perencanaan_batal_alasan" json:"perencanaan_batal_alasan"`
	NipPerawatOk                            string `form:"nip_perawat_ok" json:"nip_perawat_ok"`
	NipPerawatAnestesi                      string `form:"nip_perawat_anestesi" json:"nip_perawat_anestesi"`
}

func catatanAnestesiSedasiRules() map[string]any {
	rules := map[string]any{
		"kd_dokter_bedah":                             "required|string|max_len:20",
		"kd_dokter_anestesi":                          "required|string|max_len:20",
		"diagnosa_pre_bedah":                          "string|max_len:50",
		"tindakan_jenis_pembedahan":                   "string|max_len:50",
		"diagnosa_pasca_bedah":                        "string|max_len:50",
		"pre_induksi_jam":                             "string|max_len:10",
		"pre_induksi_kesadaran":                       "in:Compos Mentis,Somnolence,Sopor,Coma,Alert,Confusion,Voice,Pain,Unresponsive,Apatis,Delirium",
		"pre_induksi_td":                              "string|max_len:8",
		"pre_induksi_nadi":                            "string|max_len:5",
		"pre_induksi_rr":                              "string|max_len:5",
		"pre_induksi_suhu":                            "string|max_len:5",
		"pre_induksi_o2":                              "string|max_len:5",
		"pre_induksi_tb":                              "string|max_len:5",
		"pre_induksi_bb":                              "string|max_len:5",
		"pre_induksi_rhesus":                          "in:+,-",
		"pre_induksi_hb":                              "string|max_len:5",
		"pre_induksi_ht":                              "string|max_len:5",
		"pre_induksi_leko":                            "string|max_len:5",
		"pre_induksi_trombo":                          "string|max_len:5",
		"pre_induksi_btct":                            "string|max_len:5",
		"pre_induksi_gds":                             "string|max_len:5",
		"pre_induksi_lainlain":                        "string|max_len:30",
		"teknik_alat_hiopotensi":                      "in:Ya,Tidak",
		"teknik_alat_tci":                             "in:Ya,Tidak",
		"teknik_alat_cpb":                             "in:Ya,Tidak",
		"teknik_alat_ventilasi":                       "in:Ya,Tidak",
		"teknik_alat_broncoskopy":                     "in:Ya,Tidak",
		"teknik_alat_glidescopi":                      "in:Ya,Tidak",
		"teknik_alat_usg":                             "in:Ya,Tidak",
		"teknik_alat_stimulator_saraf":                "in:Ya,Tidak",
		"teknik_alat_lainlain":                        "string|max_len:100",
		"monitoring_ekg":                              "in:Ya,Tidak",
		"monitoring_ekg_keterangan":                   "string|max_len:50",
		"monitoring_arteri":                           "in:Ya,Tidak",
		"monitoring_arteri_keterangan":                "string|max_len:50",
		"monitoring_cvp":                              "in:Ya,Tidak",
		"monitoring_cvp_keterangan":                   "string|max_len:50",
		"monitoring_etco":                             "in:Ya,Tidak",
		"monitoring_stetoskop":                        "in:Ya,Tidak",
		"monitoring_nibp":                             "in:Ya,Tidak",
		"monitoring_ngt":                              "in:Ya,Tidak",
		"monitoring_bis":                              "in:Ya,Tidak",
		"monitoring_cath_a_pulmo":                     "in:Ya,Tidak",
		"monitoring_spo2":                             "in:Ya,Tidak",
		"monitoring_kateter":                          "in:Ya,Tidak",
		"monitoring_temp":                             "in:Ya,Tidak",
		"monitoring_lainlain":                         "string|max_len:100",
		"status_fisik_asa":                            "in:1,2,3,4,5,E",
		"status_fisik_alergi":                         "required|in:Tidak,Ya",
		"status_fisik_alergi_keterangan":              "string|max_len:50",
		"status_fisik_penyulit_sedasi":                "string|max_len:150",
		"perencanaan_lanjut":                          "in:Ya,Tidak",
		"perencanaan_lanjut_sedasi":                   "in:Sedang,Dalam,Tidak,Lain-lain",
		"perencanaan_lanjut_sedasi_keterangan":        "string|max_len:30",
		"perencanaan_lanjut_spinal":                   "in:Ya,Tidak",
		"perencanaan_lanjut_anestesi_umum":            "in:Ya,Tidak",
		"perencanaan_lanjut_anestesi_umum_keterangan": "string|max_len:30",
		"perencanaan_lanjut_blok_perifer":             "in:Ya,Tidak",
		"perencanaan_lanjut_blok_perifer_keterangan":  "string|max_len:30",
		"perencanaan_lanjut_epidural":                 "in:Ya,Tidak",
		"perencanaan_batal":                           "in:Ya,Tidak",
		"perencanaan_batal_alasan":                    "string|max_len:150",
		"nip_perawat_ok":                              "string|max_len:20",
		"nip_perawat_anestesi":                        "string|max_len:20",
	}
	return rules
}

// CatatanAnestesiSedasiStore simpan catatan anastesi sedasi; kolom waktu kunci kosong = sekarang.
type CatatanAnestesiSedasiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	CatatanAnestesiSedasiData
}

func (r *CatatanAnestesiSedasiStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanAnestesiSedasiStore) Rules(ctx http.Context) map[string]any {
	rules := catatanAnestesiSedasiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *CatatanAnestesiSedasiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *CatatanAnestesiSedasiStore) Payload() CatatanAnestesiSedasiData {
	return r.CatatanAnestesiSedasiData
}

func (r *CatatanAnestesiSedasiStore) DetailValues() map[string][]string { return nil }

// CatatanAnestesiSedasiUpdate ubah catatan anastesi sedasi (PUT); kunci lewat query string.
type CatatanAnestesiSedasiUpdate struct {
	CatatanAnestesiSedasiData
}

func (r *CatatanAnestesiSedasiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanAnestesiSedasiUpdate) Rules(ctx http.Context) map[string]any {
	return catatanAnestesiSedasiRules()
}

func (r *CatatanAnestesiSedasiUpdate) Payload() CatatanAnestesiSedasiData {
	return r.CatatanAnestesiSedasiData
}

func (r *CatatanAnestesiSedasiUpdate) DetailValues() map[string][]string { return nil }
