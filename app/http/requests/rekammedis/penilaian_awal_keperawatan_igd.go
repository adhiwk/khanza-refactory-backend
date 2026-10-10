package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianAwalKeperawatanIgdData isian penilaian awal keperawatan IGD.
type PenilaianAwalKeperawatanIgdData struct {
	Tanggal          string `form:"tanggal" json:"tanggal"`
	Informasi        string `form:"informasi" json:"informasi"`
	KeluhanUtama     string `form:"keluhan_utama" json:"keluhan_utama"`
	Rpd              string `form:"rpd" json:"rpd"`
	Rpo              string `form:"rpo" json:"rpo"`
	StatusKehamilan  string `form:"status_kehamilan" json:"status_kehamilan"`
	Gravida          string `form:"gravida" json:"gravida"`
	Para             string `form:"para" json:"para"`
	Abortus          string `form:"abortus" json:"abortus"`
	Hpht             string `form:"hpht" json:"hpht"`
	Tekanan          string `form:"tekanan" json:"tekanan"`
	Pupil            string `form:"pupil" json:"pupil"`
	Neurosensorik    string `form:"neurosensorik" json:"neurosensorik"`
	Integumen        string `form:"integumen" json:"integumen"`
	Turgor           string `form:"turgor" json:"turgor"`
	Edema            string `form:"edema" json:"edema"`
	Mukosa           string `form:"mukosa" json:"mukosa"`
	Perdarahan       string `form:"perdarahan" json:"perdarahan"`
	JumlahPerdarahan string `form:"jumlah_perdarahan" json:"jumlah_perdarahan"`
	WarnaPerdarahan  string `form:"warna_perdarahan" json:"warna_perdarahan"`
	Intoksikasi      string `form:"intoksikasi" json:"intoksikasi"`
	Bab              string `form:"bab" json:"bab"`
	Xbab             string `form:"xbab" json:"xbab"`
	Kbab             string `form:"kbab" json:"kbab"`
	Wbab             string `form:"wbab" json:"wbab"`
	Bak              string `form:"bak" json:"bak"`
	Xbak             string `form:"xbak" json:"xbak"`
	Wbak             string `form:"wbak" json:"wbak"`
	Lbak             string `form:"lbak" json:"lbak"`
	Psikologis       string `form:"psikologis" json:"psikologis"`
	Jiwa             string `form:"jiwa" json:"jiwa"`
	Perilaku         string `form:"perilaku" json:"perilaku"`
	Dilaporkan       string `form:"dilaporkan" json:"dilaporkan"`
	Sebutkan         string `form:"sebutkan" json:"sebutkan"`
	Hubungan         string `form:"hubungan" json:"hubungan"`
	TinggalDengan    string `form:"tinggal_dengan" json:"tinggal_dengan"`
	KetTinggal       string `form:"ket_tinggal" json:"ket_tinggal"`
	Budaya           string `form:"budaya" json:"budaya"`
	KetBudaya        string `form:"ket_budaya" json:"ket_budaya"`
	PendidikanPj     string `form:"pendidikan_pj" json:"pendidikan_pj"`
	KetPendidikanPj  string `form:"ket_pendidikan_pj" json:"ket_pendidikan_pj"`
	Edukasi          string `form:"edukasi" json:"edukasi"`
	KetEdukasi       string `form:"ket_edukasi" json:"ket_edukasi"`
	Kemampuan        string `form:"kemampuan" json:"kemampuan"`
	Aktifitas        string `form:"aktifitas" json:"aktifitas"`
	AlatBantu        string `form:"alat_bantu" json:"alat_bantu"`
	KetBantu         string `form:"ket_bantu" json:"ket_bantu"`
	Nyeri            string `form:"nyeri" json:"nyeri"`
	Provokes         string `form:"provokes" json:"provokes"`
	KetProvokes      string `form:"ket_provokes" json:"ket_provokes"`
	Quality          string `form:"quality" json:"quality"`
	KetQuality       string `form:"ket_quality" json:"ket_quality"`
	Lokasi           string `form:"lokasi" json:"lokasi"`
	Menyebar         string `form:"menyebar" json:"menyebar"`
	SkalaNyeri       string `form:"skala_nyeri" json:"skala_nyeri"`
	Durasi           string `form:"durasi" json:"durasi"`
	NyeriHilang      string `form:"nyeri_hilang" json:"nyeri_hilang"`
	KetNyeri         string `form:"ket_nyeri" json:"ket_nyeri"`
	PadaDokter       string `form:"pada_dokter" json:"pada_dokter"`
	KetDokter        string `form:"ket_dokter" json:"ket_dokter"`
	BerjalanA        string `form:"berjalan_a" json:"berjalan_a"`
	BerjalanB        string `form:"berjalan_b" json:"berjalan_b"`
	BerjalanC        string `form:"berjalan_c" json:"berjalan_c"`
	Hasil            string `form:"hasil" json:"hasil"`
	Lapor            string `form:"lapor" json:"lapor"`
	KetLapor         string `form:"ket_lapor" json:"ket_lapor"`
	Rencana          string `form:"rencana" json:"rencana"`
	Nip              string `form:"nip" json:"nip"`
}

