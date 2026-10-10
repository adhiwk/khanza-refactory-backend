package rekammedis

import "time"

// ChecklistKriteriaKeluarNicu tabel `checklist_kriteria_keluar_nicu` (checklist kriteria keluar NICU, RMChecklistKriteriaKeluarNICU).
type ChecklistKriteriaKeluarNicu struct {
	NoRawat    string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal    *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Respirasi1 string     `gorm:"column:respirasi1" json:"respirasi1"`
	Respirasi2 string     `gorm:"column:respirasi2" json:"respirasi2"`
	Respirasi3 string     `gorm:"column:respirasi3" json:"respirasi3"`
	Kardio1    string     `gorm:"column:kardio1" json:"kardio1"`
	Kardio2    string     `gorm:"column:kardio2" json:"kardio2"`
	Nutrisi1   string     `gorm:"column:nutrisi1" json:"nutrisi1"`
	Nutrisi2   string     `gorm:"column:nutrisi2" json:"nutrisi2"`
	Nutrisi3   string     `gorm:"column:nutrisi3" json:"nutrisi3"`
	Suhutubuh1 string     `gorm:"column:suhutubuh1" json:"suhutubuh1"`
	Suhutubuh2 string     `gorm:"column:suhutubuh2" json:"suhutubuh2"`
	Infeksi1   string     `gorm:"column:infeksi1" json:"infeksi1"`
	Infeksi2   string     `gorm:"column:infeksi2" json:"infeksi2"`
	Infeksi3   string     `gorm:"column:infeksi3" json:"infeksi3"`
	Keputusan  string     `gorm:"column:keputusan" json:"keputusan"`
	Keterangan *string    `gorm:"column:keterangan" json:"keterangan"`
	Nik        *string    `gorm:"column:nik" json:"nik"`
}

func (ChecklistKriteriaKeluarNicu) TableName() string {
	return "checklist_kriteria_keluar_nicu"
}
