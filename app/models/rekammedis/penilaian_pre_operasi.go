package rekammedis

import "time"

// PenilaianPreOperasi tabel `penilaian_pre_operasi` (penilaian pre operasi, RMPenilaianPreOperasi).
type PenilaianPreOperasi struct {
	NoRawat                     string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                     *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	KdDokter                    string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	RingkasanKlinik             *string    `gorm:"column:ringkasan_klinik" json:"ringkasan_klinik"`
	PemeriksaanFisik            *string    `gorm:"column:pemeriksaan_fisik" json:"pemeriksaan_fisik"`
	PemeriksaanDiagnostik       *string    `gorm:"column:pemeriksaan_diagnostik" json:"pemeriksaan_diagnostik"`
	DiagnosaPreOperasi          *string    `gorm:"column:diagnosa_pre_operasi" json:"diagnosa_pre_operasi"`
	RencanaTindakanBedah        *string    `gorm:"column:rencana_tindakan_bedah" json:"rencana_tindakan_bedah"`
	HalHalYangPerludiPersiapkan *string    `gorm:"column:hal_hal_yang_perludi_persiapkan" json:"hal_hal_yang_perludi_persiapkan"`
	TerapiPreOperasi            *string    `gorm:"column:terapi_pre_operasi" json:"terapi_pre_operasi"`
}

func (PenilaianPreOperasi) TableName() string {
	return "penilaian_pre_operasi"
}
