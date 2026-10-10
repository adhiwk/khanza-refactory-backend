package rekammedis

import "time"

// SkriningPneumoniaSeverityIndex tabel `skrining_pneumonia_severity_index` (skrining pneumonia severity index, RMSkriningPneumoniaSeverityIndex).
type SkriningPneumoniaSeverityIndex struct {
	NoRawat                      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                      *time.Time `gorm:"column:tanggal" json:"tanggal"`
	NilaiUmur                    *int       `gorm:"column:nilai_umur" json:"nilai_umur"`
	TinggalDiPantiJompo          *string    `gorm:"column:tinggal_di_panti_jompo" json:"tinggal_di_panti_jompo"`
	NilaiTinggalDiPantiJompo     *int       `gorm:"column:nilai_tinggal_di_panti_jompo" json:"nilai_tinggal_di_panti_jompo"`
	GagalJantung                 *string    `gorm:"column:gagal_jantung" json:"gagal_jantung"`
	NilaiGagalJantung            *int       `gorm:"column:nilai_gagal_jantung" json:"nilai_gagal_jantung"`
	PenyakitHati                 *string    `gorm:"column:penyakit_hati" json:"penyakit_hati"`
	NilaiPenyakitHati            *int       `gorm:"column:nilai_penyakit_hati" json:"nilai_penyakit_hati"`
	PenyakitGinjal               *string    `gorm:"column:penyakit_ginjal" json:"penyakit_ginjal"`
	NilaiPenyakitGinjal          *int       `gorm:"column:nilai_penyakit_ginjal" json:"nilai_penyakit_ginjal"`
	Kanker                       *string    `gorm:"column:kanker" json:"kanker"`
	NilaiKanker                  *int       `gorm:"column:nilai_kanker" json:"nilai_kanker"`
	PenyakitSerebrovaskuler      *string    `gorm:"column:penyakit_serebrovaskuler" json:"penyakit_serebrovaskuler"`
	NilaiPenyakitSerebrovaskuler *int       `gorm:"column:nilai_penyakit_serebrovaskuler" json:"nilai_penyakit_serebrovaskuler"`
	DisorentasiMental            *string    `gorm:"column:disorentasi_mental" json:"disorentasi_mental"`
	NilaiDisorentasiMental       *int       `gorm:"column:nilai_disorentasi_mental" json:"nilai_disorentasi_mental"`
	FrekuensiNapas               *string    `gorm:"column:frekuensi_napas" json:"frekuensi_napas"`
	NilaiFrekuensiNapas          *int       `gorm:"column:nilai_frekuensi_napas" json:"nilai_frekuensi_napas"`
	TdSistolik                   *string    `gorm:"column:td_sistolik" json:"td_sistolik"`
	NilaiTdSistolik              *int       `gorm:"column:nilai_td_sistolik" json:"nilai_td_sistolik"`
	Suhu                         *string    `gorm:"column:suhu" json:"suhu"`
	NilaiSuhu                    *int       `gorm:"column:nilai_suhu" json:"nilai_suhu"`
	Nadi                         *string    `gorm:"column:nadi" json:"nadi"`
	NilaiNadi                    *int       `gorm:"column:nilai_nadi" json:"nilai_nadi"`
	PhDarah                      *string    `gorm:"column:ph_darah" json:"ph_darah"`
	NilaiPhDarah                 *int       `gorm:"column:nilai_ph_darah" json:"nilai_ph_darah"`
	Natrium                      *string    `gorm:"column:natrium" json:"natrium"`
	NilaiNatrium                 *int       `gorm:"column:nilai_natrium" json:"nilai_natrium"`
	Bun                          *string    `gorm:"column:bun" json:"bun"`
	NilaiBun                     *int       `gorm:"column:nilai_bun" json:"nilai_bun"`
	Pao                          *string    `gorm:"column:pao" json:"pao"`
	NilaiPao                     *int       `gorm:"column:nilai_pao" json:"nilai_pao"`
	Glukosa                      *string    `gorm:"column:glukosa" json:"glukosa"`
	NilaiGlukosa                 *int       `gorm:"column:nilai_glukosa" json:"nilai_glukosa"`
	EfusiPleura                  *string    `gorm:"column:efusi_pleura" json:"efusi_pleura"`
	NilaiEfusiPleura             *int       `gorm:"column:nilai_efusi_pleura" json:"nilai_efusi_pleura"`
	Hematokrit                   *string    `gorm:"column:hematokrit" json:"hematokrit"`
	NilaiHematokrit              *int       `gorm:"column:nilai_hematokrit" json:"nilai_hematokrit"`
	TotalSkor                    *int       `gorm:"column:total_skor" json:"total_skor"`
	Kelas                        *string    `gorm:"column:kelas" json:"kelas"`
	SkorInterpretasi             *string    `gorm:"column:skor_interpretasi" json:"skor_interpretasi"`
	Mortalitas                   *string    `gorm:"column:mortalitas" json:"mortalitas"`
	Rekomendasi                  *string    `gorm:"column:rekomendasi" json:"rekomendasi"`
	Nip                          string     `gorm:"column:nip" json:"nip"`
}

func (SkriningPneumoniaSeverityIndex) TableName() string {
	return "skrining_pneumonia_severity_index"
}
