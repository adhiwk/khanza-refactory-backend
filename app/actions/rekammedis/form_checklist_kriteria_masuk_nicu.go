package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormChecklistKriteriaMasukNicu checklist kriteria masuk NICU (RMChecklistKriteriaMasukNICU).
var FormChecklistKriteriaMasukNicu = &Form[model.ChecklistKriteriaMasukNicu, request.ChecklistKriteriaMasukNicuData]{
	Slug:  "checklist-kriteria-masuk-nicu",
	Label: "checklist kriteria masuk NICU",
	Spec: repo.Spec{
		Table:  "checklist_kriteria_masuk_nicu",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nik"},
	},
	SetKey: func(m *model.ChecklistKriteriaMasukNicu, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.ChecklistKriteriaMasukNicu) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.ChecklistKriteriaMasukNicu) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.ChecklistKriteriaMasukNicu) []string { return []string{derefStr(m.Nik)} },
	Fill: func(m *model.ChecklistKriteriaMasukNicu, d request.ChecklistKriteriaMasukNicuData) error {
		m.Respirasi1 = strings.TrimSpace(d.Respirasi1)
		m.Respirasi2 = strings.TrimSpace(d.Respirasi2)
		m.Respirasi3 = strings.TrimSpace(d.Respirasi3)
		m.Respirasi4 = strings.TrimSpace(d.Respirasi4)
		m.Prematur1 = strings.TrimSpace(d.Prematur1)
		m.Prematur2 = strings.TrimSpace(d.Prematur2)
		m.Prematur3 = strings.TrimSpace(d.Prematur3)
		m.Kardio1 = strings.TrimSpace(d.Kardio1)
		m.Kardio2 = strings.TrimSpace(d.Kardio2)
		m.Kardio3 = strings.TrimSpace(d.Kardio3)
		m.Neuro1 = strings.TrimSpace(d.Neuro1)
		m.Neuro2 = strings.TrimSpace(d.Neuro2)
		m.Neuro3 = strings.TrimSpace(d.Neuro3)
		m.Metabolik1 = strings.TrimSpace(d.Metabolik1)
		m.Metabolik2 = strings.TrimSpace(d.Metabolik2)
		m.Metabolik3 = strings.TrimSpace(d.Metabolik3)
		m.Kondisilain1 = strings.TrimSpace(d.Kondisilain1)
		m.Kondisilain2 = strings.TrimSpace(d.Kondisilain2)
		m.Kondisilain3 = strings.TrimSpace(d.Kondisilain3)
		m.Kondisilain4 = strings.TrimSpace(d.Kondisilain4)
		m.Keputusan = strings.TrimSpace(d.Keputusan)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nik = support.Nullable(d.Nik)
		return nil
	},
	Refs: []Ref[model.ChecklistKriteriaMasukNicu]{
		{Column: "nik", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.ChecklistKriteriaMasukNicu) any { return StrPtr(m.Nik) }},
	},
}
