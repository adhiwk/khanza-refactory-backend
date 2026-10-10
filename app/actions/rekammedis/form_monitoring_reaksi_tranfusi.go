package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormMonitoringReaksiTranfusi monitoring reaksi tranfusi (RMDataMonitoringReaksiTranfusi).
var FormMonitoringReaksiTranfusi = &Form[model.MonitoringReaksiTranfusi, request.MonitoringReaksiTranfusiData]{
	Slug:  "monitoring-reaksi-tranfusi",
	Label: "monitoring reaksi tranfusi",
	Spec: repo.Spec{
		Table:  "monitoring_reaksi_tranfusi",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.MonitoringReaksiTranfusi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.TglPerawatan, err = KeyDate(key["tgl_perawatan"]); err != nil {
			return err
		}
		if m.JamRawat, err = KeyTime(key["jam_rawat"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.MonitoringReaksiTranfusi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.MonitoringReaksiTranfusi) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.MonitoringReaksiTranfusi) []string { return []string{m.Nip} },
	Fill: func(m *model.MonitoringReaksiTranfusi, d request.MonitoringReaksiTranfusiData) error {
		m.ProdukDarah = support.Nullable(d.ProdukDarah)
		m.NoKantong = support.Nullable(d.NoKantong)
		m.LokasiInsersi = strings.TrimSpace(d.LokasiInsersi)
		m.Td = strings.TrimSpace(d.Td)
		m.Hr = support.Nullable(d.Hr)
		m.Rr = support.Nullable(d.Rr)
		m.Suhu = support.Nullable(d.Suhu)
		m.JenisReaksiAlergi = support.Nullable(d.JenisReaksiAlergi)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.MonitoringReaksiTranfusi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.MonitoringReaksiTranfusi) any { return Str(m.Nip) }},
	},
}
