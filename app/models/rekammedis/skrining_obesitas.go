package rekammedis

import "time"

// SkriningObesitas tabel `skrining_obesitas` (skrining obesitas, RMSkriningObesitas).
type SkriningObesitas struct {
	NoRawat                            string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                            *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KebiasaanMakanManis                *string    `gorm:"column:kebiasaan_makan_manis" json:"kebiasaan_makan_manis"`
	AktifitasFisikSetiapHari           *string    `gorm:"column:aktifitas_fisik_setiap_hari" json:"aktifitas_fisik_setiap_hari"`
	IstirahatCukup                     *string    `gorm:"column:istirahat_cukup" json:"istirahat_cukup"`
	RisikoMerokok                      *string    `gorm:"column:risiko_merokok" json:"risiko_merokok"`
	RiwayatMinumAlkoholMerokokKeluarga *string    `gorm:"column:riwayat_minum_alkohol_merokok_keluarga" json:"riwayat_minum_alkohol_merokok_keluarga"`
	RiwayatPenggunaanObatSteroid       *string    `gorm:"column:riwayat_penggunaan_obat_steroid" json:"riwayat_penggunaan_obat_steroid"`
	BeratBadan                         *string    `gorm:"column:berat_badan" json:"berat_badan"`
	TinggiBadan                        *string    `gorm:"column:tinggi_badan" json:"tinggi_badan"`
	Imt                                *string    `gorm:"column:imt" json:"imt"`
	KasifikasiImt                      *string    `gorm:"column:kasifikasi_imt" json:"kasifikasi_imt"`
	LingkarPinggang                    *string    `gorm:"column:lingkar_pinggang" json:"lingkar_pinggang"`
	RisikoLingkarPinggang              *string    `gorm:"column:risiko_lingkar_pinggang" json:"risiko_lingkar_pinggang"`
	StatusObesitas                     *string    `gorm:"column:status_obesitas" json:"status_obesitas"`
	Keterangan                         *string    `gorm:"column:keterangan" json:"keterangan"`
	Nip                                string     `gorm:"column:nip" json:"nip"`
}

func (SkriningObesitas) TableName() string {
	return "skrining_obesitas"
}
