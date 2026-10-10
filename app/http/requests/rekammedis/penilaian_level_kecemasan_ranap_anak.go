package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianLevelKecemasanRanapAnakData isian penilaian level kecemasan ranap anak.
type PenilaianLevelKecemasanRanapAnakData struct {
	Cemas                          *int   `form:"cemas" json:"cemas"`
	FirasatBuruk                   *int   `form:"firasat_buruk" json:"firasat_buruk"`
	TakutPikiranSendiri            *int   `form:"takut_pikiran_sendiri" json:"takut_pikiran_sendiri"`
	MudahTersinggung               *int   `form:"mudah_tersinggung" json:"mudah_tersinggung"`
	MerasaTegang                   *int   `form:"merasa_tegang" json:"merasa_tegang"`
	Lesu                           *int   `form:"lesu" json:"lesu"`
	TakBisaIstirahatTenang         *int   `form:"tak_bisa_istirahat_tenang" json:"tak_bisa_istirahat_tenang"`
	MudahTerkejut                  *int   `form:"mudah_terkejut" json:"mudah_terkejut"`
	MudahMenangis                  *int   `form:"mudah_menangis" json:"mudah_menangis"`
	Gemetar                        *int   `form:"gemetar" json:"gemetar"`
	Gelisah                        *int   `form:"gelisah" json:"gelisah"`
	TakutPadaGelap                 *int   `form:"takut_pada_gelap" json:"takut_pada_gelap"`
	TakutPadaOrangasing            *int   `form:"takut_pada_orangasing" json:"takut_pada_orangasing"`
	TakutPadaKerumunanBanyakOrang  *int   `form:"takut_pada_kerumunan_banyak_orang" json:"takut_pada_kerumunan_banyak_orang"`
	TakutPadaBinatangBesar         *int   `form:"takut_pada_binatang_besar" json:"takut_pada_binatang_besar"`
	TakutPadaKeramaianLaluLintas   *int   `form:"takut_pada_keramaian_lalu_lintas" json:"takut_pada_keramaian_lalu_lintas"`
	TakutDitinggalSendiri          *int   `form:"takut_ditinggal_sendiri" json:"takut_ditinggal_sendiri"`
	SulitTidur                     *int   `form:"sulit_tidur" json:"sulit_tidur"`
	TerbangunMalamHari             *int   `form:"terbangun_malam_hari" json:"terbangun_malam_hari"`
	TidurTidakNyeyak               *int   `form:"tidur_tidak_nyeyak" json:"tidur_tidak_nyeyak"`
	MimpiBuruk                     *int   `form:"mimpi_buruk" json:"mimpi_buruk"`
	BangunDenganLesu               *int   `form:"bangun_dengan_lesu" json:"bangun_dengan_lesu"`
	BanyakMengalamiMimpi           *int   `form:"banyak_mengalami_mimpi" json:"banyak_mengalami_mimpi"`
	MimpiMenakutkan                *int   `form:"mimpi_menakutkan" json:"mimpi_menakutkan"`
	SulitKonsentrasi               *int   `form:"sulit_konsentrasi" json:"sulit_konsentrasi"`
	DayaIngatBuruk                 *int   `form:"daya_ingat_buruk" json:"daya_ingat_buruk"`
	HilangnyaMinat                 *int   `form:"hilangnya_minat" json:"hilangnya_minat"`
	BerkurangnyaKesenanganPadaHobi *int   `form:"berkurangnya_kesenangan_pada_hobi" json:"berkurangnya_kesenangan_pada_hobi"`
	Sedih                          *int   `form:"sedih" json:"sedih"`
	BangunDiniHari                 *int   `form:"bangun_dini_hari" json:"bangun_dini_hari"`
	PerasaanBerubah                *int   `form:"perasaan_berubah" json:"perasaan_berubah"`
	SakitNyeriDiOtot               *int   `form:"sakit_nyeri_di_otot" json:"sakit_nyeri_di_otot"`
	Kaku                           *int   `form:"kaku" json:"kaku"`
	KedutanOtot                    *int   `form:"kedutan_otot" json:"kedutan_otot"`
	GigiGemerutuk                  *int   `form:"gigi_gemerutuk" json:"gigi_gemerutuk"`
	SuaraTidakStabil               *int   `form:"suara_tidak_stabil" json:"suara_tidak_stabil"`
	Tinnitus                       *int   `form:"tinnitus" json:"tinnitus"`
	PenglihatanKabur               *int   `form:"penglihatan_kabur" json:"penglihatan_kabur"`
	MukaMerahGejalaSomatic         *int   `form:"muka_merah_gejala_somatic" json:"muka_merah_gejala_somatic"`
	MerasaLemah                    *int   `form:"merasa_lemah" json:"merasa_lemah"`
	PerasaanDitusuk                *int   `form:"perasaan_ditusuk" json:"perasaan_ditusuk"`
	Takhikardia                    *int   `form:"takhikardia" json:"takhikardia"`
	Berdebar                       *int   `form:"berdebar" json:"berdebar"`
	NyeriDiDada                    *int   `form:"nyeri_di_dada" json:"nyeri_di_dada"`
	DenyutNadiMengeras             *int   `form:"denyut_nadi_mengeras" json:"denyut_nadi_mengeras"`
	PerasaanLesu                   *int   `form:"perasaan_lesu" json:"perasaan_lesu"`
	DetakJantungMenghilang         *int   `form:"detak_jantung_menghilang" json:"detak_jantung_menghilang"`
	MerasaTertekan                 *int   `form:"merasa_tertekan" json:"merasa_tertekan"`
	PerasaanTercekik               *int   `form:"perasaan_tercekik" json:"perasaan_tercekik"`
	SeringMenarikNapas             *int   `form:"sering_menarik_napas" json:"sering_menarik_napas"`
	NapasPendek                    *int   `form:"napas_pendek" json:"napas_pendek"`
	BuluBerdiri                    *int   `form:"bulu_berdiri" json:"bulu_berdiri"`
	SulitMenelan                   *int   `form:"sulit_menelan" json:"sulit_menelan"`
	PerutMelilit                   *int   `form:"perut_melilit" json:"perut_melilit"`
	GanguanPencernaan              *int   `form:"ganguan_pencernaan" json:"ganguan_pencernaan"`
	RasaKembung                    *int   `form:"rasa_kembung" json:"rasa_kembung"`
	NyeriMakan                     *int   `form:"nyeri_makan" json:"nyeri_makan"`
	TerbakarPerut                  *int   `form:"terbakar_perut" json:"terbakar_perut"`
	SukarBab                       *int   `form:"sukar_bab" json:"sukar_bab"`
	Muntah                         *int   `form:"muntah" json:"muntah"`
	BabLembek                      *int   `form:"bab_lembek" json:"bab_lembek"`
	KehilanganBb                   *int   `form:"kehilangan_bb" json:"kehilangan_bb"`
	Mual                           *int   `form:"mual" json:"mual"`
	SeringBak                      *int   `form:"sering_bak" json:"sering_bak"`
	TidakBisaMenahanKencing        *int   `form:"tidak_bisa_menahan_kencing" json:"tidak_bisa_menahan_kencing"`
	MenjadiDingin                  *int   `form:"menjadi_dingin" json:"menjadi_dingin"`
	Manorrhagia                    *int   `form:"manorrhagia" json:"manorrhagia"`
	Amenorrhoea                    *int   `form:"amenorrhoea" json:"amenorrhoea"`
	EjakulasiPraecocks             *int   `form:"ejakulasi_praecocks" json:"ejakulasi_praecocks"`
	EreksiHilang                   *int   `form:"ereksi_hilang" json:"ereksi_hilang"`
	Impotensi                      *int   `form:"impotensi" json:"impotensi"`
	MulutKering                    *int   `form:"mulut_kering" json:"mulut_kering"`
	MukaMerahGejalaOtonom          *int   `form:"muka_merah_gejala_otonom" json:"muka_merah_gejala_otonom"`
	MudahBerkeringat               *int   `form:"mudah_berkeringat" json:"mudah_berkeringat"`
	BuluBerdiriGejalaOtonom        *int   `form:"bulu_berdiri_gejala_otonom" json:"bulu_berdiri_gejala_otonom"`
	SakitKepala                    *int   `form:"sakit_kepala" json:"sakit_kepala"`
	GelisahWawancara               *int   `form:"gelisah_wawancara" json:"gelisah_wawancara"`
	NapasPendekWawancara           *int   `form:"napas_pendek_wawancara" json:"napas_pendek_wawancara"`
	JariGemetar                    *int   `form:"jari_gemetar" json:"jari_gemetar"`
	KerutKening                    *int   `form:"kerut_kening" json:"kerut_kening"`
	MukaTegang                     *int   `form:"muka_tegang" json:"muka_tegang"`
	TonusMeningkat                 *int   `form:"tonus_meningkat" json:"tonus_meningkat"`
	TidakTenang                    *int   `form:"tidak_tenang" json:"tidak_tenang"`
	MukaMerahWawancara             *int   `form:"muka_merah_wawancara" json:"muka_merah_wawancara"`
	TotalSkor                      *int   `form:"total_skor" json:"total_skor"`
	KeteranganSkor                 string `form:"keterangan_skor" json:"keterangan_skor"`
	Nip                            string `form:"nip" json:"nip"`
}

