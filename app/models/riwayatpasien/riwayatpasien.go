// Package riwayatpasien model riwayat tingkat pasien yang diisi dari form bayi/kebidanan:
// riwayat persalinan (riwayat_persalinan_pasien) dan riwayat imunisasi (riwayat_imunisasi).
package riwayatpasien

// Persalinan satu riwayat kehamilan/persalinan; tgl_thn menjadi identitas baris per pasien.
type Persalinan struct {
	NoRkmMedis       string `gorm:"column:no_rkm_medis" json:"no_rkm_medis"`
	TglThn           string `gorm:"column:tgl_thn" json:"tgl_thn"`
	TempatPersalinan string `gorm:"column:tempat_persalinan" json:"tempat_persalinan"`
	UsiaHamil        string `gorm:"column:usia_hamil" json:"usia_hamil"`
	JenisPersalinan  string `gorm:"column:jenis_persalinan" json:"jenis_persalinan"`
	Penolong         string `gorm:"column:penolong" json:"penolong"`
	Penyulit         string `gorm:"column:penyulit" json:"penyulit"`
	Jk               string `gorm:"column:jk" json:"jk"`
	Bbpb             string `gorm:"column:bbpb" json:"bbpb"`
	Keadaan          string `gorm:"column:keadaan" json:"keadaan"`
}

type Imunisasi struct {
	NoRkmMedis    string `gorm:"column:no_rkm_medis" json:"no_rkm_medis"`
	KodeImunisasi string `gorm:"column:kode_imunisasi" json:"kode_imunisasi"`
	NamaImunisasi string `gorm:"column:nama_imunisasi" json:"nama_imunisasi"`
	NoImunisasi   int    `gorm:"column:no_imunisasi" json:"no_imunisasi"`
}
