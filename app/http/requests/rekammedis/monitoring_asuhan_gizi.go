package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// MonitoringAsuhanGiziData isian monitoring asuhan gizi.
type MonitoringAsuhanGiziData struct {
	Monitoring string `form:"monitoring" json:"monitoring"`
	Evaluasi   string `form:"evaluasi" json:"evaluasi"`
	Nip        string `form:"nip" json:"nip"`
}

func monitoringAsuhanGiziRules() map[string]any {
	rules := map[string]any{
		"monitoring": "string|max_len:500",
		"evaluasi":   "string|max_len:500",
		"nip":        "string|max_len:20",
	}
	return rules
}

// MonitoringAsuhanGiziStore simpan monitoring asuhan gizi; kolom waktu kunci kosong = sekarang.
type MonitoringAsuhanGiziStore struct {
	NoRawat string `form:"no_rawat" json:"no_rawat"`
	Tanggal string `form:"tanggal" json:"tanggal"`
	MonitoringAsuhanGiziData
}

func (r *MonitoringAsuhanGiziStore) Authorize(ctx http.Context) error { return nil }

func (r *MonitoringAsuhanGiziStore) Rules(ctx http.Context) map[string]any {
	rules := monitoringAsuhanGiziRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tanggal"] = "date"
	return rules
}

func (r *MonitoringAsuhanGiziStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tanggal": r.Tanggal}
}

func (r *MonitoringAsuhanGiziStore) Payload() MonitoringAsuhanGiziData {
	return r.MonitoringAsuhanGiziData
}

func (r *MonitoringAsuhanGiziStore) DetailValues() map[string][]string { return nil }

// MonitoringAsuhanGiziUpdate ubah monitoring asuhan gizi (PUT); kunci lewat query string.
type MonitoringAsuhanGiziUpdate struct {
	MonitoringAsuhanGiziData
}

func (r *MonitoringAsuhanGiziUpdate) Authorize(ctx http.Context) error { return nil }

func (r *MonitoringAsuhanGiziUpdate) Rules(ctx http.Context) map[string]any {
	return monitoringAsuhanGiziRules()
}

func (r *MonitoringAsuhanGiziUpdate) Payload() MonitoringAsuhanGiziData {
	return r.MonitoringAsuhanGiziData
}

func (r *MonitoringAsuhanGiziUpdate) DetailValues() map[string][]string { return nil }
