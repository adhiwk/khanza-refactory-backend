package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianPasienImunitasRendahData isian penilaian pasien imunitas rendah.
type PenilaianPasienImunitasRendahData struct {
	Tanggal                            string `form:"tanggal" json:"tanggal"`
	KdDokter                           string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis                          string `form:"anamnesis" json:"anamnesis"`
	Hubungan                           string `form:"hubungan" json:"hubungan"`
	PasienMengetahuiKondisiPenyakitnya string `form:"pasien_mengetahui_kondisi_penyakitnya" json:"pasien_mengetahui_kondisi_penyakitnya"`
	KebutuhanRuangPerawatan            string `form:"kebutuhan_ruang_perawatan" json:"kebutuhan_ruang_perawatan"`
	RiwayatPenyakitKeluhan             string `form:"riwayat_penyakit_keluhan" json:"riwayat_penyakit_keluhan"`
	RiwayatPenyakitKeluarga            string `form:"riwayat_penyakit_keluarga" json:"riwayat_penyakit_keluarga"`
	RiwayatAlergi                      string `form:"riwayat_alergi" json:"riwayat_alergi"`
	RiwayatVaksinasi                   string `form:"riwayat_vaksinasi" json:"riwayat_vaksinasi"`
	RiwayatPengobatan                  string `form:"riwayat_pengobatan" json:"riwayat_pengobatan"`
	DiagnosaUtama                      string `form:"diagnosa_utama" json:"diagnosa_utama"`
	DiagnosaTambahan                   string `form:"diagnosa_tambahan" json:"diagnosa_tambahan"`
}

func penilaianPasienImunitasRendahRules() map[string]any {
	rules := map[string]any{
		"tanggal":                               "required|date",
		"kd_dokter":                             "required|string|max_len:20",
		"anamnesis":                             "in:Autoanamnesis,Alloanamnesis",
		"hubungan":                              "string|max_len:30",
		"pasien_mengetahui_kondisi_penyakitnya": "in:Ya,Tidak",
		"kebutuhan_ruang_perawatan":             "in:Isolasi,Ruang Rawat Biasa,ICU,ICU Isolasi",
		"riwayat_penyakit_keluhan":              "string|max_len:1000",
		"riwayat_penyakit_keluarga":             "string|max_len:1000",
		"riwayat_alergi":                        "string|max_len:100",
		"riwayat_vaksinasi":                     "string|max_len:100",
		"riwayat_pengobatan":                    "string|max_len:1000",
		"diagnosa_utama":                        "string|max_len:500",
		"diagnosa_tambahan":                     "string|max_len:500",
	}
	return rules
}

// PenilaianPasienImunitasRendahStore simpan penilaian pasien imunitas rendah; kolom waktu kunci kosong = sekarang.
type PenilaianPasienImunitasRendahStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianPasienImunitasRendahData
}

func (r *PenilaianPasienImunitasRendahStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPasienImunitasRendahStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianPasienImunitasRendahRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianPasienImunitasRendahStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianPasienImunitasRendahStore) Payload() PenilaianPasienImunitasRendahData {
	return r.PenilaianPasienImunitasRendahData
}

func (r *PenilaianPasienImunitasRendahStore) DetailValues() map[string][]string { return nil }

// PenilaianPasienImunitasRendahUpdate ubah penilaian pasien imunitas rendah (PUT); kunci lewat query string.
type PenilaianPasienImunitasRendahUpdate struct {
	PenilaianPasienImunitasRendahData
}

func (r *PenilaianPasienImunitasRendahUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPasienImunitasRendahUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianPasienImunitasRendahRules()
}

func (r *PenilaianPasienImunitasRendahUpdate) Payload() PenilaianPasienImunitasRendahData {
	return r.PenilaianPasienImunitasRendahData
}

func (r *PenilaianPasienImunitasRendahUpdate) DetailValues() map[string][]string { return nil }
