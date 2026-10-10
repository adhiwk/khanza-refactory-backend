package rekammedis

import "time"

// PenilaianMedisRanapKandungan tabel `penilaian_medis_ranap_kandungan` (penilaian awal medis ranap kandungan, RMPenilaianAwalMedisRanapKandungan).
type PenilaianMedisRanapKandungan struct {
	NoRawat      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal      *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter     string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Anamnesis    string     `gorm:"column:anamnesis" json:"anamnesis"`
	Hubungan     string     `gorm:"column:hubungan" json:"hubungan"`
	KeluhanUtama string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rps          string     `gorm:"column:rps" json:"rps"`
	Rpd          string     `gorm:"column:rpd" json:"rpd"`
	Rpk          string     `gorm:"column:rpk" json:"rpk"`
	Rpo          string     `gorm:"column:rpo" json:"rpo"`
	Alergi       string     `gorm:"column:alergi" json:"alergi"`
	Keadaan      string     `gorm:"column:keadaan" json:"keadaan"`
	Gcs          string     `gorm:"column:gcs" json:"gcs"`
	Kesadaran    string     `gorm:"column:kesadaran" json:"kesadaran"`
	Td           string     `gorm:"column:td" json:"td"`
	Nadi         string     `gorm:"column:nadi" json:"nadi"`
	Rr           string     `gorm:"column:rr" json:"rr"`
	Suhu         string     `gorm:"column:suhu" json:"suhu"`
	Spo          string     `gorm:"column:spo" json:"spo"`
	Bb           string     `gorm:"column:bb" json:"bb"`
	Tb           string     `gorm:"column:tb" json:"tb"`
	Kepala       string     `gorm:"column:kepala" json:"kepala"`
	Mata         string     `gorm:"column:mata" json:"mata"`
	Gigi         string     `gorm:"column:gigi" json:"gigi"`
	Tht          string     `gorm:"column:tht" json:"tht"`
	Thoraks      string     `gorm:"column:thoraks" json:"thoraks"`
	Jantung      string     `gorm:"column:jantung" json:"jantung"`
	Paru         string     `gorm:"column:paru" json:"paru"`
	Abdomen      string     `gorm:"column:abdomen" json:"abdomen"`
	Genital      string     `gorm:"column:genital" json:"genital"`
	Ekstremitas  string     `gorm:"column:ekstremitas" json:"ekstremitas"`
	Kulit        string     `gorm:"column:kulit" json:"kulit"`
	KetFisik     string     `gorm:"column:ket_fisik" json:"ket_fisik"`
	Tfu          string     `gorm:"column:tfu" json:"tfu"`
	Tbj          string     `gorm:"column:tbj" json:"tbj"`
	His          string     `gorm:"column:his" json:"his"`
	Kontraksi    string     `gorm:"column:kontraksi" json:"kontraksi"`
	Djj          string     `gorm:"column:djj" json:"djj"`
	Inspeksi     string     `gorm:"column:inspeksi" json:"inspeksi"`
	Inspekulo    string     `gorm:"column:inspekulo" json:"inspekulo"`
	Vt           string     `gorm:"column:vt" json:"vt"`
	Rt           string     `gorm:"column:rt" json:"rt"`
	Ultra        string     `gorm:"column:ultra" json:"ultra"`
	Kardio       string     `gorm:"column:kardio" json:"kardio"`
	Lab          string     `gorm:"column:lab" json:"lab"`
	Diagnosis    string     `gorm:"column:diagnosis" json:"diagnosis"`
	Tata         string     `gorm:"column:tata" json:"tata"`
	Edukasi      string     `gorm:"column:edukasi" json:"edukasi"`
}

func (PenilaianMedisRanapKandungan) TableName() string {
	return "penilaian_medis_ranap_kandungan"
}
