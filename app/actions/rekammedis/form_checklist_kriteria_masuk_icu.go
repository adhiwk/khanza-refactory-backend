package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormChecklistKriteriaMasukIcu checklist kriteria masuk ICU (RMChecklistKriteriaMasukICU).
var FormChecklistKriteriaMasukIcu = &Form[model.ChecklistKriteriaMasukIcu, request.ChecklistKriteriaMasukIcuData]{
	Slug:  "checklist-kriteria-masuk-icu",
	Label: "checklist kriteria masuk ICU",
	Spec: repo.Spec{
		Table:  "checklist_kriteria_masuk_icu",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nik"},
	},
	SetKey: func(m *model.ChecklistKriteriaMasukIcu, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.ChecklistKriteriaMasukIcu) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.ChecklistKriteriaMasukIcu) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.ChecklistKriteriaMasukIcu) []string { return []string{derefStr(m.Nik)} },
	Fill: func(m *model.ChecklistKriteriaMasukIcu, d request.ChecklistKriteriaMasukIcuData) error {
		m.Prioritas11 = strings.TrimSpace(d.Prioritas11)
		m.Prioritas12 = strings.TrimSpace(d.Prioritas12)
		m.Prioritas13 = strings.TrimSpace(d.Prioritas13)
		m.Prioritas14 = strings.TrimSpace(d.Prioritas14)
		m.Prioritas15 = strings.TrimSpace(d.Prioritas15)
		m.Prioritas16 = strings.TrimSpace(d.Prioritas16)
		m.Prioritas21 = strings.TrimSpace(d.Prioritas21)
		m.Prioritas22 = strings.TrimSpace(d.Prioritas22)
		m.Prioritas23 = strings.TrimSpace(d.Prioritas23)
		m.Prioritas24 = strings.TrimSpace(d.Prioritas24)
		m.Prioritas25 = strings.TrimSpace(d.Prioritas25)
		m.Prioritas26 = strings.TrimSpace(d.Prioritas26)
		m.Prioritas27 = strings.TrimSpace(d.Prioritas27)
		m.Prioritas28 = strings.TrimSpace(d.Prioritas28)
		m.Prioritas31 = strings.TrimSpace(d.Prioritas31)
		m.Prioritas32 = strings.TrimSpace(d.Prioritas32)
		m.Prioritas33 = strings.TrimSpace(d.Prioritas33)
		m.Prioritas34 = strings.TrimSpace(d.Prioritas34)
		m.KriteriaFisiologisTandaVital1 = strings.TrimSpace(d.KriteriaFisiologisTandaVital1)
		m.KriteriaFisiologisTandaVital2 = strings.TrimSpace(d.KriteriaFisiologisTandaVital2)
		m.KriteriaFisiologisTandaVital3 = strings.TrimSpace(d.KriteriaFisiologisTandaVital3)
		m.KriteriaFisiologisTandaVital4 = strings.TrimSpace(d.KriteriaFisiologisTandaVital4)
		m.KriteriaFisiologisTandaVital5 = strings.TrimSpace(d.KriteriaFisiologisTandaVital5)
		m.KriteriaFisiologisLaborat1 = strings.TrimSpace(d.KriteriaFisiologisLaborat1)
		m.KriteriaFisiologisLaborat2 = strings.TrimSpace(d.KriteriaFisiologisLaborat2)
		m.KriteriaFisiologisLaborat3 = strings.TrimSpace(d.KriteriaFisiologisLaborat3)
		m.KriteriaFisiologisLaborat4 = strings.TrimSpace(d.KriteriaFisiologisLaborat4)
		m.KriteriaFisiologisLaborat5 = strings.TrimSpace(d.KriteriaFisiologisLaborat5)
		m.KriteriaFisiologisLaborat6 = strings.TrimSpace(d.KriteriaFisiologisLaborat6)
		m.KriteriaFisiologisRadiologi1 = strings.TrimSpace(d.KriteriaFisiologisRadiologi1)
		m.KriteriaFisiologisRadiologi2 = strings.TrimSpace(d.KriteriaFisiologisRadiologi2)
		m.KriteriaFisiologisKlinis1 = strings.TrimSpace(d.KriteriaFisiologisKlinis1)
		m.KriteriaFisiologisKlinis2 = strings.TrimSpace(d.KriteriaFisiologisKlinis2)
		m.KriteriaFisiologisKlinis3 = strings.TrimSpace(d.KriteriaFisiologisKlinis3)
		m.KriteriaFisiologisKlinis4 = strings.TrimSpace(d.KriteriaFisiologisKlinis4)
		m.KriteriaFisiologisKlinis5 = strings.TrimSpace(d.KriteriaFisiologisKlinis5)
		m.KriteriaFisiologisKlinis6 = strings.TrimSpace(d.KriteriaFisiologisKlinis6)
		m.KriteriaFisiologisKlinis7 = strings.TrimSpace(d.KriteriaFisiologisKlinis7)
		m.KriteriaFisiologisKlinis8 = strings.TrimSpace(d.KriteriaFisiologisKlinis8)
		m.Nik = support.Nullable(d.Nik)
		return nil
	},
	Refs: []Ref[model.ChecklistKriteriaMasukIcu]{
		{Column: "nik", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.ChecklistKriteriaMasukIcu) any { return StrPtr(m.Nik) }},
	},
}