func penilaianLevelKecemasanRanapAnakRules() map[string]any {
	rules := map[string]any{
		"cemas":                             "int",
		"firasat_buruk":                     "int",
		"takut_pikiran_sendiri":             "int",
		"mudah_tersinggung":                 "int",
		"merasa_tegang":                     "int",
		"lesu":                              "int",
		"tak_bisa_istirahat_tenang":         "int",
		"mudah_terkejut":                    "int",
		"mudah_menangis":                    "int",
		"gemetar":                           "int",
		"gelisah":                           "int",
		"takut_pada_gelap":                  "int",
		"takut_pada_orangasing":             "int",
		"takut_pada_kerumunan_banyak_orang": "int",
		"takut_pada_binatang_besar":         "int",
		"takut_pada_keramaian_lalu_lintas":  "int",
		"takut_ditinggal_sendiri":           "int",
		"sulit_tidur":                       "int",
		"terbangun_malam_hari":              "int",
		"tidur_tidak_nyeyak":                "int",
		"mimpi_buruk":                       "int",
		"bangun_dengan_lesu":                "int",
		"banyak_mengalami_mimpi":            "int",
		"mimpi_menakutkan":                  "int",
		"sulit_konsentrasi":                 "int",
		"daya_ingat_buruk":                  "int",
		"hilangnya_minat":                   "int",
		"berkurangnya_kesenangan_pada_hobi": "int",
		"sedih":                             "int",
		"bangun_dini_hari":                  "int",
		"perasaan_berubah":                  "int",
		"sakit_nyeri_di_otot":               "int",
		"kaku":                              "int",
		"kedutan_otot":                      "int",
		"gigi_gemerutuk":                    "int",
		"suara_tidak_stabil":                "int",
		"tinnitus":                          "int",
		"penglihatan_kabur":                 "int",
		"muka_merah_gejala_somatic":         "int",
		"merasa_lemah":                      "int",
		"perasaan_ditusuk":                  "int",
		"takhikardia":                       "int",
		"berdebar":                          "int",
		"nyeri_di_dada":                     "int",
		"denyut_nadi_mengeras":              "int",
		"perasaan_lesu":                     "int",
		"detak_jantung_menghilang":          "int",
		"merasa_tertekan":                   "int",
		"perasaan_tercekik":                 "int",
		"sering_menarik_napas":              "int",
		"napas_pendek":                      "int",
		"bulu_berdiri":                      "int",
		"sulit_menelan":                     "int",
		"perut_melilit":                     "int",
		"ganguan_pencernaan":                "int",
		"rasa_kembung":                      "int",
		"nyeri_makan":                       "int",
		"terbakar_perut":                    "int",
		"sukar_bab":                         "int",
		"muntah":                            "int",
		"bab_lembek":                        "int",
		"kehilangan_bb":                     "int",
		"mual":                              "int",
		"sering_bak":                        "int",
		"tidak_bisa_menahan_kencing":        "int",
		"menjadi_dingin":                    "int",
		"manorrhagia":                       "int",
		"amenorrhoea":                       "int",
		"ejakulasi_praecocks":               "int",
		"ereksi_hilang":                     "int",
		"impotensi":                         "int",
		"mulut_kering":                      "int",
		"muka_merah_gejala_otonom":          "int",
		"mudah_berkeringat":                 "int",
		"bulu_berdiri_gejala_otonom":        "int",
		"sakit_kepala":                      "int",
		"gelisah_wawancara":                 "int",
		"napas_pendek_wawancara":            "int",
		"jari_gemetar":                      "int",
		"kerut_kening":                      "int",
		"muka_tegang":                       "int",
		"tonus_meningkat":                   "int",
		"tidak_tenang":                      "int",
		"muka_merah_wawancara":              "int",
		"total_skor":                        "int",
		"keterangan_skor":                   "in:Tidak Mengalami Kecemasan,Kecemasan Ringan,Kecemasan Sedang,Kecemasan Berat,Kecemasan Sangat Berat",
		"nip":                               "required|string|max_len:20",
	}
	return rules
}

