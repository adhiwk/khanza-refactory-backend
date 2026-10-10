package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPemantauanPewsAnak pemantauan PEWS (RMPemantauanPEWS).
var FormPemantauanPewsAnak = &Form[model.PemantauanPewsAnak, request.PemantauanPewsAnakData]{
	Slug:  "pemantauan-pews-anak",
	Label: "pemantauan PEWS",
	Spec: repo.Spec{
		Table:  "pemantauan_pews_anak",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PemantauanPewsAnak, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PemantauanPewsAnak) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PemantauanPewsAnak) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PemantauanPewsAnak) []string { return []string{derefStr(m.Nip)} },
	Fill: func(m *model.PemantauanPewsAnak, d request.PemantauanPewsAnakData) error {
		m.ParameterPerilaku = support.Nullable(d.ParameterPerilaku)
		m.SkorPerilaku = support.Nullable(d.SkorPerilaku)
		m.ParameterCrtAtauWarnaKulit = support.Nullable(d.ParameterCrtAtauWarnaKulit)
		m.SkorCrtAtauWarnaKulit = support.Nullable(d.SkorCrtAtauWarnaKulit)
		m.ParameterPerespirasi = support.Nullable(d.ParameterPerespirasi)
		m.SkorPerespirasi = support.Nullable(d.SkorPerespirasi)
		m.SkorTotal = support.Nullable(d.SkorTotal)
		m.ParameterTotal = support.Nullable(d.ParameterTotal)
		m.Nip = support.Nullable(d.Nip)
		return nil
	},
	Refs: []Ref[model.PemantauanPewsAnak]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PemantauanPewsAnak) any { return StrPtr(m.Nip) }},
	},
}
