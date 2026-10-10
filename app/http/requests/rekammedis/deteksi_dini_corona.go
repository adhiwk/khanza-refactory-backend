package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// DeteksiDiniCoronaData isian deteksi dini corona.
type DeteksiDiniCoronaData struct {
	Tanggal                    string `form:"tanggal" json:"tanggal"`
	Nip                        string `form:"nip" json:"nip"`
	GejalaDemam                string `form:"gejala_demam" json:"gejala_demam"`
	GejalaBatuk                string `form:"gejala_batuk" json:"gejala_batuk"`
	GejalaSesak                string `form:"gejala_sesak" json:"gejala_sesak"`
	GejalaTanggalPertama       string `form:"gejala_tanggal_pertama" json:"gejala_tanggal_pertama"`
	GejalaRiwayatSakit         string `form:"gejala_riwayat_sakit" json:"gejala_riwayat_sakit"`
	GejalaRiwayatPeriksa       string `form:"gejala_riwayat_periksa" json:"gejala_riwayat_periksa"`
	FaktorRiwayatPerjalanan    string `form:"faktor_riwayat_perjalanan" json:"faktor_riwayat_perjalanan"`
	FaktorAsalDaerah           string `form:"faktor_asal_daerah" json:"faktor_asal_daerah"`
	FaktorTanggalKedatangan    string `form:"faktor_tanggal_kedatangan" json:"faktor_tanggal_kedatangan"`
	FaktorPaparanKontakpositif string `form:"faktor_paparan_kontakpositif" json:"faktor_paparan_kontakpositif"`
	FaktorPaparanKontakpdp     string `form:"faktor_paparan_kontakpdp" json:"faktor_paparan_kontakpdp"`
	FaktorPaparanFaskespositif string `form:"faktor_paparan_faskespositif" json:"faktor_paparan_faskespositif"`
	FaktorPaparanPerjalananln  string `form:"faktor_paparan_perjalananln" json:"faktor_paparan_perjalananln"`
	FaktorPaparanPasarhewan    string `form:"faktor_paparan_pasarhewan" json:"faktor_paparan_pasarhewan"`
	Kesimpulan                 string `form:"kesimpulan" json:"kesimpulan"`
	TindakLanjut               string `form:"tindak_lanjut" json:"tindak_lanjut"`
}

func deteksiDiniCoronaRules() map[string]any {
	rules := map[string]any{
		"tanggal":                      "required|date",
		"nip":                          "string|max_len:20",
		"gejala_demam":                 "in:Ya,Tidak",
		"gejala_batuk":                 "in:Ya,Tidak",
		"gejala_sesak":                 "in:Ya,Tidak",
		"gejala_tanggal_pertama":       "date",
		"gejala_riwayat_sakit":         "string|max_len:50",
		"gejala_riwayat_periksa":       "string|max_len:50",
		"faktor_riwayat_perjalanan":    "required|in:Ya,Tidak",
		"faktor_asal_daerah":           "string|max_len:50",
		"faktor_tanggal_kedatangan":    "required|date",
		"faktor_paparan_kontakpositif": "required|in:Ya,Tidak",
		"faktor_paparan_kontakpdp":     "required|in:Ya,Tidak",
		"faktor_paparan_faskespositif": "required|in:Ya,Tidak",
		"faktor_paparan_perjalananln":  "required|in:Ya,Tidak",
		"faktor_paparan_pasarhewan":    "required|in:Ya,Tidak",
		"kesimpulan":                   "required|in:ODP,PDP,OTG,Bukan ketiganya",
		"tindak_lanjut":                "required|in:Rujuk,Rawat Inap,Rawat Jalan",
	}
	return rules
}

// DeteksiDiniCoronaStore simpan deteksi dini corona; kolom waktu kunci kosong = sekarang.
type DeteksiDiniCoronaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	DeteksiDiniCoronaData
}

func (r *DeteksiDiniCoronaStore) Authorize(ctx http.Context) error { return nil }

func (r *DeteksiDiniCoronaStore) Rules(ctx http.Context) map[string]any {
	rules := deteksiDiniCoronaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *DeteksiDiniCoronaStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *DeteksiDiniCoronaStore) Payload() DeteksiDiniCoronaData { return r.DeteksiDiniCoronaData }

func (r *DeteksiDiniCoronaStore) DetailValues() map[string][]string { return nil }

// DeteksiDiniCoronaUpdate ubah deteksi dini corona (PUT); kunci lewat query string.
type DeteksiDiniCoronaUpdate struct {
	DeteksiDiniCoronaData
}

func (r *DeteksiDiniCoronaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *DeteksiDiniCoronaUpdate) Rules(ctx http.Context) map[string]any {
	return deteksiDiniCoronaRules()
}

func (r *DeteksiDiniCoronaUpdate) Payload() DeteksiDiniCoronaData { return r.DeteksiDiniCoronaData }

func (r *DeteksiDiniCoronaUpdate) DetailValues() map[string][]string { return nil }
