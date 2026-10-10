package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianLanjutanSkriningFungsionalData isian penilaian lanjutan skrining fungsional.
type PenilaianLanjutanSkriningFungsionalData struct {
	PenilaianSkriningSkala1     string `form:"penilaian_skrining_skala1" json:"penilaian_skrining_skala1"`
	PenilaianSkriningNilai1     *int   `form:"penilaian_skrining_nilai1" json:"penilaian_skrining_nilai1"`
	PenilaianSkriningSkala2     string `form:"penilaian_skrining_skala2" json:"penilaian_skrining_skala2"`
	PenilaianSkriningNilai2     *int   `form:"penilaian_skrining_nilai2" json:"penilaian_skrining_nilai2"`
	PenilaianSkriningSkala3     string `form:"penilaian_skrining_skala3" json:"penilaian_skrining_skala3"`
	PenilaianSkriningNilai3     *int   `form:"penilaian_skrining_nilai3" json:"penilaian_skrining_nilai3"`
	PenilaianSkriningSkala4     string `form:"penilaian_skrining_skala4" json:"penilaian_skrining_skala4"`
	PenilaianSkriningNilai4     *int   `form:"penilaian_skrining_nilai4" json:"penilaian_skrining_nilai4"`
	PenilaianSkriningSkala5     string `form:"penilaian_skrining_skala5" json:"penilaian_skrining_skala5"`
	PenilaianSkriningNilai5     *int   `form:"penilaian_skrining_nilai5" json:"penilaian_skrining_nilai5"`
	PenilaianSkriningSkala6     string `form:"penilaian_skrining_skala6" json:"penilaian_skrining_skala6"`
	PenilaianSkriningNilai6     *int   `form:"penilaian_skrining_nilai6" json:"penilaian_skrining_nilai6"`
	PenilaianSkriningSkala7     string `form:"penilaian_skrining_skala7" json:"penilaian_skrining_skala7"`
	PenilaianSkriningNilai7     *int   `form:"penilaian_skrining_nilai7" json:"penilaian_skrining_nilai7"`
	PenilaianSkriningSkala8     string `form:"penilaian_skrining_skala8" json:"penilaian_skrining_skala8"`
	PenilaianSkriningNilai8     *int   `form:"penilaian_skrining_nilai8" json:"penilaian_skrining_nilai8"`
	PenilaianSkriningSkala9     string `form:"penilaian_skrining_skala9" json:"penilaian_skrining_skala9"`
	PenilaianSkriningNilai9     *int   `form:"penilaian_skrining_nilai9" json:"penilaian_skrining_nilai9"`
	PenilaianSkriningSkala10    string `form:"penilaian_skrining_skala10" json:"penilaian_skrining_skala10"`
	PenilaianSkriningNilai10    *int   `form:"penilaian_skrining_nilai10" json:"penilaian_skrining_nilai10"`
	PenilaianSkriningTotalnilai *int   `form:"penilaian_skrining_totalnilai" json:"penilaian_skrining_totalnilai"`
	Nip                         string `form:"nip" json:"nip"`
}

func penilaianLanjutanSkriningFungsionalRules() map[string]any {
	rules := map[string]any{
		"penilaian_skrining_skala1":     "in:Tak Terkendali/Tak Teratur (Perlu Pencahar),Kadang-kadang Tak Terkendali (1x Seminggu),Terkendali Teratur",
		"penilaian_skrining_nilai1":     "int",
		"penilaian_skrining_skala2":     "in:Tak Terkendali/Pakai Kateter,Kadang-kadang Tak Terkendali (Hanya 1x/24 Jam ),Mandiri",
		"penilaian_skrining_nilai2":     "int",
		"penilaian_skrining_skala3":     "in:Butuh Pertolongan Orang Lain,Mandiri",
		"penilaian_skrining_nilai3":     "int",
		"penilaian_skrining_skala4":     "in:Tergantung Pertolongan Orang Lain,Perlu Pertolongan Pada Beberapa Kegiatan Tetapi Dapat Mengerjakan Sendiri Beberapa Kegiatan Yang Lain,Mandiri",
		"penilaian_skrining_nilai4":     "int",
		"penilaian_skrining_skala5":     "in:Tidak Mampu,Perlu Ditolong Memotong Makanan,Mandiri",
		"penilaian_skrining_nilai5":     "int",
		"penilaian_skrining_skala6":     "in:Tidak Mampu,Perlu Banyak Bantuan Untuk Bisa Duduk (2 Orang),Bantuan Minimal 1 Orang,Mandiri",
		"penilaian_skrining_nilai6":     "int",
		"penilaian_skrining_skala7":     "in:Tidak Mampu,Bisa (Pindah) Dengan Kursi Roda,Berjalan Dengan Bantuan 1 Orang,Mandiri",
		"penilaian_skrining_nilai7":     "int",
		"penilaian_skrining_skala8":     "in:Tergantung Orang Lain,Sebagian Dibantu (Misal Mengancing Baju),Mandiri",
		"penilaian_skrining_nilai8":     "int",
		"penilaian_skrining_skala9":     "in:Tidak Mampu,Butuh Pertolongan,Mandiri",
		"penilaian_skrining_nilai9":     "int",
		"penilaian_skrining_skala10":    "in:Tergantung Orang Lain,Mandiri",
		"penilaian_skrining_nilai10":    "int",
		"penilaian_skrining_totalnilai": "int",
		"nip":                           "required|string|max_len:20",
	}
	return rules
}

// PenilaianLanjutanSkriningFungsionalStore simpan penilaian lanjutan skrining fungsional; kolom waktu kunci kosong = sekarang.
type PenilaianLanjutanSkriningFungsionalStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PenilaianLanjutanSkriningFungsionalData
}

func (r *PenilaianLanjutanSkriningFungsionalStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianLanjutanSkriningFungsionalStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianLanjutanSkriningFungsionalRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PenilaianLanjutanSkriningFungsionalStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PenilaianLanjutanSkriningFungsionalStore) Payload() PenilaianLanjutanSkriningFungsionalData {
	return r.PenilaianLanjutanSkriningFungsionalData
}

func (r *PenilaianLanjutanSkriningFungsionalStore) DetailValues() map[string][]string { return nil }

// PenilaianLanjutanSkriningFungsionalUpdate ubah penilaian lanjutan skrining fungsional (PUT); kunci lewat query string.
type PenilaianLanjutanSkriningFungsionalUpdate struct {
	PenilaianLanjutanSkriningFungsionalData
}

func (r *PenilaianLanjutanSkriningFungsionalUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianLanjutanSkriningFungsionalUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianLanjutanSkriningFungsionalRules()
}

func (r *PenilaianLanjutanSkriningFungsionalUpdate) Payload() PenilaianLanjutanSkriningFungsionalData {
	return r.PenilaianLanjutanSkriningFungsionalData
}

func (r *PenilaianLanjutanSkriningFungsionalUpdate) DetailValues() map[string][]string { return nil }
