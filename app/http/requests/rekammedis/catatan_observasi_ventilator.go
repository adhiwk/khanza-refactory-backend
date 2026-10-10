package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanObservasiVentilatorData isian catatan observasi ventilator.
type CatatanObservasiVentilatorData struct {
	Mode       string `form:"mode" json:"mode"`
	Vt         string `form:"vt" json:"vt"`
	Pakar      string `form:"pakar" json:"pakar"`
	Rr         string `form:"rr" json:"rr"`
	Reefps     string `form:"reefps" json:"reefps"`
	Ee         string `form:"ee" json:"ee"`
	Keterangan string `form:"keterangan" json:"keterangan"`
	Nip        string `form:"nip" json:"nip"`
}

func catatanObservasiVentilatorRules() map[string]any {
	rules := map[string]any{
		"mode":       "in:CPAP,Nasal IMV,IMV,SIMV,A/C Atau SIPPV,PSV,Volume Guarantee,HFO,HFO + IMV",
		"vt":         "string|max_len:5",
		"pakar":      "string|max_len:30",
		"rr":         "string|max_len:5",
		"reefps":     "string|max_len:5",
		"ee":         "string|max_len:5",
		"keterangan": "string|max_len:200",
		"nip":        "required|string|max_len:20",
	}
	return rules
}

// CatatanObservasiVentilatorStore simpan catatan observasi ventilator; kolom waktu kunci kosong = sekarang.
type CatatanObservasiVentilatorStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	CatatanObservasiVentilatorData
}

func (r *CatatanObservasiVentilatorStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiVentilatorStore) Rules(ctx http.Context) map[string]any {
	rules := catatanObservasiVentilatorRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *CatatanObservasiVentilatorStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *CatatanObservasiVentilatorStore) Payload() CatatanObservasiVentilatorData {
	return r.CatatanObservasiVentilatorData
}

func (r *CatatanObservasiVentilatorStore) DetailValues() map[string][]string { return nil }

// CatatanObservasiVentilatorUpdate ubah catatan observasi ventilator (PUT); kunci lewat query string.
type CatatanObservasiVentilatorUpdate struct {
	CatatanObservasiVentilatorData
}

func (r *CatatanObservasiVentilatorUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiVentilatorUpdate) Rules(ctx http.Context) map[string]any {
	return catatanObservasiVentilatorRules()
}

func (r *CatatanObservasiVentilatorUpdate) Payload() CatatanObservasiVentilatorData {
	return r.CatatanObservasiVentilatorData
}

func (r *CatatanObservasiVentilatorUpdate) DetailValues() map[string][]string { return nil }
