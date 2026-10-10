package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianTambahanBeresikoMelarikanDiriData isian penilaian tambahan melarikan diri.
type PenilaianTambahanBeresikoMelarikanDiriData struct {
	Tanggal                                string `form:"tanggal" json:"tanggal"`
	Nip                                    string `form:"nip" json:"nip"`
	StatikRiwayatMelarikanDiri             string `form:"statik_riwayat_melarikan_diri" json:"statik_riwayat_melarikan_diri"`
	StatikSkorriwayatMelarikanDiri         *int   `form:"statik_skorriwayat_melarikan_diri" json:"statik_skorriwayat_melarikan_diri"`
	StatikRiwayatPenolakanPengobatan       string `form:"statik_riwayat_penolakan_pengobatan" json:"statik_riwayat_penolakan_pengobatan"`
	StatikSkorriwayatPenolakanPengobatan   *int   `form:"statik_skorriwayat_penolakan_pengobatan" json:"statik_skorriwayat_penolakan_pengobatan"`
	StatikUsiaDibawah35                    string `form:"statik_usia_dibawah_35" json:"statik_usia_dibawah_35"`
	StatikSkorusiaDibawah35                *int   `form:"statik_skorusia_dibawah_35" json:"statik_skorusia_dibawah_35"`
	StatikLakiLaki                         string `form:"statik_laki_laki" json:"statik_laki_laki"`
	StatikSkorlakiLaki                     *int   `form:"statik_skorlaki_laki" json:"statik_skorlaki_laki"`
	StatikDiagnosisSkizofrenia             string `form:"statik_diagnosis_skizofrenia" json:"statik_diagnosis_skizofrenia"`
	StatikSkordiagnosisSkizofrenia         *int   `form:"statik_skordiagnosis_skizofrenia" json:"statik_skordiagnosis_skizofrenia"`
	StatikBelumMenikah                     string `form:"statik_belum_menikah" json:"statik_belum_menikah"`
	StatikSkorbelumMenikah                 *int   `form:"statik_skorbelum_menikah" json:"statik_skorbelum_menikah"`
	StatikRiwayatPenggunaanNapza           string `form:"statik_riwayat_penggunaan_napza" json:"statik_riwayat_penggunaan_napza"`
	StatikSkoriwayatPenggunaanNapza        *int   `form:"statik_skoriwayat_penggunaan_napza" json:"statik_skoriwayat_penggunaan_napza"`
	StatikDiagnosisGangguanKepribadian     string `form:"statik_diagnosis_gangguan_kepribadian" json:"statik_diagnosis_gangguan_kepribadian"`
	StatikSkordiagnosisGangguanKepribadian *int   `form:"statik_skordiagnosis_gangguan_kepribadian" json:"statik_skordiagnosis_gangguan_kepribadian"`
	StatikRiwayatKriminal                  string `form:"statik_riwayat_kriminal" json:"statik_riwayat_kriminal"`
	StatikSkorriwayatKriminal              *int   `form:"statik_skorriwayat_kriminal" json:"statik_skorriwayat_kriminal"`
	StatikSkortotal                        *int   `form:"statik_skortotal" json:"statik_skortotal"`
	DinamisAntiTreatment                   string `form:"dinamis_anti_treatment" json:"dinamis_anti_treatment"`
	DinamisSkorantiTreatment               *int   `form:"dinamis_skoranti_treatment" json:"dinamis_skoranti_treatment"`
	DinamisPenggunaanNapza                 string `form:"dinamis_penggunaan_napza" json:"dinamis_penggunaan_napza"`
	DinamisSkorpenggunaanNapza             *int   `form:"dinamis_skorpenggunaan_napza" json:"dinamis_skorpenggunaan_napza"`
	DinamisKebosanan                       string `form:"dinamis_kebosanan" json:"dinamis_kebosanan"`
	DinamisSkorkebosanan                   *int   `form:"dinamis_skorkebosanan" json:"dinamis_skorkebosanan"`
	DinamisPerintahHalusinasi              string `form:"dinamis_perintah_halusinasi" json:"dinamis_perintah_halusinasi"`
	DinamisSkorperintahHalusinasi          *int   `form:"dinamis_skorperintah_halusinasi" json:"dinamis_skorperintah_halusinasi"`
	DinamisHilangnyaKontrolDiri            string `form:"dinamis_hilangnya_kontrol_diri" json:"dinamis_hilangnya_kontrol_diri"`
	DinamisSkorhilangnyaKontrolDiri        *int   `form:"dinamis_skorhilangnya_kontrol_diri" json:"dinamis_skorhilangnya_kontrol_diri"`
	DinamisSeksualTidakWajar               string `form:"dinamis_seksual_tidak_wajar" json:"dinamis_seksual_tidak_wajar"`
	DinamisSkorseksualTidakWajar           *int   `form:"dinamis_skorseksual_tidak_wajar" json:"dinamis_skorseksual_tidak_wajar"`
	DinamisKemarahanFrustasi               string `form:"dinamis_kemarahan_frustasi" json:"dinamis_kemarahan_frustasi"`
	DinamisSkorkemarahanFrustasi           *int   `form:"dinamis_skorkemarahan_frustasi" json:"dinamis_skorkemarahan_frustasi"`
	DinamisKetakutanPerawatan              string `form:"dinamis_ketakutan_perawatan" json:"dinamis_ketakutan_perawatan"`
	DinamisSkorketakutanPerawatan          *int   `form:"dinamis_skorketakutan_perawatan" json:"dinamis_skorketakutan_perawatan"`
	DinamisSkortotal                       *int   `form:"dinamis_skortotal" json:"dinamis_skortotal"`
	FaktorFaktorPencegahan                 string `form:"faktor_faktor_pencegahan" json:"faktor_faktor_pencegahan"`
	TotalSkor                              *int   `form:"total_skor" json:"total_skor"`
	LevelSkor                              string `form:"level_skor" json:"level_skor"`
}

