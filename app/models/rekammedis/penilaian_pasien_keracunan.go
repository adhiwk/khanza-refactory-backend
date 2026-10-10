package rekammedis

import "time"

// PenilaianPasienKeracunan tabel `penilaian_pasien_keracunan` (penilaian pasien keracunan, RMPenilaianPasienKeracunan).
type PenilaianPasienKeracunan struct {
	NoRawat                  string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                  *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter                 string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Anamnesis                string     `gorm:"column:anamnesis" json:"anamnesis"`
	Hubungan                 string     `gorm:"column:hubungan" json:"hubungan"`
	TempatKejadian           *string    `gorm:"column:tempat_kejadian" json:"tempat_kejadian"`
	KeteranganTempatKejadian *string    `gorm:"column:keterangan_tempat_kejadian" json:"keterangan_tempat_kejadian"`
	Keluhan                  *string    `gorm:"column:keluhan" json:"keluhan"`
	RiwayatPenyakitSekarang  *string    `gorm:"column:riwayat_penyakit_sekarang" json:"riwayat_penyakit_sekarang"`
	Hamil                    *string    `gorm:"column:hamil" json:"hamil"`
	Menyusui                 *string    `gorm:"column:menyusui" json:"menyusui"`
	Penyebab                 *string    `gorm:"column:penyebab" json:"penyebab"`
	NamaBahan                *string    `gorm:"column:nama_bahan" json:"nama_bahan"`
	JumlahBahan              *string    `gorm:"column:jumlah_bahan" json:"jumlah_bahan"`
	TipePemaparan            *string    `gorm:"column:tipe_pemaparan" json:"tipe_pemaparan"`
	KeteranganTipePemaparan  *string    `gorm:"column:keterangan_tipe_pemaparan" json:"keterangan_tipe_pemaparan"`
	TipeKejadian             *string    `gorm:"column:tipe_kejadian" json:"tipe_kejadian"`
	BauBahan                 *string    `gorm:"column:bau_bahan" json:"bau_bahan"`
	KeteranganBauBahan       *string    `gorm:"column:keterangan_bau_bahan" json:"keterangan_bau_bahan"`
	Pupil                    *string    `gorm:"column:pupil" json:"pupil"`
	KeteranganPupil          *string    `gorm:"column:keterangan_pupil" json:"keterangan_pupil"`
	Kesadaran                string     `gorm:"column:kesadaran" json:"kesadaran"`
	Td                       *string    `gorm:"column:td" json:"td"`
	Nadi                     *string    `gorm:"column:nadi" json:"nadi"`
	Rr                       string     `gorm:"column:rr" json:"rr"`
	Suhu                     *string    `gorm:"column:suhu" json:"suhu"`
	Spo                      string     `gorm:"column:spo" json:"spo"`
	Urine                    *string    `gorm:"column:urine" json:"urine"`
	PengobatanSebelumIgd     *string    `gorm:"column:pengobatan_sebelum_igd" json:"pengobatan_sebelum_igd"`
	Diagnosis                *string    `gorm:"column:diagnosis" json:"diagnosis"`
	PemeriksaanPenunjang     *string    `gorm:"column:pemeriksaan_penunjang" json:"pemeriksaan_penunjang"`
	PenatalaksanaanDiberikan *string    `gorm:"column:penatalaksanaan_diberikan" json:"penatalaksanaan_diberikan"`
	TindakLanjut             *string    `gorm:"column:tindak_lanjut" json:"tindak_lanjut"`
}

func (PenilaianPasienKeracunan) TableName() string {
	return "penilaian_pasien_keracunan"
}
