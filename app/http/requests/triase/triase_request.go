package triase

import (
	"github.com/goravel/framework/contracts/http"
)

// Data isian triase IGD. jenis "primer" (skala_level 1/2, plan Ruang Resusitasi|Ruang Kritis, keluhan = keluhan utama)
// atau "sekunder" (skala_level 3/4/5, plan Zona Kuning|Zona Hijau, keluhan = anamnesa singkat).
type Data struct {
	TglKunjungan         string   `form:"tgl_kunjungan" json:"tgl_kunjungan"`
	CaraMasuk            string   `form:"cara_masuk" json:"cara_masuk"`
	AlatTransportasi     string   `form:"alat_transportasi" json:"alat_transportasi"`
	AlasanKedatangan     string   `form:"alasan_kedatangan" json:"alasan_kedatangan"`
	KeteranganKedatangan string   `form:"keterangan_kedatangan" json:"keterangan_kedatangan"`
	KodeKasus            string   `form:"kode_kasus" json:"kode_kasus"`
	TekananDarah         string   `form:"tekanan_darah" json:"tekanan_darah"`
	Nadi                 string   `form:"nadi" json:"nadi"`
	Pernapasan           string   `form:"pernapasan" json:"pernapasan"`
	Suhu                 string   `form:"suhu" json:"suhu"`
	SaturasiO2           string   `form:"saturasi_o2" json:"saturasi_o2"`
	Nyeri                string   `form:"nyeri" json:"nyeri"`
	Jenis                string   `form:"jenis" json:"jenis"`
	Keluhan              string   `form:"keluhan" json:"keluhan"`
	KebutuhanKhusus      string   `form:"kebutuhan_khusus" json:"kebutuhan_khusus"`
	Catatan              string   `form:"catatan" json:"catatan"`
	Plan                 string   `form:"plan" json:"plan"`
	TanggalTriase        string   `form:"tanggaltriase" json:"tanggaltriase"`
	Nik                  string   `form:"nik" json:"nik"`
	SkalaLevel           int      `form:"skala_level" json:"skala_level"`
	SkalaKode            []string `form:"skala_kode" json:"skala_kode"`
}

func baseRules() map[string]any {
	return map[string]any{
		"tgl_kunjungan":         "required|date",
		"cara_masuk":            "required|in:Jalan,Brankar,Kursi Roda,Digendong",
		"alat_transportasi":     "required|in:-,AGD,Sendiri,Swasta",
		"alasan_kedatangan":     "required|in:Datang Sendiri,Polisi,Rujukan,Bidan,Puskesmas,Rumah Sakit,Poliklinik,Faskes Lain,-",
		"keterangan_kedatangan": "required|string|max_len:100",
		"kode_kasus":            "required|string|max_len:3",
		"tekanan_darah":         "required|string|max_len:8",
		"nadi":                  "required|string|max_len:3",
		"pernapasan":            "required|string|max_len:3",
		"suhu":                  "required|string|max_len:5",
		"saturasi_o2":           "required|string|max_len:3",
		"nyeri":                 "required|string|max_len:5",
		"jenis":                 "required|in:primer,sekunder",
		"keluhan":               "required|string|max_len:400",
		"kebutuhan_khusus":      "in:-,UPPA,Airborne,Dekontaminan",
		"catatan":               "required|string|max_len:100",
		"plan":                  "required|in:Ruang Resusitasi,Ruang Kritis,Zona Kuning,Zona Hijau",
		"tanggaltriase":         "required|date",
		"nik":                   "required|string|max_len:20",
		"skala_level":           "required|in:1,2,3,4,5",
		"skala_kode":            "required|slice",
		"skala_kode.*":          "required|string|max_len:3",
	}
}

type StoreRequest struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error { return nil }

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

// UpdateRequest no_rawat lewat query string.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error { return nil }

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any { return baseRules() }
