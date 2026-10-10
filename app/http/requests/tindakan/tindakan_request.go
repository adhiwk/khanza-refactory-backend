package tindakan

import (
	"github.com/goravel/framework/contracts/http"
)

// StoreRequest simpan satu/lebih tindakan pada waktu yang sama; tarif diambil dari master, bukan dari client.
type StoreRequest struct {
	NoRawat    string   `form:"no_rawat" json:"no_rawat"`
	Pelaksana  string   `form:"pelaksana" json:"pelaksana"` // dr | pr | drpr
	KdDokter   string   `form:"kd_dokter" json:"kd_dokter"`
	Nip        string   `form:"nip" json:"nip"`
	Tgl        string   `form:"tgl_perawatan" json:"tgl_perawatan"` // YYYY-MM-DD, default hari ini
	Jam        string   `form:"jam_rawat" json:"jam_rawat"`         // HH:MM:SS, default jam sekarang
	KdJenisPrw []string `form:"kd_jenis_prw" json:"kd_jenis_prw"`
}

func (r *StoreRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StoreRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rawat":       "required|string|max_len:17",
		"pelaksana":      "required|in:dr,pr,drpr",
		"kd_dokter":      "required_if:pelaksana,dr,drpr|string|max_len:20",
		"nip":            "required_if:pelaksana,pr,drpr|string|max_len:20",
		"tgl_perawatan":  "date",
		"jam_rawat":      "string|len:8",
		"kd_jenis_prw":   "required|slice",
		"kd_jenis_prw.*": "required|string|max_len:15",
	}
}

// DeleteRequest identitas satu baris tindakan (primary key tabel).
type DeleteRequest struct {
	NoRawat    string `form:"no_rawat" json:"no_rawat"`
	Pelaksana  string `form:"pelaksana" json:"pelaksana"`
	KdJenisPrw string `form:"kd_jenis_prw" json:"kd_jenis_prw"`
	KdDokter   string `form:"kd_dokter" json:"kd_dokter"`
	Nip        string `form:"nip" json:"nip"`
	Tgl        string `form:"tgl_perawatan" json:"tgl_perawatan"`
	Jam        string `form:"jam_rawat" json:"jam_rawat"`
}

func (r *DeleteRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *DeleteRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rawat":      "required|string|max_len:17",
		"pelaksana":     "required|in:dr,pr,drpr",
		"kd_jenis_prw":  "required|string|max_len:15",
		"kd_dokter":     "required_if:pelaksana,dr,drpr|string|max_len:20",
		"nip":           "required_if:pelaksana,pr,drpr|string|max_len:20",
		"tgl_perawatan": "required|date",
		"jam_rawat":     "required|string|len:8",
	}
}
