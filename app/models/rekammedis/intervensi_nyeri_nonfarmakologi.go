package rekammedis

import "time"

// IntervensiNyeriNonfarmakologi tabel `intervensi_nyeri_nonfarmakologi` (intervensi nyeri non farmakologi, RMDataIntervensiNyeriNonFarmakologi).
type IntervensiNyeriNonfarmakologi struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat     string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	Intervensi   *string    `gorm:"column:intervensi" json:"intervensi"`
	Nip          string     `gorm:"column:nip" json:"nip"`
}

func (IntervensiNyeriNonfarmakologi) TableName() string {
	return "intervensi_nyeri_nonfarmakologi"
}
