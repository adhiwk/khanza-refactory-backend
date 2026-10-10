package rekammedis

import "time"

// MppSkrining tabel `mpp_skrining` (skrining MPP, RMSkriningMPP).
type MppSkrining struct {
	NoRawat string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Param1  *string    `gorm:"column:param1" json:"param1"`
	Param2  *string    `gorm:"column:param2" json:"param2"`
	Param3  *string    `gorm:"column:param3" json:"param3"`
	Param4  *string    `gorm:"column:param4" json:"param4"`
	Param5  *string    `gorm:"column:param5" json:"param5"`
	Param6  *string    `gorm:"column:param6" json:"param6"`
	Param7  *string    `gorm:"column:param7" json:"param7"`
	Param8  *string    `gorm:"column:param8" json:"param8"`
	Param9  *string    `gorm:"column:param9" json:"param9"`
	Param10 *string    `gorm:"column:param10" json:"param10"`
	Param11 *string    `gorm:"column:param11" json:"param11"`
	Param12 *string    `gorm:"column:param12" json:"param12"`
	Param13 *string    `gorm:"column:param13" json:"param13"`
	Param14 *string    `gorm:"column:param14" json:"param14"`
	Param15 *string    `gorm:"column:param15" json:"param15"`
	Param16 *string    `gorm:"column:param16" json:"param16"`
	Nip     string     `gorm:"column:nip" json:"nip"`
}

func (MppSkrining) TableName() string {
	return "mpp_skrining"
}
