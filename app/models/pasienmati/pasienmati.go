package pasienmati

import "time"

// PasienMati tabel `pasien_mati` (pasien meninggal).
type PasienMati struct {
	Tanggal       *time.Time `gorm:"column:tanggal;type:date" json:"tanggal"`
	Jam           *string    `gorm:"column:jam;type:time" json:"jam"`
	NoRkmMedis    string     `gorm:"column:no_rkm_medis;primaryKey;type:varchar(15);not null" json:"no_rkm_medis"`
	Keterangan    *string    `gorm:"column:keterangan;type:varchar(100)" json:"keterangan"`
	TempMeninggal *string    `gorm:"column:temp_meninggal;type:enum('-','Rumah Sakit','Puskesmas','Rumah Bersalin','Rumah Tempat Tinggal','Lain-lain (Termasuk Doa)" json:"temp_meninggal"`
	Icd1          *string    `gorm:"column:icd1;type:varchar(20)" json:"icd1"`
	Icd2          *string    `gorm:"column:icd2;type:varchar(20)" json:"icd2"`
	Icd3          *string    `gorm:"column:icd3;type:varchar(20)" json:"icd3"`
	Icd4          *string    `gorm:"column:icd4;type:varchar(20)" json:"icd4"`
	KdDokter      string     `gorm:"column:kd_dokter;type:varchar(20);not null" json:"kd_dokter"`
}

func (PasienMati) TableName() string {
	return "pasien_mati"
}
