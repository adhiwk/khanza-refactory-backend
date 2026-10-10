package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianAwalKeperawatanRalanPsikiatriData isian penilaian awal keperawatan ralan psikiatri.
type PenilaianAwalKeperawatanRalanPsikiatriData struct {
	Tanggal                     string `form:"tanggal" json:"tanggal"`
	Informasi                   string `form:"informasi" json:"informasi"`
	KeluhanUtama                string `form:"keluhan_utama" json:"keluhan_utama"`
	RkdSakitSejak               string `form:"rkd_sakit_sejak" json:"rkd_sakit_sejak"`
	RkdKeluhan                  string `form:"rkd_keluhan" json:"rkd_keluhan"`
	RkdBerobat                  string `form:"rkd_berobat" json:"rkd_berobat"`
	RkdHasilPengobatan          string `form:"rkd_hasil_pengobatan" json:"rkd_hasil_pengobatan"`
	FpPutusObat                 string `form:"fp_putus_obat" json:"fp_putus_obat"`
	KetPutusObat                string `form:"ket_putus_obat" json:"ket_putus_obat"`
	FpEkonomi                   string `form:"fp_ekonomi" json:"fp_ekonomi"`
	KetMasalahEkonomi           string `form:"ket_masalah_ekonomi" json:"ket_masalah_ekonomi"`
	FpMasalahFisik              string `form:"fp_masalah_fisik" json:"fp_masalah_fisik"`
	KetMasalahFisik             string `form:"ket_masalah_fisik" json:"ket_masalah_fisik"`
	FpMasalahPsikososial        string `form:"fp_masalah_psikososial" json:"fp_masalah_psikososial"`
	KetMasalahPsikososial       string `form:"ket_masalah_psikososial" json:"ket_masalah_psikososial"`
	RhKeluarga                  string `form:"rh_keluarga" json:"rh_keluarga"`
	KetRhKeluarga               string `form:"ket_rh_keluarga" json:"ket_rh_keluarga"`
	ResikoBunuhDiri             string `form:"resiko_bunuh_diri" json:"resiko_bunuh_diri"`
	RbdIde                      string `form:"rbd_ide" json:"rbd_ide"`
	KetRbdIde                   string `form:"ket_rbd_ide" json:"ket_rbd_ide"`
	RbdRencana                  string `form:"rbd_rencana" json:"rbd_rencana"`
	KetRbdRencana               string `form:"ket_rbd_rencana" json:"ket_rbd_rencana"`
	RbdAlat                     string `form:"rbd_alat" json:"rbd_alat"`
	KetRbdAlat                  string `form:"ket_rbd_alat" json:"ket_rbd_alat"`
	RbdPercobaan                string `form:"rbd_percobaan" json:"rbd_percobaan"`
	KetRbdPercobaan             string `form:"ket_rbd_percobaan" json:"ket_rbd_percobaan"`
	RbdKeinginan                string `form:"rbd_keinginan" json:"rbd_keinginan"`
	KetRbdKeinginan             string `form:"ket_rbd_keinginan" json:"ket_rbd_keinginan"`
	RpoPenggunaan               string `form:"rpo_penggunaan" json:"rpo_penggunaan"`
	KetRpoPenggunaan            string `form:"ket_rpo_penggunaan" json:"ket_rpo_penggunaan"`
	RpoEfekSamping              string `form:"rpo_efek_samping" json:"rpo_efek_samping"`
	KetRpoEfekSamping           string `form:"ket_rpo_efek_samping" json:"ket_rpo_efek_samping"`
	RpoNapza                    string `form:"rpo_napza" json:"rpo_napza"`
	KetRpoNapza                 string `form:"ket_rpo_napza" json:"ket_rpo_napza"`
	KetLamaPemakaian            string `form:"ket_lama_pemakaian" json:"ket_lama_pemakaian"`
	KetCaraPemakaian            string `form:"ket_cara_pemakaian" json:"ket_cara_pemakaian"`
	KetLatarBelakangPemakaian   string `form:"ket_latar_belakang_pemakaian" json:"ket_latar_belakang_pemakaian"`
	RpoPenggunaanObatLainnya    string `form:"rpo_penggunaan_obat_lainnya" json:"rpo_penggunaan_obat_lainnya"`
	KetPenggunaanObatLainnya    string `form:"ket_penggunaan_obat_lainnya" json:"ket_penggunaan_obat_lainnya"`
	KetAlasanPenggunaan         string `form:"ket_alasan_penggunaan" json:"ket_alasan_penggunaan"`
	RpoAlergiObat               string `form:"rpo_alergi_obat" json:"rpo_alergi_obat"`
	KetAlergiObat               string `form:"ket_alergi_obat" json:"ket_alergi_obat"`
	RpoMerokok                  string `form:"rpo_merokok" json:"rpo_merokok"`
	KetMerokok                  string `form:"ket_merokok" json:"ket_merokok"`
	RpoMinumKopi                string `form:"rpo_minum_kopi" json:"rpo_minum_kopi"`
	KetMinumKopi                string `form:"ket_minum_kopi" json:"ket_minum_kopi"`
	Td                          string `form:"td" json:"td"`
	Nadi                        string `form:"nadi" json:"nadi"`
	Gcs                         string `form:"gcs" json:"gcs"`
	Rr                          string `form:"rr" json:"rr"`
	Suhu                        string `form:"suhu" json:"suhu"`
	PfKeluhanFisik              string `form:"pf_keluhan_fisik" json:"pf_keluhan_fisik"`
	KetKeluhanFisik             string `form:"ket_keluhan_fisik" json:"ket_keluhan_fisik"`
	SkalaNyeri                  string `form:"skala_nyeri" json:"skala_nyeri"`
	Durasi                      string `form:"durasi" json:"durasi"`
	Nyeri                       string `form:"nyeri" json:"nyeri"`
	Provokes                    string `form:"provokes" json:"provokes"`
	KetProvokes                 string `form:"ket_provokes" json:"ket_provokes"`
	Quality                     string `form:"quality" json:"quality"`
	KetQuality                  string `form:"ket_quality" json:"ket_quality"`
	Lokasi                      string `form:"lokasi" json:"lokasi"`
	Menyebar                    string `form:"menyebar" json:"menyebar"`
	PadaDokter                  string `form:"pada_dokter" json:"pada_dokter"`
	KetDokter                   string `form:"ket_dokter" json:"ket_dokter"`
	NyeriHilang                 string `form:"nyeri_hilang" json:"nyeri_hilang"`
	KetNyeri                    string `form:"ket_nyeri" json:"ket_nyeri"`
	Bb                          string `form:"bb" json:"bb"`
	Tb                          string `form:"tb" json:"tb"`
	Bmi                         string `form:"bmi" json:"bmi"`
	LaporStatusNutrisi          string `form:"lapor_status_nutrisi" json:"lapor_status_nutrisi"`
	KetLaporStatusNutrisi       string `form:"ket_lapor_status_nutrisi" json:"ket_lapor_status_nutrisi"`
	Sg1                         string `form:"sg1" json:"sg1"`
	Nilai1                      string `form:"nilai1" json:"nilai1"`
	Sg2                         string `form:"sg2" json:"sg2"`
	Nilai2                      string `form:"nilai2" json:"nilai2"`
	TotalHasil                  int    `form:"total_hasil" json:"total_hasil"`
	Resikojatuh                 string `form:"resikojatuh" json:"resikojatuh"`
	Bjm                         string `form:"bjm" json:"bjm"`
	Msa                         string `form:"msa" json:"msa"`
	Hasil                       string `form:"hasil" json:"hasil"`
	Lapor                       string `form:"lapor" json:"lapor"`
	KetLapor                    string `form:"ket_lapor" json:"ket_lapor"`
	AdlMandi                    string `form:"adl_mandi" json:"adl_mandi"`
	AdlBerpakaian               string `form:"adl_berpakaian" json:"adl_berpakaian"`
	AdlMakan                    string `form:"adl_makan" json:"adl_makan"`
	AdlBak                      string `form:"adl_bak" json:"adl_bak"`
	AdlBab                      string `form:"adl_bab" json:"adl_bab"`
	AdlHobi                     string `form:"adl_hobi" json:"adl_hobi"`
	KetAdlHobi                  string `form:"ket_adl_hobi" json:"ket_adl_hobi"`
	AdlSosialisasi              string `form:"adl_sosialisasi" json:"adl_sosialisasi"`
	KetAdlSosialisasi           string `form:"ket_adl_sosialisasi" json:"ket_adl_sosialisasi"`
	AdlKegiatan                 string `form:"adl_kegiatan" json:"adl_kegiatan"`
	KetAdlKegiatan              string `form:"ket_adl_kegiatan" json:"ket_adl_kegiatan"`
	SkPenampilan                string `form:"sk_penampilan" json:"sk_penampilan"`
	SkAlamPerasaan              string `form:"sk_alam_perasaan" json:"sk_alam_perasaan"`
	SkPembicaraan               string `form:"sk_pembicaraan" json:"sk_pembicaraan"`
	SkAfek                      string `form:"sk_afek" json:"sk_afek"`
	SkAktifitasMotorik          string `form:"sk_aktifitas_motorik" json:"sk_aktifitas_motorik"`
	SkGangguanRingan            string `form:"sk_gangguan_ringan" json:"sk_gangguan_ringan"`
	SkProsesPikir               string `form:"sk_proses_pikir" json:"sk_proses_pikir"`
	SkOrientasi                 string `form:"sk_orientasi" json:"sk_orientasi"`
	SkTingkatKesadaranOrientasi string `form:"sk_tingkat_kesadaran_orientasi" json:"sk_tingkat_kesadaran_orientasi"`
	SkMemori                    string `form:"sk_memori" json:"sk_memori"`
	SkInteraksi                 string `form:"sk_interaksi" json:"sk_interaksi"`
	SkKonsentrasi               string `form:"sk_konsentrasi" json:"sk_konsentrasi"`
	SkPersepsi                  string `form:"sk_persepsi" json:"sk_persepsi"`
	KetSkPersepsi               string `form:"ket_sk_persepsi" json:"ket_sk_persepsi"`
	SkIsiPikir                  string `form:"sk_isi_pikir" json:"sk_isi_pikir"`
	SkWaham                     string `form:"sk_waham" json:"sk_waham"`
	KetSkWaham                  string `form:"ket_sk_waham" json:"ket_sk_waham"`
	SkDayaTilikDiri             string `form:"sk_daya_tilik_diri" json:"sk_daya_tilik_diri"`
	KetSkDayaTilikDiri          string `form:"ket_sk_daya_tilik_diri" json:"ket_sk_daya_tilik_diri"`
	KkPembelajaran              string `form:"kk_pembelajaran" json:"kk_pembelajaran"`
	KetKkPembelajaran           string `form:"ket_kk_pembelajaran" json:"ket_kk_pembelajaran"`
	KetKkPembelajaranLainnya    string `form:"ket_kk_pembelajaran_lainnya" json:"ket_kk_pembelajaran_lainnya"`
	KkPenerjamah                string `form:"kk_Penerjamah" json:"kk_Penerjamah"`
	KetKkPenerjamahLainnya      string `form:"ket_kk_penerjamah_Lainnya" json:"ket_kk_penerjamah_Lainnya"`
	KkBahasaIsyarat             string `form:"kk_bahasa_isyarat" json:"kk_bahasa_isyarat"`
	KkKebutuhanEdukasi          string `form:"kk_kebutuhan_edukasi" json:"kk_kebutuhan_edukasi"`
	KetKkKebutuhanEdukasi       string `form:"ket_kk_kebutuhan_edukasi" json:"ket_kk_kebutuhan_edukasi"`
	Rencana                     string `form:"rencana" json:"rencana"`
	Nip                         string `form:"nip" json:"nip"`
}

