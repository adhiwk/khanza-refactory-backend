package rekammedis

import "time"

// PenilaianPreInduksi tabel `penilaian_pre_induksi` (penilaian pre induksi, RMPenilaianPreInduksi).
type PenilaianPreInduksi struct {
	NoRawat                      string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                      *time.Time `gorm:"column:tanggal;primaryKey;autoIncrement:false" json:"tanggal"`
	KdDokter                     string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	Tensi                        *string    `gorm:"column:tensi" json:"tensi"`
	Nadi                         *string    `gorm:"column:nadi" json:"nadi"`
	Rr                           *string    `gorm:"column:rr" json:"rr"`
	Suhu                         *string    `gorm:"column:suhu" json:"suhu"`
	Ekg                          *string    `gorm:"column:ekg" json:"ekg"`
	LainLain                     *string    `gorm:"column:lain_lain" json:"lain_lain"`
	Asesmen                      *string    `gorm:"column:asesmen" json:"asesmen"`
	Perencanaan                  *string    `gorm:"column:perencanaan" json:"perencanaan"`
	InfusPerifier                *string    `gorm:"column:infus_perifier" json:"infus_perifier"`
	Cvc                          *string    `gorm:"column:cvc" json:"cvc"`
	Posisi                       *string    `gorm:"column:posisi" json:"posisi"`
	Premedikasi                  *string    `gorm:"column:premedikasi" json:"premedikasi"`
	PremedikasiKeterangan        *string    `gorm:"column:premedikasi_keterangan" json:"premedikasi_keterangan"`
	Induksi                      *string    `gorm:"column:induksi" json:"induksi"`
	InduksiKeterangan            *string    `gorm:"column:induksi_keterangan" json:"induksi_keterangan"`
	FaceMaskNo                   *string    `gorm:"column:face_mask_no" json:"face_mask_no"`
	NasopharingNo                *string    `gorm:"column:nasopharing_no" json:"nasopharing_no"`
	EttNo                        *string    `gorm:"column:ett_no" json:"ett_no"`
	EttJenis                     *string    `gorm:"column:ett_jenis" json:"ett_jenis"`
	EttViksasi                   *string    `gorm:"column:ett_viksasi" json:"ett_viksasi"`
	LmaNo                        *string    `gorm:"column:lma_no" json:"lma_no"`
	LmaJenis                     *string    `gorm:"column:lma_jenis" json:"lma_jenis"`
	Tracheostomi                 *string    `gorm:"column:tracheostomi" json:"tracheostomi"`
	BronchoscopiFiberoptik       *string    `gorm:"column:bronchoscopi_fiberoptik" json:"bronchoscopi_fiberoptik"`
	Glidescopi                   *string    `gorm:"column:glidescopi" json:"glidescopi"`
	LainLainTatalaksana          *string    `gorm:"column:lain_lain_tatalaksana" json:"lain_lain_tatalaksana"`
	IntubasiSesudahTidur         *string    `gorm:"column:intubasi_sesudah_tidur" json:"intubasi_sesudah_tidur"`
	IntubasiOral                 *string    `gorm:"column:intubasi_oral" json:"intubasi_oral"`
	IntubasiTracheostomi         *string    `gorm:"column:intubasi_tracheostomi" json:"intubasi_tracheostomi"`
	IntubasiKeterangan           *string    `gorm:"column:intubasi_keterangan" json:"intubasi_keterangan"`
	SulitVentilasi               *string    `gorm:"column:sulit_ventilasi" json:"sulit_ventilasi"`
	SulitIntubasi                *string    `gorm:"column:sulit_intubasi" json:"sulit_intubasi"`
	Ventilasi                    *string    `gorm:"column:ventilasi" json:"ventilasi"`
	TeknikRegionalJenis          *string    `gorm:"column:teknik_regional_jenis" json:"teknik_regional_jenis"`
	TeknikRegionalLokasi         *string    `gorm:"column:teknik_regional_lokasi" json:"teknik_regional_lokasi"`
	TeknikRegionalJenisJarum     *string    `gorm:"column:teknik_regional_jenis_jarum" json:"teknik_regional_jenis_jarum"`
	TeknikRegionalKateter        *string    `gorm:"column:teknik_regional_kateter" json:"teknik_regional_kateter"`
	TeknikRegionalKateterViksasi *string    `gorm:"column:teknik_regional_kateter_viksasi" json:"teknik_regional_kateter_viksasi"`
	TeknikRegionalObatObatan     *string    `gorm:"column:teknik_regional_obat_obatan" json:"teknik_regional_obat_obatan"`
	TeknikRegionalKomplikasi     *string    `gorm:"column:teknik_regional_komplikasi" json:"teknik_regional_komplikasi"`
	TeknikRegionalHasil          *string    `gorm:"column:teknik_regional_hasil" json:"teknik_regional_hasil"`
}

func (PenilaianPreInduksi) TableName() string {
	return "penilaian_pre_induksi"
}
