// Package icd model diagnosa (ICD-10, diagnosa_pasien) dan prosedur (ICD-9, prosedur_pasien) pasien.
package icd

// Jenis diagnosa atau prosedur.
type Jenis string

const (
	Diagnosa Jenis = "diagnosa"
	Prosedur Jenis = "prosedur"
)

// Kode satu baris diagnosa_pasien / prosedur_pasien.
type Kode struct {
	NoRawat        string `gorm:"column:no_rawat" json:"no_rawat"`
	TglRegistrasi  string `gorm:"column:tgl_registrasi" json:"tgl_registrasi"`
	Kode           string `gorm:"column:kode" json:"kode"`
	Nama           string `gorm:"column:nama" json:"nama"`
	Status         string `gorm:"column:status" json:"status"`
	Prioritas      int    `gorm:"column:prioritas" json:"prioritas"`
	StatusPenyakit string `gorm:"column:status_penyakit" json:"status_penyakit,omitempty"`
	Jumlah         string `gorm:"column:jumlah" json:"jumlah,omitempty"`
}

// Referensi master ICD (penyakit / icd9) untuk pencarian.
type Referensi struct {
	Kode string `gorm:"column:kode" json:"kode"`
	Nama string `gorm:"column:nama" json:"nama"`
}

// Item input simpan.
type Item struct {
	Kode      string
	Prioritas int
	Jumlah    string
}
