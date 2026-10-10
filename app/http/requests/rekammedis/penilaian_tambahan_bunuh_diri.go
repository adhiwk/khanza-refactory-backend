package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianTambahanBunuhDiriData isian penilaian tambahan bunuh diri.
type PenilaianTambahanBunuhDiriData struct {
	Tanggal                        string `form:"tanggal" json:"tanggal"`
	Nip                            string `form:"nip" json:"nip"`
	StatikHidupSendiri             string `form:"statik_hidup_sendiri" json:"statik_hidup_sendiri"`
	StatikSkorhidupSendiri         *int   `form:"statik_skorhidup_sendiri" json:"statik_skorhidup_sendiri"`
	StatikUpayaSuicide             string `form:"statik_upaya_suicide" json:"statik_upaya_suicide"`
	StatikSkorupayaSuicide         *int   `form:"statik_skorupaya_suicide" json:"statik_skorupaya_suicide"`
	StatikKeluargaSuicide          string `form:"statik_keluarga_suicide" json:"statik_keluarga_suicide"`
	StatikSkorkeluargaSuicide      *int   `form:"statik_skorkeluarga_suicide" json:"statik_skorkeluarga_suicide"`
	StatikDiagnosaGangguanJiwa     string `form:"statik_diagnosa_gangguan_jiwa" json:"statik_diagnosa_gangguan_jiwa"`
	StatikSkordiagnosaGangguanJiwa *int   `form:"statik_skordiagnosa_gangguan_jiwa" json:"statik_skordiagnosa_gangguan_jiwa"`
	StatikDisabilitasBerat         string `form:"statik_disabilitas_berat" json:"statik_disabilitas_berat"`
	StatikSkordisabilitasBerat     *int   `form:"statik_skordisabilitas_berat" json:"statik_skordisabilitas_berat"`
	StatikBerpisah                 string `form:"statik_berpisah" json:"statik_berpisah"`
	StatikSkorberpisah             *int   `form:"statik_skorberpisah" json:"statik_skorberpisah"`
	StatikKehilanganKerja          string `form:"statik_kehilangan_kerja" json:"statik_kehilangan_kerja"`
	StatikSkorkehilanganKerja      *int   `form:"statik_skorkehilangan_kerja" json:"statik_skorkehilangan_kerja"`
	StatikSkortotal                *int   `form:"statik_skortotal" json:"statik_skortotal"`
	DinamisIdeBunuhDiri            string `form:"dinamis_ide_bunuh_diri" json:"dinamis_ide_bunuh_diri"`
	DinamisSkorideBunuhDiri        *int   `form:"dinamis_skoride_bunuh_diri" json:"dinamis_skoride_bunuh_diri"`
	DinamisMaksudSuicide           string `form:"dinamis_maksud_suicide" json:"dinamis_maksud_suicide"`
	DinamisSkormaksudSuicide       *int   `form:"dinamis_skormaksud_suicide" json:"dinamis_skormaksud_suicide"`
	DinamisStressBerat             string `form:"dinamis_stress_berat" json:"dinamis_stress_berat"`
	DinamisSkorstressBerat         *int   `form:"dinamis_skorstress_berat" json:"dinamis_skorstress_berat"`
	DinamisKeputusasaan            string `form:"dinamis_keputusasaan" json:"dinamis_keputusasaan"`
	DinamisSkorkeputusasaan        *int   `form:"dinamis_skorkeputusasaan" json:"dinamis_skorkeputusasaan"`
	DinamisKejadianSignifikan      string `form:"dinamis_kejadian_signifikan" json:"dinamis_kejadian_signifikan"`
	DinamisSkorkejadianSignifikan  *int   `form:"dinamis_skorkejadian_signifikan" json:"dinamis_skorkejadian_signifikan"`
	DinamisKehilanganKontrol       string `form:"dinamis_kehilangan_kontrol" json:"dinamis_kehilangan_kontrol"`
	DinamisSkorkehilanganKontrol   *int   `form:"dinamis_skorkehilangan_kontrol" json:"dinamis_skorkehilangan_kontrol"`
	DinamisPenggunaanNapza         string `form:"dinamis_penggunaan_napza" json:"dinamis_penggunaan_napza"`
	DinamisSkorpenggunaanNapza     *int   `form:"dinamis_skorpenggunaan_napza" json:"dinamis_skorpenggunaan_napza"`
	DinamisSkortotal               *int   `form:"dinamis_skortotal" json:"dinamis_skortotal"`
	FaktorFaktorPencegahan         string `form:"faktor_faktor_pencegahan" json:"faktor_faktor_pencegahan"`
	TotalSkor                      *int   `form:"total_skor" json:"total_skor"`
	LevelSkor                      string `form:"level_skor" json:"level_skor"`
}

