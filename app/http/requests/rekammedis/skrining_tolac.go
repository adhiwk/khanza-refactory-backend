package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningTolacData isian skrining TOLAC.
type SkriningTolacData struct {
	Tanggal                   string `form:"tanggal" json:"tanggal"`
	Gpa                       string `form:"gpa" json:"gpa"`
	Diagnosa                  string `form:"diagnosa" json:"diagnosa"`
	JumlahSc                  string `form:"jumlah_sc" json:"jumlah_sc"`
	TahunSc                   string `form:"tahun_sc" json:"tahun_sc"`
	IndikasiSc                string `form:"indikasi_sc" json:"indikasi_sc"`
	JenisInsisi               string `form:"jenis_insisi" json:"jenis_insisi"`
	RiwayatPervaginam         string `form:"riwayat_pervaginam" json:"riwayat_pervaginam"`
	TbjGram                   string `form:"tbj_gram" json:"tbj_gram"`
	PresentasiJanin           string `form:"presentasi_janin" json:"presentasi_janin"`
	InklusiRiwayatSc          string `form:"inklusi_riwayat_sc" json:"inklusi_riwayat_sc"`
	InklusiPanggulAdekuat     string `form:"inklusi_panggul_adekuat" json:"inklusi_panggul_adekuat"`
	InklusiJaninTunggalKepala string `form:"inklusi_janin_tunggal_kepala" json:"inklusi_janin_tunggal_kepala"`
	InklusiTbjSesuai          string `form:"inklusi_tbj_sesuai" json:"inklusi_tbj_sesuai"`
	EksklusiScKlasikRuptur    string `form:"eksklusi_sc_klasik_ruptur" json:"eksklusi_sc_klasik_ruptur"`
	EksklusiSc2x              string `form:"eksklusi_sc_2x" json:"eksklusi_sc_2x"`
	EksklusiPlasentaPrevia    string `form:"eksklusi_plasenta_previa" json:"eksklusi_plasenta_previa"`
	Kesimpulan                string `form:"kesimpulan" json:"kesimpulan"`
	EdukasiDiberikan          string `form:"edukasi_diberikan" json:"edukasi_diberikan"`
	Keterangan                string `form:"keterangan" json:"keterangan"`
	KdDokter                  string `form:"kd_dokter" json:"kd_dokter"`
}

func skriningTolacRules() map[string]any {
	rules := map[string]any{
		"tanggal":                      "date",
		"gpa":                          "string|max_len:15",
		"diagnosa":                     "string|max_len:100",
		"jumlah_sc":                    "string|max_len:1",
		"tahun_sc":                     "string|max_len:4",
		"indikasi_sc":                  "string|max_len:50",
		"jenis_insisi":                 "in:-,Transversal Rendah,Klasik,Tidak Diketahui",
		"riwayat_pervaginam":           "in:Sebelum & Sesudah SC,Sesudah SC Saja,Sebelum SC Saja,Tidak Pernah",
		"tbj_gram":                     "string|max_len:5",
		"presentasi_janin":             "string|max_len:30",
		"inklusi_riwayat_sc":           "in:Ya,Tidak",
		"inklusi_panggul_adekuat":      "in:Ya,Tidak",
		"inklusi_janin_tunggal_kepala": "in:Ya,Tidak",
		"inklusi_tbj_sesuai":           "in:Ya,Tidak",
		"eksklusi_sc_klasik_ruptur":    "in:Ya,Tidak",
		"eksklusi_sc_2x":               "in:Ya,Tidak",
		"eksklusi_plasenta_previa":     "in:Ya,Tidak",
		"kesimpulan":                   "in:Kandidat TOLAC,Rujuk SC Elektif",
		"edukasi_diberikan":            "in:Ya,Tidak",
		"keterangan":                   "string|max_len:100",
		"kd_dokter":                    "string|max_len:15",
	}
	return rules
}

// SkriningTolacStore simpan skrining TOLAC; kolom waktu kunci kosong = sekarang.
type SkriningTolacStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningTolacData
}

func (r *SkriningTolacStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningTolacStore) Rules(ctx http.Context) map[string]any {
	rules := skriningTolacRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningTolacStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningTolacStore) Payload() SkriningTolacData { return r.SkriningTolacData }

func (r *SkriningTolacStore) DetailValues() map[string][]string { return nil }

// SkriningTolacUpdate ubah skrining TOLAC (PUT); kunci lewat query string.
type SkriningTolacUpdate struct {
	SkriningTolacData
}

func (r *SkriningTolacUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningTolacUpdate) Rules(ctx http.Context) map[string]any { return skriningTolacRules() }

func (r *SkriningTolacUpdate) Payload() SkriningTolacData { return r.SkriningTolacData }

func (r *SkriningTolacUpdate) DetailValues() map[string][]string { return nil }