func penilaianTambahanBeresikoMelarikanDiriRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                   "required|date",
		"nip":                                       "required|string|max_len:20",
		"statik_riwayat_melarikan_diri":             "in:Ya,Tidak,Tidak Tahu",
		"statik_skorriwayat_melarikan_diri":         "int",
		"statik_riwayat_penolakan_pengobatan":       "in:Ya,Tidak,Tidak Tahu",
		"statik_skorriwayat_penolakan_pengobatan":   "int",
		"statik_usia_dibawah_35":                    "in:Ya,Tidak,Tidak Tahu",
		"statik_skorusia_dibawah_35":                "int",
		"statik_laki_laki":                          "in:Ya,Tidak,Tidak Tahu",
		"statik_skorlaki_laki":                      "int",
		"statik_diagnosis_skizofrenia":              "in:Ya,Tidak,Tidak Tahu",
		"statik_skordiagnosis_skizofrenia":          "int",
		"statik_belum_menikah":                      "in:Ya,Tidak,Tidak Tahu",
		"statik_skorbelum_menikah":                  "int",
		"statik_riwayat_penggunaan_napza":           "in:Ya,Tidak,Tidak Tahu",
		"statik_skoriwayat_penggunaan_napza":        "int",
		"statik_diagnosis_gangguan_kepribadian":     "in:Ya,Tidak,Tidak Tahu",
		"statik_skordiagnosis_gangguan_kepribadian": "int",
		"statik_riwayat_kriminal":                   "in:Ya,Tidak,Tidak Tahu",
		"statik_skorriwayat_kriminal":               "int",
		"statik_skortotal":                          "int",
		"dinamis_anti_treatment":                    "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skoranti_treatment":                "int",
		"dinamis_penggunaan_napza":                  "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorpenggunaan_napza":              "int",
		"dinamis_kebosanan":                         "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorkebosanan":                     "int",
		"dinamis_perintah_halusinasi":               "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorperintah_halusinasi":           "int",
		"dinamis_hilangnya_kontrol_diri":            "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorhilangnya_kontrol_diri":        "int",
		"dinamis_seksual_tidak_wajar":               "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorseksual_tidak_wajar":           "int",
		"dinamis_kemarahan_frustasi":                "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorkemarahan_frustasi":            "int",
		"dinamis_ketakutan_perawatan":               "in:Ya,Tidak,Tidak Tahu",
		"dinamis_skorketakutan_perawatan":           "int",
		"dinamis_skortotal":                         "int",
		"faktor_faktor_pencegahan":                  "string|max_len:500",
		"total_skor":                                "int",
		"level_skor":                                "in:Rendah(<7),Sedang(7-14),Tinggi(>14)",
	}
	return rules
}

// PenilaianTambahanBeresikoMelarikanDiriStore simpan penilaian tambahan melarikan diri; kolom waktu kunci kosong = sekarang.
type PenilaianTambahanBeresikoMelarikanDiriStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianTambahanBeresikoMelarikanDiriData
}

func (r *PenilaianTambahanBeresikoMelarikanDiriStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianTambahanBeresikoMelarikanDiriStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianTambahanBeresikoMelarikanDiriRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianTambahanBeresikoMelarikanDiriStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianTambahanBeresikoMelarikanDiriStore) Payload() PenilaianTambahanBeresikoMelarikanDiriData {
	return r.PenilaianTambahanBeresikoMelarikanDiriData
}

func (r *PenilaianTambahanBeresikoMelarikanDiriStore) DetailValues() map[string][]string { return nil }

// PenilaianTambahanBeresikoMelarikanDiriUpdate ubah penilaian tambahan melarikan diri (PUT); kunci lewat query string.
type PenilaianTambahanBeresikoMelarikanDiriUpdate struct {
	PenilaianTambahanBeresikoMelarikanDiriData
}

func (r *PenilaianTambahanBeresikoMelarikanDiriUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianTambahanBeresikoMelarikanDiriUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianTambahanBeresikoMelarikanDiriRules()
}

func (r *PenilaianTambahanBeresikoMelarikanDiriUpdate) Payload() PenilaianTambahanBeresikoMelarikanDiriData {
	return r.PenilaianTambahanBeresikoMelarikanDiriData
}

func (r *PenilaianTambahanBeresikoMelarikanDiriUpdate) DetailValues() map[string][]string { return nil }
