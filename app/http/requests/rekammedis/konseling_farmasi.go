package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// KonselingFarmasiData isian konseling farmasi.
type KonselingFarmasiData struct {
	Tanggal       string `form:"tanggal" json:"tanggal"`
	Diagnosa      string `form:"diagnosa" json:"diagnosa"`
	ObatPemakaian string `form:"obat_pemakaian" json:"obat_pemakaian"`
	RiwayatAlergi string `form:"riwayat_alergi" json:"riwayat_alergi"`
	Keluhan       string `form:"keluhan" json:"keluhan"`
	PernahDatang  string `form:"pernah_datang" json:"pernah_datang"`
	TindakLanjut  string `form:"tindak_lanjut" json:"tindak_lanjut"`
	Nip           string `form:"nip" json:"nip"`
}

func konselingFarmasiRules() map[string]any {
	rules := map[string]any{
		"tanggal":        "required|date",
		"diagnosa":       "string|max_len:40",
		"obat_pemakaian": "string|max_len:700",
		"riwayat_alergi": "string|max_len:30",
		"keluhan":        "string|max_len:300",
		"pernah_datang":  "in:Ya,Tidak",
		"tindak_lanjut":  "string|max_len:400",
		"nip":            "required|string|max_len:20",
	}
	return rules
}

// KonselingFarmasiStore simpan konseling farmasi; kolom waktu kunci kosong = sekarang.
type KonselingFarmasiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	KonselingFarmasiData
}

func (r *KonselingFarmasiStore) Authorize(ctx http.Context) error { return nil }

func (r *KonselingFarmasiStore) Rules(ctx http.Context) map[string]any {
	rules := konselingFarmasiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *KonselingFarmasiStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *KonselingFarmasiStore) Payload() KonselingFarmasiData { return r.KonselingFarmasiData }

func (r *KonselingFarmasiStore) DetailValues() map[string][]string { return nil }

// KonselingFarmasiUpdate ubah konseling farmasi (PUT); kunci lewat query string.
type KonselingFarmasiUpdate struct {
	KonselingFarmasiData
}

func (r *KonselingFarmasiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *KonselingFarmasiUpdate) Rules(ctx http.Context) map[string]any {
	return konselingFarmasiRules()
}

func (r *KonselingFarmasiUpdate) Payload() KonselingFarmasiData { return r.KonselingFarmasiData }

func (r *KonselingFarmasiUpdate) DetailValues() map[string][]string { return nil }
