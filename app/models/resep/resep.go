// Package resep model resep elektronik dokter (resep_obat + resep_dokter).
package resep

// Resep header resep; tgl_perawatan '0000-00-00' berarti belum dilayani farmasi.
type Resep struct {
	NoResep      string `gorm:"column:no_resep" json:"no_resep"`
	NoRawat      string `gorm:"column:no_rawat" json:"no_rawat"`
	KdDokter     string `gorm:"column:kd_dokter" json:"kd_dokter"`
	NmDokter     string `gorm:"column:nm_dokter" json:"nm_dokter"`
	TglPeresepan string `gorm:"column:tgl_peresepan" json:"tgl_peresepan"`
	JamPeresepan string `gorm:"column:jam_peresepan" json:"jam_peresepan"`
	TglPerawatan string `gorm:"column:tgl_perawatan" json:"tgl_perawatan"`
	Jam          string `gorm:"column:jam" json:"jam"`
	Status       string `gorm:"column:status" json:"status"`
	Items        []Item `gorm:"-" json:"items"`
}

// Dilayani true bila farmasi sudah memproses resep.
func (r Resep) Dilayani() bool {
	return r.TglPerawatan != "" && r.TglPerawatan != "0000-00-00"
}

type Item struct {
	KodeBrng    string  `gorm:"column:kode_brng" json:"kode_brng"`
	NamaBrng    string  `gorm:"column:nama_brng" json:"nama_brng"`
	Jml         float64 `gorm:"column:jml" json:"jml"`
	AturanPakai string  `gorm:"column:aturan_pakai" json:"aturan_pakai"`
}
