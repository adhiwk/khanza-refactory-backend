package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianRisikoJatuhNeonatusData isian penilaian risiko jatuh neonatus.
type PenilaianRisikoJatuhNeonatusData struct {
	Intervensi1 string `form:"intervensi1" json:"intervensi1"`
	Intervensi2 string `form:"intervensi2" json:"intervensi2"`
	Intervensi3 string `form:"intervensi3" json:"intervensi3"`
	Intervensi4 string `form:"intervensi4" json:"intervensi4"`
	Intervensi5 string `form:"intervensi5" json:"intervensi5"`
	Intervensi6 string `form:"intervensi6" json:"intervensi6"`
	Intervensi7 string `form:"intervensi7" json:"intervensi7"`
	Intervensi8 string `form:"intervensi8" json:"intervensi8"`
	Intervensi9 string `form:"intervensi9" json:"intervensi9"`
	Edukasi1    string `form:"edukasi1" json:"edukasi1"`
	Edukasi2    string `form:"edukasi2" json:"edukasi2"`
	Edukasi3    string `form:"edukasi3" json:"edukasi3"`
	Edukasi4    string `form:"edukasi4" json:"edukasi4"`
	Edukasi5    string `form:"edukasi5" json:"edukasi5"`
	Sasaran1    string `form:"sasaran1" json:"sasaran1"`
	Sasaran2    string `form:"sasaran2" json:"sasaran2"`
	Sasaran3    string `form:"sasaran3" json:"sasaran3"`
	Sasaran4    string `form:"sasaran4" json:"sasaran4"`
	Evaluasi1   string `form:"evaluasi1" json:"evaluasi1"`
	Evaluasi2   string `form:"evaluasi2" json:"evaluasi2"`
	Evaluasi3   string `form:"evaluasi3" json:"evaluasi3"`
	Nip         string `form:"nip" json:"nip"`
}

func penilaianRisikoJatuhNeonatusRules() map[string]any {
	rules := map[string]any{
		"intervensi1": "in:Tidak,Ya",
		"intervensi2": "in:Tidak,Ya",
		"intervensi3": "in:Tidak,Ya",
		"intervensi4": "in:Tidak,Ya",
		"intervensi5": "in:Tidak,Ya",
		"intervensi6": "in:Tidak,Ya",
		"intervensi7": "in:Tidak,Ya",
		"intervensi8": "in:Tidak,Ya",
		"intervensi9": "in:Tidak,Ya",
		"edukasi1":    "in:Tidak,Ya",
		"edukasi2":    "in:Tidak,Ya",
		"edukasi3":    "in:Tidak,Ya",
		"edukasi4":    "in:Tidak,Ya",
		"edukasi5":    "in:Tidak,Ya",
		"sasaran1":    "in:Tidak,Ya",
		"sasaran2":    "in:Tidak,Ya",
		"sasaran3":    "in:Tidak,Ya",
		"sasaran4":    "in:Tidak,Ya",
		"evaluasi1":   "in:Tidak,Ya",
		"evaluasi2":   "in:Tidak,Ya",
		"evaluasi3":   "in:Tidak,Ya",
		"nip":         "required|string|max_len:20",
	}
	return rules
}

// PenilaianRisikoJatuhNeonatusStore simpan penilaian risiko jatuh neonatus; kolom waktu kunci kosong = sekarang.
type PenilaianRisikoJatuhNeonatusStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PenilaianRisikoJatuhNeonatusData
}

func (r *PenilaianRisikoJatuhNeonatusStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianRisikoJatuhNeonatusStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianRisikoJatuhNeonatusRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PenilaianRisikoJatuhNeonatusStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PenilaianRisikoJatuhNeonatusStore) Payload() PenilaianRisikoJatuhNeonatusData {
	return r.PenilaianRisikoJatuhNeonatusData
}

func (r *PenilaianRisikoJatuhNeonatusStore) DetailValues() map[string][]string { return nil }

// PenilaianRisikoJatuhNeonatusUpdate ubah penilaian risiko jatuh neonatus (PUT); kunci lewat query string.
type PenilaianRisikoJatuhNeonatusUpdate struct {
	PenilaianRisikoJatuhNeonatusData
}

func (r *PenilaianRisikoJatuhNeonatusUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianRisikoJatuhNeonatusUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianRisikoJatuhNeonatusRules()
}

func (r *PenilaianRisikoJatuhNeonatusUpdate) Payload() PenilaianRisikoJatuhNeonatusData {
	return r.PenilaianRisikoJatuhNeonatusData
}

func (r *PenilaianRisikoJatuhNeonatusUpdate) DetailValues() map[string][]string { return nil }
