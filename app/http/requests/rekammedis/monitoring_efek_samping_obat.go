package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// MonitoringEfekSampingObatData isian pelaporan efek samping obat.
type MonitoringEfekSampingObatData struct {
	NoLaporan         string `form:"no_laporan" json:"no_laporan"`
	Tanggal           string `form:"tanggal" json:"tanggal"`
	Profesi           string `form:"profesi" json:"profesi"`
	Nik               string `form:"nik" json:"nik"`
	KdBangsal         string `form:"kd_bangsal" json:"kd_bangsal"`
	BeratBadan        string `form:"berat_badan" json:"berat_badan"`
	PasienHamil       string `form:"pasien_hamil" json:"pasien_hamil"`
	Kesudahan         string `form:"kesudahan" json:"kesudahan"`
	PenyakitLain      string `form:"penyakit_lain" json:"penyakit_lain"`
	PenyakitUtama     string `form:"penyakit_utama" json:"penyakit_utama"`
	TanggalKejadian   string `form:"tanggal_kejadian" json:"tanggal_kejadian"`
	Manifestasi       string `form:"manifestasi" json:"manifestasi"`
	MasalahKualitas   string `form:"masalah_kualitas" json:"masalah_kualitas"`
	RiwayatEso        string `form:"riwayat_eso" json:"riwayat_eso"`
	TanggalKesudahan  string `form:"tanggal_kesudahan" json:"tanggal_kesudahan"`
	HasilKesudahan    string `form:"hasil_kesudahan" json:"hasil_kesudahan"`
	Obat1             string `form:"obat1" json:"obat1"`
	Sedian1           string `form:"sedian1" json:"sedian1"`
	ObatJkn1          string `form:"obat_jkn1" json:"obat_jkn1"`
	Batch1            string `form:"batch1" json:"batch1"`
	Cara1             string `form:"cara1" json:"cara1"`
	Dosis1            string `form:"dosis1" json:"dosis1"`
	TanggalMulai1     string `form:"tanggal_mulai1" json:"tanggal_mulai1"`
	TanggalAkhir1     string `form:"tanggal_akhir1" json:"tanggal_akhir1"`
	Indikasi1         string `form:"indikasi1" json:"indikasi1"`
	Obat2             string `form:"obat2" json:"obat2"`
	Sedian2           string `form:"sedian2" json:"sedian2"`
	ObatJkn2          string `form:"obat_jkn2" json:"obat_jkn2"`
	Batch2            string `form:"batch2" json:"batch2"`
	Cara2             string `form:"cara2" json:"cara2"`
	Dosis2            string `form:"dosis2" json:"dosis2"`
	TanggalMulai2     string `form:"tanggal_mulai2" json:"tanggal_mulai2"`
	TanggalAkhir2     string `form:"tanggal_akhir2" json:"tanggal_akhir2"`
	Indikasi2         string `form:"indikasi2" json:"indikasi2"`
	Obat3             string `form:"obat3" json:"obat3"`
	Sedian3           string `form:"sedian3" json:"sedian3"`
	ObatJkn3          string `form:"obat_jkn3" json:"obat_jkn3"`
	Batch3            string `form:"batch3" json:"batch3"`
	Cara3             string `form:"cara3" json:"cara3"`
	Dosis3            string `form:"dosis3" json:"dosis3"`
	TanggalMulai3     string `form:"tanggal_mulai3" json:"tanggal_mulai3"`
	TanggalAkhir3     string `form:"tanggal_akhir3" json:"tanggal_akhir3"`
	Indikasi3         string `form:"indikasi3" json:"indikasi3"`
	Obat4             string `form:"obat4" json:"obat4"`
	Sedian4           string `form:"sedian4" json:"sedian4"`
	ObatJkn4          string `form:"obat_jkn4" json:"obat_jkn4"`
	Batch4            string `form:"batch4" json:"batch4"`
	Cara4             string `form:"cara4" json:"cara4"`
	Dosis4            string `form:"dosis4" json:"dosis4"`
	TanggalMulai4     string `form:"tanggal_mulai4" json:"tanggal_mulai4"`
	TanggalAkhir4     string `form:"tanggal_akhir4" json:"tanggal_akhir4"`
	Indikasi4         string `form:"indikasi4" json:"indikasi4"`
	Obat5             string `form:"obat5" json:"obat5"`
	Sedian5           string `form:"sedian5" json:"sedian5"`
	ObatJkn5          string `form:"obat_jkn5" json:"obat_jkn5"`
	Batch5            string `form:"batch5" json:"batch5"`
	Cara5             string `form:"cara5" json:"cara5"`
	Dosis5            string `form:"dosis5" json:"dosis5"`
	TanggalMulai5     string `form:"tanggal_mulai5" json:"tanggal_mulai5"`
	TanggalAkhir5     string `form:"tanggal_akhir5" json:"tanggal_akhir5"`
	Indikasi5         string `form:"indikasi5" json:"indikasi5"`
	Obat6             string `form:"obat6" json:"obat6"`
	Sedian6           string `form:"sedian6" json:"sedian6"`
	ObatJkn6          string `form:"obat_jkn6" json:"obat_jkn6"`
	Batch6            string `form:"batch6" json:"batch6"`
	Cara6             string `form:"cara6" json:"cara6"`
	Dosis6            string `form:"dosis6" json:"dosis6"`
	TanggalMulai6     string `form:"tanggal_mulai6" json:"tanggal_mulai6"`
	TanggalAkhir6     string `form:"tanggal_akhir6" json:"tanggal_akhir6"`
	Indikasi6         string `form:"indikasi6" json:"indikasi6"`
	Obat7             string `form:"obat7" json:"obat7"`
	Sedian7           string `form:"sedian7" json:"sedian7"`
	ObatJkn7          string `form:"obat_jkn7" json:"obat_jkn7"`
	Batch7            string `form:"batch7" json:"batch7"`
	Cara7             string `form:"cara7" json:"cara7"`
	Dosis7            string `form:"dosis7" json:"dosis7"`
	TanggalMulai7     string `form:"tanggal_mulai7" json:"tanggal_mulai7"`
	TanggalAkhir7     string `form:"tanggal_akhir7" json:"tanggal_akhir7"`
	Indikasi7         string `form:"indikasi7" json:"indikasi7"`
	Obat8             string `form:"obat8" json:"obat8"`
	Sedian8           string `form:"sedian8" json:"sedian8"`
	ObatJkn8          string `form:"obat_jkn8" json:"obat_jkn8"`
	Batch8            string `form:"batch8" json:"batch8"`
	Cara8             string `form:"cara8" json:"cara8"`
	Dosis8            string `form:"dosis8" json:"dosis8"`
	TanggalMulai8     string `form:"tanggal_mulai8" json:"tanggal_mulai8"`
	TanggalAkhir8     string `form:"tanggal_akhir8" json:"tanggal_akhir8"`
	Indikasi8         string `form:"indikasi8" json:"indikasi8"`
	Obat9             string `form:"obat9" json:"obat9"`
	Sedian9           string `form:"sedian9" json:"sedian9"`
	ObatJkn9          string `form:"obat_jkn9" json:"obat_jkn9"`
	Batch9            string `form:"batch9" json:"batch9"`
	Cara9             string `form:"cara9" json:"cara9"`
	Dosis9            string `form:"dosis9" json:"dosis9"`
	TanggalMulai9     string `form:"tanggal_mulai9" json:"tanggal_mulai9"`
	TanggalAkhir9     string `form:"tanggal_akhir9" json:"tanggal_akhir9"`
	Indikasi9         string `form:"indikasi9" json:"indikasi9"`
	Obat10            string `form:"obat10" json:"obat10"`
	Sedian10          string `form:"sedian10" json:"sedian10"`
	ObatJkn10         string `form:"obat_jkn10" json:"obat_jkn10"`
	Batch10           string `form:"batch10" json:"batch10"`
	Cara10            string `form:"cara10" json:"cara10"`
	Dosis10           string `form:"dosis10" json:"dosis10"`
	TanggalMulai10    string `form:"tanggal_mulai10" json:"tanggal_mulai10"`
	TanggalAkhir10    string `form:"tanggal_akhir10" json:"tanggal_akhir10"`
	Indikasi10        string `form:"indikasi10" json:"indikasi10"`
	ReaksiSetelah     string `form:"reaksi_setelah" json:"reaksi_setelah"`
	ReaksiSama        string `form:"reaksi_sama" json:"reaksi_sama"`
	Naranjo1          string `form:"naranjo1" json:"naranjo1"`
	NilaiNaranjo1     string `form:"nilai_naranjo1" json:"nilai_naranjo1"`
	Naranjo2          string `form:"naranjo2" json:"naranjo2"`
	NilaiNaranjo2     string `form:"nilai_naranjo2" json:"nilai_naranjo2"`
	Naranjo3          string `form:"naranjo3" json:"naranjo3"`
	NilaiNaranjo3     string `form:"nilai_naranjo3" json:"nilai_naranjo3"`
	Naranjo4          string `form:"naranjo4" json:"naranjo4"`
	NilaiNaranjo4     string `form:"nilai_naranjo4" json:"nilai_naranjo4"`
	Naranjo5          string `form:"naranjo5" json:"naranjo5"`
	NilaiNaranjo5     string `form:"nilai_naranjo5" json:"nilai_naranjo5"`
	Naranjo6          string `form:"naranjo6" json:"naranjo6"`
	NilaiNaranjo6     string `form:"nilai_naranjo6" json:"nilai_naranjo6"`
	Naranjo7          string `form:"naranjo7" json:"naranjo7"`
	NilaiNaranjo7     string `form:"nilai_naranjo7" json:"nilai_naranjo7"`
	Naranjo8          string `form:"naranjo8" json:"naranjo8"`
	NilaiNaranjo8     string `form:"nilai_naranjo8" json:"nilai_naranjo8"`
	Naranjo9          string `form:"naranjo9" json:"naranjo9"`
	NilaiNaranjo9     string `form:"nilai_naranjo9" json:"nilai_naranjo9"`
	Naranjo10         string `form:"naranjo10" json:"naranjo10"`
	NilaiNaranjo10    string `form:"nilai_naranjo10" json:"nilai_naranjo10"`
	TotalNilaiNaranjo string `form:"total_nilai_naranjo" json:"total_nilai_naranjo"`
	KategoriNaranjo   string `form:"kategori_naranjo" json:"kategori_naranjo"`
}

