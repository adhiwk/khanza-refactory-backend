package rekammedis

import "time"

// PenatalaksanaanTerapiOkupasi tabel `penatalaksanaan_terapi_okupasi` (penatalaksanaan terapi okupasi, RMPenatalaksanaanTerapiOkupasi).
type PenatalaksanaanTerapiOkupasi struct {
	NoRawat                  string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                  *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip                      string     `gorm:"column:nip" json:"nip"`
	KeluhanUtama             string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rpd                      string     `gorm:"column:rpd" json:"rpd"`
	Rps                      string     `gorm:"column:rps" json:"rps"`
	AnamnesaGeneral          string     `gorm:"column:anamnesa_general" json:"anamnesa_general"`
	TandaVital               string     `gorm:"column:tanda_vital" json:"tanda_vital"`
	PemeriksaanPenunjang     string     `gorm:"column:pemeriksaan_penunjang" json:"pemeriksaan_penunjang"`
	Spesialisasi             string     `gorm:"column:spesialisasi" json:"spesialisasi"`
	KeteranganSpesialisasi   string     `gorm:"column:keterangan_spesialisasi" json:"keterangan_spesialisasi"`
	PemeriksaanOkupasiTerapi string     `gorm:"column:pemeriksaan_okupasi_terapi" json:"pemeriksaan_okupasi_terapi"`
	Aset                     string     `gorm:"column:aset" json:"aset"`
	Limitasi                 string     `gorm:"column:limitasi" json:"limitasi"`
	DiagnosaTerapiOkupasi    string     `gorm:"column:diagnosa_terapi_okupasi" json:"diagnosa_terapi_okupasi"`
	RencanaIntervensi        string     `gorm:"column:rencana_intervensi" json:"rencana_intervensi"`
}

func (PenatalaksanaanTerapiOkupasi) TableName() string {
	return "penatalaksanaan_terapi_okupasi"
}
