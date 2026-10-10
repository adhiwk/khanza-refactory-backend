package riwayatpasien

import (
	"github.com/goravel/framework/contracts/http"
)

type PersalinanRequest struct {
	NoRkmMedis       string `form:"no_rkm_medis" json:"no_rkm_medis"`
	TglThn           string `form:"tgl_thn" json:"tgl_thn"`
	TempatPersalinan string `form:"tempat_persalinan" json:"tempat_persalinan"`
	UsiaHamil        string `form:"usia_hamil" json:"usia_hamil"`
	JenisPersalinan  string `form:"jenis_persalinan" json:"jenis_persalinan"`
	Penolong         string `form:"penolong" json:"penolong"`
	Penyulit         string `form:"penyulit" json:"penyulit"`
	Jk               string `form:"jk" json:"jk"`
	Bbpb             string `form:"bbpb" json:"bbpb"`
	Keadaan          string `form:"keadaan" json:"keadaan"`
}

func (r *PersalinanRequest) Authorize(ctx http.Context) error { return nil }

func (r *PersalinanRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rkm_medis":      "required|string|max_len:15",
		"tgl_thn":           "required|string|max_len:12",
		"tempat_persalinan": "string|max_len:30",
		"usia_hamil":        "string|max_len:20",
		"jenis_persalinan":  "string|max_len:20",
		"penolong":          "string|max_len:30",
		"penyulit":          "string|max_len:40",
		"jk":                "in:L,P,-",
		"bbpb":              "string|max_len:10",
		"keadaan":           "string|max_len:40",
	}
}

type ImunisasiRequest struct {
	NoRkmMedis    string `form:"no_rkm_medis" json:"no_rkm_medis"`
	KodeImunisasi string `form:"kode_imunisasi" json:"kode_imunisasi"`
	NoImunisasi   int    `form:"no_imunisasi" json:"no_imunisasi"`
}

func (r *ImunisasiRequest) Authorize(ctx http.Context) error { return nil }

func (r *ImunisasiRequest) Rules(ctx http.Context) map[string]any {
	return map[string]any{
		"no_rkm_medis":   "required|string|max_len:15",
		"kode_imunisasi": "required|string|max_len:3",
		"no_imunisasi":   "required|int|min:1|max:127",
	}
}
