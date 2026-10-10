package rekammedis

import "time"

// ChecklistPemberianFibrinolitik tabel `checklist_pemberian_fibrinolitik` (checklist pemberian fibrinolitik, RMChecklistPemberianFibrinolitik).
type ChecklistPemberianFibrinolitik struct {
	NoRawat                     string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                     *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip                         string     `gorm:"column:nip" json:"nip"`
	KontraIndikasi1             *string    `gorm:"column:kontra_indikasi1" json:"kontra_indikasi1"`
	KeteranganKontraIndikasi1   *string    `gorm:"column:keterangan_kontra_indikasi1" json:"keterangan_kontra_indikasi1"`
	KontraIndikasi2             *string    `gorm:"column:kontra_indikasi2" json:"kontra_indikasi2"`
	KeteranganKontraIndikasi2   *string    `gorm:"column:keterangan_kontra_indikasi2" json:"keterangan_kontra_indikasi2"`
	KontraIndikasi3             *string    `gorm:"column:kontra_indikasi3" json:"kontra_indikasi3"`
	KeteranganKontraIndikasi3   *string    `gorm:"column:keterangan_kontra_indikasi3" json:"keterangan_kontra_indikasi3"`
	KontraIndikasi4             *string    `gorm:"column:kontra_indikasi4" json:"kontra_indikasi4"`
	KeteranganKontraIndikasi4   *string    `gorm:"column:keterangan_kontra_indikasi4" json:"keterangan_kontra_indikasi4"`
	KontraIndikasi5             *string    `gorm:"column:kontra_indikasi5" json:"kontra_indikasi5"`
	KeteranganKontraIndikasi5   *string    `gorm:"column:keterangan_kontra_indikasi5" json:"keterangan_kontra_indikasi5"`
	KontraIndikasi6             *string    `gorm:"column:kontra_indikasi6" json:"kontra_indikasi6"`
	KeteranganKontraIndikasi6   *string    `gorm:"column:keterangan_kontra_indikasi6" json:"keterangan_kontra_indikasi6"`
	KontraIndikasi7             *string    `gorm:"column:kontra_indikasi7" json:"kontra_indikasi7"`
	KeteranganKontraIndikasi7   *string    `gorm:"column:keterangan_kontra_indikasi7" json:"keterangan_kontra_indikasi7"`
	KontraIndikasi8             *string    `gorm:"column:kontra_indikasi8" json:"kontra_indikasi8"`
	KeteranganKontraIndikasi8   *string    `gorm:"column:keterangan_kontra_indikasi8" json:"keterangan_kontra_indikasi8"`
	KontraIndikasi9             *string    `gorm:"column:kontra_indikasi9" json:"kontra_indikasi9"`
	KeteranganKontraIndikasi9   *string    `gorm:"column:keterangan_kontra_indikasi9" json:"keterangan_kontra_indikasi9"`
	KontraIndikasi10            *string    `gorm:"column:kontra_indikasi10" json:"kontra_indikasi10"`
	KeteranganKontraIndikasi10  *string    `gorm:"column:keterangan_kontra_indikasi10" json:"keterangan_kontra_indikasi10"`
	RisikoTinggi1               *string    `gorm:"column:risiko_tinggi1" json:"risiko_tinggi1"`
	KeteranganRisikoTinggi1     *string    `gorm:"column:keterangan_risiko_tinggi1" json:"keterangan_risiko_tinggi1"`
	RisikoTinggi2               *string    `gorm:"column:risiko_tinggi2" json:"risiko_tinggi2"`
	KeteranganRisikoTinggi2     *string    `gorm:"column:keterangan_risiko_tinggi2" json:"keterangan_risiko_tinggi2"`
	RisikoTinggi3               *string    `gorm:"column:risiko_tinggi3" json:"risiko_tinggi3"`
	KeteranganRisikoTinggi3     *string    `gorm:"column:keterangan_risiko_tinggi3" json:"keterangan_risiko_tinggi3"`
	RisikoTinggi4               *string    `gorm:"column:risiko_tinggi4" json:"risiko_tinggi4"`
	KeteranganRisikoTinggi4     *string    `gorm:"column:keterangan_risiko_tinggi4" json:"keterangan_risiko_tinggi4"`
	RisikoTinggi5               *string    `gorm:"column:risiko_tinggi5" json:"risiko_tinggi5"`
	KeteranganRisikoTinggi5     *string    `gorm:"column:keterangan_risiko_tinggi5" json:"keterangan_risiko_tinggi5"`
	Kesimpulan                  *string    `gorm:"column:kesimpulan" json:"kesimpulan"`
	PersyaratanEkgPreStreptase  *string    `gorm:"column:persyaratan_ekg_pre_streptase" json:"persyaratan_ekg_pre_streptase"`
	PersyaratanEkgPostStreptase *string    `gorm:"column:persyaratan_ekg_post_streptase" json:"persyaratan_ekg_post_streptase"`
	CekTroponin                 *string    `gorm:"column:cek_troponin" json:"cek_troponin"`
}

func (ChecklistPemberianFibrinolitik) TableName() string {
	return "checklist_pemberian_fibrinolitik"
}
