package rekammedis

import "time"

// DeteksiDiniCorona tabel `deteksi_dini_corona` (deteksi dini corona, RMDeteksiDiniCorona).
type DeteksiDiniCorona struct {
	NoRawat                    string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                    *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip                        *string    `gorm:"column:nip" json:"nip"`
	GejalaDemam                *string    `gorm:"column:gejala_demam" json:"gejala_demam"`
	GejalaBatuk                *string    `gorm:"column:gejala_batuk" json:"gejala_batuk"`
	GejalaSesak                *string    `gorm:"column:gejala_sesak" json:"gejala_sesak"`
	GejalaTanggalPertama       *time.Time `gorm:"column:gejala_tanggal_pertama" json:"gejala_tanggal_pertama"`
	GejalaRiwayatSakit         *string    `gorm:"column:gejala_riwayat_sakit" json:"gejala_riwayat_sakit"`
	GejalaRiwayatPeriksa       *string    `gorm:"column:gejala_riwayat_periksa" json:"gejala_riwayat_periksa"`
	FaktorRiwayatPerjalanan    string     `gorm:"column:faktor_riwayat_perjalanan" json:"faktor_riwayat_perjalanan"`
	FaktorAsalDaerah           string     `gorm:"column:faktor_asal_daerah" json:"faktor_asal_daerah"`
	FaktorTanggalKedatangan    *time.Time `gorm:"column:faktor_tanggal_kedatangan" json:"faktor_tanggal_kedatangan"`
	FaktorPaparanKontakpositif string     `gorm:"column:faktor_paparan_kontakpositif" json:"faktor_paparan_kontakpositif"`
	FaktorPaparanKontakpdp     string     `gorm:"column:faktor_paparan_kontakpdp" json:"faktor_paparan_kontakpdp"`
	FaktorPaparanFaskespositif string     `gorm:"column:faktor_paparan_faskespositif" json:"faktor_paparan_faskespositif"`
	FaktorPaparanPerjalananln  string     `gorm:"column:faktor_paparan_perjalananln" json:"faktor_paparan_perjalananln"`
	FaktorPaparanPasarhewan    string     `gorm:"column:faktor_paparan_pasarhewan" json:"faktor_paparan_pasarhewan"`
	Kesimpulan                 string     `gorm:"column:kesimpulan" json:"kesimpulan"`
	TindakLanjut               string     `gorm:"column:tindak_lanjut" json:"tindak_lanjut"`
}

func (DeteksiDiniCorona) TableName() string {
	return "deteksi_dini_corona"
}
