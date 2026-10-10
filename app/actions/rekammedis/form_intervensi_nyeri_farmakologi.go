package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormIntervensiNyeriFarmakologi intervensi nyeri farmakologi (RMDataIntervensiNyeriFarmakologi).
var FormIntervensiNyeriFarmakologi = &Form[model.IntervensiNyeriFarmakologi, request.IntervensiNyeriFarmakologiData]{
	Slug:  "intervensi-nyeri-farmakologi",
	Label: "intervensi nyeri farmakologi",
	Spec: repo.Spec{
		Table:  "intervensi_nyeri_farmakologi",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.IntervensiNyeriFarmakologi, key repo.Key) error {
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
	KeyOf: func(m *model.IntervensiNyeriFarmakologi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.IntervensiNyeriFarmakologi) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.IntervensiNyeriFarmakologi) []string { return []string{m.Nip} },
	Fill: func(m *model.IntervensiNyeriFarmakologi, d request.IntervensiNyeriFarmakologiData) error {
		m.NamaObat = support.Nullable(d.NamaObat)
		m.DosisEfek = support.Nullable(d.DosisEfek)
		m.Rute = support.Nullable(d.Rute)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.IntervensiNyeriFarmakologi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.IntervensiNyeriFarmakologi) any { return Str(m.Nip) }},
	},
}
