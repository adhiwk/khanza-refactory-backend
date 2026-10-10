package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianPsikologiKlinisData isian penilaian psikologi klinis.
type PenilaianPsikologiKlinisData struct {
	Tanggal                            string `form:"tanggal" json:"tanggal"`
	Nip                                string `form:"nip" json:"nip"`
	Anamnesis                          string `form:"anamnesis" json:"anamnesis"`
	DikirimDari                        string `form:"dikirim_dari" json:"dikirim_dari"`
	TujuanPemeriksaan                  string `form:"tujuan_pemeriksaan" json:"tujuan_pemeriksaan"`
	KetAnamnesis                       string `form:"ket_anamnesis" json:"ket_anamnesis"`
	KeluhanUtama                       string `form:"keluhan_utama" json:"keluhan_utama"`
	RiwayatPenyakit                    string `form:"riwayat_penyakit" json:"riwayat_penyakit"`
	RiwayatKeluhan                     string `form:"riwayat_keluhan" json:"riwayat_keluhan"`
	PermasalahanSaatIni                string `form:"permasalahan_saat_ini" json:"permasalahan_saat_ini"`
	PermasalahanAlasan                 string `form:"permasalahan_alasan" json:"permasalahan_alasan"`
	PermasalahanEkspektasi             string `form:"permasalahan_ekspektasi" json:"permasalahan_ekspektasi"`
	RiwayatHidupSingkat                string `form:"riwayat_hidup_singkat" json:"riwayat_hidup_singkat"`
	KondisiPsikologisPenampilan        string `form:"kondisi_psikologis_penampilan" json:"kondisi_psikologis_penampilan"`
	KondisiPsikologisEkspresiWajah     string `form:"kondisi_psikologis_ekspresi_wajah" json:"kondisi_psikologis_ekspresi_wajah"`
	KondisiPsikologisSuasanaHati       string `form:"kondisi_psikologis_suasana_hati" json:"kondisi_psikologis_suasana_hati"`
	KondisiPsikologisTingkahLaku       string `form:"kondisi_psikologis_tingkah_laku" json:"kondisi_psikologis_tingkah_laku"`
	KondisiPsikologisFungsiUmum        string `form:"kondisi_psikologis_fungsi_umum" json:"kondisi_psikologis_fungsi_umum"`
	KondisiPsikologisFungsiIntelektual string `form:"kondisi_psikologis_fungsi_intelektual" json:"kondisi_psikologis_fungsi_intelektual"`
	KondisiPsikologisPengalaman        string `form:"kondisi_psikologis_pengalaman" json:"kondisi_psikologis_pengalaman"`
	KondisiPsikologisLainnya           string `form:"kondisi_psikologis_lainnya" json:"kondisi_psikologis_lainnya"`
	KondisiPatologisDelusi             string `form:"kondisi_patologis_delusi" json:"kondisi_patologis_delusi"`
	KondisiPatologisProsesPikiran      string `form:"kondisi_patologis_proses_pikiran" json:"kondisi_patologis_proses_pikiran"`
	KondisiPatologisHalusinasi         string `form:"kondisi_patologis_halusinasi" json:"kondisi_patologis_halusinasi"`
	KondisiPatologisAfek               string `form:"kondisi_patologis_afek" json:"kondisi_patologis_afek"`
	KondisiPatologisInsight            string `form:"kondisi_patologis_insight" json:"kondisi_patologis_insight"`
	KondisiPatologisKesadaran          string `form:"kondisi_patologis_kesadaran" json:"kondisi_patologis_kesadaran"`
	KondisiPatologisOrientasi          string `form:"kondisi_patologis_orientasi" json:"kondisi_patologis_orientasi"`
	KondisiPatologisAtensi             string `form:"kondisi_patologis_atensi" json:"kondisi_patologis_atensi"`
	KondisiPatologisKontrolImpuls      string `form:"kondisi_patologis_kontrol_impuls" json:"kondisi_patologis_kontrol_impuls"`
	PsikotesTanggalPelaksanaan         string `form:"psikotes_tanggal_pelaksanaan" json:"psikotes_tanggal_pelaksanaan"`
	PsikotesNamaTes                    string `form:"psikotes_nama_tes" json:"psikotes_nama_tes"`
	PsikotesHasil                      string `form:"psikotes_hasil" json:"psikotes_hasil"`
	DinamikaPsikologis                 string `form:"dinamika_psikologis" json:"dinamika_psikologis"`
	DiagnosaPsikologis                 string `form:"diagnosa_psikologis" json:"diagnosa_psikologis"`
	ManifestasiFungsiPsikologis        string `form:"manifestasi_fungsi_psikologis" json:"manifestasi_fungsi_psikologis"`
	RencanaIntervensi                  string `form:"rencana_intervensi" json:"rencana_intervensi"`
	TahapanIntervensi1                 string `form:"tahapan_intervensi1" json:"tahapan_intervensi1"`
	TargetTerapi1                      string `form:"target_terapi1" json:"target_terapi1"`
	TahapanIntervensi2                 string `form:"tahapan_intervensi2" json:"tahapan_intervensi2"`
	TargetTerapi2                      string `form:"target_terapi2" json:"target_terapi2"`
	TahapanIntervensi3                 string `form:"tahapan_intervensi3" json:"tahapan_intervensi3"`
	TargetTerapi3                      string `form:"target_terapi3" json:"target_terapi3"`
	TahapanIntervensi4                 string `form:"tahapan_intervensi4" json:"tahapan_intervensi4"`
	TargetTerapi4                      string `form:"target_terapi4" json:"target_terapi4"`
	TahapanIntervensi5                 string `form:"tahapan_intervensi5" json:"tahapan_intervensi5"`
	TargetTerapi5                      string `form:"target_terapi5" json:"target_terapi5"`
	TahapanIntervensi6                 string `form:"tahapan_intervensi6" json:"tahapan_intervensi6"`
	TargetTerapi6                      string `form:"target_terapi6" json:"target_terapi6"`
	TahapanIntervensi7                 string `form:"tahapan_intervensi7" json:"tahapan_intervensi7"`
	TargetTerapi7                      string `form:"target_terapi7" json:"target_terapi7"`
	Evaluasi                           string `form:"evaluasi" json:"evaluasi"`
}

