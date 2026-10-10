package rekammedis

import "time"

// ChecklistKriteriaKeluarIsolasi tabel `checklist_kriteria_keluar_isolasi` (checklist kriteria keluar isolasi, RMChecklistKriteriaKeluarIsolasi).
type ChecklistKriteriaKeluarIsolasi struct {
	NoRawat                   string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                   *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	GejalaMembaik             string     `gorm:"column:gejala_membaik" json:"gejala_membaik"`
	TidakAdaIndikasiTransmisi string     `gorm:"column:tidak_ada_indikasi_transmisi" json:"tidak_ada_indikasi_transmisi"`
	HasilPenunjangMemenuhi    string     `gorm:"column:hasil_penunjang_memenuhi" json:"hasil_penunjang_memenuhi"`
	KriteriaPedomanTerpenuhi  string     `gorm:"column:kriteria_pedoman_terpenuhi" json:"kriteria_pedoman_terpenuhi"`
	PersetujuanDpjp           string     `gorm:"column:persetujuan_dpjp" json:"persetujuan_dpjp"`
	Keputusan                 string     `gorm:"column:keputusan" json:"keputusan"`
	Alasan                    *string    `gorm:"column:alasan" json:"alasan"`
	Nik                       *string    `gorm:"column:nik" json:"nik"`
}

func (ChecklistKriteriaKeluarIsolasi) TableName() string {
	return "checklist_kriteria_keluar_isolasi"
}
