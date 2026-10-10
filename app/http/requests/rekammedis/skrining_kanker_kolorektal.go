package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningKankerKolorektalData isian skrining kanker kolorektal.
type SkriningKankerKolorektalData struct {
	Tanggal                   string `form:"tanggal" json:"tanggal"`
	Nip                       string `form:"nip" json:"nip"`
	RiwayatPolipAdenomatosa   string `form:"riwayat_polip_adenomatosa" json:"riwayat_polip_adenomatosa"`
	RiwayatBabBerdarah        string `form:"riwayat_bab_berdarah" json:"riwayat_bab_berdarah"`
	RiwayatReseksiKuratif     string `form:"riwayat_reseksi_kuratif" json:"riwayat_reseksi_kuratif"`
	ColokDubur                string `form:"colok_dubur" json:"colok_dubur"`
	RiwayatKolorektalKeluarga string `form:"riwayat_kolorektal_keluarga" json:"riwayat_kolorektal_keluarga"`
	DarahSamarFeses           string `form:"darah_samar_feses" json:"darah_samar_feses"`
	RujukFaskesLanjut         string `form:"rujuk_faskes_lanjut" json:"rujuk_faskes_lanjut"`
	Kesimpulan                string `form:"kesimpulan" json:"kesimpulan"`
	KeteranganKesimpulan      string `form:"keterangan_kesimpulan" json:"keterangan_kesimpulan"`
}

func skriningKankerKolorektalRules() map[string]any {
	rules := map[string]any{
		"tanggal":                     "required|date",
		"nip":                         "required|string|max_len:20",
		"riwayat_polip_adenomatosa":   "in:Tidak,Ya",
		"riwayat_bab_berdarah":        "in:Tidak,Ya",
		"riwayat_reseksi_kuratif":     "in:Tidak,Ya",
		"colok_dubur":                 "in:Tidak,Ya",
		"riwayat_kolorektal_keluarga": "in:Tidak,Ya",
		"darah_samar_feses":           "in:Tidak,Ya",
		"rujuk_faskes_lanjut":         "in:Tidak,Ya",
		"kesimpulan":                  "in:Normal,Suspek",
		"keterangan_kesimpulan":       "string|max_len:100",
	}
	return rules
}

// SkriningKankerKolorektalStore simpan skrining kanker kolorektal; kolom waktu kunci kosong = sekarang.
type SkriningKankerKolorektalStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningKankerKolorektalData
}

func (r *SkriningKankerKolorektalStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningKankerKolorektalStore) Rules(ctx http.Context) map[string]any {
	rules := skriningKankerKolorektalRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningKankerKolorektalStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningKankerKolorektalStore) Payload() SkriningKankerKolorektalData {
	return r.SkriningKankerKolorektalData
}

func (r *SkriningKankerKolorektalStore) DetailValues() map[string][]string { return nil }

// SkriningKankerKolorektalUpdate ubah skrining kanker kolorektal (PUT); kunci lewat query string.
type SkriningKankerKolorektalUpdate struct {
	SkriningKankerKolorektalData
}

func (r *SkriningKankerKolorektalUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningKankerKolorektalUpdate) Rules(ctx http.Context) map[string]any {
	return skriningKankerKolorektalRules()
}

func (r *SkriningKankerKolorektalUpdate) Payload() SkriningKankerKolorektalData {
	return r.SkriningKankerKolorektalData
}

func (r *SkriningKankerKolorektalUpdate) DetailValues() map[string][]string { return nil }
