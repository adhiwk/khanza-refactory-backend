package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianPreInduksiData isian penilaian pre induksi.
type PenilaianPreInduksiData struct {
	KdDokter                     string `form:"kd_dokter" json:"kd_dokter"`
	Tensi                        string `form:"tensi" json:"tensi"`
	Nadi                         string `form:"nadi" json:"nadi"`
	Rr                           string `form:"rr" json:"rr"`
	Suhu                         string `form:"suhu" json:"suhu"`
	Ekg                          string `form:"ekg" json:"ekg"`
	LainLain                     string `form:"lain_lain" json:"lain_lain"`
	Asesmen                      string `form:"asesmen" json:"asesmen"`
	Perencanaan                  string `form:"perencanaan" json:"perencanaan"`
	InfusPerifier                string `form:"infus_perifier" json:"infus_perifier"`
	Cvc                          string `form:"cvc" json:"cvc"`
	Posisi                       string `form:"posisi" json:"posisi"`
	Premedikasi                  string `form:"premedikasi" json:"premedikasi"`
	PremedikasiKeterangan        string `form:"premedikasi_keterangan" json:"premedikasi_keterangan"`
	Induksi                      string `form:"induksi" json:"induksi"`
	InduksiKeterangan            string `form:"induksi_keterangan" json:"induksi_keterangan"`
	FaceMaskNo                   string `form:"face_mask_no" json:"face_mask_no"`
	NasopharingNo                string `form:"nasopharing_no" json:"nasopharing_no"`
	EttNo                        string `form:"ett_no" json:"ett_no"`
	EttJenis                     string `form:"ett_jenis" json:"ett_jenis"`
	EttViksasi                   string `form:"ett_viksasi" json:"ett_viksasi"`
	LmaNo                        string `form:"lma_no" json:"lma_no"`
	LmaJenis                     string `form:"lma_jenis" json:"lma_jenis"`
	Tracheostomi                 string `form:"tracheostomi" json:"tracheostomi"`
	BronchoscopiFiberoptik       string `form:"bronchoscopi_fiberoptik" json:"bronchoscopi_fiberoptik"`
	Glidescopi                   string `form:"glidescopi" json:"glidescopi"`
	LainLainTatalaksana          string `form:"lain_lain_tatalaksana" json:"lain_lain_tatalaksana"`
	IntubasiSesudahTidur         string `form:"intubasi_sesudah_tidur" json:"intubasi_sesudah_tidur"`
	IntubasiOral                 string `form:"intubasi_oral" json:"intubasi_oral"`
	IntubasiTracheostomi         string `form:"intubasi_tracheostomi" json:"intubasi_tracheostomi"`
	IntubasiKeterangan           string `form:"intubasi_keterangan" json:"intubasi_keterangan"`
	SulitVentilasi               string `form:"sulit_ventilasi" json:"sulit_ventilasi"`
	SulitIntubasi                string `form:"sulit_intubasi" json:"sulit_intubasi"`
	Ventilasi                    string `form:"ventilasi" json:"ventilasi"`
	TeknikRegionalJenis          string `form:"teknik_regional_jenis" json:"teknik_regional_jenis"`
	TeknikRegionalLokasi         string `form:"teknik_regional_lokasi" json:"teknik_regional_lokasi"`
	TeknikRegionalJenisJarum     string `form:"teknik_regional_jenis_jarum" json:"teknik_regional_jenis_jarum"`
	TeknikRegionalKateter        string `form:"teknik_regional_kateter" json:"teknik_regional_kateter"`
	TeknikRegionalKateterViksasi string `form:"teknik_regional_kateter_viksasi" json:"teknik_regional_kateter_viksasi"`
	TeknikRegionalObatObatan     string `form:"teknik_regional_obat_obatan" json:"teknik_regional_obat_obatan"`
	TeknikRegionalKomplikasi     string `form:"teknik_regional_komplikasi" json:"teknik_regional_komplikasi"`
	TeknikRegionalHasil          string `form:"teknik_regional_hasil" json:"teknik_regional_hasil"`
}

