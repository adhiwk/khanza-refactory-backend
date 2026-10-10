package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningRisikoKankerPayudaraData isian skrining risiko kanker payudara.
type SkriningRisikoKankerPayudaraData struct {
	Tanggal                string `form:"tanggal" json:"tanggal"`
	FaktorRisikoAwal1      string `form:"faktor_risiko_awal1" json:"faktor_risiko_awal1"`
	NilaiRisikoAwal1       string `form:"nilai_risiko_awal1" json:"nilai_risiko_awal1"`
	FaktorRisikoAwal2      string `form:"faktor_risiko_awal2" json:"faktor_risiko_awal2"`
	NilaiRisikoAwal2       string `form:"nilai_risiko_awal2" json:"nilai_risiko_awal2"`
	FaktorRisikoAwal3      string `form:"faktor_risiko_awal3" json:"faktor_risiko_awal3"`
	NilaiRisikoAwal3       string `form:"nilai_risiko_awal3" json:"nilai_risiko_awal3"`
	FaktorRisikoAwal4      string `form:"faktor_risiko_awal4" json:"faktor_risiko_awal4"`
	NilaiRisikoAwal4       string `form:"nilai_risiko_awal4" json:"nilai_risiko_awal4"`
	FaktorRisikoAwal5      string `form:"faktor_risiko_awal5" json:"faktor_risiko_awal5"`
	NilaiRisikoAwal5       string `form:"nilai_risiko_awal5" json:"nilai_risiko_awal5"`
	FaktorRisikoAwal6      string `form:"faktor_risiko_awal6" json:"faktor_risiko_awal6"`
	NilaiRisikoAwal6       string `form:"nilai_risiko_awal6" json:"nilai_risiko_awal6"`
	FaktorRisikoAwal7      string `form:"faktor_risiko_awal7" json:"faktor_risiko_awal7"`
	NilaiRisikoAwal7       string `form:"nilai_risiko_awal7" json:"nilai_risiko_awal7"`
	FaktorRisikoAwal8      string `form:"faktor_risiko_awal8" json:"faktor_risiko_awal8"`
	NilaiRisikoAwal8       string `form:"nilai_risiko_awal8" json:"nilai_risiko_awal8"`
	FaktorRisikoAwal9      string `form:"faktor_risiko_awal9" json:"faktor_risiko_awal9"`
	NilaiRisikoAwal9       string `form:"nilai_risiko_awal9" json:"nilai_risiko_awal9"`
	FaktorRisikoAwal10     string `form:"faktor_risiko_awal10" json:"faktor_risiko_awal10"`
	NilaiRisikoAwal10      string `form:"nilai_risiko_awal10" json:"nilai_risiko_awal10"`
	FaktorRisikoAwal11     string `form:"faktor_risiko_awal11" json:"faktor_risiko_awal11"`
	NilaiRisikoAwal11      string `form:"nilai_risiko_awal11" json:"nilai_risiko_awal11"`
	FaktorRisikoAwal12     string `form:"faktor_risiko_awal12" json:"faktor_risiko_awal12"`
	NilaiRisikoAwal12      string `form:"nilai_risiko_awal12" json:"nilai_risiko_awal12"`
	FaktorRisikoAwal13     string `form:"faktor_risiko_awal13" json:"faktor_risiko_awal13"`
	NilaiRisikoAwal13      string `form:"nilai_risiko_awal13" json:"nilai_risiko_awal13"`
	FaktorRisikoAwal14     string `form:"faktor_risiko_awal14" json:"faktor_risiko_awal14"`
	NilaiRisikoAwal14      string `form:"nilai_risiko_awal14" json:"nilai_risiko_awal14"`
	FaktorRisikoTinggi1    string `form:"faktor_risiko_tinggi1" json:"faktor_risiko_tinggi1"`
	NilaiRisikoTinggi1     string `form:"nilai_risiko_tinggi1" json:"nilai_risiko_tinggi1"`
	FaktorRisikoTinggi2    string `form:"faktor_risiko_tinggi2" json:"faktor_risiko_tinggi2"`
	NilaiRisikoTinggi2     string `form:"nilai_risiko_tinggi2" json:"nilai_risiko_tinggi2"`
	FaktorRisikoTinggi3    string `form:"faktor_risiko_tinggi3" json:"faktor_risiko_tinggi3"`
	NilaiRisikoTinggi3     string `form:"nilai_risiko_tinggi3" json:"nilai_risiko_tinggi3"`
	FaktorRisikoTinggi4    string `form:"faktor_risiko_tinggi4" json:"faktor_risiko_tinggi4"`
	NilaiRisikoTinggi4     string `form:"nilai_risiko_tinggi4" json:"nilai_risiko_tinggi4"`
	FaktorRisikoTinggi5    string `form:"faktor_risiko_tinggi5" json:"faktor_risiko_tinggi5"`
	NilaiRisikoTinggi5     string `form:"nilai_risiko_tinggi5" json:"nilai_risiko_tinggi5"`
	FaktorRisikoTinggi6    string `form:"faktor_risiko_tinggi6" json:"faktor_risiko_tinggi6"`
	NilaiRisikoTinggi6     string `form:"nilai_risiko_tinggi6" json:"nilai_risiko_tinggi6"`
	FaktorRisikoTinggi7    string `form:"faktor_risiko_tinggi7" json:"faktor_risiko_tinggi7"`
	NilaiRisikoTinggi7     string `form:"nilai_risiko_tinggi7" json:"nilai_risiko_tinggi7"`
	FaktorRisikoTinggi8    string `form:"faktor_risiko_tinggi8" json:"faktor_risiko_tinggi8"`
	NilaiRisikoTinggi8     string `form:"nilai_risiko_tinggi8" json:"nilai_risiko_tinggi8"`
	FaktorRisikoTinggi9    string `form:"faktor_risiko_tinggi9" json:"faktor_risiko_tinggi9"`
	NilaiRisikoTinggi9     string `form:"nilai_risiko_tinggi9" json:"nilai_risiko_tinggi9"`
	FaktorRisikoTinggi10   string `form:"faktor_risiko_tinggi10" json:"faktor_risiko_tinggi10"`
	NilaiRisikoTinggi10    string `form:"nilai_risiko_tinggi10" json:"nilai_risiko_tinggi10"`
	FaktorRisikoTinggi11   string `form:"faktor_risiko_tinggi11" json:"faktor_risiko_tinggi11"`
	NilaiRisikoTinggi11    string `form:"nilai_risiko_tinggi11" json:"nilai_risiko_tinggi11"`
	FaktorRisikoTinggi12   string `form:"faktor_risiko_tinggi12" json:"faktor_risiko_tinggi12"`
	NilaiRisikoTinggi12    string `form:"nilai_risiko_tinggi12" json:"nilai_risiko_tinggi12"`
	FaktorRisikoTinggi13   string `form:"faktor_risiko_tinggi13" json:"faktor_risiko_tinggi13"`
	NilaiRisikoTinggi13    string `form:"nilai_risiko_tinggi13" json:"nilai_risiko_tinggi13"`
	FaktorKecurigaanGanas1 string `form:"faktor_kecurigaan_ganas1" json:"faktor_kecurigaan_ganas1"`
	NilaiKecurigaanGanas1  string `form:"nilai_kecurigaan_ganas1" json:"nilai_kecurigaan_ganas1"`
	FaktorKecurigaanGanas2 string `form:"faktor_kecurigaan_ganas2" json:"faktor_kecurigaan_ganas2"`
	NilaiKecurigaanGanas2  string `form:"nilai_kecurigaan_ganas2" json:"nilai_kecurigaan_ganas2"`
	FaktorKecurigaanGanas3 string `form:"faktor_kecurigaan_ganas3" json:"faktor_kecurigaan_ganas3"`
	NilaiKecurigaanGanas3  string `form:"nilai_kecurigaan_ganas3" json:"nilai_kecurigaan_ganas3"`
	FaktorKecurigaanGanas4 string `form:"faktor_kecurigaan_ganas4" json:"faktor_kecurigaan_ganas4"`
	NilaiKecurigaanGanas4  string `form:"nilai_kecurigaan_ganas4" json:"nilai_kecurigaan_ganas4"`
	FaktorKecurigaanGanas5 string `form:"faktor_kecurigaan_ganas5" json:"faktor_kecurigaan_ganas5"`
	NilaiKecurigaanGanas5  string `form:"nilai_kecurigaan_ganas5" json:"nilai_kecurigaan_ganas5"`
	FaktorKecurigaanGanas6 string `form:"faktor_kecurigaan_ganas6" json:"faktor_kecurigaan_ganas6"`
	NilaiKecurigaanGanas6  string `form:"nilai_kecurigaan_ganas6" json:"nilai_kecurigaan_ganas6"`
	FaktorKecurigaanGanas7 string `form:"faktor_kecurigaan_ganas7" json:"faktor_kecurigaan_ganas7"`
	NilaiKecurigaanGanas7  string `form:"nilai_kecurigaan_ganas7" json:"nilai_kecurigaan_ganas7"`
	FaktorKecurigaanGanas8 string `form:"faktor_kecurigaan_ganas8" json:"faktor_kecurigaan_ganas8"`
	NilaiKecurigaanGanas8  string `form:"nilai_kecurigaan_ganas8" json:"nilai_kecurigaan_ganas8"`
	TotalSkor              string `form:"total_skor" json:"total_skor"`
	HasilSadanis           string `form:"hasil_sadanis" json:"hasil_sadanis"`
	TindakLanjutSadanis    string `form:"tindak_lanjut_sadanis" json:"tindak_lanjut_sadanis"`
	HasilSkrining          string `form:"hasil_skrining" json:"hasil_skrining"`
	Keterangan             string `form:"keterangan" json:"keterangan"`
	Nip                    string `form:"nip" json:"nip"`
}

