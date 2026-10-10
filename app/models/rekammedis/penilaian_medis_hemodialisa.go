package rekammedis

import "time"

// PenilaianMedisHemodialisa tabel `penilaian_medis_hemodialisa` (penilaian awal medis hemodialisa, RMPenilaianAwalMedisHemodialisa).
type PenilaianMedisHemodialisa struct {
	NoRawat                       string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                       *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter                      *string    `gorm:"column:kd_dokter" json:"kd_dokter"`
	Anamnesis                     *string    `gorm:"column:anamnesis" json:"anamnesis"`
	Hubungan                      *string    `gorm:"column:hubungan" json:"hubungan"`
	Ruangan                       *string    `gorm:"column:ruangan" json:"ruangan"`
	Alergi                        *string    `gorm:"column:alergi" json:"alergi"`
	Nyeri                         *string    `gorm:"column:nyeri" json:"nyeri"`
	StatusNutrisi                 *string    `gorm:"column:status_nutrisi" json:"status_nutrisi"`
	Hipertensi                    *string    `gorm:"column:hipertensi" json:"hipertensi"`
	KeteranganHipertensi          *string    `gorm:"column:keterangan_hipertensi" json:"keterangan_hipertensi"`
	Diabetes                      *string    `gorm:"column:diabetes" json:"diabetes"`
	KeteranganDiabetes            *string    `gorm:"column:keterangan_diabetes" json:"keterangan_diabetes"`
	BatuSaluranKemih              *string    `gorm:"column:batu_saluran_kemih" json:"batu_saluran_kemih"`
	KeteranganBatuSaluranKemih    *string    `gorm:"column:keterangan_batu_saluran_kemih" json:"keterangan_batu_saluran_kemih"`
	OperasiSaluranKemih           *string    `gorm:"column:operasi_saluran_kemih" json:"operasi_saluran_kemih"`
	KeteranganOperasiSaluranKemih *string    `gorm:"column:keterangan_operasi_saluran_kemih" json:"keterangan_operasi_saluran_kemih"`
	InfeksiSaluranKemih           *string    `gorm:"column:infeksi_saluran_kemih" json:"infeksi_saluran_kemih"`
	KeteranganInfeksiSaluranKemih *string    `gorm:"column:keterangan_infeksi_saluran_kemih" json:"keterangan_infeksi_saluran_kemih"`
	BengkakSeluruhTubuh           *string    `gorm:"column:bengkak_seluruh_tubuh" json:"bengkak_seluruh_tubuh"`
	KeteranganBengkakSeluruhTubuh *string    `gorm:"column:keterangan_bengkak_seluruh_tubuh" json:"keterangan_bengkak_seluruh_tubuh"`
	UrinBerdarah                  *string    `gorm:"column:urin_berdarah" json:"urin_berdarah"`
	KeteranganUrinBerdarah        *string    `gorm:"column:keterangan_urin_berdarah" json:"keterangan_urin_berdarah"`
	PenyakitGinjalLaom            *string    `gorm:"column:penyakit_ginjal_laom" json:"penyakit_ginjal_laom"`
	KeteranganPenyakitGinjalLaom  *string    `gorm:"column:keterangan_penyakit_ginjal_laom" json:"keterangan_penyakit_ginjal_laom"`
	PenyakitLain                  *string    `gorm:"column:penyakit_lain" json:"penyakit_lain"`
	KeteranganPenyakitLain        *string    `gorm:"column:keterangan_penyakit_lain" json:"keterangan_penyakit_lain"`
	KonsumsiObatNefro             *string    `gorm:"column:konsumsi_obat_nefro" json:"konsumsi_obat_nefro"`
	KeteranganKonsumsiObatNefro   *string    `gorm:"column:keterangan_konsumsi_obat_nefro" json:"keterangan_konsumsi_obat_nefro"`
	DialisisPertama               *time.Time `gorm:"column:dialisis_pertama" json:"dialisis_pertama"`
	PernahCpad                    string     `gorm:"column:pernah_cpad" json:"pernah_cpad"`
	TanggalCpad                   *time.Time `gorm:"column:tanggal_cpad" json:"tanggal_cpad"`
	PernahTransplantasi           string     `gorm:"column:pernah_transplantasi" json:"pernah_transplantasi"`
	TanggalTransplantasi          *time.Time `gorm:"column:tanggal_transplantasi" json:"tanggal_transplantasi"`
	KeadaanUmum                   string     `gorm:"column:keadaan_umum" json:"keadaan_umum"`
	Kesadaran                     string     `gorm:"column:kesadaran" json:"kesadaran"`
	Nadi                          string     `gorm:"column:nadi" json:"nadi"`
	Bb                            string     `gorm:"column:bb" json:"bb"`
	Td                            string     `gorm:"column:td" json:"td"`
	Suhu                          string     `gorm:"column:suhu" json:"suhu"`
	Napas                         string     `gorm:"column:napas" json:"napas"`
	Tb                            string     `gorm:"column:tb" json:"tb"`
	Hepatomegali                  string     `gorm:"column:hepatomegali" json:"hepatomegali"`
	Splenomegali                  string     `gorm:"column:splenomegali" json:"splenomegali"`
	Ascites                       string     `gorm:"column:ascites" json:"ascites"`
	Edema                         string     `gorm:"column:edema" json:"edema"`
	Whezzing                      string     `gorm:"column:whezzing" json:"whezzing"`
	Ronchi                        string     `gorm:"column:ronchi" json:"ronchi"`
	Ikterik                       string     `gorm:"column:ikterik" json:"ikterik"`
	TekananVena                   string     `gorm:"column:tekanan_vena" json:"tekanan_vena"`
	Anemia                        string     `gorm:"column:anemia" json:"anemia"`
	Kardiomegali                  string     `gorm:"column:kardiomegali" json:"kardiomegali"`
	Bising                        string     `gorm:"column:bising" json:"bising"`
	Thorax                        string     `gorm:"column:thorax" json:"thorax"`
	TanggalThorax                 *time.Time `gorm:"column:tanggal_thorax" json:"tanggal_thorax"`
	Ekg                           string     `gorm:"column:ekg" json:"ekg"`
	TanggalEkg                    *time.Time `gorm:"column:tanggal_ekg" json:"tanggal_ekg"`
	Bno                           string     `gorm:"column:bno" json:"bno"`
	TanggalBno                    *time.Time `gorm:"column:tanggal_bno" json:"tanggal_bno"`
	Usg                           string     `gorm:"column:usg" json:"usg"`
	TanggalUsg                    *time.Time `gorm:"column:tanggal_usg" json:"tanggal_usg"`
	Renogram                      string     `gorm:"column:renogram" json:"renogram"`
	TanggalRenogram               *time.Time `gorm:"column:tanggal_renogram" json:"tanggal_renogram"`
	Biopsi                        string     `gorm:"column:biopsi" json:"biopsi"`
	TanggalBiopsi                 *time.Time `gorm:"column:tanggal_biopsi" json:"tanggal_biopsi"`
	Ctscan                        string     `gorm:"column:ctscan" json:"ctscan"`
	TanggalCtscan                 *time.Time `gorm:"column:tanggal_ctscan" json:"tanggal_ctscan"`
	Arteriografi                  string     `gorm:"column:arteriografi" json:"arteriografi"`
	TanggalArteriografi           *time.Time `gorm:"column:tanggal_arteriografi" json:"tanggal_arteriografi"`
	KulturUrin                    string     `gorm:"column:kultur_urin" json:"kultur_urin"`
	TanggalKulturUrin             *time.Time `gorm:"column:tanggal_kultur_urin" json:"tanggal_kultur_urin"`
	Laborat                       string     `gorm:"column:laborat" json:"laborat"`
	TanggalLaborat                *time.Time `gorm:"column:tanggal_laborat" json:"tanggal_laborat"`
	Hematokrit                    string     `gorm:"column:hematokrit" json:"hematokrit"`
	Hemoglobin                    string     `gorm:"column:hemoglobin" json:"hemoglobin"`
	Leukosit                      string     `gorm:"column:leukosit" json:"leukosit"`
	Trombosit                     string     `gorm:"column:trombosit" json:"trombosit"`
	HitungJenis                   string     `gorm:"column:hitung_jenis" json:"hitung_jenis"`
	Ureum                         string     `gorm:"column:ureum" json:"ureum"`
	UrinLengkap                   string     `gorm:"column:urin_lengkap" json:"urin_lengkap"`
	Kreatinin                     string     `gorm:"column:kreatinin" json:"kreatinin"`
	Cct                           string     `gorm:"column:cct" json:"cct"`
	Sgot                          string     `gorm:"column:sgot" json:"sgot"`
	Sgpt                          string     `gorm:"column:sgpt" json:"sgpt"`
	Ct                            string     `gorm:"column:ct" json:"ct"`
	AsamUrat                      string     `gorm:"column:asam_urat" json:"asam_urat"`
	Hbsag                         string     `gorm:"column:hbsag" json:"hbsag"`
	AntiHcv                       string     `gorm:"column:anti_hcv" json:"anti_hcv"`
	Edukasi                       string     `gorm:"column:edukasi" json:"edukasi"`
}

func (PenilaianMedisHemodialisa) TableName() string {
	return "penilaian_medis_hemodialisa"
}
