// Package biayalain model tambahan biaya (tambahan_biaya) & potongan biaya (pengurangan_biaya) per no_rawat.
package biayalain

// Jenis tambahan atau potongan.
type Jenis string

const (
	Tambahan Jenis = "tambahan"
	Potongan Jenis = "potongan"
)

type BiayaLain struct {
	NoRawat string  `gorm:"column:no_rawat" json:"no_rawat"`
	Nama    string  `gorm:"column:nama" json:"nama"`
	Besar   float64 `gorm:"column:besar" json:"besar"`
}
