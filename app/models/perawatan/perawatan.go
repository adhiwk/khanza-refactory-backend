// Package perawatan model pelayanan pasien per no_rawat (tindakan, pemeriksaan) untuk rawat jalan & rawat inap.
package perawatan

import "time"

// Rawat jenis pelayanan; nilainya sama dengan enum status Khanza (Ralan/Ranap).
type Rawat string

const (
	Ralan Rawat = "Ralan"
	Ranap Rawat = "Ranap"
)

// Pelaksana pelaksana tindakan: dokter, paramedis, atau dokter + paramedis.
type Pelaksana string

const (
	Dokter          Pelaksana = "dr"
	Paramedis       Pelaksana = "pr"
	DokterParamedis Pelaksana = "drpr"
)

func (p Pelaksana) PakaiDokter() bool    { return p == Dokter || p == DokterParamedis }
func (p Pelaksana) PakaiParamedis() bool { return p == Paramedis || p == DokterParamedis }

// Konteks data registrasi yang menentukan aturan & tarif tindakan.
type Konteks struct {
	NoRawat       string     `gorm:"column:no_rawat"`
	NoRkmMedis    string     `gorm:"column:no_rkm_medis"`
	NmPasien      string     `gorm:"column:nm_pasien"`
	TglRegistrasi *time.Time `gorm:"column:tgl_registrasi"`
	JamReg        string     `gorm:"column:jam_reg"`
	KdPoli        string     `gorm:"column:kd_poli"`
	KdPj          string     `gorm:"column:kd_pj"`
	Stts          string     `gorm:"column:stts"`
	StatusLanjut  string     `gorm:"column:status_lanjut"`
	Billing       bool       `gorm:"column:billing"`
	KdBangsal     string     `gorm:"column:kd_bangsal"` // kamar inap aktif (Ranap)
	Kelas         string     `gorm:"column:kelas"`      // kelas kamar inap aktif (Ranap)
}

// Terkunci registrasi batal atau billing sudah dibuat (sekuel.cariRegistrasi).
func (k Konteks) Terkunci() bool {
	return k.Billing || k.Stts == "Batal"
}

// Tarif master tarif tindakan (jns_perawatan / jns_perawatan_inap).
type Tarif struct {
	KdJenisPrw      string  `gorm:"column:kd_jenis_prw" json:"kd_jenis_prw"`
	NmPerawatan     string  `gorm:"column:nm_perawatan" json:"nm_perawatan"`
	Material        float64 `gorm:"column:material" json:"material"`
	Bhp             float64 `gorm:"column:bhp" json:"bhp"`
	TarifTindakandr float64 `gorm:"column:tarif_tindakandr" json:"tarif_tindakandr"`
	TarifTindakanpr float64 `gorm:"column:tarif_tindakanpr" json:"tarif_tindakanpr"`
	Kso             float64 `gorm:"column:kso" json:"kso"`
	Menejemen       float64 `gorm:"column:menejemen" json:"menejemen"`
	TotalByrdr      float64 `gorm:"column:total_byrdr" json:"total_byrdr"`
	TotalByrpr      float64 `gorm:"column:total_byrpr" json:"total_byrpr"`
	TotalByrdrpr    float64 `gorm:"column:total_byrdrpr" json:"total_byrdrpr"`
}

// Tindakan satu baris rawat_jl_dr/pr/drpr atau rawat_inap_dr/pr/drpr.
type Tindakan struct {
	Pelaksana       Pelaksana `gorm:"column:pelaksana" json:"pelaksana"`
	NoRawat         string    `gorm:"column:no_rawat" json:"no_rawat"`
	KdJenisPrw      string    `gorm:"column:kd_jenis_prw" json:"kd_jenis_prw"`
	NmPerawatan     string    `gorm:"column:nm_perawatan" json:"nm_perawatan"`
	KdDokter        string    `gorm:"column:kd_dokter" json:"kd_dokter"`
	NmDokter        string    `gorm:"column:nm_dokter" json:"nm_dokter"`
	Nip             string    `gorm:"column:nip" json:"nip"`
	NmPetugas       string    `gorm:"column:nm_petugas" json:"nm_petugas"`
	TglPerawatan    string    `gorm:"column:tgl_perawatan" json:"tgl_perawatan"`
	JamRawat        string    `gorm:"column:jam_rawat" json:"jam_rawat"`
	Material        float64   `gorm:"column:material" json:"material"`
	Bhp             float64   `gorm:"column:bhp" json:"bhp"`
	TarifTindakandr float64   `gorm:"column:tarif_tindakandr" json:"tarif_tindakandr"`
	TarifTindakanpr float64   `gorm:"column:tarif_tindakanpr" json:"tarif_tindakanpr"`
	Kso             float64   `gorm:"column:kso" json:"kso"`
	Menejemen       float64   `gorm:"column:menejemen" json:"menejemen"`
	BiayaRawat      float64   `gorm:"column:biaya_rawat" json:"biaya_rawat"`
}

