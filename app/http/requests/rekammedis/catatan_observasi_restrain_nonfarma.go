package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanObservasiRestrainNonfarmaData isian catatan observasi restrain non farmakologi.
type CatatanObservasiRestrainNonfarmaData struct {
	TanganKiri        string `form:"tangan_kiri" json:"tangan_kiri"`
	TanganKanan       string `form:"tangan_kanan" json:"tangan_kanan"`
	KakiKiri          string `form:"kaki_kiri" json:"kaki_kiri"`
	KakiKanan         string `form:"kaki_kanan" json:"kaki_kanan"`
	Badan             string `form:"badan" json:"badan"`
	Edema             string `form:"edema" json:"edema"`
	Iritasi           string `form:"iritasi" json:"iritasi"`
	Sirkulasi         string `form:"sirkulasi" json:"sirkulasi"`
	KondisiKeterangan string `form:"kondisi_keterangan" json:"kondisi_keterangan"`
	Nip               string `form:"nip" json:"nip"`
}

func catatanObservasiRestrainNonfarmaRules() map[string]any {
	rules := map[string]any{
		"tangan_kiri":        "in:Ya,Tidak",
		"tangan_kanan":       "in:Ya,Tidak",
		"kaki_kiri":          "in:Ya,Tidak",
		"kaki_kanan":         "in:Ya,Tidak",
		"badan":              "in:Ya,Tidak",
		"edema":              "in:Ya,Tidak",
		"iritasi":            "in:Ya,Tidak",
		"sirkulasi":          "in:Ya,Tidak",
		"kondisi_keterangan": "string|max_len:100",
		"nip":                "required|string|max_len:20",
	}
	return rules
}

// CatatanObservasiRestrainNonfarmaStore simpan catatan observasi restrain non farmakologi; kolom waktu kunci kosong = sekarang.
type CatatanObservasiRestrainNonfarmaStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	CatatanObservasiRestrainNonfarmaData
}

func (r *CatatanObservasiRestrainNonfarmaStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiRestrainNonfarmaStore) Rules(ctx http.Context) map[string]any {
	rules := catatanObservasiRestrainNonfarmaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *CatatanObservasiRestrainNonfarmaStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *CatatanObservasiRestrainNonfarmaStore) Payload() CatatanObservasiRestrainNonfarmaData {
	return r.CatatanObservasiRestrainNonfarmaData
}

func (r *CatatanObservasiRestrainNonfarmaStore) DetailValues() map[string][]string { return nil }

// CatatanObservasiRestrainNonfarmaUpdate ubah catatan observasi restrain non farmakologi (PUT); kunci lewat query string.
type CatatanObservasiRestrainNonfarmaUpdate struct {
	CatatanObservasiRestrainNonfarmaData
}

func (r *CatatanObservasiRestrainNonfarmaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiRestrainNonfarmaUpdate) Rules(ctx http.Context) map[string]any {
	return catatanObservasiRestrainNonfarmaRules()
}

func (r *CatatanObservasiRestrainNonfarmaUpdate) Payload() CatatanObservasiRestrainNonfarmaData {
	return r.CatatanObservasiRestrainNonfarmaData
}

func (r *CatatanObservasiRestrainNonfarmaUpdate) DetailValues() map[string][]string { return nil }
