package cacatfisik

// CacatFisik tabel `cacat_fisik` (cacat fisik).
type CacatFisik struct {
	Id        int    `gorm:"column:id;primaryKey;autoIncrement;type:int(11);not null" json:"id"`
	NamaCacat string `gorm:"column:nama_cacat;type:varchar(30);not null" json:"nama_cacat"`
}

func (CacatFisik) TableName() string {
	return "cacat_fisik"
}
