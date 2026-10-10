package rekammedis

import "time"

// PenilaianTambahanBeresikoMelarikanDiri tabel `penilaian_tambahan_beresiko_melarikan_diri` (penilaian tambahan melarikan diri, RMPenilaianTambahanMelarikanDiri).
type PenilaianTambahanBeresikoMelarikanDiri struct {
	NoRawat                                string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                                *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip                                    string     `gorm:"column:nip" json:"nip"`
	StatikRiwayatMelarikanDiri             *string    `gorm:"column:statik_riwayat_melarikan_diri" json:"statik_riwayat_melarikan_diri"`
	StatikSkorriwayatMelarikanDiri         *int       `gorm:"column:statik_skorriwayat_melarikan_diri" json:"statik_skorriwayat_melarikan_diri"`
	StatikRiwayatPenolakanPengobatan       *string    `gorm:"column:statik_riwayat_penolakan_pengobatan" json:"statik_riwayat_penolakan_pengobatan"`
	StatikSkorriwayatPenolakanPengobatan   *int       `gorm:"column:statik_skorriwayat_penolakan_pengobatan" json:"statik_skorriwayat_penolakan_pengobatan"`
	StatikUsiaDibawah35                    *string    `gorm:"column:statik_usia_dibawah_35" json:"statik_usia_dibawah_35"`
	StatikSkorusiaDibawah35                *int       `gorm:"column:statik_skorusia_dibawah_35" json:"statik_skorusia_dibawah_35"`
	StatikLakiLaki                         *string    `gorm:"column:statik_laki_laki" json:"statik_laki_laki"`
	StatikSkorlakiLaki                     *int       `gorm:"column:statik_skorlaki_laki" json:"statik_skorlaki_laki"`
	StatikDiagnosisSkizofrenia             *string    `gorm:"column:statik_diagnosis_skizofrenia" json:"statik_diagnosis_skizofrenia"`
	StatikSkordiagnosisSkizofrenia         *int       `gorm:"column:statik_skordiagnosis_skizofrenia" json:"statik_skordiagnosis_skizofrenia"`
	StatikBelumMenikah                     *string    `gorm:"column:statik_belum_menikah" json:"statik_belum_menikah"`
	StatikSkorbelumMenikah                 *int       `gorm:"column:statik_skorbelum_menikah" json:"statik_skorbelum_menikah"`
	StatikRiwayatPenggunaanNapza           *string    `gorm:"column:statik_riwayat_penggunaan_napza" json:"statik_riwayat_penggunaan_napza"`
	StatikSkoriwayatPenggunaanNapza        *int       `gorm:"column:statik_skoriwayat_penggunaan_napza" json:"statik_skoriwayat_penggunaan_napza"`
	StatikDiagnosisGangguanKepribadian     *string    `gorm:"column:statik_diagnosis_gangguan_kepribadian" json:"statik_diagnosis_gangguan_kepribadian"`
	StatikSkordiagnosisGangguanKepribadian *int       `gorm:"column:statik_skordiagnosis_gangguan_kepribadian" json:"statik_skordiagnosis_gangguan_kepribadian"`
	StatikRiwayatKriminal                  *string    `gorm:"column:statik_riwayat_kriminal" json:"statik_riwayat_kriminal"`
	StatikSkorriwayatKriminal              *int       `gorm:"column:statik_skorriwayat_kriminal" json:"statik_skorriwayat_kriminal"`
	StatikSkortotal                        *int       `gorm:"column:statik_skortotal" json:"statik_skortotal"`
	DinamisAntiTreatment                   *string    `gorm:"column:dinamis_anti_treatment" json:"dinamis_anti_treatment"`
	DinamisSkorantiTreatment               *int       `gorm:"column:dinamis_skoranti_treatment" json:"dinamis_skoranti_treatment"`
	DinamisPenggunaanNapza                 *string    `gorm:"column:dinamis_penggunaan_napza" json:"dinamis_penggunaan_napza"`
	DinamisSkorpenggunaanNapza             *int       `gorm:"column:dinamis_skorpenggunaan_napza" json:"dinamis_skorpenggunaan_napza"`
	DinamisKebosanan                       *string    `gorm:"column:dinamis_kebosanan" json:"dinamis_kebosanan"`
	DinamisSkorkebosanan                   *int       `gorm:"column:dinamis_skorkebosanan" json:"dinamis_skorkebosanan"`
	DinamisPerintahHalusinasi              *string    `gorm:"column:dinamis_perintah_halusinasi" json:"dinamis_perintah_halusinasi"`
	DinamisSkorperintahHalusinasi          *int       `gorm:"column:dinamis_skorperintah_halusinasi" json:"dinamis_skorperintah_halusinasi"`
	DinamisHilangnyaKontrolDiri            *string    `gorm:"column:dinamis_hilangnya_kontrol_diri" json:"dinamis_hilangnya_kontrol_diri"`
	DinamisSkorhilangnyaKontrolDiri        *int       `gorm:"column:dinamis_skorhilangnya_kontrol_diri" json:"dinamis_skorhilangnya_kontrol_diri"`
	DinamisSeksualTidakWajar               *string    `gorm:"column:dinamis_seksual_tidak_wajar" json:"dinamis_seksual_tidak_wajar"`
	DinamisSkorseksualTidakWajar           *int       `gorm:"column:dinamis_skorseksual_tidak_wajar" json:"dinamis_skorseksual_tidak_wajar"`
	DinamisKemarahanFrustasi               *string    `gorm:"column:dinamis_kemarahan_frustasi" json:"dinamis_kemarahan_frustasi"`
	DinamisSkorkemarahanFrustasi           *int       `gorm:"column:dinamis_skorkemarahan_frustasi" json:"dinamis_skorkemarahan_frustasi"`
	DinamisKetakutanPerawatan              *string    `gorm:"column:dinamis_ketakutan_perawatan" json:"dinamis_ketakutan_perawatan"`
	DinamisSkorketakutanPerawatan          *int       `gorm:"column:dinamis_skorketakutan_perawatan" json:"dinamis_skorketakutan_perawatan"`
	DinamisSkortotal                       *int       `gorm:"column:dinamis_skortotal" json:"dinamis_skortotal"`
	FaktorFaktorPencegahan                 *string    `gorm:"column:faktor_faktor_pencegahan" json:"faktor_faktor_pencegahan"`
	TotalSkor                              *int       `gorm:"column:total_skor" json:"total_skor"`
	LevelSkor                              *string    `gorm:"column:level_skor" json:"level_skor"`
}

func (PenilaianTambahanBeresikoMelarikanDiri) TableName() string {
	return "penilaian_tambahan_beresiko_melarikan_diri"
}
