package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianAwalKeperawatanRalanGeriatriData isian penilaian awal keperawatan ralan geriatri.
type PenilaianAwalKeperawatanRalanGeriatriData struct {
	Tanggal                              string `form:"tanggal" json:"tanggal"`
	Informasi                            string `form:"informasi" json:"informasi"`
	Td                                   string `form:"td" json:"td"`
	Nadi                                 string `form:"nadi" json:"nadi"`
	Rr                                   string `form:"rr" json:"rr"`
	Suhu                                 string `form:"suhu" json:"suhu"`
	Gcs                                  string `form:"gcs" json:"gcs"`
	Bb                                   string `form:"bb" json:"bb"`
	Tb                                   string `form:"tb" json:"tb"`
	Bmi                                  string `form:"bmi" json:"bmi"`
	KeluhanUtama                         string `form:"keluhan_utama" json:"keluhan_utama"`
	Rpd                                  string `form:"rpd" json:"rpd"`
	Rpk                                  string `form:"rpk" json:"rpk"`
	Rpo                                  string `form:"rpo" json:"rpo"`
	Alergi                               string `form:"alergi" json:"alergi"`
	AlatBantu                            string `form:"alat_bantu" json:"alat_bantu"`
	KetBantu                             string `form:"ket_bantu" json:"ket_bantu"`
	Prothesa                             string `form:"prothesa" json:"prothesa"`
	KetPro                               string `form:"ket_pro" json:"ket_pro"`
	Adl                                  string `form:"adl" json:"adl"`
	StatusPsiko                          string `form:"status_psiko" json:"status_psiko"`
	KetPsiko                             string `form:"ket_psiko" json:"ket_psiko"`
	HubKeluarga                          string `form:"hub_keluarga" json:"hub_keluarga"`
	TinggalDengan                        string `form:"tinggal_dengan" json:"tinggal_dengan"`
	KetTinggal                           string `form:"ket_tinggal" json:"ket_tinggal"`
	Ekonomi                              string `form:"ekonomi" json:"ekonomi"`
	Budaya                               string `form:"budaya" json:"budaya"`
	KetBudaya                            string `form:"ket_budaya" json:"ket_budaya"`
	Edukasi                              string `form:"edukasi" json:"edukasi"`
	KetEdukasi                           string `form:"ket_edukasi" json:"ket_edukasi"`
	BerjalanA                            string `form:"berjalan_a" json:"berjalan_a"`
	BerjalanB                            string `form:"berjalan_b" json:"berjalan_b"`
	BerjalanC                            string `form:"berjalan_c" json:"berjalan_c"`
	Hasil                                string `form:"hasil" json:"hasil"`
	Lapor                                string `form:"lapor" json:"lapor"`
	KetLapor                             string `form:"ket_lapor" json:"ket_lapor"`
	Sg1                                  string `form:"sg1" json:"sg1"`
	Nilai1                               string `form:"nilai1" json:"nilai1"`
	Sg2                                  string `form:"sg2" json:"sg2"`
	Nilai2                               string `form:"nilai2" json:"nilai2"`
	TotalHasil                           int    `form:"total_hasil" json:"total_hasil"`
	Nyeri                                string `form:"nyeri" json:"nyeri"`
	Provokes                             string `form:"provokes" json:"provokes"`
	KetProvokes                          string `form:"ket_provokes" json:"ket_provokes"`
	Quality                              string `form:"quality" json:"quality"`
	KetQuality                           string `form:"ket_quality" json:"ket_quality"`
	Lokasi                               string `form:"lokasi" json:"lokasi"`
	Menyebar                             string `form:"menyebar" json:"menyebar"`
	SkalaNyeri                           string `form:"skala_nyeri" json:"skala_nyeri"`
	Durasi                               string `form:"durasi" json:"durasi"`
	NyeriHilang                          string `form:"nyeri_hilang" json:"nyeri_hilang"`
	KetNyeri                             string `form:"ket_nyeri" json:"ket_nyeri"`
	PadaDokter                           string `form:"pada_dokter" json:"pada_dokter"`
	KetDokter                            string `form:"ket_dokter" json:"ket_dokter"`
	EdukasiKemampuanBacatulis            string `form:"edukasi_kemampuan_bacatulis" json:"edukasi_kemampuan_bacatulis"`
	EdukasiKebutuhanPenerjemah           string `form:"edukasi_kebutuhan_penerjemah" json:"edukasi_kebutuhan_penerjemah"`
	EdukasiKeteranganKebutuhanPenerjemah string `form:"edukasi_keterangan_kebutuhan_penerjemah" json:"edukasi_keterangan_kebutuhan_penerjemah"`
	EdukasiHambatan                      string `form:"edukasi_hambatan" json:"edukasi_hambatan"`
	EdukasiHambatanKategori              string `form:"edukasi_hambatan_kategori" json:"edukasi_hambatan_kategori"`
	EdukasiKeteranganHambatan            string `form:"edukasi_keterangan_hambatan" json:"edukasi_keterangan_hambatan"`
	EdukasiCaraBicara                    string `form:"edukasi_cara_bicara" json:"edukasi_cara_bicara"`
	EdukasiBahasaIsyarat                 string `form:"edukasi_bahasa_isyarat" json:"edukasi_bahasa_isyarat"`
	EdukasiMenerimaInformasi             string `form:"edukasi_menerima_informasi" json:"edukasi_menerima_informasi"`
	EdukasiKeteranganMenerimaInformasi   string `form:"edukasi_keterangan_menerima_informasi" json:"edukasi_keterangan_menerima_informasi"`
	EdukasiMetodeBelajar                 string `form:"edukasi_metode_belajar" json:"edukasi_metode_belajar"`
	FrailyPhenotypeBeratBadan            string `form:"fraily_phenotype_berat_badan" json:"fraily_phenotype_berat_badan"`
	FrailyPhenotypeBeratBadanNilai       *int   `form:"fraily_phenotype_berat_badan_nilai" json:"fraily_phenotype_berat_badan_nilai"`
	FrailyPhenotypeAktifitasFisik        string `form:"fraily_phenotype_aktifitas_fisik" json:"fraily_phenotype_aktifitas_fisik"`
	FrailyPhenotypeAktifitasFisikNilai   *int   `form:"fraily_phenotype_aktifitas_fisik_nilai" json:"fraily_phenotype_aktifitas_fisik_nilai"`
	FrailyPhenotypeKelelahan             string `form:"fraily_phenotype_kelelahan" json:"fraily_phenotype_kelelahan"`
	FrailyPhenotypeKelelahanNilai        *int   `form:"fraily_phenotype_kelelahan_nilai" json:"fraily_phenotype_kelelahan_nilai"`
	FrailyPhenotypeKekuatan              string `form:"fraily_phenotype_kekuatan" json:"fraily_phenotype_kekuatan"`
	FrailyPhenotypeKekuatanNilai         *int   `form:"fraily_phenotype_kekuatan_nilai" json:"fraily_phenotype_kekuatan_nilai"`
	FrailyPhenotypeWaktuBerjalan         string `form:"fraily_phenotype_waktu_berjalan" json:"fraily_phenotype_waktu_berjalan"`
	FrailyPhenotypeWaktuBerjalanNilai    *int   `form:"fraily_phenotype_waktu_berjalan_nilai" json:"fraily_phenotype_waktu_berjalan_nilai"`
	FrailyPhenotypeNilaiTotal            *int   `form:"fraily_phenotype_nilai_total" json:"fraily_phenotype_nilai_total"`
	FrailyPhenotypeStatus                string `form:"fraily_phenotype_status" json:"fraily_phenotype_status"`
	Rencana                              string `form:"rencana" json:"rencana"`
	Nip                                  string `form:"nip" json:"nip"`
}

