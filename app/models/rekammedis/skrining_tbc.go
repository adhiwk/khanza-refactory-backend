package rekammedis

import "time"

// SkriningTbc tabel `skrining_tbc` (skrining TBC, RMSkriningTBC).
type SkriningTbc struct {
	NoRawat                                string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                                *time.Time `gorm:"column:tanggal" json:"tanggal"`
	BeratBadan                             *string    `gorm:"column:berat_badan" json:"berat_badan"`
	TinggiBadan                            *string    `gorm:"column:tinggi_badan" json:"tinggi_badan"`
	Imt                                    *string    `gorm:"column:imt" json:"imt"`
	KasifikasiImt                          *string    `gorm:"column:kasifikasi_imt" json:"kasifikasi_imt"`
	LingkarPinggang                        *string    `gorm:"column:lingkar_pinggang" json:"lingkar_pinggang"`
	RisikoLingkarPinggang                  *string    `gorm:"column:risiko_lingkar_pinggang" json:"risiko_lingkar_pinggang"`
	RiwayatKontakTbc                       *string    `gorm:"column:riwayat_kontak_tbc" json:"riwayat_kontak_tbc"`
	JenisKontakTbc                         *string    `gorm:"column:jenis_kontak_tbc" json:"jenis_kontak_tbc"`
	FaktorResikoPernahTerdiagnosaTbc       *string    `gorm:"column:faktor_resiko_pernah_terdiagnosa_tbc" json:"faktor_resiko_pernah_terdiagnosa_tbc"`
	KeteranganPernahTerdiagnosa            *string    `gorm:"column:keterangan_pernah_terdiagnosa" json:"keterangan_pernah_terdiagnosa"`
	FaktorResikoPernahBerobatTbc           *string    `gorm:"column:faktor_resiko_pernah_berobat_tbc" json:"faktor_resiko_pernah_berobat_tbc"`
	FaktorResikoMalnutrisi                 *string    `gorm:"column:faktor_resiko_malnutrisi" json:"faktor_resiko_malnutrisi"`
	FaktorResikoMerokok                    *string    `gorm:"column:faktor_resiko_merokok" json:"faktor_resiko_merokok"`
	FaktorResikoRiwayatDm                  *string    `gorm:"column:faktor_resiko_riwayat_dm" json:"faktor_resiko_riwayat_dm"`
	FaktorResikoOdhiv                      *string    `gorm:"column:faktor_resiko_odhiv" json:"faktor_resiko_odhiv"`
	FaktorResikoLansia                     *string    `gorm:"column:faktor_resiko_lansia" json:"faktor_resiko_lansia"`
	FaktorResikoIbuHamil                   *string    `gorm:"column:faktor_resiko_ibu_hamil" json:"faktor_resiko_ibu_hamil"`
	FaktorResikoWbp                        *string    `gorm:"column:faktor_resiko_wbp" json:"faktor_resiko_wbp"`
	FaktorResikoTinggalDiwilayahPadatKumuh *string    `gorm:"column:faktor_resiko_tinggal_diwilayah_padat_kumuh" json:"faktor_resiko_tinggal_diwilayah_padat_kumuh"`
	AbnormalitasTbc                        *string    `gorm:"column:abnormalitas_tbc" json:"abnormalitas_tbc"`
	GejalaTbcBatuk                         *string    `gorm:"column:gejala_tbc_batuk" json:"gejala_tbc_batuk"`
	GejalaTbcBbTurun                       *string    `gorm:"column:gejala_tbc_bb_turun" json:"gejala_tbc_bb_turun"`
	GejalaTbcDemam                         *string    `gorm:"column:gejala_tbc_demam" json:"gejala_tbc_demam"`
	GejalaTbcBerkeringatMalamHari          *string    `gorm:"column:gejala_tbc_berkeringat_malam_hari" json:"gejala_tbc_berkeringat_malam_hari"`
	KeteranganGejalaPenyakitLain           *string    `gorm:"column:keterangan_gejala_penyakit_lain" json:"keterangan_gejala_penyakit_lain"`
	KesimpulanSkrining                     *string    `gorm:"column:kesimpulan_skrining" json:"kesimpulan_skrining"`
	KeteranganHasilSkrining                *string    `gorm:"column:keterangan_hasil_skrining" json:"keterangan_hasil_skrining"`
	Nip                                    string     `gorm:"column:nip" json:"nip"`
}

func (SkriningTbc) TableName() string {
	return "skrining_tbc"
}
