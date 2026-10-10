// Package masterkode model master kode sederhana rekam medis (masalah/rencana keperawatan, triase, imunisasi).
package masterkode

// Kode satu baris master dengan bentuk seragam; KodeInduk terisi untuk master turunan
// (rencana -> masalah keperawatan, skala triase -> pemeriksaan triase).
type Kode struct {
	Kode      string `gorm:"column:kode" json:"kode"`
	Nama      string `gorm:"column:nama" json:"nama"`
	KodeInduk string `gorm:"column:kode_induk" json:"kode_induk,omitempty"`
	NamaInduk string `gorm:"column:nama_induk" json:"nama_induk,omitempty"`
}
