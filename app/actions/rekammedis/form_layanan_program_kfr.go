package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormLayananProgramKfr layanan program KFR (RMLayananProgramKFR).
var FormLayananProgramKfr = &Form[model.LayananProgramKfr, request.LayananProgramKfrData]{
	Slug:  "layanan-program-kfr",
	Label: "layanan program KFR",
	Spec: repo.Spec{
		Table:  "layanan_program_kfr",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.LayananProgramKfr, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.LayananProgramKfr) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.LayananProgramKfr) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.LayananProgramKfr) []string { return []string{m.Nip} },
	Fill: func(m *model.LayananProgramKfr, d request.LayananProgramKfrData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.NoRawatLayanan = strings.TrimSpace(d.NoRawatLayanan)
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.Program = support.Nullable(d.Program)
		return nil
	},
	Refs: []Ref[model.LayananProgramKfr]{
		{Column: "no_rawat_layanan", Table: "layanan_kedokteran_fisik_rehabilitasi", RefCol: "no_rawat", Label: "layanan kedokteran fisik rehabilitasi", Value: func(m *model.LayananProgramKfr) any { return Str(m.NoRawatLayanan) }},
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.LayananProgramKfr) any { return Str(m.Nip) }},
	},
}
