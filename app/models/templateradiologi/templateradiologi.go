// Package templateradiologi model template hasil radiologi (MasterTemplateHasilRadiologi).
package templateradiologi

type TemplateHasilRadiologi struct {
	NoTemplate             string  `gorm:"column:no_template;primaryKey" json:"no_template"`
	NamaPemeriksaan        *string `gorm:"column:nama_pemeriksaan" json:"nama_pemeriksaan"`
	TemplateHasilRadiologi *string `gorm:"column:template_hasil_radiologi" json:"template_hasil_radiologi"`
}

func (TemplateHasilRadiologi) TableName() string {
	return "template_hasil_radiologi"
}
