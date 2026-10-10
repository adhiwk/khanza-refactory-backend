package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianPasienPenyakitMenularData isian penilaian pasien penyakit menular.
type PenilaianPasienPenyakitMenularData struct {
	Tanggal                              string `form:"tanggal" json:"tanggal"`
	KdDokter                             string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis                            string `form:"anamnesis" json:"anamnesis"`
	Hubungan                             string `form:"hubungan" json:"hubungan"`
	PasienMengetahuiKondisiPenyakitnya   string `form:"pasien_mengetahui_kondisi_penyakitnya" json:"pasien_mengetahui_kondisi_penyakitnya"`
	PenyakitSamaSerumah                  string `form:"penyakit_sama_serumah" json:"penyakit_sama_serumah"`
	RiwayatKontak                        string `form:"riwayat_kontak" json:"riwayat_kontak"`
	KeteranganRiwayatKontak              string `form:"keterangan_riwayat_kontak" json:"keterangan_riwayat_kontak"`
	TransmisiPenularanPenyakit           string `form:"transmisi_penularan_penyakit" json:"transmisi_penularan_penyakit"`
	KeteranganTransmisiPenularanPenyakit string `form:"keterangan_transmisi_penularan_penyakit" json:"keterangan_transmisi_penularan_penyakit"`
	KebutuhanRuangRawat                  string `form:"kebutuhan_ruang_rawat" json:"kebutuhan_ruang_rawat"`
	KeluhanYangDirasakanSaatIni          string `form:"keluhan_yang_dirasakan_saat_ini" json:"keluhan_yang_dirasakan_saat_ini"`
	RiwayatPenyakitKeluarga              string `form:"riwayat_penyakit_keluarga" json:"riwayat_penyakit_keluarga"`
	RiwayatAlergi                        string `form:"riwayat_alergi" json:"riwayat_alergi"`
	RiwayatVaksinasi                     string `form:"riwayat_vaksinasi" json:"riwayat_vaksinasi"`
	RiwayatPengobatan                    string `form:"riwayat_pengobatan" json:"riwayat_pengobatan"`
	DiagnosaUtama                        string `form:"diagnosa_utama" json:"diagnosa_utama"`
	DiagnosaTambahan                     string `form:"diagnosa_tambahan" json:"diagnosa_tambahan"`
}

func penilaianPasienPenyakitMenularRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                 "required|date",
		"kd_dokter":                               "required|string|max_len:20",
		"anamnesis":                               "required|in:Autoanamnesis,Alloanamnesis",
		"hubungan":                                "string|max_len:30",
		"pasien_mengetahui_kondisi_penyakitnya":   "in:Ya,Tidak",
		"penyakit_sama_serumah":                   "in:Ya,Tidak",
		"riwayat_kontak":                          "in:Bekerja Di Daerah KLB,Merawat Pasien Penyakit Infeksi,Berkunjung Ke Daerah Endemik Dalam 2 Minggu Terakhir,Bekerja Di Laboratorium,Kontak Langsung,Lain-lain",
		"keterangan_riwayat_kontak":               "string|max_len:70",
		"transmisi_penularan_penyakit":            "in:Airborne,Droplet,Kontak Langsung,Cairan Tubuh Lainnya,Lain-lain",
		"keterangan_transmisi_penularan_penyakit": "string|max_len:70",
		"kebutuhan_ruang_rawat":                   "in:Isolasi,Ruang Biasa,ICU,ICU Isolasi",
		"keluhan_yang_dirasakan_saat_ini":         "string|max_len:1000",
		"riwayat_penyakit_keluarga":               "string|max_len:1000",
		"riwayat_alergi":                          "string|max_len:100",
		"riwayat_vaksinasi":                       "string|max_len:100",
		"riwayat_pengobatan":                      "string|max_len:1000",
		"diagnosa_utama":                          "string|max_len:500",
		"diagnosa_tambahan":                       "string|max_len:500",
	}
	return rules
}

// PenilaianPasienPenyakitMenularStore simpan penilaian pasien penyakit menular; kolom waktu kunci kosong = sekarang.
type PenilaianPasienPenyakitMenularStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianPasienPenyakitMenularData
}

func (r *PenilaianPasienPenyakitMenularStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPasienPenyakitMenularStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianPasienPenyakitMenularRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianPasienPenyakitMenularStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianPasienPenyakitMenularStore) Payload() PenilaianPasienPenyakitMenularData {
	return r.PenilaianPasienPenyakitMenularData
}

func (r *PenilaianPasienPenyakitMenularStore) DetailValues() map[string][]string { return nil }

// PenilaianPasienPenyakitMenularUpdate ubah penilaian pasien penyakit menular (PUT); kunci lewat query string.
type PenilaianPasienPenyakitMenularUpdate struct {
	PenilaianPasienPenyakitMenularData
}

func (r *PenilaianPasienPenyakitMenularUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPasienPenyakitMenularUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianPasienPenyakitMenularRules()
}

func (r *PenilaianPasienPenyakitMenularUpdate) Payload() PenilaianPasienPenyakitMenularData {
	return r.PenilaianPasienPenyakitMenularData
}

func (r *PenilaianPasienPenyakitMenularUpdate) DetailValues() map[string][]string { return nil }
