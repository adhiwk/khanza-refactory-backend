package rekammedis

import "time"

// PenilaianTambahanBunuhDiri tabel `penilaian_tambahan_bunuh_diri` (penilaian tambahan bunuh diri, RMPenilaianTambahanBunuhDiri).
type PenilaianTambahanBunuhDiri struct {
	NoRawat                        string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                        *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip                            string     `gorm:"column:nip" json:"nip"`
	StatikHidupSendiri             *string    `gorm:"column:statik_hidup_sendiri" json:"statik_hidup_sendiri"`
	StatikSkorhidupSendiri         *int       `gorm:"column:statik_skorhidup_sendiri" json:"statik_skorhidup_sendiri"`
	StatikUpayaSuicide             *string    `gorm:"column:statik_upaya_suicide" json:"statik_upaya_suicide"`
	StatikSkorupayaSuicide         *int       `gorm:"column:statik_skorupaya_suicide" json:"statik_skorupaya_suicide"`
	StatikKeluargaSuicide          *string    `gorm:"column:statik_keluarga_suicide" json:"statik_keluarga_suicide"`
	StatikSkorkeluargaSuicide      *int       `gorm:"column:statik_skorkeluarga_suicide" json:"statik_skorkeluarga_suicide"`
	StatikDiagnosaGangguanJiwa     *string    `gorm:"column:statik_diagnosa_gangguan_jiwa" json:"statik_diagnosa_gangguan_jiwa"`
	StatikSkordiagnosaGangguanJiwa *int       `gorm:"column:statik_skordiagnosa_gangguan_jiwa" json:"statik_skordiagnosa_gangguan_jiwa"`
	StatikDisabilitasBerat         *string    `gorm:"column:statik_disabilitas_berat" json:"statik_disabilitas_berat"`
	StatikSkordisabilitasBerat     *int       `gorm:"column:statik_skordisabilitas_berat" json:"statik_skordisabilitas_berat"`
	StatikBerpisah                 *string    `gorm:"column:statik_berpisah" json:"statik_berpisah"`
	StatikSkorberpisah             *int       `gorm:"column:statik_skorberpisah" json:"statik_skorberpisah"`
	StatikKehilanganKerja          *string    `gorm:"column:statik_kehilangan_kerja" json:"statik_kehilangan_kerja"`
	StatikSkorkehilanganKerja      *int       `gorm:"column:statik_skorkehilangan_kerja" json:"statik_skorkehilangan_kerja"`
	StatikSkortotal                *int       `gorm:"column:statik_skortotal" json:"statik_skortotal"`
	DinamisIdeBunuhDiri            *string    `gorm:"column:dinamis_ide_bunuh_diri" json:"dinamis_ide_bunuh_diri"`
	DinamisSkorideBunuhDiri        *int       `gorm:"column:dinamis_skoride_bunuh_diri" json:"dinamis_skoride_bunuh_diri"`
	DinamisMaksudSuicide           *string    `gorm:"column:dinamis_maksud_suicide" json:"dinamis_maksud_suicide"`
	DinamisSkormaksudSuicide       *int       `gorm:"column:dinamis_skormaksud_suicide" json:"dinamis_skormaksud_suicide"`
	DinamisStressBerat             *string    `gorm:"column:dinamis_stress_berat" json:"dinamis_stress_berat"`
	DinamisSkorstressBerat         *int       `gorm:"column:dinamis_skorstress_berat" json:"dinamis_skorstress_berat"`
	DinamisKeputusasaan            *string    `gorm:"column:dinamis_keputusasaan" json:"dinamis_keputusasaan"`
	DinamisSkorkeputusasaan        *int       `gorm:"column:dinamis_skorkeputusasaan" json:"dinamis_skorkeputusasaan"`
	DinamisKejadianSignifikan      *string    `gorm:"column:dinamis_kejadian_signifikan" json:"dinamis_kejadian_signifikan"`
	DinamisSkorkejadianSignifikan  *int       `gorm:"column:dinamis_skorkejadian_signifikan" json:"dinamis_skorkejadian_signifikan"`
	DinamisKehilanganKontrol       *string    `gorm:"column:dinamis_kehilangan_kontrol" json:"dinamis_kehilangan_kontrol"`
	DinamisSkorkehilanganKontrol   *int       `gorm:"column:dinamis_skorkehilangan_kontrol" json:"dinamis_skorkehilangan_kontrol"`
	DinamisPenggunaanNapza         *string    `gorm:"column:dinamis_penggunaan_napza" json:"dinamis_penggunaan_napza"`
	DinamisSkorpenggunaanNapza     *int       `gorm:"column:dinamis_skorpenggunaan_napza" json:"dinamis_skorpenggunaan_napza"`
	DinamisSkortotal               *int       `gorm:"column:dinamis_skortotal" json:"dinamis_skortotal"`
	FaktorFaktorPencegahan         *string    `gorm:"column:faktor_faktor_pencegahan" json:"faktor_faktor_pencegahan"`
	TotalSkor                      *int       `gorm:"column:total_skor" json:"total_skor"`
	LevelSkor                      *string    `gorm:"column:level_skor" json:"level_skor"`
}

func (PenilaianTambahanBunuhDiri) TableName() string {
	return "penilaian_tambahan_bunuh_diri"
}
