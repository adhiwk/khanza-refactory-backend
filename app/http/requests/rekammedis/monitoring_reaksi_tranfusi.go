package rekammedis

import (
	"github.com/goravel/framework/contracts/http"

	repo "goravel/app/repository/rekammedis"
)

// MonitoringReaksiTranfusiData isian monitoring reaksi tranfusi.
type MonitoringReaksiTranfusiData struct {
	ProdukDarah       string `form:"produk_darah" json:"produk_darah"`
	NoKantong         string `form:"no_kantong" json:"no_kantong"`
	LokasiInsersi     string `form:"lokasi_insersi" json:"lokasi_insersi"`
	Td                string `form:"td" json:"td"`
	Hr                string `form:"hr" json:"hr"`
	Rr                string `form:"rr" json:"rr"`
	Suhu              string `form:"suhu" json:"suhu"`
	JenisReaksiAlergi string `form:"jenis_reaksi_alergi" json:"jenis_reaksi_alergi"`
	Keterangan        string `form:"keterangan" json:"keterangan"`
	Nip               string `form:"nip" json:"nip"`
}

func monitoringReaksiTranfusiRules() map[string]any {
	rules := map[string]any{
		"produk_darah":        "string|max_len:40",
		"no_kantong":          "string|max_len:20",
		"lokasi_insersi":      "string|max_len:40",
		"td":                  "string|max_len:8",
		"hr":                  "string|max_len:5",
		"rr":                  "string|max_len:5",
		"suhu":                "string|max_len:5",
		"jenis_reaksi_alergi": "string|max_len:70",
		"keterangan":          "string|max_len:70",
		"nip":                 "required|string|max_len:20",
	}
	return rules
}

// MonitoringReaksiTranfusiStore simpan monitoring reaksi tranfusi; kolom waktu kunci kosong = sekarang.
type MonitoringReaksiTranfusiStore struct {
	NoRawat      string `form:"no_rawat" json:"no_rawat"`
	TglPerawatan string `form:"tgl_perawatan" json:"tgl_perawatan"`
	JamRawat     string `form:"jam_rawat" json:"jam_rawat"`
	MonitoringReaksiTranfusiData
}

func (r *MonitoringReaksiTranfusiStore) Authorize(ctx http.Context) error { return nil }

func (r *MonitoringReaksiTranfusiStore) Rules(ctx http.Context) map[string]any {
	rules := monitoringReaksiTranfusiRules()
	rules["no_rawat"] = "required|string|max_len:17"
	rules["tgl_perawatan"] = "date"
	rules["jam_rawat"] = "string|len:8"
	return rules
}

func (r *MonitoringReaksiTranfusiStore) KeyValues() repo.Key {
	return repo.Key{"no_rawat": r.NoRawat, "tgl_perawatan": r.TglPerawatan, "jam_rawat": r.JamRawat}
}

func (r *MonitoringReaksiTranfusiStore) Payload() MonitoringReaksiTranfusiData {
	return r.MonitoringReaksiTranfusiData
}

func (r *MonitoringReaksiTranfusiStore) DetailValues() map[string][]string { return nil }

// MonitoringReaksiTranfusiUpdate ubah monitoring reaksi tranfusi (PUT); kunci lewat query string.
type MonitoringReaksiTranfusiUpdate struct {
	MonitoringReaksiTranfusiData
}

func (r *MonitoringReaksiTranfusiUpdate) Authorize(ctx http.Context) error { return nil }

func (r *MonitoringReaksiTranfusiUpdate) Rules(ctx http.Context) map[string]any {
	return monitoringReaksiTranfusiRules()
}

func (r *MonitoringReaksiTranfusiUpdate) Payload() MonitoringReaksiTranfusiData {
	return r.MonitoringReaksiTranfusiData
}

func (r *MonitoringReaksiTranfusiUpdate) DetailValues() map[string][]string { return nil }
