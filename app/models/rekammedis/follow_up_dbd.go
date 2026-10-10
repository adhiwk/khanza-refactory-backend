package rekammedis

import "time"

// FollowUpDbd tabel `follow_up_dbd` (follow up DBD, RMDataFollowUpDBD).
type FollowUpDbd struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat     string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	Hemoglobin   *string    `gorm:"column:hemoglobin" json:"hemoglobin"`
	Hematokrit   string     `gorm:"column:hematokrit" json:"hematokrit"`
	Leokosit     *string    `gorm:"column:leokosit" json:"leokosit"`
	Trombosit    *string    `gorm:"column:trombosit" json:"trombosit"`
	TerapiCairan *string    `gorm:"column:terapi_cairan" json:"terapi_cairan"`
	Nip          string     `gorm:"column:nip" json:"nip"`
}

func (FollowUpDbd) TableName() string {
	return "follow_up_dbd"
}