func monitoringEfekSampingObatRules() map[string]any {
	rules := map[string]any{
		"no_laporan":          "string|max_len:17",
		"tanggal":             "required|date",
		"profesi":             "required|in:-,Dokter,Perawat/Bidan,Farmasi",
		"nik":                 "string|max_len:20",
		"kd_bangsal":          "string|max_len:5",
		"berat_badan":         "string|max_len:5",
		"pasien_hamil":        "required|in:Tidak,Ya,Tidak Tahu",
		"kesudahan":           "required|in:Sembuh,Meninggal,Sembuh Dengan Gejala Sisa,Belum Sembuh,Tidak Tahu",
		"penyakit_lain":       "required|in:Gangguan Ginjal,Gangguan Hati,Alergi,Kondisi Medis Lainnya,Faktor Industri,Pertanian,Kimia,Lain-Lainnya",
		"penyakit_utama":      "string|max_len:2000",
		"tanggal_kejadian":    "required|date",
		"manifestasi":         "string|max_len:2000",
		"masalah_kualitas":    "string|max_len:2000",
		"riwayat_eso":         "string|max_len:2000",
		"tanggal_kesudahan":   "required|date",
		"hasil_kesudahan":     "required|in:Sembuh,Meninggal,Sembuh Dengan Gejala Sisa,Belum Sembuh,Tidak Tahu",
		"obat1":               "string|max_len:500",
		"sedian1":             "string|max_len:100",
		"obat_jkn1":           "required|in:-,Ya,Tidak",
		"batch1":              "string|max_len:100",
		"cara1":               "string|max_len:100",
		"dosis1":              "string|max_len:100",
		"tanggal_mulai1":      "required|date",
		"tanggal_akhir1":      "required|date",
		"indikasi1":           "string|max_len:100",
		"obat2":               "string|max_len:500",
		"sedian2":             "string|max_len:100",
		"obat_jkn2":           "required|in:-,Ya,Tidak",
		"batch2":              "string|max_len:100",
		"cara2":               "string|max_len:100",
		"dosis2":              "string|max_len:100",
		"tanggal_mulai2":      "required|date",
		"tanggal_akhir2":      "required|date",
		"indikasi2":           "string|max_len:100",
		"obat3":               "string|max_len:500",
		"sedian3":             "string|max_len:100",
		"obat_jkn3":           "required|in:-,Ya,Tidak",
		"batch3":              "string|max_len:100",
		"cara3":               "string|max_len:100",
		"dosis3":              "string|max_len:100",
		"tanggal_mulai3":      "required|date",
		"tanggal_akhir3":      "required|date",
		"indikasi3":           "string|max_len:100",
		"obat4":               "string|max_len:500",
		"sedian4":             "string|max_len:100",
		"obat_jkn4":           "required|in:-,Ya,Tidak",
		"batch4":              "string|max_len:100",
		"cara4":               "string|max_len:100",
		"dosis4":              "string|max_len:100",
		"tanggal_mulai4":      "required|date",
		"tanggal_akhir4":      "required|date",
		"indikasi4":           "string|max_len:100",
		"obat5":               "string|max_len:500",
		"sedian5":             "string|max_len:100",
		"obat_jkn5":           "required|in:-,Ya,Tidak",
		"batch5":              "string|max_len:100",
		"cara5":               "string|max_len:100",
		"dosis5":              "string|max_len:100",
		"tanggal_mulai5":      "required|date",
		"tanggal_akhir5":      "required|date",
		"indikasi5":           "string|max_len:100",
		"obat6":               "string|max_len:500",
		"sedian6":             "string|max_len:100",
		"obat_jkn6":           "required|in:-,Ya,Tidak",
		"batch6":              "string|max_len:100",
		"cara6":               "string|max_len:100",
		"dosis6":              "string|max_len:100",
		"tanggal_mulai6":      "required|date",
		"tanggal_akhir6":      "required|date",
		"indikasi6":           "string|max_len:100",
		"obat7":               "string|max_len:500",
		"sedian7":             "string|max_len:100",
		"obat_jkn7":           "required|in:-,Ya,Tidak",
		"batch7":              "string|max_len:100",
		"cara7":               "string|max_len:100",
		"dosis7":              "string|max_len:100",
		"tanggal_mulai7":      "required|date",
		"tanggal_akhir7":      "required|date",
		"indikasi7":           "string|max_len:100",
		"obat8":               "string|max_len:500",
		"sedian8":             "string|max_len:100",
		"obat_jkn8":           "required|in:-,Ya,Tidak",
		"batch8":              "string|max_len:100",
		"cara8":               "string|max_len:100",
		"dosis8":              "string|max_len:100",
		"tanggal_mulai8":      "required|date",
		"tanggal_akhir8":      "required|date",
		"indikasi8":           "string|max_len:100",
		"obat9":               "string|max_len:500",
		"sedian9":             "string|max_len:100",
		"obat_jkn9":           "required|in:-,Ya,Tidak",
		"batch9":              "string|max_len:100",
		"cara9":               "string|max_len:100",
		"dosis9":              "string|max_len:100",
		"tanggal_mulai9":      "required|date",
		"tanggal_akhir9":      "required|date",
		"indikasi9":           "string|max_len:100",
		"obat10":              "string|max_len:500",
		"sedian10":            "string|max_len:100",
		"obat_jkn10":          "required|in:-,Ya,Tidak",
		"batch10":             "string|max_len:100",
		"cara10":              "string|max_len:100",
		"dosis10":             "string|max_len:100",
		"tanggal_mulai10":     "required|date",
		"tanggal_akhir10":     "required|date",
		"indikasi10":          "string|max_len:100",
		"reaksi_setelah":      "required|in:Tidak,Ya,Tidak Tahu",
		"reaksi_sama":         "required|in:Tidak,Ya,Tidak Tahu",
		"naranjo1":            "required|in:Tidak Tahu,Ya,Tidak",
		"nilai_naranjo1":      "string|max_len:2",
		"naranjo2":            "required|in:Tidak Tahu,Ya,Tidak",
		"nilai_naranjo2":      "string|max_len:2",
		"naranjo3":            "required|in:Tidak Tahu,Ya,Tidak",
		"nilai_naranjo3":      "string|max_len:2",
		"naranjo4":            "required|in:Tidak Tahu,Ya,Tidak",
		"nilai_naranjo4":      "string|max_len:2",
		"naranjo5":            "required|in:Tidak Tahu,Ya,Tidak",
		"nilai_naranjo5":      "string|max_len:2",
		"naranjo6":            "required|in:Tidak Tahu,Ya,Tidak",
		"nilai_naranjo6":      "string|max_len:2",
		"naranjo7":            "required|in:Tidak Tahu,Ya,Tidak",
		"nilai_naranjo7":      "string|max_len:2",
		"naranjo8":            "required|in:Tidak Tahu,Ya,Tidak",
		"nilai_naranjo8":      "string|max_len:2",
		"naranjo9":            "required|in:Tidak Tahu,Ya,Tidak",
		"nilai_naranjo9":      "string|max_len:2",
		"naranjo10":           "required|in:Tidak Tahu,Ya,Tidak",
		"nilai_naranjo10":     "string|max_len:2",
		"total_nilai_naranjo": "string|max_len:2",
		"kategori_naranjo":    "string|max_len:50",
	}
	return rules
}

