package obat

import (
	"time"
)

type Obat struct {
	KodeBrng     string     `gorm:"column:kode_brng;primaryKey;type:varchar(15);not null" json:"kode_brng"`
	NamaBrng     string     `gorm:"column:nama_brng;type:varchar(80)" json:"nama_brng"`
	KodeSatbesar string     `gorm:"column:kode_satbesar;type:char(4);not null" json:"kode_satbesar"`
	KodeSat      string     `gorm:"column:kode_sat;type:char(4)" json:"kode_sat"`
	LetakBarang  string     `gorm:"column:letak_barang;type:varchar(100)" json:"letak_barang"`
	Dasar        float64    `gorm:"column:dasar;type:double;not null" json:"dasar"`
	HBeli        float64    `gorm:"column:h_beli;type:double" json:"h_beli"`
	Ralan        float64    `gorm:"column:ralan;type:double" json:"ralan"`
	Kelas1       float64    `gorm:"column:kelas1;type:double" json:"kelas1"`
	Kelas2       float64    `gorm:"column:kelas2;type:double" json:"kelas2"`
	Kelas3       float64    `gorm:"column:kelas3;type:double" json:"kelas3"`
	Utama        float64    `gorm:"column:utama;type:double" json:"utama"`
	Vip          float64    `gorm:"column:vip;type:double" json:"vip"`
	Vvip         float64    `gorm:"column:vvip;type:double" json:"vvip"`
	BeliLuar     float64    `gorm:"column:beliluar;type:double" json:"beliluar"`
	JualBebas    float64    `gorm:"column:jualbebas;type:double" json:"jualbebas"`
	Karyawan     float64    `gorm:"column:karyawan;type:double" json:"karyawan"`
	StokMinimal  float64    `gorm:"column:stokminimal;type:double" json:"stokminimal"`
	Kdjns        string     `gorm:"column:kdjns;type:char(4)" json:"kdjns"`
	Isi          float64    `gorm:"column:isi;type:double;not null" json:"isi"`
	Kapasitas    float64    `gorm:"column:kapasitas;type:double;not null" json:"kapasitas"`
	Expire       *time.Time `gorm:"column:expire;type:date" json:"expire"`
	Status       string     `gorm:"column:status;type:enum('0','1');not null" json:"status"`
	KodeIndustri string     `gorm:"column:kode_industri;type:char(5)" json:"kode_industri"`
	KodeKategori string     `gorm:"column:kode_kategori;type:char(4)" json:"kode_kategori"`
	KodeGolongan string     `gorm:"column:kode_golongan;type:char(4)" json:"kode_golongan"`
}

// TableName menentukan nama tabel secara eksplisit di database
func (Obat) TableName() string {
	return "databarang"
}
