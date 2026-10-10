package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormIntervensiNyeriNonfarmakologi intervensi nyeri non farmakologi (RMDataIntervensiNyeriNonFarmakologi).
var FormIntervensiNyeriNonfarmakologi = &Form[model.IntervensiNyeriNonfarmakologi, request.IntervensiNyeriNonfarmakologiData]{
	Slug:  "intervensi-nyeri-nonfarmakologi",
	Label: "intervensi nyeri non farmakologi",
	Spec: repo.Spec{
		Table:  "intervensi_nyeri_nonfarmakologi",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.IntervensiNyeriNonfarmakologi, key repo.Key) error {
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
	KeyOf: func(m *model.IntervensiNyeriNonfarmakologi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.IntervensiNyeriNonfarmakologi) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.IntervensiNyeriNonfarmakologi) []string { return []string{m.Nip} },
	Fill: func(m *model.IntervensiNyeriNonfarmakologi, d request.IntervensiNyeriNonfarmakologiData) error {
		m.Intervensi = support.Nullable(d.Intervensi)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.IntervensiNyeriNonfarmakologi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.IntervensiNyeriNonfarmakologi) any { return Str(m.Nip) }},
	},
}
