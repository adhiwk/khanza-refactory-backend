package rekammedis

import "time"

// PerencanaanPemulangan tabel `perencanaan_pemulangan` (perencanaan pemulangan, RMPerencanaanPemulangan).
type PerencanaanPemulangan struct {
	NoRawat                                   string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	RencanaPulang                             *time.Time `gorm:"column:rencana_pulang" json:"rencana_pulang"`
	AlasanMasuk                               *string    `gorm:"column:alasan_masuk" json:"alasan_masuk"`
	DiagnosaMedis                             *string    `gorm:"column:diagnosa_medis" json:"diagnosa_medis"`
	PengaruhRiPasienDanKeluarga               *string    `gorm:"column:pengaruh_ri_pasien_dan_keluarga" json:"pengaruh_ri_pasien_dan_keluarga"`
	KeteranganPengaruhRiPasienDanKeluarga     *string    `gorm:"column:keterangan_pengaruh_ri_pasien_dan_keluarga" json:"keterangan_pengaruh_ri_pasien_dan_keluarga"`
	PengaruhRiPekerjaanSekolah                string     `gorm:"column:pengaruh_ri_pekerjaan_sekolah" json:"pengaruh_ri_pekerjaan_sekolah"`
	KeteranganPengaruhRiPekerjaanSekolah      string     `gorm:"column:keterangan_pengaruh_ri_pekerjaan_sekolah" json:"keterangan_pengaruh_ri_pekerjaan_sekolah"`
	PengaruhRiKeuangan                        string     `gorm:"column:pengaruh_ri_keuangan" json:"pengaruh_ri_keuangan"`
	KeteranganPengaruhRiKeuangan              string     `gorm:"column:keterangan_pengaruh_ri_keuangan" json:"keterangan_pengaruh_ri_keuangan"`
	AntisipasiMasalahSaatPulang               string     `gorm:"column:antisipasi_masalah_saat_pulang" json:"antisipasi_masalah_saat_pulang"`
	KeteranganAntisipasiMasalahSaatPulang     string     `gorm:"column:keterangan_antisipasi_masalah_saat_pulang" json:"keterangan_antisipasi_masalah_saat_pulang"`
	BantuanDiperlukanDalam                    string     `gorm:"column:bantuan_diperlukan_dalam" json:"bantuan_diperlukan_dalam"`
	KeteranganBantuanDiperlukanDalam          string     `gorm:"column:keterangan_bantuan_diperlukan_dalam" json:"keterangan_bantuan_diperlukan_dalam"`
	AdakahYangMembantuKeperluan               string     `gorm:"column:adakah_yang_membantu_keperluan" json:"adakah_yang_membantu_keperluan"`
	KeteranganAdakahYangMembantuKeperluan     string     `gorm:"column:keterangan_adakah_yang_membantu_keperluan" json:"keterangan_adakah_yang_membantu_keperluan"`
	PasienTinggalSendiri                      string     `gorm:"column:pasien_tinggal_sendiri" json:"pasien_tinggal_sendiri"`
	KeteranganPasienTinggalSendiri            string     `gorm:"column:keterangan_pasien_tinggal_sendiri" json:"keterangan_pasien_tinggal_sendiri"`
	PasienMenggunakanPeralatanMedis           string     `gorm:"column:pasien_menggunakan_peralatan_medis" json:"pasien_menggunakan_peralatan_medis"`
	KeteranganPasienMenggunakanPeralatanMedis string     `gorm:"column:keterangan_pasien_menggunakan_peralatan_medis" json:"keterangan_pasien_menggunakan_peralatan_medis"`
	PasienMemerlukanAlatBantu                 string     `gorm:"column:pasien_memerlukan_alat_bantu" json:"pasien_memerlukan_alat_bantu"`
	KeteranganPasienMemerlukanAlatBantu       string     `gorm:"column:keterangan_pasien_memerlukan_alat_bantu" json:"keterangan_pasien_memerlukan_alat_bantu"`
	MemerlukanPerawatanKhusus                 string     `gorm:"column:memerlukan_perawatan_khusus" json:"memerlukan_perawatan_khusus"`
	KeteranganMemerlukanPerawatanKhusus       string     `gorm:"column:keterangan_memerlukan_perawatan_khusus" json:"keterangan_memerlukan_perawatan_khusus"`
	BermasalahMemenuhiKebutuhan               string     `gorm:"column:bermasalah_memenuhi_kebutuhan" json:"bermasalah_memenuhi_kebutuhan"`
	KeteranganBermasalahMemenuhiKebutuhan     string     `gorm:"column:keterangan_bermasalah_memenuhi_kebutuhan" json:"keterangan_bermasalah_memenuhi_kebutuhan"`
	MemilikiNyeriKronis                       string     `gorm:"column:memiliki_nyeri_kronis" json:"memiliki_nyeri_kronis"`
	KeteranganMemilikiNyeriKronis             string     `gorm:"column:keterangan_memiliki_nyeri_kronis" json:"keterangan_memiliki_nyeri_kronis"`
	MemerlukanEdukasiKesehatan                string     `gorm:"column:memerlukan_edukasi_kesehatan" json:"memerlukan_edukasi_kesehatan"`
	KeteranganMemerlukanEdukasiKesehatan      string     `gorm:"column:keterangan_memerlukan_edukasi_kesehatan" json:"keterangan_memerlukan_edukasi_kesehatan"`
	MemerlukanKeterampilkanKhusus             string     `gorm:"column:memerlukan_keterampilkan_khusus" json:"memerlukan_keterampilkan_khusus"`
	KeteranganMemerlukanKeterampilkanKhusus   string     `gorm:"column:keterangan_memerlukan_keterampilkan_khusus" json:"keterangan_memerlukan_keterampilkan_khusus"`
	NamaPasienKeluarga                        string     `gorm:"column:nama_pasien_keluarga" json:"nama_pasien_keluarga"`
	Nip                                       string     `gorm:"column:nip" json:"nip"`
}

func (PerencanaanPemulangan) TableName() string {
	return "perencanaan_pemulangan"
}
