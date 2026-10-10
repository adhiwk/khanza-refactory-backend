package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningKesehatanGigiMulutRemajaData isian skrining kesehatan gigi mulut remaja.
type SkriningKesehatanGigiMulutRemajaData struct {
	Tanggal                    string `form:"tanggal" json:"tanggal"`
	PernahPemeriksaanGigimulut string `form:"pernah_pemeriksaan_gigimulut" json:"pernah_pemeriksaan_gigimulut"`
	JumlahGigiTumbuh           string `form:"jumlah_gigi_tumbuh" json:"jumlah_gigi_tumbuh"`
	KondisiKebersihanGigimulut string `form:"kondisi_kebersihan_gigimulut" json:"kondisi_kebersihan_gigimulut"`
	PunyaGigiBerlubang         string `form:"punya_gigi_berlubang" json:"punya_gigi_berlubang"`
	PernahGusiBerdarah         string `form:"pernah_gusi_berdarah" json:"pernah_gusi_berdarah"`
	PunyaKarangGigi            string `form:"punya_karang_gigi" json:"punya_karang_gigi"`
	GigiDepanTidakTeratur      string `form:"gigi_depan_tidak_teratur" json:"gigi_depan_tidak_teratur"`
	MenyikatGigiSebelumTidur   string `form:"menyikat_gigi_sebelum_tidur" json:"menyikat_gigi_sebelum_tidur"`
	PunyaSariawan              string `form:"punya_sariawan" json:"punya_sariawan"`
	PemeriksaanFisik           string `form:"pemeriksaan_fisik" json:"pemeriksaan_fisik"`
	PemeriksaanPenunjang       string `form:"pemeriksaan_penunjang" json:"pemeriksaan_penunjang"`
	HasilSkrining              string `form:"hasil_skrining" json:"hasil_skrining"`
	Keterangan                 string `form:"keterangan" json:"keterangan"`
	Nip                        string `form:"nip" json:"nip"`
}

func skriningKesehatanGigiMulutRemajaRules() map[string]any {
	rules := map[string]any{
		"tanggal":                      "required|date",
		"pernah_pemeriksaan_gigimulut": "in:Pernah,Tidak Pernah",
		"jumlah_gigi_tumbuh":           "in:<20,>20",
		"kondisi_kebersihan_gigimulut": "in:Bersih,Kotor",
		"punya_gigi_berlubang":         "in:Ya,Tidak",
		"pernah_gusi_berdarah":         "in:Ya,Tidak",
		"punya_karang_gigi":            "in:Ya,Tidak",
		"gigi_depan_tidak_teratur":     "in:Ya,Tidak",
		"menyikat_gigi_sebelum_tidur":  "in:Ya,Tidak",
		"punya_sariawan":               "in:Ya,Tidak",
		"pemeriksaan_fisik":            "string|max_len:600",
		"pemeriksaan_penunjang":        "string|max_len:600",
		"hasil_skrining":               "string|max_len:50",
		"keterangan":                   "string|max_len:100",
		"nip":                          "required|string|max_len:20",
	}
	return rules
}

// SkriningKesehatanGigiMulutRemajaStore simpan skrining kesehatan gigi mulut remaja; kolom waktu kunci kosong = sekarang.
type SkriningKesehatanGigiMulutRemajaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningKesehatanGigiMulutRemajaData
}

func (r *SkriningKesehatanGigiMulutRemajaStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningKesehatanGigiMulutRemajaStore) Rules(ctx http.Context) map[string]any {
	rules := skriningKesehatanGigiMulutRemajaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningKesehatanGigiMulutRemajaStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *SkriningKesehatanGigiMulutRemajaStore) Payload() SkriningKesehatanGigiMulutRemajaData {
	return r.SkriningKesehatanGigiMulutRemajaData
}

func (r *SkriningKesehatanGigiMulutRemajaStore) DetailValues() map[string][]string { return nil }

// SkriningKesehatanGigiMulutRemajaUpdate ubah skrining kesehatan gigi mulut remaja (PUT); kunci lewat query string.
type SkriningKesehatanGigiMulutRemajaUpdate struct {
	SkriningKesehatanGigiMulutRemajaData
}

func (r *SkriningKesehatanGigiMulutRemajaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningKesehatanGigiMulutRemajaUpdate) Rules(ctx http.Context) map[string]any {
	return skriningKesehatanGigiMulutRemajaRules()
}

func (r *SkriningKesehatanGigiMulutRemajaUpdate) Payload() SkriningKesehatanGigiMulutRemajaData {
	return r.SkriningKesehatanGigiMulutRemajaData
}

func (r *SkriningKesehatanGigiMulutRemajaUpdate) DetailValues() map[string][]string { return nil }
