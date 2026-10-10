package rekammedis

import "time"

// PemantauanPewsAnak tabel `pemantauan_pews_anak` (pemantauan PEWS, RMPemantauanPEWS).
type PemantauanPewsAnak struct {
	NoRawat                    string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                    *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	ParameterPerilaku          *string    `gorm:"column:parameter_perilaku" json:"parameter_perilaku"`
	SkorPerilaku               *string    `gorm:"column:skor_perilaku" json:"skor_perilaku"`
	ParameterCrtAtauWarnaKulit *string    `gorm:"column:parameter_crt_atau_warna_kulit" json:"parameter_crt_atau_warna_kulit"`
	SkorCrtAtauWarnaKulit      *string    `gorm:"column:skor_crt_atau_warna_kulit" json:"skor_crt_atau_warna_kulit"`
	ParameterPerespirasi       *string    `gorm:"column:parameter_perespirasi" json:"parameter_perespirasi"`
	SkorPerespirasi            *string    `gorm:"column:skor_perespirasi" json:"skor_perespirasi"`
	SkorTotal                  *string    `gorm:"column:skor_total" json:"skor_total"`
	ParameterTotal             *string    `gorm:"column:parameter_total" json:"parameter_total"`
	Nip                        *string    `gorm:"column:nip" json:"nip"`
}

func (PemantauanPewsAnak) TableName() string {
	return "pemantauan_pews_anak"
}
