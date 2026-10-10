package rekammedis

import "time"

// TimeoutSebelumInsisi tabel `timeout_sebelum_insisi` (time out sebelum insisi, RMTimeOutSebelumInsisi).
type TimeoutSebelumInsisi struct {
	NoRawat                   string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                   *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Sncn                      string     `gorm:"column:sncn" json:"sncn"`
	Tindakan                  string     `gorm:"column:tindakan" json:"tindakan"`
	KdDokterBedah             string     `gorm:"column:kd_dokter_bedah" json:"kd_dokter_bedah"`
	KdDokterAnestesi          string     `gorm:"column:kd_dokter_anestesi" json:"kd_dokter_anestesi"`
	VerbalIdentitas           *string    `gorm:"column:verbal_identitas" json:"verbal_identitas"`
	VerbalTindakan            *string    `gorm:"column:verbal_tindakan" json:"verbal_tindakan"`
	VerbalAreaInsisi          *string    `gorm:"column:verbal_area_insisi" json:"verbal_area_insisi"`
	PenandaanAreaOperasi      *string    `gorm:"column:penandaan_area_operasi" json:"penandaan_area_operasi"`
	LamaOperasi               string     `gorm:"column:lama_operasi" json:"lama_operasi"`
	PenayanganRadiologi       *string    `gorm:"column:penayangan_radiologi" json:"penayangan_radiologi"`
	PenayanganCtscan          *string    `gorm:"column:penayangan_ctscan" json:"penayangan_ctscan"`
	PenayanganMri             *string    `gorm:"column:penayangan_mri" json:"penayangan_mri"`
	AntibiotikProfilaks       *string    `gorm:"column:antibiotik_profilaks" json:"antibiotik_profilaks"`
	NamaAntibiotik            string     `gorm:"column:nama_antibiotik" json:"nama_antibiotik"`
	JamPemberian              string     `gorm:"column:jam_pemberian" json:"jam_pemberian"`
	AntisipasiKehilanganDarah string     `gorm:"column:antisipasi_kehilangan_darah" json:"antisipasi_kehilangan_darah"`
	HalKhusus                 *string    `gorm:"column:hal_khusus" json:"hal_khusus"`
	HalKhususDiperhatikan     string     `gorm:"column:hal_khusus_diperhatikan" json:"hal_khusus_diperhatikan"`
	TanggalSteril             *time.Time `gorm:"column:tanggal_steril" json:"tanggal_steril"`
	PetujukSterilisasi        *string    `gorm:"column:petujuk_sterilisasi" json:"petujuk_sterilisasi"`
	VerifikasiPreoperatif     *string    `gorm:"column:verifikasi_preoperatif" json:"verifikasi_preoperatif"`
	NipPerawatOk              *string    `gorm:"column:nip_perawat_ok" json:"nip_perawat_ok"`
}

func (TimeoutSebelumInsisi) TableName() string {
	return "timeout_sebelum_insisi"
}
