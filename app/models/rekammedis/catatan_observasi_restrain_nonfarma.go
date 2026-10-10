package rekammedis

import "time"

// CatatanObservasiRestrainNonfarma tabel `catatan_observasi_restrain_nonfarma` (catatan observasi restrain non farmakologi, RMDataCatatanObservasiRestrainNonFarmakologi).
type CatatanObservasiRestrainNonfarma struct {
	NoRawat           string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan      *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat          string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	TanganKiri        *string    `gorm:"column:tangan_kiri" json:"tangan_kiri"`
	TanganKanan       *string    `gorm:"column:tangan_kanan" json:"tangan_kanan"`
	KakiKiri          *string    `gorm:"column:kaki_kiri" json:"kaki_kiri"`
	KakiKanan         *string    `gorm:"column:kaki_kanan" json:"kaki_kanan"`
	Badan             *string    `gorm:"column:badan" json:"badan"`
	Edema             *string    `gorm:"column:edema" json:"edema"`
	Iritasi           *string    `gorm:"column:iritasi" json:"iritasi"`
	Sirkulasi         *string    `gorm:"column:sirkulasi" json:"sirkulasi"`
	KondisiKeterangan *string    `gorm:"column:kondisi_keterangan" json:"kondisi_keterangan"`
	Nip               string     `gorm:"column:nip" json:"nip"`
}

func (CatatanObservasiRestrainNonfarma) TableName() string {
	return "catatan_observasi_restrain_nonfarma"
}
