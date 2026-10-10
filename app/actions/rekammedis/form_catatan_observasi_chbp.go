package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanObservasiChbp catatan observasi CHBP (RMDataCatatanObservasiCHBP).
var FormCatatanObservasiChbp = &Form[model.CatatanObservasiChbp, request.CatatanObservasiChbpData]{
	Slug:  "catatan-observasi-chbp",
	Label: "catatan observasi CHBP",
	Spec: repo.Spec{
		Table:  "catatan_observasi_chbp",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanObservasiChbp, key repo.Key) error {
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
	KeyOf: func(m *model.CatatanObservasiChbp) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.CatatanObservasiChbp) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.CatatanObservasiChbp) []string { return []string{m.Nip} },
	Fill: func(m *model.CatatanObservasiChbp, d request.CatatanObservasiChbpData) error {
		m.Td = strings.TrimSpace(d.Td)
		m.Hr = support.Nullable(d.Hr)
		m.Suhu = support.Nullable(d.Suhu)
		m.Djj = strings.TrimSpace(d.Djj)
		m.His = strings.TrimSpace(d.His)
		m.Ppv = strings.TrimSpace(d.Ppv)
		m.Keterangan = strings.TrimSpace(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanObservasiChbp]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanObservasiChbp) any { return Str(m.Nip) }},
	},
}
