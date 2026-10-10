// Package pemberianobat model pemberian obat ke pasien (detail_pemberian_obat).
package pemberianobat

// Barang obat di satu depo beserta seluruh kolom harga (databarang atau data_batch).
type Barang struct {
	KodeBrng  string  `gorm:"column:kode_brng" json:"kode_brng"`
	NamaBrng  string  `gorm:"column:nama_brng" json:"nama_brng"`
	KodeSat   string  `gorm:"column:kode_sat" json:"kode_sat"`
	NoBatch   string  `gorm:"column:no_batch" json:"no_batch"`
	NoFaktur  string  `gorm:"column:no_faktur" json:"no_faktur"`
	Stok      float64 `gorm:"column:stok" json:"stok"`
	HBeli     float64 `gorm:"column:h_beli" json:"-"`
	Hpp       float64 `gorm:"column:hpp" json:"-"`
	Ralan     float64 `gorm:"column:ralan" json:"-"`
	Kelas1    float64 `gorm:"column:kelas1" json:"-"`
	Kelas2    float64 `gorm:"column:kelas2" json:"-"`
	Kelas3    float64 `gorm:"column:kelas3" json:"-"`
	Utama     float64 `gorm:"column:utama" json:"-"`
	Vip       float64 `gorm:"column:vip" json:"-"`
	Vvip      float64 `gorm:"column:vvip" json:"-"`
	Beliluar  float64 `gorm:"column:beliluar" json:"-"`
	Karyawan  float64 `gorm:"column:karyawan" json:"-"`
	HargaJual float64 `gorm:"-" json:"harga"`
}

// Kolom harga jual sesuai jenis harga (pilihan Jeniskelas DlgCariObat / DlgCariObat2).
func (b Barang) Kolom(jenis string) float64 {
	switch jenis {
	case "kelas1":
		return b.Kelas1
	case "kelas2":
		return b.Kelas2
	case "kelas3":
		return b.Kelas3
	case "utama":
		return b.Utama
	case "vip":
		return b.Vip
	case "vvip":
		return b.Vvip
	case "beliluar":
		return b.Beliluar
	case "karyawan":
		return b.Karyawan
	default:
		return b.Ralan
	}
}

// Pemberian satu baris detail_pemberian_obat.
type Pemberian struct {
	TglPerawatan string  `gorm:"column:tgl_perawatan" json:"tgl_perawatan"`
	Jam          string  `gorm:"column:jam" json:"jam"`
	NoRawat      string  `gorm:"column:no_rawat" json:"no_rawat"`
	KodeBrng     string  `gorm:"column:kode_brng" json:"kode_brng"`
	NamaBrng     string  `gorm:"column:nama_brng" json:"nama_brng"`
	HBeli        float64 `gorm:"column:h_beli" json:"h_beli"`
	BiayaObat    float64 `gorm:"column:biaya_obat" json:"biaya_obat"`
	Jml          float64 `gorm:"column:jml" json:"jml"`
	Embalase     float64 `gorm:"column:embalase" json:"embalase"`
	Tuslah       float64 `gorm:"column:tuslah" json:"tuslah"`
	Total        float64 `gorm:"column:total" json:"total"`
	Status       string  `gorm:"column:status" json:"status"`
	KdBangsal    string  `gorm:"column:kd_bangsal" json:"kd_bangsal"`
	NoBatch      string  `gorm:"column:no_batch" json:"no_batch"`
	NoFaktur     string  `gorm:"column:no_faktur" json:"no_faktur"`
	AturanPakai  string  `gorm:"column:aturan" json:"aturan_pakai"`
}

// Key primary key detail_pemberian_obat.
type Key struct {
	NoRawat      string
	TglPerawatan string
	Jam          string
	KodeBrng     string
	NoBatch      string
	NoFaktur     string
}
