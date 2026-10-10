package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PenilaianKorbanKekerasanData isian penilaian korban kekerasan.
type PenilaianKorbanKekerasanData struct {
	Tanggal                     string `form:"tanggal" json:"tanggal"`
	Informasi                   string `form:"informasi" json:"informasi"`
	HubunganDenganPasien        string `form:"hubungan_dengan_pasien" json:"hubungan_dengan_pasien"`
	JumlahSaudara               string `form:"jumlah_saudara" json:"jumlah_saudara"`
	KondisiKeluaga              string `form:"kondisi_keluaga" json:"kondisi_keluaga"`
	HubunganOrangTerdekat       string `form:"hubungan_orang_terdekat" json:"hubungan_orang_terdekat"`
	KekerasanYangDialami        string `form:"kekerasan_yang_dialami" json:"kekerasan_yang_dialami"`
	TempatKejadian              string `form:"tempat_kejadian" json:"tempat_kejadian"`
	LamaKekerasan               *int   `form:"lama_kekerasan" json:"lama_kekerasan"`
	PeriodeKekerasan            string `form:"periode_kekerasan" json:"periode_kekerasan"`
	SeberapaSeringMengalami     string `form:"seberapa_sering_mengalami" json:"seberapa_sering_mengalami"`
	PemicuKekerasan             string `form:"pemicu_kekerasan" json:"pemicu_kekerasan"`
	YangMelakukanKekerasan      string `form:"yang_melakukan_kekerasan" json:"yang_melakukan_kekerasan"`
	DampakKekerasan             string `form:"dampak_kekerasan" json:"dampak_kekerasan"`
	TandaTandaDidapatkan        string `form:"tanda_tanda_didapatkan" json:"tanda_tanda_didapatkan"`
	MemerlukanPendampingan      string `form:"memerlukan_pendampingan" json:"memerlukan_pendampingan"`
	RiwayatKelainan             string `form:"riwayat_kelainan" json:"riwayat_kelainan"`
	PemeriksaanKepala           string `form:"pemeriksaan_kepala" json:"pemeriksaan_kepala"`
	PemeriksaanThoraks          string `form:"pemeriksaan_thoraks" json:"pemeriksaan_thoraks"`
	PemeriksaanLeher            string `form:"pemeriksaan_leher" json:"pemeriksaan_leher"`
	PemeriksaanAbdomen          string `form:"pemeriksaan_abdomen" json:"pemeriksaan_abdomen"`
	PemeriksaanGenitalia        string `form:"pemeriksaan_genitalia" json:"pemeriksaan_genitalia"`
	PemeriksaanEkstrimitasAtas  string `form:"pemeriksaan_ekstrimitas_atas" json:"pemeriksaan_ekstrimitas_atas"`
	PemeriksaanEkstrimitasBawah string `form:"pemeriksaan_ekstrimitas_bawah" json:"pemeriksaan_ekstrimitas_bawah"`
	PemeriksaanAnus             string `form:"pemeriksaan_anus" json:"pemeriksaan_anus"`
	Nip                         string `form:"nip" json:"nip"`
}

func penilaianKorbanKekerasanRules() map[string]any {
	rules := map[string]any{
		"tanggal":                       "required|date",
		"informasi":                     "in:Autoanamnesis,Alloanamnesis",
		"hubungan_dengan_pasien":        "string|max_len:30",
		"jumlah_saudara":                "string|max_len:2",
		"kondisi_keluaga":               "in:Bahagia,Broken Home",
		"hubungan_orang_terdekat":       "in:Ada Masalah,Tidak Ada Masalah",
		"kekerasan_yang_dialami":        "string|max_len:350",
		"tempat_kejadian":               "string|max_len:40",
		"lama_kekerasan":                "int",
		"periode_kekerasan":             "in:Hari,Bulan,Tahun",
		"seberapa_sering_mengalami":     "string|max_len:150",
		"pemicu_kekerasan":              "string|max_len:150",
		"yang_melakukan_kekerasan":      "string|max_len:50",
		"dampak_kekerasan":              "string|max_len:200",
		"tanda_tanda_didapatkan":        "string|max_len:350",
		"memerlukan_pendampingan":       "in:Ya,Tidak",
		"riwayat_kelainan":              "string|max_len:50",
		"pemeriksaan_kepala":            "string|max_len:50",
		"pemeriksaan_thoraks":           "string|max_len:50",
		"pemeriksaan_leher":             "string|max_len:50",
		"pemeriksaan_abdomen":           "string|max_len:50",
		"pemeriksaan_genitalia":         "string|max_len:50",
		"pemeriksaan_ekstrimitas_atas":  "string|max_len:50",
		"pemeriksaan_ekstrimitas_bawah": "string|max_len:50",
		"pemeriksaan_anus":              "string|max_len:50",
		"nip":                           "required|string|max_len:20",
	}
	return rules
}

// PenilaianKorbanKekerasanStore simpan penilaian korban kekerasan; kolom waktu kunci kosong = sekarang.
type PenilaianKorbanKekerasanStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PenilaianKorbanKekerasanData
}

func (r *PenilaianKorbanKekerasanStore) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianKorbanKekerasanStore) Rules(ctx http.Context) map[string]any {
	rules := penilaianKorbanKekerasanRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PenilaianKorbanKekerasanStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PenilaianKorbanKekerasanStore) Payload() PenilaianKorbanKekerasanData {
	return r.PenilaianKorbanKekerasanData
}

func (r *PenilaianKorbanKekerasanStore) DetailValues() map[string][]string { return nil }

// PenilaianKorbanKekerasanUpdate ubah penilaian korban kekerasan (PUT); kunci lewat query string.
type PenilaianKorbanKekerasanUpdate struct {
	PenilaianKorbanKekerasanData
}

func (r *PenilaianKorbanKekerasanUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PenilaianKorbanKekerasanUpdate) Rules(ctx http.Context) map[string]any {
	return penilaianKorbanKekerasanRules()
}

func (r *PenilaianKorbanKekerasanUpdate) Payload() PenilaianKorbanKekerasanData {
	return r.PenilaianKorbanKekerasanData
}

func (r *PenilaianKorbanKekerasanUpdate) DetailValues() map[string][]string { return nil }
