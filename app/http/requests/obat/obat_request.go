package obat

import (
	"github.com/goravel/framework/contracts/http"
)

// Data field obat (databarang) yang dipakai bersama oleh store & update.
// Status aktif tidak diisi di sini: obat baru selalu aktif, ubah lewat PATCH /obat/{kode}/status.
type Data struct {
	NamaBrng     string  `form:"nama_brng" json:"nama_brng"`
	KodeSatbesar string  `form:"kode_satbesar" json:"kode_satbesar"`
	KodeSat      string  `form:"kode_sat" json:"kode_sat"`
	LetakBarang  string  `form:"letak_barang" json:"letak_barang"`
	Dasar        float64 `form:"dasar" json:"dasar"`
	HBeli        float64 `form:"h_beli" json:"h_beli"`
	Ralan        float64 `form:"ralan" json:"ralan"`
	Kelas1       float64 `form:"kelas1" json:"kelas1"`
	Kelas2       float64 `form:"kelas2" json:"kelas2"`
	Kelas3       float64 `form:"kelas3" json:"kelas3"`
	Utama        float64 `form:"utama" json:"utama"`
	Vip          float64 `form:"vip" json:"vip"`
	Vvip         float64 `form:"vvip" json:"vvip"`
	BeliLuar     float64 `form:"beliluar" json:"beliluar"`
	JualBebas    float64 `form:"jualbebas" json:"jualbebas"`
	Karyawan     float64 `form:"karyawan" json:"karyawan"`
	StokMinimal  float64 `form:"stokminimal" json:"stokminimal"`
	Kdjns        string  `form:"kdjns" json:"kdjns"`
	Isi          float64 `form:"isi" json:"isi"`
	Kapasitas    float64 `form:"kapasitas" json:"kapasitas"`
	Expire       string  `form:"expire" json:"expire"` // YYYY-MM-DD, boleh kosong
	KodeIndustri string  `form:"kode_industri" json:"kode_industri"`
	KodeKategori string  `form:"kode_kategori" json:"kode_kategori"`
	KodeGolongan string  `form:"kode_golongan" json:"kode_golongan"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nama_brng":     "required|string|max_len:80",
		"kode_satbesar": "required|string|max_len:4",
		"kode_sat":      "string|max_len:4",
		"letak_barang":  "string|max_len:100",
		"dasar":         "required|numeric",
		"h_beli":        "numeric",
		"ralan":         "numeric",
		"kelas1":        "numeric",
		"kelas2":        "numeric",
		"kelas3":        "numeric",
		"utama":         "numeric",
		"vip":           "numeric",
		"vvip":          "numeric",
		"beliluar":      "numeric",
		"jualbebas":     "numeric",
		"karyawan":      "numeric",
		"stokminimal":   "numeric",
		"kdjns":         "string|max_len:4",
		"isi":           "required|numeric",
		"kapasitas":     "required|numeric",
		"expire":        "date",
		"kode_industri": "string|max_len:5",
		"kode_kategori": "string|max_len:4",
		"kode_golongan": "string|max_len:4",
	}
}

type StoreRequest struct {
	KodeBrng string `form:"kode_brng" json:"kode_brng"`
	Data
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kode_brng"] = "required|string|max_len:15"
	return rules
}

// UpdateRequest mengganti data obat (PUT); kode_brng diambil dari route.
type UpdateRequest struct {
	Data
}

func (r *UpdateRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}

// StatusRequest PATCH /obat/{kode}/status: "1" aktif, "0" nonaktif.
type StatusRequest struct {
	Status string `form:"status" json:"status"`
}

func (r *StatusRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StatusRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{"status": "required|in:0,1"}
}
