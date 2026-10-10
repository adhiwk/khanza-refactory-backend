package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanKeperawatanRanapData isian catatan keperawatan ranap.
type CatatanKeperawatanRanapData struct {
	Uraian string `form:"uraian" json:"uraian"`
	Nip    string `form:"nip" json:"nip"`
}

func catatanKeperawatanRanapRules() map[string]any {
	rules := map[string]any{
		"uraian": "string|max_len:1000",
		"nip":    "required|string|max_len:20",
	}
	return rules
}

// CatatanKeperawatanRanapStore simpan catatan keperawatan ranap; kolom waktu kunci kosong = sekarang.
type CatatanKeperawatanRanapStore struct {
	Tanggal string `form:"tanggal" json:"tanggal"`
	Jam     string `form:"jam" json:"jam"`
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	CatatanKeperawatanRanapData
}

func (r *CatatanKeperawatanRanapStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanKeperawatanRanapStore) Rules(ctx http.Context) map[string]any {
	rules := catatanKeperawatanRanapRules()
	rules["tanggal"] = "date"
	rules["jam"] = "string|len:8"
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *CatatanKeperawatanRanapStore) KeyValues() repo.Key {
	return repo.Key{"tanggal": r.Tanggal, "jam": r.Jam, "no_rawat": r.NoRawat}
}

func (r *CatatanKeperawatanRanapStore) Payload() CatatanKeperawatanRanapData {
	return r.CatatanKeperawatanRanapData
}

func (r *CatatanKeperawatanRanapStore) DetailValues() map[string][]string { return nil }

// CatatanKeperawatanRanapUpdate ubah catatan keperawatan ranap (PUT); kunci lewat query string.
type CatatanKeperawatanRanapUpdate struct {
	CatatanKeperawatanRanapData
}

func (r *CatatanKeperawatanRanapUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanKeperawatanRanapUpdate) Rules(ctx http.Context) map[string]any {
	return catatanKeperawatanRanapRules()
}

func (r *CatatanKeperawatanRanapUpdate) Payload() CatatanKeperawatanRanapData {
	return r.CatatanKeperawatanRanapData
}

func (r *CatatanKeperawatanRanapUpdate) DetailValues() map[string][]string { return nil }
