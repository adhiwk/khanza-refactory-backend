package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// EdukasiPasienKeluargaRjData isian edukasi pasien keluarga rawat jalan.
type EdukasiPasienKeluargaRjData struct {
	Tanggal                                 string `form:"tanggal" json:"tanggal"`
	Nip                                     string `form:"nip" json:"nip"`
	Bicara                                  string `form:"bicara" json:"bicara"`
	KeteranganBicara                        string `form:"keterangan_bicara" json:"keterangan_bicara"`
	BahasaSehari                            string `form:"bahasa_sehari" json:"bahasa_sehari"`
	PerluPenerjemah                         string `form:"perlu_penerjemah" json:"perlu_penerjemah"`
	KeteranganPenerjemah                    string `form:"keterangan_penerjemah" json:"keterangan_penerjemah"`
	BahasaIsyarat                           string `form:"bahasa_isyarat" json:"bahasa_isyarat"`
	CaraBelajar                             string `form:"cara_belajar" json:"cara_belajar"`
	HambatanBelajar                         string `form:"hambatan_belajar" json:"hambatan_belajar"`
	KeteranganHambatanBelajar               string `form:"keterangan_hambatan_belajar" json:"keterangan_hambatan_belajar"`
	KemampuanBelajar                        string `form:"kemampuan_belajar" json:"kemampuan_belajar"`
	KeteranganKemampuanBelajar              string `form:"keterangan_kemampuan_belajar" json:"keterangan_kemampuan_belajar"`
	PenyakitnyaMerupakan                    string `form:"penyakitnya_merupakan" json:"penyakitnya_merupakan"`
	KeteranganPenyakitnyaMerupakan          string `form:"keterangan_penyakitnya_merupakan" json:"keterangan_penyakitnya_merupakan"`
	KeputusanMemilihLayanan                 string `form:"keputusan_memilih_layanan" json:"keputusan_memilih_layanan"`
	KeteranganKeputusanMemilihLayanan       string `form:"keterangan_keputusan_memilih_layanan" json:"keterangan_keputusan_memilih_layanan"`
	KeyakinanTerhadapTerapi                 string `form:"keyakinan_terhadap_terapi" json:"keyakinan_terhadap_terapi"`
	KeteranganKeyakinanTerhadapTerapi       string `form:"keterangan_keyakinan_terhadap_terapi" json:"keterangan_keyakinan_terhadap_terapi"`
	AspekKeyakinanDipertimbangkan           string `form:"aspek_keyakinan_dipertimbangkan" json:"aspek_keyakinan_dipertimbangkan"`
	KeteranganAspekKeyakinanDipertimbangkan string `form:"keterangan_aspek_keyakinan_dipertimbangkan" json:"keterangan_aspek_keyakinan_dipertimbangkan"`
	KesediaanMenerimaInformasi              string `form:"kesediaan_menerima_informasi" json:"kesediaan_menerima_informasi"`
	TopikEdukasiPenyakit                    string `form:"topik_edukasi_penyakit" json:"topik_edukasi_penyakit"`
	TopikEdukasiRencanaTindakan             string `form:"topik_edukasi_rencana_tindakan" json:"topik_edukasi_rencana_tindakan"`
	TopikEdukasiPengobatan                  string `form:"topik_edukasi_pengobatan" json:"topik_edukasi_pengobatan"`
	TopikEdukasiHasilLayanan                string `form:"topik_edukasi_hasil_layanan" json:"topik_edukasi_hasil_layanan"`
}

func edukasiPasienKeluargaRjRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                    "required|date",
		"nip":                                        "required|string|max_len:20",
		"bicara":                                     "in:Normal,Gangguan Bicara",
		"keterangan_bicara":                          "string|max_len:50",
		"bahasa_sehari":                              "string|max_len:50",
		"perlu_penerjemah":                           "in:Ya,Tidak",
		"keterangan_penerjemah":                      "string|max_len:50",
		"bahasa_isyarat":                             "in:Ya,Tidak",
		"cara_belajar":                               "in:Menulis,Audio-Visual/Gambar,Diskusi,Simulasi",
		"hambatan_belajar":                           "in:Tidak Ada,Takut/Gelisah,Tidak Tertarik,Nyeri Tidak Nyaman,Buta Huruf,Gangguan Kognitif,Lain-lain",
		"keterangan_hambatan_belajar":                "string|max_len:50",
		"kemampuan_belajar":                          "in:Mampu Menerima Informasi,Tidak Mampu Menerima Informasi",
		"keterangan_kemampuan_belajar":               "string|max_len:50",
		"penyakitnya_merupakan":                      "required|in:Ujian/Cobaan,Kutukan,Lain-lain",
		"keterangan_penyakitnya_merupakan":           "string|max_len:50",
		"keputusan_memilih_layanan":                  "required|in:Sendiri,Keluarga,Lain-lain",
		"keterangan_keputusan_memilih_layanan":       "string|max_len:50",
		"keyakinan_terhadap_terapi":                  "required|in:Pasrah,Yakin Sembuh Jika Kontrol Teratur,Yakin Sembuh Jika Minum Obat Teratur,Lain-lain",
		"keterangan_keyakinan_terhadap_terapi":       "string|max_len:50",
		"aspek_keyakinan_dipertimbangkan":            "required|in:Ada,Tidak",
		"keterangan_aspek_keyakinan_dipertimbangkan": "string|max_len:50",
		"kesediaan_menerima_informasi":               "required|in:Ya,Tidak",
		"topik_edukasi_penyakit":                     "required|in:Ya,Tidak",
		"topik_edukasi_rencana_tindakan":             "required|in:Ya,Tidak",
		"topik_edukasi_pengobatan":                   "required|in:Ya,Tidak",
		"topik_edukasi_hasil_layanan":                "required|in:Ya,Tidak",
	}
	return rules
}

// EdukasiPasienKeluargaRjStore simpan edukasi pasien keluarga rawat jalan; kolom waktu kunci kosong = sekarang.
type EdukasiPasienKeluargaRjStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	EdukasiPasienKeluargaRjData
}

func (r *EdukasiPasienKeluargaRjStore) Authorize(ctx http.Context) error { return nil }

func (r *EdukasiPasienKeluargaRjStore) Rules(ctx http.Context) map[string]any {
	rules := edukasiPasienKeluargaRjRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *EdukasiPasienKeluargaRjStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *EdukasiPasienKeluargaRjStore) Payload() EdukasiPasienKeluargaRjData {
	return r.EdukasiPasienKeluargaRjData
}

func (r *EdukasiPasienKeluargaRjStore) DetailValues() map[string][]string { return nil }

// EdukasiPasienKeluargaRjUpdate ubah edukasi pasien keluarga rawat jalan (PUT); kunci lewat query string.
type EdukasiPasienKeluargaRjUpdate struct {
	EdukasiPasienKeluargaRjData
}

func (r *EdukasiPasienKeluargaRjUpdate) Authorize(ctx http.Context) error { return nil }

func (r *EdukasiPasienKeluargaRjUpdate) Rules(ctx http.Context) map[string]any {
	return edukasiPasienKeluargaRjRules()
}

func (r *EdukasiPasienKeluargaRjUpdate) Payload() EdukasiPasienKeluargaRjData {
	return r.EdukasiPasienKeluargaRjData
}

func (r *EdukasiPasienKeluargaRjUpdate) DetailValues() map[string][]string { return nil }
