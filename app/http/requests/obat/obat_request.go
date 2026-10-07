package obat

import (
	"github.com/goravel/framework/contracts/http"

	obatAction "goravel/app/actions/obat"
)

// baseRules aturan validasi yang dipakai bersama oleh create & update.
func baseRules() map[string]any {
	return map[string]any{
		"nama_brng":     "required|max_len:80",
		"kode_satbesar": "required|max_len:4",
		"kode_sat":      "max_len:4",
		"letak_barang":  "max_len:100",
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
		"kdjns":         "max_len:4",
		"isi":           "required|numeric",
		"kapasitas":     "required|numeric",
		"expire":        "date",
		"status":        "required|in:0,1",
		"kode_industri": "max_len:5",
		"kode_kategori": "max_len:4",
		"kode_golongan": "max_len:4",
	}
}

// StoreObatRequest validasi untuk POST /obat.
type StoreObatRequest struct {
	obatAction.ObatRequest
}

func (r *StoreObatRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreObatRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["kode_brng"] = "required|max_len:15"
	return rules
}

// UpdateObatRequest validasi untuk PUT /obat/{kode}.
// kode_brng tidak divalidasi karena primary key diambil dari URL.
type UpdateObatRequest struct {
	obatAction.ObatRequest
}

func (r *UpdateObatRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateObatRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}

// UpdateStatusRequest validasi untuk PATCH /obat/{kode}/status.
type UpdateStatusRequest struct {
	Status string `form:"status" json:"status"`
}

func (r *UpdateStatusRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdateStatusRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"status": "required|in:0,1",
	}
}
