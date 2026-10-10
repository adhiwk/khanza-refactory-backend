package rekammedis

import "time"

// ChecklistPostOperasi tabel `checklist_post_operasi` (checklist post operasi, RMChecklistPostOperasi).
type ChecklistPostOperasi struct {
	NoRawat                               string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                               *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Sncn                                  string     `gorm:"column:sncn" json:"sncn"`
	Tindakan                              string     `gorm:"column:tindakan" json:"tindakan"`
	KdDokterBedah                         string     `gorm:"column:kd_dokter_bedah" json:"kd_dokter_bedah"`
	KdDokterAnestesi                      string     `gorm:"column:kd_dokter_anestesi" json:"kd_dokter_anestesi"`
	KeadaanUmum                           *string    `gorm:"column:keadaan_umum" json:"keadaan_umum"`
	PemeriksaanPenunjangRontgen           *string    `gorm:"column:pemeriksaan_penunjang_rontgen" json:"pemeriksaan_penunjang_rontgen"`
	KeteranganPemeriksaanPenunjangRontgen *string    `gorm:"column:keterangan_pemeriksaan_penunjang_rontgen" json:"keterangan_pemeriksaan_penunjang_rontgen"`
	PemeriksaanPenunjangEkg               *string    `gorm:"column:pemeriksaan_penunjang_ekg" json:"pemeriksaan_penunjang_ekg"`
	KeteranganPemeriksaanPenunjangEkg     *string    `gorm:"column:keterangan_pemeriksaan_penunjang_ekg" json:"keterangan_pemeriksaan_penunjang_ekg"`
	PemeriksaanPenunjangUsg               *string    `gorm:"column:pemeriksaan_penunjang_usg" json:"pemeriksaan_penunjang_usg"`
	KeteranganPemeriksaanPenunjangUsg     *string    `gorm:"column:keterangan_pemeriksaan_penunjang_usg" json:"keterangan_pemeriksaan_penunjang_usg"`
	PemeriksaanPenunjangCtscan            *string    `gorm:"column:pemeriksaan_penunjang_ctscan" json:"pemeriksaan_penunjang_ctscan"`
	KeteranganPemeriksaanPenunjangCtscan  *string    `gorm:"column:keterangan_pemeriksaan_penunjang_ctscan" json:"keterangan_pemeriksaan_penunjang_ctscan"`
	PemeriksaanPenunjangMri               *string    `gorm:"column:pemeriksaan_penunjang_mri" json:"pemeriksaan_penunjang_mri"`
	KeteranganPemeriksaanPenunjangMri     *string    `gorm:"column:keterangan_pemeriksaan_penunjang_mri" json:"keterangan_pemeriksaan_penunjang_mri"`
	JenisCairanInfus                      *string    `gorm:"column:jenis_cairan_infus" json:"jenis_cairan_infus"`
	KateterUrine                          *string    `gorm:"column:kateter_urine" json:"kateter_urine"`
	TanggalPemasanganKateter              *time.Time `gorm:"column:tanggal_pemasangan_kateter" json:"tanggal_pemasangan_kateter"`
	WarnaKateter                          *string    `gorm:"column:warna_kateter" json:"warna_kateter"`
	JumlahKateter                         *string    `gorm:"column:jumlah_kateter" json:"jumlah_kateter"`
	AreaLukaOperasi                       *string    `gorm:"column:area_luka_operasi" json:"area_luka_operasi"`
	Drain                                 *string    `gorm:"column:drain" json:"drain"`
	JumlahDrain                           *string    `gorm:"column:jumlah_drain" json:"jumlah_drain"`
	LetakDrain                            *string    `gorm:"column:letak_drain" json:"letak_drain"`
	WarnaDrain                            *string    `gorm:"column:warna_drain" json:"warna_drain"`
	JaringanPa                            *string    `gorm:"column:jaringan_pa" json:"jaringan_pa"`
	NipPerawatOk                          *string    `gorm:"column:nip_perawat_ok" json:"nip_perawat_ok"`
	NipPerawatAnestesi                    *string    `gorm:"column:nip_perawat_anestesi" json:"nip_perawat_anestesi"`
}

func (ChecklistPostOperasi) TableName() string {
	return "checklist_post_operasi"
}
