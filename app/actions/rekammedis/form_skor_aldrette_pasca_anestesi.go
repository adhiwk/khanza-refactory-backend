package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkorAldrettePascaAnestesi monitoring aldrette pasca anestesi (RMMonitoringAldrettePascaAnestesi).
var FormSkorAldrettePascaAnestesi = &Form[model.SkorAldrettePascaAnestesi, request.SkorAldrettePascaAnestesiData]{
	Slug:  "skor-aldrette-pasca-anestesi",
	Label: "monitoring aldrette pasca anestesi",
	Spec: repo.Spec{
		Table:  "skor_aldrette_pasca_anestesi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter", "nip"},
	},
	SetKey: func(m *model.SkorAldrettePascaAnestesi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.SkorAldrettePascaAnestesi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.SkorAldrettePascaAnestesi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkorAldrettePascaAnestesi) []string { return []string{m.KdDokter, m.Nip} },
	Fill: func(m *model.SkorAldrettePascaAnestesi, d request.SkorAldrettePascaAnestesiData) error {
		m.PenilaianSkala1 = support.Nullable(d.PenilaianSkala1)
		m.PenilaianNilai1 = d.PenilaianNilai1
		m.PenilaianSkala2 = support.Nullable(d.PenilaianSkala2)
		m.PenilaianNilai2 = d.PenilaianNilai2
		m.PenilaianSkala3 = support.Nullable(d.PenilaianSkala3)
		m.PenilaianNilai3 = d.PenilaianNilai3
		m.PenilaianSkala4 = support.Nullable(d.PenilaianSkala4)
		m.PenilaianNilai4 = d.PenilaianNilai4
		m.PenilaianSkala5 = support.Nullable(d.PenilaianSkala5)
		m.PenilaianNilai5 = d.PenilaianNilai5
		m.PenilaianTotalnilai = d.PenilaianTotalnilai
		m.Keluar = support.Nullable(d.Keluar)
		m.Instruksi = support.Nullable(d.Instruksi)
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkorAldrettePascaAnestesi]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.SkorAldrettePascaAnestesi) any { return Str(m.KdDokter) }},
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkorAldrettePascaAnestesi) any { return Str(m.Nip) }},
	},
}
