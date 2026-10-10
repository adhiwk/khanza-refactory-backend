package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SigninSebelumAnestesiData isian sign in sebelum anastesi.
type SigninSebelumAnestesiData struct {
	Sncn                                      string `form:"sncn" json:"sncn"`
	Tindakan                                  string `form:"tindakan" json:"tindakan"`
	KdDokterBedah                             string `form:"kd_dokter_bedah" json:"kd_dokter_bedah"`
	KdDokterAnestesi                          string `form:"kd_dokter_anestesi" json:"kd_dokter_anestesi"`
	Identitas                                 string `form:"identitas" json:"identitas"`
	PenandaanAreaOperasi                      string `form:"penandaan_area_operasi" json:"penandaan_area_operasi"`
	Alergi                                    string `form:"alergi" json:"alergi"`
	ResikoAspirasi                            string `form:"resiko_aspirasi" json:"resiko_aspirasi"`
	ResikoAspirasiRencanaAntisipasi           string `form:"resiko_aspirasi_rencana_antisipasi" json:"resiko_aspirasi_rencana_antisipasi"`
	ResikoKehilanganDarah                     string `form:"resiko_kehilangan_darah" json:"resiko_kehilangan_darah"`
	ResikoKehilanganDarahLine                 string `form:"resiko_kehilangan_darah_line" json:"resiko_kehilangan_darah_line"`
	ResikoKehilanganDarahRencanaAntisipasi    string `form:"resiko_kehilangan_darah_rencana_antisipasi" json:"resiko_kehilangan_darah_rencana_antisipasi"`
	KesiapanAlatObatAnestesi                  string `form:"kesiapan_alat_obat_anestesi" json:"kesiapan_alat_obat_anestesi"`
	KesiapanAlatObatAnestesiRencanaAntisipasi string `form:"kesiapan_alat_obat_anestesi_rencana_antisipasi" json:"kesiapan_alat_obat_anestesi_rencana_antisipasi"`
	NipPerawatOk                              string `form:"nip_perawat_ok" json:"nip_perawat_ok"`
}

func signinSebelumAnestesiRules() map[string]any {
	rules := map[string]any{
		"sncn":                               "string|max_len:25",
		"tindakan":                           "string|max_len:50",
		"kd_dokter_bedah":                    "required|string|max_len:20",
		"kd_dokter_anestesi":                 "required|string|max_len:20",
		"identitas":                          "in:Ya,Tidak",
		"penandaan_area_operasi":             "in:Ada,Tidak Ada,Tidak Diperlukan",
		"alergi":                             "string|max_len:30",
		"resiko_aspirasi":                    "in:Ada,Tidak Ada",
		"resiko_aspirasi_rencana_antisipasi": "string|max_len:50",
		"resiko_kehilangan_darah":            "in:Tidak Ada,Ada",
		"resiko_kehilangan_darah_line":       "string|max_len:30",
		"resiko_kehilangan_darah_rencana_antisipasi":     "string|max_len:50",
		"kesiapan_alat_obat_anestesi":                    "in:Lengkap,Pulsa Oximetri,Tidak Lengkap",
		"kesiapan_alat_obat_anestesi_rencana_antisipasi": "string|max_len:50",
		"nip_perawat_ok":                                 "string|max_len:20",
	}
	return rules
}

// SigninSebelumAnestesiStore simpan sign in sebelum anastesi; kolom waktu kunci kosong = sekarang.
type SigninSebelumAnestesiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	SigninSebelumAnestesiData
}

func (r *SigninSebelumAnestesiStore) Authorize(ctx http.Context) error { return nil }

func (r *SigninSebelumAnestesiStore) Rules(ctx http.Context) map[string]any {
	rules := signinSebelumAnestesiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *SigninSebelumAnestesiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *SigninSebelumAnestesiStore) Payload() SigninSebelumAnestesiData {
	return r.SigninSebelumAnestesiData
}

func (r *SigninSebelumAnestesiStore) DetailValues() map[string][]string { return nil }

// SigninSebelumAnestesiUpdate ubah sign in sebelum anastesi (PUT); kunci lewat query string.
type SigninSebelumAnestesiUpdate struct {
	SigninSebelumAnestesiData
}

func (r *SigninSebelumAnestesiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SigninSebelumAnestesiUpdate) Rules(ctx http.Context) map[string]any {
	return signinSebelumAnestesiRules()
}

func (r *SigninSebelumAnestesiUpdate) Payload() SigninSebelumAnestesiData {
	return r.SigninSebelumAnestesiData
}

func (r *SigninSebelumAnestesiUpdate) DetailValues() map[string][]string { return nil }
