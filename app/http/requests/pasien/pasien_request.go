package pasien

import (
	"github.com/goravel/framework/contracts/http"
)

// PasienData field yang dipakai bersama oleh store & update.
// Tanggal dikirim sebagai string "YYYY-MM-DD"; string kosong pada kolom nullable disimpan sebagai NULL.
type PasienData struct {
	NmPasien         string `form:"nm_pasien" json:"nm_pasien"`
	NoKtp            string `form:"no_ktp" json:"no_ktp"`
	Jk               string `form:"jk" json:"jk"`
	TmpLahir         string `form:"tmp_lahir" json:"tmp_lahir"`
	TglLahir         string `form:"tgl_lahir" json:"tgl_lahir"`
	NmIbu            string `form:"nm_ibu" json:"nm_ibu"`
	Alamat           string `form:"alamat" json:"alamat"`
	GolDarah         string `form:"gol_darah" json:"gol_darah"`
	Pekerjaan        string `form:"pekerjaan" json:"pekerjaan"`
	SttsNikah        string `form:"stts_nikah" json:"stts_nikah"`
	Agama            string `form:"agama" json:"agama"`
	TglDaftar        string `form:"tgl_daftar" json:"tgl_daftar"`
	NoTlp            string `form:"no_tlp" json:"no_tlp"`
	Umur             string `form:"umur" json:"umur"`
	Pnd              string `form:"pnd" json:"pnd"`
	Keluarga         string `form:"keluarga" json:"keluarga"`
	NamaKeluarga     string `form:"namakeluarga" json:"namakeluarga"`
	KdPj             string `form:"kd_pj" json:"kd_pj"`
	NoPeserta        string `form:"no_peserta" json:"no_peserta"`
	KdKel            int    `form:"kd_kel" json:"kd_kel"`
	KdKec            int    `form:"kd_kec" json:"kd_kec"`
	KdKab            int    `form:"kd_kab" json:"kd_kab"`
	PekerjaanPj      string `form:"pekerjaanpj" json:"pekerjaanpj"`
	AlamatPj         string `form:"alamatpj" json:"alamatpj"`
	KelurahanPj      string `form:"kelurahanpj" json:"kelurahanpj"`
	KecamatanPj      string `form:"kecamatanpj" json:"kecamatanpj"`
	KabupatenPj      string `form:"kabupatenpj" json:"kabupatenpj"`
	PerusahaanPasien string `form:"perusahaan_pasien" json:"perusahaan_pasien"`
	SukuBangsa       int    `form:"suku_bangsa" json:"suku_bangsa"`
	BahasaPasien     int    `form:"bahasa_pasien" json:"bahasa_pasien"`
	CacatFisik       int    `form:"cacat_fisik" json:"cacat_fisik"`
	Email            string `form:"email" json:"email"`
	Nip              string `form:"nip" json:"nip"`
	KdProp           int    `form:"kd_prop" json:"kd_prop"`
	PropinsiPj       string `form:"propinsipj" json:"propinsipj"`
}

func baseRules() map[string]any {
	return map[string]any{
		"nm_pasien":         "required|string|max_len:40",
		"no_ktp":            "string|max_len:20",
		"jk":                "required|in:L,P",
		"tmp_lahir":         "string|max_len:15",
		"tgl_lahir":         "required|date",
		"nm_ibu":            "required|string|max_len:40",
		"alamat":            "string|max_len:200",
		"gol_darah":         "in:A,B,O,AB,-",
		"pekerjaan":         "string|max_len:60",
		"stts_nikah":        "in:BELUM MENIKAH,MENIKAH,JANDA,DUDHA,JOMBLO",
		"agama":             "string|max_len:12",
		"tgl_daftar":        "date",
		"no_tlp":            "string|max_len:40",
		"umur":              "string|max_len:30",
		"pnd":               "required|in:TS,TK,SD,SMP,SMA,SLTA/SEDERAJAT,D1,D2,D3,D4,S1,S2,S3,-",
		"keluarga":          "in:AYAH,IBU,ISTRI,SUAMI,SAUDARA,ANAK,DIRI SENDIRI,LAIN-LAIN",
		"namakeluarga":      "string|max_len:50",
		"kd_pj":             "required|string|max_len:3",
		"no_peserta":        "string|max_len:25",
		"kd_kel":            "int",
		"kd_kec":            "int",
		"kd_kab":            "int",
		"pekerjaanpj":       "string|max_len:35",
		"alamatpj":          "string|max_len:100",
		"kelurahanpj":       "string|max_len:60",
		"kecamatanpj":       "string|max_len:60",
		"kabupatenpj":       "string|max_len:60",
		"perusahaan_pasien": "string|max_len:8",
		"suku_bangsa":       "int",
		"bahasa_pasien":     "int",
		"cacat_fisik":       "int",
		"email":             "email|max_len:50",
		"nip":               "string|max_len:30",
		"kd_prop":           "int",
		"propinsipj":        "string|max_len:30",
	}
}

type StorePasienRequest struct {
	NoRkmMedis string `form:"no_rkm_medis" json:"no_rkm_medis"`
	PasienData
}

func (r *StorePasienRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *StorePasienRequest) Rules(ctx http.Context) map[string]any {
	rules := baseRules()
	rules["no_rkm_medis"] = "required|string|max_len:15"
	return rules
}

func (r *StorePasienRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}

// UpdatePasienRequest mengganti seluruh data pasien (PUT); no_rkm_medis diambil dari route.
type UpdatePasienRequest struct {
	PasienData
}

func (r *UpdatePasienRequest) Authorize(ctx http.Context) error {
	return nil
}

func (r *UpdatePasienRequest) Rules(ctx http.Context) map[string]any {
	return baseRules()
}

func (r *UpdatePasienRequest) Filters(ctx http.Context) map[string]string {
	return map[string]string{}
}