// MonitoringEfekSampingObatStore simpan pelaporan efek samping obat; kolom waktu kunci kosong = sekarang.
type MonitoringEfekSampingObatStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	MonitoringEfekSampingObatData
}

func (r *MonitoringEfekSampingObatStore) Authorize(ctx http.Context) error { return nil }

func (r *MonitoringEfekSampingObatStore) Rules(ctx http.Context) map[string]any {
	rules := monitoringEfekSampingObatRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *MonitoringEfekSampingObatStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *MonitoringEfekSampingObatStore) Payload() MonitoringEfekSampingObatData {
	return r.MonitoringEfekSampingObatData
}

func (r *MonitoringEfekSampingObatStore) DetailValues() map[string][]string { return nil }

// MonitoringEfekSampingObatUpdate ubah pelaporan efek samping obat (PUT); kunci lewat query string.
type MonitoringEfekSampingObatUpdate struct {
	MonitoringEfekSampingObatData
}

func (r *MonitoringEfekSampingObatUpdate) Authorize(ctx http.Context) error { return nil }

func (r *MonitoringEfekSampingObatUpdate) Rules(ctx http.Context) map[string]any {
	return monitoringEfekSampingObatRules()
}

func (r *MonitoringEfekSampingObatUpdate) Payload() MonitoringEfekSampingObatData {
	return r.MonitoringEfekSampingObatData
}

func (r *MonitoringEfekSampingObatUpdate) DetailValues() map[string][]string { return nil }
