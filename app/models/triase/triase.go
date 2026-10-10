// Package triase model triase IGD (data_triase_igd, data_triase_igdprimer/sekunder, data_triase_igddetail_skala1..5).
package triase

const (
	Primer   = "primer"   // zona merah: skala 1 / 2
	Sekunder = "sekunder" // skala 3 / 4 / 5
)

// Ringkas baris daftar triase.
type Ringkas struct {
	NoRawat      string `gorm:"column:no_rawat" json:"no_rawat"`
	NoRkmMedis   string `gorm:"column:no_rkm_medis" json:"no_rkm_medis"`
	NmPasien     string `gorm:"column:nm_pasien" json:"nm_pasien"`
	TglKunjungan string `gorm:"column:tgl_kunjungan" json:"tgl_kunjungan"`
	KodeKasus    string `gorm:"column:kode_kasus" json:"kode_kasus"`
	MacamKasus   string `gorm:"column:macam_kasus" json:"macam_kasus"`
	Jenis        string `gorm:"column:jenis" json:"jenis"`
	Plan         string `gorm:"column:plan" json:"plan"`
	Nik          string `gorm:"column:nik" json:"nik"`
	NmPetugas    string `gorm:"column:nm_petugas" json:"nm_petugas"`
}

// Utama kolom data_triase_igd (tanpa id SatuSehat).
type Utama struct {
	NoRawat              string `gorm:"column:no_rawat" json:"no_rawat"`
	TglKunjungan         string `gorm:"column:tgl_kunjungan" json:"tgl_kunjungan"`
	CaraMasuk            string `gorm:"column:cara_masuk" json:"cara_masuk"`
	AlatTransportasi     string `gorm:"column:alat_transportasi" json:"alat_transportasi"`
	AlasanKedatangan     string `gorm:"column:alasan_kedatangan" json:"alasan_kedatangan"`
	KeteranganKedatangan string `gorm:"column:keterangan_kedatangan" json:"keterangan_kedatangan"`
	KodeKasus            string `gorm:"column:kode_kasus" json:"kode_kasus"`
	MacamKasus           string `gorm:"column:macam_kasus" json:"macam_kasus"`
	TekananDarah         string `gorm:"column:tekanan_darah" json:"tekanan_darah"`
	Nadi                 string `gorm:"column:nadi" json:"nadi"`
	Pernapasan           string `gorm:"column:pernapasan" json:"pernapasan"`
	Suhu                 string `gorm:"column:suhu" json:"suhu"`
	SaturasiO2           string `gorm:"column:saturasi_o2" json:"saturasi_o2"`
	Nyeri                string `gorm:"column:nyeri" json:"nyeri"`
}

// Penilaian bagian primer atau sekunder. Keluhan = keluhan_utama (primer) / anamnesa_singkat (sekunder);
// KebutuhanKhusus hanya untuk primer.
type Penilaian struct {
	Keluhan         string `gorm:"column:keluhan" json:"keluhan"`
	KebutuhanKhusus string `gorm:"column:kebutuhan_khusus" json:"kebutuhan_khusus,omitempty"`
	Catatan         string `gorm:"column:catatan" json:"catatan"`
	Plan            string `gorm:"column:plan" json:"plan"`
	TanggalTriase   string `gorm:"column:tanggaltriase" json:"tanggaltriase"`
	Nik             string `gorm:"column:nik" json:"nik"`
	NmPetugas       string `gorm:"column:nm_petugas" json:"nm_petugas"`
}

type SkalaItem struct {
	Kode            string `gorm:"column:kode" json:"kode"`
	Pengkajian      string `gorm:"column:pengkajian" json:"pengkajian"`
	KodePemeriksaan string `gorm:"column:kode_pemeriksaan" json:"kode_pemeriksaan"`
	NamaPemeriksaan string `gorm:"column:nama_pemeriksaan" json:"nama_pemeriksaan"`
}

type Skala struct {
	Level int         `json:"level"`
	Items []SkalaItem `json:"items"`
}

// Triase data lengkap satu triase.
type Triase struct {
	Utama
	NoRkmMedis string     `json:"no_rkm_medis"`
	NmPasien   string     `json:"nm_pasien"`
	Jenis      string     `json:"jenis"`
	Penilaian  *Penilaian `json:"penilaian"`
	Skala      Skala      `json:"skala"`
}

// Input simpan/ubah triase.
type Input struct {
	Utama
	Jenis      string
	Penilaian  Penilaian
	SkalaLevel int
	SkalaKode  []string
}
