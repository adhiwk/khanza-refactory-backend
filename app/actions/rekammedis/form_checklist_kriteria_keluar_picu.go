package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormChecklistKriteriaKeluarPicu checklist kriteria keluar PICU (RMChecklistKriteriaKeluarPICU).
var FormChecklistKriteriaKeluarPicu = &Form[model.ChecklistKriteriaKeluarPicu, request.ChecklistKriteriaKeluarPicuData]{
	Slug:  "checklist-kriteria-keluar-picu",
	Label: "checklist kriteria keluar PICU",
	Spec: repo.Spec{
		Table:  "checklist_kriteria_keluar_picu",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nik"},
	},
	SetKey: func(m *model.ChecklistKriteriaKeluarPicu, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.ChecklistKriteriaKeluarPicu) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.ChecklistKriteriaKeluarPicu) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.ChecklistKriteriaKeluarPicu) []string { return []string{derefStr(m.Nik)} },
	Fill: func(m *model.ChecklistKriteriaKeluarPicu, d request.ChecklistKriteriaKeluarPicuData) error {
		m.Kondisiklinis1 = strings.TrimSpace(d.Kondisiklinis1)
		m.Kondisiklinis2 = strings.TrimSpace(d.Kondisiklinis2)
		m.Kondisiklinis3 = strings.TrimSpace(d.Kondisiklinis3)
		m.Kondisiklinis4 = strings.TrimSpace(d.Kondisiklinis4)
		m.Kondisiklinis5 = strings.TrimSpace(d.Kondisiklinis5)
		m.Kondisiklinis6 = strings.TrimSpace(d.Kondisiklinis6)
		m.Kebutuhanperawatan1 = strings.TrimSpace(d.Kebutuhanperawatan1)
		m.Kebutuhanperawatan2 = strings.TrimSpace(d.Kebutuhanperawatan2)
		m.Kebutuhanperawatan3 = strings.TrimSpace(d.Kebutuhanperawatan3)
		m.Kebutuhanperawatan4 = strings.TrimSpace(d.Kebutuhanperawatan4)
		m.Tindaklanjut1 = strings.TrimSpace(d.Tindaklanjut1)
		m.Tindaklanjut2 = strings.TrimSpace(d.Tindaklanjut2)
		m.Tindaklanjut3 = strings.TrimSpace(d.Tindaklanjut3)
		m.Tindaklanjut4 = strings.TrimSpace(d.Tindaklanjut4)
		m.Keputusan = strings.TrimSpace(d.Keputusan)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nik = support.Nullable(d.Nik)
		return nil
	},
	Refs: []Ref[model.ChecklistKriteriaKeluarPicu]{
		{Column: "nik", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.ChecklistKriteriaKeluarPicu) any { return StrPtr(m.Nik) }},
	},
}
