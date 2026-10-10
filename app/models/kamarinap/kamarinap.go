// Package kamarinap model penempatan pasien rawat inap (tabel kamar_inap).
package kamarinap

// KamarInap baris kamar_inap beserta data pasien/kamar. Tanggal dikirim sebagai string karena
// Khanza memakai '0000-00-00' untuk pasien yang belum keluar.
type KamarInap struct {
	NoRawat       string  `gorm:"column:no_rawat" json:"no_rawat"`
	NoRkmMedis    string  `gorm:"column:no_rkm_medis" json:"no_rkm_medis"`
	NmPasien      string  `gorm:"column:nm_pasien" json:"nm_pasien"`
	KdKamar       string  `gorm:"column:kd_kamar" json:"kd_kamar"`
	KdBangsal     string  `gorm:"column:kd_bangsal" json:"kd_bangsal"`
	NmBangsal     string  `gorm:"column:nm_bangsal" json:"nm_bangsal"`
	Kelas         string  `gorm:"column:kelas" json:"kelas"`
	PngJawab      string  `gorm:"column:png_jawab" json:"png_jawab"`
	TrfKamar      float64 `gorm:"column:trf_kamar" json:"trf_kamar"`
	DiagnosaAwal  string  `gorm:"column:diagnosa_awal" json:"diagnosa_awal"`
	DiagnosaAkhir string  `gorm:"column:diagnosa_akhir" json:"diagnosa_akhir"`
	TglMasuk      string  `gorm:"column:tgl_masuk" json:"tgl_masuk"`
	JamMasuk      string  `gorm:"column:jam_masuk" json:"jam_masuk"`
	TglKeluar     string  `gorm:"column:tgl_keluar" json:"tgl_keluar"`
	JamKeluar     string  `gorm:"column:jam_keluar" json:"jam_keluar"`
	Lama          float64 `gorm:"column:lama" json:"lama"`
	TtlBiaya      float64 `gorm:"column:ttl_biaya" json:"ttl_biaya"`
	SttsPulang    string  `gorm:"column:stts_pulang" json:"stts_pulang"`
}

// Key primary key kamar_inap.
type Key struct {
	NoRawat  string
	TglMasuk string
	JamMasuk string
}

func (k KamarInap) Key() Key {
	return Key{NoRawat: k.NoRawat, TglMasuk: k.TglMasuk, JamMasuk: k.JamMasuk}
}

// Kamar status bed yang dikunci saat transaksi.
type Kamar struct {
	KdKamar    string  `gorm:"column:kd_kamar"`
	TrfKamar   float64 `gorm:"column:trf_kamar"`
	Status     string  `gorm:"column:status"`
	Statusdata string  `gorm:"column:statusdata"`
}

// Pengaturan set_jam_minimal (pengaturankamarinap).
type Pengaturan struct {
	JamMinimal     float64 `gorm:"column:lamajam"`
	HitungHariAwal bool    `gorm:"column:hariawal"`
}