// PenilaianAwalKeperawatanRalanPsikiatriDetail daftar kode detail (diganti seluruhnya saat simpan/ubah).
type PenilaianAwalKeperawatanRalanPsikiatriDetail struct {
	KodeMasalah []string `form:"kode_masalah" json:"kode_masalah"`
	KodeRencana []string `form:"kode_rencana" json:"kode_rencana"`
}

func penilaianAwalKeperawatanRalanPsikiatriRules() map[string]any {
	rules := map[string]any{
		"tanggal":                        "required|date",
		"informasi":                      "required|in:Autoanamnesis,Alloanamnesis",
		"keluhan_utama":                  "string|max_len:500",
		"rkd_sakit_sejak":                "string|max_len:8",
		"rkd_keluhan":                    "string|max_len:500",
		"rkd_berobat":                    "required|string",
		"rkd_hasil_pengobatan":           "required|in:Berhasil,Tidak Berhasil",
		"fp_putus_obat":                  "required|in:Tidak,Ya",
		"ket_putus_obat":                 "string|max_len:50",
		"fp_ekonomi":                     "required|in:Tidak,Ya",
		"ket_masalah_ekonomi":            "string|max_len:50",
		"fp_masalah_fisik":               "required|in:Tidak,Ya",
		"ket_masalah_fisik":              "string|max_len:50",
		"fp_masalah_psikososial":         "required|in:Tidak,Ya",
		"ket_masalah_psikososial":        "string|max_len:50",
		"rh_keluarga":                    "required|in:Tidak,Ya",
		"ket_rh_keluarga":                "string|max_len:50",
		"resiko_bunuh_diri":              "required|in:Tidak,Ya",
		"rbd_ide":                        "required|in:Tidak,Ya",
		"ket_rbd_ide":                    "string|max_len:50",
		"rbd_rencana":                    "required|in:Tidak,Ya",
		"ket_rbd_rencana":                "string|max_len:50",
		"rbd_alat":                       "required|in:Tidak,Ya",
		"ket_rbd_alat":                   "string|max_len:50",
		"rbd_percobaan":                  "required|in:Tidak,Ya",
		"ket_rbd_percobaan":              "string|max_len:15",
		"rbd_keinginan":                  "required|in:Tidak,Ya",
		"ket_rbd_keinginan":              "string|max_len:100",
		"rpo_penggunaan":                 "required|in:Tidak,Ya",
		"ket_rpo_penggunaan":             "string|max_len:20",
		"rpo_efek_samping":               "required|in:Tidak,Ya",
		"ket_rpo_efek_samping":           "string|max_len:20",
		"rpo_napza":                      "required|in:Tidak,Ya",
		"ket_rpo_napza":                  "string|max_len:25",
		"ket_lama_pemakaian":             "string|max_len:8",
		"ket_cara_pemakaian":             "string|max_len:15",
		"ket_latar_belakang_pemakaian":   "string|max_len:60",
		"rpo_penggunaan_obat_lainnya":    "required|in:Tidak,Ya",
		"ket_penggunaan_obat_lainnya":    "string|max_len:20",
		"ket_alasan_penggunaan":          "string|max_len:65",
		"rpo_alergi_obat":                "required|in:Tidak,Ya",
		"ket_alergi_obat":                "string|max_len:25",
		"rpo_merokok":                    "required|in:Tidak,Ya",
		"ket_merokok":                    "string|max_len:25",
		"rpo_minum_kopi":                 "required|in:Tidak,Ya",
		"ket_minum_kopi":                 "string|max_len:25",
		"td":                             "string|max_len:8",
		"nadi":                           "string|max_len:5",
		"gcs":                            "string|max_len:5",
		"rr":                             "string|max_len:5",
		"suhu":                           "string|max_len:5",
		"pf_keluhan_fisik":               "required|in:Tidak,Ya",
		"ket_keluhan_fisik":              "string|max_len:100",
		"skala_nyeri":                    "required|in:0,1,2,3,4,5,6,7,8,9,10",
		"durasi":                         "string|max_len:25",
		"nyeri":                          "required|in:Tidak Ada Nyeri,Nyeri Akut,Nyeri Kronis",
		"provokes":                       "required|in:Proses Penyakit,Benturan,Lain-lain",
		"ket_provokes":                   "string|max_len:40",
		"quality":                        "required|in:Seperti Tertusuk,Berdenyut,Teriris,Tertindih,Tertiban,Lain-lain",
		"ket_quality":                    "string|max_len:50",
		"lokasi":                         "string|max_len:50",
		"menyebar":                       "required|in:Tidak,Ya",
		"pada_dokter":                    "required|in:Tidak,Ya",
		"ket_dokter":                     "string|max_len:15",
		"nyeri_hilang":                   "required|in:Istirahat,Medengar Musik,Minum Obat",
		"ket_nyeri":                      "string|max_len:40",
		"bb":                             "string|max_len:5",
		"tb":                             "string|max_len:5",
		"bmi":                            "string|max_len:5",
		"lapor_status_nutrisi":           "required|in:Ya,Tidak",
		"ket_lapor_status_nutrisi":       "string|max_len:15",
		"sg1":                            "required|string",
		"nilai1":                         "required|in:0,1,2,3,4",
		"sg2":                            "required|in:Ya,Tidak",
		"nilai2":                         "required|in:0,1",
		"total_hasil":                    "int",
		"resikojatuh":                    "required|in:Ya,Tidak",
		"bjm":                            "required|in:Ya,Tidak",
		"msa":                            "required|in:Ya,Tidak",
		"hasil":                          "required|in:Tidak beresiko (tidak ditemukan a dan b),Resiko rendah (ditemukan a/b),Resiko tinggi (ditemukan a dan b)",
		"lapor":                          "required|in:Ya,Tidak",
		"ket_lapor":                      "string|max_len:15",
		"adl_mandi":                      "required|in:Mandiri,Bantuan Minimal,Bantuan Total",
		"adl_berpakaian":                 "required|in:Mandiri,Bantuan Minimal,Bantuan Total",
		"adl_makan":                      "required|in:Mandiri,Bantuan Minimal,Bantuan Total",
		"adl_bak":                        "required|in:Mandiri,Bantuan Minimal,Bantuan Total",
		"adl_bab":                        "required|in:Mandiri,Bantuan Minimal,Bantuan Total",
		"adl_hobi":                       "required|in:Ya,Tidak",
		"ket_adl_hobi":                   "string|max_len:50",
		"adl_sosialisasi":                "required|in:Ya,Tidak",
		"ket_adl_sosialisasi":            "string|max_len:50",
		"adl_kegiatan":                   "required|in:Ya,Tidak",
		"ket_adl_kegiatan":               "string|max_len:50",
		"sk_penampilan":                  "required|in:Bersih,Rapi,Tidak Rapi,Kotor,Tidak Seperti Biasanya,Pakaian Tidak Sesuai",
		"sk_alam_perasaan":               "required|in:Sesuai,Marah,Putus Asa,Tertekan,Sedih,Labil,Malu,Khawatir,Gembira Berlebihan,Merasa Tidak Mampu,Ketakutan,Tidak Berguna",
		"sk_pembicaraan":                 "required|in:Sesuai,Cepat,Lambat,Membisu,Mendominasi,Mengancam,Inkoheren,Apatis,Keras,Gagap,Tidak Mampu Memulai Pembicaraan",
		"sk_afek":                        "required|in:Sesuai,Datar,Tumpul,Labil,Tidak Sesuai",
		"sk_aktifitas_motorik":           "required|in:Normal,Tegang,Gelisah,Lesuh,Grimasem,TIK,Tremor,Agitasi,Konfulsif,Melamun,Sulit Diarahkan",
		"sk_gangguan_ringan":             "required|in:Gangguan Ringan,Gangguan Bermakna",
		"sk_proses_pikir":                "required|in:Sesuai,Sirkumsial,Kehilangan Asosiasi,Flight Of Ideas,Bloking,Pengulangan Pembicaraan,Tangensial",
		"sk_orientasi":                   "required|in:Tidak,Ya",
		"sk_tingkat_kesadaran_orientasi": "required|in:-,Bingung,Sedasi,Waktu,Stupor,Tempat,Orang",
		"sk_memori":                      "required|in:Ganguan Daya Ingat Jangka Pendek,Ganguan Daya Ingat Jangka Panjang,Ganguan Daya Ingat Saat Ini,Konfabulasi",
		"sk_interaksi":                   "required|in:Kooperatif,Tidak Kooperatif,Bermusuhan,Mudah Tersinggung,Curiga,Defensif,Kontak Mata Kurang",
		"sk_konsentrasi":                 "required|in:Konsentrasi Baik,Mudah Beralih,Tidak Mampu Berkonsentrasi,Tidak Mampu Berhitung Sederhana",
		"sk_persepsi":                    "required|in:Halusinasi,Pendengaran,Penghidung,Penglihatan,Pengecapan,Perabaan",
		"ket_sk_persepsi":                "string|max_len:70",
		"sk_isi_pikir":                   "required|in:Sesuai,Obsesi,Fobia,Hipokondria,Depersonalisasi,Pikiran Magis,Ide Yang Terkait,Waham",
		"sk_waham":                       "required|in:Kebesaran,Curiga,Agama,Nihilistik",
		"ket_sk_waham":                   "string|max_len:100",
		"sk_daya_tilik_diri":             "required|in:Mengingkari Penyakit Yang Diderita,Menyalahkan Hal-hal Diluar Dirinya",
		"ket_sk_daya_tilik_diri":         "string|max_len:100",
		"kk_pembelajaran":                "required|in:Tidak,Ya",
		"ket_kk_pembelajaran":            "required|in:-,Pendengaran,Penglihatan,Kognitif,Fisik,Budaya,Emosi,Bahasa,Lainnya",
		"ket_kk_pembelajaran_lainnya":    "string|max_len:50",
		"kk_Penerjamah":                  "required|in:Tidak,Ya",
		"ket_kk_penerjamah_Lainnya":      "string|max_len:50",
		"kk_bahasa_isyarat":              "required|in:Tidak,Ya",
		"kk_kebutuhan_edukasi":           "required|in:Diagnosa Dan Manajemen Penyakit,Obat-obatan/Terapi,Diet Dan Nutrisi,Tindakan Keperawatan,Rehabilitasi,Manajemen Nyeri,Lain-lain",
		"ket_kk_kebutuhan_edukasi":       "string|max_len:50",
		"rencana":                        "string|max_len:200",
		"nip":                            "required|string|max_len:20",
	}
	rules["kode_masalah"] = "slice"
	rules["kode_masalah.*"] = "string|max_len:20"
	rules["kode_rencana"] = "slice"
	rules["kode_rencana.*"] = "string|max_len:20"
	return rules
}

