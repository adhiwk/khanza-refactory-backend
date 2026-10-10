package rekammedis

import "time"

// PenilaianPasienTerminal tabel `penilaian_pasien_terminal` (penilaian pasien terminal, RMPenilaianPasienTerminal).
type PenilaianPasienTerminal struct {
	NoRawat                      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                      *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Diagnosa                     string     `gorm:"column:diagnosa" json:"diagnosa"`
	Rps                          string     `gorm:"column:rps" json:"rps"`
	Rpd                          string     `gorm:"column:rpd" json:"rpd"`
	KeadaanUmum                  *string    `gorm:"column:keadaan_umum" json:"keadaan_umum"`
	Kesadaran                    *string    `gorm:"column:kesadaran" json:"kesadaran"`
	Td                           string     `gorm:"column:td" json:"td"`
	Nadi                         string     `gorm:"column:nadi" json:"nadi"`
	Suhu                         string     `gorm:"column:suhu" json:"suhu"`
	Rr                           string     `gorm:"column:rr" json:"rr"`
	Spo2                         string     `gorm:"column:spo2" json:"spo2"`
	SkalaNyeri                   string     `gorm:"column:skala_nyeri" json:"skala_nyeri"`
	TahapPasienMenjelangAjal     *string    `gorm:"column:tahap_pasien_menjelang_ajal" json:"tahap_pasien_menjelang_ajal"`
	TandaKlinisMenjelangKematian *string    `gorm:"column:tanda_klinis_menjelang_kematian" json:"tanda_klinis_menjelang_kematian"`
	KebutuhanSpiritualPasien     *string    `gorm:"column:kebutuhan_spiritual_pasien" json:"kebutuhan_spiritual_pasien"`
	Nip                          string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianPasienTerminal) TableName() string {
	return "penilaian_pasien_terminal"
}
