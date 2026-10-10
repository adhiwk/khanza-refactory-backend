package rekammedis

import "time"

// PenilaianAwalKeperawatanIgd tabel `penilaian_awal_keperawatan_igd` (penilaian awal keperawatan IGD, RMPenilaianAwalKeperawatanIGD).
type PenilaianAwalKeperawatanIgd struct {
	NoRawat          string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal          *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Informasi        string     `gorm:"column:informasi" json:"informasi"`
	KeluhanUtama     string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rpd              string     `gorm:"column:rpd" json:"rpd"`
	Rpo              string     `gorm:"column:rpo" json:"rpo"`
	StatusKehamilan  string     `gorm:"column:status_kehamilan" json:"status_kehamilan"`
	Gravida          *string    `gorm:"column:gravida" json:"gravida"`
	Para             *string    `gorm:"column:para" json:"para"`
	Abortus          *string    `gorm:"column:abortus" json:"abortus"`
	Hpht             *string    `gorm:"column:hpht" json:"hpht"`
	Tekanan          string     `gorm:"column:tekanan" json:"tekanan"`
	Pupil            string     `gorm:"column:pupil" json:"pupil"`
	Neurosensorik    string     `gorm:"column:neurosensorik" json:"neurosensorik"`
	Integumen        string     `gorm:"column:integumen" json:"integumen"`
	Turgor           string     `gorm:"column:turgor" json:"turgor"`
	Edema            string     `gorm:"column:edema" json:"edema"`
	Mukosa           string     `gorm:"column:mukosa" json:"mukosa"`
	Perdarahan       string     `gorm:"column:perdarahan" json:"perdarahan"`
	JumlahPerdarahan *string    `gorm:"column:jumlah_perdarahan" json:"jumlah_perdarahan"`
	WarnaPerdarahan  *string    `gorm:"column:warna_perdarahan" json:"warna_perdarahan"`
	Intoksikasi      string     `gorm:"column:intoksikasi" json:"intoksikasi"`
	Bab              *string    `gorm:"column:bab" json:"bab"`
	Xbab             *string    `gorm:"column:xbab" json:"xbab"`
	Kbab             *string    `gorm:"column:kbab" json:"kbab"`
	Wbab             *string    `gorm:"column:wbab" json:"wbab"`
	Bak              *string    `gorm:"column:bak" json:"bak"`
	Xbak             *string    `gorm:"column:xbak" json:"xbak"`
	Wbak             *string    `gorm:"column:wbak" json:"wbak"`
	Lbak             *string    `gorm:"column:lbak" json:"lbak"`
	Psikologis       string     `gorm:"column:psikologis" json:"psikologis"`
	Jiwa             string     `gorm:"column:jiwa" json:"jiwa"`
	Perilaku         string     `gorm:"column:perilaku" json:"perilaku"`
	Dilaporkan       *string    `gorm:"column:dilaporkan" json:"dilaporkan"`
	Sebutkan         *string    `gorm:"column:sebutkan" json:"sebutkan"`
	Hubungan         string     `gorm:"column:hubungan" json:"hubungan"`
	TinggalDengan    string     `gorm:"column:tinggal_dengan" json:"tinggal_dengan"`
	KetTinggal       *string    `gorm:"column:ket_tinggal" json:"ket_tinggal"`
	Budaya           string     `gorm:"column:budaya" json:"budaya"`
	KetBudaya        string     `gorm:"column:ket_budaya" json:"ket_budaya"`
	PendidikanPj     string     `gorm:"column:pendidikan_pj" json:"pendidikan_pj"`
	KetPendidikanPj  *string    `gorm:"column:ket_pendidikan_pj" json:"ket_pendidikan_pj"`
	Edukasi          string     `gorm:"column:edukasi" json:"edukasi"`
	KetEdukasi       string     `gorm:"column:ket_edukasi" json:"ket_edukasi"`
	Kemampuan        string     `gorm:"column:kemampuan" json:"kemampuan"`
	Aktifitas        string     `gorm:"column:aktifitas" json:"aktifitas"`
	AlatBantu        string     `gorm:"column:alat_bantu" json:"alat_bantu"`
	KetBantu         *string    `gorm:"column:ket_bantu" json:"ket_bantu"`
	Nyeri            string     `gorm:"column:nyeri" json:"nyeri"`
	Provokes         string     `gorm:"column:provokes" json:"provokes"`
	KetProvokes      string     `gorm:"column:ket_provokes" json:"ket_provokes"`
	Quality          string     `gorm:"column:quality" json:"quality"`
	KetQuality       string     `gorm:"column:ket_quality" json:"ket_quality"`
	Lokasi           string     `gorm:"column:lokasi" json:"lokasi"`
	Menyebar         string     `gorm:"column:menyebar" json:"menyebar"`
	SkalaNyeri       string     `gorm:"column:skala_nyeri" json:"skala_nyeri"`
	Durasi           string     `gorm:"column:durasi" json:"durasi"`
	NyeriHilang      string     `gorm:"column:nyeri_hilang" json:"nyeri_hilang"`
	KetNyeri         *string    `gorm:"column:ket_nyeri" json:"ket_nyeri"`
	PadaDokter       string     `gorm:"column:pada_dokter" json:"pada_dokter"`
	KetDokter        *string    `gorm:"column:ket_dokter" json:"ket_dokter"`
	BerjalanA        string     `gorm:"column:berjalan_a" json:"berjalan_a"`
	BerjalanB        string     `gorm:"column:berjalan_b" json:"berjalan_b"`
	BerjalanC        string     `gorm:"column:berjalan_c" json:"berjalan_c"`
	Hasil            string     `gorm:"column:hasil" json:"hasil"`
	Lapor            string     `gorm:"column:lapor" json:"lapor"`
	KetLapor         *string    `gorm:"column:ket_lapor" json:"ket_lapor"`
	Rencana          string     `gorm:"column:rencana" json:"rencana"`
	Nip              string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianAwalKeperawatanIgd) TableName() string {
	return "penilaian_awal_keperawatan_igd"
}
