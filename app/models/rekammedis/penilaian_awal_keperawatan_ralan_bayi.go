package rekammedis

import "time"

// PenilaianAwalKeperawatanRalanBayi tabel `penilaian_awal_keperawatan_ralan_bayi` (penilaian awal keperawatan bayi anak, RMPenilaianAwalKeperawatanBayiAnak).
type PenilaianAwalKeperawatanRalanBayi struct {
	NoRawat           string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal           *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Informasi         string     `gorm:"column:informasi" json:"informasi"`
	Td                string     `gorm:"column:td" json:"td"`
	Nadi              string     `gorm:"column:nadi" json:"nadi"`
	Rr                string     `gorm:"column:rr" json:"rr"`
	Suhu              string     `gorm:"column:suhu" json:"suhu"`
	Gcs               string     `gorm:"column:gcs" json:"gcs"`
	Bb                string     `gorm:"column:bb" json:"bb"`
	Tb                string     `gorm:"column:tb" json:"tb"`
	Lp                string     `gorm:"column:lp" json:"lp"`
	Lk                string     `gorm:"column:lk" json:"lk"`
	Ld                string     `gorm:"column:ld" json:"ld"`
	KeluhanUtama      string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rpd               string     `gorm:"column:rpd" json:"rpd"`
	Rpk               string     `gorm:"column:rpk" json:"rpk"`
	Rpo               string     `gorm:"column:rpo" json:"rpo"`
	Alergi            string     `gorm:"column:alergi" json:"alergi"`
	Anakke            string     `gorm:"column:anakke" json:"anakke"`
	Darisaudara       string     `gorm:"column:darisaudara" json:"darisaudara"`
	Caralahir         string     `gorm:"column:caralahir" json:"caralahir"`
	KetCaralahir      string     `gorm:"column:ket_caralahir" json:"ket_caralahir"`
	Umurkelahiran     string     `gorm:"column:umurkelahiran" json:"umurkelahiran"`
	Kelainanbawaan    string     `gorm:"column:kelainanbawaan" json:"kelainanbawaan"`
	KetKelainanBawaan string     `gorm:"column:ket_kelainan_bawaan" json:"ket_kelainan_bawaan"`
	Usiatengkurap     string     `gorm:"column:usiatengkurap" json:"usiatengkurap"`
	Usiaduduk         string     `gorm:"column:usiaduduk" json:"usiaduduk"`
	Usiaberdiri       string     `gorm:"column:usiaberdiri" json:"usiaberdiri"`
	Usiagigipertama   string     `gorm:"column:usiagigipertama" json:"usiagigipertama"`
	Usiaberjalan      string     `gorm:"column:usiaberjalan" json:"usiaberjalan"`
	Usiabicara        string     `gorm:"column:usiabicara" json:"usiabicara"`
	Usiamembaca       string     `gorm:"column:usiamembaca" json:"usiamembaca"`
	Usiamenulis       string     `gorm:"column:usiamenulis" json:"usiamenulis"`
	Gangguanemosi     string     `gorm:"column:gangguanemosi" json:"gangguanemosi"`
	AlatBantu         string     `gorm:"column:alat_bantu" json:"alat_bantu"`
	KetBantu          string     `gorm:"column:ket_bantu" json:"ket_bantu"`
	Prothesa          string     `gorm:"column:prothesa" json:"prothesa"`
	KetPro            string     `gorm:"column:ket_pro" json:"ket_pro"`
	Adl               string     `gorm:"column:adl" json:"adl"`
	StatusPsiko       string     `gorm:"column:status_psiko" json:"status_psiko"`
	KetPsiko          string     `gorm:"column:ket_psiko" json:"ket_psiko"`
	HubKeluarga       string     `gorm:"column:hub_keluarga" json:"hub_keluarga"`
	Pengasuh          string     `gorm:"column:pengasuh" json:"pengasuh"`
	KetPengasuh       string     `gorm:"column:ket_pengasuh" json:"ket_pengasuh"`
	Ekonomi           string     `gorm:"column:ekonomi" json:"ekonomi"`
	Budaya            string     `gorm:"column:budaya" json:"budaya"`
	KetBudaya         string     `gorm:"column:ket_budaya" json:"ket_budaya"`
	Edukasi           string     `gorm:"column:edukasi" json:"edukasi"`
	KetEdukasi        string     `gorm:"column:ket_edukasi" json:"ket_edukasi"`
	BerjalanA         string     `gorm:"column:berjalan_a" json:"berjalan_a"`
	BerjalanB         string     `gorm:"column:berjalan_b" json:"berjalan_b"`
	BerjalanC         string     `gorm:"column:berjalan_c" json:"berjalan_c"`
	Hasil             string     `gorm:"column:hasil" json:"hasil"`
	Lapor             string     `gorm:"column:lapor" json:"lapor"`
	KetLapor          string     `gorm:"column:ket_lapor" json:"ket_lapor"`
	Sg1               string     `gorm:"column:sg1" json:"sg1"`
	Nilai1            string     `gorm:"column:nilai1" json:"nilai1"`
	Sg2               string     `gorm:"column:sg2" json:"sg2"`
	Nilai2            string     `gorm:"column:nilai2" json:"nilai2"`
	Sg3               string     `gorm:"column:sg3" json:"sg3"`
	Nilai3            string     `gorm:"column:nilai3" json:"nilai3"`
	Sg4               string     `gorm:"column:sg4" json:"sg4"`
	Nilai4            string     `gorm:"column:nilai4" json:"nilai4"`
	TotalHasil        int        `gorm:"column:total_hasil" json:"total_hasil"`
	Wajah             string     `gorm:"column:wajah" json:"wajah"`
	Nilaiwajah        string     `gorm:"column:nilaiwajah" json:"nilaiwajah"`
	Kaki              string     `gorm:"column:kaki" json:"kaki"`
	Nilaikaki         string     `gorm:"column:nilaikaki" json:"nilaikaki"`
	Aktifitas         string     `gorm:"column:aktifitas" json:"aktifitas"`
	Nilaiaktifitas    string     `gorm:"column:nilaiaktifitas" json:"nilaiaktifitas"`
	Menangis          string     `gorm:"column:menangis" json:"menangis"`
	Nilaimenangis     string     `gorm:"column:nilaimenangis" json:"nilaimenangis"`
	Bersuara          string     `gorm:"column:bersuara" json:"bersuara"`
	Nilaibersuara     string     `gorm:"column:nilaibersuara" json:"nilaibersuara"`
	Hasilnyeri        int        `gorm:"column:hasilnyeri" json:"hasilnyeri"`
	Nyeri             string     `gorm:"column:nyeri" json:"nyeri"`
	Lokasi            string     `gorm:"column:lokasi" json:"lokasi"`
	Durasi            string     `gorm:"column:durasi" json:"durasi"`
	Frekuensi         string     `gorm:"column:frekuensi" json:"frekuensi"`
	NyeriHilang       string     `gorm:"column:nyeri_hilang" json:"nyeri_hilang"`
	KetNyeri          string     `gorm:"column:ket_nyeri" json:"ket_nyeri"`
	PadaDokter        string     `gorm:"column:pada_dokter" json:"pada_dokter"`
	KetDokter         string     `gorm:"column:ket_dokter" json:"ket_dokter"`
	Rencana           string     `gorm:"column:rencana" json:"rencana"`
	Nip               string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianAwalKeperawatanRalanBayi) TableName() string {
	return "penilaian_awal_keperawatan_ralan_bayi"
}
