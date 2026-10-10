package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianTambahanGeriatriData isian penilaian tambahan geriatri.
type PenilaianTambahanGeriatriData struct {
	Tanggal                             string `form:"tanggal" json:"tanggal"`
	Nik                                 string `form:"nik" json:"nik"`
	AsalMasuk                           string `form:"asal_masuk" json:"asal_masuk"`
	KondisiMasuk                        string `form:"kondisi_masuk" json:"kondisi_masuk"`
	KeteranganKondisiMasuk              string `form:"keterangan_kondisi_masuk" json:"keterangan_kondisi_masuk"`
	Anamnesis                           string `form:"anamnesis" json:"anamnesis"`
	DiagnosaMedis                       string `form:"diagnosa_medis" json:"diagnosa_medis"`
	RiwayatImmunoTelinga                string `form:"riwayat_immuno_telinga" json:"riwayat_immuno_telinga"`
	RiwayatImmunoSinus                  string `form:"riwayat_immuno_sinus" json:"riwayat_immuno_sinus"`
	RiwayatImmunoAntibiotik             string `form:"riwayat_immuno_antibiotik" json:"riwayat_immuno_antibiotik"`
	RiwayatImmunoPneumonia              string `form:"riwayat_immuno_pneumonia" json:"riwayat_immuno_pneumonia"`
	RiwayatImmunoAbses                  string `form:"riwayat_immuno_abses" json:"riwayat_immuno_abses"`
	RiwayatImmunoSariawan               string `form:"riwayat_immuno_sariawan" json:"riwayat_immuno_sariawan"`
	RiwayatImmunoMemerlukanAntibiotik   string `form:"riwayat_immuno_memerlukan_antibiotik" json:"riwayat_immuno_memerlukan_antibiotik"`
	RiwayatImmunoInfeksiDalam           string `form:"riwayat_immuno_infeksi_dalam" json:"riwayat_immuno_infeksi_dalam"`
	RiwayatImmunoImmunodefisiensiPrimer string `form:"riwayat_immuno_immunodefisiensi_primer" json:"riwayat_immuno_immunodefisiensi_primer"`
	RiwayatImmunoJenisKangker           string `form:"riwayat_immuno_jenis_kangker" json:"riwayat_immuno_jenis_kangker"`
	RiwayatImmunoInfeksiOportunistik    string `form:"riwayat_immuno_infeksi_oportunistik" json:"riwayat_immuno_infeksi_oportunistik"`
	PolaAktifitasTidur                  string `form:"pola_aktifitas_tidur" json:"pola_aktifitas_tidur"`
	KeteranganPolaAktifitasTidur        string `form:"keterangan_pola_aktifitas_tidur" json:"keterangan_pola_aktifitas_tidur"`
	PolaAktifitasObatTidur              string `form:"pola_aktifitas_obat_tidur" json:"pola_aktifitas_obat_tidur"`
	KeteranganPolaAktifitasObatTidur    string `form:"keterangan_pola_aktifitas_obat_tidur" json:"keterangan_pola_aktifitas_obat_tidur"`
	PolaAktifitasOlahraga               string `form:"pola_aktifitas_olahraga" json:"pola_aktifitas_olahraga"`
	KeteranganPolaAktifitasOlahraga     string `form:"keterangan_pola_aktifitas_olahraga" json:"keterangan_pola_aktifitas_olahraga"`
	KualitasHidupMobilitas              string `form:"kualitas_hidup_mobilitas" json:"kualitas_hidup_mobilitas"`
	KualitasHidupPerawatanDiri          string `form:"kualitas_hidup_perawatan_diri" json:"kualitas_hidup_perawatan_diri"`
	KualitasHidupAktifitasSeharihari    string `form:"kualitas_hidup_aktifitas_seharihari" json:"kualitas_hidup_aktifitas_seharihari"`
	KualitasHidupRasaNyeri              string `form:"kualitas_hidup_rasa_nyeri" json:"kualitas_hidup_rasa_nyeri"`
	SkalaNyeri                          string `form:"skala_nyeri" json:"skala_nyeri"`
}

func penilaianTambahanGeriatriRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                "required|date",
		"nik":                                    "required|string|max_len:20",
		"asal_masuk":                             "in:IGD,Kamar Bersalin,Klinik",
		"kondisi_masuk":                          "in:Mandiri,Kursi Roda,Dipapah,Tempat Tidur",
		"keterangan_kondisi_masuk":               "string|max_len:50",
		"anamnesis":                              "required|in:Autoanamnesis,Alloanamnesis",
		"diagnosa_medis":                         "string|max_len:100",
		"riwayat_immuno_telinga":                 "in:Ya,Tidak",
		"riwayat_immuno_sinus":                   "in:Ya,Tidak",
		"riwayat_immuno_antibiotik":              "in:Ya,Tidak",
		"riwayat_immuno_pneumonia":               "in:Ya,Tidak",
		"riwayat_immuno_abses":                   "in:Ya,Tidak",
		"riwayat_immuno_sariawan":                "in:Ya,Tidak",
		"riwayat_immuno_memerlukan_antibiotik":   "in:Ya,Tidak",
		"riwayat_immuno_infeksi_dalam":           "in:Ya,Tidak",
		"riwayat_immuno_immunodefisiensi_primer": "in:Ya,Tidak",
		"riwayat_immuno_jenis_kangker":           "in:Ya,Tidak",
		"riwayat_immuno_infeksi_oportunistik":    "in:Ya,Tidak",
		"pola_aktifitas_tidur":                   "in:TAK,Insomnia,Lainnya",
		"keterangan_pola_aktifitas_tidur":        "string|max_len:50",
		"pola_aktifitas_obat_tidur":              "in:Tidak,Ya",
		"keterangan_pola_aktifitas_obat_tidur":   "string|max_len:50",
		"pola_aktifitas_olahraga":                "in:Tidak,Ya",
		"keterangan_pola_aktifitas_olahraga":     "string|max_len:50",
		"kualitas_hidup_mobilitas":               "in:Tidak Mempunyai Masalah Untuk Berjalan,Ada Masalah Untuk Berjalan,Hanya Mampu Berbaring",
		"kualitas_hidup_perawatan_diri":          "string",
		"kualitas_hidup_aktifitas_seharihari":    "in:Tak Mempunyai Kesulitan Dalam Melaksanakan Kegiatan Sehari-hari,Mempunyai Keterbatasan Dalam Melaksanakan Kegiatan Sehari-hari,Tak Mampu Melaksanakan Kegiatan Sehari-hari",
		"kualitas_hidup_rasa_nyeri":              "in:Tidak Mempunyai Keluhan Rasa Nyeri Atau Rasa Tak Nyaman,Suka Merasakan Agak Nyeri/Agak Kurang Nyaman,Menderita Karena Keluhan Rasa Nyeri/Tidak Nyaman",
		"skala_nyeri":                            "string",
	}
	return rules
}

// PenilaianTambahanGeriatriStore simpan penilaian tambahan geriatri; kolom waktu kunci kosong = sekarang.
type PenilaianTambahanGeriatriStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianTambahanGeriatriData
}

func (r *PenilaianTambahanGeriatriStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianTambahanGeriatriStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianTambahanGeriatriRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianTambahanGeriatriStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianTambahanGeriatriStore) Payload() PenilaianTambahanGeriatriData {
	return r.PenilaianTambahanGeriatriData
}

func (r *PenilaianTambahanGeriatriStore) DetailValues() map[string][]string { return nil }

// PenilaianTambahanGeriatriUpdate ubah penilaian tambahan geriatri (PUT); kunci lewat query string.
type PenilaianTambahanGeriatriUpdate struct {
	PenilaianTambahanGeriatriData
}

func (r *PenilaianTambahanGeriatriUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianTambahanGeriatriUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianTambahanGeriatriRules()
}

func (r *PenilaianTambahanGeriatriUpdate) Payload() PenilaianTambahanGeriatriData {
	return r.PenilaianTambahanGeriatriData
}

func (r *PenilaianTambahanGeriatriUpdate) DetailValues() map[string][]string { return nil }
