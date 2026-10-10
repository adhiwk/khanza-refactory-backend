package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianPasienTerminalData isian penilaian pasien terminal.
type PenilaianPasienTerminalData struct {
	Tanggal                      string `form:"tanggal" json:"tanggal"`
	Diagnosa                     string `form:"diagnosa" json:"diagnosa"`
	Rps                          string `form:"rps" json:"rps"`
	Rpd                          string `form:"rpd" json:"rpd"`
	KeadaanUmum                  string `form:"keadaan_umum" json:"keadaan_umum"`
	Kesadaran                    string `form:"kesadaran" json:"kesadaran"`
	Td                           string `form:"td" json:"td"`
	Nadi                         string `form:"nadi" json:"nadi"`
	Suhu                         string `form:"suhu" json:"suhu"`
	Rr                           string `form:"rr" json:"rr"`
	Spo2                         string `form:"spo2" json:"spo2"`
	SkalaNyeri                   string `form:"skala_nyeri" json:"skala_nyeri"`
	TahapPasienMenjelangAjal     string `form:"tahap_pasien_menjelang_ajal" json:"tahap_pasien_menjelang_ajal"`
	TandaKlinisMenjelangKematian string `form:"tanda_klinis_menjelang_kematian" json:"tanda_klinis_menjelang_kematian"`
	KebutuhanSpiritualPasien     string `form:"kebutuhan_spiritual_pasien" json:"kebutuhan_spiritual_pasien"`
	Nip                          string `form:"nip" json:"nip"`
}

func penilaianPasienTerminalRules() map[string]any {
	rules := map[string]any{
		"tanggal":                         "required|date",
		"diagnosa":                        "string|max_len:500",
		"rps":                             "string|max_len:500",
		"rpd":                             "string|max_len:500",
		"keadaan_umum":                    "in:Sedang,Jelek,Sangat Jelek",
		"kesadaran":                       "in:Compos Mentis,Apatis,Delirium,Samnolen,Sopor,Koma",
		"td":                              "string|max_len:8",
		"nadi":                            "string|max_len:5",
		"suhu":                            "string|max_len:5",
		"rr":                              "string|max_len:5",
		"spo2":                            "string|max_len:5",
		"skala_nyeri":                     "required|in:0,1,2,3,4,5,6,7,8,9,10",
		"tahap_pasien_menjelang_ajal":     "in:Menolak,Marah,Menawar,Depresi,Menerima",
		"tanda_klinis_menjelang_kematian": "in:Kurang/Tidak Responsif,Nadi Cepat & Melemah,Pernapasan Tidak Teratur & Dangkal/Ngorok,Kulit Pucat,Ekstrimitas Dingin,Defekasi/Berkemih Tidak Sengaja,Mata Tidak Respon Cahaya,Penurunan Tonus Otot",
		"kebutuhan_spiritual_pasien":      "string|max_len:500",
		"nip":                             "required|string|max_len:20",
	}
	return rules
}

// PenilaianPasienTerminalStore simpan penilaian pasien terminal; kolom waktu kunci kosong = sekarang.
type PenilaianPasienTerminalStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianPasienTerminalData
}

func (r *PenilaianPasienTerminalStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPasienTerminalStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianPasienTerminalRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianPasienTerminalStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianPasienTerminalStore) Payload() PenilaianPasienTerminalData {
	return r.PenilaianPasienTerminalData
}

func (r *PenilaianPasienTerminalStore) DetailValues() map[string][]string { return nil }

// PenilaianPasienTerminalUpdate ubah penilaian pasien terminal (PUT); kunci lewat query string.
type PenilaianPasienTerminalUpdate struct {
	PenilaianPasienTerminalData
}

func (r *PenilaianPasienTerminalUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianPasienTerminalUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianPasienTerminalRules()
}

func (r *PenilaianPasienTerminalUpdate) Payload() PenilaianPasienTerminalData {
	return r.PenilaianPasienTerminalData
}

func (r *PenilaianPasienTerminalUpdate) DetailValues() map[string][]string { return nil }
