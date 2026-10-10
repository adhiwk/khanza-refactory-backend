package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningPerilakuMerokokSekolahRemajaData isian skrining merokok usia sekolah remaja.
type SkriningPerilakuMerokokSekolahRemajaData struct {
	Tanggal                                         string `form:"tanggal" json:"tanggal"`
	KdSekolah                                       string `form:"kd_sekolah" json:"kd_sekolah"`
	Kelas                                           string `form:"kelas" json:"kelas"`
	ApakahAndaMerokok                               string `form:"apakah_anda_merokok" json:"apakah_anda_merokok"`
	JumlahBatangRokok                               string `form:"jumlah_batang_rokok" json:"jumlah_batang_rokok"`
	JumlahBatangRokokHariminggu                     string `form:"jumlah_batang_rokok_hariminggu" json:"jumlah_batang_rokok_hariminggu"`
	JenisRokokYangDigunakan                         string `form:"jenis_rokok_yang_digunakan" json:"jenis_rokok_yang_digunakan"`
	JenisRokokYangDigunakanKeterangan               string `form:"jenis_rokok_yang_digunakan_keterangan" json:"jenis_rokok_yang_digunakan_keterangan"`
	UsiaMulaiMerokok                                string `form:"usia_mulai_merokok" json:"usia_mulai_merokok"`
	AlasanMulaiMerokok                              string `form:"alasan_mulai_merokok" json:"alasan_mulai_merokok"`
	AlasanMulaiMerokokKeterangan                    string `form:"alasan_mulai_merokok_keterangan" json:"alasan_mulai_merokok_keterangan"`
	SudahBerapaLamaMerokok                          string `form:"sudah_berapa_lama_merokok" json:"sudah_berapa_lama_merokok"`
	BagaimanaBiasanyaMendapatkanRokok               string `form:"bagaimana_biasanya_mendapatkan_rokok" json:"bagaimana_biasanya_mendapatkan_rokok"`
	BagaimanaBiasanyaMendapatkanRokokKeterangan     string `form:"bagaimana_biasanya_mendapatkan_rokok_keterangan" json:"bagaimana_biasanya_mendapatkan_rokok_keterangan"`
	KeinginanBerhentiMerokok                        string `form:"keinginan_berhenti_merokok" json:"keinginan_berhenti_merokok"`
	AlasanUtamaBerhentiMerokok                      string `form:"alasan_utama_berhenti_merokok" json:"alasan_utama_berhenti_merokok"`
	AlasanUtamaBerhentiMerokokKeterangan            string `form:"alasan_utama_berhenti_merokok_keterangan" json:"alasan_utama_berhenti_merokok_keterangan"`
	TahuDampakKesehatanMerokok                      string `form:"tahu_dampak_kesehatan_merokok" json:"tahu_dampak_kesehatan_merokok"`
	DampakKesehatanDariMerokokYangDiketahui         string `form:"dampak_kesehatan_dari_merokok_yang_diketahui" json:"dampak_kesehatan_dari_merokok_yang_diketahui"`
	TahuMerokokPintuMasukNarkoba                    string `form:"tahu_merokok_pintu_masuk_narkoba" json:"tahu_merokok_pintu_masuk_narkoba"`
	MelihatOrangMerokokDiSekolah                    string `form:"melihat_orang_merokok_di_sekolah" json:"melihat_orang_merokok_di_sekolah"`
	OrangYangPalingSeringMerokokDisekolah           string `form:"orang_yang_paling_sering_merokok_disekolah" json:"orang_yang_paling_sering_merokok_disekolah"`
	OrangYangPalingSeringMerokokDisekolahKeterangan string `form:"orang_yang_paling_sering_merokok_disekolah_keterangan" json:"orang_yang_paling_sering_merokok_disekolah_keterangan"`
	AdaAnggotaKeluargaDiRumahYangMerokok            string `form:"ada_anggota_keluarga_di_rumah_yang_merokok" json:"ada_anggota_keluarga_di_rumah_yang_merokok"`
	TemanDekatBanyakyangMerokok                     string `form:"teman_dekat_banyakyang_merokok" json:"teman_dekat_banyakyang_merokok"`
	DilakukanPemeriksaanKadarCoPernapasan           string `form:"dilakukan_pemeriksaan_kadar_co_pernapasan" json:"dilakukan_pemeriksaan_kadar_co_pernapasan"`
	HasilPemeriksaanCoPernapasan                    string `form:"hasil_pemeriksaan_co_pernapasan" json:"hasil_pemeriksaan_co_pernapasan"`
	Nip                                             string `form:"nip" json:"nip"`
}

