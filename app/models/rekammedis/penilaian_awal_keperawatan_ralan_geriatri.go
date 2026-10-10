package rekammedis

import "time"

// PenilaianAwalKeperawatanRalanGeriatri tabel `penilaian_awal_keperawatan_ralan_geriatri` (penilaian awal keperawatan ralan geriatri, RMPenilaianAwalKeperawatanRalanGeriatri).
type PenilaianAwalKeperawatanRalanGeriatri struct {
	NoRawat                              string     `gorm:"column:no_rawat;primaryKey;autoIncrement:false" json:"no_rawat"`
	Tanggal                              *time.Time `gorm:"column:tanggal" json:"tanggal"`
	Informasi                            string     `gorm:"column:informasi" json:"informasi"`
	Td                                   string     `gorm:"column:td" json:"td"`
	Nadi                                 string     `gorm:"column:nadi" json:"nadi"`
	Rr                                   string     `gorm:"column:rr" json:"rr"`
	Suhu                                 string     `gorm:"column:suhu" json:"suhu"`
	Gcs                                  string     `gorm:"column:gcs" json:"gcs"`
	Bb                                   string     `gorm:"column:bb" json:"bb"`
	Tb                                   string     `gorm:"column:tb" json:"tb"`
	Bmi                                  string     `gorm:"column:bmi" json:"bmi"`
	KeluhanUtama                         string     `gorm:"column:keluhan_utama" json:"keluhan_utama"`
	Rpd                                  string     `gorm:"column:rpd" json:"rpd"`
	Rpk                                  string     `gorm:"column:rpk" json:"rpk"`
	Rpo                                  string     `gorm:"column:rpo" json:"rpo"`
	Alergi                               string     `gorm:"column:alergi" json:"alergi"`
	AlatBantu                            string     `gorm:"column:alat_bantu" json:"alat_bantu"`
	KetBantu                             string     `gorm:"column:ket_bantu" json:"ket_bantu"`
	Prothesa                             string     `gorm:"column:prothesa" json:"prothesa"`
	KetPro                               string     `gorm:"column:ket_pro" json:"ket_pro"`
	Adl                                  string     `gorm:"column:adl" json:"adl"`
	StatusPsiko                          string     `gorm:"column:status_psiko" json:"status_psiko"`
	KetPsiko                             string     `gorm:"column:ket_psiko" json:"ket_psiko"`
	HubKeluarga                          string     `gorm:"column:hub_keluarga" json:"hub_keluarga"`
	TinggalDengan                        string     `gorm:"column:tinggal_dengan" json:"tinggal_dengan"`
	KetTinggal                           string     `gorm:"column:ket_tinggal" json:"ket_tinggal"`
	Ekonomi                              string     `gorm:"column:ekonomi" json:"ekonomi"`
	Budaya                               string     `gorm:"column:budaya" json:"budaya"`
	KetBudaya                            string     `gorm:"column:ket_budaya" json:"ket_budaya"`
	Edukasi                              string     `gorm:"column:edukasi" json:"edukasi"`
	KetEdukasi                           string     `gorm:"column:ket_edukasi" json:"ket_edukasi"`
	BerjalanA                            string     `gorm:"column:berjalan_a" json:"berjalan_a"`
	BerjalanB                            string     `gorm:"column:berjalan_b" json:"berjalan_b"`
	BerjalanC                            string     `gorm:"column:berjalan_c" json:"berjalan_c"`
	Hasil                                string     `gorm:"column:hasil" json:"hasil"`
	Lapor                                string     `gorm:"column:lapor" json:"lapor"`
	KetLapor                             string     `gorm:"column:ket_lapor" json:"ket_lapor"`
	Sg1                                  string     `gorm:"column:sg1" json:"sg1"`
	Nilai1                               string     `gorm:"column:nilai1" json:"nilai1"`
	Sg2                                  string     `gorm:"column:sg2" json:"sg2"`
	Nilai2                               string     `gorm:"column:nilai2" json:"nilai2"`
	TotalHasil                           int        `gorm:"column:total_hasil" json:"total_hasil"`
	Nyeri                                string     `gorm:"column:nyeri" json:"nyeri"`
	Provokes                             string     `gorm:"column:provokes" json:"provokes"`
	KetProvokes                          string     `gorm:"column:ket_provokes" json:"ket_provokes"`
	Quality                              string     `gorm:"column:quality" json:"quality"`
	KetQuality                           string     `gorm:"column:ket_quality" json:"ket_quality"`
	Lokasi                               string     `gorm:"column:lokasi" json:"lokasi"`
	Menyebar                             string     `gorm:"column:menyebar" json:"menyebar"`
	SkalaNyeri                           string     `gorm:"column:skala_nyeri" json:"skala_nyeri"`
	Durasi                               string     `gorm:"column:durasi" json:"durasi"`
	NyeriHilang                          string     `gorm:"column:nyeri_hilang" json:"nyeri_hilang"`
	KetNyeri                             string     `gorm:"column:ket_nyeri" json:"ket_nyeri"`
	PadaDokter                           string     `gorm:"column:pada_dokter" json:"pada_dokter"`
	KetDokter                            string     `gorm:"column:ket_dokter" json:"ket_dokter"`
	EdukasiKemampuanBacatulis            *string    `gorm:"column:edukasi_kemampuan_bacatulis" json:"edukasi_kemampuan_bacatulis"`
	EdukasiKebutuhanPenerjemah           *string    `gorm:"column:edukasi_kebutuhan_penerjemah" json:"edukasi_kebutuhan_penerjemah"`
	EdukasiKeteranganKebutuhanPenerjemah *string    `gorm:"column:edukasi_keterangan_kebutuhan_penerjemah" json:"edukasi_keterangan_kebutuhan_penerjemah"`
	EdukasiHambatan                      *string    `gorm:"column:edukasi_hambatan" json:"edukasi_hambatan"`
	EdukasiHambatanKategori              string     `gorm:"column:edukasi_hambatan_kategori" json:"edukasi_hambatan_kategori"`
	EdukasiKeteranganHambatan            *string    `gorm:"column:edukasi_keterangan_hambatan" json:"edukasi_keterangan_hambatan"`
	EdukasiCaraBicara                    *string    `gorm:"column:edukasi_cara_bicara" json:"edukasi_cara_bicara"`
	EdukasiBahasaIsyarat                 *string    `gorm:"column:edukasi_bahasa_isyarat" json:"edukasi_bahasa_isyarat"`
	EdukasiMenerimaInformasi             *string    `gorm:"column:edukasi_menerima_informasi" json:"edukasi_menerima_informasi"`
	EdukasiKeteranganMenerimaInformasi   *string    `gorm:"column:edukasi_keterangan_menerima_informasi" json:"edukasi_keterangan_menerima_informasi"`
	EdukasiMetodeBelajar                 *string    `gorm:"column:edukasi_metode_belajar" json:"edukasi_metode_belajar"`
	FrailyPhenotypeBeratBadan            *string    `gorm:"column:fraily_phenotype_berat_badan" json:"fraily_phenotype_berat_badan"`
	FrailyPhenotypeBeratBadanNilai       *int       `gorm:"column:fraily_phenotype_berat_badan_nilai" json:"fraily_phenotype_berat_badan_nilai"`
	FrailyPhenotypeAktifitasFisik        *string    `gorm:"column:fraily_phenotype_aktifitas_fisik" json:"fraily_phenotype_aktifitas_fisik"`
	FrailyPhenotypeAktifitasFisikNilai   *int       `gorm:"column:fraily_phenotype_aktifitas_fisik_nilai" json:"fraily_phenotype_aktifitas_fisik_nilai"`
	FrailyPhenotypeKelelahan             *string    `gorm:"column:fraily_phenotype_kelelahan" json:"fraily_phenotype_kelelahan"`
	FrailyPhenotypeKelelahanNilai        *int       `gorm:"column:fraily_phenotype_kelelahan_nilai" json:"fraily_phenotype_kelelahan_nilai"`
	FrailyPhenotypeKekuatan              *string    `gorm:"column:fraily_phenotype_kekuatan" json:"fraily_phenotype_kekuatan"`
	FrailyPhenotypeKekuatanNilai         *int       `gorm:"column:fraily_phenotype_kekuatan_nilai" json:"fraily_phenotype_kekuatan_nilai"`
	FrailyPhenotypeWaktuBerjalan         *string    `gorm:"column:fraily_phenotype_waktu_berjalan" json:"fraily_phenotype_waktu_berjalan"`
	FrailyPhenotypeWaktuBerjalanNilai    *int       `gorm:"column:fraily_phenotype_waktu_berjalan_nilai" json:"fraily_phenotype_waktu_berjalan_nilai"`
	FrailyPhenotypeNilaiTotal            *int       `gorm:"column:fraily_phenotype_nilai_total" json:"fraily_phenotype_nilai_total"`
	FrailyPhenotypeStatus                *string    `gorm:"column:fraily_phenotype_status" json:"fraily_phenotype_status"`
	Rencana                              string     `gorm:"column:rencana" json:"rencana"`
	Nip                                  string     `gorm:"column:nip" json:"nip"`
}

func (PenilaianAwalKeperawatanRalanGeriatri) TableName() string {
	return "penilaian_awal_keperawatan_ralan_geriatri"
}
