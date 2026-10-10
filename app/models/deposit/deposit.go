// Package deposit model uang muka pasien (tabel deposit).
package deposit

type Deposit struct {
	NoDeposit    string  `gorm:"column:no_deposit" json:"no_deposit"`
	NoRawat      string  `gorm:"column:no_rawat" json:"no_rawat"`
	TglDeposit   string  `gorm:"column:tgl_deposit" json:"tgl_deposit"`
	NamaBayar    string  `gorm:"column:nama_bayar" json:"nama_bayar"`
	Besarppn     float64 `gorm:"column:besarppn" json:"besarppn"`
	BesarDeposit float64 `gorm:"column:besar_deposit" json:"besar_deposit"`
	Nip          string  `gorm:"column:nip" json:"nip"`
	Keterangan   string  `gorm:"column:keterangan" json:"keterangan"`
}

// AkunBayar cara bayar beserta rekening & persen PPN (akun_bayar).
type AkunBayar struct {
	NamaBayar string  `gorm:"column:nama_bayar"`
	KdRek     string  `gorm:"column:kd_rek"`
	Ppn       float64 `gorm:"column:ppn"`
}
