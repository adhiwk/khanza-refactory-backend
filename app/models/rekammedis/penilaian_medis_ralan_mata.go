package rekammedis

import "time"

// PenilaianMedisRalanMata tabel `penilaian_medis_ralan_mata` (penilaian awal medis ralan mata, RMPenilaianAwalMedisRalanMata).
type PenilaianMedisRalanMata struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal      *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter     string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Anamnesis    string     `gorm:"column:anamnesis" json:"anamnesis"`
	Hubungan     string     `gorm:"column:hubungan" json:"hubungan"`
	KeluhanUtama string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rps          string     `gorm:"column:rps" json:"rps"`
	Rpd          string     `gorm:"column:rpd" json:"rpd"`
	Rpo          string     `gorm:"column:rpo" json:"rpo"`
	Alergi       string     `gorm:"column:alergi" json:"alergi"`
	Status       string     `gorm:"column:status" json:"status"`
	Td           string     `gorm:"column:td" json:"td"`
	Nadi         string     `gorm:"column:nadi" json:"nadi"`
	Rr           string     `gorm:"column:rr" json:"rr"`
	Suhu         string     `gorm:"column:suhu" json:"suhu"`
	Nyeri        string     `gorm:"column:nyeri" json:"nyeri"`
	Bb           string     `gorm:"column:bb" json:"bb"`
	Visuskanan   string     `gorm:"column:visuskanan" json:"visuskanan"`
	Visuskiri    string     `gorm:"column:visuskiri" json:"visuskiri"`
	Cckanan      string     `gorm:"column:cckanan" json:"cckanan"`
	Cckiri       string     `gorm:"column:cckiri" json:"cckiri"`
	Palkanan     string     `gorm:"column:palkanan" json:"palkanan"`
	Palkiri      string     `gorm:"column:palkiri" json:"palkiri"`
	Conkanan     string     `gorm:"column:conkanan" json:"conkanan"`
	Conkiri      string     `gorm:"column:conkiri" json:"conkiri"`
	Corneakanan  string     `gorm:"column:corneakanan" json:"corneakanan"`
	Corneakiri   string     `gorm:"column:corneakiri" json:"corneakiri"`
	Coakanan     string     `gorm:"column:coakanan" json:"coakanan"`
	Coakiri      string     `gorm:"column:coakiri" json:"coakiri"`
	Pupilkanan   string     `gorm:"column:pupilkanan" json:"pupilkanan"`
	Pupilkiri    string     `gorm:"column:pupilkiri" json:"pupilkiri"`
	Lensakanan   string     `gorm:"column:lensakanan" json:"lensakanan"`
	Lensakiri    string     `gorm:"column:lensakiri" json:"lensakiri"`
	Funduskanan  string     `gorm:"column:funduskanan" json:"funduskanan"`
	Funduskiri   string     `gorm:"column:funduskiri" json:"funduskiri"`
	Papilkanan   string     `gorm:"column:papilkanan" json:"papilkanan"`
	Papilkiri    string     `gorm:"column:papilkiri" json:"papilkiri"`
	Retinakanan  string     `gorm:"column:retinakanan" json:"retinakanan"`
	Retinakiri   string     `gorm:"column:retinakiri" json:"retinakiri"`
	Makulakanan  string     `gorm:"column:makulakanan" json:"makulakanan"`
	Makulakiri   string     `gorm:"column:makulakiri" json:"makulakiri"`
	Tiokanan     string     `gorm:"column:tiokanan" json:"tiokanan"`
	Tiokiri      string     `gorm:"column:tiokiri" json:"tiokiri"`
	Mbokanan     string     `gorm:"column:mbokanan" json:"mbokanan"`
	Mbokiri      string     `gorm:"column:mbokiri" json:"mbokiri"`
	Lab          string     `gorm:"column:lab" json:"lab"`
	Rad          string     `gorm:"column:rad" json:"rad"`
	Penunjang    string     `gorm:"column:penunjang" json:"penunjang"`
	Tes          string     `gorm:"column:tes" json:"tes"`
	Pemeriksaan  string     `gorm:"column:pemeriksaan" json:"pemeriksaan"`
	Diagnosis    string     `gorm:"column:diagnosis" json:"diagnosis"`
	Diagnosisbdg string     `gorm:"column:diagnosisbdg" json:"diagnosisbdg"`
	Permasalahan string     `gorm:"column:permasalahan" json:"permasalahan"`
	Terapi       string     `gorm:"column:terapi" json:"terapi"`
	Tindakan     string     `gorm:"column:tindakan" json:"tindakan"`
	Edukasi      string     `gorm:"column:edukasi" json:"edukasi"`
}

func (PenilaianMedisRalanMata) TableName() string {
	return "penilaian_medis_ralan_mata"
}
