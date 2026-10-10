// Package templateoperasi model template laporan operasi (MasterTemplateLaporanOperasi).
package templateoperasi

type TemplateLaporanOperasi struct {
	NoTemplate       string `gorm:"column:no_template;primaryKey" json:"no_template"`
	NamaOperasi      string `gorm:"column:nama_operasi" json:"nama_operasi"`
	DiagnosaPreop    string `gorm:"column:diagnosa_preop" json:"diagnosa_preop"`
	DiagnosaPostop   string `gorm:"column:diagnosa_postop" json:"diagnosa_postop"`
	JaringanDieksisi string `gorm:"column:jaringan_dieksisi" json:"jaringan_dieksisi"`
	PermintaanPa     string `gorm:"column:permintaan_pa" json:"permintaan_pa"`
	LaporanOperasi   string `gorm:"column:laporan_operasi" json:"laporan_operasi"`
}

func (TemplateLaporanOperasi) TableName() string {
	return "template_laporan_operasi"
}
