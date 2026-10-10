package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanCekGds catatan cek GDS (RMDataCatatanCekGDS).
var FormCatatanCekGds = &Form[model.CatatanCekGds, request.CatatanCekGdsData]{
	Slug:  "catatan-cek-gds",
	Label: "catatan cek GDS",
	Spec: repo.Spec{
		Table:  "catatan_cek_gds",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanCekGds, key repo.Key) error {
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
	KeyOf: func(m *model.CatatanCekGds) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.CatatanCekGds) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.CatatanCekGds) []string { return []string{m.Nip} },
	Fill: func(m *model.CatatanCekGds, d request.CatatanCekGdsData) error {
		m.Gdp = support.Nullable(d.Gdp)
		m.Insulin = support.Nullable(d.Insulin)
		m.ObatGula = support.Nullable(d.ObatGula)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanCekGds]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanCekGds) any { return Str(m.Nip) }},
	},
}
