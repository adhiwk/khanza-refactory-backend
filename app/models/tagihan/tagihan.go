// Package tagihan model rincian tagihan pasien per no_rawat (bagian baca DlgBilingRalan / DlgBilingRanap).
package tagihan

// Item satu baris rincian biaya.
type Item struct {
	Kategori string  `gorm:"column:kategori" json:"kategori"`
	Nama     string  `gorm:"column:nama" json:"nama"`
	Tanggal  string  `gorm:"column:tanggal" json:"tanggal"`
	Biaya    float64 `gorm:"column:biaya" json:"biaya"`
	Jumlah   float64 `gorm:"column:jumlah" json:"jumlah"`
	Total    float64 `gorm:"column:total" json:"total"`
}

type Kategori struct {
	Kategori string  `json:"kategori"`
	Total    float64 `json:"total"`
	Items    []Item  `json:"items"`
}

type Tagihan struct {
	NoRawat      string     `json:"no_rawat"`
	NoRkmMedis   string     `json:"no_rkm_medis"`
	NmPasien     string     `json:"nm_pasien"`
	StatusLanjut string     `json:"status_lanjut"`
	StatusBayar  string     `json:"status_bayar"`
	Rincian      []Kategori `json:"rincian"`
	TotalBiaya   float64    `json:"total_biaya"`
	Potongan     float64    `json:"potongan"`
	Deposit      float64    `json:"deposit"`
	SisaTagihan  float64    `json:"sisa_tagihan"`
}
