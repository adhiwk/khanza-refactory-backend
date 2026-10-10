package rekammedis

import "time"

// PenilaianPsikologi tabel `penilaian_psikologi` (penilaian psikologi, RMPenilaianPsikologi).
type PenilaianPsikologi struct {
	NoRawat            string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal            *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip                string     `gorm:"column:nip" json:"nip"`
	Anamnesis          string     `gorm:"column:anamnesis" json:"anamnesis"`
	DikirimDari        string     `gorm:"column:dikirim_dari" json:"dikirim_dari"`
	TujuanPemeriksaan  string     `gorm:"column:tujuan_pemeriksaan" json:"tujuan_pemeriksaan"`
	KetAnamnesis       string     `gorm:"column:ket_anamnesis" json:"ket_anamnesis"`
	Rupa               string     `gorm:"column:rupa" json:"rupa"`
	BentukTubuh        string     `gorm:"column:bentuk_tubuh" json:"bentuk_tubuh"`
	Tindakan           string     `gorm:"column:tindakan" json:"tindakan"`
	Pakaian            string     `gorm:"column:pakaian" json:"pakaian"`
	Ekspresi           string     `gorm:"column:ekspresi" json:"ekspresi"`
	Berbicara          string     `gorm:"column:berbicara" json:"berbicara"`
	PenggunaanKata     string     `gorm:"column:penggunaan_kata" json:"penggunaan_kata"`
	CiriMenyolok       string     `gorm:"column:ciri_menyolok" json:"ciri_menyolok"`
	HasilPsikotes      string     `gorm:"column:hasil_psikotes" json:"hasil_psikotes"`
	Kepribadian        string     `gorm:"column:kepribadian" json:"kepribadian"`
	Psikodinamika      string     `gorm:"column:psikodinamika" json:"psikodinamika"`
	KesimpulanPsikolog string     `gorm:"column:kesimpulan_psikolog" json:"kesimpulan_psikolog"`
}

func (PenilaianPsikologi) TableName() string {
	return "penilaian_psikologi"
}
