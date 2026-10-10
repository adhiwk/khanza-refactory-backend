package rekammedis

import "time"

// ChecklistKriteriaMasukIcu tabel `checklist_kriteria_masuk_icu` (checklist kriteria masuk ICU, RMChecklistKriteriaMasukICU).
type ChecklistKriteriaMasukIcu struct {
	NoRawat                       string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                       *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Prioritas11                   string     `gorm:"column:prioritas1_1" json:"prioritas1_1"`
	Prioritas12                   string     `gorm:"column:prioritas1_2" json:"prioritas1_2"`
	Prioritas13                   string     `gorm:"column:prioritas1_3" json:"prioritas1_3"`
	Prioritas14                   string     `gorm:"column:prioritas1_4" json:"prioritas1_4"`
	Prioritas15                   string     `gorm:"column:prioritas1_5" json:"prioritas1_5"`
	Prioritas16                   string     `gorm:"column:prioritas1_6" json:"prioritas1_6"`
	Prioritas21                   string     `gorm:"column:prioritas2_1" json:"prioritas2_1"`
	Prioritas22                   string     `gorm:"column:prioritas2_2" json:"prioritas2_2"`
	Prioritas23                   string     `gorm:"column:prioritas2_3" json:"prioritas2_3"`
	Prioritas24                   string     `gorm:"column:prioritas2_4" json:"prioritas2_4"`
	Prioritas25                   string     `gorm:"column:prioritas2_5" json:"prioritas2_5"`
	Prioritas26                   string     `gorm:"column:prioritas2_6" json:"prioritas2_6"`
	Prioritas27                   string     `gorm:"column:prioritas2_7" json:"prioritas2_7"`
	Prioritas28                   string     `gorm:"column:prioritas2_8" json:"prioritas2_8"`
	Prioritas31                   string     `gorm:"column:prioritas3_1" json:"prioritas3_1"`
	Prioritas32                   string     `gorm:"column:prioritas3_2" json:"prioritas3_2"`
	Prioritas33                   string     `gorm:"column:prioritas3_3" json:"prioritas3_3"`
	Prioritas34                   string     `gorm:"column:prioritas3_4" json:"prioritas3_4"`
	KriteriaFisiologisTandaVital1 string     `gorm:"column:kriteria_fisiologis_tanda_vital_1" json:"kriteria_fisiologis_tanda_vital_1"`
	KriteriaFisiologisTandaVital2 string     `gorm:"column:kriteria_fisiologis_tanda_vital_2" json:"kriteria_fisiologis_tanda_vital_2"`
	KriteriaFisiologisTandaVital3 string     `gorm:"column:kriteria_fisiologis_tanda_vital_3" json:"kriteria_fisiologis_tanda_vital_3"`
	KriteriaFisiologisTandaVital4 string     `gorm:"column:kriteria_fisiologis_tanda_vital_4" json:"kriteria_fisiologis_tanda_vital_4"`
	KriteriaFisiologisTandaVital5 string     `gorm:"column:kriteria_fisiologis_tanda_vital_5" json:"kriteria_fisiologis_tanda_vital_5"`
	KriteriaFisiologisLaborat1    string     `gorm:"column:kriteria_fisiologis_laborat_1" json:"kriteria_fisiologis_laborat_1"`
	KriteriaFisiologisLaborat2    string     `gorm:"column:kriteria_fisiologis_laborat_2" json:"kriteria_fisiologis_laborat_2"`
	KriteriaFisiologisLaborat3    string     `gorm:"column:kriteria_fisiologis_laborat_3" json:"kriteria_fisiologis_laborat_3"`
	KriteriaFisiologisLaborat4    string     `gorm:"column:kriteria_fisiologis_laborat_4" json:"kriteria_fisiologis_laborat_4"`
	KriteriaFisiologisLaborat5    string     `gorm:"column:kriteria_fisiologis_laborat_5" json:"kriteria_fisiologis_laborat_5"`
	KriteriaFisiologisLaborat6    string     `gorm:"column:kriteria_fisiologis_laborat_6" json:"kriteria_fisiologis_laborat_6"`
	KriteriaFisiologisRadiologi1  string     `gorm:"column:kriteria_fisiologis_radiologi_1" json:"kriteria_fisiologis_radiologi_1"`
	KriteriaFisiologisRadiologi2  string     `gorm:"column:kriteria_fisiologis_radiologi_2" json:"kriteria_fisiologis_radiologi_2"`
	KriteriaFisiologisKlinis1     string     `gorm:"column:kriteria_fisiologis_klinis_1" json:"kriteria_fisiologis_klinis_1"`
	KriteriaFisiologisKlinis2     string     `gorm:"column:kriteria_fisiologis_klinis_2" json:"kriteria_fisiologis_klinis_2"`
	KriteriaFisiologisKlinis3     string     `gorm:"column:kriteria_fisiologis_klinis_3" json:"kriteria_fisiologis_klinis_3"`
	KriteriaFisiologisKlinis4     string     `gorm:"column:kriteria_fisiologis_klinis_4" json:"kriteria_fisiologis_klinis_4"`
	KriteriaFisiologisKlinis5     string     `gorm:"column:kriteria_fisiologis_klinis_5" json:"kriteria_fisiologis_klinis_5"`
	KriteriaFisiologisKlinis6     string     `gorm:"column:kriteria_fisiologis_klinis_6" json:"kriteria_fisiologis_klinis_6"`
	KriteriaFisiologisKlinis7     string     `gorm:"column:kriteria_fisiologis_klinis_7" json:"kriteria_fisiologis_klinis_7"`
	KriteriaFisiologisKlinis8     string     `gorm:"column:kriteria_fisiologis_klinis_8" json:"kriteria_fisiologis_klinis_8"`
	Nik                           *string    `gorm:"column:nik" json:"nik"`
}

func (ChecklistKriteriaMasukIcu) TableName() string {
	return "checklist_kriteria_masuk_icu"
}
