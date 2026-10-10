package rekammedis

import "time"

// SkriningPuma tabel `skrining_puma` (skrining PUMA, RMSkriningPUMA).
type SkriningPuma struct {
	NoRawat                 string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                 *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Jk                      *string    `gorm:"column:jk" json:"jk"`
	NilaiJk                 *int       `gorm:"column:nilai_jk" json:"nilai_jk"`
	Usia                    *string    `gorm:"column:usia" json:"usia"`
	NilaiUsia               *int       `gorm:"column:nilai_usia" json:"nilai_usia"`
	PernahMerokok           *string    `gorm:"column:pernah_merokok" json:"pernah_merokok"`
	NilaiPernahMerokok      *int       `gorm:"column:nilai_pernah_merokok" json:"nilai_pernah_merokok"`
	JumlahRokokPerhari      *string    `gorm:"column:jumlah_rokok_perhari" json:"jumlah_rokok_perhari"`
	LamaMerokok             *string    `gorm:"column:lama_merokok" json:"lama_merokok"`
	NapasPendek             *string    `gorm:"column:napas_pendek" json:"napas_pendek"`
	NilaiNapasPendek        *int       `gorm:"column:nilai_napas_pendek" json:"nilai_napas_pendek"`
	PunyaDahak              *string    `gorm:"column:punya_dahak" json:"punya_dahak"`
	NilaiPunyaDahak         *int       `gorm:"column:nilai_punya_dahak" json:"nilai_punya_dahak"`
	BiasaBatuk              *string    `gorm:"column:biasa_batuk" json:"biasa_batuk"`
	NilaiBiasaBatuk         *int       `gorm:"column:nilai_biasa_batuk" json:"nilai_biasa_batuk"`
	Spirometri              *string    `gorm:"column:spirometri" json:"spirometri"`
	NilaiSpirometri         *int       `gorm:"column:nilai_spirometri" json:"nilai_spirometri"`
	NilaiTotal              *int       `gorm:"column:nilai_total" json:"nilai_total"`
	KeteranganHasilSkrining *string    `gorm:"column:keterangan_hasil_skrining" json:"keterangan_hasil_skrining"`
	Nip                     string     `gorm:"column:nip" json:"nip"`
}

func (SkriningPuma) TableName() string {
	return "skrining_puma"
}
