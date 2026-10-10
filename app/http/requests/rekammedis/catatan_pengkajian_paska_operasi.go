package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanPengkajianPaskaOperasiData isian catatan pengkajian paska operasi.
type CatatanPengkajianPaskaOperasiData struct {
	KdDokter           string `form:"kd_dokter" json:"kd_dokter"`
	RawatPaskaOperasi  string `form:"rawat_paska_operasi" json:"rawat_paska_operasi"`
	Cairan             string `form:"cairan" json:"cairan"`
	Antibiotika        string `form:"antibiotika" json:"antibiotika"`
	Analgetika         string `form:"analgetika" json:"analgetika"`
	MedikamentosaLain  string `form:"medikamentosa_lain" json:"medikamentosa_lain"`
	Diet               string `form:"diet" json:"diet"`
	PemeriksaanLaborat string `form:"pemeriksaan_laborat" json:"pemeriksaan_laborat"`
	Tranfusi           string `form:"tranfusi" json:"tranfusi"`
	Lainlain           string `form:"lainlain" json:"lainlain"`
}

func catatanPengkajianPaskaOperasiRules() map[string]any {
	rules := map[string]any{
		"kd_dokter":           "required|string|max_len:20",
		"rawat_paska_operasi": "string|max_len:250",
		"cairan":              "string|max_len:500",
		"antibiotika":         "string|max_len:500",
		"analgetika":          "string|max_len:500",
		"medikamentosa_lain":  "string|max_len:500",
		"diet":                "string|max_len:500",
		"pemeriksaan_laborat": "string|max_len:500",
		"tranfusi":            "string|max_len:250",
		"lainlain":            "string|max_len:500",
	}
	return rules
}

// CatatanPengkajianPaskaOperasiStore simpan catatan pengkajian paska operasi; kolom waktu kunci kosong = sekarang.
type CatatanPengkajianPaskaOperasiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	CatatanPengkajianPaskaOperasiData
}

func (r *CatatanPengkajianPaskaOperasiStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanPengkajianPaskaOperasiStore) Rules(ctx http.Context) map[string]any {
	rules := catatanPengkajianPaskaOperasiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *CatatanPengkajianPaskaOperasiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *CatatanPengkajianPaskaOperasiStore) Payload() CatatanPengkajianPaskaOperasiData {
	return r.CatatanPengkajianPaskaOperasiData
}

func (r *CatatanPengkajianPaskaOperasiStore) DetailValues() map[string][]string { return nil }

// CatatanPengkajianPaskaOperasiUpdate ubah catatan pengkajian paska operasi (PUT); kunci lewat query string.
type CatatanPengkajianPaskaOperasiUpdate struct {
	CatatanPengkajianPaskaOperasiData
}

func (r *CatatanPengkajianPaskaOperasiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanPengkajianPaskaOperasiUpdate) Rules(ctx http.Context) map[string]any {
	return catatanPengkajianPaskaOperasiRules()
}

func (r *CatatanPengkajianPaskaOperasiUpdate) Payload() CatatanPengkajianPaskaOperasiData {
	return r.CatatanPengkajianPaskaOperasiData
}

func (r *CatatanPengkajianPaskaOperasiUpdate) DetailValues() map[string][]string { return nil }