func skriningRisikoKankerPayudaraRules() map[string]any {
	rules := map[string]any{
		"tanggal":                  "required|date",
		"faktor_risiko_awal1":      "in:Ya,Tidak",
		"nilai_risiko_awal1":       "string|max_len:1",
		"faktor_risiko_awal2":      "in:Ya,Tidak",
		"nilai_risiko_awal2":       "string|max_len:1",
		"faktor_risiko_awal3":      "in:Ya,Tidak",
		"nilai_risiko_awal3":       "string|max_len:1",
		"faktor_risiko_awal4":      "in:Ya,Tidak",
		"nilai_risiko_awal4":       "string|max_len:1",
		"faktor_risiko_awal5":      "in:Ya,Tidak",
		"nilai_risiko_awal5":       "string|max_len:1",
		"faktor_risiko_awal6":      "in:Ya,Tidak",
		"nilai_risiko_awal6":       "string|max_len:1",
		"faktor_risiko_awal7":      "in:Ya,Tidak",
		"nilai_risiko_awal7":       "string|max_len:1",
		"faktor_risiko_awal8":      "in:Ya,Tidak",
		"nilai_risiko_awal8":       "string|max_len:1",
		"faktor_risiko_awal9":      "in:Ya,Tidak",
		"nilai_risiko_awal9":       "string|max_len:1",
		"faktor_risiko_awal10":     "in:Ya,Tidak",
		"nilai_risiko_awal10":      "string|max_len:1",
		"faktor_risiko_awal11":     "in:Ya,Tidak",
		"nilai_risiko_awal11":      "string|max_len:1",
		"faktor_risiko_awal12":     "in:Ya,Tidak",
		"nilai_risiko_awal12":      "string|max_len:1",
		"faktor_risiko_awal13":     "in:Ya,Tidak",
		"nilai_risiko_awal13":      "string|max_len:1",
		"faktor_risiko_awal14":     "in:Ya,Tidak",
		"nilai_risiko_awal14":      "string|max_len:1",
		"faktor_risiko_tinggi1":    "in:Ya,Tidak",
		"nilai_risiko_tinggi1":     "string|max_len:1",
		"faktor_risiko_tinggi2":    "in:Ya,Tidak",
		"nilai_risiko_tinggi2":     "string|max_len:1",
		"faktor_risiko_tinggi3":    "in:Ya,Tidak",
		"nilai_risiko_tinggi3":     "string|max_len:1",
		"faktor_risiko_tinggi4":    "in:Ya,Tidak",
		"nilai_risiko_tinggi4":     "string|max_len:1",
		"faktor_risiko_tinggi5":    "in:Ya,Tidak",
		"nilai_risiko_tinggi5":     "string|max_len:1",
		"faktor_risiko_tinggi6":    "in:Ya,Tidak",
		"nilai_risiko_tinggi6":     "string|max_len:1",
		"faktor_risiko_tinggi7":    "in:Ya,Tidak",
		"nilai_risiko_tinggi7":     "string|max_len:1",
		"faktor_risiko_tinggi8":    "in:Ya,Tidak",
		"nilai_risiko_tinggi8":     "string|max_len:1",
		"faktor_risiko_tinggi9":    "in:Ya,Tidak",
		"nilai_risiko_tinggi9":     "string|max_len:1",
		"faktor_risiko_tinggi10":   "in:Ya,Tidak",
		"nilai_risiko_tinggi10":    "string|max_len:1",
		"faktor_risiko_tinggi11":   "in:Ya,Tidak",
		"nilai_risiko_tinggi11":    "string|max_len:1",
		"faktor_risiko_tinggi12":   "in:Ya,Tidak",
		"nilai_risiko_tinggi12":    "string|max_len:1",
		"faktor_risiko_tinggi13":   "in:Ya,Tidak",
		"nilai_risiko_tinggi13":    "string|max_len:1",
		"faktor_kecurigaan_ganas1": "in:Ya,Tidak",
		"nilai_kecurigaan_ganas1":  "string|max_len:2",
		"faktor_kecurigaan_ganas2": "in:Ya,Tidak",
		"nilai_kecurigaan_ganas2":  "string|max_len:2",
		"faktor_kecurigaan_ganas3": "in:Ya,Tidak",
		"nilai_kecurigaan_ganas3":  "string|max_len:2",
		"faktor_kecurigaan_ganas4": "in:Ya,Tidak",
		"nilai_kecurigaan_ganas4":  "string|max_len:2",
		"faktor_kecurigaan_ganas5": "in:Ya,Tidak",
		"nilai_kecurigaan_ganas5":  "string|max_len:2",
		"faktor_kecurigaan_ganas6": "in:Ya,Tidak",
		"nilai_kecurigaan_ganas6":  "string|max_len:2",
		"faktor_kecurigaan_ganas7": "in:Ya,Tidak",
		"nilai_kecurigaan_ganas7":  "string|max_len:2",
		"faktor_kecurigaan_ganas8": "in:Ya,Tidak",
		"nilai_kecurigaan_ganas8":  "string|max_len:2",
		"total_skor":               "string|max_len:3",
		"hasil_sadanis":            "in:Benjolan,Tidak Ada Benjolan,Curiga Kanker",
		"tindak_lanjut_sadanis":    "in:Dirujuk,Tidak Dirujuk",
		"hasil_skrining":           "in:Normal,Kemungkinan Kelainan Payudara Jinak,Curiga Kelainan Payudara Ganas",
		"keterangan":               "string|max_len:50",
		"nip":                      "required|string|max_len:20",
	}
	return rules
}

