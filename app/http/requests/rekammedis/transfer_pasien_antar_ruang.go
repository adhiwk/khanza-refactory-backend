package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// TransferPasienAntarRuangData isian transfer pasien antar ruang.
type TransferPasienAntarRuangData struct {
	TanggalPindah                     string `form:"tanggal_pindah" json:"tanggal_pindah"`
	AsalRuang                         string `form:"asal_ruang" json:"asal_ruang"`
	RuangSelanjutnya                  string `form:"ruang_selanjutnya" json:"ruang_selanjutnya"`
	DiagnosaUtama                     string `form:"diagnosa_utama" json:"diagnosa_utama"`
	DiagnosaSekunder                  string `form:"diagnosa_sekunder" json:"diagnosa_sekunder"`
	IndikasiPindahRuang               string `form:"indikasi_pindah_ruang" json:"indikasi_pindah_ruang"`
	KeteranganIndikasiPindahRuang     string `form:"keterangan_indikasi_pindah_ruang" json:"keterangan_indikasi_pindah_ruang"`
	ProsedurYangSudahDilakukan        string `form:"prosedur_yang_sudah_dilakukan" json:"prosedur_yang_sudah_dilakukan"`
	ObatYangTelahDiberikan            string `form:"obat_yang_telah_diberikan" json:"obat_yang_telah_diberikan"`
	MetodePemindahanPasien            string `form:"metode_pemindahan_pasien" json:"metode_pemindahan_pasien"`
	PeralatanYangMenyertai            string `form:"peralatan_yang_menyertai" json:"peralatan_yang_menyertai"`
	KeteranganPeralatanYangMenyertai  string `form:"keterangan_peralatan_yang_menyertai" json:"keterangan_peralatan_yang_menyertai"`
	PemeriksaanPenunjangYangDilakukan string `form:"pemeriksaan_penunjang_yang_dilakukan" json:"pemeriksaan_penunjang_yang_dilakukan"`
	PasienKeluargaMenyetujui          string `form:"pasien_keluarga_menyetujui" json:"pasien_keluarga_menyetujui"`
	NamaMenyetujui                    string `form:"nama_menyetujui" json:"nama_menyetujui"`
	HubunganMenyetujui                string `form:"hubungan_menyetujui" json:"hubungan_menyetujui"`
	KeluhanUtamaSebelumTransfer       string `form:"keluhan_utama_sebelum_transfer" json:"keluhan_utama_sebelum_transfer"`
	KeadaanUmumSebelumTransfer        string `form:"keadaan_umum_sebelum_transfer" json:"keadaan_umum_sebelum_transfer"`
	TdSebelumTransfer                 string `form:"td_sebelum_transfer" json:"td_sebelum_transfer"`
	NadiSebelumTransfer               string `form:"nadi_sebelum_transfer" json:"nadi_sebelum_transfer"`
	RrSebelumTransfer                 string `form:"rr_sebelum_transfer" json:"rr_sebelum_transfer"`
	SuhuSebelumTransfer               string `form:"suhu_sebelum_transfer" json:"suhu_sebelum_transfer"`
	KeluhanUtamaSesudahTransfer       string `form:"keluhan_utama_sesudah_transfer" json:"keluhan_utama_sesudah_transfer"`
	KeadaanUmumSesudahTransfer        string `form:"keadaan_umum_sesudah_transfer" json:"keadaan_umum_sesudah_transfer"`
	TdSesudahTransfer                 string `form:"td_sesudah_transfer" json:"td_sesudah_transfer"`
	NadiSesudahTransfer               string `form:"nadi_sesudah_transfer" json:"nadi_sesudah_transfer"`
	RrSesudahTransfer                 string `form:"rr_sesudah_transfer" json:"rr_sesudah_transfer"`
	SuhuSesudahTransfer               string `form:"suhu_sesudah_transfer" json:"suhu_sesudah_transfer"`
	NipMenyerahkan                    string `form:"nip_menyerahkan" json:"nip_menyerahkan"`
	NipMenerima                       string `form:"nip_menerima" json:"nip_menerima"`
}

