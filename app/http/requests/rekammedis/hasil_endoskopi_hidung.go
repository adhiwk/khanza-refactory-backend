package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilEndoskopiHidungData isian hasil endoskopi hidung.
type HasilEndoskopiHidungData struct {
	Tanggal            string `form:"tanggal" json:"tanggal"`
	KdDokter           string `form:"kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis     string `form:"diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari        string `form:"kiriman_dari" json:"kiriman_dari"`
	KondisiHidungKanan string `form:"kondisi_hidung_kanan" json:"kondisi_hidung_kanan"`
	KondisiHidungKiri  string `form:"kondisi_hidung_kiri" json:"kondisi_hidung_kiri"`
	KavumNasiKanan     string `form:"kavum_nasi_kanan" json:"kavum_nasi_kanan"`
	KavumNasiKiri      string `form:"kavum_nasi_kiri" json:"kavum_nasi_kiri"`
	KonkaInferiorKanan string `form:"konka_inferior_kanan" json:"konka_inferior_kanan"`
	KonkaInferiorKiri  string `form:"konka_inferior_kiri" json:"konka_inferior_kiri"`
	MeatusMediusKanan  string `form:"meatus_medius_kanan" json:"meatus_medius_kanan"`
	MeatusMediusKiri   string `form:"meatus_medius_kiri" json:"meatus_medius_kiri"`
	SeptumKanan        string `form:"septum_kanan" json:"septum_kanan"`
	SeptumKiri         string `form:"septum_kiri" json:"septum_kiri"`
	NasofaringKanan    string `form:"nasofaring_kanan" json:"nasofaring_kanan"`
	NasofaringKiri     string `form:"nasofaring_kiri" json:"nasofaring_kiri"`
	LainlainKanan      string `form:"lainlain_kanan" json:"lainlain_kanan"`
	LainlainKiri       string `form:"lainlain_kiri" json:"lainlain_kiri"`
	Kesimpulan         string `form:"kesimpulan" json:"kesimpulan"`
}

func hasilEndoskopiHidungRules() map[string]any {
	rules := map[string]any{
		"tanggal":              "required|date",
		"kd_dokter":            "required|string|max_len:20",
		"diagnosa_klinis":      "string|max_len:50",
		"kiriman_dari":         "string|max_len:50",
		"kondisi_hidung_kanan": "in:Lapang,Sempit,Mukosa Edema",
		"kondisi_hidung_kiri":  "in:Lapang,Sempit,Mukosa Edema",
		"kavum_nasi_kanan":     "in:Mukosa Pucat,Mukosa Hiperemis,Massa,Polip",
		"kavum_nasi_kiri":      "in:Mukosa Pucat,Mukosa Hiperemis,Massa,Polip",
		"konka_inferior_kanan": "in:Eutrofi,Hipertrofi,Atrofi",
		"konka_inferior_kiri":  "in:Eutrofi,Hipertrofi,Atrofi",
		"meatus_medius_kanan":  "in:Terbuka,Tertutup,Mukosa Edema,Polip,Sekret",
		"meatus_medius_kiri":   "in:Terbuka,Tertutup,Mukosa Edema,Polip,Sekret",
		"septum_kanan":         "in:Lurus,Deviasi,Spina",
		"septum_kiri":          "in:Lurus,Deviasi,Spina",
		"nasofaring_kanan":     "in:Normal,Adenoid,Keradangan,Massa",
		"nasofaring_kiri":      "in:Normal,Adenoid,Keradangan,Massa",
		"lainlain_kanan":       "string|max_len:100",
		"lainlain_kiri":        "string|max_len:100",
		"kesimpulan":           "string|max_len:300",
	}
	return rules
}

// HasilEndoskopiHidungStore simpan hasil endoskopi hidung; kolom waktu kunci kosong = sekarang.
type HasilEndoskopiHidungStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	HasilEndoskopiHidungData
}

func (r *HasilEndoskopiHidungStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilEndoskopiHidungStore) Rules(ctx http.Context) map[string]any {
	rules := hasilEndoskopiHidungRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *HasilEndoskopiHidungStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *HasilEndoskopiHidungStore) Payload() HasilEndoskopiHidungData {
	return r.HasilEndoskopiHidungData
}

func (r *HasilEndoskopiHidungStore) DetailValues() map[string][]string { return nil }

// HasilEndoskopiHidungUpdate ubah hasil endoskopi hidung (PUT); kunci lewat query string.
type HasilEndoskopiHidungUpdate struct {
	HasilEndoskopiHidungData
}

func (r *HasilEndoskopiHidungUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilEndoskopiHidungUpdate) Rules(ctx http.Context) map[string]any {
	return hasilEndoskopiHidungRules()
}

func (r *HasilEndoskopiHidungUpdate) Payload() HasilEndoskopiHidungData {
	return r.HasilEndoskopiHidungData
}

func (r *HasilEndoskopiHidungUpdate) DetailValues() map[string][]string { return nil }
