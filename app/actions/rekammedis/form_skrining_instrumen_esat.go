package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningInstrumenEsat skrining instrumen ESAT (RMSkriningInstrumenESAT).
var FormSkriningInstrumenEsat = &Form[model.SkriningInstrumenEsat, request.SkriningInstrumenEsatData]{
	Slug:  "skrining-instrumen-esat",
	Label: "skrining instrumen ESAT",
	Spec: repo.Spec{
		Table:  "skrining_instrumen_esat",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningInstrumenEsat, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningInstrumenEsat) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningInstrumenEsat) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningInstrumenEsat) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningInstrumenEsat, d request.SkriningInstrumenEsatData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.Pernyataanesat1 = support.Nullable(d.Pernyataanesat1)
		m.Pernyataanesat2 = support.Nullable(d.Pernyataanesat2)
		m.Pernyataanesat3 = support.Nullable(d.Pernyataanesat3)
		m.Pernyataanesat4 = support.Nullable(d.Pernyataanesat4)
		m.Pernyataanesat5 = support.Nullable(d.Pernyataanesat5)
		m.Pernyataanesat6 = support.Nullable(d.Pernyataanesat6)
		m.Pernyataanesat7 = support.Nullable(d.Pernyataanesat7)
		m.Pernyataanesat8 = support.Nullable(d.Pernyataanesat8)
		m.Pernyataanesat9 = support.Nullable(d.Pernyataanesat9)
		m.Pernyataanesat10 = support.Nullable(d.Pernyataanesat10)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		m.Catatan = support.Nullable(d.Catatan)
		return nil
	},
	Refs: []Ref[model.SkriningInstrumenEsat]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningInstrumenEsat) any { return Str(m.Nip) }},
	},
}
