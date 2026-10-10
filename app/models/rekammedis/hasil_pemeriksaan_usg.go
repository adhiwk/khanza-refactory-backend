package rekammedis

import "time"

// HasilPemeriksaanUsg tabel `hasil_pemeriksaan_usg` (hasil pemeriksaan USG, RMHasilPemeriksaanUSG).
type HasilPemeriksaanUsg struct {
	NoRawat               string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal               *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter              string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis        *string    `gorm:"column:diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari           *string    `gorm:"column:kiriman_dari" json:"kiriman_dari"`
	Hta                   *string    `gorm:"column:hta" json:"hta"`
	KantongGestasi        *string    `gorm:"column:kantong_gestasi" json:"kantong_gestasi"`
	UkuranBokongkepala    *string    `gorm:"column:ukuran_bokongkepala" json:"ukuran_bokongkepala"`
	JenisPrestasi         *string    `gorm:"column:jenis_prestasi" json:"jenis_prestasi"`
	DiameterBiparietal    *string    `gorm:"column:diameter_biparietal" json:"diameter_biparietal"`
	PanjangFemur          *string    `gorm:"column:panjang_femur" json:"panjang_femur"`
	LingkarAbdomen        *string    `gorm:"column:lingkar_abdomen" json:"lingkar_abdomen"`
	TafsiranBeratJanin    *string    `gorm:"column:tafsiran_berat_janin" json:"tafsiran_berat_janin"`
	UsiaKehamilan         *string    `gorm:"column:usia_kehamilan" json:"usia_kehamilan"`
	PlasentaBerimplatansi *string    `gorm:"column:plasenta_berimplatansi" json:"plasenta_berimplatansi"`
	DerajatMaturitas      *string    `gorm:"column:derajat_maturitas" json:"derajat_maturitas"`
	JumlahAirKetuban      *string    `gorm:"column:jumlah_air_ketuban" json:"jumlah_air_ketuban"`
	IndekCairanKetuban    *string    `gorm:"column:indek_cairan_ketuban" json:"indek_cairan_ketuban"`
	KelainanKongenital    *string    `gorm:"column:kelainan_kongenital" json:"kelainan_kongenital"`
	PeluangSex            *string    `gorm:"column:peluang_sex" json:"peluang_sex"`
	Kesimpulan            *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
}

func (HasilPemeriksaanUsg) TableName() string {
	return "hasil_pemeriksaan_usg"
}
