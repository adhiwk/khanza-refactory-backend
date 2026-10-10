package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PerencanaanPemulanganData isian perencanaan pemulangan.
type PerencanaanPemulanganData struct {
	RencanaPulang                             string `form:"rencana_pulang" json:"rencana_pulang"`
	AlasanMasuk                               string `form:"alasan_masuk" json:"alasan_masuk"`
	DiagnosaMedis                             string `form:"diagnosa_medis" json:"diagnosa_medis"`
	PengaruhRiPasienDanKeluarga               string `form:"pengaruh_ri_pasien_dan_keluarga" json:"pengaruh_ri_pasien_dan_keluarga"`
	KeteranganPengaruhRiPasienDanKeluarga     string `form:"keterangan_pengaruh_ri_pasien_dan_keluarga" json:"keterangan_pengaruh_ri_pasien_dan_keluarga"`
	PengaruhRiPekerjaanSekolah                string `form:"pengaruh_ri_pekerjaan_sekolah" json:"pengaruh_ri_pekerjaan_sekolah"`
	KeteranganPengaruhRiPekerjaanSekolah      string `form:"keterangan_pengaruh_ri_pekerjaan_sekolah" json:"keterangan_pengaruh_ri_pekerjaan_sekolah"`
	PengaruhRiKeuangan                        string `form:"pengaruh_ri_keuangan" json:"pengaruh_ri_keuangan"`
	KeteranganPengaruhRiKeuangan              string `form:"keterangan_pengaruh_ri_keuangan" json:"keterangan_pengaruh_ri_keuangan"`
	AntisipasiMasalahSaatPulang               string `form:"antisipasi_masalah_saat_pulang" json:"antisipasi_masalah_saat_pulang"`
	KeteranganAntisipasiMasalahSaatPulang     string `form:"keterangan_antisipasi_masalah_saat_pulang" json:"keterangan_antisipasi_masalah_saat_pulang"`
	BantuanDiperlukanDalam                    string `form:"bantuan_diperlukan_dalam" json:"bantuan_diperlukan_dalam"`
	KeteranganBantuanDiperlukanDalam          string `form:"keterangan_bantuan_diperlukan_dalam" json:"keterangan_bantuan_diperlukan_dalam"`
	AdakahYangMembantuKeperluan               string `form:"adakah_yang_membantu_keperluan" json:"adakah_yang_membantu_keperluan"`
	KeteranganAdakahYangMembantuKeperluan     string `form:"keterangan_adakah_yang_membantu_keperluan" json:"keterangan_adakah_yang_membantu_keperluan"`
	PasienTinggalSendiri                      string `form:"pasien_tinggal_sendiri" json:"pasien_tinggal_sendiri"`
	KeteranganPasienTinggalSendiri            string `form:"keterangan_pasien_tinggal_sendiri" json:"keterangan_pasien_tinggal_sendiri"`
	PasienMenggunakanPeralatanMedis           string `form:"pasien_menggunakan_peralatan_medis" json:"pasien_menggunakan_peralatan_medis"`
	KeteranganPasienMenggunakanPeralatanMedis string `form:"keterangan_pasien_menggunakan_peralatan_medis" json:"keterangan_pasien_menggunakan_peralatan_medis"`
	PasienMemerlukanAlatBantu                 string `form:"pasien_memerlukan_alat_bantu" json:"pasien_memerlukan_alat_bantu"`
	KeteranganPasienMemerlukanAlatBantu       string `form:"keterangan_pasien_memerlukan_alat_bantu" json:"keterangan_pasien_memerlukan_alat_bantu"`
	MemerlukanPerawatanKhusus                 string `form:"memerlukan_perawatan_khusus" json:"memerlukan_perawatan_khusus"`
	KeteranganMemerlukanPerawatanKhusus       string `form:"keterangan_memerlukan_perawatan_khusus" json:"keterangan_memerlukan_perawatan_khusus"`
	BermasalahMemenuhiKebutuhan               string `form:"bermasalah_memenuhi_kebutuhan" json:"bermasalah_memenuhi_kebutuhan"`
	KeteranganBermasalahMemenuhiKebutuhan     string `form:"keterangan_bermasalah_memenuhi_kebutuhan" json:"keterangan_bermasalah_memenuhi_kebutuhan"`
	MemilikiNyeriKronis                       string `form:"memiliki_nyeri_kronis" json:"memiliki_nyeri_kronis"`
	KeteranganMemilikiNyeriKronis             string `form:"keterangan_memiliki_nyeri_kronis" json:"keterangan_memiliki_nyeri_kronis"`
	MemerlukanEdukasiKesehatan                string `form:"memerlukan_edukasi_kesehatan" json:"memerlukan_edukasi_kesehatan"`
	KeteranganMemerlukanEdukasiKesehatan      string `form:"keterangan_memerlukan_edukasi_kesehatan" json:"keterangan_memerlukan_edukasi_kesehatan"`
	MemerlukanKeterampilkanKhusus             string `form:"memerlukan_keterampilkan_khusus" json:"memerlukan_keterampilkan_khusus"`
	KeteranganMemerlukanKeterampilkanKhusus   string `form:"keterangan_memerlukan_keterampilkan_khusus" json:"keterangan_memerlukan_keterampilkan_khusus"`
	NamaPasienKeluarga                        string `form:"nama_pasien_keluarga" json:"nama_pasien_keluarga"`
	Nip                                       string `form:"nip" json:"nip"`
}

