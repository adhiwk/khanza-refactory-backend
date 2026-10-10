package rekammedis

import "time"

// SkriningKankerKolorektal tabel `skrining_kanker_kolorektal` (skrining kanker kolorektal, RMSkriningKankerKolorektal).
type SkriningKankerKolorektal struct {
	NoRawat                   string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                   *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip                       string     `gorm:"column:nip" json:"nip"`
	RiwayatPolipAdenomatosa   *string    `gorm:"column:riwayat_polip_adenomatosa" json:"riwayat_polip_adenomatosa"`
	RiwayatBabBerdarah        *string    `gorm:"column:riwayat_bab_berdarah" json:"riwayat_bab_berdarah"`
	RiwayatReseksiKuratif     *string    `gorm:"column:riwayat_reseksi_kuratif" json:"riwayat_reseksi_kuratif"`
	ColokDubur                *string    `gorm:"column:colok_dubur" json:"colok_dubur"`
	RiwayatKolorektalKeluarga *string    `gorm:"column:riwayat_kolorektal_keluarga" json:"riwayat_kolorektal_keluarga"`
	DarahSamarFeses           *string    `gorm:"column:darah_samar_feses" json:"darah_samar_feses"`
	RujukFaskesLanjut         *string    `gorm:"column:rujuk_faskes_lanjut" json:"rujuk_faskes_lanjut"`
	Kesimpulan                *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
	KeteranganKesimpulan      *string    `gorm:"column:keterangan_kesimpulan" json:"keterangan_kesimpulan"`
}

func (SkriningKankerKolorektal) TableName() string {
	return "skrining_kanker_kolorektal"
}
