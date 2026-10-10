package rekammedis

import "time"

// MonitoringReaksiTranfusi tabel `monitoring_reaksi_tranfusi` (monitoring reaksi tranfusi, RMDataMonitoringReaksiTranfusi).
type MonitoringReaksiTranfusi struct {
	NoRawat           string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan      *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat          string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	ProdukDarah       *string    `gorm:"column:produk_darah" json:"produk_darah"`
	NoKantong         *string    `gorm:"column:no_kantong" json:"no_kantong"`
	LokasiInsersi     string     `gorm:"column:lokasi_insersi" json:"lokasi_insersi"`
	Td                string     `gorm:"column:td" json:"td"`
	Hr                *string    `gorm:"column:hr" json:"hr"`
	Rr                *string    `gorm:"column:rr" json:"rr"`
	Suhu              *string    `gorm:"column:suhu" json:"suhu"`
	JenisReaksiAlergi *string    `gorm:"column:jenis_reaksi_alergi" json:"jenis_reaksi_alergi"`
	Keterangan        *string    `gorm:"column:keterangan" json:"keterangan"`
	Nip               string     `gorm:"column:nip" json:"nip"`
}

func (MonitoringReaksiTranfusi) TableName() string {
	return "monitoring_reaksi_tranfusi"
}
