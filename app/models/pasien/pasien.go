package pasien

import (
	"time"
)

type Pasien struct {
	NoRkmMedis       string     `gorm:"column:no_rkm_medis;primaryKey;type:varchar(15);not null" json:"no_rkm_medis"`
	NmPasien         *string    `gorm:"column:nm_pasien;type:varchar(40)" json:"nm_pasien"`
	NoKtp            *string    `gorm:"column:no_ktp;type:varchar(20)" json:"no_ktp"`
	Jk               *string    `gorm:"column:jk;type:enum('L','P')" json:"jk"`
	TmpLahir         *string    `gorm:"column:tmp_lahir;type:varchar(15)" json:"tmp_lahir"`
	TglLahir         *time.Time `gorm:"column:tgl_lahir;type:date" json:"tgl_lahir"`
	NmIbu            string     `gorm:"column:nm_ibu;type:varchar(40);not null" json:"nm_ibu"`
	Alamat           *string    `gorm:"column:alamat;type:varchar(200)" json:"alamat"`
	GolDarah         *string    `gorm:"column:gol_darah;type:enum('A','B','O','AB','-')" json:"gol_darah"`
	Pekerjaan        *string    `gorm:"column:pekerjaan;type:varchar(60)" json:"pekerjaan"`
	SttsNikah        *string    `gorm:"column:stts_nikah;type:enum('BELUM MENIKAH','MENIKAH','JANDA','DUDHA','JOMBLO')" json:"stts_nikah"`
	Agama            *string    `gorm:"column:agama;type:varchar(12)" json:"agama"`
	TglDaftar        *time.Time `gorm:"column:tgl_daftar;type:date" json:"tgl_daftar"`
	NoTlp            *string    `gorm:"column:no_tlp;type:varchar(40)" json:"no_tlp"`
	Umur             string     `gorm:"column:umur;type:varchar(30);not null" json:"umur"`
	Pnd              string     `gorm:"column:pnd;type:enum('TS','TK','SD','SMP','SMA','SLTA/SEDERAJAT','D1','D2','D3','D4','S1','S2','S3','-');not null" json:"pnd"`
	Keluarga         *string    `gorm:"column:keluarga;type:enum('AYAH','IBU','ISTRI','SUAMI','SAUDARA','ANAK','DIRI SENDIRI','LAIN-LAIN')" json:"keluarga"`
	NamaKeluarga     string     `gorm:"column:namakeluarga;type:varchar(50);not null" json:"namakeluarga"`
	KdPj             string     `gorm:"column:kd_pj;type:char(3);not null" json:"kd_pj"`
	NoPeserta        *string    `gorm:"column:no_peserta;type:varchar(25)" json:"no_peserta"`
	KdKel            int        `gorm:"column:kd_kel;type:int(11);not null" json:"kd_kel"`
	KdKec            int        `gorm:"column:kd_kec;type:int(11);not null" json:"kd_kec"`
	KdKab            int        `gorm:"column:kd_kab;type:int(11);not null" json:"kd_kab"`
	PekerjaanPj      string     `gorm:"column:pekerjaanpj;type:varchar(35);not null" json:"pekerjaanpj"`
	AlamatPj         string     `gorm:"column:alamatpj;type:varchar(100);not null" json:"alamatpj"`
	KelurahanPj      string     `gorm:"column:kelurahanpj;type:varchar(60);not null" json:"kelurahanpj"`
	KecamatanPj      string     `gorm:"column:kecamatanpj;type:varchar(60);not null" json:"kecamatanpj"`
	KabupatenPj      string     `gorm:"column:kabupatenpj;type:varchar(60);not null" json:"kabupatenpj"`
	PerusahaanPasien string     `gorm:"column:perusahaan_pasien;type:varchar(8);not null" json:"perusahaan_pasien"`
	SukuBangsa       int        `gorm:"column:suku_bangsa;type:int(11);not null" json:"suku_bangsa"`
	BahasaPasien     int        `gorm:"column:bahasa_pasien;type:int(11);not null" json:"bahasa_pasien"`
	CacatFisik       int        `gorm:"column:cacat_fisik;type:int(11);not null" json:"cacat_fisik"`
	Email            string     `gorm:"column:email;type:varchar(50);not null" json:"email"`
	Nip              string     `gorm:"column:nip;type:varchar(30);not null" json:"nip"`
	KdProp           int        `gorm:"column:kd_prop;type:int(11);not null" json:"kd_prop"`
	PropinsiPj       string     `gorm:"column:propinsipj;type:varchar(30);not null" json:"propinsipj"`
}

func (Pasien) TableName() string {
	return "pasien"
}

func (Pasien) Connection() string {
	return "mysql_kedua"
}
