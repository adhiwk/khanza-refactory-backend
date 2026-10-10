package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningIndraPendengaranData isian skrining indra pendengaran.
type SkriningIndraPendengaranData struct {
	Tanggal                          string `form:"tanggal" json:"tanggal"`
	CurigaTuliTelingaKiri            string `form:"curiga_tuli_telinga_kiri" json:"curiga_tuli_telinga_kiri"`
	CurigaTuliTelingaKanan           string `form:"curiga_tuli_telinga_kanan" json:"curiga_tuli_telinga_kanan"`
	CurigaTuliTelingaRujuk           string `form:"curiga_tuli_telinga_rujuk" json:"curiga_tuli_telinga_rujuk"`
	PenurunanPendengaranTelingaKiri  string `form:"penurunan_pendengaran_telinga_kiri" json:"penurunan_pendengaran_telinga_kiri"`
	PenurunanPendengaranTelingaKanan string `form:"penurunan_pendengaran_telinga_kanan" json:"penurunan_pendengaran_telinga_kanan"`
	MendengarBisikanTelingaKiri      string `form:"mendengar_bisikan_telinga_kiri" json:"mendengar_bisikan_telinga_kiri"`
	MendengarBisikanTelingaKanan     string `form:"mendengar_bisikan_telinga_kanan" json:"mendengar_bisikan_telinga_kanan"`
	CongekTelingaKiri                string `form:"congek_telinga_kiri" json:"congek_telinga_kiri"`
	CongekTelingaKanan               string `form:"congek_telinga_kanan" json:"congek_telinga_kanan"`
	CongekTelingaRujuk               string `form:"congek_telinga_rujuk" json:"congek_telinga_rujuk"`
	SumbatanSerumenTelingaKiri       string `form:"sumbatan_serumen_telinga_kiri" json:"sumbatan_serumen_telinga_kiri"`
	SumbatanSerumenTelingaKanan      string `form:"sumbatan_serumen_telinga_kanan" json:"sumbatan_serumen_telinga_kanan"`
	SumbatanSerumenTelingaRujuk      string `form:"sumbatan_serumen_telinga_rujuk" json:"sumbatan_serumen_telinga_rujuk"`
	HasilSkrining                    string `form:"hasil_skrining" json:"hasil_skrining"`
	Keterangan                       string `form:"keterangan" json:"keterangan"`
	Nip                              string `form:"nip" json:"nip"`
}

func skriningIndraPendengaranRules() map[string]any {
	rules := map[string]any{
		"tanggal":                             "required|date",
		"curiga_tuli_telinga_kiri":            "in:Tidak,Ya",
		"curiga_tuli_telinga_kanan":           "in:Tidak,Ya",
		"curiga_tuli_telinga_rujuk":           "in:Tidak,Ya",
		"penurunan_pendengaran_telinga_kiri":  "in:Tidak,Ya",
		"penurunan_pendengaran_telinga_kanan": "in:Tidak,Ya",
		"mendengar_bisikan_telinga_kiri":      "in:Normal,Gangguan Pendengaran",
		"mendengar_bisikan_telinga_kanan":     "in:Normal,Gangguan Pendengaran",
		"congek_telinga_kiri":                 "in:Tidak,Ya",
		"congek_telinga_kanan":                "in:Tidak,Ya",
		"congek_telinga_rujuk":                "in:Tidak,Ya",
		"sumbatan_serumen_telinga_kiri":       "in:Tidak,Ya",
		"sumbatan_serumen_telinga_kanan":      "in:Tidak,Ya",
		"sumbatan_serumen_telinga_rujuk":      "in:Tidak,Ya",
		"hasil_skrining":                      "string|max_len:40",
		"keterangan":                          "string|max_len:100",
		"nip":                                 "required|string|max_len:20",
	}
	return rules
}

// SkriningIndraPendengaranStore simpan skrining indra pendengaran; kolom waktu kunci kosong = sekarang.
type SkriningIndraPendengaranStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningIndraPendengaranData
}

func (r *SkriningIndraPendengaranStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningIndraPendengaranStore) Rules(ctx http.Context) map[string]any {
	rules := skriningIndraPendengaranRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningIndraPendengaranStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningIndraPendengaranStore) Payload() SkriningIndraPendengaranData {
	return r.SkriningIndraPendengaranData
}

func (r *SkriningIndraPendengaranStore) DetailValues() map[string][]string { return nil }

// SkriningIndraPendengaranUpdate ubah skrining indra pendengaran (PUT); kunci lewat query string.
type SkriningIndraPendengaranUpdate struct {
	SkriningIndraPendengaranData
}

func (r *SkriningIndraPendengaranUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningIndraPendengaranUpdate) Rules(ctx http.Context) map[string]any {
	return skriningIndraPendengaranRules()
}

func (r *SkriningIndraPendengaranUpdate) Payload() SkriningIndraPendengaranData {
	return r.SkriningIndraPendengaranData
}

func (r *SkriningIndraPendengaranUpdate) DetailValues() map[string][]string { return nil }
