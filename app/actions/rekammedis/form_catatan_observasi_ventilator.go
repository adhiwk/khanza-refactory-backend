package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanObservasiVentilator catatan observasi ventilator (RMDataCatatanObservasiVentilator).
var FormCatatanObservasiVentilator = &Form[model.CatatanObservasiVentilator, request.CatatanObservasiVentilatorData]{
	Slug:  "catatan-observasi-ventilator",
	Label: "catatan observasi ventilator",
	Spec: repo.Spec{
		Table:  "catatan_observasi_ventilator",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanObservasiVentilator, key repo.Key) error {
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
	KeyOf: func(m *model.CatatanObservasiVentilator) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.CatatanObservasiVentilator) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.CatatanObservasiVentilator) []string { return []string{m.Nip} },
	Fill: func(m *model.CatatanObservasiVentilator, d request.CatatanObservasiVentilatorData) error {
		m.Mode = support.Nullable(d.Mode)
		m.Vt = support.Nullable(d.Vt)
		m.Pakar = support.Nullable(d.Pakar)
		m.Rr = support.Nullable(d.Rr)
		m.Reefps = support.Nullable(d.Reefps)
		m.Ee = support.Nullable(d.Ee)
		m.Keterangan = strings.TrimSpace(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanObservasiVentilator]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanObservasiVentilator) any { return Str(m.Nip) }},
	},
}
