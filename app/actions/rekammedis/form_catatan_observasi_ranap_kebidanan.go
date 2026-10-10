package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanObservasiRanapKebidanan catatan observasi ranap kebidanan (RMDataCatatanObservasiRanapKebidanan).
var FormCatatanObservasiRanapKebidanan = &Form[model.CatatanObservasiRanapKebidanan, request.CatatanObservasiRanapKebidananData]{
	Slug:  "catatan-observasi-ranap-kebidanan",
	Label: "catatan observasi ranap kebidanan",
	Spec: repo.Spec{
		Table:  "catatan_observasi_ranap_kebidanan",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanObservasiRanapKebidanan, key repo.Key) error {
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
	KeyOf: func(m *model.CatatanObservasiRanapKebidanan) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.CatatanObservasiRanapKebidanan) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.CatatanObservasiRanapKebidanan) []string { return []string{m.Nip} },
	Fill: func(m *model.CatatanObservasiRanapKebidanan, d request.CatatanObservasiRanapKebidananData) error {
		m.Gcs = support.Nullable(d.Gcs)
		m.Td = strings.TrimSpace(d.Td)
		m.Hr = support.Nullable(d.Hr)
		m.Rr = support.Nullable(d.Rr)
		m.Suhu = support.Nullable(d.Suhu)
		m.Spo2 = strings.TrimSpace(d.Spo2)
		m.Kontraksi = strings.TrimSpace(d.Kontraksi)
		m.Bjj = strings.TrimSpace(d.Bjj)
		m.Ppv = strings.TrimSpace(d.Ppv)
		m.Vt = strings.TrimSpace(d.Vt)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanObservasiRanapKebidanan]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanObservasiRanapKebidanan) any { return Str(m.Nip) }},
	},
}
