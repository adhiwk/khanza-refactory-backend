package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianLanjutanResikoJatuhPsikiatriData isian penilaian lanjutan risiko jatuh psikiatri.
type PenilaianLanjutanResikoJatuhPsikiatriData struct {
	PenilaianJatuhedmonsonSkala1     string `form:"penilaian_jatuhedmonson_skala1" json:"penilaian_jatuhedmonson_skala1"`
	PenilaianJatuhedmonsonNilai1     *int   `form:"penilaian_jatuhedmonson_nilai1" json:"penilaian_jatuhedmonson_nilai1"`
	PenilaianJatuhedmonsonSkala2     string `form:"penilaian_jatuhedmonson_skala2" json:"penilaian_jatuhedmonson_skala2"`
	PenilaianJatuhedmonsonNilai2     *int   `form:"penilaian_jatuhedmonson_nilai2" json:"penilaian_jatuhedmonson_nilai2"`
	PenilaianJatuhedmonsonSkala3     string `form:"penilaian_jatuhedmonson_skala3" json:"penilaian_jatuhedmonson_skala3"`
	PenilaianJatuhedmonsonNilai3     *int   `form:"penilaian_jatuhedmonson_nilai3" json:"penilaian_jatuhedmonson_nilai3"`
	PenilaianJatuhedmonsonSkala4     string `form:"penilaian_jatuhedmonson_skala4" json:"penilaian_jatuhedmonson_skala4"`
	PenilaianJatuhedmonsonNilai4     *int   `form:"penilaian_jatuhedmonson_nilai4" json:"penilaian_jatuhedmonson_nilai4"`
	PenilaianJatuhedmonsonSkala5     string `form:"penilaian_jatuhedmonson_skala5" json:"penilaian_jatuhedmonson_skala5"`
	PenilaianJatuhedmonsonNilai5     *int   `form:"penilaian_jatuhedmonson_nilai5" json:"penilaian_jatuhedmonson_nilai5"`
	PenilaianJatuhedmonsonSkala6     string `form:"penilaian_jatuhedmonson_skala6" json:"penilaian_jatuhedmonson_skala6"`
	PenilaianJatuhedmonsonNilai6     *int   `form:"penilaian_jatuhedmonson_nilai6" json:"penilaian_jatuhedmonson_nilai6"`
	PenilaianJatuhedmonsonTotalnilai *int   `form:"penilaian_jatuhedmonson_totalnilai" json:"penilaian_jatuhedmonson_totalnilai"`
	HasilSkrining                    string `form:"hasil_skrining" json:"hasil_skrining"`
	Saran                            string `form:"saran" json:"saran"`
	Nip                              string `form:"nip" json:"nip"`
}

func penilaianLanjutanResikoJatuhPsikiatriRules() map[string]any {
	rules := map[string]any{
		"penilaian_jatuhedmonson_skala1":     "in:-,Kurang Dari 50 Th,50 - 70 Th,Lebih Dari 70 Th",
		"penilaian_jatuhedmonson_nilai1":     "int",
		"penilaian_jatuhedmonson_skala2":     "in:-,Kesadaran/Orientasi Baik Setiap Saat,Agitasi/Ansietas,Kadang-kadang Bingung,Bingung/Disorientasi",
		"penilaian_jatuhedmonson_nilai2":     "int",
		"penilaian_jatuhedmonson_skala3":     "in:-,Mandiri & Mampu Mengontrol BAB/BAK,Dower Catheter/Colostomy,Eliminasi Dengan Bantuan,Gangguan Eliminasi (Inkontinensia/Nokturia/Frekuensi),Inkontinensia Tetapi Mampu Untuk Mobilisasi",
		"penilaian_jatuhedmonson_nilai3":     "int",
		"penilaian_jatuhedmonson_skala4":     "in:-,Tanpa Obat-obatan,Obat-obatan Jantung,Obat-obatan Psikotropika (Termasuk Benzodiazepine & Antidepresan),Mendapat Tambahan Obat-obatan & Atau Obat PRN Selama 24 Jam Terakhir",
		"penilaian_jatuhedmonson_nilai4":     "int",
		"penilaian_jatuhedmonson_skala5":     "string",
		"penilaian_jatuhedmonson_nilai5":     "int",
		"penilaian_jatuhedmonson_skala6":     "string",
		"penilaian_jatuhedmonson_nilai6":     "int",
		"penilaian_jatuhedmonson_totalnilai": "int",
		"hasil_skrining":                     "string|max_len:200",
		"saran":                              "string|max_len:200",
		"nip":                                "required|string|max_len:20",
	}
	return rules
}

// PenilaianLanjutanResikoJatuhPsikiatriStore simpan penilaian lanjutan risiko jatuh psikiatri; kolom waktu kunci kosong = sekarang.
type PenilaianLanjutanResikoJatuhPsikiatriStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PenilaianLanjutanResikoJatuhPsikiatriData
}

func (r *PenilaianLanjutanResikoJatuhPsikiatriStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianLanjutanResikoJatuhPsikiatriStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianLanjutanResikoJatuhPsikiatriRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PenilaianLanjutanResikoJatuhPsikiatriStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PenilaianLanjutanResikoJatuhPsikiatriStore) Payload() PenilaianLanjutanResikoJatuhPsikiatriData {
	return r.PenilaianLanjutanResikoJatuhPsikiatriData
}

func (r *PenilaianLanjutanResikoJatuhPsikiatriStore) DetailValues() map[string][]string { return nil }

// PenilaianLanjutanResikoJatuhPsikiatriUpdate ubah penilaian lanjutan risiko jatuh psikiatri (PUT); kunci lewat query string.
type PenilaianLanjutanResikoJatuhPsikiatriUpdate struct {
	PenilaianLanjutanResikoJatuhPsikiatriData
}

func (r *PenilaianLanjutanResikoJatuhPsikiatriUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianLanjutanResikoJatuhPsikiatriUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianLanjutanResikoJatuhPsikiatriRules()
}

func (r *PenilaianLanjutanResikoJatuhPsikiatriUpdate) Payload() PenilaianLanjutanResikoJatuhPsikiatriData {
	return r.PenilaianLanjutanResikoJatuhPsikiatriData
}

func (r *PenilaianLanjutanResikoJatuhPsikiatriUpdate) DetailValues() map[string][]string { return nil }
