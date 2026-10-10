package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanCekGdsData isian catatan cek GDS.
type CatatanCekGdsData struct {
	Gdp      string `form:"gdp" json:"gdp"`
	Insulin  string `form:"insulin" json:"insulin"`
	ObatGula string `form:"obat_gula" json:"obat_gula"`
	Nip      string `form:"nip" json:"nip"`
}

func catatanCekGdsRules() map[string]any {
	rules := map[string]any{
		"gdp":       "string|max_len:5",
		"insulin":   "string|max_len:30",
		"obat_gula": "string|max_len:30",
		"nip":       "required|string|max_len:20",
	}
	return rules
}

// CatatanCekGdsStore simpan catatan cek GDS; kolom waktu kunci kosong = sekarang.
type CatatanCekGdsStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	CatatanCekGdsData
}

func (r *CatatanCekGdsStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanCekGdsStore) Rules(ctx http.Context) map[string]any {
	rules := catatanCekGdsRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *CatatanCekGdsStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *CatatanCekGdsStore) Payload() CatatanCekGdsData { return r.CatatanCekGdsData }

func (r *CatatanCekGdsStore) DetailValues() map[string][]string { return nil }

// CatatanCekGdsUpdate ubah catatan cek GDS (PUT); kunci lewat query string.
type CatatanCekGdsUpdate struct {
	CatatanCekGdsData
}

func (r *CatatanCekGdsUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanCekGdsUpdate) Rules(ctx http.Context) map[string]any { return catatanCekGdsRules() }

func (r *CatatanCekGdsUpdate) Payload() CatatanCekGdsData { return r.CatatanCekGdsData }

func (r *CatatanCekGdsUpdate) DetailValues() map[string][]string { return nil }
