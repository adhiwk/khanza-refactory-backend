package rekammedis

import "time"

// SkriningIndraPendengaran tabel `skrining_indra_pendengaran` (skrining indra pendengaran, RMSkriningIndraPendengaran).
type SkriningIndraPendengaran struct {
	NoRawat                          string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                          *time.Time `gorm:"column:tanggal" json:"tanggal"`
	CurigaTuliTelingaKiri            *string    `gorm:"column:curiga_tuli_telinga_kiri" json:"curiga_tuli_telinga_kiri"`
	CurigaTuliTelingaKanan           *string    `gorm:"column:curiga_tuli_telinga_kanan" json:"curiga_tuli_telinga_kanan"`
	CurigaTuliTelingaRujuk           *string    `gorm:"column:curiga_tuli_telinga_rujuk" json:"curiga_tuli_telinga_rujuk"`
	PenurunanPendengaranTelingaKiri  *string    `gorm:"column:penurunan_pendengaran_telinga_kiri" json:"penurunan_pendengaran_telinga_kiri"`
	PenurunanPendengaranTelingaKanan *string    `gorm:"column:penurunan_pendengaran_telinga_kanan" json:"penurunan_pendengaran_telinga_kanan"`
	MendengarBisikanTelingaKiri      *string    `gorm:"column:mendengar_bisikan_telinga_kiri" json:"mendengar_bisikan_telinga_kiri"`
	MendengarBisikanTelingaKanan     *string    `gorm:"column:mendengar_bisikan_telinga_kanan" json:"mendengar_bisikan_telinga_kanan"`
	CongekTelingaKiri                *string    `gorm:"column:congek_telinga_kiri" json:"congek_telinga_kiri"`
	CongekTelingaKanan               *string    `gorm:"column:congek_telinga_kanan" json:"congek_telinga_kanan"`
	CongekTelingaRujuk               *string    `gorm:"column:congek_telinga_rujuk" json:"congek_telinga_rujuk"`
	SumbatanSerumenTelingaKiri       *string    `gorm:"column:sumbatan_serumen_telinga_kiri" json:"sumbatan_serumen_telinga_kiri"`
	SumbatanSerumenTelingaKanan      *string    `gorm:"column:sumbatan_serumen_telinga_kanan" json:"sumbatan_serumen_telinga_kanan"`
	SumbatanSerumenTelingaRujuk      *string    `gorm:"column:sumbatan_serumen_telinga_rujuk" json:"sumbatan_serumen_telinga_rujuk"`
	HasilSkrining                    string     `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Keterangan                       *string    `gorm:"column:keterangan" json:"keterangan"`
	Nip                              string     `gorm:"column:nip" json:"nip"`
}

func (SkriningIndraPendengaran) TableName() string {
	return "skrining_indra_pendengaran"
}
