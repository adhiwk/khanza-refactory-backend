package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningAdiksiNikotinData isian skrining adiksi nikotin.
type SkriningAdiksiNikotinData struct {
	Tanggal                 string `form:"tanggal" json:"tanggal"`
	RokokDihisap            string `form:"rokok_dihisap" json:"rokok_dihisap"`
	NilaiRokokDihisap       *int   `form:"nilai_rokok_dihisap" json:"nilai_rokok_dihisap"`
	MenyalakanRokok         string `form:"menyalakan_rokok" json:"menyalakan_rokok"`
	NilaiMenyalakanRokok    *int   `form:"nilai_menyalakan_rokok" json:"nilai_menyalakan_rokok"`
	TidakRela               string `form:"tidak_rela" json:"tidak_rela"`
	NilaiTidakRela          *int   `form:"nilai_tidak_rela" json:"nilai_tidak_rela"`
	JamPertama              string `form:"jam_pertama" json:"jam_pertama"`
	NilaiJamPertama         *int   `form:"nilai_jam_pertama" json:"nilai_jam_pertama"`
	RasaIngin               string `form:"rasa_ingin" json:"rasa_ingin"`
	NilaiRasaIngin          *int   `form:"nilai_rasa_ingin" json:"nilai_rasa_ingin"`
	SakitBerat              string `form:"sakit_berat" json:"sakit_berat"`
	NilaiSakitBerat         *int   `form:"nilai_sakit_berat" json:"nilai_sakit_berat"`
	NilaiTotal              *int   `form:"nilai_total" json:"nilai_total"`
	KeteranganHasilSkrining string `form:"keterangan_hasil_skrining" json:"keterangan_hasil_skrining"`
	SkalaMotivasi           string `form:"skala_motivasi" json:"skala_motivasi"`
	Nip                     string `form:"nip" json:"nip"`
}

func skriningAdiksiNikotinRules() map[string]any {
	rules := map[string]any{
		"tanggal":                   "required|date",
		"rokok_dihisap":             "in:1-10,11-20,21-30,>=31",
		"nilai_rokok_dihisap":       "int",
		"menyalakan_rokok":          "in:Dalam 5 Menit,6 Hingga 30 Menit,31 Hingga 60 Menit,Setelah 60 Menit",
		"nilai_menyalakan_rokok":    "int",
		"tidak_rela":                "in:Lainnya,Rokok Pertama Pada Pagi Hari",
		"nilai_tidak_rela":          "int",
		"jam_pertama":               "in:Ya,Tidak",
		"nilai_jam_pertama":         "int",
		"rasa_ingin":                "in:Ya,Tidak",
		"nilai_rasa_ingin":          "int",
		"sakit_berat":               "in:Ya,Tidak",
		"nilai_sakit_berat":         "int",
		"nilai_total":               "int",
		"keterangan_hasil_skrining": "in:Ketergantungan Berat,Ketergantungan Sedang,Ketergantungan Rendah",
		"skala_motivasi":            "string",
		"nip":                       "required|string|max_len:20",
	}
	return rules
}

// SkriningAdiksiNikotinStore simpan skrining adiksi nikotin; kolom waktu kunci kosong = sekarang.
type SkriningAdiksiNikotinStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningAdiksiNikotinData
}

func (r *SkriningAdiksiNikotinStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningAdiksiNikotinStore) Rules(ctx http.Context) map[string]any {
	rules := skriningAdiksiNikotinRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningAdiksiNikotinStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningAdiksiNikotinStore) Payload() SkriningAdiksiNikotinData {
	return r.SkriningAdiksiNikotinData
}

func (r *SkriningAdiksiNikotinStore) DetailValues() map[string][]string { return nil }

// SkriningAdiksiNikotinUpdate ubah skrining adiksi nikotin (PUT); kunci lewat query string.
type SkriningAdiksiNikotinUpdate struct {
	SkriningAdiksiNikotinData
}

func (r *SkriningAdiksiNikotinUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningAdiksiNikotinUpdate) Rules(ctx http.Context) map[string]any {
	return skriningAdiksiNikotinRules()
}

func (r *SkriningAdiksiNikotinUpdate) Payload() SkriningAdiksiNikotinData {
	return r.SkriningAdiksiNikotinData
}

func (r *SkriningAdiksiNikotinUpdate) DetailValues() map[string][]string { return nil }