// PenilaianAwalKeperawatanIgdDetail daftar kode detail (diganti seluruhnya saat simpan/ubah).
type PenilaianAwalKeperawatanIgdDetail struct {
	KodeMasalah []string `form:"kode_masalah" json:"kode_masalah"`
	KodeRencana []string `form:"kode_rencana" json:"kode_rencana"`
}

func penilaianAwalKeperawatanIgdRules() map[string]any {
	rules := map[string]any{
		"tanggal":           "required|date",
		"informasi":         "required|in:Autoanamnesis,Alloanamnesis",
		"keluhan_utama":     "string",
		"rpd":               "string",
		"rpo":               "string",
		"status_kehamilan":  "required|in:Tidak Hamil,Hamil",
		"gravida":           "string|max_len:20",
		"para":              "string|max_len:20",
		"abortus":           "string|max_len:20",
		"hpht":              "string|max_len:20",
		"tekanan":           "required|in:TAK,Sakit Kepala,Muntah,Pusing,Bingung",
		"pupil":             "required|in:Normal,Miosis,Isokor,Anisokor",
		"neurosensorik":     "required|in:TAK,Spasme Otot,Perubahan Sensorik,Perubahan Motorik,Perubahan Bentuk Ekstremitas,Penurunan Tingkat Kesadaran,Fraktur/Dislokasi,Luksasio,Kerusakan Jaringan/Luka",
		"integumen":         "required|in:TAK,Luka Bakar,Luka Robek,Lecet,Luka Decubitus,Luka Gangren",
		"turgor":            "required|in:Baik,Menurun",
		"edema":             "required|in:Tidak Ada,Ekstremitas,Seluruh Tubuh,Asites,Palpebrae",
		"mukosa":            "required|in:Lembab,Kering",
		"perdarahan":        "required|in:Tidak Ada,Ada",
		"jumlah_perdarahan": "string|max_len:5",
		"warna_perdarahan":  "string|max_len:40",
		"intoksikasi":       "required|in:Tidak Ada,Ada,Gigitan Binatang,Zat Kimia,Gas,Obat",
		"bab":               "string|max_len:2",
		"xbab":              "string|max_len:10",
		"kbab":              "string|max_len:40",
		"wbab":              "string|max_len:40",
		"bak":               "string|max_len:2",
		"xbak":              "string|max_len:10",
		"wbak":              "string|max_len:40",
		"lbak":              "string|max_len:40",
		"psikologis":        "required|in:Tidak Ada Masalah,Marah,Takut,Depresi,Cepat Lelah,Cemas,Gelisah,Lain-lain",
		"jiwa":              "required|in:Ya,Tidak",
		"perilaku":          "required|in:Perilaku Kekerasan,Gangguan Efek,Gangguan Memori,Halusinasi,Kecenderungan Percobaan Bunuh Diri,Lainnya,-",
		"dilaporkan":        "string|max_len:50",
		"sebutkan":          "string|max_len:50",
		"hubungan":          "required|in:Harmonis,Kurang Harmonis,Tidak Harmonis,Konflik Besar",
		"tinggal_dengan":    "required|in:Sendiri,Orang Tua,Suami / Istri,Lainnya",
		"ket_tinggal":       "string|max_len:50",
		"budaya":            "required|in:Tidak Ada,Ada",
		"ket_budaya":        "string|max_len:50",
		"pendidikan_pj":     "required|in:-,TS,TK,SD,SMP,SMA,SLTA/SEDERAJAT,D1,D2,D3,D4,S1,S2,S3",
		"ket_pendidikan_pj": "string|max_len:50",
		"edukasi":           "required|in:Pasien,Keluarga",
		"ket_edukasi":       "string|max_len:50",
		"kemampuan":         "required|in:Mandiri,Bantuan Minimal,Bantuan Sebagian,Ketergantungan Total",
		"aktifitas":         "required|in:Tirah Baring,Duduk,Berjalan",
		"alat_bantu":        "required|in:Tidak,Ya",
		"ket_bantu":         "string|max_len:50",
		"nyeri":             "required|in:Tidak Ada Nyeri,Nyeri Akut,Nyeri Kronis",
		"provokes":          "required|in:Proses Penyakit,Benturan,Lain-lain",
		"ket_provokes":      "string|max_len:40",
		"quality":           "required|in:Seperti Tertusuk,Berdenyut,Teriris,Tertindih,Tertiban,Lain-lain",
		"ket_quality":       "string|max_len:50",
		"lokasi":            "string|max_len:50",
		"menyebar":          "required|in:Tidak,Ya",
		"skala_nyeri":       "required|in:0,1,2,3,4,5,6,7,8,9,10",
		"durasi":            "string|max_len:25",
		"nyeri_hilang":      "required|in:Istirahat,Medengar Musik,Minum Obat",
		"ket_nyeri":         "string|max_len:40",
		"pada_dokter":       "required|in:Tidak,Ya",
		"ket_dokter":        "string|max_len:15",
		"berjalan_a":        "required|in:Ya,Tidak",
		"berjalan_b":        "required|in:Ya,Tidak",
		"berjalan_c":        "required|in:Ya,Tidak",
		"hasil":             "required|in:Tidak beresiko (tidak ditemukan a dan b),Resiko rendah (ditemukan a/b),Resiko tinggi (ditemukan a dan b)",
		"lapor":             "required|in:Ya,Tidak",
		"ket_lapor":         "string|max_len:15",
		"rencana":           "string",
		"nip":               "required|string|max_len:20",
	}
	rules["kode_masalah"] = "slice"
	rules["kode_masalah.*"] = "string|max_len:20"
	rules["kode_rencana"] = "slice"
	rules["kode_rencana.*"] = "string|max_len:20"
	return rules
}

