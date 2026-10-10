package rekammedis

import "time"

// ChecklistPreOperasi tabel `checklist_pre_operasi` (checklist pre operasi, RMChecklistPreOperasi).
type ChecklistPreOperasi struct {
	NoRawat                               string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                               *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	Sncn                                  string     `gorm:"column:sncn" json:"sncn"`
	Tindakan                              string     `gorm:"column:tindakan" json:"tindakan"`
	KdDokterBedah                         string     `gorm:"column:kd_dokter_bedah" json:"kd_dokter_bedah"`
	KdDokterAnestesi                      string     `gorm:"column:kd_dokter_anestesi" json:"kd_dokter_anestesi"`
	Identitas                             *string    `gorm:"column:identitas" json:"identitas"`
	SuratIjinBedah                        *string    `gorm:"column:surat_ijin_bedah" json:"surat_ijin_bedah"`
	SuratIjinAnestesi                     *string    `gorm:"column:surat_ijin_anestesi" json:"surat_ijin_anestesi"`
	SuratIjinTransfusi                    *string    `gorm:"column:surat_ijin_transfusi" json:"surat_ijin_transfusi"`
	PenandaanAreaOperasi                  *string    `gorm:"column:penandaan_area_operasi" json:"penandaan_area_operasi"`
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
	PersiapanDarah                        *string    `gorm:"column:persiapan_darah" json:"persiapan_darah"`
	KeteranganPersiapanDarah              *string    `gorm:"column:keterangan_persiapan_darah" json:"keterangan_persiapan_darah"`
	PerlengkapanKhusus                    *string    `gorm:"column:perlengkapan_khusus" json:"perlengkapan_khusus"`
	NipPetugasRuangan                     *string    `gorm:"column:nip_petugas_ruangan" json:"nip_petugas_ruangan"`
	NipPerawatOk                          *string    `gorm:"column:nip_perawat_ok" json:"nip_perawat_ok"`
}

func (ChecklistPreOperasi) TableName() string {
	return "checklist_pre_operasi"
}
