package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPemantauanMeowsObstetri pemantauan MEOWS (RMPemantauanMEOWS).
var FormPemantauanMeowsObstetri = &Form[model.PemantauanMeowsObstetri, request.PemantauanMeowsObstetriData]{
	Slug:  "pemantauan-meows-obstetri",
	Label: "pemantauan MEOWS",
	Spec: repo.Spec{
		Table:  "pemantauan_meows_obstetri",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PemantauanMeowsObstetri, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PemantauanMeowsObstetri) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PemantauanMeowsObstetri) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PemantauanMeowsObstetri) []string { return []string{derefStr(m.Nip)} },
	Fill: func(m *model.PemantauanMeowsObstetri, d request.PemantauanMeowsObstetriData) error {
		m.ParameterPernapasan = support.Nullable(d.ParameterPernapasan)
		m.SkorPernapasan = support.Nullable(d.SkorPernapasan)
		m.ParameterSaturasi = support.Nullable(d.ParameterSaturasi)
		m.SkorSaturasi = support.Nullable(d.SkorSaturasi)
		m.ParameterTemperatur = support.Nullable(d.ParameterTemperatur)
		m.SkorTemperatur = support.Nullable(d.SkorTemperatur)
		m.ParameterTekananDarahSistole = support.Nullable(d.ParameterTekananDarahSistole)
		m.SkorTekananDarahSistole = support.Nullable(d.SkorTekananDarahSistole)
		m.ParameterTekananDarahDiastole = support.Nullable(d.ParameterTekananDarahDiastole)
		m.SkorTekananDarahDiastole = support.Nullable(d.SkorTekananDarahDiastole)
		m.ParameterDenyutJantung = support.Nullable(d.ParameterDenyutJantung)
		m.SkorDenyutJantung = support.Nullable(d.SkorDenyutJantung)
		m.ParameterKesadaran = support.Nullable(d.ParameterKesadaran)
		m.SkorKesadaran = support.Nullable(d.SkorKesadaran)
		m.ParameterKetuban = support.Nullable(d.ParameterKetuban)
		m.SkorKetuban = support.Nullable(d.SkorKetuban)
		m.ParameterDischarge = support.Nullable(d.ParameterDischarge)
		m.SkorDischarge = support.Nullable(d.SkorDischarge)
		m.ParameterProteinuria = support.Nullable(d.ParameterProteinuria)
		m.SkorProteinuria = support.Nullable(d.SkorProteinuria)
		m.SkorTotal = strings.TrimSpace(d.SkorTotal)
		m.ParameterTotal = support.Nullable(d.ParameterTotal)
		m.CodeBlue = strings.TrimSpace(d.CodeBlue)
		m.Nip = support.Nullable(d.Nip)
		return nil
	},
	Refs: []Ref[model.PemantauanMeowsObstetri]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PemantauanMeowsObstetri) any { return StrPtr(m.Nip) }},
	},
}
