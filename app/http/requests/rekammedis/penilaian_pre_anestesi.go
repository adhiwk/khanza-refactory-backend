package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianPreAnestesiData isian penilaian pre anastesi.
type PenilaianPreAnestesiData struct {
	KdDokter                     string `form:"kd_dokter" json:"kd_dokter"`
	TanggalOperasi               string `form:"tanggal_operasi" json:"tanggal_operasi"`
	Diagnosa                     string `form:"diagnosa" json:"diagnosa"`
	RencanaTindakan              string `form:"rencana_tindakan" json:"rencana_tindakan"`
	Tb                           string `form:"tb" json:"tb"`
	Bb                           string `form:"bb" json:"bb"`
	Td                           string `form:"td" json:"td"`
	Io2                          string `form:"io2" json:"io2"`
	Nadi                         string `form:"nadi" json:"nadi"`
	Pernapasan                   string `form:"pernapasan" json:"pernapasan"`
	Suhu                         string `form:"suhu" json:"suhu"`
	FisikCardiovasculer          string `form:"fisik_cardiovasculer" json:"fisik_cardiovasculer"`
	FisikParu                    string `form:"fisik_paru" json:"fisik_paru"`
	FisikAbdomen                 string `form:"fisik_abdomen" json:"fisik_abdomen"`
	FisikExtrimitas              string `form:"fisik_extrimitas" json:"fisik_extrimitas"`
	FisikEndokrin                string `form:"fisik_endokrin" json:"fisik_endokrin"`
	FisikGinjal                  string `form:"fisik_ginjal" json:"fisik_ginjal"`
	FisikObatobatan              string `form:"fisik_obatobatan" json:"fisik_obatobatan"`
	FisikLaborat                 string `form:"fisik_laborat" json:"fisik_laborat"`
	FisikPenunjang               string `form:"fisik_penunjang" json:"fisik_penunjang"`
	RiwayatPenyakitAlergiobat    string `form:"riwayat_penyakit_alergiobat" json:"riwayat_penyakit_alergiobat"`
	RiwayatPenyakitAlergilainnya string `form:"riwayat_penyakit_alergilainnya" json:"riwayat_penyakit_alergilainnya"`
	RiwayatPenyakitTerapi        string `form:"riwayat_penyakit_terapi" json:"riwayat_penyakit_terapi"`
	RiwayatKebiasaanMerokok      string `form:"riwayat_kebiasaan_merokok" json:"riwayat_kebiasaan_merokok"`
	RiwayatKebiasaanKetMerokok   string `form:"riwayat_kebiasaan_ket_merokok" json:"riwayat_kebiasaan_ket_merokok"`
	RiwayatKebiasaanAlkohol      string `form:"riwayat_kebiasaan_alkohol" json:"riwayat_kebiasaan_alkohol"`
	RiwayatKebiasaanKetAlkohol   string `form:"riwayat_kebiasaan_ket_alkohol" json:"riwayat_kebiasaan_ket_alkohol"`
	RiwayatKebiasaanObat         string `form:"riwayat_kebiasaan_obat" json:"riwayat_kebiasaan_obat"`
	RiwayatKebiasaanKetObat      string `form:"riwayat_kebiasaan_ket_obat" json:"riwayat_kebiasaan_ket_obat"`
	RiwayatMedisCardiovasculer   string `form:"riwayat_medis_cardiovasculer" json:"riwayat_medis_cardiovasculer"`
	RiwayatMedisRespiratory      string `form:"riwayat_medis_respiratory" json:"riwayat_medis_respiratory"`
	RiwayatMedisEndocrine        string `form:"riwayat_medis_endocrine" json:"riwayat_medis_endocrine"`
	RiwayatMedisLainnya          string `form:"riwayat_medis_lainnya" json:"riwayat_medis_lainnya"`
	Asa                          string `form:"asa" json:"asa"`
	Puasa                        string `form:"puasa" json:"puasa"`
	RencanaAnestesi              string `form:"rencana_anestesi" json:"rencana_anestesi"`
	RencanaPerawatan             string `form:"rencana_perawatan" json:"rencana_perawatan"`
	CatatanKhusus                string `form:"catatan_khusus" json:"catatan_khusus"`
}

