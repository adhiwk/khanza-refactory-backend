// Package templateedukasi model template informasi edukasi (MasterTemplateInformasiEdukasi).
package templateedukasi

type TemplatePelaksanaanInformasiEdukasi struct {
	NoTemplate    string  `gorm:"column:no_template;primaryKey" json:"no_template"`
	MateriEdukasi *string `gorm:"column:materi_edukasi" json:"materi_edukasi"`
	LamaEdukasi   *string `gorm:"column:lama_edukasi" json:"lama_edukasi"`
	MetodeEdukasi *string `gorm:"column:metode_edukasi" json:"metode_edukasi"`
}

func (TemplatePelaksanaanInformasiEdukasi) TableName() string {
	return "template_pelaksanaan_informasi_edukasi"
}
