package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ChecklistPreOperasiData isian checklist pre operasi.
type ChecklistPreOperasiData struct {
	Sncn                                  string `form:"sncn" json:"sncn"`
	Tindakan                              string `form:"tindakan" json:"tindakan"`
	KdDokterBedah                         string `form:"kd_dokter_bedah" json:"kd_dokter_bedah"`
	KdDokterAnestesi                      string `form:"kd_dokter_anestesi" json:"kd_dokter_anestesi"`
	Identitas                             string `form:"identitas" json:"identitas"`
	SuratIjinBedah                        string `form:"surat_ijin_bedah" json:"surat_ijin_bedah"`
	SuratIjinAnestesi                     string `form:"surat_ijin_anestesi" json:"surat_ijin_anestesi"`
	SuratIjinTransfusi                    string `form:"surat_ijin_transfusi" json:"surat_ijin_transfusi"`
	PenandaanAreaOperasi                  string `form:"penandaan_area_operasi" json:"penandaan_area_operasi"`
	KeadaanUmum                           string `form:"keadaan_umum" json:"keadaan_umum"`
	PemeriksaanPenunjangRontgen           string `form:"pemeriksaan_penunjang_rontgen" json:"pemeriksaan_penunjang_rontgen"`
	KeteranganPemeriksaanPenunjangRontgen string `form:"keterangan_pemeriksaan_penunjang_rontgen" json:"keterangan_pemeriksaan_penunjang_rontgen"`
	PemeriksaanPenunjangEkg               string `form:"pemeriksaan_penunjang_ekg" json:"pemeriksaan_penunjang_ekg"`
	KeteranganPemeriksaanPenunjangEkg     string `form:"keterangan_pemeriksaan_penunjang_ekg" json:"keterangan_pemeriksaan_penunjang_ekg"`
	PemeriksaanPenunjangUsg               string `form:"pemeriksaan_penunjang_usg" json:"pemeriksaan_penunjang_usg"`
	KeteranganPemeriksaanPenunjangUsg     string `form:"keterangan_pemeriksaan_penunjang_usg" json:"keterangan_pemeriksaan_penunjang_usg"`
	PemeriksaanPenunjangCtscan            string `form:"pemeriksaan_penunjang_ctscan" json:"pemeriksaan_penunjang_ctscan"`
	KeteranganPemeriksaanPenunjangCtscan  string `form:"keterangan_pemeriksaan_penunjang_ctscan" json:"keterangan_pemeriksaan_penunjang_ctscan"`
	PemeriksaanPenunjangMri               string `form:"pemeriksaan_penunjang_mri" json:"pemeriksaan_penunjang_mri"`
	KeteranganPemeriksaanPenunjangMri     string `form:"keterangan_pemeriksaan_penunjang_mri" json:"keterangan_pemeriksaan_penunjang_mri"`
	PersiapanDarah                        string `form:"persiapan_darah" json:"persiapan_darah"`
	KeteranganPersiapanDarah              string `form:"keterangan_persiapan_darah" json:"keterangan_persiapan_darah"`
	PerlengkapanKhusus                    string `form:"perlengkapan_khusus" json:"perlengkapan_khusus"`
	NipPetugasRuangan                     string `form:"nip_petugas_ruangan" json:"nip_petugas_ruangan"`
	NipPerawatOk                          string `form:"nip_perawat_ok" json:"nip_perawat_ok"`
}

func checklistPreOperasiRules() map[string]any {
	rules := map[string]any{
		"sncn":                          "string|max_len:25",
		"tindakan":                      "string|max_len:50",
		"kd_dokter_bedah":               "required|string|max_len:20",
		"kd_dokter_anestesi":            "required|string|max_len:20",
		"identitas":                     "in:Ya,Tidak",
		"surat_ijin_bedah":              "in:Ada,Tidak Ada",
		"surat_ijin_anestesi":           "in:Ada,Tidak Ada",
		"surat_ijin_transfusi":          "in:Ada,Tidak Ada,Tidak Diperlukan",
		"penandaan_area_operasi":        "in:Ada,Tidak Ada,Tidak Diperlukan",
		"keadaan_umum":                  "in:Baik,Sedang,Lemah",
		"pemeriksaan_penunjang_rontgen": "in:Ada,Tidak Ada,Tidak Diperlukan",
		"keterangan_pemeriksaan_penunjang_rontgen": "string|max_len:20",
		"pemeriksaan_penunjang_ekg":                "in:Ada,Tidak Ada,Tidak Diperlukan",
		"keterangan_pemeriksaan_penunjang_ekg":     "string|max_len:20",
		"pemeriksaan_penunjang_usg":                "in:Ada,Tidak Ada,Tidak Diperlukan",
		"keterangan_pemeriksaan_penunjang_usg":     "string|max_len:20",
		"pemeriksaan_penunjang_ctscan":             "in:Ada,Tidak Ada,Tidak Diperlukan",
		"keterangan_pemeriksaan_penunjang_ctscan":  "string|max_len:20",
		"pemeriksaan_penunjang_mri":                "in:Ada,Tidak Ada,Tidak Diperlukan",
		"keterangan_pemeriksaan_penunjang_mri":     "string|max_len:20",
		"persiapan_darah":                          "in:Ada,Tidak Ada,Tidak Diperlukan",
		"keterangan_persiapan_darah":               "string|max_len:20",
		"perlengkapan_khusus":                      "in:Ada,Tidak Ada,Tidak Diperlukan",
		"nip_petugas_ruangan":                      "string|max_len:20",
		"nip_perawat_ok":                           "string|max_len:20",
	}
	return rules
}

// ChecklistPreOperasiStore simpan checklist pre operasi; kolom waktu kunci kosong = sekarang.
type ChecklistPreOperasiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	ChecklistPreOperasiData
}

func (r *ChecklistPreOperasiStore) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistPreOperasiStore) Rules(ctx http.Context) map[string]any {
	rules := checklistPreOperasiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *ChecklistPreOperasiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *ChecklistPreOperasiStore) Payload() ChecklistPreOperasiData {
	return r.ChecklistPreOperasiData
}

func (r *ChecklistPreOperasiStore) DetailValues() map[string][]string { return nil }

// ChecklistPreOperasiUpdate ubah checklist pre operasi (PUT); kunci lewat query string.
type ChecklistPreOperasiUpdate struct {
	ChecklistPreOperasiData
}

func (r *ChecklistPreOperasiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistPreOperasiUpdate) Rules(ctx http.Context) map[string]any {
	return checklistPreOperasiRules()
}

func (r *ChecklistPreOperasiUpdate) Payload() ChecklistPreOperasiData {
	return r.ChecklistPreOperasiData
}

func (r *ChecklistPreOperasiUpdate) DetailValues() map[string][]string { return nil }
