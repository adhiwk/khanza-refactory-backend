package rekammedis

import "time"

// PenilaianMedisRalanGeriatri tabel `penilaian_medis_ralan_geriatri` (penilaian awal medis ralan geriatri, RMPenilaianAwalMedisRalanGeriatri).
type PenilaianMedisRalanGeriatri struct {
	NoRawat                 string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                 *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter                string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Anamnesis               string     `gorm:"column:anamnesis" json:"anamnesis"`
	Hubungan                string     `gorm:"column:hubungan" json:"hubungan"`
	KeluhanUtama            string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rps                     string     `gorm:"column:rps" json:"rps"`
	Rpd                     string     `gorm:"column:rpd" json:"rpd"`
	Rpo                     string     `gorm:"column:rpo" json:"rpo"`
	Alergi                  string     `gorm:"column:alergi" json:"alergi"`
	TulangBelakang          string     `gorm:"column:tulang_belakang" json:"tulang_belakang"`
	Td                      string     `gorm:"column:td" json:"td"`
	Nadi                    string     `gorm:"column:nadi" json:"nadi"`
	Suhu                    string     `gorm:"column:suhu" json:"suhu"`
	Rr                      string     `gorm:"column:rr" json:"rr"`
	KondisiUmum             string     `gorm:"column:kondisi_umum" json:"kondisi_umum"`
	StatusPsikologisGds     string     `gorm:"column:status_psikologis_gds" json:"status_psikologis_gds"`
	KondisiSosial           string     `gorm:"column:kondisi_sosial" json:"kondisi_sosial"`
	StatusKognitifMmse      string     `gorm:"column:status_kognitif_mmse" json:"status_kognitif_mmse"`
	Kepala                  string     `gorm:"column:kepala" json:"kepala"`
	KeteranganKepala        string     `gorm:"column:keterangan_kepala" json:"keterangan_kepala"`
	Thoraks                 string     `gorm:"column:thoraks" json:"thoraks"`
	KeteranganThoraks       string     `gorm:"column:keterangan_thoraks" json:"keterangan_thoraks"`
	Abdomen                 string     `gorm:"column:abdomen" json:"abdomen"`
	KeteranganAbdomen       string     `gorm:"column:keterangan_abdomen" json:"keterangan_abdomen"`
	Ekstremitas             string     `gorm:"column:ekstremitas" json:"ekstremitas"`
	KeteranganEkstremitas   string     `gorm:"column:keterangan_ekstremitas" json:"keterangan_ekstremitas"`
	IntegumentKebersihan    string     `gorm:"column:Integument_kebersihan" json:"Integument_kebersihan"`
	IntegumentWarna         string     `gorm:"column:Integument_warna" json:"Integument_warna"`
	IntegumentKelembaban    string     `gorm:"column:Integument_kelembaban" json:"Integument_kelembaban"`
	IntegumentGangguanKulit string     `gorm:"column:Integument_gangguan_kulit" json:"Integument_gangguan_kulit"`
	StatusFungsional        string     `gorm:"column:status_fungsional" json:"status_fungsional"`
	SkriningJatuh           string     `gorm:"column:skrining_jatuh" json:"skrining_jatuh"`
	StatusNutrisi           string     `gorm:"column:status_nutrisi" json:"status_nutrisi"`
	Lainnya                 string     `gorm:"column:lainnya" json:"lainnya"`
	Lab                     string     `gorm:"column:lab" json:"lab"`
	Rad                     string     `gorm:"column:rad" json:"rad"`
	Pemeriksaan             string     `gorm:"column:pemeriksaan" json:"pemeriksaan"`
	Diagnosis               string     `gorm:"column:diagnosis" json:"diagnosis"`
	Diagnosis2              string     `gorm:"column:diagnosis2" json:"diagnosis2"`
	Permasalahan            string     `gorm:"column:permasalahan" json:"permasalahan"`
	Terapi                  string     `gorm:"column:terapi" json:"terapi"`
	Tindakan                string     `gorm:"column:tindakan" json:"tindakan"`
	Edukasi                 string     `gorm:"column:edukasi" json:"edukasi"`
}

func (PenilaianMedisRalanGeriatri) TableName() string {
	return "penilaian_medis_ralan_geriatri"
}
