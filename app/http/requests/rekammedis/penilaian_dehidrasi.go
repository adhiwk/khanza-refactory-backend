package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianDehidrasiData isian penilaian derajat dehidrasi.
type PenilaianDehidrasiData struct {
	Penilaian1          string `form:"penilaian1" json:"penilaian1"`
	PenilaianNilai1     *int   `form:"penilaian_nilai1" json:"penilaian_nilai1"`
	Penilaian2          string `form:"penilaian2" json:"penilaian2"`
	PenilaianNilai2     *int   `form:"penilaian_nilai2" json:"penilaian_nilai2"`
	Penilaian3          string `form:"penilaian3" json:"penilaian3"`
	PenilaianNilai3     *int   `form:"penilaian_nilai3" json:"penilaian_nilai3"`
	Penilaian4          string `form:"penilaian4" json:"penilaian4"`
	PenilaianNilai4     *int   `form:"penilaian_nilai4" json:"penilaian_nilai4"`
	Penilaian5          string `form:"penilaian5" json:"penilaian5"`
	PenilaianNilai5     *int   `form:"penilaian_nilai5" json:"penilaian_nilai5"`
	Penilaian6          string `form:"penilaian6" json:"penilaian6"`
	PenilaianNilai6     *int   `form:"penilaian_nilai6" json:"penilaian_nilai6"`
	PenilaianTotalnilai *int   `form:"penilaian_totalnilai" json:"penilaian_totalnilai"`
	HasilPenilaian      string `form:"hasil_penilaian" json:"hasil_penilaian"`
	KdDokter            string `form:"kd_dokter" json:"kd_dokter"`
}

func penilaianDehidrasiRules() map[string]any {
	rules := map[string]any{
		"penilaian1":           "string",
		"penilaian_nilai1":     "int",
		"penilaian2":           "in:Biasa,Cekung,Sangat Cekung",
		"penilaian_nilai2":     "int",
		"penilaian3":           "in:Biasa,Kering,Sangat Kering",
		"penilaian_nilai3":     "int",
		"penilaian4":           "in:< 30 x/menit,30 - 40 x/menit,> 40 x/menit",
		"penilaian_nilai4":     "int",
		"penilaian5":           "in:Baik,Kurang,Jelek",
		"penilaian_nilai5":     "int",
		"penilaian6":           "in:< 120 x/menit,120 - 140 x/menit,> 140 x/menit",
		"penilaian_nilai6":     "int",
		"penilaian_totalnilai": "int",
		"hasil_penilaian":      "string|max_len:200",
		"kd_dokter":            "required|string|max_len:20",
	}
	return rules
}

// PenilaianDehidrasiStore simpan penilaian derajat dehidrasi; kolom waktu kunci kosong = sekarang.
type PenilaianDehidrasiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PenilaianDehidrasiData
}

func (r *PenilaianDehidrasiStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianDehidrasiStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianDehidrasiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PenilaianDehidrasiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PenilaianDehidrasiStore) Payload() PenilaianDehidrasiData { return r.PenilaianDehidrasiData }

func (r *PenilaianDehidrasiStore) DetailValues() map[string][]string { return nil }

// PenilaianDehidrasiUpdate ubah penilaian derajat dehidrasi (PUT); kunci lewat query string.
type PenilaianDehidrasiUpdate struct {
	PenilaianDehidrasiData
}

func (r *PenilaianDehidrasiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianDehidrasiUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianDehidrasiRules()
}

func (r *PenilaianDehidrasiUpdate) Payload() PenilaianDehidrasiData { return r.PenilaianDehidrasiData }

func (r *PenilaianDehidrasiUpdate) DetailValues() map[string][]string { return nil }
