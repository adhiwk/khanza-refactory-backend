package rekammedis

import "time"

// ChecklistKriteriaKeluarIcu tabel `checklist_kriteria_keluar_icu` (checklist kriteria keluar ICU, RMChecklistKriteriaKeluarICU).
type ChecklistKriteriaKeluarIcu struct {
	NoRawat    string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal    *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Kriteria1  string     `gorm:"column:kriteria1" json:"kriteria1"`
	Kriteria2  string     `gorm:"column:kriteria2" json:"kriteria2"`
	Kriteria3  string     `gorm:"column:kriteria3" json:"kriteria3"`
	Kriteria4  string     `gorm:"column:kriteria4" json:"kriteria4"`
	Kriteria5  string     `gorm:"column:kriteria5" json:"kriteria5"`
	Kriteria6  string     `gorm:"column:kriteria6" json:"kriteria6"`
	Kriteria7  string     `gorm:"column:kriteria7" json:"kriteria7"`
	Kriteria8  string     `gorm:"column:kriteria8" json:"kriteria8"`
	Kriteria9  string     `gorm:"column:kriteria9" json:"kriteria9"`
	Kriteria10 string     `gorm:"column:kriteria10" json:"kriteria10"`
	Kriteria11 string     `gorm:"column:kriteria11" json:"kriteria11"`
	Nik        *string    `gorm:"column:nik" json:"nik"`
}

func (ChecklistKriteriaKeluarIcu) TableName() string {
	return "checklist_kriteria_keluar_icu"
}