func penilaianPreAnestesiRules() map[string]any {
	rules := map[string]any{
		"kd_dokter":                      "required|string|max_len:20",
		"tanggal_operasi":                "date",
		"diagnosa":                       "string|max_len:100",
		"rencana_tindakan":               "string|max_len:100",
		"tb":                             "string|max_len:5",
		"bb":                             "string|max_len:5",
		"td":                             "string|max_len:8",
		"io2":                            "string|max_len:5",
		"nadi":                           "string|max_len:5",
		"pernapasan":                     "string|max_len:5",
		"suhu":                           "string|max_len:5",
		"fisik_cardiovasculer":           "string|max_len:100",
		"fisik_paru":                     "string|max_len:100",
		"fisik_abdomen":                  "string|max_len:100",
		"fisik_extrimitas":               "string|max_len:100",
		"fisik_endokrin":                 "string|max_len:100",
		"fisik_ginjal":                   "string|max_len:100",
		"fisik_obatobatan":               "string|max_len:100",
		"fisik_laborat":                  "string|max_len:100",
		"fisik_penunjang":                "string|max_len:100",
		"riwayat_penyakit_alergiobat":    "string|max_len:50",
		"riwayat_penyakit_alergilainnya": "string|max_len:50",
		"riwayat_penyakit_terapi":        "string|max_len:100",
		"riwayat_kebiasaan_merokok":      "required|in:Tidak,Ya",
		"riwayat_kebiasaan_ket_merokok":  "string|max_len:5",
		"riwayat_kebiasaan_alkohol":      "required|in:Tidak,Ya",
		"riwayat_kebiasaan_ket_alkohol":  "string|max_len:5",
		"riwayat_kebiasaan_obat":         "required|in:-,Obat Obatan,Vitamin,Jamu Jamuan",
		"riwayat_kebiasaan_ket_obat":     "string|max_len:100",
		"riwayat_medis_cardiovasculer":   "string|max_len:100",
		"riwayat_medis_respiratory":      "string|max_len:100",
		"riwayat_medis_endocrine":        "string|max_len:100",
		"riwayat_medis_lainnya":          "string|max_len:100",
		"asa":                            "in:1,2,3,4,5,E,2E,3E,4E,5E",
		"puasa":                          "date",
		"rencana_anestesi":               "in:GA,RA Spinal,RA Epidural,RA Combined,Blok Syaraf",
		"rencana_perawatan":              "string|max_len:40",
		"catatan_khusus":                 "string|max_len:100",
	}
	return rules
}

// PenilaianPreAnestesiStore simpan penilaian pre anastesi; kolom waktu kunci kosong = sekarang.
type PenilaianPreAnestesiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PenilaianPreAnestesiData
}

func (r *PenilaianPreAnestesiStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPreAnestesiStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianPreAnestesiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PenilaianPreAnestesiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PenilaianPreAnestesiStore) Payload() PenilaianPreAnestesiData {
	return r.PenilaianPreAnestesiData
}

func (r *PenilaianPreAnestesiStore) DetailValues() map[string][]string { return nil }

// PenilaianPreAnestesiUpdate ubah penilaian pre anastesi (PUT); kunci lewat query string.
type PenilaianPreAnestesiUpdate struct {
	PenilaianPreAnestesiData
}

func (r *PenilaianPreAnestesiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPreAnestesiUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianPreAnestesiRules()
}

func (r *PenilaianPreAnestesiUpdate) Payload() PenilaianPreAnestesiData {
	return r.PenilaianPreAnestesiData
}

func (r *PenilaianPreAnestesiUpdate) DetailValues() map[string][]string { return nil }
