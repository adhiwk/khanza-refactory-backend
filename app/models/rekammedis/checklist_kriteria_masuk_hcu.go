package rekammedis

import "time"

// ChecklistKriteriaMasukHcu tabel `checklist_kriteria_masuk_hcu` (checklist kriteria masuk HCU, RMChecklistKriteriaMasukHCU).
type ChecklistKriteriaMasukHcu struct {
	NoRawat     string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal     *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Kardiologi1 string     `gorm:"column:kardiologi1" json:"kardiologi1"`
	Kardiologi2 string     `gorm:"column:kardiologi2" json:"kardiologi2"`
	Kardiologi3 string     `gorm:"column:kardiologi3" json:"kardiologi3"`
	Kardiologi4 string     `gorm:"column:kardiologi4" json:"kardiologi4"`
	Kardiologi5 string     `gorm:"column:kardiologi5" json:"kardiologi5"`
	Kardiologi6 string     `gorm:"column:kardiologi6" json:"kardiologi6"`
	Pernapasan1 string     `gorm:"column:pernapasan1" json:"pernapasan1"`
	Pernapasan2 string     `gorm:"column:pernapasan2" json:"pernapasan2"`
	Pernapasan3 string     `gorm:"column:pernapasan3" json:"pernapasan3"`
	Syaraf1     string     `gorm:"column:syaraf1" json:"syaraf1"`
	Syaraf2     string     `gorm:"column:syaraf2" json:"syaraf2"`
	Syaraf3     string     `gorm:"column:syaraf3" json:"syaraf3"`
	Syaraf4     string     `gorm:"column:syaraf4" json:"syaraf4"`
	Pencernaan1 string     `gorm:"column:pencernaan1" json:"pencernaan1"`
	Pencernaan2 string     `gorm:"column:pencernaan2" json:"pencernaan2"`
	Pencernaan3 string     `gorm:"column:pencernaan3" json:"pencernaan3"`
	Pencernaan4 string     `gorm:"column:pencernaan4" json:"pencernaan4"`
	Pembedahan1 string     `gorm:"column:pembedahan1" json:"pembedahan1"`
	Pembedahan2 string     `gorm:"column:pembedahan2" json:"pembedahan2"`
	Hematologi1 string     `gorm:"column:hematologi1" json:"hematologi1"`
	Hematologi2 string     `gorm:"column:hematologi2" json:"hematologi2"`
	Infeksi     string     `gorm:"column:infeksi" json:"infeksi"`
	Nik         *string    `gorm:"column:nik" json:"nik"`
}

func (ChecklistKriteriaMasukHcu) TableName() string {
	return "checklist_kriteria_masuk_hcu"
}
