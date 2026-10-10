package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormFollowUpDbd follow up DBD (RMDataFollowUpDBD).
var FormFollowUpDbd = &Form[model.FollowUpDbd, request.FollowUpDbdData]{
	Slug:  "follow-up-dbd",
	Label: "follow up DBD",
	Spec: repo.Spec{
		Table:  "follow_up_dbd",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.FollowUpDbd, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.TglPerawatan, err = KeyDate(key["tgl_perawatan"]); err != nil {
			return err
		}
		if m.JamRawat, err = KeyTime(key["jam_rawat"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.FollowUpDbd) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.FollowUpDbd) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.FollowUpDbd) []string { return []string{m.Nip} },
	Fill: func(m *model.FollowUpDbd, d request.FollowUpDbdData) error {
		m.Hemoglobin = support.Nullable(d.Hemoglobin)
		m.Hematokrit = strings.TrimSpace(d.Hematokrit)
		m.Leokosit = support.Nullable(d.Leokosit)
		m.Trombosit = support.Nullable(d.Trombosit)
		m.TerapiCairan = support.Nullable(d.TerapiCairan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.FollowUpDbd]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.FollowUpDbd) any { return Str(m.Nip) }},
	},
}
