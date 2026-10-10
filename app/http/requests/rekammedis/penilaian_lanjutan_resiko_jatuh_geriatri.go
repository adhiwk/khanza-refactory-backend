package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianLanjutanResikoJatuhGeriatriData isian penilaian lanjutan risiko jatuh geriatri.
type PenilaianLanjutanResikoJatuhGeriatriData struct {
	PenilaianJatuhSkala1     string `form:"penilaian_jatuh_skala1" json:"penilaian_jatuh_skala1"`
	PenilaianJatuhNilai1     *int   `form:"penilaian_jatuh_nilai1" json:"penilaian_jatuh_nilai1"`
	PenilaianJatuhSkala2     string `form:"penilaian_jatuh_skala2" json:"penilaian_jatuh_skala2"`
	PenilaianJatuhNilai2     *int   `form:"penilaian_jatuh_nilai2" json:"penilaian_jatuh_nilai2"`
	PenilaianJatuhSkala3     string `form:"penilaian_jatuh_skala3" json:"penilaian_jatuh_skala3"`
	PenilaianJatuhNilai3     *int   `form:"penilaian_jatuh_nilai3" json:"penilaian_jatuh_nilai3"`
	PenilaianJatuhSkala4     string `form:"penilaian_jatuh_skala4" json:"penilaian_jatuh_skala4"`
	PenilaianJatuhNilai4     *int   `form:"penilaian_jatuh_nilai4" json:"penilaian_jatuh_nilai4"`
	PenilaianJatuhSkala5     string `form:"penilaian_jatuh_skala5" json:"penilaian_jatuh_skala5"`
	PenilaianJatuhNilai5     *int   `form:"penilaian_jatuh_nilai5" json:"penilaian_jatuh_nilai5"`
	PenilaianJatuhSkala6     string `form:"penilaian_jatuh_skala6" json:"penilaian_jatuh_skala6"`
	PenilaianJatuhNilai6     *int   `form:"penilaian_jatuh_nilai6" json:"penilaian_jatuh_nilai6"`
	PenilaianJatuhSkala7     string `form:"penilaian_jatuh_skala7" json:"penilaian_jatuh_skala7"`
	PenilaianJatuhNilai7     *int   `form:"penilaian_jatuh_nilai7" json:"penilaian_jatuh_nilai7"`
	PenilaianJatuhSkala8     string `form:"penilaian_jatuh_skala8" json:"penilaian_jatuh_skala8"`
	PenilaianJatuhNilai8     *int   `form:"penilaian_jatuh_nilai8" json:"penilaian_jatuh_nilai8"`
	PenilaianJatuhSkala9     string `form:"penilaian_jatuh_skala9" json:"penilaian_jatuh_skala9"`
	PenilaianJatuhNilai9     *int   `form:"penilaian_jatuh_nilai9" json:"penilaian_jatuh_nilai9"`
	PenilaianJatuhSkala10    string `form:"penilaian_jatuh_skala10" json:"penilaian_jatuh_skala10"`
	PenilaianJatuhNilai10    *int   `form:"penilaian_jatuh_nilai10" json:"penilaian_jatuh_nilai10"`
	PenilaianJatuhSkala11    string `form:"penilaian_jatuh_skala11" json:"penilaian_jatuh_skala11"`
	PenilaianJatuhNilai11    *int   `form:"penilaian_jatuh_nilai11" json:"penilaian_jatuh_nilai11"`
	PenilaianJatuhTotalnilai *int   `form:"penilaian_jatuh_totalnilai" json:"penilaian_jatuh_totalnilai"`
	HasilSkrining            string `form:"hasil_skrining" json:"hasil_skrining"`
	Saran                    string `form:"saran" json:"saran"`
	Nip                      string `form:"nip" json:"nip"`
}

func penilaianLanjutanResikoJatuhGeriatriRules() map[string]any {
	rules := map[string]any{
		"penilaian_jatuh_skala1":     "in:Tidak,Ya",
		"penilaian_jatuh_nilai1":     "int",
		"penilaian_jatuh_skala2":     "in:Tidak,Ya",
		"penilaian_jatuh_nilai2":     "int",
		"penilaian_jatuh_skala3":     "in:Tidak,Ya",
		"penilaian_jatuh_nilai3":     "int",
		"penilaian_jatuh_skala4":     "in:Tidak,Ya",
		"penilaian_jatuh_nilai4":     "int",
		"penilaian_jatuh_skala5":     "in:Tidak,Ya",
		"penilaian_jatuh_nilai5":     "int",
		"penilaian_jatuh_skala6":     "in:Tidak,Ya",
		"penilaian_jatuh_nilai6":     "int",
		"penilaian_jatuh_skala7":     "in:Tidak,Ya",
		"penilaian_jatuh_nilai7":     "int",
		"penilaian_jatuh_skala8":     "in:Tidak,Ya",
		"penilaian_jatuh_nilai8":     "int",
		"penilaian_jatuh_skala9":     "in:Tidak,Ya",
		"penilaian_jatuh_nilai9":     "int",
		"penilaian_jatuh_skala10":    "in:Tidak,Ya",
		"penilaian_jatuh_nilai10":    "int",
		"penilaian_jatuh_skala11":    "in:Tidak,Ya",
		"penilaian_jatuh_nilai11":    "int",
		"penilaian_jatuh_totalnilai": "int",
		"hasil_skrining":             "string|max_len:200",
		"saran":                      "string|max_len:200",
		"nip":                        "required|string|max_len:20",
	}
	return rules
}

// PenilaianLanjutanResikoJatuhGeriatriStore simpan penilaian lanjutan risiko jatuh geriatri; kolom waktu kunci kosong = sekarang.
type PenilaianLanjutanResikoJatuhGeriatriStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PenilaianLanjutanResikoJatuhGeriatriData
}

func (r *PenilaianLanjutanResikoJatuhGeriatriStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianLanjutanResikoJatuhGeriatriStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianLanjutanResikoJatuhGeriatriRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PenilaianLanjutanResikoJatuhGeriatriStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PenilaianLanjutanResikoJatuhGeriatriStore) Payload() PenilaianLanjutanResikoJatuhGeriatriData {
	return r.PenilaianLanjutanResikoJatuhGeriatriData
}

func (r *PenilaianLanjutanResikoJatuhGeriatriStore) DetailValues() map[string][]string { return nil }

// PenilaianLanjutanResikoJatuhGeriatriUpdate ubah penilaian lanjutan risiko jatuh geriatri (PUT); kunci lewat query string.
type PenilaianLanjutanResikoJatuhGeriatriUpdate struct {
	PenilaianLanjutanResikoJatuhGeriatriData
}

func (r *PenilaianLanjutanResikoJatuhGeriatriUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianLanjutanResikoJatuhGeriatriUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianLanjutanResikoJatuhGeriatriRules()
}

func (r *PenilaianLanjutanResikoJatuhGeriatriUpdate) Payload() PenilaianLanjutanResikoJatuhGeriatriData {
	return r.PenilaianLanjutanResikoJatuhGeriatriData
}

func (r *PenilaianLanjutanResikoJatuhGeriatriUpdate) DetailValues() map[string][]string { return nil }
