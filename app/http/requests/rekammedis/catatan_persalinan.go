package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanPersalinanData isian catatan persalinan.
type CatatanPersalinanData struct {
	Mulai                 string `form:"mulai" json:"mulai"`
	Selesai               string `form:"selesai" json:"selesai"`
	KdDokter              string `form:"kd_dokter" json:"kd_dokter"`
	Nip                   string `form:"nip" json:"nip"`
	Catatan               string `form:"catatan" json:"catatan"`
	WaktuPersalinanKala1  string `form:"waktu_persalinan_kala_1" json:"waktu_persalinan_kala_1"`
	WaktuPersalinanKala2  string `form:"waktu_persalinan_kala_2" json:"waktu_persalinan_kala_2"`
	WaktuPersalinanKala3  string `form:"waktu_persalinan_kala_3" json:"waktu_persalinan_kala_3"`
	WaktuPersalinanJumlah string `form:"waktu_persalinan_jumlah" json:"waktu_persalinan_jumlah"`
	Perineum              string `form:"perineum" json:"perineum"`
	JahitanLuar1          string `form:"jahitan_luar_1" json:"jahitan_luar_1"`
	JahitanLuar2          string `form:"jahitan_luar_2" json:"jahitan_luar_2"`
	JahitanDalam1         string `form:"jahitan_dalam_1" json:"jahitan_dalam_1"`
	JahitanDalam2         string `form:"jahitan_dalam_2" json:"jahitan_dalam_2"`
	Anak                  string `form:"anak" json:"anak"`
	StatusLahir           string `form:"status_lahir" json:"status_lahir"`
	ApgarScore            string `form:"apgar_score" json:"apgar_score"`
	Bb                    string `form:"bb" json:"bb"`
	Pb                    string `form:"pb" json:"pb"`
	Kelainan              string `form:"kelainan" json:"kelainan"`
	Ketuban               string `form:"ketuban" json:"ketuban"`
	Placenta              string `form:"placenta" json:"placenta"`
	Ukuran                string `form:"ukuran" json:"ukuran"`
	TaliPusat             string `form:"tali_pusat" json:"tali_pusat"`
	Insertio              string `form:"insertio" json:"insertio"`
	DarahKeluarKala1      string `form:"darah_keluar_kala_1" json:"darah_keluar_kala_1"`
	DarahKeluarKala2      string `form:"darah_keluar_kala_2" json:"darah_keluar_kala_2"`
	DarahKeluarKala3      string `form:"darah_keluar_kala_3" json:"darah_keluar_kala_3"`
	DarahKeluarKala4      string `form:"darah_keluar_kala_4" json:"darah_keluar_kala_4"`
	DarahKeluarJumlah     string `form:"darah_keluar_jumlah" json:"darah_keluar_jumlah"`
	KondisiUmum           string `form:"kondisi_umum" json:"kondisi_umum"`
	Td                    string `form:"td" json:"td"`
	Nadi                  string `form:"nadi" json:"nadi"`
	Rr                    string `form:"rr" json:"rr"`
	Suhu                  string `form:"suhu" json:"suhu"`
	KontraksiUterus       string `form:"kontraksi_uterus" json:"kontraksi_uterus"`
	Ppv                   string `form:"ppv" json:"ppv"`
	Pengobatan            string `form:"pengobatan" json:"pengobatan"`
}

func catatanPersalinanRules() map[string]any {
	rules := map[string]any{
		"mulai":                   "required|date",
		"selesai":                 "required|date",
		"kd_dokter":               "required|string|max_len:20",
		"nip":                     "required|string|max_len:20",
		"catatan":                 "string",
		"waktu_persalinan_kala_1": "string|max_len:5",
		"waktu_persalinan_kala_2": "string|max_len:5",
		"waktu_persalinan_kala_3": "string|max_len:5",
		"waktu_persalinan_jumlah": "string|max_len:5",
		"perineum":                "in:Utuh,Rupture,Episiotomi",
		"jahitan_luar_1":          "string|max_len:5",
		"jahitan_luar_2":          "string|max_len:5",
		"jahitan_dalam_1":         "string|max_len:5",
		"jahitan_dalam_2":         "string|max_len:5",
		"anak":                    "in:Laki-laki,Perempuan",
		"status_lahir":            "in:Hidup,Mati",
		"apgar_score":             "string|max_len:20",
		"bb":                      "string|max_len:5",
		"pb":                      "string|max_len:5",
		"kelainan":                "string|max_len:100",
		"ketuban":                 "string|max_len:20",
		"placenta":                "string|max_len:20",
		"ukuran":                  "string|max_len:5",
		"tali_pusat":              "string|max_len:5",
		"insertio":                "string|max_len:20",
		"darah_keluar_kala_1":     "string|max_len:5",
		"darah_keluar_kala_2":     "string|max_len:5",
		"darah_keluar_kala_3":     "string|max_len:5",
		"darah_keluar_kala_4":     "string|max_len:5",
		"darah_keluar_jumlah":     "string|max_len:5",
		"kondisi_umum":            "string|max_len:100",
		"td":                      "string|max_len:8",
		"nadi":                    "string|max_len:5",
		"rr":                      "string|max_len:5",
		"suhu":                    "string|max_len:5",
		"kontraksi_uterus":        "string|max_len:100",
		"ppv":                     "string|max_len:100",
		"pengobatan":              "string|max_len:600",
	}
	return rules
}

// CatatanPersalinanStore simpan catatan persalinan; kolom waktu kunci kosong = sekarang.
type CatatanPersalinanStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	CatatanPersalinanData
}

func (r *CatatanPersalinanStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanPersalinanStore) Rules(ctx http.Context) map[string]any {
	rules := catatanPersalinanRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *CatatanPersalinanStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *CatatanPersalinanStore) Payload() CatatanPersalinanData { return r.CatatanPersalinanData }

func (r *CatatanPersalinanStore) DetailValues() map[string][]string { return nil }

// CatatanPersalinanUpdate ubah catatan persalinan (PUT); kunci lewat query string.
type CatatanPersalinanUpdate struct {
	CatatanPersalinanData
}

func (r *CatatanPersalinanUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanPersalinanUpdate) Rules(ctx http.Context) map[string]any {
	return catatanPersalinanRules()
}

func (r *CatatanPersalinanUpdate) Payload() CatatanPersalinanData { return r.CatatanPersalinanData }

func (r *CatatanPersalinanUpdate) DetailValues() map[string][]string { return nil }
