package rekammedis

import "time"

// PenilaianMedisRalanTht tabel `penilaian_medis_ralan_tht` (penilaian awal medis ralan THT, RMPenilaianAwalMedisRalanTHT).
type PenilaianMedisRalanTht struct {
	NoRawat          string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal          *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter         string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Anamnesis        string     `gorm:"column:anamnesis" json:"anamnesis"`
	Hubungan         string     `gorm:"column:hubungan" json:"hubungan"`
	KeluhanUtama     string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rps              string     `gorm:"column:rps" json:"rps"`
	Rpd              string     `gorm:"column:rpd" json:"rpd"`
	Rpo              string     `gorm:"column:rpo" json:"rpo"`
	Alergi           string     `gorm:"column:alergi" json:"alergi"`
	Td               string     `gorm:"column:td" json:"td"`
	Nadi             string     `gorm:"column:nadi" json:"nadi"`
	Rr               string     `gorm:"column:rr" json:"rr"`
	Suhu             string     `gorm:"column:suhu" json:"suhu"`
	Bb               string     `gorm:"column:bb" json:"bb"`
	Tb               string     `gorm:"column:tb" json:"tb"`
	Nyeri            string     `gorm:"column:nyeri" json:"nyeri"`
	StatusNutrisi    string     `gorm:"column:status_nutrisi" json:"status_nutrisi"`
	Kondisi          string     `gorm:"column:kondisi" json:"kondisi"`
	KetLokalis       string     `gorm:"column:ket_lokalis" json:"ket_lokalis"`
	Lab              string     `gorm:"column:lab" json:"lab"`
	Rad              string     `gorm:"column:rad" json:"rad"`
	TesPendengaran   string     `gorm:"column:tes_pendengaran" json:"tes_pendengaran"`
	Penunjang        string     `gorm:"column:penunjang" json:"penunjang"`
	Diagnosis        string     `gorm:"column:diagnosis" json:"diagnosis"`
	Diagnosisbanding string     `gorm:"column:diagnosisbanding" json:"diagnosisbanding"`
	Permasalahan     string     `gorm:"column:permasalahan" json:"permasalahan"`
	Terapi           string     `gorm:"column:terapi" json:"terapi"`
	Tindakan         string     `gorm:"column:tindakan" json:"tindakan"`
	Tatalaksana      string     `gorm:"column:tatalaksana" json:"tatalaksana"`
	Edukasi          string     `gorm:"column:edukasi" json:"edukasi"`
}

func (PenilaianMedisRalanTht) TableName() string {
	return "penilaian_medis_ralan_tht"
}
