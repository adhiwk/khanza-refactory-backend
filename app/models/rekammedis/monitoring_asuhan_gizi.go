package rekammedis

import "time"

// MonitoringAsuhanGizi tabel `monitoring_asuhan_gizi` (monitoring asuhan gizi, RMDataMonitoringAsuhanGizi).
type MonitoringAsuhanGizi struct {
	NoRawat    string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal    *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Monitoring *string    `gorm:"column:monitoring" json:"monitoring"`
	Evaluasi   *string    `gorm:"column:evaluasi" json:"evaluasi"`
	Nip        *string    `gorm:"column:nip" json:"nip"`
}

func (MonitoringAsuhanGizi) TableName() string {
	return "monitoring_asuhan_gizi"
}
