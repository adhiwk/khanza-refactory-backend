package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningTbcData isian skrining TBC.
type SkriningTbcData struct {
	Tanggal                                string `form:"tanggal" json:"tanggal"`
	BeratBadan                             string `form:"berat_badan" json:"berat_badan"`
	TinggiBadan                            string `form:"tinggi_badan" json:"tinggi_badan"`
	Imt                                    string `form:"imt" json:"imt"`
	KasifikasiImt                          string `form:"kasifikasi_imt" json:"kasifikasi_imt"`
	LingkarPinggang                        string `form:"lingkar_pinggang" json:"lingkar_pinggang"`
	RisikoLingkarPinggang                  string `form:"risiko_lingkar_pinggang" json:"risiko_lingkar_pinggang"`
	RiwayatKontakTbc                       string `form:"riwayat_kontak_tbc" json:"riwayat_kontak_tbc"`
	JenisKontakTbc                         string `form:"jenis_kontak_tbc" json:"jenis_kontak_tbc"`
	FaktorResikoPernahTerdiagnosaTbc       string `form:"faktor_resiko_pernah_terdiagnosa_tbc" json:"faktor_resiko_pernah_terdiagnosa_tbc"`
	KeteranganPernahTerdiagnosa            string `form:"keterangan_pernah_terdiagnosa" json:"keterangan_pernah_terdiagnosa"`
	FaktorResikoPernahBerobatTbc           string `form:"faktor_resiko_pernah_berobat_tbc" json:"faktor_resiko_pernah_berobat_tbc"`
	FaktorResikoMalnutrisi                 string `form:"faktor_resiko_malnutrisi" json:"faktor_resiko_malnutrisi"`
	FaktorResikoMerokok                    string `form:"faktor_resiko_merokok" json:"faktor_resiko_merokok"`
	FaktorResikoRiwayatDm                  string `form:"faktor_resiko_riwayat_dm" json:"faktor_resiko_riwayat_dm"`
	FaktorResikoOdhiv                      string `form:"faktor_resiko_odhiv" json:"faktor_resiko_odhiv"`
	FaktorResikoLansia                     string `form:"faktor_resiko_lansia" json:"faktor_resiko_lansia"`
	FaktorResikoIbuHamil                   string `form:"faktor_resiko_ibu_hamil" json:"faktor_resiko_ibu_hamil"`
	FaktorResikoWbp                        string `form:"faktor_resiko_wbp" json:"faktor_resiko_wbp"`
	FaktorResikoTinggalDiwilayahPadatKumuh string `form:"faktor_resiko_tinggal_diwilayah_padat_kumuh" json:"faktor_resiko_tinggal_diwilayah_padat_kumuh"`
	AbnormalitasTbc                        string `form:"abnormalitas_tbc" json:"abnormalitas_tbc"`
	GejalaTbcBatuk                         string `form:"gejala_tbc_batuk" json:"gejala_tbc_batuk"`
	GejalaTbcBbTurun                       string `form:"gejala_tbc_bb_turun" json:"gejala_tbc_bb_turun"`
	GejalaTbcDemam                         string `form:"gejala_tbc_demam" json:"gejala_tbc_demam"`
	GejalaTbcBerkeringatMalamHari          string `form:"gejala_tbc_berkeringat_malam_hari" json:"gejala_tbc_berkeringat_malam_hari"`
	KeteranganGejalaPenyakitLain           string `form:"keterangan_gejala_penyakit_lain" json:"keterangan_gejala_penyakit_lain"`
	KesimpulanSkrining                     string `form:"kesimpulan_skrining" json:"kesimpulan_skrining"`
	KeteranganHasilSkrining                string `form:"keterangan_hasil_skrining" json:"keterangan_hasil_skrining"`
	Nip                                    string `form:"nip" json:"nip"`
}

func skriningTbcRules() map[string]any {
	rules := map[string]any{
		"tanggal":                                     "required|date",
		"berat_badan":                                 "string|max_len:6",
		"tinggi_badan":                                "string|max_len:8",
		"imt":                                         "string|max_len:6",
		"kasifikasi_imt":                              "in:Berat Badan Kurang,Berat Badan Normal,Kelebihan Berat Badan,Obesitas I,Obesitas II",
		"lingkar_pinggang":                            "string|max_len:6",
		"risiko_lingkar_pinggang":                     "in:Rendah,Cukup,Meningkat,Moderat,Berat,Sangat",
		"riwayat_kontak_tbc":                          "in:Ya,Tidak",
		"jenis_kontak_tbc":                            "in:Tidak,TBC Paru Bakteriologis,TBC Paru Klinis,TBC Paru Ekstraparu",
		"faktor_resiko_pernah_terdiagnosa_tbc":        "in:Ya,Tidak",
		"keterangan_pernah_terdiagnosa":               "string|max_len:40",
		"faktor_resiko_pernah_berobat_tbc":            "in:Ya,Tidak",
		"faktor_resiko_malnutrisi":                    "in:Ya,Tidak",
		"faktor_resiko_merokok":                       "in:Ya,Tidak",
		"faktor_resiko_riwayat_dm":                    "in:Ya,Tidak",
		"faktor_resiko_odhiv":                         "in:Ya,Tidak",
		"faktor_resiko_lansia":                        "in:Ya,Tidak",
		"faktor_resiko_ibu_hamil":                     "in:Ya,Tidak",
		"faktor_resiko_wbp":                           "in:Ya,Tidak",
		"faktor_resiko_tinggal_diwilayah_padat_kumuh": "in:Ya,Tidak",
		"abnormalitas_tbc":                            "in:Normal,Abnormalitas TBC,Abnormalitas Bukan TBC",
		"gejala_tbc_batuk":                            "in:Ya,Tidak",
		"gejala_tbc_bb_turun":                         "in:Ya,Tidak",
		"gejala_tbc_demam":                            "in:Ya,Tidak",
		"gejala_tbc_berkeringat_malam_hari":           "in:Ya,Tidak",
		"keterangan_gejala_penyakit_lain":             "string|max_len:30",
		"kesimpulan_skrining":                         "in:Bukan Terduga TBC,Kontak Erat,Terduga TBC",
		"keterangan_hasil_skrining":                   "string|max_len:50",
		"nip":                                         "required|string|max_len:20",
	}
	return rules
}

// SkriningTbcStore simpan skrining TBC; kolom waktu kunci kosong = sekarang.
type SkriningTbcStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningTbcData
}

func (r *SkriningTbcStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningTbcStore) Rules(ctx http.Context) map[string]any {
	rules := skriningTbcRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningTbcStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningTbcStore) Payload() SkriningTbcData { return r.SkriningTbcData }

func (r *SkriningTbcStore) DetailValues() map[string][]string { return nil }

// SkriningTbcUpdate ubah skrining TBC (PUT); kunci lewat query string.
type SkriningTbcUpdate struct {
	SkriningTbcData
}

func (r *SkriningTbcUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningTbcUpdate) Rules(ctx http.Context) map[string]any { return skriningTbcRules() }

func (r *SkriningTbcUpdate) Payload() SkriningTbcData { return r.SkriningTbcData }

func (r *SkriningTbcUpdate) DetailValues() map[string][]string { return nil }
