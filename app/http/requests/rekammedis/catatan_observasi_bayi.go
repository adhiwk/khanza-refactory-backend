package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// CatatanObservasiBayiData isian catatan observasi bayi.
type CatatanObservasiBayiData struct {
	Gcs           string `form:"gcs" json:"gcs"`
	Td            string `form:"td" json:"td"`
	Hr            string `form:"hr" json:"hr"`
	Rr            string `form:"rr" json:"rr"`
	Suhu          string `form:"suhu" json:"suhu"`
	Spo2          string `form:"spo2" json:"spo2"`
	Nch           string `form:"nch" json:"nch"`
	IkterikStatus string `form:"ikterik_status" json:"ikterik_status"`
	RetraksiDada  string `form:"retraksi_dada" json:"retraksi_dada"`
	OgtResidu     string `form:"ogt_residu" json:"ogt_residu"`
	AsiJumlah     string `form:"asi_jumlah" json:"asi_jumlah"`
	PasiJumlah    string `form:"pasi_jumlah" json:"pasi_jumlah"`
	BakStatus     string `form:"bak_status" json:"bak_status"`
	BabStatus     string `form:"bab_status" json:"bab_status"`
	Nip           string `form:"nip" json:"nip"`
}

func catatanObservasiBayiRules() map[string]any {
	rules := map[string]any{
		"gcs":            "string|max_len:10",
		"td":             "string|max_len:8",
		"hr":             "string|max_len:5",
		"rr":             "string|max_len:5",
		"suhu":           "string|max_len:5",
		"spo2":           "string|max_len:3",
		"nch":            "string|max_len:70",
		"ikterik_status": "string|max_len:30",
		"retraksi_dada":  "string|max_len:30",
		"ogt_residu":     "string|max_len:30",
		"asi_jumlah":     "string|max_len:30",
		"pasi_jumlah":    "string|max_len:30",
		"bak_status":     "string|max_len:30",
		"bab_status":     "string|max_len:30",
		"nip":            "string|max_len:20",
	}
	return rules
}

// CatatanObservasiBayiStore simpan catatan observasi bayi; kolom waktu kunci kosong = sekarang.
type CatatanObservasiBayiStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	CatatanObservasiBayiData
}

func (r *CatatanObservasiBayiStore) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiBayiStore) Rules(ctx http.Context) map[string]any {
	rules := catatanObservasiBayiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *CatatanObservasiBayiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *CatatanObservasiBayiStore) Payload() CatatanObservasiBayiData {
	return r.CatatanObservasiBayiData
}

func (r *CatatanObservasiBayiStore) DetailValues() map[string][]string { return nil }

// CatatanObservasiBayiUpdate ubah catatan observasi bayi (PUT); kunci lewat query string.
type CatatanObservasiBayiUpdate struct {
	CatatanObservasiBayiData
}

func (r *CatatanObservasiBayiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *CatatanObservasiBayiUpdate) Rules(ctx http.Context) map[string]any {
	return catatanObservasiBayiRules()
}

func (r *CatatanObservasiBayiUpdate) Payload() CatatanObservasiBayiData {
	return r.CatatanObservasiBayiData
}

func (r *CatatanObservasiBayiUpdate) DetailValues() map[string][]string { return nil }
