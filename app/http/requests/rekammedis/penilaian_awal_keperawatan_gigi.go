package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianAwalKeperawatanGigiData isian penilaian awal keperawatan gigi.
type PenilaianAwalKeperawatanGigiData struct {
	Tanggal                 string `form:"tanggal" json:"tanggal"`
	Informasi               string `form:"informasi" json:"informasi"`
	Td                      string `form:"td" json:"td"`
	Nadi                    string `form:"nadi" json:"nadi"`
	Rr                      string `form:"rr" json:"rr"`
	Suhu                    string `form:"suhu" json:"suhu"`
	Bb                      string `form:"bb" json:"bb"`
	Tb                      string `form:"tb" json:"tb"`
	Bmi                     string `form:"bmi" json:"bmi"`
	KeluhanUtama            string `form:"keluhan_utama" json:"keluhan_utama"`
	RiwayatPenyakit         string `form:"riwayat_penyakit" json:"riwayat_penyakit"`
	KetRiwayatPenyakit      string `form:"ket_riwayat_penyakit" json:"ket_riwayat_penyakit"`
	Alergi                  string `form:"alergi" json:"alergi"`
	RiwayatPerawatanGigi    string `form:"riwayat_perawatan_gigi" json:"riwayat_perawatan_gigi"`
	KetRiwayatPerawatanGigi string `form:"ket_riwayat_perawatan_gigi" json:"ket_riwayat_perawatan_gigi"`
	KebiasaanSikatGigi      string `form:"kebiasaan_sikat_gigi" json:"kebiasaan_sikat_gigi"`
	KebiasaanLain           string `form:"kebiasaan_lain" json:"kebiasaan_lain"`
	KetKebiasaanLain        string `form:"ket_kebiasaan_lain" json:"ket_kebiasaan_lain"`
	ObatYangDiminumSaatini  string `form:"obat_yang_diminum_saatini" json:"obat_yang_diminum_saatini"`
	AlatBantu               string `form:"alat_bantu" json:"alat_bantu"`
	KetAlatBantu            string `form:"ket_alat_bantu" json:"ket_alat_bantu"`
	Prothesa                string `form:"prothesa" json:"prothesa"`
	KetPro                  string `form:"ket_pro" json:"ket_pro"`
	StatusPsiko             string `form:"status_psiko" json:"status_psiko"`
	KetPsiko                string `form:"ket_psiko" json:"ket_psiko"`
	HubKeluarga             string `form:"hub_keluarga" json:"hub_keluarga"`
	TinggalDengan           string `form:"tinggal_dengan" json:"tinggal_dengan"`
	KetTinggal              string `form:"ket_tinggal" json:"ket_tinggal"`
	Ekonomi                 string `form:"ekonomi" json:"ekonomi"`
	Budaya                  string `form:"budaya" json:"budaya"`
	KetBudaya               string `form:"ket_budaya" json:"ket_budaya"`
	Edukasi                 string `form:"edukasi" json:"edukasi"`
	KetEdukasi              string `form:"ket_edukasi" json:"ket_edukasi"`
	BerjalanA               string `form:"berjalan_a" json:"berjalan_a"`
	BerjalanB               string `form:"berjalan_b" json:"berjalan_b"`
	BerjalanC               string `form:"berjalan_c" json:"berjalan_c"`
	Hasil                   string `form:"hasil" json:"hasil"`
	Lapor                   string `form:"lapor" json:"lapor"`
	KetLapor                string `form:"ket_lapor" json:"ket_lapor"`
	Nyeri                   string `form:"nyeri" json:"nyeri"`
	Lokasi                  string `form:"lokasi" json:"lokasi"`
	SkalaNyeri              string `form:"skala_nyeri" json:"skala_nyeri"`
	Durasi                  string `form:"durasi" json:"durasi"`
	Frekuensi               string `form:"frekuensi" json:"frekuensi"`
	NyeriHilang             string `form:"nyeri_hilang" json:"nyeri_hilang"`
	KetNyeri                string `form:"ket_nyeri" json:"ket_nyeri"`
	PadaDokter              string `form:"pada_dokter" json:"pada_dokter"`
	KetDokter               string `form:"ket_dokter" json:"ket_dokter"`
	KebersihanMulut         string `form:"kebersihan_mulut" json:"kebersihan_mulut"`
	MukosaMulut             string `form:"mukosa_mulut" json:"mukosa_mulut"`
	Karies                  string `form:"karies" json:"karies"`
	KarangGigi              string `form:"karang_gigi" json:"karang_gigi"`
	Gingiva                 string `form:"gingiva" json:"gingiva"`
	Palatum                 string `form:"palatum" json:"palatum"`
	Rencana                 string `form:"rencana" json:"rencana"`
	Nip                     string `form:"nip" json:"nip"`
}

// PenilaianAwalKeperawatanGigiDetail daftar kode detail (diganti seluruhnya saat simpan/ubah).
type PenilaianAwalKeperawatanGigiDetail struct {
	KodeMasalah []string `form:"kode_masalah" json:"kode_masalah"`
	KodeRencana []string `form:"kode_rencana" json:"kode_rencana"`
}

