package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// PengkajianRestrainData isian pengkajian restrain.
type PengkajianRestrainData struct {
	Tanggal                          string `form:"tanggal" json:"tanggal"`
	Nip                              string `form:"nip" json:"nip"`
	Gcs                              string `form:"gcs" json:"gcs"`
	ReflekaCahayaKa                  string `form:"refleka_cahaya_ka" json:"refleka_cahaya_ka"`
	ReflekaCahayaKi                  string `form:"refleka_cahaya_ki" json:"refleka_cahaya_ki"`
	UkuranPupilKa                    string `form:"ukuran_pupil_ka" json:"ukuran_pupil_ka"`
	UkuranPupilKi                    string `form:"ukuran_pupil_ki" json:"ukuran_pupil_ki"`
	Td                               string `form:"td" json:"td"`
	Suhu                             string `form:"suhu" json:"suhu"`
	Rr                               string `form:"rr" json:"rr"`
	Nadi                             string `form:"nadi" json:"nadi"`
	HasilObservasi                   string `form:"hasil_observasi" json:"hasil_observasi"`
	PertimbanganKlinis               string `form:"pertimbangan_klinis" json:"pertimbangan_klinis"`
	RestrainNonFarmakologi           string `form:"restrain_non_farmakologi" json:"restrain_non_farmakologi"`
	RestrainNonFarmakologiKeterangan string `form:"restrain_non_farmakologi_keterangan" json:"restrain_non_farmakologi_keterangan"`
	RestrainFarmakologi              string `form:"restrain_farmakologi" json:"restrain_farmakologi"`
	SudahDijelaskanKeluarga          string `form:"sudah_dijelaskan_keluarga" json:"sudah_dijelaskan_keluarga"`
	KeluargaYangMenyetujui           string `form:"keluarga_yang_menyetujui" json:"keluarga_yang_menyetujui"`
}

func pengkajianRestrainRules() map[string]any {
	rules := map[string]any{
		"tanggal":                             "required|date",
		"nip":                                 "required|string|max_len:20",
		"gcs":                                 "string|max_len:5",
		"refleka_cahaya_ka":                   "string|max_len:3",
		"refleka_cahaya_ki":                   "string|max_len:3",
		"ukuran_pupil_ka":                     "string|max_len:3",
		"ukuran_pupil_ki":                     "string|max_len:3",
		"td":                                  "string|max_len:8",
		"suhu":                                "string|max_len:5",
		"rr":                                  "string|max_len:5",
		"nadi":                                "string|max_len:5",
		"hasil_observasi":                     "in:Pasien Gelisah/Delirium Dan Berontak,Pasien Tidak Kooperatif,Ketidakmampuan Dalam Mengikuti Perintah Untuk Tidak Meninggalkan Tempat Tidur",
		"pertimbangan_klinis":                 "in:Membahayakan Diri Sendiri,Membahayakan Orang Lain",
		"restrain_non_farmakologi":            "in:Restrain Pergelangan Tangan,Restrain Pergelangan Kaki,Restrain Badan,Lain-lain",
		"restrain_non_farmakologi_keterangan": "string|max_len:50",
		"restrain_farmakologi":                "string|max_len:200",
		"sudah_dijelaskan_keluarga":           "required|in:Sudah,Belum",
		"keluarga_yang_menyetujui":            "string|max_len:100",
	}
	return rules
}

// PengkajianRestrainStore simpan pengkajian restrain; kolom waktu kunci kosong = sekarang.
type PengkajianRestrainStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	PengkajianRestrainData
}

func (r *PengkajianRestrainStore) Authorize(ctx http.Context) error { return nil }

func (r *PengkajianRestrainStore) Rules(ctx http.Context) map[string]any {
	rules := pengkajianRestrainRules()
	rules["no_rawat"] = "required|string|max_len:17"
	return rules
}

func (r *PengkajianRestrainStore) KeyValues() repo.Key { return repo.Key{"no_rawat": r.NoRawat} }

func (r *PengkajianRestrainStore) Payload() PengkajianRestrainData { return r.PengkajianRestrainData }

func (r *PengkajianRestrainStore) DetailValues() map[string][]string { return nil }

// PengkajianRestrainUpdate ubah pengkajian restrain (PUT); kunci lewat query string.
type PengkajianRestrainUpdate struct {
	PengkajianRestrainData
}

func (r *PengkajianRestrainUpdate) Authorize(ctx http.Context) error { return nil }

func (r *PengkajianRestrainUpdate) Rules(ctx http.Context) map[string]any {
	return pengkajianRestrainRules()
}

func (r *PengkajianRestrainUpdate) Payload() PengkajianRestrainData { return r.PengkajianRestrainData }

func (r *PengkajianRestrainUpdate) DetailValues() map[string][]string { return nil }
