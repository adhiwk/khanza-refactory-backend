package rekammedis

import "time"

// SkriningAnemia tabel `skrining_anemia` (skrining anemia, RMSkriningAnemia).
type SkriningAnemia struct {
	NoRawat            string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal            *time.Time `gorm:"column:tanggal" json:"tanggal"`
	MudahLelah         *string    `gorm:"column:mudah_lelah" json:"mudah_lelah"`
	BuahSayur          *string    `gorm:"column:buah_sayur" json:"buah_sayur"`
	ProteinHewani      *string    `gorm:"column:protein_hewani" json:"protein_hewani"`
	MasalahPubertas    *string    `gorm:"column:masalah_pubertas" json:"masalah_pubertas"`
	RisikoIms          *string    `gorm:"column:risiko_ims" json:"risiko_ims"`
	KekerasanSeksual   *string    `gorm:"column:kekerasan_seksual" json:"kekerasan_seksual"`
	SudahMenstruasi    *string    `gorm:"column:sudah_menstruasi" json:"sudah_menstruasi"`
	GangguanMenstruasi *string    `gorm:"column:gangguan_menstruasi" json:"gangguan_menstruasi"`
	TambahDarah        *string    `gorm:"column:tambah_darah" json:"tambah_darah"`
	KelainanDarah      *string    `gorm:"column:kelainan_darah" json:"kelainan_darah"`
	KeluargaThalasemia *string    `gorm:"column:keluarga_thalasemia" json:"keluarga_thalasemia"`
	Rambut             *string    `gorm:"column:rambut" json:"rambut"`
	Kulit              *string    `gorm:"column:kulit" json:"kulit"`
	BekasSutikan       *string    `gorm:"column:bekas_sutikan" json:"bekas_sutikan"`
	Kuku               *string    `gorm:"column:kuku" json:"kuku"`
	TandaKlinis        string     `gorm:"column:tanda_klinis" json:"tanda_klinis"`
	PemeriksaanHb      *string    `gorm:"column:pemeriksaan_hb" json:"pemeriksaan_hb"`
	KadarHb            *string    `gorm:"column:kadar_hb" json:"kadar_hb"`
	JenisAnemia        *string    `gorm:"column:jenis_anemia" json:"jenis_anemia"`
	HasilSkrining      string     `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Keterangan         *string    `gorm:"column:keterangan" json:"keterangan"`
	Nip                string     `gorm:"column:nip" json:"nip"`
}

func (SkriningAnemia) TableName() string {
	return "skrining_anemia"
}
