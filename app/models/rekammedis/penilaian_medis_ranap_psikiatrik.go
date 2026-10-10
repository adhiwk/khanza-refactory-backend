package rekammedis

import "time"

// PenilaianMedisRanapPsikiatrik tabel `penilaian_medis_ranap_psikiatrik` (penilaian awal medis ranap psikiatrik, RMPenilaianAwalMedisRanapPsikiatrik).
type PenilaianMedisRanapPsikiatrik struct {
	NoRawat              string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal              *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter             string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Anamnesis            string     `gorm:"column:anamnesis" json:"anamnesis"`
	Hubungan             string     `gorm:"column:hubungan" json:"hubungan"`
	KeluhanUtama         string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rps                  string     `gorm:"column:rps" json:"rps"`
	Rpd                  string     `gorm:"column:rpd" json:"rpd"`
	Rpk                  string     `gorm:"column:rpk" json:"rpk"`
	Rpo                  string     `gorm:"column:rpo" json:"rpo"`
	Alergi               string     `gorm:"column:alergi" json:"alergi"`
	Penampilan           string     `gorm:"column:penampilan" json:"penampilan"`
	Pembicaraan          string     `gorm:"column:pembicaraan" json:"pembicaraan"`
	Psikomotor           string     `gorm:"column:psikomotor" json:"psikomotor"`
	Sikap                string     `gorm:"column:sikap" json:"sikap"`
	Mood                 string     `gorm:"column:mood" json:"mood"`
	FungsiKognitif       string     `gorm:"column:fungsi_kognitif" json:"fungsi_kognitif"`
	GangguanPersepsi     string     `gorm:"column:gangguan_persepsi" json:"gangguan_persepsi"`
	ProsesPikir          string     `gorm:"column:proses_pikir" json:"proses_pikir"`
	PengendalianImpuls   string     `gorm:"column:pengendalian_impuls" json:"pengendalian_impuls"`
	Tilikan              string     `gorm:"column:tilikan" json:"tilikan"`
	Rta                  string     `gorm:"column:rta" json:"rta"`
	SkalaPenilaianKhusus string     `gorm:"column:skala_penilaian_khusus" json:"skala_penilaian_khusus"`
	Keadaan              string     `gorm:"column:keadaan" json:"keadaan"`
	Gcs                  string     `gorm:"column:gcs" json:"gcs"`
	Kesadaran            string     `gorm:"column:kesadaran" json:"kesadaran"`
	Td                   string     `gorm:"column:td" json:"td"`
	Nadi                 string     `gorm:"column:nadi" json:"nadi"`
	Rr                   string     `gorm:"column:rr" json:"rr"`
	Suhu                 string     `gorm:"column:suhu" json:"suhu"`
	Spo                  string     `gorm:"column:spo" json:"spo"`
	Bb                   string     `gorm:"column:bb" json:"bb"`
	Tb                   string     `gorm:"column:tb" json:"tb"`
	Kepala               string     `gorm:"column:kepala" json:"kepala"`
	Gigi                 string     `gorm:"column:gigi" json:"gigi"`
	Tht                  string     `gorm:"column:tht" json:"tht"`
	Thoraks              string     `gorm:"column:thoraks" json:"thoraks"`
	Abdomen              string     `gorm:"column:abdomen" json:"abdomen"`
	Genital              string     `gorm:"column:genital" json:"genital"`
	Ekstremitas          string     `gorm:"column:ekstremitas" json:"ekstremitas"`
	Kulit                string     `gorm:"column:kulit" json:"kulit"`
	KetFisik             string     `gorm:"column:ket_fisik" json:"ket_fisik"`
	Penunjang            string     `gorm:"column:penunjang" json:"penunjang"`
	Diagnosis            string     `gorm:"column:diagnosis" json:"diagnosis"`
	Tata                 string     `gorm:"column:tata" json:"tata"`
	Konsulrujuk          string     `gorm:"column:konsulrujuk" json:"konsulrujuk"`
}

func (PenilaianMedisRanapPsikiatrik) TableName() string {
	return "penilaian_medis_ranap_psikiatrik"
}
