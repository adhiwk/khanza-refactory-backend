package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PelaksanaanInformasiEdukasiData isian pelaksanaan informasi edukasi.
type PelaksanaanInformasiEdukasiData struct {
	Nik                     string `form:"nik" json:"nik"`
	MateriEdukasi           string `form:"materi_edukasi" json:"materi_edukasi"`
	Keterangan              string `form:"keterangan" json:"keterangan"`
	DiberikanPada           string `form:"diberikan_pada" json:"diberikan_pada"`
	KeteranganDiberikanPada string `form:"keterangan_diberikan_pada" json:"keterangan_diberikan_pada"`
	LamaEdukasi             string `form:"lama_edukasi" json:"lama_edukasi"`
	MetodeEdukasi           string `form:"metode_edukasi" json:"metode_edukasi"`
	HasilVerifikasi         string `form:"hasil_verifikasi" json:"hasil_verifikasi"`
	Status                  string `form:"status" json:"status"`
}

func pelaksanaanInformasiEdukasiRules() map[string]any {
	rules := map[string]any{
		"nik":                       "required|string|max_len:20",
		"materi_edukasi":            "string|max_len:1000",
		"keterangan":                "string|max_len:50",
		"diberikan_pada":            "required|in:Pasien,Keluarga,Lain-lain",
		"keterangan_diberikan_pada": "string|max_len:40",
		"lama_edukasi":              "string|max_len:10",
		"metode_edukasi":            "in:Ceramah,Diskusi,Demonstrasi",
		"hasil_verifikasi":          "in:Sudah Mengerti,Re-Edukasi,Re-Demonstrasi",
		"status":                    "in:Awal,Ulang",
	}
	return rules
}

// PelaksanaanInformasiEdukasiStore simpan pelaksanaan informasi edukasi; kolom waktu kunci kosong = sekarang.
type PelaksanaanInformasiEdukasiStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PelaksanaanInformasiEdukasiData
}

func (r *PelaksanaanInformasiEdukasiStore) Authorize(ctx http.Context) error { return nil }

func (r *PelaksanaanInformasiEdukasiStore) Rules(ctx http.Context) map[string]any {
	rules := pelaksanaanInformasiEdukasiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PelaksanaanInformasiEdukasiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PelaksanaanInformasiEdukasiStore) Payload() PelaksanaanInformasiEdukasiData {
	return r.PelaksanaanInformasiEdukasiData
}

func (r *PelaksanaanInformasiEdukasiStore) DetailValues() map[string][]string { return nil }

// PelaksanaanInformasiEdukasiUpdate ubah pelaksanaan informasi edukasi (PUT); kunci lewat query string.
type PelaksanaanInformasiEdukasiUpdate struct {
	PelaksanaanInformasiEdukasiData
}

func (r *PelaksanaanInformasiEdukasiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PelaksanaanInformasiEdukasiUpdate) Rules(ctx http.Context) map[string]any {
	return pelaksanaanInformasiEdukasiRules()
}

func (r *PelaksanaanInformasiEdukasiUpdate) Payload() PelaksanaanInformasiEdukasiData {
	return r.PelaksanaanInformasiEdukasiData
}

func (r *PelaksanaanInformasiEdukasiUpdate) DetailValues() map[string][]string { return nil }
