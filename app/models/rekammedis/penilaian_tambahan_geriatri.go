package rekammedis

import "time"

// PenilaianTambahanGeriatri tabel `penilaian_tambahan_geriatri` (penilaian tambahan geriatri, RMPenilaianTambahanGeriatri).
type PenilaianTambahanGeriatri struct {
	NoRawat                             string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                             *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nik                                 string     `gorm:"column:nik" json:"nik"`
	AsalMasuk                           *string    `gorm:"column:asal_masuk" json:"asal_masuk"`
	KondisiMasuk                        *string    `gorm:"column:kondisi_masuk" json:"kondisi_masuk"`
	KeteranganKondisiMasuk              *string    `gorm:"column:keterangan_kondisi_masuk" json:"keterangan_kondisi_masuk"`
	Anamnesis                           string     `gorm:"column:anamnesis" json:"anamnesis"`
	DiagnosaMedis                       *string    `gorm:"column:diagnosa_medis" json:"diagnosa_medis"`
	RiwayatImmunoTelinga                *string    `gorm:"column:riwayat_immuno_telinga" json:"riwayat_immuno_telinga"`
	RiwayatImmunoSinus                  *string    `gorm:"column:riwayat_immuno_sinus" json:"riwayat_immuno_sinus"`
	RiwayatImmunoAntibiotik             *string    `gorm:"column:riwayat_immuno_antibiotik" json:"riwayat_immuno_antibiotik"`
	RiwayatImmunoPneumonia              *string    `gorm:"column:riwayat_immuno_pneumonia" json:"riwayat_immuno_pneumonia"`
	RiwayatImmunoAbses                  *string    `gorm:"column:riwayat_immuno_abses" json:"riwayat_immuno_abses"`
	RiwayatImmunoSariawan               *string    `gorm:"column:riwayat_immuno_sariawan" json:"riwayat_immuno_sariawan"`
	RiwayatImmunoMemerlukanAntibiotik   *string    `gorm:"column:riwayat_immuno_memerlukan_antibiotik" json:"riwayat_immuno_memerlukan_antibiotik"`
	RiwayatImmunoInfeksiDalam           *string    `gorm:"column:riwayat_immuno_infeksi_dalam" json:"riwayat_immuno_infeksi_dalam"`
	RiwayatImmunoImmunodefisiensiPrimer *string    `gorm:"column:riwayat_immuno_immunodefisiensi_primer" json:"riwayat_immuno_immunodefisiensi_primer"`
	RiwayatImmunoJenisKangker           *string    `gorm:"column:riwayat_immuno_jenis_kangker" json:"riwayat_immuno_jenis_kangker"`
	RiwayatImmunoInfeksiOportunistik    *string    `gorm:"column:riwayat_immuno_infeksi_oportunistik" json:"riwayat_immuno_infeksi_oportunistik"`
	PolaAktifitasTidur                  *string    `gorm:"column:pola_aktifitas_tidur" json:"pola_aktifitas_tidur"`
	KeteranganPolaAktifitasTidur        *string    `gorm:"column:keterangan_pola_aktifitas_tidur" json:"keterangan_pola_aktifitas_tidur"`
	PolaAktifitasObatTidur              *string    `gorm:"column:pola_aktifitas_obat_tidur" json:"pola_aktifitas_obat_tidur"`
	KeteranganPolaAktifitasObatTidur    *string    `gorm:"column:keterangan_pola_aktifitas_obat_tidur" json:"keterangan_pola_aktifitas_obat_tidur"`
	PolaAktifitasOlahraga               *string    `gorm:"column:pola_aktifitas_olahraga" json:"pola_aktifitas_olahraga"`
	KeteranganPolaAktifitasOlahraga     *string    `gorm:"column:keterangan_pola_aktifitas_olahraga" json:"keterangan_pola_aktifitas_olahraga"`
	KualitasHidupMobilitas              *string    `gorm:"column:kualitas_hidup_mobilitas" json:"kualitas_hidup_mobilitas"`
	KualitasHidupPerawatanDiri          *string    `gorm:"column:kualitas_hidup_perawatan_diri" json:"kualitas_hidup_perawatan_diri"`
	KualitasHidupAktifitasSeharihari    *string    `gorm:"column:kualitas_hidup_aktifitas_seharihari" json:"kualitas_hidup_aktifitas_seharihari"`
	KualitasHidupRasaNyeri              *string    `gorm:"column:kualitas_hidup_rasa_nyeri" json:"kualitas_hidup_rasa_nyeri"`
	SkalaNyeri                          *string    `gorm:"column:skala_nyeri" json:"skala_nyeri"`
}

func (PenilaianTambahanGeriatri) TableName() string {
	return "penilaian_tambahan_geriatri"
}
