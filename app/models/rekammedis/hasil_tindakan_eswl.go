package rekammedis

import "time"

// HasilTindakanEswl tabel `hasil_tindakan_eswl` (hasil tindakan ESWL, RMHasilTindakanESWL).
type HasilTindakanEswl struct {
	NoRawat             string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Mulai               *time.Time `gorm:"column:mulai;primaryKey;autoIncrement:false" json:"mulai"`
	Selesai             *time.Time `gorm:"column:selesai" json:"selesai"`
	KdDokter            string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Nip                 string     `gorm:"column:nip" json:"nip"`
	Diagnosa            string     `gorm:"column:diagnosa" json:"diagnosa"`
	Tindakan            string     `gorm:"column:tindakan" json:"tindakan"`
	ObatAnalgesik       string     `gorm:"column:obat_analgesik" json:"obat_analgesik"`
	ObatLain            string     `gorm:"column:obat_lain" json:"obat_lain"`
	UraianTindakan      string     `gorm:"column:uraian_tindakan" json:"uraian_tindakan"`
	UraianTindakanFocus string     `gorm:"column:uraian_tindakan_focus" json:"uraian_tindakan_focus"`
	UraianTindakanRate  string     `gorm:"column:uraian_tindakan_rate" json:"uraian_tindakan_rate"`
	UraianTindakanPower string     `gorm:"column:uraian_tindakan_power" json:"uraian_tindakan_power"`
	UraianTindakanShock string     `gorm:"column:uraian_tindakan_shock" json:"uraian_tindakan_shock"`
	Diintegrasi         string     `gorm:"column:diintegrasi" json:"diintegrasi"`
	Kekurangan          string     `gorm:"column:kekurangan" json:"kekurangan"`
	Anjungan            string     `gorm:"column:anjungan" json:"anjungan"`
}

func (HasilTindakanEswl) TableName() string {
	return "hasil_tindakan_eswl"
}
