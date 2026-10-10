package rekammedis

import "time"

// SkriningAdiksiNikotin tabel `skrining_adiksi_nikotin` (skrining adiksi nikotin, RMSkriningAdiksiNikotin).
type SkriningAdiksiNikotin struct {
	NoRawat                 string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                 *time.Time `gorm:"column:tanggal" json:"tanggal"`
	RokokDihisap            *string    `gorm:"column:rokok_dihisap" json:"rokok_dihisap"`
	NilaiRokokDihisap       *int       `gorm:"column:nilai_rokok_dihisap" json:"nilai_rokok_dihisap"`
	MenyalakanRokok         *string    `gorm:"column:menyalakan_rokok" json:"menyalakan_rokok"`
	NilaiMenyalakanRokok    *int       `gorm:"column:nilai_menyalakan_rokok" json:"nilai_menyalakan_rokok"`
	TidakRela               *string    `gorm:"column:tidak_rela" json:"tidak_rela"`
	NilaiTidakRela          *int       `gorm:"column:nilai_tidak_rela" json:"nilai_tidak_rela"`
	JamPertama              *string    `gorm:"column:jam_pertama" json:"jam_pertama"`
	NilaiJamPertama         *int       `gorm:"column:nilai_jam_pertama" json:"nilai_jam_pertama"`
	RasaIngin               *string    `gorm:"column:rasa_ingin" json:"rasa_ingin"`
	NilaiRasaIngin          *int       `gorm:"column:nilai_rasa_ingin" json:"nilai_rasa_ingin"`
	SakitBerat              *string    `gorm:"column:sakit_berat" json:"sakit_berat"`
	NilaiSakitBerat         *int       `gorm:"column:nilai_sakit_berat" json:"nilai_sakit_berat"`
	NilaiTotal              *int       `gorm:"column:nilai_total" json:"nilai_total"`
	KeteranganHasilSkrining *string    `gorm:"column:keterangan_hasil_skrining" json:"keterangan_hasil_skrining"`
	SkalaMotivasi           *string    `gorm:"column:skala_motivasi" json:"skala_motivasi"`
	Nip                     string     `gorm:"column:nip" json:"nip"`
}

func (SkriningAdiksiNikotin) TableName() string {
	return "skrining_adiksi_nikotin"
}
