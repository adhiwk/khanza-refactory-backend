package rekammedis

import "time"

// CatatanObservasiBayi tabel `catatan_observasi_bayi` (catatan observasi bayi, RMDataCatatanObservasiBayi).
type CatatanObservasiBayi struct {
	NoRawat       string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	TglPerawatan  *time.Time `gorm:"column:tgl_perawatan;primaryKey;autoIncrement:false" json:"tgl_perawatan"`
	JamRawat      string     `gorm:"column:jam_rawat;primaryKey;autoIncrement:false" json:"jam_rawat"`
	Gcs           *string    `gorm:"column:gcs" json:"gcs"`
	Td            *string    `gorm:"column:td" json:"td"`
	Hr            *string    `gorm:"column:hr" json:"hr"`
	Rr            *string    `gorm:"column:rr" json:"rr"`
	Suhu          *string    `gorm:"column:suhu" json:"suhu"`
	Spo2          *string    `gorm:"column:spo2" json:"spo2"`
	Nch           *string    `gorm:"column:nch" json:"nch"`
	IkterikStatus *string    `gorm:"column:ikterik_status" json:"ikterik_status"`
	RetraksiDada  *string    `gorm:"column:retraksi_dada" json:"retraksi_dada"`
	OgtResidu     *string    `gorm:"column:ogt_residu" json:"ogt_residu"`
	AsiJumlah     *string    `gorm:"column:asi_jumlah" json:"asi_jumlah"`
	PasiJumlah    *string    `gorm:"column:pasi_jumlah" json:"pasi_jumlah"`
	BakStatus     *string    `gorm:"column:bak_status" json:"bak_status"`
	BabStatus     *string    `gorm:"column:bab_status" json:"bab_status"`
	Nip           *string    `gorm:"column:nip" json:"nip"`
}

func (CatatanObservasiBayi) TableName() string {
	return "catatan_observasi_bayi"
}
