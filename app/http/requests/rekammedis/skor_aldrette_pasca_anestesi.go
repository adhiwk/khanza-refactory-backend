package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkorAldrettePascaAnestesiData isian monitoring aldrette pasca anestesi.
type SkorAldrettePascaAnestesiData struct {
	PenilaianSkala1     string `form:"penilaian_skala1" json:"penilaian_skala1"`
	PenilaianNilai1     *int   `form:"penilaian_nilai1" json:"penilaian_nilai1"`
	PenilaianSkala2     string `form:"penilaian_skala2" json:"penilaian_skala2"`
	PenilaianNilai2     *int   `form:"penilaian_nilai2" json:"penilaian_nilai2"`
	PenilaianSkala3     string `form:"penilaian_skala3" json:"penilaian_skala3"`
	PenilaianNilai3     *int   `form:"penilaian_nilai3" json:"penilaian_nilai3"`
	PenilaianSkala4     string `form:"penilaian_skala4" json:"penilaian_skala4"`
	PenilaianNilai4     *int   `form:"penilaian_nilai4" json:"penilaian_nilai4"`
	PenilaianSkala5     string `form:"penilaian_skala5" json:"penilaian_skala5"`
	PenilaianNilai5     *int   `form:"penilaian_nilai5" json:"penilaian_nilai5"`
	PenilaianTotalnilai *int   `form:"penilaian_totalnilai" json:"penilaian_totalnilai"`
	Keluar              string `form:"keluar" json:"keluar"`
	Instruksi           string `form:"instruksi" json:"instruksi"`
	KdDokter            string `form:"kd_dokter" json:"kd_dokter"`
	Nip                 string `form:"nip" json:"nip"`
}

func skorAldrettePascaAnestesiRules() map[string]any {
	rules := map[string]any{
		"penilaian_skala1":     "in:Tidak Sanggup Menggerakan Satupun Anggota Gerak,Sanggup Gerak 2 Anggota Tubuh,Sanggup Gerak 4 Anggota Tubuh",
		"penilaian_nilai1":     "int",
		"penilaian_skala2":     "in:Apnea Atau Napas Tidak Adekuat,Sesak Atau Pernapasan Sedikit Terbatas,Sanggup Bernafas Dalam Serta Disuruh Batuk",
		"penilaian_nilai2":     "int",
		"penilaian_skala3":     "in:Â± 50% Tekanan Darah Pra Anestesi,Â± 20% - 50% Tekanan Darah Pra Anestesi,Â± 20% Tekanan Darah Pra Anestesi",
		"penilaian_nilai3":     "int",
		"penilaian_skala4":     "in:Tidak Ada Respon,Respon Terhadap Panggilan,Sadar Penuh",
		"penilaian_nilai4":     "int",
		"penilaian_skala5":     "in:Cianosis,Pucat,Kemerahan / Normal",
		"penilaian_nilai5":     "int",
		"penilaian_totalnilai": "int",
		"keluar":               "string|max_len:200",
		"instruksi":            "string|max_len:250",
		"kd_dokter":            "required|string|max_len:20",
		"nip":                  "required|string|max_len:20",
	}
	return rules
}

// SkorAldrettePascaAnestesiStore simpan monitoring aldrette pasca anestesi; kolom waktu kunci kosong = sekarang.
type SkorAldrettePascaAnestesiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	SkorAldrettePascaAnestesiData
}

func (r *SkorAldrettePascaAnestesiStore) Authorize(ctx http.Context) error { return nil }

func (r *SkorAldrettePascaAnestesiStore) Rules(ctx http.Context) map[string]any {
	rules := skorAldrettePascaAnestesiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *SkorAldrettePascaAnestesiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *SkorAldrettePascaAnestesiStore) Payload() SkorAldrettePascaAnestesiData {
	return r.SkorAldrettePascaAnestesiData
}

func (r *SkorAldrettePascaAnestesiStore) DetailValues() map[string][]string { return nil }

// SkorAldrettePascaAnestesiUpdate ubah monitoring aldrette pasca anestesi (PUT); kunci lewat query string.
type SkorAldrettePascaAnestesiUpdate struct {
	SkorAldrettePascaAnestesiData
}

func (r *SkorAldrettePascaAnestesiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkorAldrettePascaAnestesiUpdate) Rules(ctx http.Context) map[string]any {
	return skorAldrettePascaAnestesiRules()
}

func (r *SkorAldrettePascaAnestesiUpdate) Payload() SkorAldrettePascaAnestesiData {
	return r.SkorAldrettePascaAnestesiData
}

func (r *SkorAldrettePascaAnestesiUpdate) DetailValues() map[string][]string { return nil }
