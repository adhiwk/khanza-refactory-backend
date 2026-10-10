package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanKeperawatanRalan catatan keperawatan ralan (RMDataCatatanKeperawatanRalan).
var FormCatatanKeperawatanRalan = &Form[model.CatatanKeperawatanRalan, request.CatatanKeperawatanRalanData]{
	Slug:  "catatan-keperawatan-ralan",
	Label: "catatan keperawatan ralan",
	Spec: repo.Spec{
		Table:  "catatan_keperawatan_ralan",
		Keys:   []string{"tanggal", "jam", "no_rawat"},
		Waktu:  "concat(tanggal,' ',jam)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanKeperawatanRalan, key repo.Key) error {
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
	KeyOf: func(m *model.CatatanKeperawatanRalan) repo.Key {
		return repo.Key{"tanggal": FmtDate(m.Tanggal), "jam": m.Jam, "no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.CatatanKeperawatanRalan) time.Time { return TglJam(m.Tanggal, m.Jam) },
	Petugas: func(m *model.CatatanKeperawatanRalan) []string { return []string{m.Nip} },
	Fill: func(m *model.CatatanKeperawatanRalan, d request.CatatanKeperawatanRalanData) error {
		m.Uraian = support.Nullable(d.Uraian)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanKeperawatanRalan]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanKeperawatanRalan) any { return Str(m.Nip) }},
	},
}
