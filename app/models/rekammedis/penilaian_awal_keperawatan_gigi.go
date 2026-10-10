package rekammedis

import "time"

// PenilaianAwalKeperawatanGigi tabel `penilaian_awal_keperawatan_gigi` (penilaian awal keperawatan gigi, RMPenilaianAwalKeperawatanGigi).
type PenilaianAwalKeperawatanGigi struct {
	NoRawat                 string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                 *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Informasi               string     `gorm:"column:informasi" json:"informasi"`
	Td                      string     `gorm:"column:td" json:"td"`
	Nadi                    string     `gorm:"column:nadi" json:"nadi"`
	Rr                      string     `gorm:"column:rr" json:"rr"`
	Suhu                    string     `gorm:"column:suhu" json:"suhu"`
	Bb                      string     `gorm:"column:bb" json:"bb"`
	Tb                      string     `gorm:"column:tb" json:"tb"`
	Bmi                     string     `gorm:"column:bmi" json:"bmi"`
	KeluhanUtama            string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	RiwayatPenyakit         *string    `gorm:"column:riwayat_penyakit" json:"riwayat_penyakit"`
	KetRiwayatPenyakit      string     `gorm:"column:ket_riwayat_penyakit" json:"ket_riwayat_penyakit"`
	Alergi                  string     `gorm:"column:alergi" json:"alergi"`
	RiwayatPerawatanGigi    string     `gorm:"column:riwayat_perawatan_gigi" json:"riwayat_perawatan_gigi"`
	KetRiwayatPerawatanGigi string     `gorm:"column:ket_riwayat_perawatan_gigi" json:"ket_riwayat_perawatan_gigi"`
	KebiasaanSikatGigi      string     `gorm:"column:kebiasaan_sikat_gigi" json:"kebiasaan_sikat_gigi"`
	KebiasaanLain           *string    `gorm:"column:kebiasaan_lain" json:"kebiasaan_lain"`
	KetKebiasaanLain        string     `gorm:"column:ket_kebiasaan_lain" json:"ket_kebiasaan_lain"`
	ObatYangDiminumSaatini  *string    `gorm:"column:obat_yang_diminum_saatini" json:"obat_yang_diminum_saatini"`
	AlatBantu               string     `gorm:"column:alat_bantu" json:"alat_bantu"`
	KetAlatBantu            string     `gorm:"column:ket_alat_bantu" json:"ket_alat_bantu"`
	Prothesa                string     `gorm:"column:prothesa" json:"prothesa"`
	KetPro                  string     `gorm:"column:ket_pro" json:"ket_pro"`
	StatusPsiko             string     `gorm:"column:status_psiko" json:"status_psiko"`
	KetPsiko                string     `gorm:"column:ket_psiko" json:"ket_psiko"`
	HubKeluarga             string     `gorm:"column:hub_keluarga" json:"hub_keluarga"`
	TinggalDengan           string     `gorm:"column:tinggal_dengan" json:"tinggal_dengan"`
	KetTinggal              string     `gorm:"column:ket_tinggal" json:"ket_tinggal"`
	Ekonomi                 string     `gorm:"column:ekonomi" json:"ekonomi"`
	Budaya                  string     `gorm:"column:budaya" json:"budaya"`
	KetBudaya               string     `gorm:"column:ket_budaya" json:"ket_budaya"`
	Edukasi                 string     `gorm:"column:edukasi" json:"edukasi"`
	KetEdukasi              string     `gorm:"column:ket_edukasi" json:"ket_edukasi"`
	BerjalanA               string     `gorm:"column:berjalan_a" json:"berjalan_a"`
	BerjalanB               string     `gorm:"column:berjalan_b" json:"berjalan_b"`
	BerjalanC               string     `gorm:"column:berjalan_c" json:"berjalan_c"`
	Hasil                   string     `gorm:"column:hasil" json:"hasil"`
	Lapor                   string     `gorm:"column:lapor" json:"lapor"`
	KetLapor                string     `gorm:"column:ket_lapor" json:"ket_lapor"`
	Nyeri                   string     `gorm:"column:nyeri" json:"nyeri"`
	Lokasi                  string     `gorm:"column:lokasi" json:"lokasi"`
	SkalaNyeri              string     `gorm:"column:skala_nyeri" json:"skala_nyeri"`
	Durasi                  string     `gorm:"column:durasi" json:"durasi"`
	Frekuensi               string     `gorm:"column:frekuensi" json:"frekuensi"`
	NyeriHilang             string     `gorm:"column:nyeri_hilang" json:"nyeri_hilang"`
	KetNyeri                string     `gorm:"column:ket_nyeri" json:"ket_nyeri"`
	PadaDokter              string     `gorm:"column:pada_dokter" json:"pada_dokter"`
	KetDokter               string     `gorm:"column:ket_dokter" json:"ket_dokter"`
	KebersihanMulut         string     `gorm:"column:kebersihan_mulut" json:"kebersihan_mulut"`
	MukosaMulut             string     `gorm:"column:mukosa_mulut" json:"mukosa_mulut"`
	Karies                  string     `gorm:"column:karies" json:"karies"`
	KarangGigi              string     `gorm:"column:karang_gigi" json:"karang_gigi"`
	Gingiva                 string     `gorm:"column:gingiva" json:"gingiva"`
	Palatum                 string     `gorm:"column:palatum" json:"palatum"`
	Rencana                 string     `gorm:"column:rencana" json:"rencana"`
	Nip                     string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianAwalKeperawatanGigi) TableName() string {
	return "penilaian_awal_keperawatan_gigi"
}
