package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianAwalKeperawatanKebidananData isian penilaian awal keperawatan kebidanan.
type PenilaianAwalKeperawatanKebidananData struct {
	Tanggal       string `form:"tanggal" json:"tanggal"`
	Informasi     string `form:"informasi" json:"informasi"`
	Td            string `form:"td" json:"td"`
	Nadi          string `form:"nadi" json:"nadi"`
	Rr            string `form:"rr" json:"rr"`
	Suhu          string `form:"suhu" json:"suhu"`
	Gcs           string `form:"gcs" json:"gcs"`
	Bb            string `form:"bb" json:"bb"`
	Tb            string `form:"tb" json:"tb"`
	Lila          string `form:"lila" json:"lila"`
	Bmi           string `form:"bmi" json:"bmi"`
	Tfu           string `form:"tfu" json:"tfu"`
	Tbj           string `form:"tbj" json:"tbj"`
	Letak         string `form:"letak" json:"letak"`
	Presentasi    string `form:"presentasi" json:"presentasi"`
	Penurunan     string `form:"penurunan" json:"penurunan"`
	His           string `form:"his" json:"his"`
	Kekuatan      string `form:"kekuatan" json:"kekuatan"`
	Lamanya       string `form:"lamanya" json:"lamanya"`
	Bjj           string `form:"bjj" json:"bjj"`
	KetBjj        string `form:"ket_bjj" json:"ket_bjj"`
	Portio        string `form:"portio" json:"portio"`
	Serviks       string `form:"serviks" json:"serviks"`
	Ketuban       string `form:"ketuban" json:"ketuban"`
	Hodge         string `form:"hodge" json:"hodge"`
	Inspekulo     string `form:"inspekulo" json:"inspekulo"`
	KetInspekulo  string `form:"ket_inspekulo" json:"ket_inspekulo"`
	Ctg           string `form:"ctg" json:"ctg"`
	KetCtg        string `form:"ket_ctg" json:"ket_ctg"`
	Usg           string `form:"usg" json:"usg"`
	KetUsg        string `form:"ket_usg" json:"ket_usg"`
	Lab           string `form:"lab" json:"lab"`
	KetLab        string `form:"ket_lab" json:"ket_lab"`
	Lakmus        string `form:"lakmus" json:"lakmus"`
	KetLakmus     string `form:"ket_lakmus" json:"ket_lakmus"`
	Panggul       string `form:"panggul" json:"panggul"`
	KeluhanUtama  string `form:"keluhan_utama" json:"keluhan_utama"`
	Umur          string `form:"umur" json:"umur"`
	Lama          string `form:"lama" json:"lama"`
	Banyaknya     string `form:"banyaknya" json:"banyaknya"`
	Haid          string `form:"haid" json:"haid"`
	Siklus        string `form:"siklus" json:"siklus"`
	KetSiklus     string `form:"ket_siklus" json:"ket_siklus"`
	KetSiklus1    string `form:"ket_siklus1" json:"ket_siklus1"`
	Status        string `form:"status" json:"status"`
	Kali          string `form:"kali" json:"kali"`
	Usia1         string `form:"usia1" json:"usia1"`
	Ket1          string `form:"ket1" json:"ket1"`
	Usia2         string `form:"usia2" json:"usia2"`
	Ket2          string `form:"ket2" json:"ket2"`
	Usia3         string `form:"usia3" json:"usia3"`
	Ket3          string `form:"ket3" json:"ket3"`
	Hpht          string `form:"hpht" json:"hpht"`
	UsiaKehamilan string `form:"usia_kehamilan" json:"usia_kehamilan"`
	Tp            string `form:"tp" json:"tp"`
	Imunisasi     string `form:"imunisasi" json:"imunisasi"`
	KetImunisasi  string `form:"ket_imunisasi" json:"ket_imunisasi"`
	G             string `form:"g" json:"g"`
	P             string `form:"p" json:"p"`
	A             string `form:"a" json:"a"`
	Hidup         string `form:"hidup" json:"hidup"`
	Ginekologi    string `form:"ginekologi" json:"ginekologi"`
	Kebiasaan     string `form:"kebiasaan" json:"kebiasaan"`
	KetKebiasaan  string `form:"ket_kebiasaan" json:"ket_kebiasaan"`
	Kebiasaan1    string `form:"kebiasaan1" json:"kebiasaan1"`
	KetKebiasaan1 string `form:"ket_kebiasaan1" json:"ket_kebiasaan1"`
	Kebiasaan2    string `form:"kebiasaan2" json:"kebiasaan2"`
	KetKebiasaan2 string `form:"ket_kebiasaan2" json:"ket_kebiasaan2"`
	Kebiasaan3    string `form:"kebiasaan3" json:"kebiasaan3"`
	Kb            string `form:"kb" json:"kb"`
	KetKb         string `form:"ket_kb" json:"ket_kb"`
	Komplikasi    string `form:"komplikasi" json:"komplikasi"`
	KetKomplikasi string `form:"ket_komplikasi" json:"ket_komplikasi"`
	Berhenti      string `form:"berhenti" json:"berhenti"`
	Alasan        string `form:"alasan" json:"alasan"`
	AlatBantu     string `form:"alat_bantu" json:"alat_bantu"`
	KetBantu      string `form:"ket_bantu" json:"ket_bantu"`
	Prothesa      string `form:"prothesa" json:"prothesa"`
	KetPro        string `form:"ket_pro" json:"ket_pro"`
	Adl           string `form:"adl" json:"adl"`
	StatusPsiko   string `form:"status_psiko" json:"status_psiko"`
	KetPsiko      string `form:"ket_psiko" json:"ket_psiko"`
	HubKeluarga   string `form:"hub_keluarga" json:"hub_keluarga"`
	TinggalDengan string `form:"tinggal_dengan" json:"tinggal_dengan"`
	KetTinggal    string `form:"ket_tinggal" json:"ket_tinggal"`
	Ekonomi       string `form:"ekonomi" json:"ekonomi"`
	Budaya        string `form:"budaya" json:"budaya"`
	KetBudaya     string `form:"ket_budaya" json:"ket_budaya"`
	Edukasi       string `form:"edukasi" json:"edukasi"`
	KetEdukasi    string `form:"ket_edukasi" json:"ket_edukasi"`
	BerjalanA     string `form:"berjalan_a" json:"berjalan_a"`
	BerjalanB     string `form:"berjalan_b" json:"berjalan_b"`
	BerjalanC     string `form:"berjalan_c" json:"berjalan_c"`
	Hasil         string `form:"hasil" json:"hasil"`
	Lapor         string `form:"lapor" json:"lapor"`
	KetLapor      string `form:"ket_lapor" json:"ket_lapor"`
	Sg1           string `form:"sg1" json:"sg1"`
	Nilai1        string `form:"nilai1" json:"nilai1"`
	Sg2           string `form:"sg2" json:"sg2"`
	Nilai2        string `form:"nilai2" json:"nilai2"`
	TotalHasil    string `form:"total_hasil" json:"total_hasil"`
	Nyeri         string `form:"nyeri" json:"nyeri"`
	Provokes      string `form:"provokes" json:"provokes"`
	KetProvokes   string `form:"ket_provokes" json:"ket_provokes"`
	Quality       string `form:"quality" json:"quality"`
	KetQuality    string `form:"ket_quality" json:"ket_quality"`
	Lokasi        string `form:"lokasi" json:"lokasi"`
	Menyebar      string `form:"menyebar" json:"menyebar"`
	SkalaNyeri    string `form:"skala_nyeri" json:"skala_nyeri"`
	Durasi        string `form:"durasi" json:"durasi"`
	NyeriHilang   string `form:"nyeri_hilang" json:"nyeri_hilang"`
	KetNyeri      string `form:"ket_nyeri" json:"ket_nyeri"`
	PadaDokter    string `form:"pada_dokter" json:"pada_dokter"`
	KetDokter     string `form:"ket_dokter" json:"ket_dokter"`
	Masalah       string `form:"masalah" json:"masalah"`
	Tindakan      string `form:"tindakan" json:"tindakan"`
	Nip           string `form:"nip" json:"nip"`
}

