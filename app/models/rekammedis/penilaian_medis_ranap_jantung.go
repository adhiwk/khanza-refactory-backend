package rekammedis

import "time"

// PenilaianMedisRanapJantung tabel `penilaian_medis_ranap_jantung` (penilaian awal medis ranap jantung, RMPenilaianAwalMedisRanapJantung).
type PenilaianMedisRanapJantung struct {
	NoRawat               string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal               *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter              string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Anamnesis             string     `gorm:"column:anamnesis" json:"anamnesis"`
	Hubungan              string     `gorm:"column:hubungan" json:"hubungan"`
	KeluhanUtama          string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rps                   string     `gorm:"column:rps" json:"rps"`
	Rpk                   string     `gorm:"column:rpk" json:"rpk"`
	Rpd                   string     `gorm:"column:rpd" json:"rpd"`
	Rpo                   string     `gorm:"column:rpo" json:"rpo"`
	Alergi                string     `gorm:"column:alergi" json:"alergi"`
	Td                    string     `gorm:"column:td" json:"td"`
	Bb                    string     `gorm:"column:bb" json:"bb"`
	Tb                    string     `gorm:"column:tb" json:"tb"`
	Suhu                  string     `gorm:"column:suhu" json:"suhu"`
	Nadi                  string     `gorm:"column:nadi" json:"nadi"`
	Rr                    string     `gorm:"column:rr" json:"rr"`
	KeadaanUmum           *string    `gorm:"column:keadaan_umum" json:"keadaan_umum"`
	Nyeri                 string     `gorm:"column:nyeri" json:"nyeri"`
	StatusNutrisi         string     `gorm:"column:status_nutrisi" json:"status_nutrisi"`
	Jantung               string     `gorm:"column:jantung" json:"jantung"`
	KeteranganJantung     *string    `gorm:"column:keterangan_jantung" json:"keterangan_jantung"`
	Paru                  string     `gorm:"column:paru" json:"paru"`
	KeteranganParu        *string    `gorm:"column:keterangan_paru" json:"keterangan_paru"`
	Ekstrimitas           string     `gorm:"column:ekstrimitas" json:"ekstrimitas"`
	KeteranganEkstrimitas *string    `gorm:"column:keterangan_ekstrimitas" json:"keterangan_ekstrimitas"`
	Lainnya               string     `gorm:"column:lainnya" json:"lainnya"`
	Lab                   string     `gorm:"column:lab" json:"lab"`
	Ekg                   string     `gorm:"column:ekg" json:"ekg"`
	PenunjangLain         string     `gorm:"column:penunjang_lain" json:"penunjang_lain"`
	Diagnosis             string     `gorm:"column:diagnosis" json:"diagnosis"`
	Diagnosis2            string     `gorm:"column:diagnosis2" json:"diagnosis2"`
	Permasalahan          string     `gorm:"column:permasalahan" json:"permasalahan"`
	Terapi                string     `gorm:"column:terapi" json:"terapi"`
	Tindakan              string     `gorm:"column:tindakan" json:"tindakan"`
	Edukasi               string     `gorm:"column:edukasi" json:"edukasi"`
}

func (PenilaianMedisRanapJantung) TableName() string {
	return "penilaian_medis_ranap_jantung"
}
