package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SignoutSebelumMenutupLukaData isian sign out sebelum menutup luka.
type SignoutSebelumMenutupLukaData struct {
	Sncn                             string `form:"sncn" json:"sncn"`
	Tindakan                         string `form:"tindakan" json:"tindakan"`
	KdDokterBedah                    string `form:"kd_dokter_bedah" json:"kd_dokter_bedah"`
	KdDokterAnestesi                 string `form:"kd_dokter_anestesi" json:"kd_dokter_anestesi"`
	VerbalTindakan                   string `form:"verbal_tindakan" json:"verbal_tindakan"`
	VerbalKelengkapanKasa            string `form:"verbal_kelengkapan_kasa" json:"verbal_kelengkapan_kasa"`
	VerbalInstrumen                  string `form:"verbal_instrumen" json:"verbal_instrumen"`
	VerbalAlatTajam                  string `form:"verbal_alat_tajam" json:"verbal_alat_tajam"`
	KelengkapanSpecimenLabel         string `form:"kelengkapan_specimen_label" json:"kelengkapan_specimen_label"`
	KelengkapanSpecimenFormulir      string `form:"kelengkapan_specimen_formulir" json:"kelengkapan_specimen_formulir"`
	PeninjauanKegiatanDokterBedah    string `form:"peninjauan_kegiatan_dokter_bedah" json:"peninjauan_kegiatan_dokter_bedah"`
	PeninjauanKegiatanDokterAnestesi string `form:"peninjauan_kegiatan_dokter_anestesi" json:"peninjauan_kegiatan_dokter_anestesi"`
	PeninjauanKegiatanPerawatKamarOk string `form:"peninjauan_kegiatan_perawat_kamar_ok" json:"peninjauan_kegiatan_perawat_kamar_ok"`
	PerhatianUtamaFasePemulihan      string `form:"perhatian_utama_fase_pemulihan" json:"perhatian_utama_fase_pemulihan"`
	NipPerawatOk                     string `form:"nip_perawat_ok" json:"nip_perawat_ok"`
}

func signoutSebelumMenutupLukaRules() map[string]any {
	rules := map[string]any{
		"sncn":                                 "string|max_len:25",
		"tindakan":                             "string|max_len:50",
		"kd_dokter_bedah":                      "required|string|max_len:20",
		"kd_dokter_anestesi":                   "required|string|max_len:20",
		"verbal_tindakan":                      "in:Ya,Tidak",
		"verbal_kelengkapan_kasa":              "in:Ya,Tidak",
		"verbal_instrumen":                     "in:Ya,Tidak",
		"verbal_alat_tajam":                    "in:Ya,Tidak",
		"kelengkapan_specimen_label":           "in:Lengkap,Tidak Lengkap,Tidak Ada Pemeriksaan Spesimen",
		"kelengkapan_specimen_formulir":        "in:Lengkap,Tidak Lengkap,Tidak Ada Pemeriksaan Spesimen",
		"peninjauan_kegiatan_dokter_bedah":     "in:Ya,Tidak",
		"peninjauan_kegiatan_dokter_anestesi":  "in:Ya,Tidak",
		"peninjauan_kegiatan_perawat_kamar_ok": "in:Ya,Tidak",
		"perhatian_utama_fase_pemulihan":       "string|max_len:100",
		"nip_perawat_ok":                       "string|max_len:20",
	}
	return rules
}

// SignoutSebelumMenutupLukaStore simpan sign out sebelum menutup luka; kolom waktu kunci kosong = sekarang.
type SignoutSebelumMenutupLukaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	SignoutSebelumMenutupLukaData
}

func (r *SignoutSebelumMenutupLukaStore) Authorize(ctx http.Context) error { return nil }

func (r *SignoutSebelumMenutupLukaStore) Rules(ctx http.Context) map[string]any {
	rules := signoutSebelumMenutupLukaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *SignoutSebelumMenutupLukaStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *SignoutSebelumMenutupLukaStore) Payload() SignoutSebelumMenutupLukaData {
	return r.SignoutSebelumMenutupLukaData
}

func (r *SignoutSebelumMenutupLukaStore) DetailValues() map[string][]string { return nil }

// SignoutSebelumMenutupLukaUpdate ubah sign out sebelum menutup luka (PUT); kunci lewat query string.
type SignoutSebelumMenutupLukaUpdate struct {
	SignoutSebelumMenutupLukaData
}

func (r *SignoutSebelumMenutupLukaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SignoutSebelumMenutupLukaUpdate) Rules(ctx http.Context) map[string]any {
	return signoutSebelumMenutupLukaRules()
}

func (r *SignoutSebelumMenutupLukaUpdate) Payload() SignoutSebelumMenutupLukaData {
	return r.SignoutSebelumMenutupLukaData
}

func (r *SignoutSebelumMenutupLukaUpdate) DetailValues() map[string][]string { return nil }
