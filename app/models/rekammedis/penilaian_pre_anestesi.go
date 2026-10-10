package rekammedis

import "time"

// PenilaianPreAnestesi tabel `penilaian_pre_anestesi` (penilaian pre anastesi, RMPenilaianPreAnastesi).
type PenilaianPreAnestesi struct {
	NoRawat                      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                      *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	KdDokter                     string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	TanggalOperasi               *time.Time `gorm:"column:tanggal_operasi" json:"tanggal_operasi"`
	Diagnosa                     *string    `gorm:"column:diagnosa" json:"diagnosa"`
	RencanaTindakan              *string    `gorm:"column:rencana_tindakan" json:"rencana_tindakan"`
	Tb                           string     `gorm:"column:tb" json:"tb"`
	Bb                           string     `gorm:"column:bb" json:"bb"`
	Td                           string     `gorm:"column:td" json:"td"`
	Io2                          string     `gorm:"column:io2" json:"io2"`
	Nadi                         string     `gorm:"column:nadi" json:"nadi"`
	Pernapasan                   string     `gorm:"column:pernapasan" json:"pernapasan"`
	Suhu                         string     `gorm:"column:suhu" json:"suhu"`
	FisikCardiovasculer          *string    `gorm:"column:fisik_cardiovasculer" json:"fisik_cardiovasculer"`
	FisikParu                    *string    `gorm:"column:fisik_paru" json:"fisik_paru"`
	FisikAbdomen                 *string    `gorm:"column:fisik_abdomen" json:"fisik_abdomen"`
	FisikExtrimitas              *string    `gorm:"column:fisik_extrimitas" json:"fisik_extrimitas"`
	FisikEndokrin                *string    `gorm:"column:fisik_endokrin" json:"fisik_endokrin"`
	FisikGinjal                  *string    `gorm:"column:fisik_ginjal" json:"fisik_ginjal"`
	FisikObatobatan              *string    `gorm:"column:fisik_obatobatan" json:"fisik_obatobatan"`
	FisikLaborat                 *string    `gorm:"column:fisik_laborat" json:"fisik_laborat"`
	FisikPenunjang               *string    `gorm:"column:fisik_penunjang" json:"fisik_penunjang"`
	RiwayatPenyakitAlergiobat    *string    `gorm:"column:riwayat_penyakit_alergiobat" json:"riwayat_penyakit_alergiobat"`
	RiwayatPenyakitAlergilainnya *string    `gorm:"column:riwayat_penyakit_alergilainnya" json:"riwayat_penyakit_alergilainnya"`
	RiwayatPenyakitTerapi        *string    `gorm:"column:riwayat_penyakit_terapi" json:"riwayat_penyakit_terapi"`
	RiwayatKebiasaanMerokok      string     `gorm:"column:riwayat_kebiasaan_merokok" json:"riwayat_kebiasaan_merokok"`
	RiwayatKebiasaanKetMerokok   string     `gorm:"column:riwayat_kebiasaan_ket_merokok" json:"riwayat_kebiasaan_ket_merokok"`
	RiwayatKebiasaanAlkohol      string     `gorm:"column:riwayat_kebiasaan_alkohol" json:"riwayat_kebiasaan_alkohol"`
	RiwayatKebiasaanKetAlkohol   string     `gorm:"column:riwayat_kebiasaan_ket_alkohol" json:"riwayat_kebiasaan_ket_alkohol"`
	RiwayatKebiasaanObat         string     `gorm:"column:riwayat_kebiasaan_obat" json:"riwayat_kebiasaan_obat"`
	RiwayatKebiasaanKetObat      string     `gorm:"column:riwayat_kebiasaan_ket_obat" json:"riwayat_kebiasaan_ket_obat"`
	RiwayatMedisCardiovasculer   *string    `gorm:"column:riwayat_medis_cardiovasculer" json:"riwayat_medis_cardiovasculer"`
	RiwayatMedisRespiratory      *string    `gorm:"column:riwayat_medis_respiratory" json:"riwayat_medis_respiratory"`
	RiwayatMedisEndocrine        *string    `gorm:"column:riwayat_medis_endocrine" json:"riwayat_medis_endocrine"`
	RiwayatMedisLainnya          *string    `gorm:"column:riwayat_medis_lainnya" json:"riwayat_medis_lainnya"`
	Asa                          *string    `gorm:"column:asa" json:"asa"`
	Puasa                        *time.Time `gorm:"column:puasa" json:"puasa"`
	RencanaAnestesi              *string    `gorm:"column:rencana_anestesi" json:"rencana_anestesi"`
	RencanaPerawatan             *string    `gorm:"column:rencana_perawatan" json:"rencana_perawatan"`
	CatatanKhusus                *string    `gorm:"column:catatan_khusus" json:"catatan_khusus"`
}

func (PenilaianPreAnestesi) TableName() string {
	return "penilaian_pre_anestesi"
}
