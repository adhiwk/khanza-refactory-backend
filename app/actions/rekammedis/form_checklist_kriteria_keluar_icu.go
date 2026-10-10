package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormChecklistKriteriaKeluarIcu checklist kriteria keluar ICU (RMChecklistKriteriaKeluarICU).
var FormChecklistKriteriaKeluarIcu = &Form[model.ChecklistKriteriaKeluarIcu, request.ChecklistKriteriaKeluarIcuData]{
	Slug:  "checklist-kriteria-keluar-icu",
	Label: "checklist kriteria keluar ICU",
	Spec: repo.Spec{
		Table:  "checklist_kriteria_keluar_icu",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nik"},
	},
	SetKey: func(m *model.ChecklistKriteriaKeluarIcu, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.ChecklistKriteriaKeluarIcu) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.ChecklistKriteriaKeluarIcu) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.ChecklistKriteriaKeluarIcu) []string { return []string{derefStr(m.Nik)} },
	Fill: func(m *model.ChecklistKriteriaKeluarIcu, d request.ChecklistKriteriaKeluarIcuData) error {
		m.Kriteria1 = strings.TrimSpace(d.Kriteria1)
		m.Kriteria2 = strings.TrimSpace(d.Kriteria2)
		m.Kriteria3 = strings.TrimSpace(d.Kriteria3)
		m.Kriteria4 = strings.TrimSpace(d.Kriteria4)
		m.Kriteria5 = strings.TrimSpace(d.Kriteria5)
		m.Kriteria6 = strings.TrimSpace(d.Kriteria6)
		m.Kriteria7 = strings.TrimSpace(d.Kriteria7)
		m.Kriteria8 = strings.TrimSpace(d.Kriteria8)
		m.Kriteria9 = strings.TrimSpace(d.Kriteria9)
		m.Kriteria10 = strings.TrimSpace(d.Kriteria10)
		m.Kriteria11 = strings.TrimSpace(d.Kriteria11)
		m.Nik = support.Nullable(d.Nik)
		return nil
	},
	Refs: []Ref[model.ChecklistKriteriaKeluarIcu]{
		{Column: "nik", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.ChecklistKriteriaKeluarIcu) any { return StrPtr(m.Nik) }},
	},
}
