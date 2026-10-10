package rekammedis

import "time"

// TransferPasienAntarRuang tabel `transfer_pasien_antar_ruang` (transfer pasien antar ruang, RMTransferPasienAntarRuang).
type TransferPasienAntarRuang struct {
	NoRawat                           string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TanggalMasuk                      *time.Time `gorm:"column:tanggal_masuk;primaryKey;autoIncrement:false" json:"tanggal_masuk"`
	TanggalPindah                     *time.Time `gorm:"column:tanggal_pindah" json:"tanggal_pindah"`
	AsalRuang                         *string    `gorm:"column:asal_ruang" json:"asal_ruang"`
	RuangSelanjutnya                  *string    `gorm:"column:ruang_selanjutnya" json:"ruang_selanjutnya"`
	DiagnosaUtama                     *string    `gorm:"column:diagnosa_utama" json:"diagnosa_utama"`
	DiagnosaSekunder                  *string    `gorm:"column:diagnosa_sekunder" json:"diagnosa_sekunder"`
	IndikasiPindahRuang               *string    `gorm:"column:indikasi_pindah_ruang" json:"indikasi_pindah_ruang"`
	KeteranganIndikasiPindahRuang     *string    `gorm:"column:keterangan_indikasi_pindah_ruang" json:"keterangan_indikasi_pindah_ruang"`
	ProsedurYangSudahDilakukan        *string    `gorm:"column:prosedur_yang_sudah_dilakukan" json:"prosedur_yang_sudah_dilakukan"`
	ObatYangTelahDiberikan            *string    `gorm:"column:obat_yang_telah_diberikan" json:"obat_yang_telah_diberikan"`
	MetodePemindahanPasien            *string    `gorm:"column:metode_pemindahan_pasien" json:"metode_pemindahan_pasien"`
	PeralatanYangMenyertai            *string    `gorm:"column:peralatan_yang_menyertai" json:"peralatan_yang_menyertai"`
	KeteranganPeralatanYangMenyertai  *string    `gorm:"column:keterangan_peralatan_yang_menyertai" json:"keterangan_peralatan_yang_menyertai"`
	PemeriksaanPenunjangYangDilakukan *string    `gorm:"column:pemeriksaan_penunjang_yang_dilakukan" json:"pemeriksaan_penunjang_yang_dilakukan"`
	PasienKeluargaMenyetujui          *string    `gorm:"column:pasien_keluarga_menyetujui" json:"pasien_keluarga_menyetujui"`
	NamaMenyetujui                    *string    `gorm:"column:nama_menyetujui" json:"nama_menyetujui"`
	HubunganMenyetujui                *string    `gorm:"column:hubungan_menyetujui" json:"hubungan_menyetujui"`
	KeluhanUtamaSebelumTransfer       *string    `gorm:"column:keluhan_utama_sebelum_transfer" json:"keluhan_utama_sebelum_transfer"`
	KeadaanUmumSebelumTransfer        *string    `gorm:"column:keadaan_umum_sebelum_transfer" json:"keadaan_umum_sebelum_transfer"`
	TdSebelumTransfer                 *string    `gorm:"column:td_sebelum_transfer" json:"td_sebelum_transfer"`
	NadiSebelumTransfer               *string    `gorm:"column:nadi_sebelum_transfer" json:"nadi_sebelum_transfer"`
	RrSebelumTransfer                 *string    `gorm:"column:rr_sebelum_transfer" json:"rr_sebelum_transfer"`
	SuhuSebelumTransfer               *string    `gorm:"column:suhu_sebelum_transfer" json:"suhu_sebelum_transfer"`
	KeluhanUtamaSesudahTransfer       *string    `gorm:"column:keluhan_utama_sesudah_transfer" json:"keluhan_utama_sesudah_transfer"`
	KeadaanUmumSesudahTransfer        *string    `gorm:"column:keadaan_umum_sesudah_transfer" json:"keadaan_umum_sesudah_transfer"`
	TdSesudahTransfer                 *string    `gorm:"column:td_sesudah_transfer" json:"td_sesudah_transfer"`
	NadiSesudahTransfer               *string    `gorm:"column:nadi_sesudah_transfer" json:"nadi_sesudah_transfer"`
	RrSesudahTransfer                 *string    `gorm:"column:rr_sesudah_transfer" json:"rr_sesudah_transfer"`
	SuhuSesudahTransfer               *string    `gorm:"column:suhu_sesudah_transfer" json:"suhu_sesudah_transfer"`
	NipMenyerahkan                    string     `gorm:"column:nip_menyerahkan" json:"nip_menyerahkan"`
	NipMenerima                       string     `gorm:"column:nip_menerima" json:"nip_menerima"`
}

func (TransferPasienAntarRuang) TableName() string {
	return "transfer_pasien_antar_ruang"
}
