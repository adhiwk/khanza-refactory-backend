package rekammedis

import "time"

// SkriningPerilakuMerokokSekolahRemaja tabel `skrining_perilaku_merokok_sekolah_remaja` (skrining merokok usia sekolah remaja, RMSkriningMerokokUsiaSekolahRemaja).
type SkriningPerilakuMerokokSekolahRemaja struct {
	NoRawat                                         string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                                         *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdSekolah                                       *string    `gorm:"column:kd_sekolah" json:"kd_sekolah"`
	Kelas                                           *string    `gorm:"column:kelas" json:"kelas"`
	ApakahAndaMerokok                               string     `gorm:"column:apakah_anda_merokok" json:"apakah_anda_merokok"`
	JumlahBatangRokok                               string     `gorm:"column:jumlah_batang_rokok" json:"jumlah_batang_rokok"`
	JumlahBatangRokokHariminggu                     string     `gorm:"column:jumlah_batang_rokok_hariminggu" json:"jumlah_batang_rokok_hariminggu"`
	JenisRokokYangDigunakan                         string     `gorm:"column:jenis_rokok_yang_digunakan" json:"jenis_rokok_yang_digunakan"`
	JenisRokokYangDigunakanKeterangan               string     `gorm:"column:jenis_rokok_yang_digunakan_keterangan" json:"jenis_rokok_yang_digunakan_keterangan"`
	UsiaMulaiMerokok                                string     `gorm:"column:usia_mulai_merokok" json:"usia_mulai_merokok"`
	AlasanMulaiMerokok                              string     `gorm:"column:alasan_mulai_merokok" json:"alasan_mulai_merokok"`
	AlasanMulaiMerokokKeterangan                    string     `gorm:"column:alasan_mulai_merokok_keterangan" json:"alasan_mulai_merokok_keterangan"`
	SudahBerapaLamaMerokok                          string     `gorm:"column:sudah_berapa_lama_merokok" json:"sudah_berapa_lama_merokok"`
	BagaimanaBiasanyaMendapatkanRokok               string     `gorm:"column:bagaimana_biasanya_mendapatkan_rokok" json:"bagaimana_biasanya_mendapatkan_rokok"`
	BagaimanaBiasanyaMendapatkanRokokKeterangan     string     `gorm:"column:bagaimana_biasanya_mendapatkan_rokok_keterangan" json:"bagaimana_biasanya_mendapatkan_rokok_keterangan"`
	KeinginanBerhentiMerokok                        string     `gorm:"column:keinginan_berhenti_merokok" json:"keinginan_berhenti_merokok"`
	AlasanUtamaBerhentiMerokok                      string     `gorm:"column:alasan_utama_berhenti_merokok" json:"alasan_utama_berhenti_merokok"`
	AlasanUtamaBerhentiMerokokKeterangan            string     `gorm:"column:alasan_utama_berhenti_merokok_keterangan" json:"alasan_utama_berhenti_merokok_keterangan"`
	TahuDampakKesehatanMerokok                      string     `gorm:"column:tahu_dampak_kesehatan_merokok" json:"tahu_dampak_kesehatan_merokok"`
	DampakKesehatanDariMerokokYangDiketahui         string     `gorm:"column:dampak_kesehatan_dari_merokok_yang_diketahui" json:"dampak_kesehatan_dari_merokok_yang_diketahui"`
	TahuMerokokPintuMasukNarkoba                    string     `gorm:"column:tahu_merokok_pintu_masuk_narkoba" json:"tahu_merokok_pintu_masuk_narkoba"`
	MelihatOrangMerokokDiSekolah                    string     `gorm:"column:melihat_orang_merokok_di_sekolah" json:"melihat_orang_merokok_di_sekolah"`
	OrangYangPalingSeringMerokokDisekolah           string     `gorm:"column:orang_yang_paling_sering_merokok_disekolah" json:"orang_yang_paling_sering_merokok_disekolah"`
	OrangYangPalingSeringMerokokDisekolahKeterangan string     `gorm:"column:orang_yang_paling_sering_merokok_disekolah_keterangan" json:"orang_yang_paling_sering_merokok_disekolah_keterangan"`
	AdaAnggotaKeluargaDiRumahYangMerokok            string     `gorm:"column:ada_anggota_keluarga_di_rumah_yang_merokok" json:"ada_anggota_keluarga_di_rumah_yang_merokok"`
	TemanDekatBanyakyangMerokok                     string     `gorm:"column:teman_dekat_banyakyang_merokok" json:"teman_dekat_banyakyang_merokok"`
	DilakukanPemeriksaanKadarCoPernapasan           string     `gorm:"column:dilakukan_pemeriksaan_kadar_co_pernapasan" json:"dilakukan_pemeriksaan_kadar_co_pernapasan"`
	HasilPemeriksaanCoPernapasan                    string     `gorm:"column:hasil_pemeriksaan_co_pernapasan" json:"hasil_pemeriksaan_co_pernapasan"`
	Nip                                             string     `gorm:"column:nip" json:"nip"`
}

func (SkriningPerilakuMerokokSekolahRemaja) TableName() string {
	return "skrining_perilaku_merokok_sekolah_remaja"
}
