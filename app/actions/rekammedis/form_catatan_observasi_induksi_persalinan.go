package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanObservasiInduksiPersalinan catatan observasi induksi persalinan (RMDataCatatanObservasiInduksiPersalinan).
var FormCatatanObservasiInduksiPersalinan = &Form[model.CatatanObservasiInduksiPersalinan, request.CatatanObservasiInduksiPersalinanData]{
	Slug:  "catatan-observasi-induksi-persalinan",
	Label: "catatan observasi induksi persalinan",
	Spec: repo.Spec{
		Table:  "catatan_observasi_induksi_persalinan",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanObservasiInduksiPersalinan, key repo.Key) error {
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
	KeyOf: func(m *model.CatatanObservasiInduksiPersalinan) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.CatatanObservasiInduksiPersalinan) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.CatatanObservasiInduksiPersalinan) []string { return []string{m.Nip} },
	Fill: func(m *model.CatatanObservasiInduksiPersalinan, d request.CatatanObservasiInduksiPersalinanData) error {
		m.Obat = support.Nullable(d.Obat)
		m.Cairan = strings.TrimSpace(d.Cairan)
		m.Dosis = support.Nullable(d.Dosis)
		m.His = support.Nullable(d.His)
		m.Djj = support.Nullable(d.Djj)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanObservasiInduksiPersalinan]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanObservasiInduksiPersalinan) any { return Str(m.Nip) }},
	},
}
