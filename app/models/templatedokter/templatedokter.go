// Package templatedokter model template pemeriksaan dokter (MasterTemplatePemeriksaanDokter): SOAP + diagnosa,
// prosedur, permintaan lab/radiologi, resep, racikan, dan tindakan yang bisa dipakai ulang dokter.
package templatedokter

type Template struct {
	NoTemplate  string `gorm:"column:no_template" json:"no_template"`
	KdDokter    string `gorm:"column:kd_dokter" json:"kd_dokter"`
	NmDokter    string `gorm:"column:nm_dokter" json:"nm_dokter"`
	Keluhan     string `gorm:"column:keluhan" json:"keluhan"`
	Pemeriksaan string `gorm:"column:pemeriksaan" json:"pemeriksaan"`
	Penilaian   string `gorm:"column:penilaian" json:"penilaian"`
	Rencana     string `gorm:"column:rencana" json:"rencana"`
	Instruksi   string `gorm:"column:instruksi" json:"instruksi"`
	Evaluasi    string `gorm:"column:evaluasi" json:"evaluasi"`
	Isi         `gorm:"-"`
}

// Isi bagian detail template (tabel template_pemeriksaan_dokter_*).
type Isi struct {
	Diagnosa  []Kode    `gorm:"-" json:"diagnosa"`
	Prosedur  []Kode    `gorm:"-" json:"prosedur"`
	Radiologi []Kode    `gorm:"-" json:"radiologi"`
	Lab       []Lab     `gorm:"-" json:"lab"`
	Resep     []Resep   `gorm:"-" json:"resep"`
	Racikan   []Racikan `gorm:"-" json:"racikan"`
	Tindakan  []Kode    `gorm:"-" json:"tindakan"`
}

// Kode item berkode (ICD-10, ICD-9, radiologi, tindakan); Urut & Jumlah hanya untuk diagnosa/prosedur.
type Kode struct {
	Kode   string `gorm:"column:kode" json:"kode"`
	Nama   string `gorm:"column:nama" json:"nama"`
	Urut   int    `gorm:"column:urut" json:"urut,omitempty"`
	Jumlah string `gorm:"column:jumlah" json:"jumlah,omitempty"`
}

// Lab permintaan pemeriksaan lab beserta item template laboratorium yang dipilih.
type Lab struct {
	KdJenisPrw string `gorm:"column:kd_jenis_prw" json:"kd_jenis_prw"`
	Nama       string `gorm:"column:nama" json:"nama"`
	IDTemplate []int  `gorm:"-" json:"id_template"`
}

type Resep struct {
	KodeBrng    string  `gorm:"column:kode_brng" json:"kode_brng"`
	NamaBrng    string  `gorm:"column:nama_brng" json:"nama_brng"`
	Jml         float64 `gorm:"column:jml" json:"jml"`
	AturanPakai string  `gorm:"column:aturan_pakai" json:"aturan_pakai"`
}

type Racikan struct {
	NoRacik     string          `gorm:"column:no_racik" json:"no_racik"`
	NamaRacik   string          `gorm:"column:nama_racik" json:"nama_racik"`
	KdRacik     string          `gorm:"column:kd_racik" json:"kd_racik"`
	JmlDr       int             `gorm:"column:jml_dr" json:"jml_dr"`
	AturanPakai string          `gorm:"column:aturan_pakai" json:"aturan_pakai"`
	Keterangan  string          `gorm:"column:keterangan" json:"keterangan"`
	Detail      []RacikanDetail `gorm:"-" json:"detail"`
}

type RacikanDetail struct {
	KodeBrng  string  `gorm:"column:kode_brng" json:"kode_brng"`
	NamaBrng  string  `gorm:"column:nama_brng" json:"nama_brng"`
	P1        float64 `gorm:"column:p1" json:"p1"`
	P2        float64 `gorm:"column:p2" json:"p2"`
	Kandungan string  `gorm:"column:kandungan" json:"kandungan"`
	Jml       float64 `gorm:"column:jml" json:"jml"`
}
