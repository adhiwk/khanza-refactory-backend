package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningRisikoKankerServiksData isian skrining risiko kanker serviks.
type SkriningRisikoKankerServiksData struct {
	Tanggal                 string `form:"tanggal" json:"tanggal"`
	RiwayatPenyakitKeluarga string `form:"riwayat_penyakit_keluarga" json:"riwayat_penyakit_keluarga"`
	RiwayatPenyakitSendiri  string `form:"riwayat_penyakit_sendiri" json:"riwayat_penyakit_sendiri"`
	RisikoMerokok           string `form:"risiko_merokok" json:"risiko_merokok"`
	RisikoKurangFisik       string `form:"risiko_kurang_fisik" json:"risiko_kurang_fisik"`
	RisikoGulaBerlebihan    string `form:"risiko_gula_berlebihan" json:"risiko_gula_berlebihan"`
	RisikoGaramBerlebihan   string `form:"risiko_garam_berlebihan" json:"risiko_garam_berlebihan"`
	RisikoLemakBerlebihan   string `form:"risiko_lemak_berlebihan" json:"risiko_lemak_berlebihan"`
	RisikoKurangBuahSayur   string `form:"risiko_kurang_buah_sayur" json:"risiko_kurang_buah_sayur"`
	RisikoAlkohol           string `form:"risiko_alkohol" json:"risiko_alkohol"`
	HasilIva                string `form:"hasil_iva" json:"hasil_iva"`
	HasilSkrining           string `form:"hasil_skrining" json:"hasil_skrining"`
	Keterangan              string `form:"keterangan" json:"keterangan"`
	Nip                     string `form:"nip" json:"nip"`
}

func skriningRisikoKankerServiksRules() map[string]any {
	rules := map[string]any{
		"tanggal":                   "required|date",
		"riwayat_penyakit_keluarga": "in:Kanker,Benjolan Abnormal Pada Payudara,-",
		"riwayat_penyakit_sendiri":  "in:Kanker,Benjolan Abnormal Pada Payudara,-",
		"risiko_merokok":            "in:Ya,Tidak",
		"risiko_kurang_fisik":       "in:Ya,Tidak",
		"risiko_gula_berlebihan":    "in:Ya,Tidak",
		"risiko_garam_berlebihan":   "in:Ya,Tidak",
		"risiko_lemak_berlebihan":   "in:Ya,Tidak",
		"risiko_kurang_buah_sayur":  "in:Ya,Tidak",
		"risiko_alkohol":            "in:Ya,Tidak",
		"hasil_iva":                 "in:Positif,Negatif,Curiga Kanker",
		"hasil_skrining":            "string|max_len:50",
		"keterangan":                "string|max_len:100",
		"nip":                       "required|string|max_len:20",
	}
	return rules
}

// SkriningRisikoKankerServiksStore simpan skrining risiko kanker serviks; kolom waktu kunci kosong = sekarang.
type SkriningRisikoKankerServiksStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningRisikoKankerServiksData
}

func (r *SkriningRisikoKankerServiksStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningRisikoKankerServiksStore) Rules(ctx http.Context) map[string]any {
	rules := skriningRisikoKankerServiksRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningRisikoKankerServiksStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *SkriningRisikoKankerServiksStore) Payload() SkriningRisikoKankerServiksData {
	return r.SkriningRisikoKankerServiksData
}

func (r *SkriningRisikoKankerServiksStore) DetailValues() map[string][]string { return nil }

// SkriningRisikoKankerServiksUpdate ubah skrining risiko kanker serviks (PUT); kunci lewat query string.
type SkriningRisikoKankerServiksUpdate struct {
	SkriningRisikoKankerServiksData
}

func (r *SkriningRisikoKankerServiksUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningRisikoKankerServiksUpdate) Rules(ctx http.Context) map[string]any {
	return skriningRisikoKankerServiksRules()
}

func (r *SkriningRisikoKankerServiksUpdate) Payload() SkriningRisikoKankerServiksData {
	return r.SkriningRisikoKankerServiksData
}

func (r *SkriningRisikoKankerServiksUpdate) DetailValues() map[string][]string { return nil }
