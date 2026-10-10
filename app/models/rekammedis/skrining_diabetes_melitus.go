package rekammedis

import "time"

// SkriningDiabetesMelitus tabel `skrining_diabetes_melitus` (skrining diabetes melitus, RMSkriningDiabetesMelitus).
type SkriningDiabetesMelitus struct {
	NoRawat            string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal            *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip                string     `gorm:"column:nip" json:"nip"`
	Anamnesis1         *string    `gorm:"column:anamnesis1" json:"anamnesis1"`
	Anamnesis2         *string    `gorm:"column:anamnesis2" json:"anamnesis2"`
	Anamnesis3         *string    `gorm:"column:anamnesis3" json:"anamnesis3"`
	Anamnesis4         *string    `gorm:"column:anamnesis4" json:"anamnesis4"`
	Anamnesis5         *string    `gorm:"column:anamnesis5" json:"anamnesis5"`
	Anamnesis6         *string    `gorm:"column:anamnesis6" json:"anamnesis6"`
	Anamnesis7         *string    `gorm:"column:anamnesis7" json:"anamnesis7"`
	Anamnesis8         *string    `gorm:"column:anamnesis8" json:"anamnesis8"`
	Anamnesis9         *string    `gorm:"column:anamnesis9" json:"anamnesis9"`
	Anamnesis10        *string    `gorm:"column:anamnesis10" json:"anamnesis10"`
	Anamnesis11        *string    `gorm:"column:anamnesis11" json:"anamnesis11"`
	Anamnesis12        *string    `gorm:"column:anamnesis12" json:"anamnesis12"`
	BeratBadan         *string    `gorm:"column:berat_badan" json:"berat_badan"`
	TinggiBadan        *string    `gorm:"column:tinggi_badan" json:"tinggi_badan"`
	Imt                *string    `gorm:"column:imt" json:"imt"`
	KasifikasiImt      *string    `gorm:"column:kasifikasi_imt" json:"kasifikasi_imt"`
	HasilGds           *string    `gorm:"column:hasil_gds" json:"hasil_gds"`
	KeteranganGds      *string    `gorm:"column:keterangan_gds" json:"keterangan_gds"`
	HasilGdp           *string    `gorm:"column:hasil_gdp" json:"hasil_gdp"`
	KeteranganGdp      *string    `gorm:"column:keterangan_gdp" json:"keterangan_gdp"`
	HasilSkrining      *string    `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	KeteranganSkrining *string    `gorm:"column:keterangan_skrining" json:"keterangan_skrining"`
}

func (SkriningDiabetesMelitus) TableName() string {
	return "skrining_diabetes_melitus"
}
