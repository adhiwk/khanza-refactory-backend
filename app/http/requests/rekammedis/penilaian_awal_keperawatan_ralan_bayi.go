package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianAwalKeperawatanRalanBayiData isian penilaian awal keperawatan bayi anak.
type PenilaianAwalKeperawatanRalanBayiData struct {
	Tanggal           string `form:"tanggal" json:"tanggal"`
	Informasi         string `form:"informasi" json:"informasi"`
	Td                string `form:"td" json:"td"`
	Nadi              string `form:"nadi" json:"nadi"`
	Rr                string `form:"rr" json:"rr"`
	Suhu              string `form:"suhu" json:"suhu"`
	Gcs               string `form:"gcs" json:"gcs"`
	Bb                string `form:"bb" json:"bb"`
	Tb                string `form:"tb" json:"tb"`
	Lp                string `form:"lp" json:"lp"`
	Lk                string `form:"lk" json:"lk"`
	Ld                string `form:"ld" json:"ld"`
	KeluhanUtama      string `form:"keluhan_utama" json:"keluhan_utama"`
	Rpd               string `form:"rpd" json:"rpd"`
	Rpk               string `form:"rpk" json:"rpk"`
	Rpo               string `form:"rpo" json:"rpo"`
	Alergi            string `form:"alergi" json:"alergi"`
	Anakke            string `form:"anakke" json:"anakke"`
	Darisaudara       string `form:"darisaudara" json:"darisaudara"`
	Caralahir         string `form:"caralahir" json:"caralahir"`
	KetCaralahir      string `form:"ket_caralahir" json:"ket_caralahir"`
	Umurkelahiran     string `form:"umurkelahiran" json:"umurkelahiran"`
	Kelainanbawaan    string `form:"kelainanbawaan" json:"kelainanbawaan"`
	KetKelainanBawaan string `form:"ket_kelainan_bawaan" json:"ket_kelainan_bawaan"`
	Usiatengkurap     string `form:"usiatengkurap" json:"usiatengkurap"`
	Usiaduduk         string `form:"usiaduduk" json:"usiaduduk"`
	Usiaberdiri       string `form:"usiaberdiri" json:"usiaberdiri"`
	Usiagigipertama   string `form:"usiagigipertama" json:"usiagigipertama"`
	Usiaberjalan      string `form:"usiaberjalan" json:"usiaberjalan"`
	Usiabicara        string `form:"usiabicara" json:"usiabicara"`
	Usiamembaca       string `form:"usiamembaca" json:"usiamembaca"`
	Usiamenulis       string `form:"usiamenulis" json:"usiamenulis"`
	Gangguanemosi     string `form:"gangguanemosi" json:"gangguanemosi"`
	AlatBantu         string `form:"alat_bantu" json:"alat_bantu"`
	KetBantu          string `form:"ket_bantu" json:"ket_bantu"`
	Prothesa          string `form:"prothesa" json:"prothesa"`
	KetPro            string `form:"ket_pro" json:"ket_pro"`
	Adl               string `form:"adl" json:"adl"`
	StatusPsiko       string `form:"status_psiko" json:"status_psiko"`
	KetPsiko          string `form:"ket_psiko" json:"ket_psiko"`
	HubKeluarga       string `form:"hub_keluarga" json:"hub_keluarga"`
	Pengasuh          string `form:"pengasuh" json:"pengasuh"`
	KetPengasuh       string `form:"ket_pengasuh" json:"ket_pengasuh"`
	Ekonomi           string `form:"ekonomi" json:"ekonomi"`
	Budaya            string `form:"budaya" json:"budaya"`
	KetBudaya         string `form:"ket_budaya" json:"ket_budaya"`
	Edukasi           string `form:"edukasi" json:"edukasi"`
	KetEdukasi        string `form:"ket_edukasi" json:"ket_edukasi"`
	BerjalanA         string `form:"berjalan_a" json:"berjalan_a"`
	BerjalanB         string `form:"berjalan_b" json:"berjalan_b"`
	BerjalanC         string `form:"berjalan_c" json:"berjalan_c"`
	Hasil             string `form:"hasil" json:"hasil"`
	Lapor             string `form:"lapor" json:"lapor"`
	KetLapor          string `form:"ket_lapor" json:"ket_lapor"`
	Sg1               string `form:"sg1" json:"sg1"`
	Nilai1            string `form:"nilai1" json:"nilai1"`
	Sg2               string `form:"sg2" json:"sg2"`
	Nilai2            string `form:"nilai2" json:"nilai2"`
	Sg3               string `form:"sg3" json:"sg3"`
	Nilai3            string `form:"nilai3" json:"nilai3"`
	Sg4               string `form:"sg4" json:"sg4"`
	Nilai4            string `form:"nilai4" json:"nilai4"`
	TotalHasil        int    `form:"total_hasil" json:"total_hasil"`
	Wajah             string `form:"wajah" json:"wajah"`
	Nilaiwajah        string `form:"nilaiwajah" json:"nilaiwajah"`
	Kaki              string `form:"kaki" json:"kaki"`
	Nilaikaki         string `form:"nilaikaki" json:"nilaikaki"`
	Aktifitas         string `form:"aktifitas" json:"aktifitas"`
	Nilaiaktifitas    string `form:"nilaiaktifitas" json:"nilaiaktifitas"`
	Menangis          string `form:"menangis" json:"menangis"`
	Nilaimenangis     string `form:"nilaimenangis" json:"nilaimenangis"`
	Bersuara          string `form:"bersuara" json:"bersuara"`
	Nilaibersuara     string `form:"nilaibersuara" json:"nilaibersuara"`
	Hasilnyeri        int    `form:"hasilnyeri" json:"hasilnyeri"`
	Nyeri             string `form:"nyeri" json:"nyeri"`
	Lokasi            string `form:"lokasi" json:"lokasi"`
	Durasi            string `form:"durasi" json:"durasi"`
	Frekuensi         string `form:"frekuensi" json:"frekuensi"`
	NyeriHilang       string `form:"nyeri_hilang" json:"nyeri_hilang"`
	KetNyeri          string `form:"ket_nyeri" json:"ket_nyeri"`
	PadaDokter        string `form:"pada_dokter" json:"pada_dokter"`
	KetDokter         string `form:"ket_dokter" json:"ket_dokter"`
	Rencana           string `form:"rencana" json:"rencana"`
	Nip               string `form:"nip" json:"nip"`
}

