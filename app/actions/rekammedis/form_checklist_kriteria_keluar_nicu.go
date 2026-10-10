package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormChecklistKriteriaKeluarNicu checklist kriteria keluar NICU (RMChecklistKriteriaKeluarNICU).
var FormChecklistKriteriaKeluarNicu = &Form[model.ChecklistKriteriaKeluarNicu, request.ChecklistKriteriaKeluarNicuData]{
	Slug:  "checklist-kriteria-keluar-nicu",
	Label: "checklist kriteria keluar NICU",
	Spec: repo.Spec{
		Table:  "checklist_kriteria_keluar_nicu",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nik"},
	},
	SetKey: func(m *model.ChecklistKriteriaKeluarNicu, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.ChecklistKriteriaKeluarNicu) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.ChecklistKriteriaKeluarNicu) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.ChecklistKriteriaKeluarNicu) []string { return []string{derefStr(m.Nik)} },
	Fill: func(m *model.ChecklistKriteriaKeluarNicu, d request.ChecklistKriteriaKeluarNicuData) error {
		m.Respirasi1 = strings.TrimSpace(d.Respirasi1)
		m.Respirasi2 = strings.TrimSpace(d.Respirasi2)
		m.Respirasi3 = strings.TrimSpace(d.Respirasi3)
		m.Kardio1 = strings.TrimSpace(d.Kardio1)
		m.Kardio2 = strings.TrimSpace(d.Kardio2)
		m.Nutrisi1 = strings.TrimSpace(d.Nutrisi1)
		m.Nutrisi2 = strings.TrimSpace(d.Nutrisi2)
		m.Nutrisi3 = strings.TrimSpace(d.Nutrisi3)
		m.Suhutubuh1 = strings.TrimSpace(d.Suhutubuh1)
		m.Suhutubuh2 = strings.TrimSpace(d.Suhutubuh2)
		m.Infeksi1 = strings.TrimSpace(d.Infeksi1)
		m.Infeksi2 = strings.TrimSpace(d.Infeksi2)
		m.Infeksi3 = strings.TrimSpace(d.Infeksi3)
		m.Keputusan = strings.TrimSpace(d.Keputusan)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nik = support.Nullable(d.Nik)
		return nil
	},
	Refs: []Ref[model.ChecklistKriteriaKeluarNicu]{
		{Column: "nik", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.ChecklistKriteriaKeluarNicu) any { return StrPtr(m.Nik) }},
	},
}
