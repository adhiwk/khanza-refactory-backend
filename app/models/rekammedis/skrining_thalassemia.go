package rekammedis

import "time"

// SkriningThalassemia tabel `skrining_thalassemia` (skrining talasemia, RMSkriningTalasemia).
type SkriningThalassemia struct {
	NoRawat                string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip                    string     `gorm:"column:nip" json:"nip"`
	TransfusiDarah         *string    `gorm:"column:transfusi_darah" json:"transfusi_darah"`
	RutinTransfusi         *string    `gorm:"column:rutin_transfusi" json:"rutin_transfusi"`
	SaudaraThalassemia     *string    `gorm:"column:saudara_thalassemia" json:"saudara_thalassemia"`
	TumbuhKembangTerlambat *string    `gorm:"column:tumbuh_kembang_terlambat" json:"tumbuh_kembang_terlambat"`
	Anemia                 *string    `gorm:"column:anemia" json:"anemia"`
	Ikterus                *string    `gorm:"column:ikterus" json:"ikterus"`
	PerutBuncit            *string    `gorm:"column:perut_buncit" json:"perut_buncit"`
	GiziKurang             *string    `gorm:"column:gizi_kurang" json:"gizi_kurang"`
	FaciesCooley           *string    `gorm:"column:facies_cooley" json:"facies_cooley"`
	PerawakanPendek        *string    `gorm:"column:perawakan_pendek" json:"perawakan_pendek"`
	HiperpigmentasiKulit   *string    `gorm:"column:hiperpigmentasi_kulit" json:"hiperpigmentasi_kulit"`
	Hemoglobin             *string    `gorm:"column:hemoglobin" json:"hemoglobin"`
	Mvc                    *string    `gorm:"column:mvc" json:"mvc"`
	Mchc                   *string    `gorm:"column:mchc" json:"mchc"`
	DarahTepi              *string    `gorm:"column:darah_tepi" json:"darah_tepi"`
	TindakLanjut           *string    `gorm:"column:tindak_lanjut" json:"tindak_lanjut"`
}

func (SkriningThalassemia) TableName() string {
	return "skrining_thalassemia"
}
