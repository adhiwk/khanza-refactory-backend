package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianLanjutanResikoJatuhDewasaData isian penilaian lanjutan risiko jatuh dewasa.
type PenilaianLanjutanResikoJatuhDewasaData struct {
	PenilaianJatuhmorseSkala1     string `form:"penilaian_jatuhmorse_skala1" json:"penilaian_jatuhmorse_skala1"`
	PenilaianJatuhmorseNilai1     *int   `form:"penilaian_jatuhmorse_nilai1" json:"penilaian_jatuhmorse_nilai1"`
	PenilaianJatuhmorseSkala2     string `form:"penilaian_jatuhmorse_skala2" json:"penilaian_jatuhmorse_skala2"`
	PenilaianJatuhmorseNilai2     *int   `form:"penilaian_jatuhmorse_nilai2" json:"penilaian_jatuhmorse_nilai2"`
	PenilaianJatuhmorseSkala3     string `form:"penilaian_jatuhmorse_skala3" json:"penilaian_jatuhmorse_skala3"`
	PenilaianJatuhmorseNilai3     *int   `form:"penilaian_jatuhmorse_nilai3" json:"penilaian_jatuhmorse_nilai3"`
	PenilaianJatuhmorseSkala4     string `form:"penilaian_jatuhmorse_skala4" json:"penilaian_jatuhmorse_skala4"`
	PenilaianJatuhmorseNilai4     *int   `form:"penilaian_jatuhmorse_nilai4" json:"penilaian_jatuhmorse_nilai4"`
	PenilaianJatuhmorseSkala5     string `form:"penilaian_jatuhmorse_skala5" json:"penilaian_jatuhmorse_skala5"`
	PenilaianJatuhmorseNilai5     *int   `form:"penilaian_jatuhmorse_nilai5" json:"penilaian_jatuhmorse_nilai5"`
	PenilaianJatuhmorseSkala6     string `form:"penilaian_jatuhmorse_skala6" json:"penilaian_jatuhmorse_skala6"`
	PenilaianJatuhmorseNilai6     *int   `form:"penilaian_jatuhmorse_nilai6" json:"penilaian_jatuhmorse_nilai6"`
	PenilaianJatuhmorseTotalnilai *int   `form:"penilaian_jatuhmorse_totalnilai" json:"penilaian_jatuhmorse_totalnilai"`
	HasilSkrining                 string `form:"hasil_skrining" json:"hasil_skrining"`
	Saran                         string `form:"saran" json:"saran"`
	Nip                           string `form:"nip" json:"nip"`
}

func penilaianLanjutanResikoJatuhDewasaRules() map[string]any {
	rules := map[string]any{
		"penilaian_jatuhmorse_skala1":     "in:Tidak,Ya",
		"penilaian_jatuhmorse_nilai1":     "int",
		"penilaian_jatuhmorse_skala2":     "in:Tidak,Ya",
		"penilaian_jatuhmorse_nilai2":     "int",
		"penilaian_jatuhmorse_skala3":     "in:Tidak Ada/Kursi Roda/Perawat/Tirah Baring,Tongkat/Alat Penopang,Berpegangan Pada Perabot",
		"penilaian_jatuhmorse_nilai3":     "int",
		"penilaian_jatuhmorse_skala4":     "in:Tidak,Ya",
		"penilaian_jatuhmorse_nilai4":     "int",
		"penilaian_jatuhmorse_skala5":     "in:Normal/Tirah Baring/Imobilisasi,Lemah,Terganggu",
		"penilaian_jatuhmorse_nilai5":     "int",
		"penilaian_jatuhmorse_skala6":     "in:Sadar Akan Kemampuan Diri Sendiri,Sering Lupa Akan Keterbatasan Yang Dimiliki",
		"penilaian_jatuhmorse_nilai6":     "int",
		"penilaian_jatuhmorse_totalnilai": "int",
		"hasil_skrining":                  "string|max_len:200",
		"saran":                           "string|max_len:200",
		"nip":                             "required|string|max_len:20",
	}
	return rules
}

// PenilaianLanjutanResikoJatuhDewasaStore simpan penilaian lanjutan risiko jatuh dewasa; kolom waktu kunci kosong = sekarang.
type PenilaianLanjutanResikoJatuhDewasaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PenilaianLanjutanResikoJatuhDewasaData
}

func (r *PenilaianLanjutanResikoJatuhDewasaStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianLanjutanResikoJatuhDewasaStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianLanjutanResikoJatuhDewasaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PenilaianLanjutanResikoJatuhDewasaStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PenilaianLanjutanResikoJatuhDewasaStore) Payload() PenilaianLanjutanResikoJatuhDewasaData {
	return r.PenilaianLanjutanResikoJatuhDewasaData
}

func (r *PenilaianLanjutanResikoJatuhDewasaStore) DetailValues() map[string][]string { return nil }

// PenilaianLanjutanResikoJatuhDewasaUpdate ubah penilaian lanjutan risiko jatuh dewasa (PUT); kunci lewat query string.
type PenilaianLanjutanResikoJatuhDewasaUpdate struct {
	PenilaianLanjutanResikoJatuhDewasaData
}

func (r *PenilaianLanjutanResikoJatuhDewasaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianLanjutanResikoJatuhDewasaUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianLanjutanResikoJatuhDewasaRules()
}

func (r *PenilaianLanjutanResikoJatuhDewasaUpdate) Payload() PenilaianLanjutanResikoJatuhDewasaData {
	return r.PenilaianLanjutanResikoJatuhDewasaData
}

func (r *PenilaianLanjutanResikoJatuhDewasaUpdate) DetailValues() map[string][]string { return nil }