func penilaianPsikologiKlinisRules() map[string]any {
	rules := map[string]any{
		"tanggal":                               "required|date",
		"nip":                                   "required|string|max_len:20",
		"anamnesis":                             "required|in:Autoanamnesis,Alloanamnesis",
		"dikirim_dari":                          "required|in:Ruang Rawat,Poliklinik,Rehabilitasi,After Care,Dokter",
		"tujuan_pemeriksaan":                    "required|in:Klinik,Bimbingan,Forensik",
		"ket_anamnesis":                         "string|max_len:200",
		"keluhan_utama":                         "string|max_len:2000",
		"riwayat_penyakit":                      "string|max_len:1000",
		"riwayat_keluhan":                       "string|max_len:1000",
		"permasalahan_saat_ini":                 "in:Sangat Serius,Serius,Cukup Serius",
		"permasalahan_alasan":                   "string|max_len:100",
		"permasalahan_ekspektasi":               "string|max_len:100",
		"riwayat_hidup_singkat":                 "string|max_len:1000",
		"kondisi_psikologis_penampilan":         "string|max_len:150",
		"kondisi_psikologis_ekspresi_wajah":     "string|max_len:150",
		"kondisi_psikologis_suasana_hati":       "string|max_len:150",
		"kondisi_psikologis_tingkah_laku":       "string|max_len:150",
		"kondisi_psikologis_fungsi_umum":        "string|max_len:150",
		"kondisi_psikologis_fungsi_intelektual": "string|max_len:150",
		"kondisi_psikologis_pengalaman":         "string|max_len:150",
		"kondisi_psikologis_lainnya":            "string|max_len:150",
		"kondisi_patologis_delusi":              "string|max_len:150",
		"kondisi_patologis_proses_pikiran":      "string|max_len:150",
		"kondisi_patologis_halusinasi":          "string|max_len:150",
		"kondisi_patologis_afek":                "string|max_len:150",
		"kondisi_patologis_insight":             "string|max_len:150",
		"kondisi_patologis_kesadaran":           "string|max_len:150",
		"kondisi_patologis_orientasi":           "string|max_len:150",
		"kondisi_patologis_atensi":              "string|max_len:150",
		"kondisi_patologis_kontrol_impuls":      "string|max_len:150",
		"psikotes_tanggal_pelaksanaan":          "date",
		"psikotes_nama_tes":                     "string|max_len:100",
		"psikotes_hasil":                        "string|max_len:200",
		"dinamika_psikologis":                   "string|max_len:1000",
		"diagnosa_psikologis":                   "string|max_len:1000",
		"manifestasi_fungsi_psikologis":         "string|max_len:1000",
		"rencana_intervensi":                    "string|max_len:1000",
		"tahapan_intervensi1":                   "string|max_len:100",
		"target_terapi1":                        "string|max_len:100",
		"tahapan_intervensi2":                   "string|max_len:100",
		"target_terapi2":                        "string|max_len:100",
		"tahapan_intervensi3":                   "string|max_len:100",
		"target_terapi3":                        "string|max_len:100",
		"tahapan_intervensi4":                   "string|max_len:100",
		"target_terapi4":                        "string|max_len:100",
		"tahapan_intervensi5":                   "string|max_len:100",
		"target_terapi5":                        "string|max_len:100",
		"tahapan_intervensi6":                   "string|max_len:100",
		"target_terapi6":                        "string|max_len:100",
		"tahapan_intervensi7":                   "string|max_len:100",
		"target_terapi7":                        "string|max_len:100",
		"evaluasi":                              "string|max_len:1000",
	}
	return rules
}

// PenilaianPsikologiKlinisStore simpan penilaian psikologi klinis; kolom waktu kunci kosong = sekarang.
type PenilaianPsikologiKlinisStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianPsikologiKlinisData
}

func (r *PenilaianPsikologiKlinisStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPsikologiKlinisStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianPsikologiKlinisRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianPsikologiKlinisStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianPsikologiKlinisStore) Payload() PenilaianPsikologiKlinisData {
	return r.PenilaianPsikologiKlinisData
}

func (r *PenilaianPsikologiKlinisStore) DetailValues() map[string][]string { return nil }

// PenilaianPsikologiKlinisUpdate ubah penilaian psikologi klinis (PUT); kunci lewat query string.
type PenilaianPsikologiKlinisUpdate struct {
	PenilaianPsikologiKlinisData
}

func (r *PenilaianPsikologiKlinisUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPsikologiKlinisUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianPsikologiKlinisRules()
}

func (r *PenilaianPsikologiKlinisUpdate) Payload() PenilaianPsikologiKlinisData {
	return r.PenilaianPsikologiKlinisData
}

func (r *PenilaianPsikologiKlinisUpdate) DetailValues() map[string][]string { return nil }
