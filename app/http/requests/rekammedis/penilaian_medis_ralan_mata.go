package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianMedisRalanMataData isian penilaian awal medis ralan mata.
type PenilaianMedisRalanMataData struct {
	Tanggal      string `form:"tanggal" json:"tanggal"`
	KdDokter     string `form:"kd_dokter" json:"kd_dokter"`
	Anamnesis    string `form:"anamnesis" json:"anamnesis"`
	Hubungan     string `form:"hubungan" json:"hubungan"`
	KeluhanUtama string `form:"keluhan_utama" json:"keluhan_utama"`
	Rps          string `form:"rps" json:"rps"`
	Rpd          string `form:"rpd" json:"rpd"`
	Rpo          string `form:"rpo" json:"rpo"`
	Alergi       string `form:"alergi" json:"alergi"`
	Status       string `form:"status" json:"status"`
	Td           string `form:"td" json:"td"`
	Nadi         string `form:"nadi" json:"nadi"`
	Rr           string `form:"rr" json:"rr"`
	Suhu         string `form:"suhu" json:"suhu"`
	Nyeri        string `form:"nyeri" json:"nyeri"`
	Bb           string `form:"bb" json:"bb"`
	Visuskanan   string `form:"visuskanan" json:"visuskanan"`
	Visuskiri    string `form:"visuskiri" json:"visuskiri"`
	Cckanan      string `form:"cckanan" json:"cckanan"`
	Cckiri       string `form:"cckiri" json:"cckiri"`
	Palkanan     string `form:"palkanan" json:"palkanan"`
	Palkiri      string `form:"palkiri" json:"palkiri"`
	Conkanan     string `form:"conkanan" json:"conkanan"`
	Conkiri      string `form:"conkiri" json:"conkiri"`
	Corneakanan  string `form:"corneakanan" json:"corneakanan"`
	Corneakiri   string `form:"corneakiri" json:"corneakiri"`
	Coakanan     string `form:"coakanan" json:"coakanan"`
	Coakiri      string `form:"coakiri" json:"coakiri"`
	Pupilkanan   string `form:"pupilkanan" json:"pupilkanan"`
	Pupilkiri    string `form:"pupilkiri" json:"pupilkiri"`
	Lensakanan   string `form:"lensakanan" json:"lensakanan"`
	Lensakiri    string `form:"lensakiri" json:"lensakiri"`
	Funduskanan  string `form:"funduskanan" json:"funduskanan"`
	Funduskiri   string `form:"funduskiri" json:"funduskiri"`
	Papilkanan   string `form:"papilkanan" json:"papilkanan"`
	Papilkiri    string `form:"papilkiri" json:"papilkiri"`
	Retinakanan  string `form:"retinakanan" json:"retinakanan"`
	Retinakiri   string `form:"retinakiri" json:"retinakiri"`
	Makulakanan  string `form:"makulakanan" json:"makulakanan"`
	Makulakiri   string `form:"makulakiri" json:"makulakiri"`
	Tiokanan     string `form:"tiokanan" json:"tiokanan"`
	Tiokiri      string `form:"tiokiri" json:"tiokiri"`
	Mbokanan     string `form:"mbokanan" json:"mbokanan"`
	Mbokiri      string `form:"mbokiri" json:"mbokiri"`
	Lab          string `form:"lab" json:"lab"`
	Rad          string `form:"rad" json:"rad"`
	Penunjang    string `form:"penunjang" json:"penunjang"`
	Tes          string `form:"tes" json:"tes"`
	Pemeriksaan  string `form:"pemeriksaan" json:"pemeriksaan"`
	Diagnosis    string `form:"diagnosis" json:"diagnosis"`
	Diagnosisbdg string `form:"diagnosisbdg" json:"diagnosisbdg"`
	Permasalahan string `form:"permasalahan" json:"permasalahan"`
	Terapi       string `form:"terapi" json:"terapi"`
	Tindakan     string `form:"tindakan" json:"tindakan"`
	Edukasi      string `form:"edukasi" json:"edukasi"`
}

