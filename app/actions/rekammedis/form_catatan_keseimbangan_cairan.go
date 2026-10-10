package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormCatatanKeseimbanganCairan catatan keseimbangan cairan (RMDataCatatanKeseimbanganCairan).
var FormCatatanKeseimbanganCairan = &Form[model.CatatanKeseimbanganCairan, request.CatatanKeseimbanganCairanData]{
	Slug:  "catatan-keseimbangan-cairan",
	Label: "catatan keseimbangan cairan",
	Spec: repo.Spec{
		Table:  "catatan_keseimbangan_cairan",
		Keys:   []string{"no_rawat", "tgl_perawatan", "jam_rawat"},
		Waktu:  "concat(tgl_perawatan,' ',jam_rawat)",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.CatatanKeseimbanganCairan, key repo.Key) error {
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
	KeyOf: func(m *model.CatatanKeseimbanganCairan) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tgl_perawatan": FmtDate(m.TglPerawatan), "jam_rawat": m.JamRawat}
	},
	Waktu:   func(m *model.CatatanKeseimbanganCairan) time.Time { return TglJam(m.TglPerawatan, m.JamRawat) },
	Petugas: func(m *model.CatatanKeseimbanganCairan) []string { return []string{m.Nip} },
	Fill: func(m *model.CatatanKeseimbanganCairan, d request.CatatanKeseimbanganCairanData) error {
		m.Infus = support.Nullable(d.Infus)
		m.Tranfusi = strings.TrimSpace(d.Tranfusi)
		m.Minum = support.Nullable(d.Minum)
		m.Urine = support.Nullable(d.Urine)
		m.Drain = support.Nullable(d.Drain)
		m.Ngt = strings.TrimSpace(d.Ngt)
		m.Iwl = strings.TrimSpace(d.Iwl)
		m.Keseimbangan = strings.TrimSpace(d.Keseimbangan)
		m.Keterangan = strings.TrimSpace(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.CatatanKeseimbanganCairan]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.CatatanKeseimbanganCairan) any { return Str(m.Nip) }},
	},
}
