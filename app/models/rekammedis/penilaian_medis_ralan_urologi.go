package rekammedis

import "time"

// PenilaianMedisRalanUrologi tabel `penilaian_medis_ralan_urologi` (penilaian awal medis ralan urologi, RMPenilaianAwalMedisRalanUrologi).
type PenilaianMedisRalanUrologi struct {
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
	RiwayatKebiasaan      string     `gorm:"column:riwayat_kebiasaan" json:"riwayat_kebiasaan"`
	RiwayatOperasiUrologi string     `gorm:"column:riwayat_operasi_urologi" json:"riwayat_operasi_urologi"`
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
	Thoraks               string     `gorm:"column:thoraks" json:"thoraks"`
	KeteranganThoraks     *string    `gorm:"column:keterangan_thoraks" json:"keterangan_thoraks"`
	Abdomen               string     `gorm:"column:abdomen" json:"abdomen"`
	KeteranganAbdomen     *string    `gorm:"column:keterangan_abdomen" json:"keterangan_abdomen"`
	Ekstrimitas           string     `gorm:"column:ekstrimitas" json:"ekstrimitas"`
	KeteranganEkstrimitas *string    `gorm:"column:keterangan_ekstrimitas" json:"keterangan_ekstrimitas"`
	NyeriKetokCva         *string    `gorm:"column:nyeri_ketok_cva" json:"nyeri_ketok_cva"`
	GenitaliaEksternal    *string    `gorm:"column:genitalia_eksternal" json:"genitalia_eksternal"`
	ColokDubur            *string    `gorm:"column:colok_dubur" json:"colok_dubur"`
	Lainnya               string     `gorm:"column:lainnya" json:"lainnya"`
	Urinalisis            string     `gorm:"column:urinalisis" json:"urinalisis"`
	Darah                 string     `gorm:"column:darah" json:"darah"`
	UsgUrologi            string     `gorm:"column:usg_urologi" json:"usg_urologi"`
	Radiologi             string     `gorm:"column:radiologi" json:"radiologi"`
	PenunjangLain         string     `gorm:"column:penunjang_lain" json:"penunjang_lain"`
	Diagnosis             string     `gorm:"column:diagnosis" json:"diagnosis"`
	Diagnosis2            string     `gorm:"column:diagnosis2" json:"diagnosis2"`
	Permasalahan          string     `gorm:"column:permasalahan" json:"permasalahan"`
	Terapi                string     `gorm:"column:terapi" json:"terapi"`
	Tindakan              string     `gorm:"column:tindakan" json:"tindakan"`
	Edukasi               string     `gorm:"column:edukasi" json:"edukasi"`
}

func (PenilaianMedisRalanUrologi) TableName() string {
	return "penilaian_medis_ralan_urologi"
}