func perencanaanPemulanganRules() map[string]any {
	rules := map[string]any{
		"rencana_pulang":                                "required|date",
		"alasan_masuk":                                  "string|max_len:150",
		"diagnosa_medis":                                "string|max_len:50",
		"pengaruh_ri_pasien_dan_keluarga":               "in:Tidak,Ya",
		"keterangan_pengaruh_ri_pasien_dan_keluarga":    "string|max_len:100",
		"pengaruh_ri_pekerjaan_sekolah":                 "required|in:Tidak,Ya",
		"keterangan_pengaruh_ri_pekerjaan_sekolah":      "string|max_len:100",
		"pengaruh_ri_keuangan":                          "required|in:Tidak,Ya",
		"keterangan_pengaruh_ri_keuangan":               "string|max_len:100",
		"antisipasi_masalah_saat_pulang":                "required|in:Tidak,Ya",
		"keterangan_antisipasi_masalah_saat_pulang":     "string|max_len:100",
		"bantuan_diperlukan_dalam":                      "required|in:Menyiapkan Makanan,Edukasi Kesehatan,Makan,Mandi,Diet,Berpakaian,Menyiapkan Obat,Transportasi,Minum Obat",
		"keterangan_bantuan_diperlukan_dalam":           "string|max_len:100",
		"adakah_yang_membantu_keperluan":                "required|in:Tidak,Ada",
		"keterangan_adakah_yang_membantu_keperluan":     "string|max_len:100",
		"pasien_tinggal_sendiri":                        "required|in:Tidak,Ya",
		"keterangan_pasien_tinggal_sendiri":             "string|max_len:100",
		"pasien_menggunakan_peralatan_medis":            "required|in:Tidak,Ya",
		"keterangan_pasien_menggunakan_peralatan_medis": "string|max_len:100",
		"pasien_memerlukan_alat_bantu":                  "required|in:Tidak,Ya",
		"keterangan_pasien_memerlukan_alat_bantu":       "string|max_len:100",
		"memerlukan_perawatan_khusus":                   "required|in:Tidak,Ya",
		"keterangan_memerlukan_perawatan_khusus":        "string|max_len:100",
		"bermasalah_memenuhi_kebutuhan":                 "required|in:Tidak,Ya",
		"keterangan_bermasalah_memenuhi_kebutuhan":      "string|max_len:100",
		"memiliki_nyeri_kronis":                         "required|in:Tidak,Ya",
		"keterangan_memiliki_nyeri_kronis":              "string|max_len:100",
		"memerlukan_edukasi_kesehatan":                  "required|in:Tidak,Ya",
		"keterangan_memerlukan_edukasi_kesehatan":       "string|max_len:100",
		"memerlukan_keterampilkan_khusus":               "required|in:Tidak,Ya",
		"keterangan_memerlukan_keterampilkan_khusus":    "string|max_len:100",
		"nama_pasien_keluarga":                          "string|max_len:50",
		"nip":                                           "required|string|max_len:20",
	}
	return rules
}

// PerencanaanPemulanganStore simpan perencanaan pemulangan; kolom waktu kunci kosong = sekarang.
type PerencanaanPemulanganStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PerencanaanPemulanganData
}

func (r *PerencanaanPemulanganStore) Authorize(ctx http.Context) error { return nil }

func (r *PerencanaanPemulanganStore) Rules(ctx http.Context) map[string]any {
	rules := perencanaanPemulanganRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PerencanaanPemulanganStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PerencanaanPemulanganStore) Payload() PerencanaanPemulanganData {
	return r.PerencanaanPemulanganData
}

func (r *PerencanaanPemulanganStore) DetailValues() map[string][]string { return nil }

// PerencanaanPemulanganUpdate ubah perencanaan pemulangan (PUT); kunci lewat query string.
type PerencanaanPemulanganUpdate struct {
	PerencanaanPemulanganData
}

func (r *PerencanaanPemulanganUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PerencanaanPemulanganUpdate) Rules(ctx http.Context) map[string]any {
	return perencanaanPemulanganRules()
}

func (r *PerencanaanPemulanganUpdate) Payload() PerencanaanPemulanganData {
	return r.PerencanaanPemulanganData
}

func (r *PerencanaanPemulanganUpdate) DetailValues() map[string][]string { return nil }
