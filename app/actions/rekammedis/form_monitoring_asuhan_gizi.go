package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormMonitoringAsuhanGizi monitoring asuhan gizi (RMDataMonitoringAsuhanGizi).
var FormMonitoringAsuhanGizi = &Form[model.MonitoringAsuhanGizi, request.MonitoringAsuhanGiziData]{
	Slug:  "monitoring-asuhan-gizi",
	Label: "monitoring asuhan gizi",
	Spec: repo.Spec{
		Table:  "monitoring_asuhan_gizi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.MonitoringAsuhanGizi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.MonitoringAsuhanGizi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.MonitoringAsuhanGizi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.MonitoringAsuhanGizi) []string { return []string{derefStr(m.Nip)} },
	Fill: func(m *model.MonitoringAsuhanGizi, d request.MonitoringAsuhanGiziData) error {
		m.Monitoring = support.Nullable(d.Monitoring)
		m.Evaluasi = support.Nullable(d.Evaluasi)
		m.Nip = support.Nullable(d.Nip)
		return nil
	},
	Refs: []Ref[model.MonitoringAsuhanGizi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.MonitoringAsuhanGizi) any { return StrPtr(m.Nip) }},
	},
}
