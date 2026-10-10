package rekammedis

import "time"

// PenilaianTambahanPerilakuKekerasan tabel `penilaian_tambahan_perilaku_kekerasan` (penilaian tambahan perilaku kekerasan, RMPenilaianTambahanPerilakuKekerasan).
type PenilaianTambahanPerilakuKekerasan struct {
	NoRawat                            string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                            *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Nip                                string     `gorm:"column:nip" json:"nip"`
	StatikInsidenKekerasanBaruIni      *string    `gorm:"column:statik_insiden_kekerasan_baru_ini" json:"statik_insiden_kekerasan_baru_ini"`
	StatikSkorinsidenKekerasanBaruIni  *int       `gorm:"column:statik_skorinsiden_kekerasan_baru_ini" json:"statik_skorinsiden_kekerasan_baru_ini"`
	StatikRiwayatPenggunaanSenjata     *string    `gorm:"column:statik_riwayat_penggunaan_senjata" json:"statik_riwayat_penggunaan_senjata"`
	StatikSkorriwayatPenggunaanSenjata *int       `gorm:"column:statik_skorriwayat_penggunaan_senjata" json:"statik_skorriwayat_penggunaan_senjata"`
	StatikLakiLaki                     *string    `gorm:"column:statik_laki_laki" json:"statik_laki_laki"`
	StatikSkorlakiLaki                 *int       `gorm:"column:statik_skorlaki_laki" json:"statik_skorlaki_laki"`
	StatikUsiaDibawah35                *string    `gorm:"column:statik_usia_dibawah_35" json:"statik_usia_dibawah_35"`
	StatikSkorusiaDibawah35            *int       `gorm:"column:statik_skorusia_dibawah_35" json:"statik_skorusia_dibawah_35"`
	StatikRiwayatKriminal              *string    `gorm:"column:statik_riwayat_kriminal" json:"statik_riwayat_kriminal"`
	StatikSkorriwayatKriminal          *int       `gorm:"column:statik_skorriwayat_kriminal" json:"statik_skorriwayat_kriminal"`
	StatikIdeKekerasan                 *string    `gorm:"column:statik_ide_kekerasan" json:"statik_ide_kekerasan"`
	StatikSkorideKekerasan             *int       `gorm:"column:statik_skoride_kekerasan" json:"statik_skoride_kekerasan"`
	StatikKekerasanAnakAnak            *string    `gorm:"column:statik_kekerasan_anak_anak" json:"statik_kekerasan_anak_anak"`
	StatikSkorkekerasanAnakAnak        *int       `gorm:"column:statik_skorkekerasan_anak_anak" json:"statik_skorkekerasan_anak_anak"`
	StatikPeranDalamHidup              *string    `gorm:"column:statik_peran_dalam_hidup" json:"statik_peran_dalam_hidup"`
	StatikSkorperanDalamHidup          *int       `gorm:"column:statik_skorperan_dalam_hidup" json:"statik_skorperan_dalam_hidup"`
	StatikPenggunaanNapza              *string    `gorm:"column:statik_penggunaan_napza" json:"statik_penggunaan_napza"`
	StatikSkorpenggunaanNapza          *int       `gorm:"column:statik_skorpenggunaan_napza" json:"statik_skorpenggunaan_napza"`
	StatikSkortotal                    *int       `gorm:"column:statik_skortotal" json:"statik_skortotal"`
	DinamisIdeMelukaiOrangLain         *string    `gorm:"column:dinamis_ide_melukai_orang_lain" json:"dinamis_ide_melukai_orang_lain"`
	DinamisSkorideMelukaiOrangLain     *int       `gorm:"column:dinamis_skoride_melukai_orang_lain" json:"dinamis_skoride_melukai_orang_lain"`
	DinamisAksesKekerasan              *string    `gorm:"column:dinamis_akses_kekerasan" json:"dinamis_akses_kekerasan"`
	DinamisSkoraksesKekerasan          *int       `gorm:"column:dinamis_skorakses_kekerasan" json:"dinamis_skorakses_kekerasan"`
	DinamisIdeParanoid                 *string    `gorm:"column:dinamis_ide_paranoid" json:"dinamis_ide_paranoid"`
	DinamisSkorideParanoid             *int       `gorm:"column:dinamis_skoride_paranoid" json:"dinamis_skoride_paranoid"`
	DinamisPerintahHalusinasi          *string    `gorm:"column:dinamis_perintah_halusinasi" json:"dinamis_perintah_halusinasi"`
	DinamisSkorperintahHalusinasi      *int       `gorm:"column:dinamis_skorperintah_halusinasi" json:"dinamis_skorperintah_halusinasi"`
	DinamisFrustasiAgitasi             *string    `gorm:"column:dinamis_frustasi_agitasi" json:"dinamis_frustasi_agitasi"`
	DinamisSkorfrustasiAgitasi         *int       `gorm:"column:dinamis_skorfrustasi_agitasi" json:"dinamis_skorfrustasi_agitasi"`
	DinamisKesenanganKekerasan         *string    `gorm:"column:dinamis_kesenangan_kekerasan" json:"dinamis_kesenangan_kekerasan"`
	DinamisSkorkesenanganKekerasan     *int       `gorm:"column:dinamis_skorkesenangan_kekerasan" json:"dinamis_skorkesenangan_kekerasan"`
	DinamisSeksualTidakWajar           *string    `gorm:"column:dinamis_seksual_tidak_wajar" json:"dinamis_seksual_tidak_wajar"`
	DinamisSkorseksualTidakWajar       *int       `gorm:"column:dinamis_skorseksual_tidak_wajar" json:"dinamis_skorseksual_tidak_wajar"`
	DinamisHilangnyaKontrolDiri        *string    `gorm:"column:dinamis_hilangnya_kontrol_diri" json:"dinamis_hilangnya_kontrol_diri"`
	DinamisSkorhilangnyaKontrolDiri    *int       `gorm:"column:dinamis_skorhilangnya_kontrol_diri" json:"dinamis_skorhilangnya_kontrol_diri"`
	DinamisPengguaanNapza              *string    `gorm:"column:dinamis_pengguaan_napza" json:"dinamis_pengguaan_napza"`
	DinamisSkorpengguaanNapza          *int       `gorm:"column:dinamis_skorpengguaan_napza" json:"dinamis_skorpengguaan_napza"`
	DinamisSkortotal                   *int       `gorm:"column:dinamis_skortotal" json:"dinamis_skortotal"`
	FaktorFaktorPencegahan             *string    `gorm:"column:faktor_faktor_pencegahan" json:"faktor_faktor_pencegahan"`
	TotalSkor                          *int       `gorm:"column:total_skor" json:"total_skor"`
	LevelSkor                          *string    `gorm:"column:level_skor" json:"level_skor"`
}

func (PenilaianTambahanPerilakuKekerasan) TableName() string {
	return "penilaian_tambahan_perilaku_kekerasan"
}
