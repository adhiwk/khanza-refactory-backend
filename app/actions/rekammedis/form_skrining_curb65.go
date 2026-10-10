package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningCurb65 skrining 65 (RMSkriningCURB65).
var FormSkriningCurb65 = &Form[model.SkriningCurb65, request.SkriningCurb65Data]{
	Slug:  "skrining-curb65",
	Label: "skrining 65",
	Spec: repo.Spec{
		Table:  "skrining_curb65",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningCurb65, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningCurb65) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningCurb65) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningCurb65) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningCurb65, d request.SkriningCurb65Data) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.Pernyataancurb651 = support.Nullable(d.Pernyataancurb651)
		m.NilaiCurb651 = d.NilaiCurb651
		m.Pernyataancurb652 = support.Nullable(d.Pernyataancurb652)
		m.NilaiCurb652 = d.NilaiCurb652
		m.Pernyataancurb653 = support.Nullable(d.Pernyataancurb653)
		m.NilaiCurb653 = d.NilaiCurb653
		m.Pernyataancurb654 = support.Nullable(d.Pernyataancurb654)
		m.NilaiCurb654 = d.NilaiCurb654
		m.Pernyataancurb655 = support.Nullable(d.Pernyataancurb655)
		m.NilaiCurb655 = d.NilaiCurb655
		m.NilaiTotalCurb65 = d.NilaiTotalCurb65
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		return nil
	},
	Refs: []Ref[model.SkriningCurb65]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningCurb65) any { return Str(m.Nip) }},
	},
}
