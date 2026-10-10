package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianLanjutanResikoJatuhAnakData isian penilaian lanjutan risiko jatuh anak.
type PenilaianLanjutanResikoJatuhAnakData struct {
	PenilaianHumptydumptySkala1     string `form:"penilaian_humptydumpty_skala1" json:"penilaian_humptydumpty_skala1"`
	PenilaianHumptydumptyNilai1     *int   `form:"penilaian_humptydumpty_nilai1" json:"penilaian_humptydumpty_nilai1"`
	PenilaianHumptydumptySkala2     string `form:"penilaian_humptydumpty_skala2" json:"penilaian_humptydumpty_skala2"`
	PenilaianHumptydumptyNilai2     *int   `form:"penilaian_humptydumpty_nilai2" json:"penilaian_humptydumpty_nilai2"`
	PenilaianHumptydumptySkala3     string `form:"penilaian_humptydumpty_skala3" json:"penilaian_humptydumpty_skala3"`
	PenilaianHumptydumptyNilai3     *int   `form:"penilaian_humptydumpty_nilai3" json:"penilaian_humptydumpty_nilai3"`
	PenilaianHumptydumptySkala4     string `form:"penilaian_humptydumpty_skala4" json:"penilaian_humptydumpty_skala4"`
	PenilaianHumptydumptyNilai4     *int   `form:"penilaian_humptydumpty_nilai4" json:"penilaian_humptydumpty_nilai4"`
	PenilaianHumptydumptySkala5     string `form:"penilaian_humptydumpty_skala5" json:"penilaian_humptydumpty_skala5"`
	PenilaianHumptydumptyNilai5     *int   `form:"penilaian_humptydumpty_nilai5" json:"penilaian_humptydumpty_nilai5"`
	PenilaianHumptydumptySkala6     string `form:"penilaian_humptydumpty_skala6" json:"penilaian_humptydumpty_skala6"`
	PenilaianHumptydumptyNilai6     *int   `form:"penilaian_humptydumpty_nilai6" json:"penilaian_humptydumpty_nilai6"`
	PenilaianHumptydumptySkala7     string `form:"penilaian_humptydumpty_skala7" json:"penilaian_humptydumpty_skala7"`
	PenilaianHumptydumptyNilai7     *int   `form:"penilaian_humptydumpty_nilai7" json:"penilaian_humptydumpty_nilai7"`
	PenilaianHumptydumptyTotalnilai *int   `form:"penilaian_humptydumpty_totalnilai" json:"penilaian_humptydumpty_totalnilai"`
	HasilSkrining                   string `form:"hasil_skrining" json:"hasil_skrining"`
	Saran                           string `form:"saran" json:"saran"`
	Nip                             string `form:"nip" json:"nip"`
}

func penilaianLanjutanResikoJatuhAnakRules() map[string]any {
	rules := map[string]any{
		"penilaian_humptydumpty_skala1":     "in:0 - 3 Tahun,3 - 7 Tahun,7 - 13 Tahun,> 13 Tahun",
		"penilaian_humptydumpty_nilai1":     "int",
		"penilaian_humptydumpty_skala2":     "in:Laki-laki,Perempuan",
		"penilaian_humptydumpty_nilai2":     "int",
		"penilaian_humptydumpty_skala3":     "string",
		"penilaian_humptydumpty_nilai3":     "int",
		"penilaian_humptydumpty_skala4":     "in:Tidak Sadar Terhadap Keterbatasan,Lupa Keterbatasan,Mengetahui Kemampuan Diri",
		"penilaian_humptydumpty_nilai4":     "int",
		"penilaian_humptydumpty_skala5":     "in:Riwayat Jatuh Dari Tempat Tidur Saat Bayi/Anak,Pasien Menggunakan Alat Bantu/Box/Mebel,Pasien Berada Di Tempat Tidur,Di Luar Ruang Rawat",
		"penilaian_humptydumpty_nilai5":     "int",
		"penilaian_humptydumpty_skala6":     "in:Dalam 24 Jam,Dalam 48 Jam,> 48 Jam",
		"penilaian_humptydumpty_nilai6":     "int",
		"penilaian_humptydumpty_skala7":     "string",
		"penilaian_humptydumpty_nilai7":     "int",
		"penilaian_humptydumpty_totalnilai": "int",
		"hasil_skrining":                    "string|max_len:200",
		"saran":                             "string|max_len:200",
		"nip":                               "required|string|max_len:20",
	}
	return rules
}

// PenilaianLanjutanResikoJatuhAnakStore simpan penilaian lanjutan risiko jatuh anak; kolom waktu kunci kosong = sekarang.
type PenilaianLanjutanResikoJatuhAnakStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PenilaianLanjutanResikoJatuhAnakData
}

func (r *PenilaianLanjutanResikoJatuhAnakStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianLanjutanResikoJatuhAnakStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianLanjutanResikoJatuhAnakRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PenilaianLanjutanResikoJatuhAnakStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PenilaianLanjutanResikoJatuhAnakStore) Payload() PenilaianLanjutanResikoJatuhAnakData {
	return r.PenilaianLanjutanResikoJatuhAnakData
}

func (r *PenilaianLanjutanResikoJatuhAnakStore) DetailValues() map[string][]string { return nil }

// PenilaianLanjutanResikoJatuhAnakUpdate ubah penilaian lanjutan risiko jatuh anak (PUT); kunci lewat query string.
type PenilaianLanjutanResikoJatuhAnakUpdate struct {
	PenilaianLanjutanResikoJatuhAnakData
}

func (r *PenilaianLanjutanResikoJatuhAnakUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianLanjutanResikoJatuhAnakUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianLanjutanResikoJatuhAnakRules()
}

func (r *PenilaianLanjutanResikoJatuhAnakUpdate) Payload() PenilaianLanjutanResikoJatuhAnakData {
	return r.PenilaianLanjutanResikoJatuhAnakData
}

func (r *PenilaianLanjutanResikoJatuhAnakUpdate) DetailValues() map[string][]string { return nil }