func penilaianMedisRalanMataRules() map[string]any {
	rules := map[string]any{
		"tanggal":       "required|date",
		"kd_dokter":     "required|string|max_len:20",
		"anamnesis":     "required|in:Autoanamnesis,Alloanamnesis",
		"hubungan":      "string|max_len:30",
		"keluhan_utama": "string|max_len:2000",
		"rps":           "string|max_len:2000",
		"rpd":           "string|max_len:1000",
		"rpo":           "string|max_len:1000",
		"alergi":        "string|max_len:50",
		"status":        "string|max_len:50",
		"td":            "string|max_len:8",
		"nadi":          "string|max_len:5",
		"rr":            "string|max_len:5",
		"suhu":          "string|max_len:5",
		"nyeri":         "string|max_len:50",
		"bb":            "string|max_len:5",
		"visuskanan":    "string|max_len:100",
		"visuskiri":     "string|max_len:100",
		"cckanan":       "string|max_len:100",
		"cckiri":        "string|max_len:100",
		"palkanan":      "string|max_len:100",
		"palkiri":       "string|max_len:100",
		"conkanan":      "string|max_len:100",
		"conkiri":       "string|max_len:100",
		"corneakanan":   "string|max_len:100",
		"corneakiri":    "string|max_len:100",
		"coakanan":      "string|max_len:100",
		"coakiri":       "string|max_len:100",
		"pupilkanan":    "string|max_len:100",
		"pupilkiri":     "string|max_len:100",
		"lensakanan":    "string|max_len:100",
		"lensakiri":     "string|max_len:100",
		"funduskanan":   "string|max_len:100",
		"funduskiri":    "string|max_len:100",
		"papilkanan":    "string|max_len:100",
		"papilkiri":     "string|max_len:100",
		"retinakanan":   "string|max_len:100",
		"retinakiri":    "string|max_len:100",
		"makulakanan":   "string|max_len:100",
		"makulakiri":    "string|max_len:100",
		"tiokanan":      "string|max_len:100",
		"tiokiri":       "string|max_len:100",
		"mbokanan":      "string|max_len:100",
		"mbokiri":       "string|max_len:100",
		"lab":           "string",
		"rad":           "string",
		"penunjang":     "string",
		"tes":           "string",
		"pemeriksaan":   "string",
		"diagnosis":     "string|max_len:500",
		"diagnosisbdg":  "string|max_len:500",
		"permasalahan":  "string",
		"terapi":        "string",
		"tindakan":      "string",
		"edukasi":       "string|max_len:1000",
	}
	return rules
}

// PenilaianMedisRalanMataStore simpan penilaian awal medis ralan mata; kolom waktu kunci kosong = sekarang.
type PenilaianMedisRalanMataStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianMedisRalanMataData
}

func (r *PenilaianMedisRalanMataStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanMataStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianMedisRalanMataRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianMedisRalanMataStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianMedisRalanMataStore) Payload() PenilaianMedisRalanMataData {
	return r.PenilaianMedisRalanMataData
}

func (r *PenilaianMedisRalanMataStore) DetailValues() map[string][]string { return nil }

// PenilaianMedisRalanMataUpdate ubah penilaian awal medis ralan mata (PUT); kunci lewat query string.
type PenilaianMedisRalanMataUpdate struct {
	PenilaianMedisRalanMataData
}

func (r *PenilaianMedisRalanMataUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianMedisRalanMataUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianMedisRalanMataRules()
}

func (r *PenilaianMedisRalanMataUpdate) Payload() PenilaianMedisRalanMataData {
	return r.PenilaianMedisRalanMataData
}

func (r *PenilaianMedisRalanMataUpdate) DetailValues() map[string][]string { return nil }
