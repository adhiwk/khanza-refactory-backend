package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// SkriningAnemiaData isian skrining anemia.
type SkriningAnemiaData struct {
	Tanggal            string `form:"tanggal" json:"tanggal"`
	MudahLelah         string `form:"mudah_lelah" json:"mudah_lelah"`
	BuahSayur          string `form:"buah_sayur" json:"buah_sayur"`
	ProteinHewani      string `form:"protein_hewani" json:"protein_hewani"`
	MasalahPubertas    string `form:"masalah_pubertas" json:"masalah_pubertas"`
	RisikoIms          string `form:"risiko_ims" json:"risiko_ims"`
	KekerasanSeksual   string `form:"kekerasan_seksual" json:"kekerasan_seksual"`
	SudahMenstruasi    string `form:"sudah_menstruasi" json:"sudah_menstruasi"`
	GangguanMenstruasi string `form:"gangguan_menstruasi" json:"gangguan_menstruasi"`
	TambahDarah        string `form:"tambah_darah" json:"tambah_darah"`
	KelainanDarah      string `form:"kelainan_darah" json:"kelainan_darah"`
	KeluargaThalasemia string `form:"keluarga_thalasemia" json:"keluarga_thalasemia"`
	Rambut             string `form:"rambut" json:"rambut"`
	Kulit              string `form:"kulit" json:"kulit"`
	BekasSutikan       string `form:"bekas_sutikan" json:"bekas_sutikan"`
	Kuku               string `form:"kuku" json:"kuku"`
	TandaKlinis        string `form:"tanda_klinis" json:"tanda_klinis"`
	PemeriksaanHb      string `form:"pemeriksaan_hb" json:"pemeriksaan_hb"`
	KadarHb            string `form:"kadar_hb" json:"kadar_hb"`
	JenisAnemia        string `form:"jenis_anemia" json:"jenis_anemia"`
	HasilSkrining      string `form:"hasil_skrining" json:"hasil_skrining"`
	Keterangan         string `form:"keterangan" json:"keterangan"`
	Nip                string `form:"nip" json:"nip"`
}

func skriningAnemiaRules() map[string]any {
	rules := map[string]any{
		"tanggal":             "required|date",
		"mudah_lelah":         "in:Tidak,Ya",
		"buah_sayur":          "in:Tidak,Ya",
		"protein_hewani":      "in:Tidak,Ya",
		"masalah_pubertas":    "in:Tidak,Ya",
		"risiko_ims":          "in:Tidak,Ya",
		"kekerasan_seksual":   "in:Tidak,Ya",
		"sudah_menstruasi":    "in:Tidak,Ya",
		"gangguan_menstruasi": "in:Tidak,Ya",
		"tambah_darah":        "in:Tidak,Ya",
		"kelainan_darah":      "in:Tidak,Ya",
		"keluarga_thalasemia": "in:Tidak,Ya",
		"rambut":              "in:Sehat,Tidak Sehat",
		"kulit":               "in:Sehat,Tidak Sehat",
		"bekas_sutikan":       "in:Tidak,Ya",
		"kuku":                "in:Sehat,Tidak Sehat",
		"tanda_klinis":        "required|in:Tidak,Ya",
		"pemeriksaan_hb":      "string|max_len:8",
		"kadar_hb":            "string",
		"jenis_anemia":        "in:Normal,Ringan,Sedang,Berat",
		"hasil_skrining":      "string|max_len:40",
		"keterangan":          "string|max_len:100",
		"nip":                 "required|string|max_len:20",
	}
	return rules
}

// SkriningAnemiaStore simpan skrining anemia; kolom waktu kunci kosong = sekarang.
type SkriningAnemiaStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	SkriningAnemiaData
}

func (r *SkriningAnemiaStore) Authorize(ctx http.Context) error { return nil }

func (r *SkriningAnemiaStore) Rules(ctx http.Context) map[string]any {
	rules := skriningAnemiaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *SkriningAnemiaStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *SkriningAnemiaStore) Payload() SkriningAnemiaData { return r.SkriningAnemiaData }

func (r *SkriningAnemiaStore) DetailValues() map[string][]string { return nil }

// SkriningAnemiaUpdate ubah skrining anemia (PUT); kunci lewat query string.
type SkriningAnemiaUpdate struct {
	SkriningAnemiaData
}

func (r *SkriningAnemiaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *SkriningAnemiaUpdate) Rules(ctx http.Context) map[string]any { return skriningAnemiaRules() }

func (r *SkriningAnemiaUpdate) Payload() SkriningAnemiaData { return r.SkriningAnemiaData }

func (r *SkriningAnemiaUpdate) DetailValues() map[string][]string { return nil }