// PenilaianAwalKeperawatanRalanPsikiatriStore simpan penilaian awal keperawatan ralan psikiatri; kolom waktu kunci kosong = sekarang.
type PenilaianAwalKeperawatanRalanPsikiatriStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianAwalKeperawatanRalanPsikiatriData
	PenilaianAwalKeperawatanRalanPsikiatriDetail
}

func (r *PenilaianAwalKeperawatanRalanPsikiatriStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanRalanPsikiatriStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianAwalKeperawatanRalanPsikiatriRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianAwalKeperawatanRalanPsikiatriStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianAwalKeperawatanRalanPsikiatriStore) Payload() PenilaianAwalKeperawatanRalanPsikiatriData {
	return r.PenilaianAwalKeperawatanRalanPsikiatriData
}

func (r *PenilaianAwalKeperawatanRalanPsikiatriStore) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}

// PenilaianAwalKeperawatanRalanPsikiatriUpdate ubah penilaian awal keperawatan ralan psikiatri (PUT); kunci lewat query string.
type PenilaianAwalKeperawatanRalanPsikiatriUpdate struct {
	PenilaianAwalKeperawatanRalanPsikiatriData
	PenilaianAwalKeperawatanRalanPsikiatriDetail
}

func (r *PenilaianAwalKeperawatanRalanPsikiatriUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanRalanPsikiatriUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianAwalKeperawatanRalanPsikiatriRules()
}

func (r *PenilaianAwalKeperawatanRalanPsikiatriUpdate) Payload() PenilaianAwalKeperawatanRalanPsikiatriData {
	return r.PenilaianAwalKeperawatanRalanPsikiatriData
}

func (r *PenilaianAwalKeperawatanRalanPsikiatriUpdate) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}