func transferPasienAntarRuangRules() map[string]any {
	rules := map[string]any{
		"tanggal_pindah":                       "date",
		"asal_ruang":                           "string|max_len:30",
		"ruang_selanjutnya":                    "string|max_len:30",
		"diagnosa_utama":                       "string|max_len:100",
		"diagnosa_sekunder":                    "string|max_len:150",
		"indikasi_pindah_ruang":                "in:Kondisi Pasien Stabil,Kondisi Pasien Tidak Ada Perubahan,Kondisi Pasien Memburuk,Fasilitas Kurang Memadai,Fasilitas Butuh Lebih Baik,Tenaga Membutuhkan Yang Lebih Ahli,Tenaga Kurang,Lain-lain",
		"keterangan_indikasi_pindah_ruang":     "string|max_len:100",
		"prosedur_yang_sudah_dilakukan":        "string|max_len:800",
		"obat_yang_telah_diberikan":            "string|max_len:1200",
		"metode_pemindahan_pasien":             "in:Kursi Roda,Tempat Tidur,Brankar,Jalan Sendiri,-",
		"peralatan_yang_menyertai":             "in:Oksigen Portable,Infus,NGT,Syringe Pump,Suction,Kateter Urin,Tidak Ada",
		"keterangan_peralatan_yang_menyertai":  "string|max_len:90",
		"pemeriksaan_penunjang_yang_dilakukan": "string|max_len:700",
		"pasien_keluarga_menyetujui":           "in:Ya,Tidak",
		"nama_menyetujui":                      "string|max_len:50",
		"hubungan_menyetujui":                  "in:Kakak,Adik,Saudara,Keluarga,Kakek,Nenek,Orang Tua,Suami,Istri,Penanggung Jawab,Menantu,Ipar,Mertua,-",
		"keluhan_utama_sebelum_transfer":       "string|max_len:300",
		"keadaan_umum_sebelum_transfer":        "in:Compos Mentis,Gelisah,Delirium,Koma",
		"td_sebelum_transfer":                  "string|max_len:7",
		"nadi_sebelum_transfer":                "string|max_len:5",
		"rr_sebelum_transfer":                  "string|max_len:5",
		"suhu_sebelum_transfer":                "string|max_len:5",
		"keluhan_utama_sesudah_transfer":       "string|max_len:300",
		"keadaan_umum_sesudah_transfer":        "in:Compos Mentis,Gelisah,Delirium,Koma",
		"td_sesudah_transfer":                  "string|max_len:7",
		"nadi_sesudah_transfer":                "string|max_len:5",
		"rr_sesudah_transfer":                  "string|max_len:5",
		"suhu_sesudah_transfer":                "string|max_len:5",
		"nip_menyerahkan":                      "required|string|max_len:20",
		"nip_menerima":                         "required|string|max_len:20",
	}
	return rules
}

// TransferPasienAntarRuangStore simpan transfer pasien antar ruang; kolom waktu kunci kosong = sekarang.
type TransferPasienAntarRuangStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TanggalMasuk string `form:"tanggal_masuk" json:"tanggal_masuk"`
	TransferPasienAntarRuangData
}

func (r *TransferPasienAntarRuangStore) Authorize(ctx http.Context) error { return nil }

func (r *TransferPasienAntarRuangStore) Rules(ctx http.Context) map[string]any {
	rules := transferPasienAntarRuangRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal_masuk"] = "date"
	return rules
}

func (r *TransferPasienAntarRuangStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal_masuk": r.TanggalMasuk}
}

func (r *TransferPasienAntarRuangStore) Payload() TransferPasienAntarRuangData {
	return r.TransferPasienAntarRuangData
}

func (r *TransferPasienAntarRuangStore) DetailValues() map[string][]string { return nil }

// TransferPasienAntarRuangUpdate ubah transfer pasien antar ruang (PUT); kunci lewat query string.
type TransferPasienAntarRuangUpdate struct {
	TransferPasienAntarRuangData
}

func (r *TransferPasienAntarRuangUpdate) Authorize(ctx http.Context) error { return nil }

func (r *TransferPasienAntarRuangUpdate) Rules(ctx http.Context) map[string]any {
	return transferPasienAntarRuangRules()
}

func (r *TransferPasienAntarRuangUpdate) Payload() TransferPasienAntarRuangData {
	return r.TransferPasienAntarRuangData
}

func (r *TransferPasienAntarRuangUpdate) DetailValues() map[string][]string { return nil }
