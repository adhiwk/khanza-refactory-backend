package rekammedis

import "time"

// SkriningKesehatanGigiMulutRemaja tabel `skrining_kesehatan_gigi_mulut_remaja` (skrining kesehatan gigi mulut remaja, RMSkriningKesehatanGigiMulutRemaja).
type SkriningKesehatanGigiMulutRemaja struct {
	NoRawat                    string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                    *time.Time `gorm:"column:tanggal" json:"tanggal"`
	PernahPemeriksaanGigimulut *string    `gorm:"column:pernah_pemeriksaan_gigimulut" json:"pernah_pemeriksaan_gigimulut"`
	JumlahGigiTumbuh           *string    `gorm:"column:jumlah_gigi_tumbuh" json:"jumlah_gigi_tumbuh"`
	KondisiKebersihanGigimulut *string    `gorm:"column:kondisi_kebersihan_gigimulut" json:"kondisi_kebersihan_gigimulut"`
	PunyaGigiBerlubang         *string    `gorm:"column:punya_gigi_berlubang" json:"punya_gigi_berlubang"`
	PernahGusiBerdarah         *string    `gorm:"column:pernah_gusi_berdarah" json:"pernah_gusi_berdarah"`
	PunyaKarangGigi            *string    `gorm:"column:punya_karang_gigi" json:"punya_karang_gigi"`
	GigiDepanTidakTeratur      *string    `gorm:"column:gigi_depan_tidak_teratur" json:"gigi_depan_tidak_teratur"`
	MenyikatGigiSebelumTidur   *string    `gorm:"column:menyikat_gigi_sebelum_tidur" json:"menyikat_gigi_sebelum_tidur"`
	PunyaSariawan              *string    `gorm:"column:punya_sariawan" json:"punya_sariawan"`
	PemeriksaanFisik           *string    `gorm:"column:pemeriksaan_fisik" json:"pemeriksaan_fisik"`
	PemeriksaanPenunjang       *string    `gorm:"column:pemeriksaan_penunjang" json:"pemeriksaan_penunjang"`
	HasilSkrining              *string    `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Keterangan                 string     `gorm:"column:keterangan" json:"keterangan"`
	Nip                        string     `gorm:"column:nip" json:"nip"`
}

func (SkriningKesehatanGigiMulutRemaja) TableName() string {
	return "skrining_kesehatan_gigi_mulut_remaja"
}
