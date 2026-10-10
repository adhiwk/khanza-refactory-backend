package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningObesitasData isian skrining obesitas.
type SkriningObesitasData struct {
	Tanggal                            string `form:"tanggal" json:"tanggal"`
	KebiasaanMakanManis                string `form:"kebiasaan_makan_manis" json:"kebiasaan_makan_manis"`
	AktifitasFisikSetiapHari           string `form:"aktifitas_fisik_setiap_hari" json:"aktifitas_fisik_setiap_hari"`
	IstirahatCukup                     string `form:"istirahat_cukup" json:"istirahat_cukup"`
	RisikoMerokok                      string `form:"risiko_merokok" json:"risiko_merokok"`
	RiwayatMinumAlkoholMerokokKeluarga string `form:"riwayat_minum_alkohol_merokok_keluarga" json:"riwayat_minum_alkohol_merokok_keluarga"`
	RiwayatPenggunaanObatSteroid       string `form:"riwayat_penggunaan_obat_steroid" json:"riwayat_penggunaan_obat_steroid"`
	BeratBadan                         string `form:"berat_badan" json:"berat_badan"`
	TinggiBadan                        string `form:"tinggi_badan" json:"tinggi_badan"`
	Imt                                string `form:"imt" json:"imt"`
	KasifikasiImt                      string `form:"kasifikasi_imt" json:"kasifikasi_imt"`
	LingkarPinggang                    string `form:"lingkar_pinggang" json:"lingkar_pinggang"`
	RisikoLingkarPinggang              string `form:"risiko_lingkar_pinggang" json:"risiko_lingkar_pinggang"`
	StatusObesitas                     string `form:"status_obesitas" json:"status_obesitas"`
	Keterangan                         string `form:"keterangan" json:"keterangan"`
	Nip                                string `form:"nip" json:"nip"`
}

func skriningObesitasRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                "required|date",
		"kebiasaan_makan_manis":                  "in:Ya,Tidak",
		"aktifitas_fisik_setiap_hari":            "in:Ya,Tidak",
		"istirahat_cukup":                        "in:Ya,Tidak",
		"risiko_merokok":                         "in:Ya,Tidak",
		"riwayat_minum_alkohol_merokok_keluarga": "in:Ya,Tidak",
		"riwayat_penggunaan_obat_steroid":        "in:Ya,Tidak",
		"berat_badan":                            "string|max_len:6",
		"tinggi_badan":                           "string|max_len:8",
		"imt":                                    "string|max_len:6",
		"kasifikasi_imt":                         "in:Berat Badan Kurang,Berat Badan Normal,Kelebihan Berat Badan,Obesitas I,Obesitas II",
		"lingkar_pinggang":                       "string|max_len:6",
		"risiko_lingkar_pinggang":                "in:Rendah,Cukup,Meningkat,Moderat,Berat,Sangat",
		"status_obesitas":                        "in:Normal,Berisiko",
		"keterangan":                             "string|max_len:40",
		"nip":                                    "required|string|max_len:20",
	}
	return rules
}

// SkriningObesitasStore simpan skrining obesitas; kolom waktu kunci kosong = sekarang.
type SkriningObesitasStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningObesitasData
}

func (r *SkriningObesitasStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningObesitasStore) Rules(ctx http.Context) map[string]any {
	rules := skriningObesitasRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningObesitasStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningObesitasStore) Payload() SkriningObesitasData { return r.SkriningObesitasData }

func (r *SkriningObesitasStore) DetailValues() map[string][]string { return nil }

// SkriningObesitasUpdate ubah skrining obesitas (PUT); kunci lewat query string.
type SkriningObesitasUpdate struct {
	SkriningObesitasData
}

func (r *SkriningObesitasUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningObesitasUpdate) Rules(ctx http.Context) map[string]any {
	return skriningObesitasRules()
}

func (r *SkriningObesitasUpdate) Payload() SkriningObesitasData { return r.SkriningObesitasData }

func (r *SkriningObesitasUpdate) DetailValues() map[string][]string { return nil }
