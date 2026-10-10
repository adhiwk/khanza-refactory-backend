package rekammedis

import "time"

// SkriningFrailtySyndrome tabel `skrining_frailty_syndrome` (skrining frailty syndrome, RMSkriningFrailtySyndrome).
type SkriningFrailtySyndrome struct {
	NoRawat                 string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                 *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Resistensi              *string    `gorm:"column:resistensi" json:"resistensi"`
	NilaiResistensi         *int       `gorm:"column:nilai_resistensi" json:"nilai_resistensi"`
	Aktivitas               *string    `gorm:"column:aktivitas" json:"aktivitas"`
	NilaiAktivitas          *int       `gorm:"column:nilai_aktivitas" json:"nilai_aktivitas"`
	PenyakitTidakPernah     *string    `gorm:"column:penyakit_tidak_pernah" json:"penyakit_tidak_pernah"`
	PenyakitKanker          *string    `gorm:"column:penyakit_kanker" json:"penyakit_kanker"`
	PenyakitGagalJantung    *string    `gorm:"column:penyakit_gagal_jantung" json:"penyakit_gagal_jantung"`
	PenyakitGinjal          *string    `gorm:"column:penyakit_ginjal" json:"penyakit_ginjal"`
	PenyakitNyeriDada       *string    `gorm:"column:penyakit_nyeri_dada" json:"penyakit_nyeri_dada"`
	PenyakitSeranganJantung *string    `gorm:"column:penyakit_serangan_jantung" json:"penyakit_serangan_jantung"`
	PenyakitStroke          *string    `gorm:"column:penyakit_stroke" json:"penyakit_stroke"`
	PenyakitAsma            *string    `gorm:"column:penyakit_asma" json:"penyakit_asma"`
	PenyakitNyeriSendi      *string    `gorm:"column:penyakit_nyeri_sendi" json:"penyakit_nyeri_sendi"`
	PenyakitParuKronis      *string    `gorm:"column:penyakit_paru_kronis" json:"penyakit_paru_kronis"`
	PenyakitHipertensi      *string    `gorm:"column:penyakit_hipertensi" json:"penyakit_hipertensi"`
	PenyakitDiabetes        *string    `gorm:"column:penyakit_diabetes" json:"penyakit_diabetes"`
	NilaiPenyakit           *int       `gorm:"column:nilai_penyakit" json:"nilai_penyakit"`
	UsahaBerjalan           *string    `gorm:"column:usaha_berjalan" json:"usaha_berjalan"`
	NilaiUsahaBerjalan      *int       `gorm:"column:nilai_usaha_berjalan" json:"nilai_usaha_berjalan"`
	BeratBadan              *string    `gorm:"column:berat_badan" json:"berat_badan"`
	NilaiBeratBadan         *int       `gorm:"column:nilai_berat_badan" json:"nilai_berat_badan"`
	NilaiTotal              *int       `gorm:"column:nilai_total" json:"nilai_total"`
	HasilSkrining           *string    `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Keterangan              *string    `gorm:"column:keterangan" json:"keterangan"`
	Nip                     string     `gorm:"column:nip" json:"nip"`
}

func (SkriningFrailtySyndrome) TableName() string {
	return "skrining_frailty_syndrome"
}