func skriningPerilakuMerokokSekolahRemajaRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                               "required|date",
		"kd_sekolah":                                            "string|max_len:5",
		"kelas":                                                 "in:1,2,3,4,5,6,7,8,9,10,11,12",
		"apakah_anda_merokok":                                   "required|string",
		"jumlah_batang_rokok":                                   "string|max_len:4",
		"jumlah_batang_rokok_hariminggu":                        "string",
		"jenis_rokok_yang_digunakan":                            "string",
		"jenis_rokok_yang_digunakan_keterangan":                 "string|max_len:40",
		"usia_mulai_merokok":                                    "string|max_len:3",
		"alasan_mulai_merokok":                                  "string",
		"alasan_mulai_merokok_keterangan":                       "string|max_len:40",
		"sudah_berapa_lama_merokok":                             "string|max_len:5",
		"bagaimana_biasanya_mendapatkan_rokok":                  "string",
		"bagaimana_biasanya_mendapatkan_rokok_keterangan":       "string|max_len:40",
		"keinginan_berhenti_merokok":                            "string",
		"alasan_utama_berhenti_merokok":                         "string",
		"alasan_utama_berhenti_merokok_keterangan":              "string|max_len:40",
		"tahu_dampak_kesehatan_merokok":                         "required|in:Ya,Tidak",
		"dampak_kesehatan_dari_merokok_yang_diketahui":          "string",
		"tahu_merokok_pintu_masuk_narkoba":                      "required|in:Ya,Tidak",
		"melihat_orang_merokok_di_sekolah":                      "required|in:Ya,Tidak",
		"orang_yang_paling_sering_merokok_disekolah":            "string",
		"orang_yang_paling_sering_merokok_disekolah_keterangan": "string|max_len:40",
		"ada_anggota_keluarga_di_rumah_yang_merokok":            "required|in:Ya,Tidak",
		"teman_dekat_banyakyang_merokok":                        "required|in:Ya,Tidak",
		"dilakukan_pemeriksaan_kadar_co_pernapasan":             "required|string",
		"hasil_pemeriksaan_co_pernapasan":                       "string|max_len:5",
		"nip":                                                   "required|string|max_len:20",
	}
	return rules
}

// SkriningPerilakuMerokokSekolahRemajaStore simpan skrining merokok usia sekolah remaja; kolom waktu kunci kosong = sekarang.
type SkriningPerilakuMerokokSekolahRemajaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningPerilakuMerokokSekolahRemajaData
}

func (r *SkriningPerilakuMerokokSekolahRemajaStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningPerilakuMerokokSekolahRemajaStore) Rules(ctx http.Context) map[string]any {
	rules := skriningPerilakuMerokokSekolahRemajaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningPerilakuMerokokSekolahRemajaStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *SkriningPerilakuMerokokSekolahRemajaStore) Payload() SkriningPerilakuMerokokSekolahRemajaData {
	return r.SkriningPerilakuMerokokSekolahRemajaData
}

func (r *SkriningPerilakuMerokokSekolahRemajaStore) DetailValues() map[string][]string { return nil }

// SkriningPerilakuMerokokSekolahRemajaUpdate ubah skrining merokok usia sekolah remaja (PUT); kunci lewat query string.
type SkriningPerilakuMerokokSekolahRemajaUpdate struct {
	SkriningPerilakuMerokokSekolahRemajaData
}

func (r *SkriningPerilakuMerokokSekolahRemajaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningPerilakuMerokokSekolahRemajaUpdate) Rules(ctx http.Context) map[string]any {
	return skriningPerilakuMerokokSekolahRemajaRules()
}

func (r *SkriningPerilakuMerokokSekolahRemajaUpdate) Payload() SkriningPerilakuMerokokSekolahRemajaData {
	return r.SkriningPerilakuMerokokSekolahRemajaData
}

func (r *SkriningPerilakuMerokokSekolahRemajaUpdate) DetailValues() map[string][]string { return nil }
