package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanKeperawatanRalanData isian catatan keperawatan ralan.
type CatatanKeperawatanRalanData struct {
	Uraian string `form:"uraian" json:"uraian"`
	Nip    string `form:"nip" json:"nip"`
}

func catatanKeperawatanRalanRules() map[string]any {
	rules := map[string]any{
		"uraian": "string|max_len:1000",
		"nip":    "required|string|max_len:20",
	}
	return rules
}

// CatatanKeperawatanRalanStore simpan catatan keperawatan ralan; kolom waktu kunci kosong = sekarang.
type CatatanKeperawatanRalanStore struct {
	Tanggal string `form:"tanggal" json:"tanggal"`
	Jam     string `form:"jam" json:"jam"`
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	CatatanKeperawatanRalanData
}

func (r *CatatanKeperawatanRalanStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanKeperawatanRalanStore) Rules(ctx http.Context) map[string]any {
	rules := catatanKeperawatanRalanRules()
	rules["tanggal"] = "date"
	rules["jam"] = "string|len:8"
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *CatatanKeperawatanRalanStore) KeyValues() repo.Key {
	return repo.Key{"tanggal": r.Tanggal, "jam": r.Jam, "no_rawat": r.NoRawat}
}

func (r *CatatanKeperawatanRalanStore) Payload() CatatanKeperawatanRalanData {
	return r.CatatanKeperawatanRalanData
}

func (r *CatatanKeperawatanRalanStore) DetailValues() map[string][]string { return nil }

// CatatanKeperawatanRalanUpdate ubah catatan keperawatan ralan (PUT); kunci lewat query string.
type CatatanKeperawatanRalanUpdate struct {
	CatatanKeperawatanRalanData
}

func (r *CatatanKeperawatanRalanUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanKeperawatanRalanUpdate) Rules(ctx http.Context) map[string]any {
	return catatanKeperawatanRalanRules()
}

func (r *CatatanKeperawatanRalanUpdate) Payload() CatatanKeperawatanRalanData {
	return r.CatatanKeperawatanRalanData
}

func (r *CatatanKeperawatanRalanUpdate) DetailValues() map[string][]string { return nil }