// PenilaianAwalKeperawatanRalanGeriatriDetail daftar kode detail (diganti seluruhnya saat simpan/ubah).
type PenilaianAwalKeperawatanRalanGeriatriDetail struct {
	KodeMasalah []string `form:"kode_masalah" json:"kode_masalah"`
	KodeRencana []string `form:"kode_rencana" json:"kode_rencana"`
}

func penilaianAwalKeperawatanRalanGeriatriRules() map[string]any {
	rules := map[string]any{
		"tanggal":                      "required|date",
		"informasi":                    "required|in:Autoanamnesis,Alloanamnesis",
		"td":                           "string|max_len:8",
		"nadi":                         "string|max_len:5",
		"rr":                           "string|max_len:5",
		"suhu":                         "string|max_len:5",
		"gcs":                          "string|max_len:5",
		"bb":                           "string|max_len:5",
		"tb":                           "string|max_len:5",
		"bmi":                          "string|max_len:10",
		"keluhan_utama":                "string|max_len:150",
		"rpd":                          "string|max_len:100",
		"rpk":                          "string|max_len:100",
		"rpo":                          "string|max_len:100",
		"alergi":                       "string|max_len:25",
		"alat_bantu":                   "required|in:Tidak,Ya",
		"ket_bantu":                    "string|max_len:50",
		"prothesa":                     "required|in:Tidak,Ya",
		"ket_pro":                      "string|max_len:50",
		"adl":                          "required|in:Mandiri,Dibantu",
		"status_psiko":                 "required|in:Tenang,Takut,Cemas,Depresi,Lain-lain",
		"ket_psiko":                    "string|max_len:70",
		"hub_keluarga":                 "required|in:Baik,Tidak Baik",
		"tinggal_dengan":               "required|in:Sendiri,Orang Tua,Suami / Istri,Lainnya",
		"ket_tinggal":                  "string|max_len:40",
		"ekonomi":                      "required|in:Baik,Cukup,Kurang",
		"budaya":                       "required|in:Tidak Ada,Ada",
		"ket_budaya":                   "string|max_len:50",
		"edukasi":                      "required|in:Pasien,Keluarga",
		"ket_edukasi":                  "string|max_len:50",
		"berjalan_a":                   "required|in:Ya,Tidak",
		"berjalan_b":                   "required|in:Ya,Tidak",
		"berjalan_c":                   "required|in:Ya,Tidak",
		"hasil":                        "required|in:Tidak beresiko (tidak ditemukan a dan b),Resiko rendah (ditemukan a/b),Resiko tinggi (ditemukan a dan b)",
		"lapor":                        "required|in:Ya,Tidak",
		"ket_lapor":                    "string|max_len:15",
		"sg1":                          "required|string",
		"nilai1":                       "required|in:0,1,2,3,4",
		"sg2":                          "required|in:Ya,Tidak",
		"nilai2":                       "required|in:0,1",
		"total_hasil":                  "int",
		"nyeri":                        "required|in:Tidak Ada Nyeri,Nyeri Akut,Nyeri Kronis",
		"provokes":                     "required|in:Proses Penyakit,Benturan,Lain-lain",
		"ket_provokes":                 "string|max_len:40",
		"quality":                      "required|in:Seperti Tertusuk,Berdenyut,Teriris,Tertindih,Tertiban,Lain-lain",
		"ket_quality":                  "string|max_len:50",
		"lokasi":                       "string|max_len:50",
		"menyebar":                     "required|in:Tidak,Ya",
		"skala_nyeri":                  "required|in:0,1,2,3,4,5,6,7,8,9,10",
		"durasi":                       "string|max_len:25",
		"nyeri_hilang":                 "required|in:Istirahat,Medengar Musik,Minum Obat",
		"ket_nyeri":                    "string|max_len:40",
		"pada_dokter":                  "required|in:Tidak,Ya",
		"ket_dokter":                   "string|max_len:15",
		"edukasi_kemampuan_bacatulis":  "in:Baik,Kurang,Tidak Bisa",
		"edukasi_kebutuhan_penerjemah": "in:Ya,Tidak",
		"edukasi_keterangan_kebutuhan_penerjemah": "string|max_len:30",
		"edukasi_hambatan":                        "in:Tidak,Ya",
		"edukasi_hambatan_kategori":               "required|in:-,Pendengaran,Penglihatan,Kognitif,Fisik,Budaya,Emosi,Bahasa,Lainnya",
		"edukasi_keterangan_hambatan":             "string|max_len:30",
		"edukasi_cara_bicara":                     "in:Normal,Gangguan Bicara",
		"edukasi_bahasa_isyarat":                  "in:Tidak,Ya",
		"edukasi_menerima_informasi":              "in:Ya,Tidak",
		"edukasi_keterangan_menerima_informasi":   "string|max_len:30",
		"edukasi_metode_belajar":                  "in:Audio,Lisan,Visual,Demonstrasi,Tulisan",
		"fraily_phenotype_berat_badan":            "in:Tidak,Ya",
		"fraily_phenotype_berat_badan_nilai":      "int",
		"fraily_phenotype_aktifitas_fisik":        "in:Tidak Terbatas/Sedikit Terbatas,Sangat Terbatas",
		"fraily_phenotype_aktifitas_fisik_nilai":  "int",
		"fraily_phenotype_kelelahan":              "in:0 - 2 Hari,3 - 7 Hari",
		"fraily_phenotype_kelelahan_nilai":        "int",
		"fraily_phenotype_kekuatan":               "in:Melemah < 20 %,Melemah > 20 %",
		"fraily_phenotype_kekuatan_nilai":         "int",
		"fraily_phenotype_waktu_berjalan":         "in:Tidak Melambat,Melambat",
		"fraily_phenotype_waktu_berjalan_nilai":   "int",
		"fraily_phenotype_nilai_total":            "int",
		"fraily_phenotype_status":                 "in:Sehat,Sedikit Lemah,Lemah,Sangat Lemah",
		"rencana":                                 "string|max_len:200",
		"nip":                                     "required|string|max_len:20",
	}
	rules["kode_masalah"] = "slice"
	rules["kode_masalah.*"] = "string|max_len:20"
	rules["kode_rencana"] = "slice"
	rules["kode_rencana.*"] = "string|max_len:20"
	return rules
}

