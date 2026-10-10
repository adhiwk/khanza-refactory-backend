package rekammedis

import "time"

// PenilaianLevelKecemasanRanapAnak tabel `penilaian_level_kecemasan_ranap_anak` (penilaian level kecemasan ranap anak, RMPenilaianLevelKecemasanRanapAnak).
type PenilaianLevelKecemasanRanapAnak struct {
	NoRawat                        string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                        *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Cemas                          *int       `gorm:"column:cemas" json:"cemas"`
	FirasatBuruk                   *int       `gorm:"column:firasat_buruk" json:"firasat_buruk"`
	TakutPikiranSendiri            *int       `gorm:"column:takut_pikiran_sendiri" json:"takut_pikiran_sendiri"`
	MudahTersinggung               *int       `gorm:"column:mudah_tersinggung" json:"mudah_tersinggung"`
	MerasaTegang                   *int       `gorm:"column:merasa_tegang" json:"merasa_tegang"`
	Lesu                           *int       `gorm:"column:lesu" json:"lesu"`
	TakBisaIstirahatTenang         *int       `gorm:"column:tak_bisa_istirahat_tenang" json:"tak_bisa_istirahat_tenang"`
	MudahTerkejut                  *int       `gorm:"column:mudah_terkejut" json:"mudah_terkejut"`
	MudahMenangis                  *int       `gorm:"column:mudah_menangis" json:"mudah_menangis"`
	Gemetar                        *int       `gorm:"column:gemetar" json:"gemetar"`
	Gelisah                        *int       `gorm:"column:gelisah" json:"gelisah"`
	TakutPadaGelap                 *int       `gorm:"column:takut_pada_gelap" json:"takut_pada_gelap"`
	TakutPadaOrangasing            *int       `gorm:"column:takut_pada_orangasing" json:"takut_pada_orangasing"`
	TakutPadaKerumunanBanyakOrang  *int       `gorm:"column:takut_pada_kerumunan_banyak_orang" json:"takut_pada_kerumunan_banyak_orang"`
	TakutPadaBinatangBesar         *int       `gorm:"column:takut_pada_binatang_besar" json:"takut_pada_binatang_besar"`
	TakutPadaKeramaianLaluLintas   *int       `gorm:"column:takut_pada_keramaian_lalu_lintas" json:"takut_pada_keramaian_lalu_lintas"`
	TakutDitinggalSendiri          *int       `gorm:"column:takut_ditinggal_sendiri" json:"takut_ditinggal_sendiri"`
	SulitTidur                     *int       `gorm:"column:sulit_tidur" json:"sulit_tidur"`
	TerbangunMalamHari             *int       `gorm:"column:terbangun_malam_hari" json:"terbangun_malam_hari"`
	TidurTidakNyeyak               *int       `gorm:"column:tidur_tidak_nyeyak" json:"tidur_tidak_nyeyak"`
	MimpiBuruk                     *int       `gorm:"column:mimpi_buruk" json:"mimpi_buruk"`
	BangunDenganLesu               *int       `gorm:"column:bangun_dengan_lesu" json:"bangun_dengan_lesu"`
	BanyakMengalamiMimpi           *int       `gorm:"column:banyak_mengalami_mimpi" json:"banyak_mengalami_mimpi"`
	MimpiMenakutkan                *int       `gorm:"column:mimpi_menakutkan" json:"mimpi_menakutkan"`
	SulitKonsentrasi               *int       `gorm:"column:sulit_konsentrasi" json:"sulit_konsentrasi"`
	DayaIngatBuruk                 *int       `gorm:"column:daya_ingat_buruk" json:"daya_ingat_buruk"`
	HilangnyaMinat                 *int       `gorm:"column:hilangnya_minat" json:"hilangnya_minat"`
	BerkurangnyaKesenanganPadaHobi *int       `gorm:"column:berkurangnya_kesenangan_pada_hobi" json:"berkurangnya_kesenangan_pada_hobi"`
	Sedih                          *int       `gorm:"column:sedih" json:"sedih"`
	BangunDiniHari                 *int       `gorm:"column:bangun_dini_hari" json:"bangun_dini_hari"`
	PerasaanBerubah                *int       `gorm:"column:perasaan_berubah" json:"perasaan_berubah"`
	SakitNyeriDiOtot               *int       `gorm:"column:sakit_nyeri_di_otot" json:"sakit_nyeri_di_otot"`
	Kaku                           *int       `gorm:"column:kaku" json:"kaku"`
	KedutanOtot                    *int       `gorm:"column:kedutan_otot" json:"kedutan_otot"`
	GigiGemerutuk                  *int       `gorm:"column:gigi_gemerutuk" json:"gigi_gemerutuk"`
	SuaraTidakStabil               *int       `gorm:"column:suara_tidak_stabil" json:"suara_tidak_stabil"`
	Tinnitus                       *int       `gorm:"column:tinnitus" json:"tinnitus"`
	PenglihatanKabur               *int       `gorm:"column:penglihatan_kabur" json:"penglihatan_kabur"`
	MukaMerahGejalaSomatic         *int       `gorm:"column:muka_merah_gejala_somatic" json:"muka_merah_gejala_somatic"`
	MerasaLemah                    *int       `gorm:"column:merasa_lemah" json:"merasa_lemah"`
	PerasaanDitusuk                *int       `gorm:"column:perasaan_ditusuk" json:"perasaan_ditusuk"`
	Takhikardia                    *int       `gorm:"column:takhikardia" json:"takhikardia"`
	Berdebar                       *int       `gorm:"column:berdebar" json:"berdebar"`
	NyeriDiDada                    *int       `gorm:"column:nyeri_di_dada" json:"nyeri_di_dada"`
	DenyutNadiMengeras             *int       `gorm:"column:denyut_nadi_mengeras" json:"denyut_nadi_mengeras"`
	PerasaanLesu                   *int       `gorm:"column:perasaan_lesu" json:"perasaan_lesu"`
	DetakJantungMenghilang         *int       `gorm:"column:detak_jantung_menghilang" json:"detak_jantung_menghilang"`
	MerasaTertekan                 *int       `gorm:"column:merasa_tertekan" json:"merasa_tertekan"`
	PerasaanTercekik               *int       `gorm:"column:perasaan_tercekik" json:"perasaan_tercekik"`
	SeringMenarikNapas             *int       `gorm:"column:sering_menarik_napas" json:"sering_menarik_napas"`
	NapasPendek                    *int       `gorm:"column:napas_pendek" json:"napas_pendek"`
	BuluBerdiri                    *int       `gorm:"column:bulu_berdiri" json:"bulu_berdiri"`
	SulitMenelan                   *int       `gorm:"column:sulit_menelan" json:"sulit_menelan"`
	PerutMelilit                   *int       `gorm:"column:perut_melilit" json:"perut_melilit"`
	GanguanPencernaan              *int       `gorm:"column:ganguan_pencernaan" json:"ganguan_pencernaan"`
	RasaKembung                    *int       `gorm:"column:rasa_kembung" json:"rasa_kembung"`
	NyeriMakan                     *int       `gorm:"column:nyeri_makan" json:"nyeri_makan"`
	TerbakarPerut                  *int       `gorm:"column:terbakar_perut" json:"terbakar_perut"`
	SukarBab                       *int       `gorm:"column:sukar_bab" json:"sukar_bab"`
	Muntah                         *int       `gorm:"column:muntah" json:"muntah"`
	BabLembek                      *int       `gorm:"column:bab_lembek" json:"bab_lembek"`
	KehilanganBb                   *int       `gorm:"column:kehilangan_bb" json:"kehilangan_bb"`
	Mual                           *int       `gorm:"column:mual" json:"mual"`
	SeringBak                      *int       `gorm:"column:sering_bak" json:"sering_bak"`
	TidakBisaMenahanKencing        *int       `gorm:"column:tidak_bisa_menahan_kencing" json:"tidak_bisa_menahan_kencing"`
	MenjadiDingin                  *int       `gorm:"column:menjadi_dingin" json:"menjadi_dingin"`
	Manorrhagia                    *int       `gorm:"column:manorrhagia" json:"manorrhagia"`
	Amenorrhoea                    *int       `gorm:"column:amenorrhoea" json:"amenorrhoea"`
	EjakulasiPraecocks             *int       `gorm:"column:ejakulasi_praecocks" json:"ejakulasi_praecocks"`
	EreksiHilang                   *int       `gorm:"column:ereksi_hilang" json:"ereksi_hilang"`
	Impotensi                      *int       `gorm:"column:impotensi" json:"impotensi"`
	MulutKering                    *int       `gorm:"column:mulut_kering" json:"mulut_kering"`
	MukaMerahGejalaOtonom          *int       `gorm:"column:muka_merah_gejala_otonom" json:"muka_merah_gejala_otonom"`
	MudahBerkeringat               *int       `gorm:"column:mudah_berkeringat" json:"mudah_berkeringat"`
	BuluBerdiriGejalaOtonom        *int       `gorm:"column:bulu_berdiri_gejala_otonom" json:"bulu_berdiri_gejala_otonom"`
	SakitKepala                    *int       `gorm:"column:sakit_kepala" json:"sakit_kepala"`
	GelisahWawancara               *int       `gorm:"column:gelisah_wawancara" json:"gelisah_wawancara"`
	NapasPendekWawancara           *int       `gorm:"column:napas_pendek_wawancara" json:"napas_pendek_wawancara"`
	JariGemetar                    *int       `gorm:"column:jari_gemetar" json:"jari_gemetar"`
	KerutKening                    *int       `gorm:"column:kerut_kening" json:"kerut_kening"`
	MukaTegang                     *int       `gorm:"column:muka_tegang" json:"muka_tegang"`
	TonusMeningkat                 *int       `gorm:"column:tonus_meningkat" json:"tonus_meningkat"`
	TidakTenang                    *int       `gorm:"column:tidak_tenang" json:"tidak_tenang"`
	MukaMerahWawancara             *int       `gorm:"column:muka_merah_wawancara" json:"muka_merah_wawancara"`
	TotalSkor                      *int       `gorm:"column:total_skor" json:"total_skor"`
	KeteranganSkor                 *string    `gorm:"column:keterangan_skor" json:"keterangan_skor"`
	Nip                            string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianLevelKecemasanRanapAnak) TableName() string {
	return "penilaian_level_kecemasan_ranap_anak"
}
