package rekammedis

import "time"

// PenilaianUlangNyeri tabel `penilaian_ulang_nyeri` (penilaian ulang nyeri, RMPenilaianUlangNyeri).
type PenilaianUlangNyeri struct {
	NoRawat     string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal     *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Nyeri       string     `gorm:"column:nyeri" json:"nyeri"`
	Provokes    string     `gorm:"column:provokes" json:"provokes"`
	KetProvokes string     `gorm:"column:ket_provokes" json:"ket_provokes"`
	Quality     string     `gorm:"column:quality" json:"quality"`
	KetQuality  string     `gorm:"column:ket_quality" json:"ket_quality"`
	Lokasi      string     `gorm:"column:lokasi" json:"lokasi"`
	Menyebar    string     `gorm:"column:menyebar" json:"menyebar"`
	SkalaNyeri  string     `gorm:"column:skala_nyeri" json:"skala_nyeri"`
	Durasi      string     `gorm:"column:durasi" json:"durasi"`
	NyeriHilang string     `gorm:"column:nyeri_hilang" json:"nyeri_hilang"`
	KetNyeri    string     `gorm:"column:ket_nyeri" json:"ket_nyeri"`
	Nip         string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianUlangNyeri) TableName() string {
	return "penilaian_ulang_nyeri"
}
