package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanKeperawatanRanap catatan keperawatan ranap (RMDataCatatanKeperawatanRanap).
var FormCatatanKeperawatanRanap = &Form[model.CatatanKeperawatanRanap, request.CatatanKeperawatanRanapData]{
	Slug:  "catatan-keperawatan-ranap",
	Label: "catatan keperawatan ranap",
	Spec: repo.Spec{
		Table:  "catatan_keperawatan_ranap",
		Keys:   []string{"tanggal", "jam", "no_rawat"},
		Waktu:  "concat(tanggal,' ',jam)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanKeperawatanRanap, key repo.Key) error {
		var err error
		if m.Tanggal, err = KeyDate(key["tanggal"]); err != nil {
			return err
		}
		if m.Jam, err = KeyTime(key["jam"]); err != nil {
			return err
		}
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return err
	},
	KeyOf: func(m *model.CatatanKeperawatanRanap) repo.Key {
		return repo.Key{"tanggal": FmtDate(m.Tanggal), "jam": m.Jam, "no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.CatatanKeperawatanRanap) time.Time { return TglJam(m.Tanggal, m.Jam) },
	Petugas: func(m *model.CatatanKeperawatanRanap) []string { return []string{m.Nip} },
	Fill: func(m *model.CatatanKeperawatanRanap, d request.CatatanKeperawatanRanapData) error {
		m.Uraian = support.Nullable(d.Uraian)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanKeperawatanRanap]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanKeperawatanRanap) any { return Str(m.Nip) }},
	},
}
