package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanObservasiHemodialisaData isian catatan observasi hemodialisa.
type CatatanObservasiHemodialisaData struct {
	Qb            string `form:"qb" json:"qb"`
	Qd            string `form:"qd" json:"qd"`
	TekananArteri string `form:"tekanan_arteri" json:"tekanan_arteri"`
	TekananVena   string `form:"tekanan_vena" json:"tekanan_vena"`
	Tmp           string `form:"tmp" json:"tmp"`
	Ufr           string `form:"ufr" json:"ufr"`
	Tensi         string `form:"tensi" json:"tensi"`
	Nadi          string `form:"nadi" json:"nadi"`
	Suhu          string `form:"suhu" json:"suhu"`
	Spo2          string `form:"spo2" json:"spo2"`
	Tindakan      string `form:"tindakan" json:"tindakan"`
	Ufg           string `form:"ufg" json:"ufg"`
	BarcodeHf     string `form:"barcode_hf" json:"barcode_hf"`
	Nip           string `form:"nip" json:"nip"`
}

func catatanObservasiHemodialisaRules() map[string]any {
	rules := map[string]any{
		"qb":             "string|max_len:10",
		"qd":             "string|max_len:8",
		"tekanan_arteri": "string|max_len:5",
		"tekanan_vena":   "string|max_len:5",
		"tmp":            "string|max_len:5",
		"ufr":            "string|max_len:5",
		"tensi":          "string|max_len:8",
		"nadi":           "string|max_len:6",
		"suhu":           "string|max_len:5",
		"spo2":           "string|max_len:5",
		"tindakan":       "string|max_len:100",
		"ufg":            "string|max_len:10",
		"barcode_hf":     "string|max_len:50",
		"nip":            "required|string|max_len:20",
	}
	return rules
}

// CatatanObservasiHemodialisaStore simpan catatan observasi hemodialisa; kolom waktu kunci kosong = sekarang.
type CatatanObservasiHemodialisaStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	CatatanObservasiHemodialisaData
}

func (r *CatatanObservasiHemodialisaStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiHemodialisaStore) Rules(ctx http.Context) map[string]any {
	rules := catatanObservasiHemodialisaRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *CatatanObservasiHemodialisaStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *CatatanObservasiHemodialisaStore) Payload() CatatanObservasiHemodialisaData {
	return r.CatatanObservasiHemodialisaData
}

func (r *CatatanObservasiHemodialisaStore) DetailValues() map[string][]string { return nil }

// CatatanObservasiHemodialisaUpdate ubah catatan observasi hemodialisa (PUT); kunci lewat query string.
type CatatanObservasiHemodialisaUpdate struct {
	CatatanObservasiHemodialisaData
}

func (r *CatatanObservasiHemodialisaUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiHemodialisaUpdate) Rules(ctx http.Context) map[string]any {
	return catatanObservasiHemodialisaRules()
}

func (r *CatatanObservasiHemodialisaUpdate) Payload() CatatanObservasiHemodialisaData {
	return r.CatatanObservasiHemodialisaData
}

func (r *CatatanObservasiHemodialisaUpdate) DetailValues() map[string][]string { return nil }
