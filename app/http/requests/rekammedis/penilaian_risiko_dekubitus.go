package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianRisikoDekubitusData isian penilaian risiko dekubitus.
type PenilaianRisikoDekubitusData struct {
	KondisiFisik       string `form:"kondisi_fisik" json:"kondisi_fisik"`
	KondisiFisikNilai  *int   `form:"kondisi_fisik_nilai" json:"kondisi_fisik_nilai"`
	StatusMental       string `form:"status_mental" json:"status_mental"`
	StatusMentalNilai  *int   `form:"status_mental_nilai" json:"status_mental_nilai"`
	Aktifitas          string `form:"aktifitas" json:"aktifitas"`
	AktifitasNilai     *int   `form:"aktifitas_nilai" json:"aktifitas_nilai"`
	Mobilitas          string `form:"mobilitas" json:"mobilitas"`
	MobilitasNilai     *int   `form:"mobilitas_nilai" json:"mobilitas_nilai"`
	Inkontinensia      string `form:"inkontinensia" json:"inkontinensia"`
	InkontinensiaNilai *int   `form:"inkontinensia_nilai" json:"inkontinensia_nilai"`
	Totalnilai         *int   `form:"totalnilai" json:"totalnilai"`
	Kategorinilai      string `form:"kategorinilai" json:"kategorinilai"`
	Nip                string `form:"nip" json:"nip"`
}

func penilaianRisikoDekubitusRules() map[string]any {
	rules := map[string]any{
		"kondisi_fisik":       "in:Baik,Sedang,Buruk,Sangat Buruk",
		"kondisi_fisik_nilai": "int",
		"status_mental":       "in:Sadar,Apatis,Bingung,Stupor",
		"status_mental_nilai": "int",
		"aktifitas":           "in:Jalan Sendiri,Jalan Dengan Bantuan,Kursi Roda,Di Tempat Tidur",
		"aktifitas_nilai":     "int",
		"mobilitas":           "in:Bebas Bergerak,Agar Terbatas,Sangat Terbatas,Tidak Mampu Bergerak",
		"mobilitas_nilai":     "int",
		"inkontinensia":       "in:Kontinen,Kadang-kadang Inkontinensia Urine,Selalu Inkontenesia Urine,Inkontinensia Alvi Dan Urine",
		"inkontinensia_nilai": "int",
		"totalnilai":          "int",
		"kategorinilai":       "in:Risiko Rendah,Risiko Sedang,Risiko Tinggi",
		"nip":                 "required|string|max_len:20",
	}
	return rules
}

// PenilaianRisikoDekubitusStore simpan penilaian risiko dekubitus; kolom waktu kunci kosong = sekarang.
type PenilaianRisikoDekubitusStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PenilaianRisikoDekubitusData
}

func (r *PenilaianRisikoDekubitusStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianRisikoDekubitusStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianRisikoDekubitusRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PenilaianRisikoDekubitusStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PenilaianRisikoDekubitusStore) Payload() PenilaianRisikoDekubitusData {
	return r.PenilaianRisikoDekubitusData
}

func (r *PenilaianRisikoDekubitusStore) DetailValues() map[string][]string { return nil }

// PenilaianRisikoDekubitusUpdate ubah penilaian risiko dekubitus (PUT); kunci lewat query string.
type PenilaianRisikoDekubitusUpdate struct {
	PenilaianRisikoDekubitusData
}

func (r *PenilaianRisikoDekubitusUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianRisikoDekubitusUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianRisikoDekubitusRules()
}

func (r *PenilaianRisikoDekubitusUpdate) Payload() PenilaianRisikoDekubitusData {
	return r.PenilaianRisikoDekubitusData
}

func (r *PenilaianRisikoDekubitusUpdate) DetailValues() map[string][]string { return nil }
