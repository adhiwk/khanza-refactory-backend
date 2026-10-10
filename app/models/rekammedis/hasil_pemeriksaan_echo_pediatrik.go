package rekammedis

import "time"

// HasilPemeriksaanEchoPediatrik tabel `hasil_pemeriksaan_echo_pediatrik` (hasil pemeriksaan echo pediatrik, RMHasilPemeriksaanEchoPediatrik).
type HasilPemeriksaanEchoPediatrik struct {
	NoRawat                string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter               string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis         *string    `gorm:"column:diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari            *string    `gorm:"column:kiriman_dari" json:"kiriman_dari"`
	Situs                  *string    `gorm:"column:situs" json:"situs"`
	AvVa                   *string    `gorm:"column:av_va" json:"av_va"`
	DrainaseVenaPulmonalis *string    `gorm:"column:drainase_vena_pulmonalis" json:"drainase_vena_pulmonalis"`
	KatupMitral            *string    `gorm:"column:katup_mitral" json:"katup_mitral"`
	KatupAorta             *string    `gorm:"column:katup_aorta" json:"katup_aorta"`
	KatupTricuspid         *string    `gorm:"column:katup_tricuspid" json:"katup_tricuspid"`
	KatupPulmonal          *string    `gorm:"column:katup_pulmonal" json:"katup_pulmonal"`
	KatupSeptumAtrium      *string    `gorm:"column:katup_septum_atrium" json:"katup_septum_atrium"`
	KatupSeptumVentrikal   *string    `gorm:"column:katup_septum_ventrikal" json:"katup_septum_ventrikal"`
	KatupArkusAorta        *string    `gorm:"column:katup_arkus_aorta" json:"katup_arkus_aorta"`
	KatupKeteranganLainnya *string    `gorm:"column:katup_keterangan_lainnya" json:"katup_keterangan_lainnya"`
	RuangJantung           *string    `gorm:"column:ruang_jantung" json:"ruang_jantung"`
	ModeIvds               *string    `gorm:"column:mode_ivds" json:"mode_ivds"`
	ModeIvss               *string    `gorm:"column:mode_ivss" json:"mode_ivss"`
	ModeLvidDextra         *string    `gorm:"column:mode_lvid_dextra" json:"mode_lvid_dextra"`
	ModeLvidSinistra       *string    `gorm:"column:mode_lvid_sinistra" json:"mode_lvid_sinistra"`
	ModeLvpwDextra         *string    `gorm:"column:mode_lvpw_dextra" json:"mode_lvpw_dextra"`
	ModeLvpwSinistra       *string    `gorm:"column:mode_lvpw_sinistra" json:"mode_lvpw_sinistra"`
	ModeEjectionFraction   *string    `gorm:"column:mode_ejection_fraction" json:"mode_ejection_fraction"`
	ModeFractionShotening  *string    `gorm:"column:mode_fraction_shotening" json:"mode_fraction_shotening"`
	Doppler                *string    `gorm:"column:doppler" json:"doppler"`
	Kesimpulan             *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
	Saran                  *string    `gorm:"column:saran" json:"saran"`
}

func (HasilPemeriksaanEchoPediatrik) TableName() string {
	return "hasil_pemeriksaan_echo_pediatrik"
}
