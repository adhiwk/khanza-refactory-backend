package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormChecklistKriteriaMasukPicu checklist kriteria masuk PICU (RMChecklistKriteriaMasukPICU).
var FormChecklistKriteriaMasukPicu = &Form[model.ChecklistKriteriaMasukPicu, request.ChecklistKriteriaMasukPicuData]{
	Slug:  "checklist-kriteria-masuk-picu",
	Label: "checklist kriteria masuk PICU",
	Spec: repo.Spec{
		Table:  "checklist_kriteria_masuk_picu",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nik"},
	},
	SetKey: func(m *model.ChecklistKriteriaMasukPicu, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.ChecklistKriteriaMasukPicu) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.ChecklistKriteriaMasukPicu) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.ChecklistKriteriaMasukPicu) []string { return []string{derefStr(m.Nik)} },
	Fill: func(m *model.ChecklistKriteriaMasukPicu, d request.ChecklistKriteriaMasukPicuData) error {
		m.Kriteriaumum1 = strings.TrimSpace(d.Kriteriaumum1)
		m.Kriteriaumum2 = strings.TrimSpace(d.Kriteriaumum2)
		m.Kriteriaumum3 = strings.TrimSpace(d.Kriteriaumum3)
		m.Respirasi1 = strings.TrimSpace(d.Respirasi1)
		m.Respirasi2 = strings.TrimSpace(d.Respirasi2)
		m.Respirasi3 = strings.TrimSpace(d.Respirasi3)
		m.Respirasi4 = strings.TrimSpace(d.Respirasi4)
		m.Kardio1 = strings.TrimSpace(d.Kardio1)
		m.Kardio2 = strings.TrimSpace(d.Kardio2)
		m.Kardio3 = strings.TrimSpace(d.Kardio3)
		m.Kardio4 = strings.TrimSpace(d.Kardio4)
		m.Neuro1 = strings.TrimSpace(d.Neuro1)
		m.Neuro2 = strings.TrimSpace(d.Neuro2)
		m.Neuro3 = strings.TrimSpace(d.Neuro3)
		m.Neuro4 = strings.TrimSpace(d.Neuro4)
		m.Bedah1 = strings.TrimSpace(d.Bedah1)
		m.Bedah2 = strings.TrimSpace(d.Bedah2)
		m.Bedah3 = strings.TrimSpace(d.Bedah3)
		m.Kondisilain1 = strings.TrimSpace(d.Kondisilain1)
		m.Kondisilain2 = strings.TrimSpace(d.Kondisilain2)
		m.Kondisilain3 = strings.TrimSpace(d.Kondisilain3)
		m.Keputusan = strings.TrimSpace(d.Keputusan)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nik = support.Nullable(d.Nik)
		return nil
	},
	Refs: []Ref[model.ChecklistKriteriaMasukPicu]{
		{Column: "nik", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.ChecklistKriteriaMasukPicu) any { return StrPtr(m.Nik) }},
	},
}
