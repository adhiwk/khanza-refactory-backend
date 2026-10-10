package rekammedis

import "time"

// HasilEndoskopiFaringLaring tabel `hasil_endoskopi_faring_laring` (hasil endoskopi faring laring, RMHasilEndoskopiFaringLaring).
type HasilEndoskopiFaringLaring struct {
	NoRawat                  string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                  *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter                 string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	DiagnosaKlinis           string     `gorm:"column:diagnosa_klinis" json:"diagnosa_klinis"`
	KirimanDari              string     `gorm:"column:kiriman_dari" json:"kiriman_dari"`
	FaringUvula              *string    `gorm:"column:faring_uvula" json:"faring_uvula"`
	FaringArkusFaring        *string    `gorm:"column:faring_arkus_faring" json:"faring_arkus_faring"`
	FaringDindingPosterior   *string    `gorm:"column:faring_dinding_posterior" json:"faring_dinding_posterior"`
	FaringTonsil             *string    `gorm:"column:faring_tonsil" json:"faring_tonsil"`
	LaringTonsilLingual      *string    `gorm:"column:laring_tonsil_lingual" json:"laring_tonsil_lingual"`
	LaringValekula           *string    `gorm:"column:laring_valekula" json:"laring_valekula"`
	LaringSinusPiriformis    *string    `gorm:"column:laring_sinus_piriformis" json:"laring_sinus_piriformis"`
	LaringEpiglotis          *string    `gorm:"column:laring_epiglotis" json:"laring_epiglotis"`
	LaringArytenoid          *string    `gorm:"column:laring_arytenoid" json:"laring_arytenoid"`
	LaringPlikaVentrikularis *string    `gorm:"column:laring_plika_ventrikularis" json:"laring_plika_ventrikularis"`
	LaringPitaSuara          *string    `gorm:"column:laring_pita_suara" json:"laring_pita_suara"`
	LaringRimaVocalis        *string    `gorm:"column:laring_rima_vocalis" json:"laring_rima_vocalis"`
	LaringLainlain           *string    `gorm:"column:laring_lainlain" json:"laring_lainlain"`
	Kesan                    *string    `gorm:"column:kesan" json:"kesan"`
	Saran                    *string    `gorm:"column:saran" json:"saran"`
}

func (HasilEndoskopiFaringLaring) TableName() string {
	return "hasil_endoskopi_faring_laring"
}
