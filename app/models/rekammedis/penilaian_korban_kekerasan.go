package rekammedis

import "time"

// PenilaianKorbanKekerasan tabel `penilaian_korban_kekerasan` (penilaian korban kekerasan, RMPenilaianKorbanKekerasan).
type PenilaianKorbanKekerasan struct {
	NoRawat                     string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                     *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Informasi                   *string    `gorm:"column:informasi" json:"informasi"`
	HubunganDenganPasien        *string    `gorm:"column:hubungan_dengan_pasien" json:"hubungan_dengan_pasien"`
	JumlahSaudara               *string    `gorm:"column:jumlah_saudara" json:"jumlah_saudara"`
	KondisiKeluaga              *string    `gorm:"column:kondisi_keluaga" json:"kondisi_keluaga"`
	HubunganOrangTerdekat       *string    `gorm:"column:hubungan_orang_terdekat" json:"hubungan_orang_terdekat"`
	KekerasanYangDialami        *string    `gorm:"column:kekerasan_yang_dialami" json:"kekerasan_yang_dialami"`
	TempatKejadian              *string    `gorm:"column:tempat_kejadian" json:"tempat_kejadian"`
	LamaKekerasan               *int       `gorm:"column:lama_kekerasan" json:"lama_kekerasan"`
	PeriodeKekerasan            *string    `gorm:"column:periode_kekerasan" json:"periode_kekerasan"`
	SeberapaSeringMengalami     *string    `gorm:"column:seberapa_sering_mengalami" json:"seberapa_sering_mengalami"`
	PemicuKekerasan             *string    `gorm:"column:pemicu_kekerasan" json:"pemicu_kekerasan"`
	YangMelakukanKekerasan      *string    `gorm:"column:yang_melakukan_kekerasan" json:"yang_melakukan_kekerasan"`
	DampakKekerasan             *string    `gorm:"column:dampak_kekerasan" json:"dampak_kekerasan"`
	TandaTandaDidapatkan        *string    `gorm:"column:tanda_tanda_didapatkan" json:"tanda_tanda_didapatkan"`
	MemerlukanPendampingan      *string    `gorm:"column:memerlukan_pendampingan" json:"memerlukan_pendampingan"`
	RiwayatKelainan             *string    `gorm:"column:riwayat_kelainan" json:"riwayat_kelainan"`
	PemeriksaanKepala           *string    `gorm:"column:pemeriksaan_kepala" json:"pemeriksaan_kepala"`
	PemeriksaanThoraks          *string    `gorm:"column:pemeriksaan_thoraks" json:"pemeriksaan_thoraks"`
	PemeriksaanLeher            *string    `gorm:"column:pemeriksaan_leher" json:"pemeriksaan_leher"`
	PemeriksaanAbdomen          *string    `gorm:"column:pemeriksaan_abdomen" json:"pemeriksaan_abdomen"`
	PemeriksaanGenitalia        *string    `gorm:"column:pemeriksaan_genitalia" json:"pemeriksaan_genitalia"`
	PemeriksaanEkstrimitasAtas  *string    `gorm:"column:pemeriksaan_ekstrimitas_atas" json:"pemeriksaan_ekstrimitas_atas"`
	PemeriksaanEkstrimitasBawah string     `gorm:"column:pemeriksaan_ekstrimitas_bawah" json:"pemeriksaan_ekstrimitas_bawah"`
	PemeriksaanAnus             *string    `gorm:"column:pemeriksaan_anus" json:"pemeriksaan_anus"`
	Nip                         string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianKorbanKekerasan) TableName() string {
	return "penilaian_korban_kekerasan"
}
