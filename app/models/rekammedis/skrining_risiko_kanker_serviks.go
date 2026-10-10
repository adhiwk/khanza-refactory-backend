package rekammedis

import "time"

// SkriningRisikoKankerServiks tabel `skrining_risiko_kanker_serviks` (skrining risiko kanker serviks, RMSkriningRisikoKankerServiks).
type SkriningRisikoKankerServiks struct {
	NoRawat                 string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                 *time.Time `gorm:"column:tanggal" json:"tanggal"`
	RiwayatPenyakitKeluarga *string    `gorm:"column:riwayat_penyakit_keluarga" json:"riwayat_penyakit_keluarga"`
	RiwayatPenyakitSendiri  *string    `gorm:"column:riwayat_penyakit_sendiri" json:"riwayat_penyakit_sendiri"`
	RisikoMerokok           *string    `gorm:"column:risiko_merokok" json:"risiko_merokok"`
	RisikoKurangFisik       *string    `gorm:"column:risiko_kurang_fisik" json:"risiko_kurang_fisik"`
	RisikoGulaBerlebihan    *string    `gorm:"column:risiko_gula_berlebihan" json:"risiko_gula_berlebihan"`
	RisikoGaramBerlebihan   *string    `gorm:"column:risiko_garam_berlebihan" json:"risiko_garam_berlebihan"`
	RisikoLemakBerlebihan   *string    `gorm:"column:risiko_lemak_berlebihan" json:"risiko_lemak_berlebihan"`
	RisikoKurangBuahSayur   *string    `gorm:"column:risiko_kurang_buah_sayur" json:"risiko_kurang_buah_sayur"`
	RisikoAlkohol           *string    `gorm:"column:risiko_alkohol" json:"risiko_alkohol"`
	HasilIva                *string    `gorm:"column:hasil_iva" json:"hasil_iva"`
	HasilSkrining           *string    `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Keterangan              string     `gorm:"column:keterangan" json:"keterangan"`
	Nip                     string     `gorm:"column:nip" json:"nip"`
}

func (SkriningRisikoKankerServiks) TableName() string {
	return "skrining_risiko_kanker_serviks"
}
