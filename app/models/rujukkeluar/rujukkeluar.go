package rujukkeluar

import "time"

// Rujuk tabel `rujuk` (rujukan keluar).
type Rujuk struct {
	NoRujuk            string     `gorm:"column:no_rujuk;primaryKey;type:varchar(40);not null" json:"no_rujuk"`
	NoRawat            *string    `gorm:"column:no_rawat;type:varchar(17)" json:"no_rawat"`
	RujukKe            *string    `gorm:"column:rujuk_ke;type:varchar(150)" json:"rujuk_ke"`
	TglRujuk           *time.Time `gorm:"column:tgl_rujuk;type:date" json:"tgl_rujuk"`
	KeteranganDiagnosa *string    `gorm:"column:keterangan_diagnosa;type:text" json:"keterangan_diagnosa"`
	KdDokter           *string    `gorm:"column:kd_dokter;type:varchar(20)" json:"kd_dokter"`
	KatRujuk           *string    `gorm:"column:kat_rujuk;type:enum('-','Bedah','Non Bedah','Kebidanan','Anak')" json:"kat_rujuk"`
	Ambulance          *string    `gorm:"column:ambulance;type:enum('-','AGD','SENDIRI','SWASTA')" json:"ambulance"`
	Keterangan         *string    `gorm:"column:keterangan;type:text" json:"keterangan"`
	Jam                *string    `gorm:"column:jam;type:time" json:"jam"`
}

func (Rujuk) TableName() string {
	return "rujuk"
}
