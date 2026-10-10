package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// HasilPemeriksaanUsgData isian hasil pemeriksaan USG.
type HasilPemeriksaanUsgData struct {
	Tanggal               string `form:"tanggal" json:"tanggal"`
	KdDokter              string `form:"kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis        string `form:"diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari           string `form:"kiriman_dari" json:"kiriman_dari"`
	Hta                   string `form:"hta" json:"hta"`
	KantongGestasi        string `form:"kantong_gestasi" json:"kantong_gestasi"`
	UkuranBokongkepala    string `form:"ukuran_bokongkepala" json:"ukuran_bokongkepala"`
	JenisPrestasi         string `form:"jenis_prestasi" json:"jenis_prestasi"`
	DiameterBiparietal    string `form:"diameter_biparietal" json:"diameter_biparietal"`
	PanjangFemur          string `form:"panjang_femur" json:"panjang_femur"`
	LingkarAbdomen        string `form:"lingkar_abdomen" json:"lingkar_abdomen"`
	TafsiranBeratJanin    string `form:"tafsiran_berat_janin" json:"tafsiran_berat_janin"`
	UsiaKehamilan         string `form:"usia_kehamilan" json:"usia_kehamilan"`
	PlasentaBerimplatansi string `form:"plasenta_berimplatansi" json:"plasenta_berimplatansi"`
	DerajatMaturitas      string `form:"derajat_maturitas" json:"derajat_maturitas"`
	JumlahAirKetuban      string `form:"jumlah_air_ketuban" json:"jumlah_air_ketuban"`
	IndekCairanKetuban    string `form:"indek_cairan_ketuban" json:"indek_cairan_ketuban"`
	KelainanKongenital    string `form:"kelainan_kongenital" json:"kelainan_kongenital"`
	PeluangSex            string `form:"peluang_sex" json:"peluang_sex"`
	Kesimpulan            string `form:"kesimpulan" json:"kesimpulan"`
}

func hasilPemeriksaanUsgRules() map[string]any {
	rules := map[string]any{
		"tanggal":                "required|date",
		"kd_dokter":              "required|string|max_len:20",
		"diagnosa_klinis":        "string|max_len:50",
		"kiriman_dari":           "string|max_len:50",
		"hta":                    "string|max_len:40",
		"kantong_gestasi":        "string|max_len:6",
		"ukuran_bokongkepala":    "string|max_len:6",
		"jenis_prestasi":         "string|max_len:30",
		"diameter_biparietal":    "string|max_len:6",
		"panjang_femur":          "string|max_len:6",
		"lingkar_abdomen":        "string|max_len:6",
		"tafsiran_berat_janin":   "string|max_len:6",
		"usia_kehamilan":         "string|max_len:15",
		"plasenta_berimplatansi": "string|max_len:50",
		"derajat_maturitas":      "in:0,1,2,3",
		"jumlah_air_ketuban":     "in:Cukup,Berkurang",
		"indek_cairan_ketuban":   "string|max_len:40",
		"kelainan_kongenital":    "string|max_len:60",
		"peluang_sex":            "in:Laki-laki,Perempuan,-",
		"kesimpulan":             "string|max_len:200",
	}
	return rules
}

// HasilPemeriksaanUsgStore simpan hasil pemeriksaan USG; kolom waktu kunci kosong = sekarang.
type HasilPemeriksaanUsgStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	HasilPemeriksaanUsgData
}

func (r *HasilPemeriksaanUsgStore) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanUsgStore) Rules(ctx http.Context) map[string]any {
	rules := hasilPemeriksaanUsgRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *HasilPemeriksaanUsgStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *HasilPemeriksaanUsgStore) Payload() HasilPemeriksaanUsgData {
	return r.HasilPemeriksaanUsgData
}

func (r *HasilPemeriksaanUsgStore) DetailValues() map[string][]string { return nil }

// HasilPemeriksaanUsgUpdate ubah hasil pemeriksaan USG (PUT); kunci lewat query string.
type HasilPemeriksaanUsgUpdate struct {
	HasilPemeriksaanUsgData
}

func (r *HasilPemeriksaanUsgUpdate) Authorize(ctx http.Context) error { return nil }

func (r *HasilPemeriksaanUsgUpdate) Rules(ctx http.Context) map[string]any {
	return hasilPemeriksaanUsgRules()
}

func (r *HasilPemeriksaanUsgUpdate) Payload() HasilPemeriksaanUsgData {
	return r.HasilPemeriksaanUsgData
}

func (r *HasilPemeriksaanUsgUpdate) DetailValues() map[string][]string { return nil }
