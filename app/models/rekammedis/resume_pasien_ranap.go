package rekammedis

import "time"

// ResumePasienRanap tabel `resume_pasien_ranap` (resume pasien ranap, RMDataResumePasienRanap).
type ResumePasienRanap struct {
	NoRawat              string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	KdDokter             string     `gorm:"column:kd_dokter" json:"kd_dokter"`
	DiagnosaAwal         string     `gorm:"column:diagnosa_awal" json:"diagnosa_awal"`
	Alasan               string     `gorm:"column:alasan" json:"alasan"`
	KeluhanUtama         string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	PemeriksaanFisik     string     `gorm:"column:pemeriksaan_fisik" json:"pemeriksaan_fisik"`
	JalannyaPenyakit     string     `gorm:"column:jalannya_penyakit" json:"jalannya_penyakit"`
	PemeriksaanPenunjang string     `gorm:"column:pemeriksaan_penunjang" json:"pemeriksaan_penunjang"`
	HasilLaborat         string     `gorm:"column:hasil_laborat" json:"hasil_laborat"`
	TindakanDanOperasi   string     `gorm:"column:tindakan_dan_operasi" json:"tindakan_dan_operasi"`
	ObatDiRs             string     `gorm:"column:obat_di_rs" json:"obat_di_rs"`
	DiagnosaUtama        string     `gorm:"column:diagnosa_utama" json:"diagnosa_utama"`
	KdDiagnosaUtama      string     `gorm:"column:kd_diagnosa_utama" json:"kd_diagnosa_utama"`
	DiagnosaSekunder     string     `gorm:"column:diagnosa_sekunder" json:"diagnosa_sekunder"`
	KdDiagnosaSekunder   string     `gorm:"column:kd_diagnosa_sekunder" json:"kd_diagnosa_sekunder"`
	DiagnosaSekunder2    string     `gorm:"column:diagnosa_sekunder2" json:"diagnosa_sekunder2"`
	KdDiagnosaSekunder2  string     `gorm:"column:kd_diagnosa_sekunder2" json:"kd_diagnosa_sekunder2"`
	DiagnosaSekunder3    string     `gorm:"column:diagnosa_sekunder3" json:"diagnosa_sekunder3"`
	KdDiagnosaSekunder3  string     `gorm:"column:kd_diagnosa_sekunder3" json:"kd_diagnosa_sekunder3"`
	DiagnosaSekunder4    string     `gorm:"column:diagnosa_sekunder4" json:"diagnosa_sekunder4"`
	KdDiagnosaSekunder4  string     `gorm:"column:kd_diagnosa_sekunder4" json:"kd_diagnosa_sekunder4"`
	ProsedurUtama        string     `gorm:"column:prosedur_utama" json:"prosedur_utama"`
	KdProsedurUtama      string     `gorm:"column:kd_prosedur_utama" json:"kd_prosedur_utama"`
	ProsedurSekunder     string     `gorm:"column:prosedur_sekunder" json:"prosedur_sekunder"`
	KdProsedurSekunder   string     `gorm:"column:kd_prosedur_sekunder" json:"kd_prosedur_sekunder"`
	ProsedurSekunder2    string     `gorm:"column:prosedur_sekunder2" json:"prosedur_sekunder2"`
	KdProsedurSekunder2  string     `gorm:"column:kd_prosedur_sekunder2" json:"kd_prosedur_sekunder2"`
	ProsedurSekunder3    string     `gorm:"column:prosedur_sekunder3" json:"prosedur_sekunder3"`
	KdProsedurSekunder3  string     `gorm:"column:kd_prosedur_sekunder3" json:"kd_prosedur_sekunder3"`
	Alergi               string     `gorm:"column:alergi" json:"alergi"`
	Diet                 string     `gorm:"column:diet" json:"diet"`
	LabBelum             string     `gorm:"column:lab_belum" json:"lab_belum"`
	Edukasi              string     `gorm:"column:edukasi" json:"edukasi"`
	CaraKeluar           string     `gorm:"column:cara_keluar" json:"cara_keluar"`
	KetKeluar            *string    `gorm:"column:ket_keluar" json:"ket_keluar"`
	Keadaan              string     `gorm:"column:keadaan" json:"keadaan"`
	KetKeadaan           *string    `gorm:"column:ket_keadaan" json:"ket_keadaan"`
	Dilanjutkan          string     `gorm:"column:dilanjutkan" json:"dilanjutkan"`
	KetDilanjutkan       *string    `gorm:"column:ket_dilanjutkan" json:"ket_dilanjutkan"`
	Kontrol              *time.Time `gorm:"column:kontrol" json:"kontrol"`
	ObatPulang           string     `gorm:"column:obat_pulang" json:"obat_pulang"`
}

func (ResumePasienRanap) TableName() string {
	return "resume_pasien_ranap"
}
