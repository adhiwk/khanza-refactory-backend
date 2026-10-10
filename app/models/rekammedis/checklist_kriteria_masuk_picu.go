package rekammedis

import "time"

// ChecklistKriteriaMasukPicu tabel `checklist_kriteria_masuk_picu` (checklist kriteria masuk PICU, RMChecklistKriteriaMasukPICU).
type ChecklistKriteriaMasukPicu struct {
	NoRawat       string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal       *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Kriteriaumum1 string     `gorm:"column:kriteriaumum1" json:"kriteriaumum1"`
	Kriteriaumum2 string     `gorm:"column:kriteriaumum2" json:"kriteriaumum2"`
	Kriteriaumum3 string     `gorm:"column:kriteriaumum3" json:"kriteriaumum3"`
	Respirasi1    string     `gorm:"column:respirasi1" json:"respirasi1"`
	Respirasi2    string     `gorm:"column:respirasi2" json:"respirasi2"`
	Respirasi3    string     `gorm:"column:respirasi3" json:"respirasi3"`
	Respirasi4    string     `gorm:"column:respirasi4" json:"respirasi4"`
	Kardio1       string     `gorm:"column:kardio1" json:"kardio1"`
	Kardio2       string     `gorm:"column:kardio2" json:"kardio2"`
	Kardio3       string     `gorm:"column:kardio3" json:"kardio3"`
	Kardio4       string     `gorm:"column:kardio4" json:"kardio4"`
	Neuro1        string     `gorm:"column:neuro1" json:"neuro1"`
	Neuro2        string     `gorm:"column:neuro2" json:"neuro2"`
	Neuro3        string     `gorm:"column:neuro3" json:"neuro3"`
	Neuro4        string     `gorm:"column:neuro4" json:"neuro4"`
	Bedah1        string     `gorm:"column:bedah1" json:"bedah1"`
	Bedah2        string     `gorm:"column:bedah2" json:"bedah2"`
	Bedah3        string     `gorm:"column:bedah3" json:"bedah3"`
	Kondisilain1  string     `gorm:"column:kondisilain1" json:"kondisilain1"`
	Kondisilain2  string     `gorm:"column:kondisilain2" json:"kondisilain2"`
	Kondisilain3  string     `gorm:"column:kondisilain3" json:"kondisilain3"`
	Keputusan     string     `gorm:"column:keputusan" json:"keputusan"`
	Keterangan    *string    `gorm:"column:keterangan" json:"keterangan"`
	Nik           *string    `gorm:"column:nik" json:"nik"`
}

func (ChecklistKriteriaMasukPicu) TableName() string {
	return "checklist_kriteria_masuk_picu"
}