// SkriningRisikoKankerPayudaraStore simpan skrining risiko kanker payudara; kolom waktu kunci kosong = sekarang.
type SkriningRisikoKankerPayudaraStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningRisikoKankerPayudaraData
}

func (r *SkriningRisikoKankerPayudaraStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningRisikoKankerPayudaraStore) Rules(ctx http.Context) map[string]any {
	rules := skriningRisikoKankerPayudaraRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningRisikoKankerPayudaraStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat}
}

func (r *SkriningRisikoKankerPayudaraStore) Payload() SkriningRisikoKankerPayudaraData {
	return r.SkriningRisikoKankerPayudaraData
}

func (r *SkriningRisikoKankerPayudaraStore) DetailValues() map[string][]string { return nil }

// SkriningRisikoKankerPayudaraUpdate ubah skrining risiko kanker payudara (PUT); kunci lewat query string.
type SkriningRisikoKankerPayudaraUpdate struct {
	SkriningRisikoKankerPayudaraData
}

func (r *SkriningRisikoKankerPayudaraUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningRisikoKankerPayudaraUpdate) Rules(ctx http.Context) map[string]any {
	return skriningRisikoKankerPayudaraRules()
}

func (r *SkriningRisikoKankerPayudaraUpdate) Payload() SkriningRisikoKankerPayudaraData {
	return r.SkriningRisikoKankerPayudaraData
}

func (r *SkriningRisikoKankerPayudaraUpdate) DetailValues() map[string][]string { return nil }