func penilaianAwalKeperawatanKebidananRules() map[string]any {
	rules := map[string]any{
		"tanggal":        "required|date",
		"informasi":      "required|in:Autoanamnesis,Alloanamnesis",
		"td":             "string|max_len:8",
		"nadi":           "string|max_len:5",
		"rr":             "string|max_len:5",
		"suhu":           "string|max_len:5",
		"gcs":            "string|max_len:10",
		"bb":             "string|max_len:5",
		"tb":             "string|max_len:5",
		"lila":           "string|max_len:5",
		"bmi":            "string|max_len:10",
		"tfu":            "string|max_len:10",
		"tbj":            "string|max_len:10",
		"letak":          "string|max_len:10",
		"presentasi":     "string|max_len:10",
		"penurunan":      "string|max_len:10",
		"his":            "string|max_len:10",
		"kekuatan":       "string|max_len:10",
		"lamanya":        "string|max_len:10",
		"bjj":            "string|max_len:10",
		"ket_bjj":        "required|in:Teratur,Tidak Teratur",
		"portio":         "string|max_len:10",
		"serviks":        "string|max_len:10",
		"ketuban":        "string|max_len:10",
		"hodge":          "string|max_len:10",
		"inspekulo":      "required|in:Dilakukan,Tidak",
		"ket_inspekulo":  "string|max_len:50",
		"ctg":            "required|in:Dilakukan,Tidak",
		"ket_ctg":        "string|max_len:50",
		"usg":            "required|in:Dilakukan,Tidak",
		"ket_usg":        "string|max_len:50",
		"lab":            "required|in:Dilakukan,Tidak",
		"ket_lab":        "string|max_len:50",
		"lakmus":         "required|in:Dilakukan,Tidak",
		"ket_lakmus":     "string|max_len:50",
		"panggul":        "required|in:Luas,Sedang,Sempit,Tidak Dilakukan Pemeriksaan",
		"keluhan_utama":  "string|max_len:1000",
		"umur":           "string|max_len:10",
		"lama":           "string|max_len:10",
		"banyaknya":      "string|max_len:10",
		"haid":           "string|max_len:20",
		"siklus":         "string|max_len:10",
		"ket_siklus":     "required|in:Teratur,Tidak Teratur",
		"ket_siklus1":    "required|in:Tidak Ada Masalah,Dismenorhea,Spotting,Menorhagia,PMS",
		"status":         "required|in:Menikah,Tidak / Belum Menikah",
		"kali":           "string|max_len:5",
		"usia1":          "string|max_len:5",
		"ket1":           "required|in:-,Masih Menikah,Cerai,Meninggal",
		"usia2":          "string|max_len:5",
		"ket2":           "in:-,Masih Menikah,Cerai,Meninggal",
		"usia3":          "string|max_len:5",
		"ket3":           "in:-,Masih Menikah,Cerai,Meninggal",
		"hpht":           "date",
		"usia_kehamilan": "string|max_len:10",
		"tp":             "date",
		"imunisasi":      "required|in:Tidak,Ya",
		"ket_imunisasi":  "string|max_len:10",
		"g":              "string|max_len:10",
		"p":              "string|max_len:10",
		"a":              "string|max_len:10",
		"hidup":          "string|max_len:10",
		"ginekologi":     "required|in:Tidak Ada,Infertilitas,Infeksi Virus,PMS,Cervisitis Kronis,Endometriosis,Mioma,Polip Cervix,Kanker Kandungan,Operasi Kandungan",
		"kebiasaan":      "required|in:-,Obat Obatan,Vitamin,Jamu Jamuan",
		"ket_kebiasaan":  "string|max_len:50",
		"kebiasaan1":     "required|in:Tidak,Ya",
		"ket_kebiasaan1": "string|max_len:5",
		"kebiasaan2":     "required|in:Tidak,Ya",
		"ket_kebiasaan2": "string|max_len:5",
		"kebiasaan3":     "required|in:Tidak,Ya",
		"kb":             "required|in:Belum Pernah,Suntik,Pil,AKDR,MOW,Implan,Kondom,Kalender,MAL,Coitus Interuptus",
		"ket_kb":         "string|max_len:10",
		"komplikasi":     "required|in:Tidak Ada,Ada",
		"ket_komplikasi": "string|max_len:50",
		"berhenti":       "string|max_len:20",
		"alasan":         "string|max_len:50",
		"alat_bantu":     "required|in:Tidak,Ya",
		"ket_bantu":      "string|max_len:50",
		"prothesa":       "required|in:Tidak,Ya",
		"ket_pro":        "string|max_len:50",
		"adl":            "required|in:Mandiri,Dibantu",
		"status_psiko":   "required|in:Tenang,Takut,Cemas,Depresi,Lain-Lain",
		"ket_psiko":      "string|max_len:50",
		"hub_keluarga":   "required|in:Baik,Tidak Baik",
		"tinggal_dengan": "required|in:Sendiri,Orang Tua,Suami / Istri,Lainnya",
		"ket_tinggal":    "string|max_len:50",
		"ekonomi":        "required|in:Baik,Cukup,Kurang",
		"budaya":         "required|in:Tidak Ada,Ada",
		"ket_budaya":     "string|max_len:50",
		"edukasi":        "required|in:Pasien,Keluarga",
		"ket_edukasi":    "string|max_len:50",
		"berjalan_a":     "required|in:Ya,Tidak",
		"berjalan_b":     "required|in:Ya,Tidak",
		"berjalan_c":     "required|in:Ya,Tidak",
		"hasil":          "required|in:Tidak Beresiko (Tidak Ditemukan A Dan B),Resiko Rendah (Ditemukan A/B),Resiko Tinggi (Ditemukan A Dan B)",
		"lapor":          "required|in:Ya,Tidak",
		"ket_lapor":      "string|max_len:10",
		"sg1":            "required|string",
		"nilai1":         "required|in:0,1,2,3,4",
		"sg2":            "required|in:Ya,Tidak",
		"nilai2":         "required|in:0,1",
		"total_hasil":    "string|max_len:5",
		"nyeri":          "required|in:Tidak Ada Nyeri,Nyeri Akut,Nyeri Kronis",
		"provokes":       "required|in:Proses Penyakit,Benturan,Lain-Lain",
		"ket_provokes":   "string|max_len:40",
		"quality":        "required|in:Seperti Tertusuk,Berdenyut,Teriris,Tertindih,Tertiban,Lain-Lain",
		"ket_quality":    "string|max_len:40",
		"lokasi":         "string|max_len:40",
		"menyebar":       "required|in:Tidak,Ya",
		"skala_nyeri":    "required|in:0,1,2,3,4,5,6,7,8,9,10",
		"durasi":         "string|max_len:5",
		"nyeri_hilang":   "required|in:Istirahat,Mendengar Musik,Minum Obat",
		"ket_nyeri":      "string|max_len:40",
		"pada_dokter":    "required|in:Tidak,Ya",
		"ket_dokter":     "string|max_len:10",
		"masalah":        "string|max_len:1000",
		"tindakan":       "string|max_len:1000",
		"nip":            "required|string|max_len:20",
	}
	return rules
}