// TindakanKey identitas satu baris tindakan (primary key tabel).
type TindakanKey struct {
	Pelaksana    Pelaksana
	NoRawat      string
	KdJenisPrw   string
	KdDokter     string
	Nip          string
	TglPerawatan string
	JamRawat     string
}

// SebelumRegistrasi true bila waktu input lebih awal dari waktu registrasi (sekuel.cekTanggalRegistrasi).
func (k Konteks) SebelumRegistrasi(waktu time.Time) bool {
	if k.TglRegistrasi == nil {
		return false
	}
	reg, err := time.ParseInLocation("2006-01-02 15:04:05", k.TglRegistrasi.Format("2006-01-02")+" "+k.JamReg, time.Local)
	return err == nil && waktu.Before(reg)
}

// Pemeriksaan SOAP & tanda vital (pemeriksaan_ralan / pemeriksaan_ranap). LingkarPerut hanya ada di rawat jalan.
type Pemeriksaan struct {
	NoRawat      string  `gorm:"column:no_rawat" json:"no_rawat"`
	TglPerawatan string  `gorm:"column:tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string  `gorm:"column:jam_rawat" json:"jam_rawat"`
	SuhuTubuh    *string `gorm:"column:suhu_tubuh" json:"suhu_tubuh"`
	Tensi        string  `gorm:"column:tensi" json:"tensi"`
	Nadi         *string `gorm:"column:nadi" json:"nadi"`
	Respirasi    *string `gorm:"column:respirasi" json:"respirasi"`
	Tinggi       *string `gorm:"column:tinggi" json:"tinggi"`
	Berat        *string `gorm:"column:berat" json:"berat"`
	Spo2         string  `gorm:"column:spo2" json:"spo2"`
	Gcs          *string `gorm:"column:gcs" json:"gcs"`
	Kesadaran    string  `gorm:"column:kesadaran" json:"kesadaran"`
	Keluhan      *string `gorm:"column:keluhan" json:"keluhan"`
	Pemeriksaan  *string `gorm:"column:pemeriksaan" json:"pemeriksaan"`
	Alergi       *string `gorm:"column:alergi" json:"alergi"`
	LingkarPerut *string `gorm:"column:lingkar_perut" json:"lingkar_perut,omitempty"`
	Rtl          string  `gorm:"column:rtl" json:"rtl"`
	Penilaian    string  `gorm:"column:penilaian" json:"penilaian"`
	Instruksi    string  `gorm:"column:instruksi" json:"instruksi"`
	Evaluasi     string  `gorm:"column:evaluasi" json:"evaluasi"`
	Nip          string  `gorm:"column:nip" json:"nip"`
	NmPegawai    string  `gorm:"column:nm_pegawai;->" json:"nm_pegawai"`
}

// Kosong true bila tidak ada satupun isian klinis (form Khanza menolak simpan).
func (p Pemeriksaan) Kosong() bool {
	for _, s := range []*string{p.SuhuTubuh, p.Nadi, p.Respirasi, p.Tinggi, p.Berat, p.Gcs, p.Keluhan, p.Pemeriksaan, p.Alergi, p.LingkarPerut} {
		if s != nil && *s != "" {
			return false
		}
	}
	for _, s := range []string{p.Tensi, p.Spo2, p.Rtl, p.Penilaian, p.Instruksi, p.Evaluasi} {
		if s != "" {
			return false
		}
	}
	return true
}