// PenilaianAwalKeperawatanRalanBayiDetail daftar kode detail (diganti seluruhnya saat simpan/ubah).
type PenilaianAwalKeperawatanRalanBayiDetail struct {
	KodeMasalah []string `form:"kode_masalah" json:"kode_masalah"`
	KodeRencana []string `form:"kode_rencana" json:"kode_rencana"`
}

func penilaianAwalKeperawatanRalanBayiRules() map[string]any {
	rules := map[string]any{
		"tanggal":             "required|date",
		"informasi":           "required|in:Autoanamnesis,Alloanamnesis",
		"td":                  "string|max_len:8",
		"nadi":                "string|max_len:5",
		"rr":                  "string|max_len:5",
		"suhu":                "string|max_len:5",
		"gcs":                 "string|max_len:5",
		"bb":                  "string|max_len:5",
		"tb":                  "string|max_len:5",
		"lp":                  "string|max_len:5",
		"lk":                  "string|max_len:5",
		"ld":                  "string|max_len:5",
		"keluhan_utama":       "string|max_len:150",
		"rpd":                 "string|max_len:100",
		"rpk":                 "string|max_len:100",
		"rpo":                 "string|max_len:100",
		"alergi":              "string|max_len:25",
		"anakke":              "string|max_len:4",
		"darisaudara":         "string|max_len:4",
		"caralahir":           "required|in:Spontan,Sectio Caesaria,Lain-Lain",
		"ket_caralahir":       "string|max_len:30",
		"umurkelahiran":       "required|in:Cukup Bulan,Kurang Bulan",
		"kelainanbawaan":      "required|in:Tidak Ada,Ada",
		"ket_kelainan_bawaan": "string|max_len:30",
		"usiatengkurap":       "string|max_len:15",
		"usiaduduk":           "string|max_len:15",
		"usiaberdiri":         "string|max_len:15",
		"usiagigipertama":     "string|max_len:15",
		"usiaberjalan":        "string|max_len:15",
		"usiabicara":          "string|max_len:15",
		"usiamembaca":         "string|max_len:15",
		"usiamenulis":         "string|max_len:15",
		"gangguanemosi":       "string|max_len:50",
		"alat_bantu":          "required|in:Tidak,Ya",
		"ket_bantu":           "string|max_len:50",
		"prothesa":            "required|in:Tidak,Ya",
		"ket_pro":             "string|max_len:50",
		"adl":                 "required|in:Mandiri,Dibantu",
		"status_psiko":        "required|in:Tenang,Takut,Tempertantrum,Cemas,Depresi,Lain-lain",
		"ket_psiko":           "string|max_len:70",
		"hub_keluarga":        "required|in:Baik,Tidak Baik",
		"pengasuh":            "required|in:Orang Tua,Kakek/Nenek,Keluarga Lainnya",
		"ket_pengasuh":        "string|max_len:40",
		"ekonomi":             "required|in:Baik,Cukup,Kurang",
		"budaya":              "required|in:Tidak Ada,Ada",
		"ket_budaya":          "string|max_len:50",
		"edukasi":             "required|in:Orang Tua,Keluarga",
		"ket_edukasi":         "string|max_len:50",
		"berjalan_a":          "required|in:Ya,Tidak",
		"berjalan_b":          "required|in:Ya,Tidak",
		"berjalan_c":          "required|in:Ya,Tidak",
		"hasil":               "required|in:Tidak beresiko (tidak ditemukan a dan b),Resiko rendah (ditemukan a/b),Resiko tinggi (ditemukan a dan b)",
		"lapor":               "required|in:Ya,Tidak",
		"ket_lapor":           "string|max_len:15",
		"sg1":                 "required|in:Tidak,Ya",
		"nilai1":              "required|in:0,1",
		"sg2":                 "required|in:Tidak,Ya",
		"nilai2":              "required|in:0,1",
		"sg3":                 "required|in:Tidak,Ya",
		"nilai3":              "required|in:0,1",
		"sg4":                 "required|in:Tidak,Ya",
		"nilai4":              "required|in:0,1",
		"total_hasil":         "int",
		"wajah":               "required|in:Tersenyum/tidak ada ekspresi khusus,Terkadang meringis/menarik diri,Sering menggetarkan dagu dan mengatupkan rahang",
		"nilaiwajah":          "required|in:0,1,2",
		"kaki":                "required|in:Gerakan normal/relaksasi,Tidak tenang/tegang,Kaki dibuat menendang/menarik",
		"nilaikaki":           "required|in:0,1,2",
		"aktifitas":           "required|string",
		"nilaiaktifitas":      "required|in:0,1,2",
		"menangis":            "required|string",
		"nilaimenangis":       "required|in:0,1,2",
		"bersuara":            "required|string",
		"nilaibersuara":       "required|in:0,1,2",
		"hasilnyeri":          "int",
		"nyeri":               "required|in:Tidak Ada Nyeri,Nyeri Akut,Nyeri Kronis",
		"lokasi":              "string|max_len:50",
		"durasi":              "string|max_len:25",
		"frekuensi":           "string|max_len:25",
		"nyeri_hilang":        "required|in:Minum Obat,Istirahat,Mendengar Music,Berubah Posisi/Tidur,Lain-lain",
		"ket_nyeri":           "string|max_len:40",
		"pada_dokter":         "required|in:Tidak,Ya",
		"ket_dokter":          "string|max_len:15",
		"rencana":             "string|max_len:200",
		"nip":                 "required|string|max_len:20",
	}
	rules["kode_masalah"] = "slice"
	rules["kode_masalah.*"] = "string|max_len:20"
	rules["kode_rencana"] = "slice"
	rules["kode_rencana.*"] = "string|max_len:20"
	return rules
}

