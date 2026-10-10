package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// ChecklistPostOperasiData isian checklist post operasi.
type ChecklistPostOperasiData struct {
	Sncn                                  string `form:"sncn" json:"sncn"`
	Tindakan                              string `form:"tindakan" json:"tindakan"`
	KdDokterBedah                         string `form:"kd_dokter_bedah" json:"kd_dokter_bedah"`
	KdDokterAnestesi                      string `form:"kd_dokter_anestesi" json:"kd_dokter_anestesi"`
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
	JenisCairanInfus                      string `form:"jenis_cairan_infus" json:"jenis_cairan_infus"`
	KateterUrine                          string `form:"kateter_urine" json:"kateter_urine"`
	TanggalPemasanganKateter              string `form:"tanggal_pemasangan_kateter" json:"tanggal_pemasangan_kateter"`
	WarnaKateter                          string `form:"warna_kateter" json:"warna_kateter"`
	JumlahKateter                         string `form:"jumlah_kateter" json:"jumlah_kateter"`
	AreaLukaOperasi                       string `form:"area_luka_operasi" json:"area_luka_operasi"`
	Drain                                 string `form:"drain" json:"drain"`
	JumlahDrain                           string `form:"jumlah_drain" json:"jumlah_drain"`
	LetakDrain                            string `form:"letak_drain" json:"letak_drain"`
	WarnaDrain                            string `form:"warna_drain" json:"warna_drain"`
	JaringanPa                            string `form:"jaringan_pa" json:"jaringan_pa"`
	NipPerawatOk                          string `form:"nip_perawat_ok" json:"nip_perawat_ok"`
	NipPerawatAnestesi                    string `form:"nip_perawat_anestesi" json:"nip_perawat_anestesi"`
}

func checklistPostOperasiRules() map[string]any {
	rules := map[string]any{
		"sncn":                          "string|max_len:25",
		"tindakan":                      "string|max_len:50",
		"kd_dokter_bedah":               "required|string|max_len:20",
		"kd_dokter_anestesi":            "required|string|max_len:20",
		"keadaan_umum":                  "in:Sadar,Tidur,Terintubasi",
		"pemeriksaan_penunjang_rontgen": "in:Ada,Tidak Ada",
		"keterangan_pemeriksaan_penunjang_rontgen": "string|max_len:20",
		"pemeriksaan_penunjang_ekg":                "in:Ada,Tidak Ada",
		"keterangan_pemeriksaan_penunjang_ekg":     "string|max_len:20",
		"pemeriksaan_penunjang_usg":                "in:Ada,Tidak Ada",
		"keterangan_pemeriksaan_penunjang_usg":     "string|max_len:20",
		"pemeriksaan_penunjang_ctscan":             "in:Ada,Tidak Ada",
		"keterangan_pemeriksaan_penunjang_ctscan":  "string|max_len:20",
		"pemeriksaan_penunjang_mri":                "in:Ada,Tidak Ada",
		"keterangan_pemeriksaan_penunjang_mri":     "string|max_len:20",
		"jenis_cairan_infus":                       "string|max_len:40",
		"kateter_urine":                            "in:Ada,Tidak Ada",
		"tanggal_pemasangan_kateter":               "date",
		"warna_kateter":                            "in:Jernih,Keruh,-",
		"jumlah_kateter":                           "string|max_len:4",
		"area_luka_operasi":                        "string|max_len:120",
		"drain":                                    "in:Ada,Tidak Ada",
		"jumlah_drain":                             "string|max_len:2",
		"letak_drain":                              "string|max_len:40",
		"warna_drain":                              "string|max_len:30",
		"jaringan_pa":                              "in:Ada,Tidak Ada",
		"nip_perawat_ok":                           "string|max_len:20",
		"nip_perawat_anestesi":                     "string|max_len:20",
	}
	return rules
}

// ChecklistPostOperasiStore simpan checklist post operasi; kolom waktu kunci kosong = sekarang.
type ChecklistPostOperasiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	ChecklistPostOperasiData
}

func (r *ChecklistPostOperasiStore) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistPostOperasiStore) Rules(ctx http.Context) map[string]any {
	rules := checklistPostOperasiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *ChecklistPostOperasiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *ChecklistPostOperasiStore) Payload() ChecklistPostOperasiData {
	return r.ChecklistPostOperasiData
}

func (r *ChecklistPostOperasiStore) DetailValues() map[string][]string { return nil }

// ChecklistPostOperasiUpdate ubah checklist post operasi (PUT); kunci lewat query string.
type ChecklistPostOperasiUpdate struct {
	ChecklistPostOperasiData
}

func (r *ChecklistPostOperasiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *ChecklistPostOperasiUpdate) Rules(ctx http.Context) map[string]any {
	return checklistPostOperasiRules()
}

func (r *ChecklistPostOperasiUpdate) Payload() ChecklistPostOperasiData {
	return r.ChecklistPostOperasiData
}

func (r *ChecklistPostOperasiUpdate) DetailValues() map[string][]string { return nil }
