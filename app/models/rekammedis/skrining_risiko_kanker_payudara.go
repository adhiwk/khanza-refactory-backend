package rekammedis

import "time"

// SkriningRisikoKankerPayudara tabel `skrining_risiko_kanker_payudara` (skrining risiko kanker payudara, RMSkriningRisikoKankerPayudara).
type SkriningRisikoKankerPayudara struct {
	NoRawat                string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                *time.Time `gorm:"column:tanggal" json:"tanggal"`
	FaktorRisikoAwal1      *string    `gorm:"column:faktor_risiko_awal1" json:"faktor_risiko_awal1"`
	NilaiRisikoAwal1       *string    `gorm:"column:nilai_risiko_awal1" json:"nilai_risiko_awal1"`
	FaktorRisikoAwal2      *string    `gorm:"column:faktor_risiko_awal2" json:"faktor_risiko_awal2"`
	NilaiRisikoAwal2       *string    `gorm:"column:nilai_risiko_awal2" json:"nilai_risiko_awal2"`
	FaktorRisikoAwal3      *string    `gorm:"column:faktor_risiko_awal3" json:"faktor_risiko_awal3"`
	NilaiRisikoAwal3       *string    `gorm:"column:nilai_risiko_awal3" json:"nilai_risiko_awal3"`
	FaktorRisikoAwal4      *string    `gorm:"column:faktor_risiko_awal4" json:"faktor_risiko_awal4"`
	NilaiRisikoAwal4       *string    `gorm:"column:nilai_risiko_awal4" json:"nilai_risiko_awal4"`
	FaktorRisikoAwal5      *string    `gorm:"column:faktor_risiko_awal5" json:"faktor_risiko_awal5"`
	NilaiRisikoAwal5       *string    `gorm:"column:nilai_risiko_awal5" json:"nilai_risiko_awal5"`
	FaktorRisikoAwal6      *string    `gorm:"column:faktor_risiko_awal6" json:"faktor_risiko_awal6"`
	NilaiRisikoAwal6       *string    `gorm:"column:nilai_risiko_awal6" json:"nilai_risiko_awal6"`
	FaktorRisikoAwal7      *string    `gorm:"column:faktor_risiko_awal7" json:"faktor_risiko_awal7"`
	NilaiRisikoAwal7       *string    `gorm:"column:nilai_risiko_awal7" json:"nilai_risiko_awal7"`
	FaktorRisikoAwal8      *string    `gorm:"column:faktor_risiko_awal8" json:"faktor_risiko_awal8"`
	NilaiRisikoAwal8       *string    `gorm:"column:nilai_risiko_awal8" json:"nilai_risiko_awal8"`
	FaktorRisikoAwal9      *string    `gorm:"column:faktor_risiko_awal9" json:"faktor_risiko_awal9"`
	NilaiRisikoAwal9       *string    `gorm:"column:nilai_risiko_awal9" json:"nilai_risiko_awal9"`
	FaktorRisikoAwal10     *string    `gorm:"column:faktor_risiko_awal10" json:"faktor_risiko_awal10"`
	NilaiRisikoAwal10      *string    `gorm:"column:nilai_risiko_awal10" json:"nilai_risiko_awal10"`
	FaktorRisikoAwal11     *string    `gorm:"column:faktor_risiko_awal11" json:"faktor_risiko_awal11"`
	NilaiRisikoAwal11      *string    `gorm:"column:nilai_risiko_awal11" json:"nilai_risiko_awal11"`
	FaktorRisikoAwal12     *string    `gorm:"column:faktor_risiko_awal12" json:"faktor_risiko_awal12"`
	NilaiRisikoAwal12      *string    `gorm:"column:nilai_risiko_awal12" json:"nilai_risiko_awal12"`
	FaktorRisikoAwal13     *string    `gorm:"column:faktor_risiko_awal13" json:"faktor_risiko_awal13"`
	NilaiRisikoAwal13      *string    `gorm:"column:nilai_risiko_awal13" json:"nilai_risiko_awal13"`
	FaktorRisikoAwal14     *string    `gorm:"column:faktor_risiko_awal14" json:"faktor_risiko_awal14"`
	NilaiRisikoAwal14      *string    `gorm:"column:nilai_risiko_awal14" json:"nilai_risiko_awal14"`
	FaktorRisikoTinggi1    *string    `gorm:"column:faktor_risiko_tinggi1" json:"faktor_risiko_tinggi1"`
	NilaiRisikoTinggi1     *string    `gorm:"column:nilai_risiko_tinggi1" json:"nilai_risiko_tinggi1"`
	FaktorRisikoTinggi2    *string    `gorm:"column:faktor_risiko_tinggi2" json:"faktor_risiko_tinggi2"`
	NilaiRisikoTinggi2     *string    `gorm:"column:nilai_risiko_tinggi2" json:"nilai_risiko_tinggi2"`
	FaktorRisikoTinggi3    *string    `gorm:"column:faktor_risiko_tinggi3" json:"faktor_risiko_tinggi3"`
	NilaiRisikoTinggi3     *string    `gorm:"column:nilai_risiko_tinggi3" json:"nilai_risiko_tinggi3"`
	FaktorRisikoTinggi4    *string    `gorm:"column:faktor_risiko_tinggi4" json:"faktor_risiko_tinggi4"`
	NilaiRisikoTinggi4     *string    `gorm:"column:nilai_risiko_tinggi4" json:"nilai_risiko_tinggi4"`
	FaktorRisikoTinggi5    *string    `gorm:"column:faktor_risiko_tinggi5" json:"faktor_risiko_tinggi5"`
	NilaiRisikoTinggi5     *string    `gorm:"column:nilai_risiko_tinggi5" json:"nilai_risiko_tinggi5"`
	FaktorRisikoTinggi6    *string    `gorm:"column:faktor_risiko_tinggi6" json:"faktor_risiko_tinggi6"`
	NilaiRisikoTinggi6     *string    `gorm:"column:nilai_risiko_tinggi6" json:"nilai_risiko_tinggi6"`
	FaktorRisikoTinggi7    *string    `gorm:"column:faktor_risiko_tinggi7" json:"faktor_risiko_tinggi7"`
	NilaiRisikoTinggi7     *string    `gorm:"column:nilai_risiko_tinggi7" json:"nilai_risiko_tinggi7"`
	FaktorRisikoTinggi8    *string    `gorm:"column:faktor_risiko_tinggi8" json:"faktor_risiko_tinggi8"`
	NilaiRisikoTinggi8     *string    `gorm:"column:nilai_risiko_tinggi8" json:"nilai_risiko_tinggi8"`
	FaktorRisikoTinggi9    *string    `gorm:"column:faktor_risiko_tinggi9" json:"faktor_risiko_tinggi9"`
	NilaiRisikoTinggi9     *string    `gorm:"column:nilai_risiko_tinggi9" json:"nilai_risiko_tinggi9"`
	FaktorRisikoTinggi10   *string    `gorm:"column:faktor_risiko_tinggi10" json:"faktor_risiko_tinggi10"`
	NilaiRisikoTinggi10    *string    `gorm:"column:nilai_risiko_tinggi10" json:"nilai_risiko_tinggi10"`
	FaktorRisikoTinggi11   *string    `gorm:"column:faktor_risiko_tinggi11" json:"faktor_risiko_tinggi11"`
	NilaiRisikoTinggi11    *string    `gorm:"column:nilai_risiko_tinggi11" json:"nilai_risiko_tinggi11"`
	FaktorRisikoTinggi12   *string    `gorm:"column:faktor_risiko_tinggi12" json:"faktor_risiko_tinggi12"`
	NilaiRisikoTinggi12    *string    `gorm:"column:nilai_risiko_tinggi12" json:"nilai_risiko_tinggi12"`
	FaktorRisikoTinggi13   *string    `gorm:"column:faktor_risiko_tinggi13" json:"faktor_risiko_tinggi13"`
	NilaiRisikoTinggi13    *string    `gorm:"column:nilai_risiko_tinggi13" json:"nilai_risiko_tinggi13"`
	FaktorKecurigaanGanas1 *string    `gorm:"column:faktor_kecurigaan_ganas1" json:"faktor_kecurigaan_ganas1"`
	NilaiKecurigaanGanas1  *string    `gorm:"column:nilai_kecurigaan_ganas1" json:"nilai_kecurigaan_ganas1"`
	FaktorKecurigaanGanas2 *string    `gorm:"column:faktor_kecurigaan_ganas2" json:"faktor_kecurigaan_ganas2"`
	NilaiKecurigaanGanas2  *string    `gorm:"column:nilai_kecurigaan_ganas2" json:"nilai_kecurigaan_ganas2"`
	FaktorKecurigaanGanas3 *string    `gorm:"column:faktor_kecurigaan_ganas3" json:"faktor_kecurigaan_ganas3"`
	NilaiKecurigaanGanas3  *string    `gorm:"column:nilai_kecurigaan_ganas3" json:"nilai_kecurigaan_ganas3"`
	FaktorKecurigaanGanas4 *string    `gorm:"column:faktor_kecurigaan_ganas4" json:"faktor_kecurigaan_ganas4"`
	NilaiKecurigaanGanas4  *string    `gorm:"column:nilai_kecurigaan_ganas4" json:"nilai_kecurigaan_ganas4"`
	FaktorKecurigaanGanas5 *string    `gorm:"column:faktor_kecurigaan_ganas5" json:"faktor_kecurigaan_ganas5"`
	NilaiKecurigaanGanas5  *string    `gorm:"column:nilai_kecurigaan_ganas5" json:"nilai_kecurigaan_ganas5"`
	FaktorKecurigaanGanas6 *string    `gorm:"column:faktor_kecurigaan_ganas6" json:"faktor_kecurigaan_ganas6"`
	NilaiKecurigaanGanas6  *string    `gorm:"column:nilai_kecurigaan_ganas6" json:"nilai_kecurigaan_ganas6"`
	FaktorKecurigaanGanas7 *string    `gorm:"column:faktor_kecurigaan_ganas7" json:"faktor_kecurigaan_ganas7"`
	NilaiKecurigaanGanas7  *string    `gorm:"column:nilai_kecurigaan_ganas7" json:"nilai_kecurigaan_ganas7"`
	FaktorKecurigaanGanas8 *string    `gorm:"column:faktor_kecurigaan_ganas8" json:"faktor_kecurigaan_ganas8"`
	NilaiKecurigaanGanas8  *string    `gorm:"column:nilai_kecurigaan_ganas8" json:"nilai_kecurigaan_ganas8"`
	TotalSkor              *string    `gorm:"column:total_skor" json:"total_skor"`
	HasilSadanis           *string    `gorm:"column:hasil_sadanis" json:"hasil_sadanis"`
	TindakLanjutSadanis    *string    `gorm:"column:tindak_lanjut_sadanis" json:"tindak_lanjut_sadanis"`
	HasilSkrining          *string    `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Keterangan             string     `gorm:"column:keterangan" json:"keterangan"`
	Nip                    string     `gorm:"column:nip" json:"nip"`
}

func (SkriningRisikoKankerPayudara) TableName() string {
	return "skrining_risiko_kanker_payudara"
}
