package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// TimeoutSebelumInsisiData isian time out sebelum insisi.
type TimeoutSebelumInsisiData struct {
	Sncn                      string `form:"sncn" json:"sncn"`
	Tindakan                  string `form:"tindakan" json:"tindakan"`
	KdDokterBedah             string `form:"kd_dokter_bedah" json:"kd_dokter_bedah"`
	KdDokterAnestesi          string `form:"kd_dokter_anestesi" json:"kd_dokter_anestesi"`
	VerbalIdentitas           string `form:"verbal_identitas" json:"verbal_identitas"`
	VerbalTindakan            string `form:"verbal_tindakan" json:"verbal_tindakan"`
	VerbalAreaInsisi          string `form:"verbal_area_insisi" json:"verbal_area_insisi"`
	PenandaanAreaOperasi      string `form:"penandaan_area_operasi" json:"penandaan_area_operasi"`
	LamaOperasi               string `form:"lama_operasi" json:"lama_operasi"`
	PenayanganRadiologi       string `form:"penayangan_radiologi" json:"penayangan_radiologi"`
	PenayanganCtscan          string `form:"penayangan_ctscan" json:"penayangan_ctscan"`
	PenayanganMri             string `form:"penayangan_mri" json:"penayangan_mri"`
	AntibiotikProfilaks       string `form:"antibiotik_profilaks" json:"antibiotik_profilaks"`
	NamaAntibiotik            string `form:"nama_antibiotik" json:"nama_antibiotik"`
	JamPemberian              string `form:"jam_pemberian" json:"jam_pemberian"`
	AntisipasiKehilanganDarah string `form:"antisipasi_kehilangan_darah" json:"antisipasi_kehilangan_darah"`
	HalKhusus                 string `form:"hal_khusus" json:"hal_khusus"`
	HalKhususDiperhatikan     string `form:"hal_khusus_diperhatikan" json:"hal_khusus_diperhatikan"`
	TanggalSteril             string `form:"tanggal_steril" json:"tanggal_steril"`
	PetujukSterilisasi        string `form:"petujuk_sterilisasi" json:"petujuk_sterilisasi"`
	VerifikasiPreoperatif     string `form:"verifikasi_preoperatif" json:"verifikasi_preoperatif"`
	NipPerawatOk              string `form:"nip_perawat_ok" json:"nip_perawat_ok"`
}

func timeoutSebelumInsisiRules() map[string]any {
	rules := map[string]any{
		"sncn":                        "string|max_len:25",
		"tindakan":                    "string|max_len:50",
		"kd_dokter_bedah":             "required|string|max_len:20",
		"kd_dokter_anestesi":          "required|string|max_len:20",
		"verbal_identitas":            "in:Ya,Tidak",
		"verbal_tindakan":             "in:Ya,Tidak",
		"verbal_area_insisi":          "in:Ya,Tidak",
		"penandaan_area_operasi":      "in:Ada,Tidak Ada,Tidak Diperlukan",
		"lama_operasi":                "string|max_len:10",
		"penayangan_radiologi":        "in:Ditayangkan,Benar,Tidak Diperlukan",
		"penayangan_ctscan":           "in:Ditayangkan,Benar,Tidak Diperlukan",
		"penayangan_mri":              "in:Ditayangkan,Benar,Tidak Diperlukan",
		"antibiotik_profilaks":        "in:Ya,Tidak",
		"nama_antibiotik":             "string|max_len:50",
		"jam_pemberian":               "string|max_len:10",
		"antisipasi_kehilangan_darah": "string|max_len:50",
		"hal_khusus":                  "in:Ada,Tidak Ada",
		"hal_khusus_diperhatikan":     "string|max_len:100",
		"tanggal_steril":              "date",
		"petujuk_sterilisasi":         "in:Ya,Tidak",
		"verifikasi_preoperatif":      "in:Ya,Tidak",
		"nip_perawat_ok":              "string|max_len:20",
	}
	return rules
}

// TimeoutSebelumInsisiStore simpan time out sebelum insisi; kolom waktu kunci kosong = sekarang.
type TimeoutSebelumInsisiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	TimeoutSebelumInsisiData
}

func (r *TimeoutSebelumInsisiStore) Authorize(ctx http.Context) error { return nil }

func (r *TimeoutSebelumInsisiStore) Rules(ctx http.Context) map[string]any {
	rules := timeoutSebelumInsisiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *TimeoutSebelumInsisiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *TimeoutSebelumInsisiStore) Payload() TimeoutSebelumInsisiData {
	return r.TimeoutSebelumInsisiData
}

func (r *TimeoutSebelumInsisiStore) DetailValues() map[string][]string { return nil }

// TimeoutSebelumInsisiUpdate ubah time out sebelum insisi (PUT); kunci lewat query string.
type TimeoutSebelumInsisiUpdate struct {
	TimeoutSebelumInsisiData
}

func (r *TimeoutSebelumInsisiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *TimeoutSebelumInsisiUpdate) Rules(ctx http.Context) map[string]any {
	return timeoutSebelumInsisiRules()
}

func (r *TimeoutSebelumInsisiUpdate) Payload() TimeoutSebelumInsisiData {
	return r.TimeoutSebelumInsisiData
}

func (r *TimeoutSebelumInsisiUpdate) DetailValues() map[string][]string { return nil }
