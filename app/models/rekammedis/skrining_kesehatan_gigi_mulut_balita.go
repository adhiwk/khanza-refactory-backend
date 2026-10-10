package rekammedis

import "time"

// SkriningKesehatanGigiMulutBalita tabel `skrining_kesehatan_gigi_mulut_balita` (skrining kesehatan gigi mulut balita, RMSkriningKesehatanGigiMulutBalita).
type SkriningKesehatanGigiMulutBalita struct {
	NoRawat                    string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                    *time.Time `gorm:"column:tanggal" json:"tanggal"`
	PernahPemeriksaanGigimulut *string    `gorm:"column:pernah_pemeriksaan_gigimulut" json:"pernah_pemeriksaan_gigimulut"`
	SudahTumbuhGigi            *string    `gorm:"column:sudah_tumbuh_gigi" json:"sudah_tumbuh_gigi"`
	JumlahGigiTumbuh           *string    `gorm:"column:jumlah_gigi_tumbuh" json:"jumlah_gigi_tumbuh"`
	KondisiKebersihanGigimulut *string    `gorm:"column:kondisi_kebersihan_gigimulut" json:"kondisi_kebersihan_gigimulut"`
	KebiasaanSusuBotol         *string    `gorm:"column:kebiasaan_susu_botol" json:"kebiasaan_susu_botol"`
	MengemilManis              *string    `gorm:"column:mengemil_manis" json:"mengemil_manis"`
	MenyikatGigiSebelumTidur   *string    `gorm:"column:menyikat_gigi_sebelum_tidur" json:"menyikat_gigi_sebelum_tidur"`
	MengemutMakanan            *string    `gorm:"column:mengemut_makanan" json:"mengemut_makanan"`
	LidahKotor                 *string    `gorm:"column:lidah_kotor" json:"lidah_kotor"`
	CelahBibir                 *string    `gorm:"column:celah_bibir" json:"celah_bibir"`
	HasilSkrining              *string    `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Keterangan                 string     `gorm:"column:keterangan" json:"keterangan"`
	Nip                        string     `gorm:"column:nip" json:"nip"`
}

func (SkriningKesehatanGigiMulutBalita) TableName() string {
	return "skrining_kesehatan_gigi_mulut_balita"
}
