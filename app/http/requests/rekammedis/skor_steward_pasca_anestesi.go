package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkorStewardPascaAnestesiData isian monitoring steward pasca anestesi.
type SkorStewardPascaAnestesiData struct {
	PenilaianSkala1     string `form:"penilaian_skala1" json:"penilaian_skala1"`
	PenilaianNilai1     *int   `form:"penilaian_nilai1" json:"penilaian_nilai1"`
	PenilaianSkala2     string `form:"penilaian_skala2" json:"penilaian_skala2"`
	PenilaianNilai2     *int   `form:"penilaian_nilai2" json:"penilaian_nilai2"`
	PenilaianSkala3     string `form:"penilaian_skala3" json:"penilaian_skala3"`
	PenilaianNilai3     *int   `form:"penilaian_nilai3" json:"penilaian_nilai3"`
	PenilaianTotalnilai *int   `form:"penilaian_totalnilai" json:"penilaian_totalnilai"`
	Keluar              string `form:"keluar" json:"keluar"`
	Instruksi           string `form:"instruksi" json:"instruksi"`
	KdDokter            string `form:"kd_dokter" json:"kd_dokter"`
	Nip                 string `form:"nip" json:"nip"`
}

func skorStewardPascaAnestesiRules() map[string]any {
	rules := map[string]any{
		"penilaian_skala1":     "in:Belum Respon,Bangun Jika Dipanggil,Sadar Penuh",
		"penilaian_nilai1":     "int",
		"penilaian_skala2":     "in:Perlu Bantuan Bernafas,Berusaha Bernafas,Batuk / Menangis",
		"penilaian_nilai2":     "int",
		"penilaian_skala3":     "in:Tidak Bergerak,Gerakan Tanpa Tujuan,Gerakan Beraturan",
		"penilaian_nilai3":     "int",
		"penilaian_totalnilai": "int",
		"keluar":               "string|max_len:200",
		"instruksi":            "string|max_len:200",
		"kd_dokter":            "required|string|max_len:20",
		"nip":                  "required|string|max_len:20",
	}
	return rules
}

// SkorStewardPascaAnestesiStore simpan monitoring steward pasca anestesi; kolom waktu kunci kosong = sekarang.
type SkorStewardPascaAnestesiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	SkorStewardPascaAnestesiData
}

func (r *SkorStewardPascaAnestesiStore) Authorize(ctx http.Context) error { return nil }

func (r *SkorStewardPascaAnestesiStore) Rules(ctx http.Context) map[string]any {
	rules := skorStewardPascaAnestesiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *SkorStewardPascaAnestesiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *SkorStewardPascaAnestesiStore) Payload() SkorStewardPascaAnestesiData {
	return r.SkorStewardPascaAnestesiData
}

func (r *SkorStewardPascaAnestesiStore) DetailValues() map[string][]string { return nil }

// SkorStewardPascaAnestesiUpdate ubah monitoring steward pasca anestesi (PUT); kunci lewat query string.
type SkorStewardPascaAnestesiUpdate struct {
	SkorStewardPascaAnestesiData
}

func (r *SkorStewardPascaAnestesiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkorStewardPascaAnestesiUpdate) Rules(ctx http.Context) map[string]any {
	return skorStewardPascaAnestesiRules()
}

func (r *SkorStewardPascaAnestesiUpdate) Payload() SkorStewardPascaAnestesiData {
	return r.SkorStewardPascaAnestesiData
}

func (r *SkorStewardPascaAnestesiUpdate) DetailValues() map[string][]string { return nil }
