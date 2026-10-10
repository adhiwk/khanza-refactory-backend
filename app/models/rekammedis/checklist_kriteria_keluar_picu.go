package rekammedis

import "time"

// ChecklistKriteriaKeluarPicu tabel `checklist_kriteria_keluar_picu` (checklist kriteria keluar PICU, RMChecklistKriteriaKeluarPICU).
type ChecklistKriteriaKeluarPicu struct {
	NoRawat             string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal             *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Kondisiklinis1      string     `gorm:"column:kondisiklinis1" json:"kondisiklinis1"`
	Kondisiklinis2      string     `gorm:"column:kondisiklinis2" json:"kondisiklinis2"`
	Kondisiklinis3      string     `gorm:"column:kondisiklinis3" json:"kondisiklinis3"`
	Kondisiklinis4      string     `gorm:"column:kondisiklinis4" json:"kondisiklinis4"`
	Kondisiklinis5      string     `gorm:"column:kondisiklinis5" json:"kondisiklinis5"`
	Kondisiklinis6      string     `gorm:"column:kondisiklinis6" json:"kondisiklinis6"`
	Kebutuhanperawatan1 string     `gorm:"column:kebutuhanperawatan1" json:"kebutuhanperawatan1"`
	Kebutuhanperawatan2 string     `gorm:"column:kebutuhanperawatan2" json:"kebutuhanperawatan2"`
	Kebutuhanperawatan3 string     `gorm:"column:kebutuhanperawatan3" json:"kebutuhanperawatan3"`
	Kebutuhanperawatan4 string     `gorm:"column:kebutuhanperawatan4" json:"kebutuhanperawatan4"`
	Tindaklanjut1       string     `gorm:"column:tindaklanjut1" json:"tindaklanjut1"`
	Tindaklanjut2       string     `gorm:"column:tindaklanjut2" json:"tindaklanjut2"`
	Tindaklanjut3       string     `gorm:"column:tindaklanjut3" json:"tindaklanjut3"`
	Tindaklanjut4       string     `gorm:"column:tindaklanjut4" json:"tindaklanjut4"`
	Keputusan           string     `gorm:"column:keputusan" json:"keputusan"`
	Keterangan          *string    `gorm:"column:keterangan" json:"keterangan"`
	Nik                 *string    `gorm:"column:nik" json:"nik"`
}

func (ChecklistKriteriaKeluarPicu) TableName() string {
	return "checklist_kriteria_keluar_picu"
}
