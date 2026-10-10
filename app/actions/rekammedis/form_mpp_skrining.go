package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormMppSkrining skrining MPP (RMSkriningMPP).
var FormMppSkrining = &Form[model.MppSkrining, request.MppSkriningData]{
	Slug:  "mpp-skrining",
	Label: "skrining MPP",
	Spec: repo.Spec{
		Table:  "mpp_skrining",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.MppSkrining, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDate(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.MppSkrining) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDate(m.Tanggal)}
	},
	Waktu:   func(m *model.MppSkrining) time.Time { return time.Time{} },
	Petugas: func(m *model.MppSkrining) []string { return []string{m.Nip} },
	Fill: func(m *model.MppSkrining, d request.MppSkriningData) error {
		m.Param1 = support.Nullable(d.Param1)
		m.Param2 = support.Nullable(d.Param2)
		m.Param3 = support.Nullable(d.Param3)
		m.Param4 = support.Nullable(d.Param4)
		m.Param5 = support.Nullable(d.Param5)
		m.Param6 = support.Nullable(d.Param6)
		m.Param7 = support.Nullable(d.Param7)
		m.Param8 = support.Nullable(d.Param8)
		m.Param9 = support.Nullable(d.Param9)
		m.Param10 = support.Nullable(d.Param10)
		m.Param11 = support.Nullable(d.Param11)
		m.Param12 = support.Nullable(d.Param12)
		m.Param13 = support.Nullable(d.Param13)
		m.Param14 = support.Nullable(d.Param14)
		m.Param15 = support.Nullable(d.Param15)
		m.Param16 = support.Nullable(d.Param16)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.MppSkrining]{
		{Column: "nip", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.MppSkrining) any { return Str(m.Nip) }},
	},
}
