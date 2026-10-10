package rekammedis

import "time"

// CatatanCairanHemodialisa tabel `catatan_cairan_hemodialisa` (catatan cairan hemodialisa, RMDataCatatanCairanHemodialisa).
type CatatanCairanHemodialisa struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat     string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	Minum        *string    `gorm:"column:minum" json:"minum"`
	Infus        *string    `gorm:"column:infus" json:"infus"`
	Tranfusi     *string    `gorm:"column:tranfusi" json:"tranfusi"`
	SisaPriming  *string    `gorm:"column:sisa_priming" json:"sisa_priming"`
	WashOut      *string    `gorm:"column:wash_out" json:"wash_out"`
	Urine        *string    `gorm:"column:urine" json:"urine"`
	Pendarahan   *string    `gorm:"column:pendarahan" json:"pendarahan"`
	Muntah       *string    `gorm:"column:muntah" json:"muntah"`
	Keterangan   *string    `gorm:"column:keterangan" json:"keterangan"`
	Nip          *string    `gorm:"column:nip" json:"nip"`
}

func (CatatanCairanHemodialisa) TableName() string {
	return "catatan_cairan_hemodialisa"
}
