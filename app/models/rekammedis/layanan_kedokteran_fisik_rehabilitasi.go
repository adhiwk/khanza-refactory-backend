package rekammedis

import "time"

// LayananKedokteranFisikRehabilitasi tabel `layanan_kedokteran_fisik_rehabilitasi` (layanan kedokteran fisik rehabilitasi, RMLayananKedokteranFisikRehabilitasi).
type LayananKedokteranFisikRehabilitasi struct {
	NoRawat                       string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                       *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter                      string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Pendamping                    *string    `gorm:"column:pendamping" json:"pendamping"`
	KeteranganPendamping          *string    `gorm:"column:keterangan_pendamping" json:"keterangan_pendamping"`
	Anamnesa                      *string    `gorm:"column:anamnesa" json:"anamnesa"`
	PemeriksaanFisik              *string    `gorm:"column:pemeriksaan_fisik" json:"pemeriksaan_fisik"`
	DiagnosaMedis                 *string    `gorm:"column:diagnosa_medis" json:"diagnosa_medis"`
	DiagnosaFungsi                *string    `gorm:"column:diagnosa_fungsi" json:"diagnosa_fungsi"`
	Tatalaksana                   *string    `gorm:"column:tatalaksana" json:"tatalaksana"`
	Anjuran                       string     `gorm:"column:anjuran" json:"anjuran"`
	Evaluasi                      string     `gorm:"column:evaluasi" json:"evaluasi"`
	SuspekPenyakitKerja           *string    `gorm:"column:suspek_penyakit_kerja" json:"suspek_penyakit_kerja"`
	KeteranganSuspekPenyakitKerja *string    `gorm:"column:keterangan_suspek_penyakit_kerja" json:"keterangan_suspek_penyakit_kerja"`
	StatusProgram                 string     `gorm:"column:status_program" json:"status_program"`
}

func (LayananKedokteranFisikRehabilitasi) TableName() string {
	return "layanan_kedokteran_fisik_rehabilitasi"
}
