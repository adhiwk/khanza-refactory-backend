package rekammedis

import "time"

// SkriningKekerasanPadaPerempuan tabel `skrining_kekerasan_pada_perempuan` (skrining kekerasan pada perempuan, RMSkriningKekerasanPadaPerempuan).
type SkriningKekerasanPadaPerempuan struct {
	NoRawat                               string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                               *time.Time `gorm:"column:tanggal" json:"tanggal"`
	MenggambarkanHubungan                 *string    `gorm:"column:menggambarkan_hubungan" json:"menggambarkan_hubungan"`
	SkorMenggambarkanHubungan             string     `gorm:"column:skor_menggambarkan_hubungan" json:"skor_menggambarkan_hubungan"`
	BerdebatDenganPasangan                *string    `gorm:"column:berdebat_dengan_pasangan" json:"berdebat_dengan_pasangan"`
	SkorBerdebatDenganPasangan            string     `gorm:"column:skor_berdebat_dengan_pasangan" json:"skor_berdebat_dengan_pasangan"`
	PertengkaranMembuatSedih              *string    `gorm:"column:pertengkaran_membuat_sedih" json:"pertengkaran_membuat_sedih"`
	SkorPertengkaranMembuatSedih          string     `gorm:"column:skor_pertengkaran_membuat_sedih" json:"skor_pertengkaran_membuat_sedih"`
	PertengkaranMenghasilkanPukulan       *string    `gorm:"column:pertengkaran_menghasilkan_pukulan" json:"pertengkaran_menghasilkan_pukulan"`
	SkorPertengkaranMenghasilkanPukulan   string     `gorm:"column:skor_pertengkaran_menghasilkan_pukulan" json:"skor_pertengkaran_menghasilkan_pukulan"`
	PernahMerasaTakutDenganPasangan       *string    `gorm:"column:pernah_merasa_takut_dengan_pasangan" json:"pernah_merasa_takut_dengan_pasangan"`
	SkorPernahMerasaTakutDenganPasangan   string     `gorm:"column:skor_pernah_merasa_takut_dengan_pasangan" json:"skor_pernah_merasa_takut_dengan_pasangan"`
	PasanganMelecehkanSecaraFisik         *string    `gorm:"column:pasangan_melecehkan_secara_fisik" json:"pasangan_melecehkan_secara_fisik"`
	SkorPasanganMelecehkanSecaraFisik     string     `gorm:"column:skor_pasangan_melecehkan_secara_fisik" json:"skor_pasangan_melecehkan_secara_fisik"`
	PasanganMelecehkanSecaraImosional     *string    `gorm:"column:pasangan_melecehkan_secara_imosional" json:"pasangan_melecehkan_secara_imosional"`
	SkorPasanganMelecehkanSecaraImosional string     `gorm:"column:skor_pasangan_melecehkan_secara_imosional" json:"skor_pasangan_melecehkan_secara_imosional"`
	PasanganMelecehkanSecaraSeksual       *string    `gorm:"column:pasangan_melecehkan_secara_seksual" json:"pasangan_melecehkan_secara_seksual"`
	SkorPasanganMelecehkanSecaraSeksual   string     `gorm:"column:skor_pasangan_melecehkan_secara_seksual" json:"skor_pasangan_melecehkan_secara_seksual"`
	Totalskor                             string     `gorm:"column:totalskor" json:"totalskor"`
	HasilSkrining                         string     `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Nip                                   string     `gorm:"column:nip" json:"nip"`
}

func (SkriningKekerasanPadaPerempuan) TableName() string {
	return "skrining_kekerasan_pada_perempuan"
}
