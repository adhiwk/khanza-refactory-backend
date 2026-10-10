package rekammedis

import "time"

// SkriningRisikoKankerParu tabel `skrining_risiko_kanker_paru` (skrining risiko kanker paru, RMSkriningRisikoKankerParu).
type SkriningRisikoKankerParu struct {
	NoRawat                                 string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                                 *time.Time `gorm:"column:tanggal" json:"tanggal"`
	JenisKelamin                            *string    `gorm:"column:jenis_kelamin" json:"jenis_kelamin"`
	NilaiJenisKelamin                       *string    `gorm:"column:nilai_jenis_kelamin" json:"nilai_jenis_kelamin"`
	Umur                                    *string    `gorm:"column:umur" json:"umur"`
	NilaiUmur                               *string    `gorm:"column:nilai_umur" json:"nilai_umur"`
	PernahKanker                            *string    `gorm:"column:pernah_kanker" json:"pernah_kanker"`
	NilaiPernahKanker                       *string    `gorm:"column:nilai_pernah_kanker" json:"nilai_pernah_kanker"`
	AdaKeluargaKanker                       *string    `gorm:"column:ada_keluarga_kanker" json:"ada_keluarga_kanker"`
	NilaiAdaKeluargaKanker                  *string    `gorm:"column:nilai_ada_keluarga_kanker" json:"nilai_ada_keluarga_kanker"`
	RiwayatRokok                            *string    `gorm:"column:riwayat_rokok" json:"riwayat_rokok"`
	NilaiRiwayatRokok                       *string    `gorm:"column:nilai_riwayat_rokok" json:"nilai_riwayat_rokok"`
	RiwayatBekerjaMengandungKarsinogen      *string    `gorm:"column:riwayat_bekerja_mengandung_karsinogen" json:"riwayat_bekerja_mengandung_karsinogen"`
	NilaiRiwayatBekerjaMengandungKarsinogen *string    `gorm:"column:nilai_riwayat_bekerja_mengandung_karsinogen" json:"nilai_riwayat_bekerja_mengandung_karsinogen"`
	LingkunganTinggalPolusiTinggi           *string    `gorm:"column:lingkungan_tinggal_polusi_tinggi" json:"lingkungan_tinggal_polusi_tinggi"`
	NilaiLingkunganTinggalPolusiTinggi      *string    `gorm:"column:nilai_lingkungan_tinggal_polusi_tinggi" json:"nilai_lingkungan_tinggal_polusi_tinggi"`
	LingkunganRumahTidakSehat               *string    `gorm:"column:lingkungan_rumah_tidak_sehat" json:"lingkungan_rumah_tidak_sehat"`
	NilaiLingkunganRumahTidakSehat          *string    `gorm:"column:nilai_lingkungan_rumah_tidak_sehat" json:"nilai_lingkungan_rumah_tidak_sehat"`
	PernahParuKronik                        *string    `gorm:"column:pernah_paru_kronik" json:"pernah_paru_kronik"`
	NilaiPernahParuKronik                   *string    `gorm:"column:nilai_pernah_paru_kronik" json:"nilai_pernah_paru_kronik"`
	TotalSkor                               *string    `gorm:"column:total_skor" json:"total_skor"`
	HasilSkrining                           *string    `gorm:"column:hasil_skrining" json:"hasil_skrining"`
	Keterangan                              string     `gorm:"column:keterangan" json:"keterangan"`
	Nip                                     string     `gorm:"column:nip" json:"nip"`
}

func (SkriningRisikoKankerParu) TableName() string {
	return "skrining_risiko_kanker_paru"
}
