package rekammedis

import "time"

// PengkajianRestrain tabel `pengkajian_restrain` (pengkajian restrain, RMPengkajianRestrain).
type PengkajianRestrain struct {
	NoRawat                          string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                          *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip                              string     `gorm:"column:nip" json:"nip"`
	Gcs                              *string    `gorm:"column:gcs" json:"gcs"`
	ReflekaCahayaKa                  *string    `gorm:"column:refleka_cahaya_ka" json:"refleka_cahaya_ka"`
	ReflekaCahayaKi                  *string    `gorm:"column:refleka_cahaya_ki" json:"refleka_cahaya_ki"`
	UkuranPupilKa                    *string    `gorm:"column:ukuran_pupil_ka" json:"ukuran_pupil_ka"`
	UkuranPupilKi                    *string    `gorm:"column:ukuran_pupil_ki" json:"ukuran_pupil_ki"`
	Td                               *string    `gorm:"column:td" json:"td"`
	Suhu                             *string    `gorm:"column:suhu" json:"suhu"`
	Rr                               *string    `gorm:"column:rr" json:"rr"`
	Nadi                             *string    `gorm:"column:nadi" json:"nadi"`
	HasilObservasi                   *string    `gorm:"column:hasil_observasi" json:"hasil_observasi"`
	PertimbanganKlinis               *string    `gorm:"column:pertimbangan_klinis" json:"pertimbangan_klinis"`
	RestrainNonFarmakologi           *string    `gorm:"column:restrain_non_farmakologi" json:"restrain_non_farmakologi"`
	RestrainNonFarmakologiKeterangan *string    `gorm:"column:restrain_non_farmakologi_keterangan" json:"restrain_non_farmakologi_keterangan"`
	RestrainFarmakologi              *string    `gorm:"column:restrain_farmakologi" json:"restrain_farmakologi"`
	SudahDijelaskanKeluarga          string     `gorm:"column:sudah_dijelaskan_keluarga" json:"sudah_dijelaskan_keluarga"`
	KeluargaYangMenyetujui           *string    `gorm:"column:keluarga_yang_menyetujui" json:"keluarga_yang_menyetujui"`
}

func (PengkajianRestrain) TableName() string {
	return "pengkajian_restrain"
}
