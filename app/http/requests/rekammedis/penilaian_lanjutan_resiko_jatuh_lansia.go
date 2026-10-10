package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianLanjutanResikoJatuhLansiaData isian penilaian lanjutan risiko jatuh lansia.
type PenilaianLanjutanResikoJatuhLansiaData struct {
	PenilaianJatuhmorseSkala1     string `form:"penilaian_jatuhmorse_skala1" json:"penilaian_jatuhmorse_skala1"`
	PenilaianJatuhmorseNilai1     *int   `form:"penilaian_jatuhmorse_nilai1" json:"penilaian_jatuhmorse_nilai1"`
	PenilaianJatuhmorseSkala2     string `form:"penilaian_jatuhmorse_skala2" json:"penilaian_jatuhmorse_skala2"`
	PenilaianJatuhmorseNilai2     *int   `form:"penilaian_jatuhmorse_nilai2" json:"penilaian_jatuhmorse_nilai2"`
	PenilaianJatuhmorseSkala3     string `form:"penilaian_jatuhmorse_skala3" json:"penilaian_jatuhmorse_skala3"`
	PenilaianJatuhmorseNilai3     *int   `form:"penilaian_jatuhmorse_nilai3" json:"penilaian_jatuhmorse_nilai3"`
	PenilaianJatuhmorseSkala4     string `form:"penilaian_jatuhmorse_skala4" json:"penilaian_jatuhmorse_skala4"`
	PenilaianJatuhmorseNilai4     *int   `form:"penilaian_jatuhmorse_nilai4" json:"penilaian_jatuhmorse_nilai4"`
	PenilaianJatuhmorseSkala5     string `form:"penilaian_jatuhmorse_skala5" json:"penilaian_jatuhmorse_skala5"`
	PenilaianJatuhmorseNilai5     *int   `form:"penilaian_jatuhmorse_nilai5" json:"penilaian_jatuhmorse_nilai5"`
	PenilaianJatuhmorseSkala6     string `form:"penilaian_jatuhmorse_skala6" json:"penilaian_jatuhmorse_skala6"`
	PenilaianJatuhmorseNilai6     *int   `form:"penilaian_jatuhmorse_nilai6" json:"penilaian_jatuhmorse_nilai6"`
	PenilaianJatuhmorseTotalnilai *int   `form:"penilaian_jatuhmorse_totalnilai" json:"penilaian_jatuhmorse_totalnilai"`
	HasilSkrining                 string `form:"hasil_skrining" json:"hasil_skrining"`
	Saran                         string `form:"saran" json:"saran"`
	Nip                           string `form:"nip" json:"nip"`
}

func penilaianLanjutanResikoJatuhLansiaRules() map[string]any {
	rules := map[string]any{
		"penilaian_jatuhmorse_skala1":     "in:Tidak,Pasien Datang Karena Jatuh,Pasien Jatuh Dalam 2 Bulan Terakhir",
		"penilaian_jatuhmorse_nilai1":     "int",
		"penilaian_jatuhmorse_skala2":     "in:Tidak,Pasien Delirium,Pasien Disorientasi,Pasien Agitasi",
		"penilaian_jatuhmorse_nilai2":     "int",
		"penilaian_jatuhmorse_skala3":     "in:Tidak,Memakai Kaca Mata,Penglihatan Kabur,Memiliki Glukoma/Katarak/Degenerasi Makula",
		"penilaian_jatuhmorse_nilai3":     "int",
		"penilaian_jatuhmorse_skala4":     "in:Tidak,Prilaku Berkemih/Frekuensi/Urgensi/Incontinensia/Nokturia",
		"penilaian_jatuhmorse_nilai4":     "int",
		"penilaian_jatuhmorse_skala5":     "in:Mandiri,Memerlukan Bantuan 1 Orang/Pengawasan,Memerlukan Bantuan 2 Orang,Memerlukan Bantuan Total",
		"penilaian_jatuhmorse_nilai5":     "int",
		"penilaian_jatuhmorse_skala6":     "in:Mandiri,Berjalan Dengan Bantuan 1 Orang,Menggunakan Kursi Roda,Imobilisasi",
		"penilaian_jatuhmorse_nilai6":     "int",
		"penilaian_jatuhmorse_totalnilai": "int",
		"hasil_skrining":                  "string|max_len:200",
		"saran":                           "string|max_len:200",
		"nip":                             "required|string|max_len:20",
	}
	return rules
}

// PenilaianLanjutanResikoJatuhLansiaStore simpan penilaian lanjutan risiko jatuh lansia; kolom waktu kunci kosong = sekarang.
type PenilaianLanjutanResikoJatuhLansiaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	PenilaianLanjutanResikoJatuhLansiaData
}

func (r *PenilaianLanjutanResikoJatuhLansiaStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianLanjutanResikoJatuhLansiaStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianLanjutanResikoJatuhLansiaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *PenilaianLanjutanResikoJatuhLansiaStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *PenilaianLanjutanResikoJatuhLansiaStore) Payload() PenilaianLanjutanResikoJatuhLansiaData {
	return r.PenilaianLanjutanResikoJatuhLansiaData
}

func (r *PenilaianLanjutanResikoJatuhLansiaStore) DetailValues() map[string][]string { return nil }

// PenilaianLanjutanResikoJatuhLansiaUpdate ubah penilaian lanjutan risiko jatuh lansia (PUT); kunci lewat query string.
type PenilaianLanjutanResikoJatuhLansiaUpdate struct {
	PenilaianLanjutanResikoJatuhLansiaData
}

func (r *PenilaianLanjutanResikoJatuhLansiaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianLanjutanResikoJatuhLansiaUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianLanjutanResikoJatuhLansiaRules()
}

func (r *PenilaianLanjutanResikoJatuhLansiaUpdate) Payload() PenilaianLanjutanResikoJatuhLansiaData {
	return r.PenilaianLanjutanResikoJatuhLansiaData
}

func (r *PenilaianLanjutanResikoJatuhLansiaUpdate) DetailValues() map[string][]string { return nil }
