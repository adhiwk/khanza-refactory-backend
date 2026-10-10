package rekammedis

import "time"

// AsuhanGizi tabel `asuhan_gizi` (asuhan gizi, RMDataAsuhanGizi).
type AsuhanGizi struct {
	NoRawat             string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal             *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	AntropometriBb      *string    `gorm:"column:antropometri_bb" json:"antropometri_bb"`
	AntropometriTb      *string    `gorm:"column:antropometri_tb" json:"antropometri_tb"`
	AntropometriImt     *string    `gorm:"column:antropometri_imt" json:"antropometri_imt"`
	AntropometriLla     *string    `gorm:"column:antropometri_lla" json:"antropometri_lla"`
	AntropometriTl      *string    `gorm:"column:antropometri_tl" json:"antropometri_tl"`
	AntropometriUlna    string     `gorm:"column:antropometri_ulna" json:"antropometri_ulna"`
	AntropometriBbideal string     `gorm:"column:antropometri_bbideal" json:"antropometri_bbideal"`
	AntropometriBbperu  string     `gorm:"column:antropometri_bbperu" json:"antropometri_bbperu"`
	AntropometriTbperu  string     `gorm:"column:antropometri_tbperu" json:"antropometri_tbperu"`
	AntropometriBbpertb string     `gorm:"column:antropometri_bbpertb" json:"antropometri_bbpertb"`
	AntropometriLlaperu string     `gorm:"column:antropometri_llaperu" json:"antropometri_llaperu"`
	Biokimia            *string    `gorm:"column:biokimia" json:"biokimia"`
	FisikKlinis         *string    `gorm:"column:fisik_klinis" json:"fisik_klinis"`
	AlergiTelur         *string    `gorm:"column:alergi_telur" json:"alergi_telur"`
	AlergiSusuSapi      *string    `gorm:"column:alergi_susu_sapi" json:"alergi_susu_sapi"`
	AlergiKacang        *string    `gorm:"column:alergi_kacang" json:"alergi_kacang"`
	AlergiGluten        *string    `gorm:"column:alergi_gluten" json:"alergi_gluten"`
	AlergiUdang         *string    `gorm:"column:alergi_udang" json:"alergi_udang"`
	AlergiIkan          *string    `gorm:"column:alergi_ikan" json:"alergi_ikan"`
	AlergiHazelnut      *string    `gorm:"column:alergi_hazelnut" json:"alergi_hazelnut"`
	PolaMakan           *string    `gorm:"column:pola_makan" json:"pola_makan"`
	RiwayatPersonal     *string    `gorm:"column:riwayat_personal" json:"riwayat_personal"`
	Diagnosis           *string    `gorm:"column:diagnosis" json:"diagnosis"`
	IntervensiGizi      *string    `gorm:"column:intervensi_gizi" json:"intervensi_gizi"`
	MonitoringEvaluasi  *string    `gorm:"column:monitoring_evaluasi" json:"monitoring_evaluasi"`
	Nip                 string     `gorm:"column:nip" json:"nip"`
}

func (AsuhanGizi) TableName() string {
	return "asuhan_gizi"
}
