package rekammedis

import "time"

// CatatanAnestesiSedasi tabel `catatan_anestesi_sedasi` (catatan anastesi sedasi, RMCatatanAnastesiSedasi).
type CatatanAnestesiSedasi struct {
	NoRawat                                 string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                                 *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	KdDokterBedah                           string     `gorm:"column:kd_dokter_bedah" json:"kd_dokter_bedah"`
	KdDokterAnestesi                        string     `gorm:"column:kd_dokter_anestesi" json:"kd_dokter_anestesi"`
	DiagnosaPreBedah                        string     `gorm:"column:diagnosa_pre_bedah" json:"diagnosa_pre_bedah"`
	TindakanJenisPembedahan                 string     `gorm:"column:tindakan_jenis_pembedahan" json:"tindakan_jenis_pembedahan"`
	DiagnosaPascaBedah                      string     `gorm:"column:diagnosa_pasca_bedah" json:"diagnosa_pasca_bedah"`
	PreInduksiJam                           *string    `gorm:"column:pre_induksi_jam" json:"pre_induksi_jam"`
	PreInduksiKesadaran                     *string    `gorm:"column:pre_induksi_kesadaran" json:"pre_induksi_kesadaran"`
	PreInduksiTd                            *string    `gorm:"column:pre_induksi_td" json:"pre_induksi_td"`
	PreInduksiNadi                          *string    `gorm:"column:pre_induksi_nadi" json:"pre_induksi_nadi"`
	PreInduksiRr                            *string    `gorm:"column:pre_induksi_rr" json:"pre_induksi_rr"`
	PreInduksiSuhu                          *string    `gorm:"column:pre_induksi_suhu" json:"pre_induksi_suhu"`
	PreInduksiO2                            *string    `gorm:"column:pre_induksi_o2" json:"pre_induksi_o2"`
	PreInduksiTb                            *string    `gorm:"column:pre_induksi_tb" json:"pre_induksi_tb"`
	PreInduksiBb                            *string    `gorm:"column:pre_induksi_bb" json:"pre_induksi_bb"`
	PreInduksiRhesus                        *string    `gorm:"column:pre_induksi_rhesus" json:"pre_induksi_rhesus"`
	PreInduksiHb                            *string    `gorm:"column:pre_induksi_hb" json:"pre_induksi_hb"`
	PreInduksiHt                            *string    `gorm:"column:pre_induksi_ht" json:"pre_induksi_ht"`
	PreInduksiLeko                          *string    `gorm:"column:pre_induksi_leko" json:"pre_induksi_leko"`
	PreInduksiTrombo                        *string    `gorm:"column:pre_induksi_trombo" json:"pre_induksi_trombo"`
	PreInduksiBtct                          *string    `gorm:"column:pre_induksi_btct" json:"pre_induksi_btct"`
	PreInduksiGds                           *string    `gorm:"column:pre_induksi_gds" json:"pre_induksi_gds"`
	PreInduksiLainlain                      *string    `gorm:"column:pre_induksi_lainlain" json:"pre_induksi_lainlain"`
	TeknikAlatHiopotensi                    *string    `gorm:"column:teknik_alat_hiopotensi" json:"teknik_alat_hiopotensi"`
	TeknikAlatTci                           *string    `gorm:"column:teknik_alat_tci" json:"teknik_alat_tci"`
	TeknikAlatCpb                           *string    `gorm:"column:teknik_alat_cpb" json:"teknik_alat_cpb"`
	TeknikAlatVentilasi                     *string    `gorm:"column:teknik_alat_ventilasi" json:"teknik_alat_ventilasi"`
	TeknikAlatBroncoskopy                   *string    `gorm:"column:teknik_alat_broncoskopy" json:"teknik_alat_broncoskopy"`
	TeknikAlatGlidescopi                    *string    `gorm:"column:teknik_alat_glidescopi" json:"teknik_alat_glidescopi"`
	TeknikAlatUsg                           *string    `gorm:"column:teknik_alat_usg" json:"teknik_alat_usg"`
	TeknikAlatStimulatorSaraf               *string    `gorm:"column:teknik_alat_stimulator_saraf" json:"teknik_alat_stimulator_saraf"`
	TeknikAlatLainlain                      *string    `gorm:"column:teknik_alat_lainlain" json:"teknik_alat_lainlain"`
	MonitoringEkg                           *string    `gorm:"column:monitoring_ekg" json:"monitoring_ekg"`
	MonitoringEkgKeterangan                 *string    `gorm:"column:monitoring_ekg_keterangan" json:"monitoring_ekg_keterangan"`
	MonitoringArteri                        *string    `gorm:"column:monitoring_arteri" json:"monitoring_arteri"`
	MonitoringArteriKeterangan              *string    `gorm:"column:monitoring_arteri_keterangan" json:"monitoring_arteri_keterangan"`
	MonitoringCvp                           *string    `gorm:"column:monitoring_cvp" json:"monitoring_cvp"`
	MonitoringCvpKeterangan                 *string    `gorm:"column:monitoring_cvp_keterangan" json:"monitoring_cvp_keterangan"`
	MonitoringEtco                          *string    `gorm:"column:monitoring_etco" json:"monitoring_etco"`
	MonitoringStetoskop                     *string    `gorm:"column:monitoring_stetoskop" json:"monitoring_stetoskop"`
	MonitoringNibp                          *string    `gorm:"column:monitoring_nibp" json:"monitoring_nibp"`
	MonitoringNgt                           *string    `gorm:"column:monitoring_ngt" json:"monitoring_ngt"`
	MonitoringBis                           *string    `gorm:"column:monitoring_bis" json:"monitoring_bis"`
	MonitoringCathAPulmo                    *string    `gorm:"column:monitoring_cath_a_pulmo" json:"monitoring_cath_a_pulmo"`
	MonitoringSpo2                          *string    `gorm:"column:monitoring_spo2" json:"monitoring_spo2"`
	MonitoringKateter                       *string    `gorm:"column:monitoring_kateter" json:"monitoring_kateter"`
	MonitoringTemp                          *string    `gorm:"column:monitoring_temp" json:"monitoring_temp"`
	MonitoringLainlain                      *string    `gorm:"column:monitoring_lainlain" json:"monitoring_lainlain"`
	StatusFisikAsa                          *string    `gorm:"column:status_fisik_asa" json:"status_fisik_asa"`
	StatusFisikAlergi                       string     `gorm:"column:status_fisik_alergi" json:"status_fisik_alergi"`
	StatusFisikAlergiKeterangan             *string    `gorm:"column:status_fisik_alergi_keterangan" json:"status_fisik_alergi_keterangan"`
	StatusFisikPenyulitSedasi               *string    `gorm:"column:status_fisik_penyulit_sedasi" json:"status_fisik_penyulit_sedasi"`
	PerencanaanLanjut                       *string    `gorm:"column:perencanaan_lanjut" json:"perencanaan_lanjut"`
	PerencanaanLanjutSedasi                 *string    `gorm:"column:perencanaan_lanjut_sedasi" json:"perencanaan_lanjut_sedasi"`
	PerencanaanLanjutSedasiKeterangan       *string    `gorm:"column:perencanaan_lanjut_sedasi_keterangan" json:"perencanaan_lanjut_sedasi_keterangan"`
	PerencanaanLanjutSpinal                 *string    `gorm:"column:perencanaan_lanjut_spinal" json:"perencanaan_lanjut_spinal"`
	PerencanaanLanjutAnestesiUmum           *string    `gorm:"column:perencanaan_lanjut_anestesi_umum" json:"perencanaan_lanjut_anestesi_umum"`
	PerencanaanLanjutAnestesiUmumKeterangan *string    `gorm:"column:perencanaan_lanjut_anestesi_umum_keterangan" json:"perencanaan_lanjut_anestesi_umum_keterangan"`
	PerencanaanLanjutBlokPerifer            *string    `gorm:"column:perencanaan_lanjut_blok_perifer" json:"perencanaan_lanjut_blok_perifer"`
	PerencanaanLanjutBlokPeriferKeterangan  *string    `gorm:"column:perencanaan_lanjut_blok_perifer_keterangan" json:"perencanaan_lanjut_blok_perifer_keterangan"`
	PerencanaanLanjutEpidural               *string    `gorm:"column:perencanaan_lanjut_epidural" json:"perencanaan_lanjut_epidural"`
	PerencanaanBatal                        *string    `gorm:"column:perencanaan_batal" json:"perencanaan_batal"`
	PerencanaanBatalAlasan                  *string    `gorm:"column:perencanaan_batal_alasan" json:"perencanaan_batal_alasan"`
	NipPerawatOk                            *string    `gorm:"column:nip_perawat_ok" json:"nip_perawat_ok"`
	NipPerawatAnestesi                      *string    `gorm:"column:nip_perawat_anestesi" json:"nip_perawat_anestesi"`
}

func (CatatanAnestesiSedasi) TableName() string {
	return "catatan_anestesi_sedasi"
}
