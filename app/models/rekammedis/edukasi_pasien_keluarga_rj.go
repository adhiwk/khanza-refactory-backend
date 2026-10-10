package rekammedis

import "time"

// EdukasiPasienKeluargaRj tabel `edukasi_pasien_keluarga_rj` (edukasi pasien keluarga rawat jalan, RMEdukasiPasienKeluargaRawatJalan).
type EdukasiPasienKeluargaRj struct {
	NoRawat                                 string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                                 *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip                                     string     `gorm:"column:nip" json:"nip"`
	Bicara                                  *string    `gorm:"column:bicara" json:"bicara"`
	KeteranganBicara                        *string    `gorm:"column:keterangan_bicara" json:"keterangan_bicara"`
	BahasaSehari                            *string    `gorm:"column:bahasa_sehari" json:"bahasa_sehari"`
	PerluPenerjemah                         *string    `gorm:"column:perlu_penerjemah" json:"perlu_penerjemah"`
	KeteranganPenerjemah                    *string    `gorm:"column:keterangan_penerjemah" json:"keterangan_penerjemah"`
	BahasaIsyarat                           *string    `gorm:"column:bahasa_isyarat" json:"bahasa_isyarat"`
	CaraBelajar                             *string    `gorm:"column:cara_belajar" json:"cara_belajar"`
	HambatanBelajar                         *string    `gorm:"column:hambatan_belajar" json:"hambatan_belajar"`
	KeteranganHambatanBelajar               *string    `gorm:"column:keterangan_hambatan_belajar" json:"keterangan_hambatan_belajar"`
	KemampuanBelajar                        *string    `gorm:"column:kemampuan_belajar" json:"kemampuan_belajar"`
	KeteranganKemampuanBelajar              *string    `gorm:"column:keterangan_kemampuan_belajar" json:"keterangan_kemampuan_belajar"`
	PenyakitnyaMerupakan                    string     `gorm:"column:penyakitnya_merupakan" json:"penyakitnya_merupakan"`
	KeteranganPenyakitnyaMerupakan          string     `gorm:"column:keterangan_penyakitnya_merupakan" json:"keterangan_penyakitnya_merupakan"`
	KeputusanMemilihLayanan                 string     `gorm:"column:keputusan_memilih_layanan" json:"keputusan_memilih_layanan"`
	KeteranganKeputusanMemilihLayanan       string     `gorm:"column:keterangan_keputusan_memilih_layanan" json:"keterangan_keputusan_memilih_layanan"`
	KeyakinanTerhadapTerapi                 string     `gorm:"column:keyakinan_terhadap_terapi" json:"keyakinan_terhadap_terapi"`
	KeteranganKeyakinanTerhadapTerapi       string     `gorm:"column:keterangan_keyakinan_terhadap_terapi" json:"keterangan_keyakinan_terhadap_terapi"`
	AspekKeyakinanDipertimbangkan           string     `gorm:"column:aspek_keyakinan_dipertimbangkan" json:"aspek_keyakinan_dipertimbangkan"`
	KeteranganAspekKeyakinanDipertimbangkan string     `gorm:"column:keterangan_aspek_keyakinan_dipertimbangkan" json:"keterangan_aspek_keyakinan_dipertimbangkan"`
	KesediaanMenerimaInformasi              string     `gorm:"column:kesediaan_menerima_informasi" json:"kesediaan_menerima_informasi"`
	TopikEdukasiPenyakit                    string     `gorm:"column:topik_edukasi_penyakit" json:"topik_edukasi_penyakit"`
	TopikEdukasiRencanaTindakan             string     `gorm:"column:topik_edukasi_rencana_tindakan" json:"topik_edukasi_rencana_tindakan"`
	TopikEdukasiPengobatan                  string     `gorm:"column:topik_edukasi_pengobatan" json:"topik_edukasi_pengobatan"`
	TopikEdukasiHasilLayanan                string     `gorm:"column:topik_edukasi_hasil_layanan" json:"topik_edukasi_hasil_layanan"`
}

func (EdukasiPasienKeluargaRj) TableName() string {
	return "edukasi_pasien_keluarga_rj"
}
