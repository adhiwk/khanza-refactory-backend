package rekammedis

import "time"

// PenilaianMedisRalanRehabMedik tabel `penilaian_medis_ralan_rehab_medik` (penilaian awal medis ralan rehab medik, RMPenilaianAwalMedisRalanRehabMedik).
type PenilaianMedisRalanRehabMedik struct {
	NoRawat               string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal               *time.Time `gorm:"column:tanggal" json:"tanggal"`
	KdDokter              *string    `gorm:"column:kd_dokter" json:"kd_dokter"`
	Anamnesis             *string    `gorm:"column:anamnesis" json:"anamnesis"`
	Hubungan              *string    `gorm:"column:hubungan" json:"hubungan"`
	KeluhanUtama          *string    `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rps                   *string    `gorm:"column:rps" json:"rps"`
	Rpd                   *string    `gorm:"column:rpd" json:"rpd"`
	Alergi                *string    `gorm:"column:alergi" json:"alergi"`
	Kesadaran             *string    `gorm:"column:kesadaran" json:"kesadaran"`
	Nyeri                 *string    `gorm:"column:nyeri" json:"nyeri"`
	SkalaNyeri            *string    `gorm:"column:skala_nyeri" json:"skala_nyeri"`
	Td                    *string    `gorm:"column:td" json:"td"`
	Nadi                  *string    `gorm:"column:nadi" json:"nadi"`
	Suhu                  *string    `gorm:"column:suhu" json:"suhu"`
	Rr                    *string    `gorm:"column:rr" json:"rr"`
	Bb                    *string    `gorm:"column:bb" json:"bb"`
	Kepala                *string    `gorm:"column:kepala" json:"kepala"`
	KeteranganKepala      *string    `gorm:"column:keterangan_kepala" json:"keterangan_kepala"`
	Thoraks               *string    `gorm:"column:thoraks" json:"thoraks"`
	KeteranganThoraks     *string    `gorm:"column:keterangan_thoraks" json:"keterangan_thoraks"`
	Abdomen               *string    `gorm:"column:abdomen" json:"abdomen"`
	KeteranganAbdomen     *string    `gorm:"column:keterangan_abdomen" json:"keterangan_abdomen"`
	Ekstremitas           *string    `gorm:"column:ekstremitas" json:"ekstremitas"`
	KeteranganEkstremitas *string    `gorm:"column:keterangan_ekstremitas" json:"keterangan_ekstremitas"`
	Columna               *string    `gorm:"column:columna" json:"columna"`
	KeteranganColumna     *string    `gorm:"column:keterangan_columna" json:"keterangan_columna"`
	Muskulos              *string    `gorm:"column:muskulos" json:"muskulos"`
	KeteranganMuskulos    *string    `gorm:"column:keterangan_muskulos" json:"keterangan_muskulos"`
	Lainnya               *string    `gorm:"column:lainnya" json:"lainnya"`
	ResikoJatuh           *string    `gorm:"column:resiko_jatuh" json:"resiko_jatuh"`
	ResikoNutrisional     *string    `gorm:"column:resiko_nutrisional" json:"resiko_nutrisional"`
	KebutuhanFungsional   *string    `gorm:"column:kebutuhan_fungsional" json:"kebutuhan_fungsional"`
	DiagnosaMedis         *string    `gorm:"column:diagnosa_medis" json:"diagnosa_medis"`
	DiagnosaFungsi        *string    `gorm:"column:diagnosa_fungsi" json:"diagnosa_fungsi"`
	PenunjangLain         *string    `gorm:"column:penunjang_lain" json:"penunjang_lain"`
	Fisio                 *string    `gorm:"column:fisio" json:"fisio"`
	Okupasi               *string    `gorm:"column:okupasi" json:"okupasi"`
	Wicara                *string    `gorm:"column:wicara" json:"wicara"`
	Akupuntur             *string    `gorm:"column:akupuntur" json:"akupuntur"`
	Tatalain              *string    `gorm:"column:tatalain" json:"tatalain"`
	FrekuensiTerapi       *string    `gorm:"column:frekuensi_terapi" json:"frekuensi_terapi"`
	Fisioterapi           *time.Time `gorm:"column:fisioterapi" json:"fisioterapi"`
	TerapiOkupasi         *time.Time `gorm:"column:terapi_okupasi" json:"terapi_okupasi"`
	TerapiWicara          *time.Time `gorm:"column:terapi_wicara" json:"terapi_wicara"`
	TerapiAkupuntur       *time.Time `gorm:"column:terapi_akupuntur" json:"terapi_akupuntur"`
	TerapiLainnya         *time.Time `gorm:"column:terapi_lainnya" json:"terapi_lainnya"`
	Edukasi               *string    `gorm:"column:edukasi" json:"edukasi"`
}

func (PenilaianMedisRalanRehabMedik) TableName() string {
	return "penilaian_medis_ralan_rehab_medik"
}
