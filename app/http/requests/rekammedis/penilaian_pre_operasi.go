package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianPreOperasiData isian penilaian pre operasi.
type PenilaianPreOperasiData struct {
	KdDokter                    string `form:"kd_dokter" json:"kd_dokter"`
	RingkasanKlinik             string `form:"ringkasan_klinik" json:"ringkasan_klinik"`
	PemeriksaanFisik            string `form:"pemeriksaan_fisik" json:"pemeriksaan_fisik"`
	PemeriksaanDiagnostik       string `form:"pemeriksaan_diagnostik" json:"pemeriksaan_diagnostik"`
	DiagnosaPreOperasi          string `form:"diagnosa_pre_operasi" json:"diagnosa_pre_operasi"`
	RencanaTindakanBedah        string `form:"rencana_tindakan_bedah" json:"rencana_tindakan_bedah"`
	HalHalYangPerludiPersiapkan string `form:"hal_hal_yang_perludi_persiapkan" json:"hal_hal_yang_perludi_persiapkan"`
	TerapiPreOperasi            string `form:"terapi_pre_operasi" json:"terapi_pre_operasi"`
}

func penilaianPreOperasiRules() map[string]any {
	rules := map[string]any{
		"kd_dokter":                       "required|string|max_len:20",
		"ringkasan_klinik":                "string|max_len:500",
		"pemeriksaan_fisik":               "string|max_len:500",
		"pemeriksaan_diagnostik":          "string|max_len:500",
		"diagnosa_pre_operasi":            "string|max_len:500",
		"rencana_tindakan_bedah":          "string|max_len:500",
		"hal_hal_yang_perludi_persiapkan": "string|max_len:500",
		"terapi_pre_operasi":              "string|max_len:500",
	}
	return rules
}

// PenilaianPreOperasiStore simpan penilaian pre operasi; kolom waktu kunci kosong = sekarang.
type PenilaianPreOperasiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PenilaianPreOperasiData
}

func (r *PenilaianPreOperasiStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPreOperasiStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianPreOperasiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PenilaianPreOperasiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PenilaianPreOperasiStore) Payload() PenilaianPreOperasiData {
	return r.PenilaianPreOperasiData
}

func (r *PenilaianPreOperasiStore) DetailValues() map[string][]string { return nil }

// PenilaianPreOperasiUpdate ubah penilaian pre operasi (PUT); kunci lewat query string.
type PenilaianPreOperasiUpdate struct {
	PenilaianPreOperasiData
}

func (r *PenilaianPreOperasiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPreOperasiUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianPreOperasiRules()
}

func (r *PenilaianPreOperasiUpdate) Payload() PenilaianPreOperasiData {
	return r.PenilaianPreOperasiData
}

func (r *PenilaianPreOperasiUpdate) DetailValues() map[string][]string { return nil }