// PenilaianAwalKeperawatanRalanBayiStore simpan penilaian awal keperawatan bayi anak; kolom waktu kunci kosong = sekarang.
type PenilaianAwalKeperawatanRalanBayiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianAwalKeperawatanRalanBayiData
	PenilaianAwalKeperawatanRalanBayiDetail
}

func (r *PenilaianAwalKeperawatanRalanBayiStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanRalanBayiStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianAwalKeperawatanRalanBayiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianAwalKeperawatanRalanBayiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianAwalKeperawatanRalanBayiStore) Payload() PenilaianAwalKeperawatanRalanBayiData {
	return r.PenilaianAwalKeperawatanRalanBayiData
}

func (r *PenilaianAwalKeperawatanRalanBayiStore) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}

// PenilaianAwalKeperawatanRalanBayiUpdate ubah penilaian awal keperawatan bayi anak (PUT); kunci lewat query string.
type PenilaianAwalKeperawatanRalanBayiUpdate struct {
	PenilaianAwalKeperawatanRalanBayiData
	PenilaianAwalKeperawatanRalanBayiDetail
}

func (r *PenilaianAwalKeperawatanRalanBayiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanRalanBayiUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianAwalKeperawatanRalanBayiRules()
}

func (r *PenilaianAwalKeperawatanRalanBayiUpdate) Payload() PenilaianAwalKeperawatanRalanBayiData {
	return r.PenilaianAwalKeperawatanRalanBayiData
}

func (r *PenilaianAwalKeperawatanRalanBayiUpdate) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}
