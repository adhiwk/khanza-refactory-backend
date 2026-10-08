package registrasi

import (
	"time"
)

type RegPeriksa struct {
	NoReg         *string    `gorm:"column:no_reg;type:varchar(8)" json:"no_reg"`
	NoRawat       string     `gorm:"column:no_rawat;primaryKey;type:varchar(17);not null" json:"no_rawat"`
	TglRegistrasi *time.Time `gorm:"column:tgl_registrasi;type:date" json:"tgl_registrasi"`
	JamReg        *string    `gorm:"column:jam_reg;type:time" json:"jam_reg"`
	KdDokter      *string    `gorm:"column:kd_dokter;type:varchar(20)" json:"kd_dokter"`
	NoRkmMedis    *string    `gorm:"column:no_rkm_medis;type:varchar(15)" json:"no_rkm_medis"`
	KdPoli        *string    `gorm:"column:kd_poli;type:char(5)" json:"kd_poli"`
	PJawab        *string    `gorm:"column:p_jawab;type:varchar(100)" json:"p_jawab"`
	AlmtPj        *string    `gorm:"column:almt_pj;type:varchar(200)" json:"almt_pj"`
	HubunganPj    *string    `gorm:"column:hubunganpj;type:varchar(20)" json:"hubunganpj"`
	BiayaReg      *float64   `gorm:"column:biaya_reg;type:double" json:"biaya_reg"`
	Stts          *string    `gorm:"column:stts;type:enum('Belum','Sudah','Batal','Berkas Diterima','Dirujuk','Meninggal','Dirawat','Pulang Paksa')" json:"stts"`
	SttsDaftar    string     `gorm:"column:stts_daftar;type:enum('-','Lama','Baru');not null" json:"stts_daftar"`
	StatusLanjut  string     `gorm:"column:status_lanjut;type:enum('Ralan','Ranap');not null" json:"status_lanjut"`
	KdPj          string     `gorm:"column:kd_pj;type:char(3);not null" json:"kd_pj"`
	UmurDaftar    *int       `gorm:"column:umurdaftar;type:int(11)" json:"umurdaftar"`
	SttsUmur      *string    `gorm:"column:sttsumur;type:enum('Th','Bl','Hr')" json:"sttsumur"`
	StatusBayar   string     `gorm:"column:status_bayar;type:enum('Sudah Bayar','Belum Bayar');not null" json:"status_bayar"`
	StatusPoli    string     `gorm:"column:status_poli;type:enum('Lama','Baru');not null" json:"status_poli"`
}

func (RegPeriksa) TableName() string {
	return "reg_periksa"
}

func (RegPeriksa) Connection() string {
	return "mysql_kedua"
}

// RegPeriksaView baris list registrasi beserta nama dokter / pasien / poli / penjab.
type RegPeriksaView struct {
	RegPeriksa
	NmDokter *string `gorm:"column:nm_dokter" json:"nm_dokter"`
	NmPasien *string `gorm:"column:nm_pasien" json:"nm_pasien"`
	Jk       *string `gorm:"column:jk" json:"jk"`
	NoTlp    *string `gorm:"column:no_tlp" json:"no_tlp"`
	NmPoli   *string `gorm:"column:nm_poli" json:"nm_poli"`
	PngJawab *string `gorm:"column:png_jawab" json:"png_jawab"`
}
