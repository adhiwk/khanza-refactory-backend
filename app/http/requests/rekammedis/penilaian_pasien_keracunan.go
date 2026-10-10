package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianPasienKeracunanData isian penilaian pasien keracunan.
type PenilaianPasienKeracunanData struct {
	Tanggal                  string `form:"tanggal" json:"tanggal"`
	KdDokter                 string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis                string `form:"anamnesis" json:"anamnesis"`
	Hubungan                 string `form:"hubungan" json:"hubungan"`
	TempatKejadian           string `form:"tempat_kejadian" json:"tempat_kejadian"`
	KeteranganTempatKejadian string `form:"keterangan_tempat_kejadian" json:"keterangan_tempat_kejadian"`
	Keluhan                  string `form:"keluhan" json:"keluhan"`
	RiwayatPenyakitSekarang  string `form:"riwayat_penyakit_sekarang" json:"riwayat_penyakit_sekarang"`
	Hamil                    string `form:"hamil" json:"hamil"`
	Menyusui                 string `form:"menyusui" json:"menyusui"`
	Penyebab                 string `form:"penyebab" json:"penyebab"`
	NamaBahan                string `form:"nama_bahan" json:"nama_bahan"`
	JumlahBahan              string `form:"jumlah_bahan" json:"jumlah_bahan"`
	TipePemaparan            string `form:"tipe_pemaparan" json:"tipe_pemaparan"`
	KeteranganTipePemaparan  string `form:"keterangan_tipe_pemaparan" json:"keterangan_tipe_pemaparan"`
	TipeKejadian             string `form:"tipe_kejadian" json:"tipe_kejadian"`
	BauBahan                 string `form:"bau_bahan" json:"bau_bahan"`
	KeteranganBauBahan       string `form:"keterangan_bau_bahan" json:"keterangan_bau_bahan"`
	Pupil                    string `form:"pupil" json:"pupil"`
	KeteranganPupil          string `form:"keterangan_pupil" json:"keterangan_pupil"`
	Kesadaran                string `form:"kesadaran" json:"kesadaran"`
	Td                       string `form:"td" json:"td"`
	Nadi                     string `form:"nadi" json:"nadi"`
	Rr                       string `form:"rr" json:"rr"`
	Suhu                     string `form:"suhu" json:"suhu"`
	Spo                      string `form:"spo" json:"spo"`
	Urine                    string `form:"urine" json:"urine"`
	PengobatanSebelumIgd     string `form:"pengobatan_sebelum_igd" json:"pengobatan_sebelum_igd"`
	Diagnosis                string `form:"diagnosis" json:"diagnosis"`
	PemeriksaanPenunjang     string `form:"pemeriksaan_penunjang" json:"pemeriksaan_penunjang"`
	PenatalaksanaanDiberikan string `form:"penatalaksanaan_diberikan" json:"penatalaksanaan_diberikan"`
	TindakLanjut             string `form:"tindak_lanjut" json:"tindak_lanjut"`
}

func penilaianPasienKeracunanRules() map[string]any {
	rules := map[string]any{
		"tanggal":                    "required|date",
		"kd_dokter":                  "required|string|max_len:20",
		"anamnesis":                  "required|in:Autoanamnesis,Alloanamnesis",
		"hubungan":                   "string|max_len:30",
		"tempat_kejadian":            "in:Rumah,Kantor,Tempat Kerja,Tempat Hiburan,Lain-lain",
		"keterangan_tempat_kejadian": "string|max_len:50",
		"keluhan":                    "string|max_len:2000",
		"riwayat_penyakit_sekarang":  "string|max_len:2000",
		"hamil":                      "in:Ya,Tidak",
		"menyusui":                   "in:Ya,Tidak",
		"penyebab":                   "in:NAPZA,Obat,Obat Tradisional,Makanan/Minuman,Suplemen Makanan/Vitamin,Kosmetik,Bahan Kimia,Pestisida,Gigitan Ular,Binatang Selain Ular,Tumbuhan Beracun,Pencemar Lingkungan,Gas,Tidak Diketahui",
		"nama_bahan":                 "string|max_len:100",
		"jumlah_bahan":               "string|max_len:15",
		"tipe_pemaparan":             "in:Mulut,Mata,Gigitan,Injeksi,Inhalasi,Sengatan,Kulit,Lain-lain",
		"keterangan_tipe_pemaparan":  "string|max_len:50",
		"tipe_kejadian":              "in:Tidak Disengaja,Disengaja,Tidak Diketahui",
		"bau_bahan":                  "in:Tidak Ada,Ada",
		"keterangan_bau_bahan":       "string|max_len:30",
		"pupil":                      "in:Normal,Isokor,Unisokor,Miosis,Midriasis,Lainnya",
		"keterangan_pupil":           "string|max_len:30",
		"kesadaran":                  "required|in:Compos Mentis,Apatis,Somnolen,Sopor,Koma",
		"td":                         "string|max_len:8",
		"nadi":                       "string|max_len:5",
		"rr":                         "string|max_len:5",
		"suhu":                       "string|max_len:5",
		"spo":                        "string|max_len:5",
		"urine":                      "string|max_len:5",
		"pengobatan_sebelum_igd":     "string|max_len:500",
		"diagnosis":                  "string|max_len:500",
		"pemeriksaan_penunjang":      "string|max_len:500",
		"penatalaksanaan_diberikan":  "string|max_len:500",
		"tindak_lanjut":              "in:Rawat Jalan,Rawat Inap,Dirujuk,Pulang APS,Pulang Sembuh,Meninggal",
	}
	return rules
}

// PenilaianPasienKeracunanStore simpan penilaian pasien keracunan; kolom waktu kunci kosong = sekarang.
type PenilaianPasienKeracunanStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianPasienKeracunanData
}

func (r *PenilaianPasienKeracunanStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPasienKeracunanStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianPasienKeracunanRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianPasienKeracunanStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianPasienKeracunanStore) Payload() PenilaianPasienKeracunanData {
	return r.PenilaianPasienKeracunanData
}

func (r *PenilaianPasienKeracunanStore) DetailValues() map[string][]string { return nil }

// PenilaianPasienKeracunanUpdate ubah penilaian pasien keracunan (PUT); kunci lewat query string.
type PenilaianPasienKeracunanUpdate struct {
	PenilaianPasienKeracunanData
}

func (r *PenilaianPasienKeracunanUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPasienKeracunanUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianPasienKeracunanRules()
}

func (r *PenilaianPasienKeracunanUpdate) Payload() PenilaianPasienKeracunanData {
	return r.PenilaianPasienKeracunanData
}

func (r *PenilaianPasienKeracunanUpdate) DetailValues() map[string][]string { return nil }
