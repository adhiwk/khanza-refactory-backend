package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilPemeriksaanEchoPediatrikData isian hasil pemeriksaan echo pediatrik.
type HasilPemeriksaanEchoPediatrikData struct {
	Tanggal                string `form:"tanggal" json:"tanggal"`
	KdDokter               string `form:"kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis         string `form:"diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari            string `form:"kiriman_dari" json:"kiriman_dari"`
	Situs                  string `form:"situs" json:"situs"`
	AvVa                   string `form:"av_va" json:"av_va"`
	DrainaseVenaPulmonalis string `form:"drainase_vena_pulmonalis" json:"drainase_vena_pulmonalis"`
	KatupMitral            string `form:"katup_mitral" json:"katup_mitral"`
	KatupAorta             string `form:"katup_aorta" json:"katup_aorta"`
	KatupTricuspid         string `form:"katup_tricuspid" json:"katup_tricuspid"`
	KatupPulmonal          string `form:"katup_pulmonal" json:"katup_pulmonal"`
	KatupSeptumAtrium      string `form:"katup_septum_atrium" json:"katup_septum_atrium"`
	KatupSeptumVentrikal   string `form:"katup_septum_ventrikal" json:"katup_septum_ventrikal"`
	KatupArkusAorta        string `form:"katup_arkus_aorta" json:"katup_arkus_aorta"`
	KatupKeteranganLainnya string `form:"katup_keterangan_lainnya" json:"katup_keterangan_lainnya"`
	RuangJantung           string `form:"ruang_jantung" json:"ruang_jantung"`
	ModeIvds               string `form:"mode_ivds" json:"mode_ivds"`
	ModeIvss               string `form:"mode_ivss" json:"mode_ivss"`
	ModeLvidDextra         string `form:"mode_lvid_dextra" json:"mode_lvid_dextra"`
	ModeLvidSinistra       string `form:"mode_lvid_sinistra" json:"mode_lvid_sinistra"`
	ModeLvpwDextra         string `form:"mode_lvpw_dextra" json:"mode_lvpw_dextra"`
	ModeLvpwSinistra       string `form:"mode_lvpw_sinistra" json:"mode_lvpw_sinistra"`
	ModeEjectionFraction   string `form:"mode_ejection_fraction" json:"mode_ejection_fraction"`
	ModeFractionShotening  string `form:"mode_fraction_shotening" json:"mode_fraction_shotening"`
	Doppler                string `form:"doppler" json:"doppler"`
	Kesimpulan             string `form:"kesimpulan" json:"kesimpulan"`
	Saran                  string `form:"saran" json:"saran"`
}

func hasilPemeriksaanEchoPediatrikRules() map[string]any {
	rules := map[string]any{
		"tanggal":                  "required|date",
		"kd_dokter":                "required|string|max_len:20",
		"diagnosa_klinis":          "string|max_len:50",
		"kiriman_dari":             "string|max_len:50",
		"situs":                    "string|max_len:100",
		"av_va":                    "string|max_len:100",
		"drainase_vena_pulmonalis": "string|max_len:100",
		"katup_mitral":             "string|max_len:100",
		"katup_aorta":              "string|max_len:100",
		"katup_tricuspid":          "string|max_len:100",
		"katup_pulmonal":           "string|max_len:100",
		"katup_septum_atrium":      "string|max_len:100",
		"katup_septum_ventrikal":   "string|max_len:100",
		"katup_arkus_aorta":        "string|max_len:100",
		"katup_keterangan_lainnya": "string|max_len:100",
		"ruang_jantung":            "string|max_len:100",
		"mode_ivds":                "string|max_len:20",
		"mode_ivss":                "string|max_len:20",
		"mode_lvid_dextra":         "string|max_len:20",
		"mode_lvid_sinistra":       "string|max_len:20",
		"mode_lvpw_dextra":         "string|max_len:20",
		"mode_lvpw_sinistra":       "string|max_len:20",
		"mode_ejection_fraction":   "string|max_len:20",
		"mode_fraction_shotening":  "string|max_len:20",
		"doppler":                  "string|max_len:100",
		"kesimpulan":               "string|max_len:250",
		"saran":                    "string|max_len:100",
	}
	return rules
}

// HasilPemeriksaanEchoPediatrikStore simpan hasil pemeriksaan echo pediatrik; kolom waktu kunci kosong = sekarang.
type HasilPemeriksaanEchoPediatrikStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	HasilPemeriksaanEchoPediatrikData
}

func (r *HasilPemeriksaanEchoPediatrikStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanEchoPediatrikStore) Rules(ctx http.Context) map[string]any {
	rules := hasilPemeriksaanEchoPediatrikRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *HasilPemeriksaanEchoPediatrikStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *HasilPemeriksaanEchoPediatrikStore) Payload() HasilPemeriksaanEchoPediatrikData {
	return r.HasilPemeriksaanEchoPediatrikData
}

func (r *HasilPemeriksaanEchoPediatrikStore) DetailValues() map[string][]string { return nil }

// HasilPemeriksaanEchoPediatrikUpdate ubah hasil pemeriksaan echo pediatrik (PUT); kunci lewat query string.
type HasilPemeriksaanEchoPediatrikUpdate struct {
	HasilPemeriksaanEchoPediatrikData
}

func (r *HasilPemeriksaanEchoPediatrikUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanEchoPediatrikUpdate) Rules(ctx http.Context) map[string]any {
	return hasilPemeriksaanEchoPediatrikRules()
}

func (r *HasilPemeriksaanEchoPediatrikUpdate) Payload() HasilPemeriksaanEchoPediatrikData {
	return r.HasilPemeriksaanEchoPediatrikData
}

func (r *HasilPemeriksaanEchoPediatrikUpdate) DetailValues() map[string][]string { return nil }
