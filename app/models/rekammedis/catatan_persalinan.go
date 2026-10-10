package rekammedis

import "time"

// CatatanPersalinan tabel `catatan_persalinan` (catatan persalinan, RMCatatanPersalinan).
type CatatanPersalinan struct {
	NoRawat               string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Mulai                 *time.Time `gorm:"column:mulai" json:"mulai"`
	Selesai               *time.Time `gorm:"column:selesai" json:"selesai"`
	KdDokter              string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Nip                   string     `gorm:"column:nip" json:"nip"`
	Catatan               *string    `gorm:"column:catatan" json:"catatan"`
	WaktuPersalinanKala1  *string    `gorm:"column:waktu_persalinan_kala_1" json:"waktu_persalinan_kala_1"`
	WaktuPersalinanKala2  *string    `gorm:"column:waktu_persalinan_kala_2" json:"waktu_persalinan_kala_2"`
	WaktuPersalinanKala3  *string    `gorm:"column:waktu_persalinan_kala_3" json:"waktu_persalinan_kala_3"`
	WaktuPersalinanJumlah *string    `gorm:"column:waktu_persalinan_jumlah" json:"waktu_persalinan_jumlah"`
	Perineum              *string    `gorm:"column:perineum" json:"perineum"`
	JahitanLuar1          *string    `gorm:"column:jahitan_luar_1" json:"jahitan_luar_1"`
	JahitanLuar2          *string    `gorm:"column:jahitan_luar_2" json:"jahitan_luar_2"`
	JahitanDalam1         *string    `gorm:"column:jahitan_dalam_1" json:"jahitan_dalam_1"`
	JahitanDalam2         *string    `gorm:"column:jahitan_dalam_2" json:"jahitan_dalam_2"`
	Anak                  *string    `gorm:"column:anak" json:"anak"`
	StatusLahir           *string    `gorm:"column:status_lahir" json:"status_lahir"`
	ApgarScore            *string    `gorm:"column:apgar_score" json:"apgar_score"`
	Bb                    *string    `gorm:"column:bb" json:"bb"`
	Pb                    *string    `gorm:"column:pb" json:"pb"`
	Kelainan              *string    `gorm:"column:kelainan" json:"kelainan"`
	Ketuban               *string    `gorm:"column:ketuban" json:"ketuban"`
	Placenta              *string    `gorm:"column:placenta" json:"placenta"`
	Ukuran                *string    `gorm:"column:ukuran" json:"ukuran"`
	TaliPusat             *string    `gorm:"column:tali_pusat" json:"tali_pusat"`
	Insertio              *string    `gorm:"column:insertio" json:"insertio"`
	DarahKeluarKala1      *string    `gorm:"column:darah_keluar_kala_1" json:"darah_keluar_kala_1"`
	DarahKeluarKala2      *string    `gorm:"column:darah_keluar_kala_2" json:"darah_keluar_kala_2"`
	DarahKeluarKala3      *string    `gorm:"column:darah_keluar_kala_3" json:"darah_keluar_kala_3"`
	DarahKeluarKala4      *string    `gorm:"column:darah_keluar_kala_4" json:"darah_keluar_kala_4"`
	DarahKeluarJumlah     *string    `gorm:"column:darah_keluar_jumlah" json:"darah_keluar_jumlah"`
	KondisiUmum           *string    `gorm:"column:kondisi_umum" json:"kondisi_umum"`
	Td                    *string    `gorm:"column:td" json:"td"`
	Nadi                  *string    `gorm:"column:nadi" json:"nadi"`
	Rr                    *string    `gorm:"column:rr" json:"rr"`
	Suhu                  *string    `gorm:"column:suhu" json:"suhu"`
	KontraksiUterus       *string    `gorm:"column:kontraksi_uterus" json:"kontraksi_uterus"`
	Ppv                   *string    `gorm:"column:ppv" json:"ppv"`
	Pengobatan            *string    `gorm:"column:pengobatan" json:"pengobatan"`
}

func (CatatanPersalinan) TableName() string {
	return "catatan_persalinan"
}
