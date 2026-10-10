package rekammedis

import "time"

// PenilaianMedisRalanPenyakitDalam tabel `penilaian_medis_ralan_penyakit_dalam` (penilaian awal medis ralan penyakit dalam, RMPenilaianAwalMedisRalanPenyakitDalam).
type PenilaianMedisRalanPenyakitDalam struct {
	NoRawat               string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal               *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter              string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Anamnesis             string     `gorm:"column:anamnesis" json:"anamnesis"`
	Hubungan              string     `gorm:"column:hubungan" json:"hubungan"`
	KeluhanUtama          string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rps                   string     `gorm:"column:rps" json:"rps"`
	Rpd                   string     `gorm:"column:rpd" json:"rpd"`
	Rpo                   string     `gorm:"column:rpo" json:"rpo"`
	Alergi                string     `gorm:"column:alergi" json:"alergi"`
	Kondisi               string     `gorm:"column:kondisi" json:"kondisi"`
	Status                string     `gorm:"column:status" json:"status"`
	Td                    string     `gorm:"column:td" json:"td"`
	Nadi                  string     `gorm:"column:nadi" json:"nadi"`
	Suhu                  string     `gorm:"column:suhu" json:"suhu"`
	Rr                    string     `gorm:"column:rr" json:"rr"`
	Bb                    string     `gorm:"column:bb" json:"bb"`
	Nyeri                 string     `gorm:"column:nyeri" json:"nyeri"`
	Gcs                   string     `gorm:"column:gcs" json:"gcs"`
	Kepala                string     `gorm:"column:kepala" json:"kepala"`
	KeteranganKepala      string     `gorm:"column:keterangan_kepala" json:"keterangan_kepala"`
	Thoraks               string     `gorm:"column:thoraks" json:"thoraks"`
	KeteranganThorak      string     `gorm:"column:keterangan_thorak" json:"keterangan_thorak"`
	Abdomen               string     `gorm:"column:abdomen" json:"abdomen"`
	KeteranganAbdomen     string     `gorm:"column:keterangan_abdomen" json:"keterangan_abdomen"`
	Ekstremitas           string     `gorm:"column:ekstremitas" json:"ekstremitas"`
	KeteranganEkstremitas string     `gorm:"column:keterangan_ekstremitas" json:"keterangan_ekstremitas"`
	Lainnya               string     `gorm:"column:lainnya" json:"lainnya"`
	Lab                   string     `gorm:"column:lab" json:"lab"`
	Rad                   string     `gorm:"column:rad" json:"rad"`
	Penunjanglain         string     `gorm:"column:penunjanglain" json:"penunjanglain"`
	Diagnosis             string     `gorm:"column:diagnosis" json:"diagnosis"`
	Diagnosis2            string     `gorm:"column:diagnosis2" json:"diagnosis2"`
	Permasalahan          string     `gorm:"column:permasalahan" json:"permasalahan"`
	Terapi                string     `gorm:"column:terapi" json:"terapi"`
	Tindakan              string     `gorm:"column:tindakan" json:"tindakan"`
	Edukasi               string     `gorm:"column:edukasi" json:"edukasi"`
}

func (PenilaianMedisRalanPenyakitDalam) TableName() string {
	return "penilaian_medis_ralan_penyakit_dalam"
}
