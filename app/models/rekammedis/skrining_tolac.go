package rekammedis

import "time"

// SkriningTolac tabel `skrining_tolac` (skrining TOLAC, RMSkriningTOLAC).
type SkriningTolac struct {
	NoRawat                   string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                   *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Gpa                       *string    `gorm:"column:gpa" json:"gpa"`
	Diagnosa                  *string    `gorm:"column:diagnosa" json:"diagnosa"`
	JumlahSc                  *string    `gorm:"column:jumlah_sc" json:"jumlah_sc"`
	TahunSc                   *string    `gorm:"column:tahun_sc" json:"tahun_sc"`
	IndikasiSc                *string    `gorm:"column:indikasi_sc" json:"indikasi_sc"`
	JenisInsisi               *string    `gorm:"column:jenis_insisi" json:"jenis_insisi"`
	RiwayatPervaginam         *string    `gorm:"column:riwayat_pervaginam" json:"riwayat_pervaginam"`
	TbjGram                   *string    `gorm:"column:tbj_gram" json:"tbj_gram"`
	PresentasiJanin           *string    `gorm:"column:presentasi_janin" json:"presentasi_janin"`
	InklusiRiwayatSc          *string    `gorm:"column:inklusi_riwayat_sc" json:"inklusi_riwayat_sc"`
	InklusiPanggulAdekuat     *string    `gorm:"column:inklusi_panggul_adekuat" json:"inklusi_panggul_adekuat"`
	InklusiJaninTunggalKepala *string    `gorm:"column:inklusi_janin_tunggal_kepala" json:"inklusi_janin_tunggal_kepala"`
	InklusiTbjSesuai          *string    `gorm:"column:inklusi_tbj_sesuai" json:"inklusi_tbj_sesuai"`
	EksklusiScKlasikRuptur    *string    `gorm:"column:eksklusi_sc_klasik_ruptur" json:"eksklusi_sc_klasik_ruptur"`
	EksklusiSc2x              *string    `gorm:"column:eksklusi_sc_2x" json:"eksklusi_sc_2x"`
	EksklusiPlasentaPrevia    *string    `gorm:"column:eksklusi_plasenta_previa" json:"eksklusi_plasenta_previa"`
	Kesimpulan                *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
	EdukasiDiberikan          *string    `gorm:"column:edukasi_diberikan" json:"edukasi_diberikan"`
	Keterangan                string     `gorm:"column:keterangan" json:"keterangan"`
	KdDokter                  *string    `gorm:"column:kd_dokter" json:"kd_dokter"`
}

func (SkriningTolac) TableName() string {
	return "skrining_tolac"
}
