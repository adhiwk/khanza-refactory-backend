// Package skriningralan model skrining pasien rawat jalan sebelum registrasi (skrining_rawat_jalan).
package skriningralan

type Skrining struct {
	Tanggal     string `gorm:"column:tanggal" json:"tanggal"`
	Jam         string `gorm:"column:jam" json:"jam"`
	NoRkmMedis  string `gorm:"column:no_rkm_medis" json:"no_rkm_medis"`
	NmPasien    string `gorm:"column:nm_pasien" json:"nm_pasien"`
	Geriatri    string `gorm:"column:geriatri" json:"geriatri"`
	Kesadaran   string `gorm:"column:kesadaran" json:"kesadaran"`
	Pernapasan  string `gorm:"column:pernapasan" json:"pernapasan"`
	NyeriDada   string `gorm:"column:nyeri_dada" json:"nyeri_dada"`
	SkalaNyeri  string `gorm:"column:skala_nyeri" json:"skala_nyeri"`
	Batuk       string `gorm:"column:batuk" json:"batuk"`
	RisikoJatuh string `gorm:"column:risiko_jatuh" json:"risiko_jatuh"`
	Keputusan   string `gorm:"column:keputusan" json:"keputusan"`
	Nip         string `gorm:"column:nip" json:"nip"`
	NmPetugas   string `gorm:"column:nm_petugas" json:"nm_petugas"`
}

// Key primary key skrining_rawat_jalan.
type Key struct {
	Tanggal    string
	Jam        string
	NoRkmMedis string
}
