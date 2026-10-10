package rekammedis

import "time"

// IntervensiNyeriFarmakologi tabel `intervensi_nyeri_farmakologi` (intervensi nyeri farmakologi, RMDataIntervensiNyeriFarmakologi).
type IntervensiNyeriFarmakologi struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat     string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	NamaObat     *string    `gorm:"column:nama_obat" json:"nama_obat"`
	DosisEfek    *string    `gorm:"column:dosis_efek" json:"dosis_efek"`
	Rute         *string    `gorm:"column:rute" json:"rute"`
	Nip          string     `gorm:"column:nip" json:"nip"`
}

func (IntervensiNyeriFarmakologi) TableName() string {
	return "intervensi_nyeri_farmakologi"
}
