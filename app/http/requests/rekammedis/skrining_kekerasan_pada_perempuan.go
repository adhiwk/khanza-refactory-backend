package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningKekerasanPadaPerempuanData isian skrining kekerasan pada perempuan.
type SkriningKekerasanPadaPerempuanData struct {
	Tanggal                               string `form:"tanggal" json:"tanggal"`
	MenggambarkanHubungan                 string `form:"menggambarkan_hubungan" json:"menggambarkan_hubungan"`
	SkorMenggambarkanHubungan             string `form:"skor_menggambarkan_hubungan" json:"skor_menggambarkan_hubungan"`
	BerdebatDenganPasangan                string `form:"berdebat_dengan_pasangan" json:"berdebat_dengan_pasangan"`
	SkorBerdebatDenganPasangan            string `form:"skor_berdebat_dengan_pasangan" json:"skor_berdebat_dengan_pasangan"`
	PertengkaranMembuatSedih              string `form:"pertengkaran_membuat_sedih" json:"pertengkaran_membuat_sedih"`
	SkorPertengkaranMembuatSedih          string `form:"skor_pertengkaran_membuat_sedih" json:"skor_pertengkaran_membuat_sedih"`
	PertengkaranMenghasilkanPukulan       string `form:"pertengkaran_menghasilkan_pukulan" json:"pertengkaran_menghasilkan_pukulan"`
	SkorPertengkaranMenghasilkanPukulan   string `form:"skor_pertengkaran_menghasilkan_pukulan" json:"skor_pertengkaran_menghasilkan_pukulan"`
	PernahMerasaTakutDenganPasangan       string `form:"pernah_merasa_takut_dengan_pasangan" json:"pernah_merasa_takut_dengan_pasangan"`
	SkorPernahMerasaTakutDenganPasangan   string `form:"skor_pernah_merasa_takut_dengan_pasangan" json:"skor_pernah_merasa_takut_dengan_pasangan"`
	PasanganMelecehkanSecaraFisik         string `form:"pasangan_melecehkan_secara_fisik" json:"pasangan_melecehkan_secara_fisik"`
	SkorPasanganMelecehkanSecaraFisik     string `form:"skor_pasangan_melecehkan_secara_fisik" json:"skor_pasangan_melecehkan_secara_fisik"`
	PasanganMelecehkanSecaraImosional     string `form:"pasangan_melecehkan_secara_imosional" json:"pasangan_melecehkan_secara_imosional"`
	SkorPasanganMelecehkanSecaraImosional string `form:"skor_pasangan_melecehkan_secara_imosional" json:"skor_pasangan_melecehkan_secara_imosional"`
	PasanganMelecehkanSecaraSeksual       string `form:"pasangan_melecehkan_secara_seksual" json:"pasangan_melecehkan_secara_seksual"`
	SkorPasanganMelecehkanSecaraSeksual   string `form:"skor_pasangan_melecehkan_secara_seksual" json:"skor_pasangan_melecehkan_secara_seksual"`
	Totalskor                             string `form:"totalskor" json:"totalskor"`
	HasilSkrining                         string `form:"hasil_skrining" json:"hasil_skrining"`
	Nip                                   string `form:"nip" json:"nip"`
}

func skriningKekerasanPadaPerempuanRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                   "required|date",
		"menggambarkan_hubungan":                    "in:Tidak Ada Ketegangan,Beberapa Ketegangan,Banyak Ketegangan",
		"skor_menggambarkan_hubungan":               "required|in:1,2,3",
		"berdebat_dengan_pasangan":                  "in:Tidak Ada Kesulitan,Beberapa Kesulitan,Kesulitan Besar",
		"skor_berdebat_dengan_pasangan":             "required|in:1,2,3",
		"pertengkaran_membuat_sedih":                "in:Tidak Pernah,Kadang-kadang,Sering",
		"skor_pertengkaran_membuat_sedih":           "required|in:1,2,3",
		"pertengkaran_menghasilkan_pukulan":         "in:Tidak Pernah,Kadang-kadang,Sering",
		"skor_pertengkaran_menghasilkan_pukulan":    "required|in:1,2,3",
		"pernah_merasa_takut_dengan_pasangan":       "in:Tidak Pernah,Kadang-kadang,Sering",
		"skor_pernah_merasa_takut_dengan_pasangan":  "required|in:1,2,3",
		"pasangan_melecehkan_secara_fisik":          "in:Tidak Pernah,Kadang-kadang,Sering",
		"skor_pasangan_melecehkan_secara_fisik":     "required|in:1,2,3",
		"pasangan_melecehkan_secara_imosional":      "in:Tidak Pernah,Kadang-kadang,Sering",
		"skor_pasangan_melecehkan_secara_imosional": "required|in:1,2,3",
		"pasangan_melecehkan_secara_seksual":        "in:Tidak Pernah,Kadang-kadang,Sering",
		"skor_pasangan_melecehkan_secara_seksual":   "required|in:1,2,3",
		"totalskor":      "string|max_len:3",
		"hasil_skrining": "required|in:Pasien Tidak Terindikasi Mengalami Kekerasan,Pasien Terindikasi Mengalami Kekerasan",
		"nip":            "required|string|max_len:20",
	}
	return rules
}

// SkriningKekerasanPadaPerempuanStore simpan skrining kekerasan pada perempuan; kolom waktu kunci kosong = sekarang.
type SkriningKekerasanPadaPerempuanStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningKekerasanPadaPerempuanData
}

func (r *SkriningKekerasanPadaPerempuanStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningKekerasanPadaPerempuanStore) Rules(ctx http.Context) map[string]any {
	rules := skriningKekerasanPadaPerempuanRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningKekerasanPadaPerempuanStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *SkriningKekerasanPadaPerempuanStore) Payload() SkriningKekerasanPadaPerempuanData {
	return r.SkriningKekerasanPadaPerempuanData
}

func (r *SkriningKekerasanPadaPerempuanStore) DetailValues() map[string][]string { return nil }

// SkriningKekerasanPadaPerempuanUpdate ubah skrining kekerasan pada perempuan (PUT); kunci lewat query string.
type SkriningKekerasanPadaPerempuanUpdate struct {
	SkriningKekerasanPadaPerempuanData
}

func (r *SkriningKekerasanPadaPerempuanUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningKekerasanPadaPerempuanUpdate) Rules(ctx http.Context) map[string]any {
	return skriningKekerasanPadaPerempuanRules()
}

func (r *SkriningKekerasanPadaPerempuanUpdate) Payload() SkriningKekerasanPadaPerempuanData {
	return r.SkriningKekerasanPadaPerempuanData
}

func (r *SkriningKekerasanPadaPerempuanUpdate) DetailValues() map[string][]string { return nil }
