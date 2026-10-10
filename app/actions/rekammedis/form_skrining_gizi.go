package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningGizi skrining gizi lanjut (RMDataSkriningGiziLanjut).
var FormSkriningGizi = &Form[model.SkriningGizi, request.SkriningGiziData]{
	Slug:  "skrining-gizi",
	Label: "skrining gizi lanjut",
	Spec: repo.Spec{
		Table:  "skrining_gizi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningGizi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.SkriningGizi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.SkriningGizi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningGizi) []string { return []string{derefStr(m.Nip)} },
	Fill: func(m *model.SkriningGizi, d request.SkriningGiziData) error {
		if err := Enum("parameter_imt", d.ParameterImt, "IMT > 20/z score > 2", "IMT 18,5-20/-2 =< z score =< 2", "IMT < 18,5/z score < -2"); err != nil {
			return err
		}
		m.SkriningBb = support.Nullable(d.SkriningBb)
		m.SkriningTb = support.Nullable(d.SkriningTb)
		m.Alergi = support.Nullable(d.Alergi)
		m.ParameterImt = support.Nullable(d.ParameterImt)
		m.SkorImt = support.Nullable(d.SkorImt)
		m.ParameterBb = support.Nullable(d.ParameterBb)
		m.SkorBb = support.Nullable(d.SkorBb)
		m.ParameterPenyakit = support.Nullable(d.ParameterPenyakit)
		m.SkorPenyakit = support.Nullable(d.SkorPenyakit)
		m.SkorTotal = support.Nullable(d.SkorTotal)
		m.ParameterTotal = support.Nullable(d.ParameterTotal)
		m.Nip = support.Nullable(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningGizi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningGizi) any { return StrPtr(m.Nip) }},
	},
}