// PenilaianAwalKeperawatanRalanGeriatriStore simpan penilaian awal keperawatan ralan geriatri; kolom waktu kunci kosong = sekarang.
type PenilaianAwalKeperawatanRalanGeriatriStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianAwalKeperawatanRalanGeriatriData
	PenilaianAwalKeperawatanRalanGeriatriDetail
}

func (r *PenilaianAwalKeperawatanRalanGeriatriStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanRalanGeriatriStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianAwalKeperawatanRalanGeriatriRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianAwalKeperawatanRalanGeriatriStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianAwalKeperawatanRalanGeriatriStore) Payload() PenilaianAwalKeperawatanRalanGeriatriData {
	return r.PenilaianAwalKeperawatanRalanGeriatriData
}

func (r *PenilaianAwalKeperawatanRalanGeriatriStore) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}

// PenilaianAwalKeperawatanRalanGeriatriUpdate ubah penilaian awal keperawatan ralan geriatri (PUT); kunci lewat query string.
type PenilaianAwalKeperawatanRalanGeriatriUpdate struct {
	PenilaianAwalKeperawatanRalanGeriatriData
	PenilaianAwalKeperawatanRalanGeriatriDetail
}

func (r *PenilaianAwalKeperawatanRalanGeriatriUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanRalanGeriatriUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianAwalKeperawatanRalanGeriatriRules()
}

func (r *PenilaianAwalKeperawatanRalanGeriatriUpdate) Payload() PenilaianAwalKeperawatanRalanGeriatriData {
	return r.PenilaianAwalKeperawatanRalanGeriatriData
}

func (r *PenilaianAwalKeperawatanRalanGeriatriUpdate) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}
