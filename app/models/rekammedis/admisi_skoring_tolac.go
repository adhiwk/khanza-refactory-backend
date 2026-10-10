package rekammedis

import "time"

// AdmisiSkoringTolac tabel `admisi_skoring_tolac` (admisi skoring TOLAC, RMAdmisiSkoringTOLAC).
type AdmisiSkoringTolac struct {
	NoRawat                  string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                  *time.Time `gorm:"column:tanggal" json:"tanggal"`
	HisFrekuensi             *string    `gorm:"column:his_frekuensi" json:"his_frekuensi"`
	HisDurasiDetik           *string    `gorm:"column:his_durasi_detik" json:"his_durasi_detik"`
	Djj                      *string    `gorm:"column:djj" json:"djj"`
	PembukaanCm              *string    `gorm:"column:pembukaan_cm" json:"pembukaan_cm"`
	PendataranPersen         *string    `gorm:"column:pendataran_persen" json:"pendataran_persen"`
	PenurunanKepala          *string    `gorm:"column:penurunan_kepala" json:"penurunan_kepala"`
	PilihanUsia              *string    `gorm:"column:pilihan_usia" json:"pilihan_usia"`
	SkorUsia                 *string    `gorm:"column:skor_usia" json:"skor_usia"`
	PilihanRiwayatPervaginam *string    `gorm:"column:pilihan_riwayat_pervaginam" json:"pilihan_riwayat_pervaginam"`
	SkorRiwayatPervaginam    *string    `gorm:"column:skor_riwayat_pervaginam" json:"skor_riwayat_pervaginam"`
	PilihanIndikasiSc        *string    `gorm:"column:pilihan_indikasi_sc" json:"pilihan_indikasi_sc"`
	SkorIndikasiSc           *string    `gorm:"column:skor_indikasi_sc" json:"skor_indikasi_sc"`
	PilihanPendataran        *string    `gorm:"column:pilihan_pendataran" json:"pilihan_pendataran"`
	SkorPendataran           *string    `gorm:"column:skor_pendataran" json:"skor_pendataran"`
	PilihanPembukaan         *string    `gorm:"column:pilihan_pembukaan" json:"pilihan_pembukaan"`
	SkorPembukaan            *string    `gorm:"column:skor_pembukaan" json:"skor_pembukaan"`
	TotalSkor                *string    `gorm:"column:total_skor" json:"total_skor"`
	PeluangVbac              string     `gorm:"column:peluang_vbac" json:"peluang_vbac"`
	Keputusan                *string    `gorm:"column:keputusan" json:"keputusan"`
	Keterangan               *string    `gorm:"column:keterangan" json:"keterangan"`
	KdDokter                 *string    `gorm:"column:kd_dokter" json:"kd_dokter"`
}

func (AdmisiSkoringTolac) TableName() string {
	return "admisi_skoring_tolac"
}
