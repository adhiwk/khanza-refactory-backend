package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkorBromagePascaAnestesiData isian monitoring bromage pasca anestesi.
type SkorBromagePascaAnestesiData struct {
	PenilaianSkala1 string `form:"penilaian_skala1" json:"penilaian_skala1"`
	PenilaianNilai1 *int   `form:"penilaian_nilai1" json:"penilaian_nilai1"`
	Keluar          string `form:"keluar" json:"keluar"`
	Instruksi       string `form:"instruksi" json:"instruksi"`
	KdDokter        string `form:"kd_dokter" json:"kd_dokter"`
	Nip             string `form:"nip" json:"nip"`
}

func skorBromagePascaAnestesiRules() map[string]any {
	rules := map[string]any{
		"penilaian_skala1": "in:Gerakan Penuh Dari Tungkai,Tidak Mampu Extensi Tungkai,Tidak Mampu Flexi Lutut,Tidak Mampu Flexi Pergelangan Kaki",
		"penilaian_nilai1": "int",
		"keluar":           "string|max_len:200",
		"instruksi":        "string|max_len:200",
		"kd_dokter":        "required|string|max_len:20",
		"nip":              "required|string|max_len:20",
	}
	return rules
}

// SkorBromagePascaAnestesiStore simpan monitoring bromage pasca anestesi; kolom waktu kunci kosong = sekarang.
type SkorBromagePascaAnestesiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	SkorBromagePascaAnestesiData
}

func (r *SkorBromagePascaAnestesiStore) Authorize(ctx http.Context) error { return nil }

func (r *SkorBromagePascaAnestesiStore) Rules(ctx http.Context) map[string]any {
	rules := skorBromagePascaAnestesiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *SkorBromagePascaAnestesiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *SkorBromagePascaAnestesiStore) Payload() SkorBromagePascaAnestesiData {
	return r.SkorBromagePascaAnestesiData
}

func (r *SkorBromagePascaAnestesiStore) DetailValues() map[string][]string { return nil }

// SkorBromagePascaAnestesiUpdate ubah monitoring bromage pasca anestesi (PUT); kunci lewat query string.
type SkorBromagePascaAnestesiUpdate struct {
	SkorBromagePascaAnestesiData
}

func (r *SkorBromagePascaAnestesiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkorBromagePascaAnestesiUpdate) Rules(ctx http.Context) map[string]any {
	return skorBromagePascaAnestesiRules()
}

func (r *SkorBromagePascaAnestesiUpdate) Payload() SkorBromagePascaAnestesiData {
	return r.SkorBromagePascaAnestesiData
}

func (r *SkorBromagePascaAnestesiUpdate) DetailValues() map[string][]string { return nil }
