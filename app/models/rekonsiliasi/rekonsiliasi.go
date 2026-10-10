// Package rekonsiliasi model rekonsiliasi obat (rekonsiliasi_obat, _detail_obat, _konfirmasi).
package rekonsiliasi

// Rekonsiliasi header beserta daftar obat & konfirmasi farmasi.
type Rekonsiliasi struct {
	NoRekonsiliasi       string      `gorm:"column:no_rekonsiliasi" json:"no_rekonsiliasi"`
	NoRawat              string      `gorm:"column:no_rawat" json:"no_rawat"`
	NoRkmMedis           string      `gorm:"column:no_rkm_medis" json:"no_rkm_medis"`
	NmPasien             string      `gorm:"column:nm_pasien" json:"nm_pasien"`
	TanggalWawancara     string      `gorm:"column:tanggal_wawancara" json:"tanggal_wawancara"`
	RekonsiliasiObatSaat string      `gorm:"column:rekonsiliasi_obat_saat" json:"rekonsiliasi_obat_saat"`
	AlergiObat           string      `gorm:"column:alergi_obat" json:"alergi_obat"`
	ManifestasiAlergi    string      `gorm:"column:manifestasi_alergi" json:"manifestasi_alergi"`
	DampakAlergi         string      `gorm:"column:dampak_alergi" json:"dampak_alergi"`
	Nip                  string      `gorm:"column:nip" json:"nip"`
	NmPetugas            string      `gorm:"column:nm_petugas" json:"nm_petugas"`
	Dikonfirmasi         bool        `gorm:"column:dikonfirmasi" json:"dikonfirmasi"`
	Obat                 []Obat      `gorm:"-" json:"obat,omitempty"`
	Konfirmasi           *Konfirmasi `gorm:"-" json:"konfirmasi,omitempty"`
}

type Obat struct {
	NamaObat               string `gorm:"column:nama_obat" json:"nama_obat"`
	DosisObat              string `gorm:"column:dosis_obat" json:"dosis_obat"`
	Frekuensi              string `gorm:"column:frekuensi" json:"frekuensi"`
	CaraPemberian          string `gorm:"column:cara_pemberian" json:"cara_pemberian"`
	WaktuPemberianTerakhir string `gorm:"column:waktu_pemberian_terakhir" json:"waktu_pemberian_terakhir"`
	TindakLanjut           string `gorm:"column:tindak_lanjut" json:"tindak_lanjut"`
	PerubahanAturanPakai   string `gorm:"column:perubahan_aturan_pakai" json:"perubahan_aturan_pakai"`
}

// Konfirmasi tahapan farmasi (RMCariRekonsiliasiObat).
type Konfirmasi struct {
	DiterimaFarmasi      string `gorm:"column:diterima_farmasi" json:"diterima_farmasi"`
	DikonfirmasiApoteker string `gorm:"column:dikonfirmasi_apoteker" json:"dikonfirmasi_apoteker"`
	DiserahkanPasien     string `gorm:"column:diserahkan_pasien" json:"diserahkan_pasien"`
	Nip                  string `gorm:"column:nip" json:"nip"`
	NmPetugas            string `gorm:"column:nm_petugas" json:"nm_petugas"`
}
