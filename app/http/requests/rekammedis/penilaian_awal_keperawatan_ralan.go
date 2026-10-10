package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianAwalKeperawatanRalanData isian penilaian awal keperawatan ralan.
type PenilaianAwalKeperawatanRalanData struct {
	Tanggal       string `form:"tanggal" json:"tanggal"`
	Informasi     string `form:"informasi" json:"informasi"`
	Td            string `form:"td" json:"td"`
	Nadi          string `form:"nadi" json:"nadi"`
	Rr            string `form:"rr" json:"rr"`
	Suhu          string `form:"suhu" json:"suhu"`
	Gcs           string `form:"gcs" json:"gcs"`
	Bb            string `form:"bb" json:"bb"`
	Tb            string `form:"tb" json:"tb"`
	Bmi           string `form:"bmi" json:"bmi"`
	KeluhanUtama  string `form:"keluhan_utama" json:"keluhan_utama"`
	Rpd           string `form:"rpd" json:"rpd"`
	Rpk           string `form:"rpk" json:"rpk"`
	Rpo           string `form:"rpo" json:"rpo"`
	Alergi        string `form:"alergi" json:"alergi"`
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
	TotalHasil    int    `form:"total_hasil" json:"total_hasil"`
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
	Rencana       string `form:"rencana" json:"rencana"`
	Nip           string `form:"nip" json:"nip"`
}

// PenilaianAwalKeperawatanRalanDetail daftar kode detail (diganti seluruhnya saat simpan/ubah).
type PenilaianAwalKeperawatanRalanDetail struct {
	KodeMasalah []string `form:"kode_masalah" json:"kode_masalah"`
	KodeRencana []string `form:"kode_rencana" json:"kode_rencana"`
}

func penilaianAwalKeperawatanRalanRules() map[string]any {
	rules := map[string]any{
		"tanggal":        "required|date",
		"informasi":      "required|in:Autoanamnesis,Alloanamnesis",
		"td":             "string|max_len:8",
		"nadi":           "string|max_len:5",
		"rr":             "string|max_len:5",
		"suhu":           "string|max_len:5",
		"gcs":            "string|max_len:5",
		"bb":             "string|max_len:5",
		"tb":             "string|max_len:5",
		"bmi":            "string|max_len:10",
		"keluhan_utama":  "string|max_len:150",
		"rpd":            "string|max_len:100",
		"rpk":            "string|max_len:100",
		"rpo":            "string|max_len:100",
		"alergi":         "string|max_len:25",
		"alat_bantu":     "required|in:Tidak,Ya",
		"ket_bantu":      "string|max_len:50",
		"prothesa":       "required|in:Tidak,Ya",
		"ket_pro":        "string|max_len:50",
		"adl":            "required|in:Mandiri,Dibantu",
		"status_psiko":   "required|in:Tenang,Takut,Cemas,Depresi,Lain-lain",
		"ket_psiko":      "string|max_len:70",
		"hub_keluarga":   "required|in:Baik,Tidak Baik",
		"tinggal_dengan": "required|in:Sendiri,Orang Tua,Suami / Istri,Lainnya",
		"ket_tinggal":    "string|max_len:40",
		"ekonomi":        "required|in:Baik,Cukup,Kurang",
		"budaya":         "required|in:Tidak Ada,Ada",
		"ket_budaya":     "string|max_len:50",
		"edukasi":        "required|in:Pasien,Keluarga",
		"ket_edukasi":    "string|max_len:50",
		"berjalan_a":     "required|in:Ya,Tidak",
		"berjalan_b":     "required|in:Ya,Tidak",
		"berjalan_c":     "required|in:Ya,Tidak",
		"hasil":          "required|in:Tidak beresiko (tidak ditemukan a dan b),Resiko rendah (ditemukan a/b),Resiko tinggi (ditemukan a dan b)",
		"lapor":          "required|in:Ya,Tidak",
		"ket_lapor":      "string|max_len:15",
		"sg1":            "required|string",
		"nilai1":         "required|in:0,1,2,3,4",
		"sg2":            "required|in:Ya,Tidak",
		"nilai2":         "required|in:0,1",
		"total_hasil":    "int",
		"nyeri":          "required|in:Tidak Ada Nyeri,Nyeri Akut,Nyeri Kronis",
		"provokes":       "required|in:Proses Penyakit,Benturan,Lain-lain",
		"ket_provokes":   "string|max_len:40",
		"quality":        "required|in:Seperti Tertusuk,Berdenyut,Teriris,Tertindih,Tertiban,Lain-lain",
		"ket_quality":    "string|max_len:50",
		"lokasi":         "string|max_len:50",
		"menyebar":       "required|in:Tidak,Ya",
		"skala_nyeri":    "required|in:0,1,2,3,4,5,6,7,8,9,10",
		"durasi":         "string|max_len:25",
		"nyeri_hilang":   "required|in:Istirahat,Medengar Musik,Minum Obat",
		"ket_nyeri":      "string|max_len:40",
		"pada_dokter":    "required|in:Tidak,Ya",
		"ket_dokter":     "string|max_len:15",
		"rencana":        "string|max_len:200",
		"nip":            "required|string|max_len:20",
	}
	rules["kode_masalah"] = "slice"
	rules["kode_masalah.*"] = "string|max_len:20"
	rules["kode_rencana"] = "slice"
	rules["kode_rencana.*"] = "string|max_len:20"
	return rules
}

// PenilaianAwalKeperawatanRalanStore simpan penilaian awal keperawatan ralan; kolom waktu kunci kosong = sekarang.
type PenilaianAwalKeperawatanRalanStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianAwalKeperawatanRalanData
	PenilaianAwalKeperawatanRalanDetail
}

func (r *PenilaianAwalKeperawatanRalanStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanRalanStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianAwalKeperawatanRalanRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianAwalKeperawatanRalanStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianAwalKeperawatanRalanStore) Payload() PenilaianAwalKeperawatanRalanData {
	return r.PenilaianAwalKeperawatanRalanData
}

func (r *PenilaianAwalKeperawatanRalanStore) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}

// PenilaianAwalKeperawatanRalanUpdate ubah penilaian awal keperawatan ralan (PUT); kunci lewat query string.
type PenilaianAwalKeperawatanRalanUpdate struct {
	PenilaianAwalKeperawatanRalanData
	PenilaianAwalKeperawatanRalanDetail
}

func (r *PenilaianAwalKeperawatanRalanUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanRalanUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianAwalKeperawatanRalanRules()
}

func (r *PenilaianAwalKeperawatanRalanUpdate) Payload() PenilaianAwalKeperawatanRalanData {
	return r.PenilaianAwalKeperawatanRalanData
}

func (r *PenilaianAwalKeperawatanRalanUpdate) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}
