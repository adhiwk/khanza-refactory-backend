package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningInstrumenAcrs skrining instrumen ACRS (RMSkriningInstrumenACRS).
var FormSkriningInstrumenAcrs = &Form[model.SkriningInstrumenAcrs, request.SkriningInstrumenAcrsData]{
	Slug:  "skrining-instrumen-acrs",
	Label: "skrining instrumen ACRS",
	Spec: repo.Spec{
		Table:  "skrining_instrumen_acrs",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningInstrumenAcrs, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningInstrumenAcrs) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningInstrumenAcrs) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningInstrumenAcrs) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningInstrumenAcrs, d request.SkriningInstrumenAcrsData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.Pernyataanacrs1 = support.Nullable(d.Pernyataanacrs1)
		m.NilaiAcrs1 = d.NilaiAcrs1
		m.Pernyataanacrs2 = support.Nullable(d.Pernyataanacrs2)
		m.NilaiAcrs2 = d.NilaiAcrs2
		m.Pernyataanacrs3 = support.Nullable(d.Pernyataanacrs3)
		m.NilaiAcrs3 = d.NilaiAcrs3
		m.Pernyataanacrs4 = support.Nullable(d.Pernyataanacrs4)
		m.NilaiAcrs4 = d.NilaiAcrs4
		m.Pernyataanacrs5 = support.Nullable(d.Pernyataanacrs5)
		m.NilaiAcrs5 = d.NilaiAcrs5
		m.Pernyataanacrs6 = support.Nullable(d.Pernyataanacrs6)
		m.NilaiAcrs6 = d.NilaiAcrs6
		m.Pernyataanacrs7 = support.Nullable(d.Pernyataanacrs7)
		m.NilaiAcrs7 = d.NilaiAcrs7
		m.Pernyataanacrs8 = support.Nullable(d.Pernyataanacrs8)
		m.NilaiAcrs8 = d.NilaiAcrs8
		m.Pernyataanacrs9 = support.Nullable(d.Pernyataanacrs9)
		m.NilaiAcrs9 = d.NilaiAcrs9
		m.Pernyataanacrs10 = support.Nullable(d.Pernyataanacrs10)
		m.NilaiAcrs10 = d.NilaiAcrs10
		m.NilaiTotalAcrs = d.NilaiTotalAcrs
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		return nil
	},
	Refs: []Ref[model.SkriningInstrumenAcrs]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningInstrumenAcrs) any { return Str(m.Nip) }},
	},
}
