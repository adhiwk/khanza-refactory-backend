package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPemantauanPewsDewasa pemantauan EWSD (RMPemantauanEWSD).
var FormPemantauanPewsDewasa = &Form[model.PemantauanPewsDewasa, request.PemantauanPewsDewasaData]{
	Slug:  "pemantauan-pews-dewasa",
	Label: "pemantauan EWSD",
	Spec: repo.Spec{
		Table:  "pemantauan_pews_dewasa",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PemantauanPewsDewasa, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PemantauanPewsDewasa) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PemantauanPewsDewasa) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PemantauanPewsDewasa) []string { return []string{derefStr(m.Nip)} },
	Fill: func(m *model.PemantauanPewsDewasa, d request.PemantauanPewsDewasaData) error {
		m.ParameterLajuRespirasi = support.Nullable(d.ParameterLajuRespirasi)
		m.SkorLajuRespirasi = support.Nullable(d.SkorLajuRespirasi)
		m.ParameterSaturasiOksigen = support.Nullable(d.ParameterSaturasiOksigen)
		m.SkorSaturasiOksigen = support.Nullable(d.SkorSaturasiOksigen)
		m.ParameterSuplemenOksigen = support.Nullable(d.ParameterSuplemenOksigen)
		m.SkorSuplemenOksigen = support.Nullable(d.SkorSuplemenOksigen)
		m.ParameterTekananDarahSistolik = support.Nullable(d.ParameterTekananDarahSistolik)
		m.SkorTekananDarahSistolik = support.Nullable(d.SkorTekananDarahSistolik)
		m.ParameterLajuJantung = support.Nullable(d.ParameterLajuJantung)
		m.SkorLajuJantung = support.Nullable(d.SkorLajuJantung)
		m.ParameterKesadaran = support.Nullable(d.ParameterKesadaran)
		m.SkorKesadaran = support.Nullable(d.SkorKesadaran)
		m.ParameterTemperatur = support.Nullable(d.ParameterTemperatur)
		m.SkorTemperatur = support.Nullable(d.SkorTemperatur)
		m.SkorTotal = support.Nullable(d.SkorTotal)
		m.ParameterTotal = support.Nullable(d.ParameterTotal)
		m.Nip = support.Nullable(d.Nip)
		return nil
	},
	Refs: []Ref[model.PemantauanPewsDewasa]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PemantauanPewsDewasa) any { return StrPtr(m.Nip) }},
	},
}
