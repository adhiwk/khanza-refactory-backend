package templatedokter

import (
	"github.com/goravel/framework/contracts/http"
)

type Kode struct {
	Kode   string `form:"kode" json:"kode"`
	Urut   int    `form:"urut" json:"urut"`
	Jumlah string `form:"jumlah" json:"jumlah"`
}

type Lab struct {
	KdJenisPrw string `form:"kd_jenis_prw" json:"kd_jenis_prw"`
	IDTemplate []int  `form:"id_template" json:"id_template"`
}

type Resep struct {
	KodeBrng    string  `form:"kode_brng" json:"kode_brng"`
	Jml         float64 `form:"jml" json:"jml"`
	AturanPakai string  `form:"aturan_pakai" json:"aturan_pakai"`
}

type RacikanDetail struct {
	KodeBrng  string  `form:"kode_brng" json:"kode_brng"`
	P1        float64 `form:"p1" json:"p1"`
	P2        float64 `form:"p2" json:"p2"`
	Kandungan string  `form:"kandungan" json:"kandungan"`
	Jml       float64 `form:"jml" json:"jml"`
}

type Racikan struct {
	NoRacik     string          `form:"no_racik" json:"no_racik"`
	NamaRacik   string          `form:"nama_racik" json:"nama_racik"`
	KdRacik     string          `form:"kd_racik" json:"kd_racik"`
	JmlDr       int             `form:"jml_dr" json:"jml_dr"`
	AturanPakai string          `form:"aturan_pakai" json:"aturan_pakai"`
	Keterangan  string          `form:"keterangan" json:"keterangan"`
	Detail      []RacikanDetail `form:"detail" json:"detail"`
}

// Request simpan/ubah template pemeriksaan dokter; no_template dibuat otomatis (TPD + 16 digit).
type Request struct {
	KdDokter    string    `form:"kd_dokter" json:"kd_dokter"`
	Keluhan     string    `form:"keluhan" json:"keluhan"`
	Pemeriksaan string    `form:"pemeriksaan" json:"pemeriksaan"`
	Penilaian   string    `form:"penilaian" json:"penilaian"`
	Rencana     string    `form:"rencana" json:"rencana"`
	Instruksi   string    `form:"instruksi" json:"instruksi"`
	Evaluasi    string    `form:"evaluasi" json:"evaluasi"`
	Diagnosa    []Kode    `form:"diagnosa" json:"diagnosa"`
	Prosedur    []Kode    `form:"prosedur" json:"prosedur"`
	Radiologi   []Kode    `form:"radiologi" json:"radiologi"`
	Lab         []Lab     `form:"lab" json:"lab"`
	Resep       []Resep   `form:"resep" json:"resep"`
	Racikan     []Racikan `form:"racikan" json:"racikan"`
	Tindakan    []Kode    `form:"tindakan" json:"tindakan"`
}

func (r *Request) Authorize(ctx http.Context) error { return nil }

// Rules daftar detail opsional: rule "required" pada wildcard Goravel ikut gagal untuk daftar kosong,
// sehingga item tanpa kode dibuang di Action.
func (r *Request) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"kd_dokter":                    "required|string|max_len:20",
		"keluhan":                      "string|max_len:2000",
		"pemeriksaan":                  "string|max_len:2000",
		"penilaian":                    "string|max_len:2000",
		"rencana":                      "required|string|max_len:2000",
		"instruksi":                    "required|string|max_len:2000",
		"evaluasi":                     "string|max_len:2000",
		"diagnosa":                     "slice",
		"diagnosa.*.kode":              "string|max_len:15",
		"prosedur":                     "slice",
		"prosedur.*.kode":              "string|max_len:8",
		"prosedur.*.jumlah":            "string|max_len:3",
		"radiologi":                    "slice",
		"radiologi.*.kode":             "string|max_len:15",
		"tindakan":                     "slice",
		"tindakan.*.kode":              "string|max_len:15",
		"lab":                          "slice",
		"lab.*.kd_jenis_prw":           "string|max_len:15",
		"resep":                        "slice",
		"resep.*.kode_brng":            "string|max_len:15",
		"resep.*.jml":                  "numeric|min:0",
		"resep.*.aturan_pakai":         "string|max_len:150",
		"racikan":                      "slice",
		"racikan.*.no_racik":           "string|max_len:2",
		"racikan.*.nama_racik":         "string|max_len:100",
		"racikan.*.kd_racik":           "string|max_len:3",
		"racikan.*.aturan_pakai":       "string|max_len:150",
		"racikan.*.keterangan":         "string|max_len:50",
		"racikan.*.detail":             "slice",
		"racikan.*.detail.*.kode_brng": "string|max_len:15",
		"racikan.*.detail.*.kandungan": "string|max_len:10",
	}
}
