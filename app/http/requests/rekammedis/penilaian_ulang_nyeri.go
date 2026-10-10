package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianUlangNyeriData isian penilaian ulang nyeri.
type PenilaianUlangNyeriData struct {
	Nyeri       string `form:"nyeri" json:"nyeri"`
	Provokes    string `form:"provokes" json:"provokes"`
	KetProvokes string `form:"ket_provokes" json:"ket_provokes"`
	Quality     string `form:"quality" json:"quality"`
	KetQuality  string `form:"ket_quality" json:"ket_quality"`
	Lokasi      string `form:"lokasi" json:"lokasi"`
	Menyebar    string `form:"menyebar" json:"menyebar"`
	SkalaNyeri  string `form:"skala_nyeri" json:"skala_nyeri"`
	Durasi      string `form:"durasi" json:"durasi"`
	NyeriHilang string `form:"nyeri_hilang" json:"nyeri_hilang"`
	KetNyeri    string `form:"ket_nyeri" json:"ket_nyeri"`
	Nip         string `form:"nip" json:"nip"`
}

func penilaianUlangNyeriRules() map[string]any {
	rules := map[string]any{
		"nyeri":        "required|in:Tidak Ada Nyeri,Nyeri Akut,Nyeri Kronis",
		"provokes":     "required|in:Proses Penyakit,Benturan,Lain-lain,-",
		"ket_provokes": "string|max_len:40",
		"quality":      "required|in:Seperti Tertusuk,Berdenyut,Teriris,Tertindih,Tertiban,Lain-lain,-",
		"ket_quality":  "string|max_len:50",
		"lokasi":       "string|max_len:50",
		"menyebar":     "required|in:Tidak,Ya",
		"skala_nyeri":  "required|in:0,1,2,3,4,5,6,7,8,9,10",
		"durasi":       "string|max_len:25",
		"nyeri_hilang": "required|in:Istirahat,Medengar Musik,Minum Obat,-",
		"ket_nyeri":    "string|max_len:40",
		"nip":          "required|string|max_len:20",
	}
	return rules
}

// PenilaianUlangNyeriStore simpan penilaian ulang nyeri; kolom waktu kunci kosong = sekarang.
type PenilaianUlangNyeriStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PenilaianUlangNyeriData
}

func (r *PenilaianUlangNyeriStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianUlangNyeriStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianUlangNyeriRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PenilaianUlangNyeriStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PenilaianUlangNyeriStore) Payload() PenilaianUlangNyeriData {
	return r.PenilaianUlangNyeriData
}

func (r *PenilaianUlangNyeriStore) DetailValues() map[string][]string { return nil }

// PenilaianUlangNyeriUpdate ubah penilaian ulang nyeri (PUT); kunci lewat query string.
type PenilaianUlangNyeriUpdate struct {
	PenilaianUlangNyeriData
}

func (r *PenilaianUlangNyeriUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianUlangNyeriUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianUlangNyeriRules()
}

func (r *PenilaianUlangNyeriUpdate) Payload() PenilaianUlangNyeriData {
	return r.PenilaianUlangNyeriData
}

func (r *PenilaianUlangNyeriUpdate) DetailValues() map[string][]string { return nil }