// PenilaianLevelKecemasanRanapAnakStore simpan penilaian level kecemasan ranap anak; kolom waktu kunci kosong = sekarang.
type PenilaianLevelKecemasanRanapAnakStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PenilaianLevelKecemasanRanapAnakData
}

func (r *PenilaianLevelKecemasanRanapAnakStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianLevelKecemasanRanapAnakStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianLevelKecemasanRanapAnakRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PenilaianLevelKecemasanRanapAnakStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PenilaianLevelKecemasanRanapAnakStore) Payload() PenilaianLevelKecemasanRanapAnakData {
	return r.PenilaianLevelKecemasanRanapAnakData
}

func (r *PenilaianLevelKecemasanRanapAnakStore) DetailValues() map[string][]string { return nil }

// PenilaianLevelKecemasanRanapAnakUpdate ubah penilaian level kecemasan ranap anak (PUT); kunci lewat query string.
type PenilaianLevelKecemasanRanapAnakUpdate struct {
	PenilaianLevelKecemasanRanapAnakData
}

func (r *PenilaianLevelKecemasanRanapAnakUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianLevelKecemasanRanapAnakUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianLevelKecemasanRanapAnakRules()
}

func (r *PenilaianLevelKecemasanRanapAnakUpdate) Payload() PenilaianLevelKecemasanRanapAnakData {
	return r.PenilaianLevelKecemasanRanapAnakData
}

func (r *PenilaianLevelKecemasanRanapAnakUpdate) DetailValues() map[string][]string { return nil }