func penilaianTambahanBunuhDiriRules() map[string]any {
	rules := map[string]any{
		"tanggal":                           "required|date",
		"nip":                               "required|string|max_len:20",
		"statik_hidup_sendiri":              "in:Ya,Tidak,Tidak Tahu",
		"statik_skorhidup_sendiri":          "int",
		"statik_upaya_suicide":              "in:Ya,Tidak,Tidak Tahu",
		"statik_skorupaya_suicide":          "int",
		"statik_keluarga_suicide":           "in:Ya,Tidak,Tidak Tahu",
		"statik_skorkeluarga_suicide":       "int",
		"statik_diagnosa_gangguan_jiwa":     "in:Ya,Tidak,Tidak Tahu",
		"statik_skordiagnosa_gangguan_jiwa": "int",
		"statik_disabilitas_berat":          "in:Ya,Tidak,Tidak Tahu",
		"statik_skordisabilitas_berat":      "int",
		"statik_berpisah":                   "in:Ya,Tidak,Tidak Tahu",
		"statik_skorberpisah":               "int",
		"statik_kehilangan_kerja":           "in:Ya,Tidak,Tidak Tahu",
		"statik_skorkehilangan_kerja":       "int",
		"statik_skortotal":                  "int",
		"dinamis_ide_bunuh_diri":            "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skoride_bunuh_diri":        "int",
		"dinamis_maksud_suicide":            "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skormaksud_suicide":        "int",
		"dinamis_stress_berat":              "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorstress_berat":          "int",
		"dinamis_keputusasaan":              "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorkeputusasaan":          "int",
		"dinamis_kejadian_signifikan":       "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorkejadian_signifikan":   "int",
		"dinamis_kehilangan_kontrol":        "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorkehilangan_kontrol":    "int",
		"dinamis_penggunaan_napza":          "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorpenggunaan_napza":      "int",
		"dinamis_skortotal":                 "int",
		"faktor_faktor_pencegahan":          "string|max_len:500",
		"total_skor":                        "int",
		"level_skor":                        "in:Rendah(<7),Sedang(7-14),Tinggi(>14)",
	}
	return rules
}

// PenilaianTambahanBunuhDiriStore simpan penilaian tambahan bunuh diri; kolom waktu kunci kosong = sekarang.
type PenilaianTambahanBunuhDiriStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianTambahanBunuhDiriData
}

func (r *PenilaianTambahanBunuhDiriStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianTambahanBunuhDiriStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianTambahanBunuhDiriRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianTambahanBunuhDiriStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianTambahanBunuhDiriStore) Payload() PenilaianTambahanBunuhDiriData {
	return r.PenilaianTambahanBunuhDiriData
}

func (r *PenilaianTambahanBunuhDiriStore) DetailValues() map[string][]string { return nil }

// PenilaianTambahanBunuhDiriUpdate ubah penilaian tambahan bunuh diri (PUT); kunci lewat query string.
type PenilaianTambahanBunuhDiriUpdate struct {
	PenilaianTambahanBunuhDiriData
}

func (r *PenilaianTambahanBunuhDiriUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianTambahanBunuhDiriUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianTambahanBunuhDiriRules()
}

func (r *PenilaianTambahanBunuhDiriUpdate) Payload() PenilaianTambahanBunuhDiriData {
	return r.PenilaianTambahanBunuhDiriData
}

func (r *PenilaianTambahanBunuhDiriUpdate) DetailValues() map[string][]string { return nil }
