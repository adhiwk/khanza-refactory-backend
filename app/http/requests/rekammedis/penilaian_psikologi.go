package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianPsikologiData isian penilaian psikologi.
type PenilaianPsikologiData struct {
	Tanggal            string `form:"tanggal" json:"tanggal"`
	Nip                string `form:"nip" json:"nip"`
	Anamnesis          string `form:"anamnesis" json:"anamnesis"`
	DikirimDari        string `form:"dikirim_dari" json:"dikirim_dari"`
	TujuanPemeriksaan  string `form:"tujuan_pemeriksaan" json:"tujuan_pemeriksaan"`
	KetAnamnesis       string `form:"ket_anamnesis" json:"ket_anamnesis"`
	Rupa               string `form:"rupa" json:"rupa"`
	BentukTubuh        string `form:"bentuk_tubuh" json:"bentuk_tubuh"`
	Tindakan           string `form:"tindakan" json:"tindakan"`
	Pakaian            string `form:"pakaian" json:"pakaian"`
	Ekspresi           string `form:"ekspresi" json:"ekspresi"`
	Berbicara          string `form:"berbicara" json:"berbicara"`
	PenggunaanKata     string `form:"penggunaan_kata" json:"penggunaan_kata"`
	CiriMenyolok       string `form:"ciri_menyolok" json:"ciri_menyolok"`
	HasilPsikotes      string `form:"hasil_psikotes" json:"hasil_psikotes"`
	Kepribadian        string `form:"kepribadian" json:"kepribadian"`
	Psikodinamika      string `form:"psikodinamika" json:"psikodinamika"`
	KesimpulanPsikolog string `form:"kesimpulan_psikolog" json:"kesimpulan_psikolog"`
}

func penilaianPsikologiRules() map[string]any {
	rules := map[string]any{
		"tanggal":             "required|date",
		"nip":                 "required|string|max_len:20",
		"anamnesis":           "required|in:Autoanamnesis,Alloanamnesis",
		"dikirim_dari":        "required|in:Ruang Rawat,Poliklinik,Rehabilitasi,After Care,Dokter",
		"tujuan_pemeriksaan":  "required|in:Klinik,Bimbingan,Forensik",
		"ket_anamnesis":       "string",
		"rupa":                "required|in:Tampan,Buruk,Menarik,Memuakkan,Biasa",
		"bentuk_tubuh":        "required|in:Sangat Tinggi,Sangat Pendek,Sangat Kurus,Sangat Gemuk,Tinggi,Sedang,Atletik,Pendek,Langsing,Gemuk",
		"tindakan":            "required|in:Sopan,Tidak Sopan,Kurang Tahu Aturan,Canggung,Bebas,Tegas,Garang,Percaya Diri,Tertekan,Ragu-Ragu,Pasti,Kaku,Ceroboh,Dingin,Malu-Malu",
		"pakaian":             "required|in:Rapi,Serampangan,Terpelihara,Tidak Terpelihara,Teratur,Tidak Rapi,Sederhana,Biasa,Bersih,Kotor",
		"ekspresi":            "required|in:Sangat Mudah,Hati-Hati Dan Membatasi Diri,Sukar Mencari Kata-Kata,Mudah,Terbuka",
		"berbicara":           "required|in:Tenang,Acuh Tak Acuh,Gugup,Lancar,Ribut Dengan Banyak Gerak dan Isyarat",
		"penggunaan_kata":     "required|in:Ramah,Dibuat-Buat,Dengan Tekanan Suara,Terpengaruh Bahasa Daerah,Disertai Dengan Istilah Bahasa Asing",
		"ciri_menyolok":       "string|max_len:500",
		"hasil_psikotes":      "string",
		"kepribadian":         "string",
		"psikodinamika":       "string",
		"kesimpulan_psikolog": "string",
	}
	return rules
}

// PenilaianPsikologiStore simpan penilaian psikologi; kolom waktu kunci kosong = sekarang.
type PenilaianPsikologiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianPsikologiData
}

func (r *PenilaianPsikologiStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPsikologiStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianPsikologiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianPsikologiStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianPsikologiStore) Payload() PenilaianPsikologiData { return r.PenilaianPsikologiData }

func (r *PenilaianPsikologiStore) DetailValues() map[string][]string { return nil }

// PenilaianPsikologiUpdate ubah penilaian psikologi (PUT); kunci lewat query string.
type PenilaianPsikologiUpdate struct {
	PenilaianPsikologiData
}

func (r *PenilaianPsikologiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPsikologiUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianPsikologiRules()
}

func (r *PenilaianPsikologiUpdate) Payload() PenilaianPsikologiData { return r.PenilaianPsikologiData }

func (r *PenilaianPsikologiUpdate) DetailValues() map[string][]string { return nil }
