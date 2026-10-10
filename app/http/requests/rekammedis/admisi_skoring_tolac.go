package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// AdmisiSkoringTolacData isian admisi skoring TOLAC.
type AdmisiSkoringTolacData struct {
	Tanggal                  string `form:"tanggal" json:"tanggal"`
	HisFrekuensi             string `form:"his_frekuensi" json:"his_frekuensi"`
	HisDurasiDetik           string `form:"his_durasi_detik" json:"his_durasi_detik"`
	Djj                      string `form:"djj" json:"djj"`
	PembukaanCm              string `form:"pembukaan_cm" json:"pembukaan_cm"`
	PendataranPersen         string `form:"pendataran_persen" json:"pendataran_persen"`
	PenurunanKepala          string `form:"penurunan_kepala" json:"penurunan_kepala"`
	PilihanUsia              string `form:"pilihan_usia" json:"pilihan_usia"`
	SkorUsia                 string `form:"skor_usia" json:"skor_usia"`
	PilihanRiwayatPervaginam string `form:"pilihan_riwayat_pervaginam" json:"pilihan_riwayat_pervaginam"`
	SkorRiwayatPervaginam    string `form:"skor_riwayat_pervaginam" json:"skor_riwayat_pervaginam"`
	PilihanIndikasiSc        string `form:"pilihan_indikasi_sc" json:"pilihan_indikasi_sc"`
	SkorIndikasiSc           string `form:"skor_indikasi_sc" json:"skor_indikasi_sc"`
	PilihanPendataran        string `form:"pilihan_pendataran" json:"pilihan_pendataran"`
	SkorPendataran           string `form:"skor_pendataran" json:"skor_pendataran"`
	PilihanPembukaan         string `form:"pilihan_pembukaan" json:"pilihan_pembukaan"`
	SkorPembukaan            string `form:"skor_pembukaan" json:"skor_pembukaan"`
	TotalSkor                string `form:"total_skor" json:"total_skor"`
	PeluangVbac              string `form:"peluang_vbac" json:"peluang_vbac"`
	Keputusan                string `form:"keputusan" json:"keputusan"`
	Keterangan               string `form:"keterangan" json:"keterangan"`
	KdDokter                 string `form:"kd_dokter" json:"kd_dokter"`
}

func admisiSkoringTolacRules() map[string]any {
	rules := map[string]any{
		"tanggal":                    "date",
		"his_frekuensi":              "string|max_len:2",
		"his_durasi_detik":           "string|max_len:5",
		"djj":                        "string|max_len:5",
		"pembukaan_cm":               "string|max_len:2",
		"pendataran_persen":          "string|max_len:3",
		"penurunan_kepala":           "string|max_len:20",
		"pilihan_usia":               "in:< 40 Tahun,>= 40 Tahun",
		"skor_usia":                  "string|max_len:1",
		"pilihan_riwayat_pervaginam": "in:Sebelum & Sesudah SC,Sesudah SC Saja,Sebelum SC Saja,Tidak Pernah",
		"skor_riwayat_pervaginam":    "string|max_len:1",
		"pilihan_indikasi_sc":        "in:Bukan Kegagalan Kemajuan Persalinan,Kegagalan Kemajuan Persalinan / CPD-distosia",
		"skor_indikasi_sc":           "string|max_len:1",
		"pilihan_pendataran":         "in:>75%,25â75%,<25%",
		"skor_pendataran":            "string|max_len:1",
		"pilihan_pembukaan":          "in:>=4 cm,<4 cm",
		"skor_pembukaan":             "string|max_len:1",
		"total_skor":                 "string|max_len:2",
		"peluang_vbac":               "string|max_len:6",
		"keputusan":                  "in:Lanjut TOLAC,Lakukan SC",
		"keterangan":                 "string|max_len:100",
		"kd_dokter":                  "string|max_len:15",
	}
	return rules
}

// AdmisiSkoringTolacStore simpan admisi skoring TOLAC; kolom waktu kunci kosong = sekarang.
type AdmisiSkoringTolacStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	AdmisiSkoringTolacData
}

func (r *AdmisiSkoringTolacStore) Authorize(ctx http.Context) error { return nil }

func (r *AdmisiSkoringTolacStore) Rules(ctx http.Context) map[string]any {
	rules := admisiSkoringTolacRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *AdmisiSkoringTolacStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *AdmisiSkoringTolacStore) Payload() AdmisiSkoringTolacData { return r.AdmisiSkoringTolacData }

func (r *AdmisiSkoringTolacStore) DetailValues() map[string][]string { return nil }

// AdmisiSkoringTolacUpdate ubah admisi skoring TOLAC (PUT); kunci lewat query string.
type AdmisiSkoringTolacUpdate struct {
	AdmisiSkoringTolacData
}

func (r *AdmisiSkoringTolacUpdate) Authorize(ctx http.Context) error { return nil }

func (r *AdmisiSkoringTolacUpdate) Rules(ctx http.Context) map[string]any {
	return admisiSkoringTolacRules()
}

func (r *AdmisiSkoringTolacUpdate) Payload() AdmisiSkoringTolacData { return r.AdmisiSkoringTolacData }

func (r *AdmisiSkoringTolacUpdate) DetailValues() map[string][]string { return nil }