// PenilaianAwalKeperawatanIgdStore simpan penilaian awal keperawatan IGD; kolom waktu kunci kosong = sekarang.
type PenilaianAwalKeperawatanIgdStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianAwalKeperawatanIgdData
	PenilaianAwalKeperawatanIgdDetail
}

func (r *PenilaianAwalKeperawatanIgdStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanIgdStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianAwalKeperawatanIgdRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianAwalKeperawatanIgdStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *PenilaianAwalKeperawatanIgdStore) Payload() PenilaianAwalKeperawatanIgdData {
	return r.PenilaianAwalKeperawatanIgdData
}

func (r *PenilaianAwalKeperawatanIgdStore) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}

// PenilaianAwalKeperawatanIgdUpdate ubah penilaian awal keperawatan IGD (PUT); kunci lewat query string.
type PenilaianAwalKeperawatanIgdUpdate struct {
	PenilaianAwalKeperawatanIgdData
	PenilaianAwalKeperawatanIgdDetail
}

func (r *PenilaianAwalKeperawatanIgdUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianAwalKeperawatanIgdUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianAwalKeperawatanIgdRules()
}

func (r *PenilaianAwalKeperawatanIgdUpdate) Payload() PenilaianAwalKeperawatanIgdData {
	return r.PenilaianAwalKeperawatanIgdData
}

func (r *PenilaianAwalKeperawatanIgdUpdate) DetailValues() map[string][]string {
	return map[string][]string{"kode_masalah": r.KodeMasalah, "kode_rencana": r.KodeRencana}
}