// PenilaianAwalKeperawatanKebidananStore simpan penilaian awal keperawatan kebidanan; kolom waktu kunci kosong = sekarang.
type PenilaianAwalKeperawatanKebidananStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianAwalKeperawatanKebidananData
}

func (r *PenilaianAwalKeperawatanKebidananStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanKebidananStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianAwalKeperawatanKebidananRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianAwalKeperawatanKebidananStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianAwalKeperawatanKebidananStore) Payload() PenilaianAwalKeperawatanKebidananData {
	return r.PenilaianAwalKeperawatanKebidananData
}

func (r *PenilaianAwalKeperawatanKebidananStore) DetailValues() map[string][]string { return nil }

// PenilaianAwalKeperawatanKebidananUpdate ubah penilaian awal keperawatan kebidanan (PUT); kunci lewat query string.
type PenilaianAwalKeperawatanKebidananUpdate struct {
	PenilaianAwalKeperawatanKebidananData
}

func (r *PenilaianAwalKeperawatanKebidananUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanKebidananUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianAwalKeperawatanKebidananRules()
}

func (r *PenilaianAwalKeperawatanKebidananUpdate) Payload() PenilaianAwalKeperawatanKebidananData {
	return r.PenilaianAwalKeperawatanKebidananData
}

func (r *PenilaianAwalKeperawatanKebidananUpdate) DetailValues() map[string][]string { return nil }
