package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanAdimeGiziData isian catatan ADIME gizi.
type CatatanAdimeGiziData struct {
	Asesmen    string `form:"asesmen" json:"asesmen"`
	Diagnosis  string `form:"diagnosis" json:"diagnosis"`
	Intervensi string `form:"intervensi" json:"intervensi"`
	Monitoring string `form:"monitoring" json:"monitoring"`
	Evaluasi   string `form:"evaluasi" json:"evaluasi"`
	Instruksi  string `form:"instruksi" json:"instruksi"`
	Nip        string `form:"nip" json:"nip"`
}

func catatanAdimeGiziRules() map[string]any {
	rules := map[string]any{
		"asesmen":    "string|max_len:1000",
		"diagnosis":  "string|max_len:1000",
		"intervensi": "string|max_len:1000",
		"monitoring": "string|max_len:1000",
		"evaluasi":   "string|max_len:1000",
		"instruksi":  "string|max_len:1000",
		"nip":        "string|max_len:20",
	}
	return rules
}

// CatatanAdimeGiziStore simpan catatan ADIME gizi; kolom waktu kunci kosong = sekarang.
type CatatanAdimeGiziStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	CatatanAdimeGiziData
}

func (r *CatatanAdimeGiziStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanAdimeGiziStore) Rules(ctx http.Context) map[string]any {
	rules := catatanAdimeGiziRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *CatatanAdimeGiziStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *CatatanAdimeGiziStore) Payload() CatatanAdimeGiziData { return r.CatatanAdimeGiziData }

func (r *CatatanAdimeGiziStore) DetailValues() map[string][]string { return nil }

// CatatanAdimeGiziUpdate ubah catatan ADIME gizi (PUT); kunci lewat query string.
type CatatanAdimeGiziUpdate struct {
	CatatanAdimeGiziData
}

func (r *CatatanAdimeGiziUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanAdimeGiziUpdate) Rules(ctx http.Context) map[string]any {
	return catatanAdimeGiziRules()
}

func (r *CatatanAdimeGiziUpdate) Payload() CatatanAdimeGiziData { return r.CatatanAdimeGiziData }

func (r *CatatanAdimeGiziUpdate) DetailValues() map[string][]string { return nil }