func penilaianAwalKeperawatanGigiRules() map[string]any {
	rules := map[string]any{
		"tanggal":                    "required|date",
		"informasi":                  "required|in:Autoanamnesis,Alloanamnesis",
		"td":                         "string|max_len:8",
		"nadi":                       "string|max_len:5",
		"rr":                         "string|max_len:5",
		"suhu":                       "string|max_len:5",
		"bb":                         "string|max_len:5",
		"tb":                         "string|max_len:5",
		"bmi":                        "string|max_len:10",
		"keluhan_utama":              "string|max_len:150",
		"riwayat_penyakit":           "in:Tidak Ada,Diabetes Melitus,Hipertensi,Penyakit Jantung,HIV,Hepatitis,Haemophilia,Lain-lain",
		"ket_riwayat_penyakit":       "string|max_len:30",
		"alergi":                     "string|max_len:25",
		"riwayat_perawatan_gigi":     "required|string",
		"ket_riwayat_perawatan_gigi": "string|max_len:50",
		"kebiasaan_sikat_gigi":       "required|in:1x,2x,3x,Mandi,Setelah Makan,Sebelum Tidur",
		"kebiasaan_lain":             "in:Tidak ada,Minum kopi/teh,Minum alkohol,Bruxism,Menggigit pensil,Mengunyah 1 sisi rahang,Merokok,Lain-lain",
		"ket_kebiasaan_lain":         "string|max_len:30",
		"obat_yang_diminum_saatini":  "string|max_len:100",
		"alat_bantu":                 "required|in:Tidak,Ya",
		"ket_alat_bantu":             "string|max_len:30",
		"prothesa":                   "required|in:Tidak,Ya",
		"ket_pro":                    "string|max_len:50",
		"status_psiko":               "required|in:Tenang,Takut,Cemas,Depresi,Lain-lain",
		"ket_psiko":                  "string|max_len:70",
		"hub_keluarga":               "required|in:Baik,Tidak Baik",
		"tinggal_dengan":             "required|in:Sendiri,Orang Tua,Suami / Istri,Lainnya",
		"ket_tinggal":                "string|max_len:40",
		"ekonomi":                    "required|in:Baik,Cukup,Kurang",
		"budaya":                     "required|in:Tidak Ada,Ada",
		"ket_budaya":                 "string|max_len:50",
		"edukasi":                    "required|in:Pasien,Keluarga",
		"ket_edukasi":                "string|max_len:50",
		"berjalan_a":                 "required|in:Ya,Tidak",
		"berjalan_b":                 "required|in:Ya,Tidak",
		"berjalan_c":                 "required|in:Ya,Tidak",
		"hasil":                      "required|in:Tidak beresiko (tidak ditemukan a dan b),Resiko rendah (ditemukan a/b),Resiko tinggi (ditemukan a dan b)",
		"lapor":                      "required|in:Ya,Tidak",
		"ket_lapor":                  "string|max_len:15",
		"nyeri":                      "required|in:Tidak Ada Nyeri,Nyeri Akut,Nyeri Kronis",
		"lokasi":                     "string|max_len:50",
		"skala_nyeri":                "required|in:0,1,2,3,4,5,6,7,8,9,10",
		"durasi":                     "string|max_len:25",
		"frekuensi":                  "string|max_len:25",
		"nyeri_hilang":               "required|in:Istirahat,Medengar Musik,Minum Obat,Tidak ada nyeri,Lain-lain",
		"ket_nyeri":                  "string|max_len:40",
		"pada_dokter":                "required|in:Tidak,Ya",
		"ket_dokter":                 "string|max_len:15",
		"kebersihan_mulut":           "required|in:Baik,Cukup,Kurang",
		"mukosa_mulut":               "required|in:Normal,Pigmentasi,Radang",
		"karies":                     "required|in:Ada,Tidak",
		"karang_gigi":                "required|in:Ada,Tidak",
		"gingiva":                    "required|in:Normal,Radang",
		"palatum":                    "required|in:Normal,Radang",
		"rencana":                    "string|max_len:200",
		"nip":                        "required|string|max_len:20",
	}
	rules["kode_masalah"] = "slice"
	rules["kode_masalah.*"] = "string|max_len:20"
	rules["kode_rencana"] = "slice"
	rules["kode_rencana.*"] = "string|max_len:20"
	return rules
}

// PenilaianAwalKeperawatanGigiStore simpan penilaian awal keperawatan gigi; kolom waktu kunci kosong = sekarang.
type PenilaianAwalKeperawatanGigiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianAwalKeperawatanGigiData
	PenilaianAwalKeperawatanGigiDetail
}

func (r *PenilaianAwalKeperawatanGigiStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanGigiStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianAwalKeperawatanGigiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianAwalKeperawatanGigiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianAwalKeperawatanGigiStore) Payload() PenilaianAwalKeperawatanGigiData {
	return r.PenilaianAwalKeperawatanGigiData
}

func (r *PenilaianAwalKeperawatanGigiStore) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}

// PenilaianAwalKeperawatanGigiUpdate ubah penilaian awal keperawatan gigi (PUT); kunci lewat query string.
type PenilaianAwalKeperawatanGigiUpdate struct {
	PenilaianAwalKeperawatanGigiData
	PenilaianAwalKeperawatanGigiDetail
}

func (r *PenilaianAwalKeperawatanGigiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanGigiUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianAwalKeperawatanGigiRules()
}

func (r *PenilaianAwalKeperawatanGigiUpdate) Payload() PenilaianAwalKeperawatanGigiData {
	return r.PenilaianAwalKeperawatanGigiData
}

func (r *PenilaianAwalKeperawatanGigiUpdate) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}
