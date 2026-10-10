package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningKesehatanPenglihatanData isian skrining kesehatan penglihatan.
type SkriningKesehatanPenglihatanData struct {
	Tanggal        string `form:"tanggal" json:"tanggal"`
	MataLuar       string `form:"mata_luar" json:"mata_luar"`
	TajamKiri      string `form:"tajam_kiri" json:"tajam_kiri"`
	TajamKanan     string `form:"tajam_kanan" json:"tajam_kanan"`
	ButaWarnaKiri  string `form:"buta_warna_kiri" json:"buta_warna_kiri"`
	ButaWarnaKanan string `form:"buta_warna_kanan" json:"buta_warna_kanan"`
	Kacamata       string `form:"kacamata" json:"kacamata"`
	VisusKiri      string `form:"visus_kiri" json:"visus_kiri"`
	VisusKanan     string `form:"visus_kanan" json:"visus_kanan"`
	RefraksiKiri   string `form:"refraksi_kiri" json:"refraksi_kiri"`
	RefraksiKanan  string `form:"refraksi_kanan" json:"refraksi_kanan"`
	RujukRefraksi  string `form:"rujuk_refraksi" json:"rujuk_refraksi"`
	KatarakKiri    string `form:"katarak_kiri" json:"katarak_kiri"`
	KatarakKanan   string `form:"katarak_kanan" json:"katarak_kanan"`
	RujukKatarak   string `form:"rujuk_katarak" json:"rujuk_katarak"`
	HasilSkrining  string `form:"hasil_skrining" json:"hasil_skrining"`
	Keterangan     string `form:"keterangan" json:"keterangan"`
	Nip            string `form:"nip" json:"nip"`
}

func skriningKesehatanPenglihatanRules() map[string]any {
	rules := map[string]any{
		"tanggal":          "required|date",
		"mata_luar":        "in:Normal,Tidak Sehat",
		"tajam_kiri":       "in:Normal (6/6 - 6/18),Kelainan Refraksi (< 6/18 - 6/60),Low Vision (6/60 - 3/60),Kebutaan (< 3/60)",
		"tajam_kanan":      "in:Normal (6/6 - 6/18),Kelainan Refraksi (< 6/18 - 6/60),Low Vision (6/60 - 3/60),Kebutaan (< 3/60)",
		"buta_warna_kiri":  "in:Tidak,Ya",
		"buta_warna_kanan": "in:Tidak,Ya",
		"kacamata":         "in:Tidak,Ya",
		"visus_kiri":       "in:Normal (6/6 - 6/18),Kelainan Refraksi (< 6/18 - 6/60),Low Vision (6/60 - 3/60),Kebutaan (< 3/60)",
		"visus_kanan":      "in:Normal (6/6 - 6/18),Kelainan Refraksi (< 6/18 - 6/60),Low Vision (6/60 - 3/60),Kebutaan (< 3/60)",
		"refraksi_kiri":    "in:Tidak,Ya",
		"refraksi_kanan":   "in:Tidak,Ya",
		"rujuk_refraksi":   "in:Tidak,Ya",
		"katarak_kiri":     "in:Tidak,Ya",
		"katarak_kanan":    "in:Tidak,Ya",
		"rujuk_katarak":    "in:Tidak,Ya",
		"hasil_skrining":   "string|max_len:40",
		"keterangan":       "string|max_len:100",
		"nip":              "required|string|max_len:20",
	}
	return rules
}

// SkriningKesehatanPenglihatanStore simpan skrining kesehatan penglihatan; kolom waktu kunci kosong = sekarang.
type SkriningKesehatanPenglihatanStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningKesehatanPenglihatanData
}

func (r *SkriningKesehatanPenglihatanStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningKesehatanPenglihatanStore) Rules(ctx http.Context) map[string]any {
	rules := skriningKesehatanPenglihatanRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningKesehatanPenglihatanStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *SkriningKesehatanPenglihatanStore) Payload() SkriningKesehatanPenglihatanData {
	return r.SkriningKesehatanPenglihatanData
}

func (r *SkriningKesehatanPenglihatanStore) DetailValues() map[string][]string { return nil }

// SkriningKesehatanPenglihatanUpdate ubah skrining kesehatan penglihatan (PUT); kunci lewat query string.
type SkriningKesehatanPenglihatanUpdate struct {
	SkriningKesehatanPenglihatanData
}

func (r *SkriningKesehatanPenglihatanUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningKesehatanPenglihatanUpdate) Rules(ctx http.Context) map[string]any {
	return skriningKesehatanPenglihatanRules()
}

func (r *SkriningKesehatanPenglihatanUpdate) Payload() SkriningKesehatanPenglihatanData {
	return r.SkriningKesehatanPenglihatanData
}

func (r *SkriningKesehatanPenglihatanUpdate) DetailValues() map[string][]string { return nil }
