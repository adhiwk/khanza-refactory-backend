package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkorBromagePascaAnestesi monitoring bromage pasca anestesi (RMMonitoringBromagePascaAnestesi).
var FormSkorBromagePascaAnestesi = &Form[model.SkorBromagePascaAnestesi, request.SkorBromagePascaAnestesiData]{
	Slug:  "skor-bromage-pasca-anestesi",
	Label: "monitoring bromage pasca anestesi",
	Spec: repo.Spec{
		Table:  "skor_bromage_pasca_anestesi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "kd_dokter", "nip"},
	},
	SetKey: func(m *model.SkorBromagePascaAnestesi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.SkorBromagePascaAnestesi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.SkorBromagePascaAnestesi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkorBromagePascaAnestesi) []string { return []string{m.KdDokter, m.Nip} },
	Fill: func(m *model.SkorBromagePascaAnestesi, d request.SkorBromagePascaAnestesiData) error {
		m.PenilaianSkala1 = support.Nullable(d.PenilaianSkala1)
		m.PenilaianNilai1 = d.PenilaianNilai1
		m.Keluar = support.Nullable(d.Keluar)
		m.Instruksi = support.Nullable(d.Instruksi)
		m.KdDokter = strings.TrimSpace(d.KdDokter)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkorBromagePascaAnestesi]{
		{Column: "kd_dokter", Table: "dokter", RefCol: "kd_dokter", Label: "dokter", Value: func(m *model.SkorBromagePascaAnestesi) any { return Str(m.KdDokter) }},
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkorBromagePascaAnestesi) any { return Str(m.Nip) }},
	},
}
