package rekammedis

// ResumePasien tabel `resume_pasien` (resume pasien, RMDataResumePasien).
type ResumePasien struct {
	NoRawat              string `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	KdDokter             string `gorm:"column:kd_dokter" json:"kd_dokter"`
	KeluhanUtama         string `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	JalannyaPenyakit     string `gorm:"column:jalannya_penyakit" json:"jalannya_penyakit"`
	PemeriksaanPenunjang string `gorm:"column:pemeriksaan_penunjang" json:"pemeriksaan_penunjang"`
	HasilLaborat         string `gorm:"column:hasil_laborat" json:"hasil_laborat"`
	DiagnosaUtama        string `gorm:"column:diagnosa_utama" json:"diagnosa_utama"`
	KdDiagnosaUtama      string `gorm:"column:kd_diagnosa_utama" json:"kd_diagnosa_utama"`
	DiagnosaSekunder     string `gorm:"column:diagnosa_sekunder" json:"diagnosa_sekunder"`
	KdDiagnosaSekunder   string `gorm:"column:kd_diagnosa_sekunder" json:"kd_diagnosa_sekunder"`
	DiagnosaSekunder2    string `gorm:"column:diagnosa_sekunder2" json:"diagnosa_sekunder2"`
	KdDiagnosaSekunder2  string `gorm:"column:kd_diagnosa_sekunder2" json:"kd_diagnosa_sekunder2"`
	DiagnosaSekunder3    string `gorm:"column:diagnosa_sekunder3" json:"diagnosa_sekunder3"`
	KdDiagnosaSekunder3  string `gorm:"column:kd_diagnosa_sekunder3" json:"kd_diagnosa_sekunder3"`
	DiagnosaSekunder4    string `gorm:"column:diagnosa_sekunder4" json:"diagnosa_sekunder4"`
	KdDiagnosaSekunder4  string `gorm:"column:kd_diagnosa_sekunder4" json:"kd_diagnosa_sekunder4"`
	ProsedurUtama        string `gorm:"column:prosedur_utama" json:"prosedur_utama"`
	KdProsedurUtama      string `gorm:"column:kd_prosedur_utama" json:"kd_prosedur_utama"`
	ProsedurSekunder     string `gorm:"column:prosedur_sekunder" json:"prosedur_sekunder"`
	KdProsedurSekunder   string `gorm:"column:kd_prosedur_sekunder" json:"kd_prosedur_sekunder"`
	ProsedurSekunder2    string `gorm:"column:prosedur_sekunder2" json:"prosedur_sekunder2"`
	KdProsedurSekunder2  string `gorm:"column:kd_prosedur_sekunder2" json:"kd_prosedur_sekunder2"`
	ProsedurSekunder3    string `gorm:"column:prosedur_sekunder3" json:"prosedur_sekunder3"`
	KdProsedurSekunder3  string `gorm:"column:kd_prosedur_sekunder3" json:"kd_prosedur_sekunder3"`
	KondisiPulang        string `gorm:"column:kondisi_pulang" json:"kondisi_pulang"`
	ObatPulang           string `gorm:"column:obat_pulang" json:"obat_pulang"`
}

func (ResumePasien) TableName() string {
	return "resume_pasien"
}
