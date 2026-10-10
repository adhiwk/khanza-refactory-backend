package rekammedis

import "time"

// ChecklistKriteriaMasukNicu tabel `checklist_kriteria_masuk_nicu` (checklist kriteria masuk NICU, RMChecklistKriteriaMasukNICU).
type ChecklistKriteriaMasukNicu struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal      *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Respirasi1   string     `gorm:"column:respirasi1" json:"respirasi1"`
	Respirasi2   string     `gorm:"column:respirasi2" json:"respirasi2"`
	Respirasi3   string     `gorm:"column:respirasi3" json:"respirasi3"`
	Respirasi4   string     `gorm:"column:respirasi4" json:"respirasi4"`
	Prematur1    string     `gorm:"column:prematur1" json:"prematur1"`
	Prematur2    string     `gorm:"column:prematur2" json:"prematur2"`
	Prematur3    string     `gorm:"column:prematur3" json:"prematur3"`
	Kardio1      string     `gorm:"column:kardio1" json:"kardio1"`
	Kardio2      string     `gorm:"column:kardio2" json:"kardio2"`
	Kardio3      string     `gorm:"column:kardio3" json:"kardio3"`
	Neuro1       string     `gorm:"column:neuro1" json:"neuro1"`
	Neuro2       string     `gorm:"column:neuro2" json:"neuro2"`
	Neuro3       string     `gorm:"column:neuro3" json:"neuro3"`
	Metabolik1   string     `gorm:"column:metabolik1" json:"metabolik1"`
	Metabolik2   string     `gorm:"column:metabolik2" json:"metabolik2"`
	Metabolik3   string     `gorm:"column:metabolik3" json:"metabolik3"`
	Kondisilain1 string     `gorm:"column:kondisilain1" json:"kondisilain1"`
	Kondisilain2 string     `gorm:"column:kondisilain2" json:"kondisilain2"`
	Kondisilain3 string     `gorm:"column:kondisilain3" json:"kondisilain3"`
	Kondisilain4 string     `gorm:"column:kondisilain4" json:"kondisilain4"`
	Keputusan    string     `gorm:"column:keputusan" json:"keputusan"`
	Keterangan   *string    `gorm:"column:keterangan" json:"keterangan"`
	Nik          *string    `gorm:"column:nik" json:"nik"`
}

func (ChecklistKriteriaMasukNicu) TableName() string {
	return "checklist_kriteria_masuk_nicu"
}
