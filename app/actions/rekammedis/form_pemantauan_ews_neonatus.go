package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPemantauanEwsNeonatus pemantauan EWS neonatus (RMPemantauanEWSNeonatus).
var FormPemantauanEwsNeonatus = &Form[model.PemantauanEwsNeonatus, request.PemantauanEwsNeonatusData]{
	Slug:  "pemantauan-ews-neonatus",
	Label: "pemantauan EWS neonatus",
	Spec: repo.Spec{
		Table:  "pemantauan_ews_neonatus",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PemantauanEwsNeonatus, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PemantauanEwsNeonatus) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PemantauanEwsNeonatus) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PemantauanEwsNeonatus) []string { return []string{derefStr(m.Nip)} },
	Fill: func(m *model.PemantauanEwsNeonatus, d request.PemantauanEwsNeonatusData) error {
		if err := Enum("parameter7", d.Parameter7, "<= 36,5", "36,5 - 37,5", ">= 37,5"); err != nil {
			return err
		}
		m.Parameter1 = support.Nullable(d.Parameter1)
		m.Skor1 = support.Nullable(d.Skor1)
		m.Parameter2 = support.Nullable(d.Parameter2)
		m.Skor2 = support.Nullable(d.Skor2)
		m.Parameter3 = support.Nullable(d.Parameter3)
		m.Skor3 = support.Nullable(d.Skor3)
		m.Parameter4 = support.Nullable(d.Parameter4)
		m.Skor4 = support.Nullable(d.Skor4)
		m.Parameter5 = support.Nullable(d.Parameter5)
		m.Skor5 = support.Nullable(d.Skor5)
		m.Parameter6 = support.Nullable(d.Parameter6)
		m.Skor6 = support.Nullable(d.Skor6)
		m.Parameter7 = support.Nullable(d.Parameter7)
		m.Skor7 = support.Nullable(d.Skor7)
		m.Parameter8 = support.Nullable(d.Parameter8)
		m.Skor8 = support.Nullable(d.Skor8)
		m.SkorTotal = support.Nullable(d.SkorTotal)
		m.ParameterTotal = support.Nullable(d.ParameterTotal)
		m.CodeBlue = strings.TrimSpace(d.CodeBlue)
		m.Nip = support.Nullable(d.Nip)
		return nil
	},
	Refs: []Ref[model.PemantauanEwsNeonatus]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PemantauanEwsNeonatus) any { return StrPtr(m.Nip) }},
	},
}
