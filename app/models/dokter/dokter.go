package dokter

import "time"

// Dokter tabel `dokter` (dokter).
type Dokter struct {
	KdDokter     string     `gorm:"column:kd_dokter;primaryKey;type:varchar(20);not null" json:"kd_dokter"`
	NmDokter     *string    `gorm:"column:nm_dokter;type:varchar(50)" json:"nm_dokter"`
	Jk           *string    `gorm:"column:jk;type:enum('L','P')" json:"jk"`
	TmpLahir     *string    `gorm:"column:tmp_lahir;type:varchar(20)" json:"tmp_lahir"`
	TglLahir     *time.Time `gorm:"column:tgl_lahir;type:date" json:"tgl_lahir"`
	GolDrh       *string    `gorm:"column:gol_drh;type:enum('A','B','O','AB','-')" json:"gol_drh"`
	Agama        *string    `gorm:"column:agama;type:varchar(12)" json:"agama"`
	AlmtTgl      *string    `gorm:"column:almt_tgl;type:varchar(60)" json:"almt_tgl"`
	NoTelp       *string    `gorm:"column:no_telp;type:varchar(13)" json:"no_telp"`
	Email        string     `gorm:"column:email;type:varchar(70);not null" json:"email"`
	SttsNikah    *string    `gorm:"column:stts_nikah;type:enum('BELUM MENIKAH','MENIKAH','JANDA','DUDHA','JOMBLO')" json:"stts_nikah"`
	KdSps        *string    `gorm:"column:kd_sps;type:char(5)" json:"kd_sps"`
	Alumni       *string    `gorm:"column:alumni;type:varchar(60)" json:"alumni"`
	NoIjnPraktek *string    `gorm:"column:no_ijn_praktek;type:varchar(120)" json:"no_ijn_praktek"`
	Status       string     `gorm:"column:status;type:enum('0','1');not null" json:"status"`
}

func (Dokter) TableName() string {
	return "dokter"
}
