package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningRisikoKankerParuData isian skrining risiko kanker paru.
type SkriningRisikoKankerParuData struct {
	Tanggal                                 string `form:"tanggal" json:"tanggal"`
	JenisKelamin                            string `form:"jenis_kelamin" json:"jenis_kelamin"`
	NilaiJenisKelamin                       string `form:"nilai_jenis_kelamin" json:"nilai_jenis_kelamin"`
	Umur                                    string `form:"umur" json:"umur"`
	NilaiUmur                               string `form:"nilai_umur" json:"nilai_umur"`
	PernahKanker                            string `form:"pernah_kanker" json:"pernah_kanker"`
	NilaiPernahKanker                       string `form:"nilai_pernah_kanker" json:"nilai_pernah_kanker"`
	AdaKeluargaKanker                       string `form:"ada_keluarga_kanker" json:"ada_keluarga_kanker"`
	NilaiAdaKeluargaKanker                  string `form:"nilai_ada_keluarga_kanker" json:"nilai_ada_keluarga_kanker"`
	RiwayatRokok                            string `form:"riwayat_rokok" json:"riwayat_rokok"`
	NilaiRiwayatRokok                       string `form:"nilai_riwayat_rokok" json:"nilai_riwayat_rokok"`
	RiwayatBekerjaMengandungKarsinogen      string `form:"riwayat_bekerja_mengandung_karsinogen" json:"riwayat_bekerja_mengandung_karsinogen"`
	NilaiRiwayatBekerjaMengandungKarsinogen string `form:"nilai_riwayat_bekerja_mengandung_karsinogen" json:"nilai_riwayat_bekerja_mengandung_karsinogen"`
	LingkunganTinggalPolusiTinggi           string `form:"lingkungan_tinggal_polusi_tinggi" json:"lingkungan_tinggal_polusi_tinggi"`
	NilaiLingkunganTinggalPolusiTinggi      string `form:"nilai_lingkungan_tinggal_polusi_tinggi" json:"nilai_lingkungan_tinggal_polusi_tinggi"`
	LingkunganRumahTidakSehat               string `form:"lingkungan_rumah_tidak_sehat" json:"lingkungan_rumah_tidak_sehat"`
	NilaiLingkunganRumahTidakSehat          string `form:"nilai_lingkungan_rumah_tidak_sehat" json:"nilai_lingkungan_rumah_tidak_sehat"`
	PernahParuKronik                        string `form:"pernah_paru_kronik" json:"pernah_paru_kronik"`
	NilaiPernahParuKronik                   string `form:"nilai_pernah_paru_kronik" json:"nilai_pernah_paru_kronik"`
	TotalSkor                               string `form:"total_skor" json:"total_skor"`
	HasilSkrining                           string `form:"hasil_skrining" json:"hasil_skrining"`
	Keterangan                              string `form:"keterangan" json:"keterangan"`
	Nip                                     string `form:"nip" json:"nip"`
}

func skriningRisikoKankerParuRules() map[string]any {
	rules := map[string]any{
		"tanggal":                               "required|date",
		"jenis_kelamin":                         "in:Laki-laki,Perempuan",
		"nilai_jenis_kelamin":                   "string|max_len:1",
		"umur":                                  "in:> 65 Tahun,45 - 65 Tahun,< 45 Tahun",
		"nilai_umur":                            "string|max_len:1",
		"pernah_kanker":                         "string",
		"nilai_pernah_kanker":                   "string|max_len:1",
		"ada_keluarga_kanker":                   "string",
		"nilai_ada_keluarga_kanker":             "string|max_len:1",
		"riwayat_rokok":                         "string",
		"nilai_riwayat_rokok":                   "string|max_len:1",
		"riwayat_bekerja_mengandung_karsinogen": "in:Ya,Tidak Yakin/Ragu-ragu,Tidak",
		"nilai_riwayat_bekerja_mengandung_karsinogen": "string|max_len:1",
		"lingkungan_tinggal_polusi_tinggi":            "in:Ya,Tidak Yakin/Ragu-ragu,Tidak",
		"nilai_lingkungan_tinggal_polusi_tinggi":      "string|max_len:1",
		"lingkungan_rumah_tidak_sehat":                "in:Ya,Tidak Yakin/Ragu-ragu,Tidak",
		"nilai_lingkungan_rumah_tidak_sehat":          "string|max_len:1",
		"pernah_paru_kronik":                          "string",
		"nilai_pernah_paru_kronik":                    "string|max_len:1",
		"total_skor":                                  "string|max_len:2",
		"hasil_skrining":                              "in:Risiko Ringan,Risiko Sedang,Risiko Berat",
		"keterangan":                                  "string|max_len:50",
		"nip":                                         "required|string|max_len:20",
	}
	return rules
}

// SkriningRisikoKankerParuStore simpan skrining risiko kanker paru; kolom waktu kunci kosong = sekarang.
type SkriningRisikoKankerParuStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningRisikoKankerParuData
}

func (r *SkriningRisikoKankerParuStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningRisikoKankerParuStore) Rules(ctx http.Context) map[string]any {
	rules := skriningRisikoKankerParuRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningRisikoKankerParuStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningRisikoKankerParuStore) Payload() SkriningRisikoKankerParuData {
	return r.SkriningRisikoKankerParuData
}

func (r *SkriningRisikoKankerParuStore) DetailValues() map[string][]string { return nil }

// SkriningRisikoKankerParuUpdate ubah skrining risiko kanker paru (PUT); kunci lewat query string.
type SkriningRisikoKankerParuUpdate struct {
	SkriningRisikoKankerParuData
}

func (r *SkriningRisikoKankerParuUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningRisikoKankerParuUpdate) Rules(ctx http.Context) map[string]any {
	return skriningRisikoKankerParuRules()
}

func (r *SkriningRisikoKankerParuUpdate) Payload() SkriningRisikoKankerParuData {
	return r.SkriningRisikoKankerParuData
}

func (r *SkriningRisikoKankerParuUpdate) DetailValues() map[string][]string { return nil }