func penilaianPreInduksiRules() map[string]any {
	rules := map[string]any{
		"kd_dokter":                       "required|string|max_len:20",
		"tensi":                           "string|max_len:8",
		"nadi":                            "string|max_len:5",
		"rr":                              "string|max_len:5",
		"suhu":                            "string|max_len:5",
		"ekg":                             "string|max_len:50",
		"lain_lain":                       "string|max_len:50",
		"asesmen":                         "in:Sesuai Asesmen Pre Sedasi/Anestesi,Tidak Sesuai Asesmen Pre Sedasi/Anestesi",
		"perencanaan":                     "string|max_len:300",
		"infus_perifier":                  "string|max_len:300",
		"cvc":                             "string|max_len:70",
		"posisi":                          "in:Supine,Lithotomi,Lateral,Prone,Perlindungan Mata,Kanan,Kiri,Lain-lain",
		"premedikasi":                     "in:Oral,IM,IV",
		"premedikasi_keterangan":          "string|max_len:50",
		"induksi":                         "in:Intravena,Inhalasi",
		"induksi_keterangan":              "string|max_len:70",
		"face_mask_no":                    "string|max_len:20",
		"nasopharing_no":                  "string|max_len:20",
		"ett_no":                          "string|max_len:20",
		"ett_jenis":                       "string|max_len:20",
		"ett_viksasi":                     "string|max_len:25",
		"lma_no":                          "string|max_len:20",
		"lma_jenis":                       "string|max_len:20",
		"tracheostomi":                    "string|max_len:60",
		"bronchoscopi_fiberoptik":         "string|max_len:60",
		"glidescopi":                      "string|max_len:60",
		"lain_lain_tatalaksana":           "string|max_len:100",
		"intubasi_sesudah_tidur":          "in:Ya,Tidak",
		"intubasi_oral":                   "in:Ya,Tidak",
		"intubasi_tracheostomi":           "in:Ya,Tidak",
		"intubasi_keterangan":             "string|max_len:200",
		"sulit_ventilasi":                 "string|max_len:100",
		"sulit_intubasi":                  "string|max_len:100",
		"ventilasi":                       "string|max_len:100",
		"teknik_regional_jenis":           "string|max_len:100",
		"teknik_regional_lokasi":          "string|max_len:40",
		"teknik_regional_jenis_jarum":     "string|max_len:30",
		"teknik_regional_kateter":         "in:Ya,Tidak",
		"teknik_regional_kateter_viksasi": "string|max_len:40",
		"teknik_regional_obat_obatan":     "string|max_len:400",
		"teknik_regional_komplikasi":      "string|max_len:200",
		"teknik_regional_hasil":           "string|max_len:100",
	}
	return rules
}

// PenilaianPreInduksiStore simpan penilaian pre induksi; kolom waktu kunci kosong = sekarang.
type PenilaianPreInduksiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PenilaianPreInduksiData
}

func (r *PenilaianPreInduksiStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPreInduksiStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianPreInduksiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PenilaianPreInduksiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PenilaianPreInduksiStore) Payload() PenilaianPreInduksiData {
	return r.PenilaianPreInduksiData
}

func (r *PenilaianPreInduksiStore) DetailValues() map[string][]string { return nil }

// PenilaianPreInduksiUpdate ubah penilaian pre induksi (PUT); kunci lewat query string.
type PenilaianPreInduksiUpdate struct {
	PenilaianPreInduksiData
}

func (r *PenilaianPreInduksiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPreInduksiUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianPreInduksiRules()
}

func (r *PenilaianPreInduksiUpdate) Payload() PenilaianPreInduksiData {
	return r.PenilaianPreInduksiData
}

func (r *PenilaianPreInduksiUpdate) DetailValues() map[string][]string { return nil }
