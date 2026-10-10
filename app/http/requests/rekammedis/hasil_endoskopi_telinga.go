package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilEndoskopiTelingaData isian hasil endoskopi telinga.
type HasilEndoskopiTelingaData struct {
	Tanggal                                string `form:"tanggal" json:"tanggal"`
	KdDokter                               string `form:"kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis                         string `form:"diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari                            string `form:"kiriman_dari" json:"kiriman_dari"`
	BentukLiangTelingaKanan                string `form:"bentuk_liang_telinga_kanan" json:"bentuk_liang_telinga_kanan"`
	BentukLiangTelingaKiri                 string `form:"bentuk_liang_telinga_kiri" json:"bentuk_liang_telinga_kiri"`
	KondisiLiangTelingaKanan               string `form:"kondisi_liang_telinga_kanan" json:"kondisi_liang_telinga_kanan"`
	KeteranganKondisiLiangTelingaKanan     string `form:"keterangan_kondisi_liang_telinga_kanan" json:"keterangan_kondisi_liang_telinga_kanan"`
	KondisiLiangTelingaKiri                string `form:"kondisi_liang_telinga_kiri" json:"kondisi_liang_telinga_kiri"`
	KeteranganKondisiLiangTelingaKiri      string `form:"keterangan_kondisi_liang_telinga_kiri" json:"keterangan_kondisi_liang_telinga_kiri"`
	MembranTimpaniIntakKanan               string `form:"membran_timpani_intak_kanan" json:"membran_timpani_intak_kanan"`
	MembranTimpaniIntakKiri                string `form:"membran_timpani_intak_kiri" json:"membran_timpani_intak_kiri"`
	MembranTimpaniPerforasiKanan           string `form:"membran_timpani_perforasi_kanan" json:"membran_timpani_perforasi_kanan"`
	KeteranganMembranTimpaniPerforasiKanan string `form:"keterangan_membran_timpani_perforasi_kanan" json:"keterangan_membran_timpani_perforasi_kanan"`
	MembranTimpaniPerforasiKiri            string `form:"membran_timpani_perforasi_kiri" json:"membran_timpani_perforasi_kiri"`
	KeteranganMembranTimpaniPerforasiKiri  string `form:"keterangan_membran_timpani_perforasi_kiri" json:"keterangan_membran_timpani_perforasi_kiri"`
	KavumTimpaniMukosaKanan                string `form:"kavum_timpani_mukosa_kanan" json:"kavum_timpani_mukosa_kanan"`
	KavumTimpaniMukosaKiri                 string `form:"kavum_timpani_mukosa_kiri" json:"kavum_timpani_mukosa_kiri"`
	KavumTimpaniOsikelKanan                string `form:"kavum_timpani_osikel_kanan" json:"kavum_timpani_osikel_kanan"`
	KavumTimpaniOsikelKiri                 string `form:"kavum_timpani_osikel_kiri" json:"kavum_timpani_osikel_kiri"`
	KavumTimpaniIsthmusKanan               string `form:"kavum_timpani_isthmus_kanan" json:"kavum_timpani_isthmus_kanan"`
	KavumTimpaniIsthmusKiri                string `form:"kavum_timpani_isthmus_kiri" json:"kavum_timpani_isthmus_kiri"`
	KavumTimpaniAnteriorKanan              string `form:"kavum_timpani_anterior_kanan" json:"kavum_timpani_anterior_kanan"`
	KavumTimpaniAnteriorKiri               string `form:"kavum_timpani_anterior_kiri" json:"kavum_timpani_anterior_kiri"`
	KavumTimpaniPosteriorKanan             string `form:"kavum_timpani_posterior_kanan" json:"kavum_timpani_posterior_kanan"`
	KavumTimpaniPosteriorKiri              string `form:"kavum_timpani_posterior_kiri" json:"kavum_timpani_posterior_kiri"`
	LainlainKanan                          string `form:"lainlain_kanan" json:"lainlain_kanan"`
	LainlainKiri                           string `form:"lainlain_kiri" json:"lainlain_kiri"`
	Kesimpulan                             string `form:"kesimpulan" json:"kesimpulan"`
	Anjuran                                string `form:"anjuran" json:"anjuran"`
}

func hasilEndoskopiTelingaRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                    "required|date",
		"kd_dokter":                                  "required|string|max_len:20",
		"diagnosa_klinis":                            "string|max_len:50",
		"kiriman_dari":                               "string|max_len:50",
		"bentuk_liang_telinga_kanan":                 "in:Lapang,Sempit,Destruksi",
		"bentuk_liang_telinga_kiri":                  "in:Lapang,Sempit,Destruksi",
		"kondisi_liang_telinga_kanan":                "in:Serumen,Sekret,Jamur,Kolesteatoma,Massa/Jaringan,Benda Asing,Lainnya",
		"keterangan_kondisi_liang_telinga_kanan":     "string|max_len:30",
		"kondisi_liang_telinga_kiri":                 "in:Serumen,Sekret,Jamur,Kolesteatoma,Massa/Jaringan,Benda Asing,Lainnya",
		"keterangan_kondisi_liang_telinga_kiri":      "string|max_len:30",
		"membran_timpani_intak_kanan":                "in:Normal,Hiperemis,Bulging,Retraksi,Sklerotik",
		"membran_timpani_intak_kiri":                 "in:Normal,Hiperemis,Bulging,Retraksi,Sklerotik",
		"membran_timpani_perforasi_kanan":            "in:Sentral,Atik,Marginal,Lainnya",
		"keterangan_membran_timpani_perforasi_kanan": "string|max_len:30",
		"membran_timpani_perforasi_kiri":             "in:Sentral,Atik,Marginal,Lainnya",
		"keterangan_membran_timpani_perforasi_kiri":  "string|max_len:30",
		"kavum_timpani_mukosa_kanan":                 "string|max_len:40",
		"kavum_timpani_mukosa_kiri":                  "string|max_len:40",
		"kavum_timpani_osikel_kanan":                 "string|max_len:40",
		"kavum_timpani_osikel_kiri":                  "string|max_len:40",
		"kavum_timpani_isthmus_kanan":                "string|max_len:40",
		"kavum_timpani_isthmus_kiri":                 "string|max_len:40",
		"kavum_timpani_anterior_kanan":               "string|max_len:40",
		"kavum_timpani_anterior_kiri":                "string|max_len:40",
		"kavum_timpani_posterior_kanan":              "string|max_len:40",
		"kavum_timpani_posterior_kiri":               "string|max_len:40",
		"lainlain_kanan":                             "string|max_len:100",
		"lainlain_kiri":                              "string|max_len:100",
		"kesimpulan":                                 "string|max_len:300",
		"anjuran":                                    "string|max_len:300",
	}
	return rules
}

// HasilEndoskopiTelingaStore simpan hasil endoskopi telinga; kolom waktu kunci kosong = sekarang.
type HasilEndoskopiTelingaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	HasilEndoskopiTelingaData
}

func (r *HasilEndoskopiTelingaStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilEndoskopiTelingaStore) Rules(ctx http.Context) map[string]any {
	rules := hasilEndoskopiTelingaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *HasilEndoskopiTelingaStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *HasilEndoskopiTelingaStore) Payload() HasilEndoskopiTelingaData {
	return r.HasilEndoskopiTelingaData
}

func (r *HasilEndoskopiTelingaStore) DetailValues() map[string][]string { return nil }

// HasilEndoskopiTelingaUpdate ubah hasil endoskopi telinga (PUT); kunci lewat query string.
type HasilEndoskopiTelingaUpdate struct {
	HasilEndoskopiTelingaData
}

func (r *HasilEndoskopiTelingaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilEndoskopiTelingaUpdate) Rules(ctx http.Context) map[string]any {
	return hasilEndoskopiTelingaRules()
}

func (r *HasilEndoskopiTelingaUpdate) Payload() HasilEndoskopiTelingaData {
	return r.HasilEndoskopiTelingaData
}

func (r *HasilEndoskopiTelingaUpdate) DetailValues() map[string][]string { return nil }
