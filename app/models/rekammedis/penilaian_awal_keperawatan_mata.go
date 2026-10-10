package rekammedis

import "time"

// PenilaianAwalKeperawatanMata tabel `penilaian_awal_keperawatan_mata` (penilaian awal keperawatan mata, RMPenilaianAwalKeperawatanMata).
type PenilaianAwalKeperawatanMata struct {
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
	Bmi               string     `gorm:"column:bmi" json:"bmi"`
	KeluhanUtama      string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rpd               string     `gorm:"column:rpd" json:"rpd"`
	Rps               string     `gorm:"column:rps" json:"rps"`
	Rpk               string     `gorm:"column:rpk" json:"rpk"`
	Rpo               string     `gorm:"column:rpo" json:"rpo"`
	Alergi            string     `gorm:"column:alergi" json:"alergi"`
	AlatBantu         string     `gorm:"column:alat_bantu" json:"alat_bantu"`
	KetBantu          string     `gorm:"column:ket_bantu" json:"ket_bantu"`
	Prothesa          string     `gorm:"column:prothesa" json:"prothesa"`
	KetPro            string     `gorm:"column:ket_pro" json:"ket_pro"`
	Adl               string     `gorm:"column:adl" json:"adl"`
	StatusPsiko       string     `gorm:"column:status_psiko" json:"status_psiko"`
	KetPsiko          string     `gorm:"column:ket_psiko" json:"ket_psiko"`
	HubKeluarga       string     `gorm:"column:hub_keluarga" json:"hub_keluarga"`
	TinggalDengan     string     `gorm:"column:tinggal_dengan" json:"tinggal_dengan"`
	KetTinggal        string     `gorm:"column:ket_tinggal" json:"ket_tinggal"`
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
	Nyeri             string     `gorm:"column:nyeri" json:"nyeri"`
	Provokes          string     `gorm:"column:provokes" json:"provokes"`
	KetProvokes       string     `gorm:"column:ket_provokes" json:"ket_provokes"`
	Quality           string     `gorm:"column:quality" json:"quality"`
	KetQuality        string     `gorm:"column:ket_quality" json:"ket_quality"`
	Lokasi            string     `gorm:"column:lokasi" json:"lokasi"`
	Menyebar          string     `gorm:"column:menyebar" json:"menyebar"`
	SkalaNyeri        string     `gorm:"column:skala_nyeri" json:"skala_nyeri"`
	Durasi            string     `gorm:"column:durasi" json:"durasi"`
	NyeriHilang       string     `gorm:"column:nyeri_hilang" json:"nyeri_hilang"`
	KetNyeri          string     `gorm:"column:ket_nyeri" json:"ket_nyeri"`
	PadaDokter        string     `gorm:"column:pada_dokter" json:"pada_dokter"`
	KetDokter         string     `gorm:"column:ket_dokter" json:"ket_dokter"`
	Visuskanan        string     `gorm:"column:visuskanan" json:"visuskanan"`
	Visuskiri         string     `gorm:"column:visuskiri" json:"visuskiri"`
	Refraksikanan     string     `gorm:"column:refraksikanan" json:"refraksikanan"`
	Refraksikiri      string     `gorm:"column:refraksikiri" json:"refraksikiri"`
	Tiokanan          string     `gorm:"column:tiokanan" json:"tiokanan"`
	Tiokiri           string     `gorm:"column:tiokiri" json:"tiokiri"`
	Palberakanan      string     `gorm:"column:palberakanan" json:"palberakanan"`
	Palberakiri       string     `gorm:"column:palberakiri" json:"palberakiri"`
	Konjungtivakanan  string     `gorm:"column:konjungtivakanan" json:"konjungtivakanan"`
	Konjungtivakiri   string     `gorm:"column:konjungtivakiri" json:"konjungtivakiri"`
	Sklerakanan       string     `gorm:"column:sklerakanan" json:"sklerakanan"`
	Sklerakiri        string     `gorm:"column:sklerakiri" json:"sklerakiri"`
	Korneakanan       string     `gorm:"column:korneakanan" json:"korneakanan"`
	Korneakiri        string     `gorm:"column:korneakiri" json:"korneakiri"`
	Bmdkanan          string     `gorm:"column:bmdkanan" json:"bmdkanan"`
	Bmdkiri           string     `gorm:"column:bmdkiri" json:"bmdkiri"`
	Iriskanan         string     `gorm:"column:iriskanan" json:"iriskanan"`
	Iriskiri          string     `gorm:"column:iriskiri" json:"iriskiri"`
	Pupilkanan        string     `gorm:"column:pupilkanan" json:"pupilkanan"`
	Pupilkiri         string     `gorm:"column:pupilkiri" json:"pupilkiri"`
	Lensakanan        string     `gorm:"column:lensakanan" json:"lensakanan"`
	Lensakiri         string     `gorm:"column:lensakiri" json:"lensakiri"`
	Oftalmoskopikanan string     `gorm:"column:oftalmoskopikanan" json:"oftalmoskopikanan"`
	Oftalmoskopikiri  string     `gorm:"column:oftalmoskopikiri" json:"oftalmoskopikiri"`
	Rencana           string     `gorm:"column:rencana" json:"rencana"`
	Nip               string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianAwalKeperawatanMata) TableName() string {
	return "penilaian_awal_keperawatan_mata"
}
