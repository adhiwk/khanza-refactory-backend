package rekammedis

import "time"

// ChecklistKriteriaMasukIsolasi tabel `checklist_kriteria_masuk_isolasi` (checklist kriteria masuk isolasi, RMChecklistKriteriaMasukIsolasi).
type ChecklistKriteriaMasukIsolasi struct {
	NoRawat                  string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                  *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	AirborneTb               string     `gorm:"column:airborne_tb" json:"airborne_tb"`
	AirborneCampak           string     `gorm:"column:airborne_campak" json:"airborne_campak"`
	AirborneVarisela         string     `gorm:"column:airborne_varisela" json:"airborne_varisela"`
	AirborneZosterDiseminata string     `gorm:"column:airborne_zoster_diseminata" json:"airborne_zoster_diseminata"`
	AirborneLainnya          string     `gorm:"column:airborne_lainnya" json:"airborne_lainnya"`
	DropletCovid19           string     `gorm:"column:droplet_covid19" json:"droplet_covid19"`
	DropletInfluenza         string     `gorm:"column:droplet_influenza" json:"droplet_influenza"`
	DropletDifteri           string     `gorm:"column:droplet_difteri" json:"droplet_difteri"`
	DropletPertusis          string     `gorm:"column:droplet_pertusis" json:"droplet_pertusis"`
	DropletMeningitis        string     `gorm:"column:droplet_meningitis" json:"droplet_meningitis"`
	DropletLainnya           string     `gorm:"column:droplet_lainnya" json:"droplet_lainnya"`
	KontakMdro               string     `gorm:"column:kontak_mdro" json:"kontak_mdro"`
	KontakClostridium        string     `gorm:"column:kontak_clostridium" json:"kontak_clostridium"`
	KontakScabies            string     `gorm:"column:kontak_scabies" json:"kontak_scabies"`
	KontakLukaDrainase       string     `gorm:"column:kontak_luka_drainase" json:"kontak_luka_drainase"`
	KontakDiareInfeksi       string     `gorm:"column:kontak_diare_infeksi" json:"kontak_diare_infeksi"`
	KontakLainnya            string     `gorm:"column:kontak_lainnya" json:"kontak_lainnya"`
	KontakErat               string     `gorm:"column:kontak_erat" json:"kontak_erat"`
	RiwayatPerjalananWabah   string     `gorm:"column:riwayat_perjalanan_wabah" json:"riwayat_perjalanan_wabah"`
	RiwayatMdro              string     `gorm:"column:riwayat_mdro" json:"riwayat_mdro"`
	HasilLabRadiologiPositif string     `gorm:"column:hasil_lab_radiologi_positif" json:"hasil_lab_radiologi_positif"`
	GejalaKlinisMenular      string     `gorm:"column:gejala_klinis_menular" json:"gejala_klinis_menular"`
	PasienImunokompromis     string     `gorm:"column:pasien_imunokompromis" json:"pasien_imunokompromis"`
	InstruksiDpjp            string     `gorm:"column:instruksi_dpjp" json:"instruksi_dpjp"`
	PersetujuanIsolasi       string     `gorm:"column:persetujuan_isolasi" json:"persetujuan_isolasi"`
	KelengkapanJaminan       string     `gorm:"column:kelengkapan_jaminan" json:"kelengkapan_jaminan"`
	KetersediaanApd          string     `gorm:"column:ketersediaan_apd" json:"ketersediaan_apd"`
	FasilitasCuciTangan      string     `gorm:"column:fasilitas_cuci_tangan" json:"fasilitas_cuci_tangan"`
	TekananNegatifBerfungsi  string     `gorm:"column:tekanan_negatif_berfungsi" json:"tekanan_negatif_berfungsi"`
	PintuOtomatisBerfungsi   string     `gorm:"column:pintu_otomatis_berfungsi" json:"pintu_otomatis_berfungsi"`
	IndikasiIsolasi          string     `gorm:"column:indikasi_isolasi" json:"indikasi_isolasi"`
	JenisIsolasi             string     `gorm:"column:jenis_isolasi" json:"jenis_isolasi"`
	DiagnosaIsolasi          string     `gorm:"column:diagnosa_isolasi" json:"diagnosa_isolasi"`
	Keterangan               *string    `gorm:"column:keterangan" json:"keterangan"`
	Nik                      *string    `gorm:"column:nik" json:"nik"`
}

func (ChecklistKriteriaMasukIsolasi) TableName() string {
	return "checklist_kriteria_masuk_isolasi"
}
