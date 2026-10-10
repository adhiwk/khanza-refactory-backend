package rekammedis

import "time"

// PenilaianMedisRalanBedahMulut tabel `penilaian_medis_ralan_bedah_mulut` (penilaian awal medis ralan bedah mulut, RMPenilaianAwalMedisRalanBedahMulut).
type PenilaianMedisRalanBedahMulut struct {
	NoRawat               string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal               *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter              string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Anamnesis             string     `gorm:"column:anamnesis" json:"anamnesis"`
	Hubungan              string     `gorm:"column:hubungan" json:"hubungan"`
	KeluhanUtama          string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rps                   string     `gorm:"column:rps" json:"rps"`
	Rpk                   string     `gorm:"column:rpk" json:"rpk"`
	Alergi                string     `gorm:"column:alergi" json:"alergi"`
	Keadaan               string     `gorm:"column:keadaan" json:"keadaan"`
	Kesadaran             string     `gorm:"column:kesadaran" json:"kesadaran"`
	Nyeri                 string     `gorm:"column:nyeri" json:"nyeri"`
	Td                    string     `gorm:"column:td" json:"td"`
	Nadi                  string     `gorm:"column:nadi" json:"nadi"`
	Suhu                  string     `gorm:"column:suhu" json:"suhu"`
	Rr                    string     `gorm:"column:rr" json:"rr"`
	Bb                    string     `gorm:"column:bb" json:"bb"`
	Tb                    string     `gorm:"column:tb" json:"tb"`
	StatusNutrisi         string     `gorm:"column:status_nutrisi" json:"status_nutrisi"`
	Kulit                 string     `gorm:"column:kulit" json:"kulit"`
	KeteranganKulit       string     `gorm:"column:keterangan_kulit" json:"keterangan_kulit"`
	Kepala                string     `gorm:"column:kepala" json:"kepala"`
	KeteranganKepala      string     `gorm:"column:keterangan_kepala" json:"keterangan_kepala"`
	Mata                  string     `gorm:"column:mata" json:"mata"`
	KeteranganMata        string     `gorm:"column:keterangan_mata" json:"keterangan_mata"`
	Leher                 string     `gorm:"column:leher" json:"leher"`
	KeteranganLeher       string     `gorm:"column:keterangan_leher" json:"keterangan_leher"`
	Kelenjar              string     `gorm:"column:kelenjar" json:"kelenjar"`
	KeteranganKelenjar    string     `gorm:"column:keterangan_kelenjar" json:"keterangan_kelenjar"`
	Dada                  string     `gorm:"column:dada" json:"dada"`
	KeteranganDada        string     `gorm:"column:keterangan_dada" json:"keterangan_dada"`
	Perut                 string     `gorm:"column:perut" json:"perut"`
	KeteranganPerut       string     `gorm:"column:keterangan_perut" json:"keterangan_perut"`
	Ekstremitas           string     `gorm:"column:ekstremitas" json:"ekstremitas"`
	KeteranganEkstremitas string     `gorm:"column:keterangan_ekstremitas" json:"keterangan_ekstremitas"`
	Wajah                 string     `gorm:"column:wajah" json:"wajah"`
	Intra                 string     `gorm:"column:intra" json:"intra"`
	Gigigeligi            string     `gorm:"column:gigigeligi" json:"gigigeligi"`
	Lab                   string     `gorm:"column:lab" json:"lab"`
	Rad                   string     `gorm:"column:rad" json:"rad"`
	Penunjang             string     `gorm:"column:penunjang" json:"penunjang"`
	Diagnosis             string     `gorm:"column:diagnosis" json:"diagnosis"`
	Diagnosis2            string     `gorm:"column:diagnosis2" json:"diagnosis2"`
	Permasalahan          string     `gorm:"column:permasalahan" json:"permasalahan"`
	Terapi                string     `gorm:"column:terapi" json:"terapi"`
	Tindakan              string     `gorm:"column:tindakan" json:"tindakan"`
	Edukasi               string     `gorm:"column:edukasi" json:"edukasi"`
}

func (PenilaianMedisRalanBedahMulut) TableName() string {
	return "penilaian_medis_ralan_bedah_mulut"
}
