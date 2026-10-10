package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormChecklistKriteriaMasukHcu checklist kriteria masuk HCU (RMChecklistKriteriaMasukHCU).
var FormChecklistKriteriaMasukHcu = &Form[model.ChecklistKriteriaMasukHcu, request.ChecklistKriteriaMasukHcuData]{
	Slug:  "checklist-kriteria-masuk-hcu",
	Label: "checklist kriteria masuk HCU",
	Spec: repo.Spec{
		Table:  "checklist_kriteria_masuk_hcu",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nik"},
	},
	SetKey: func(m *model.ChecklistKriteriaMasukHcu, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.ChecklistKriteriaMasukHcu) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.ChecklistKriteriaMasukHcu) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.ChecklistKriteriaMasukHcu) []string { return []string{derefStr(m.Nik)} },
	Fill: func(m *model.ChecklistKriteriaMasukHcu, d request.ChecklistKriteriaMasukHcuData) error {
		m.Kardiologi1 = strings.TrimSpace(d.Kardiologi1)
		m.Kardiologi2 = strings.TrimSpace(d.Kardiologi2)
		m.Kardiologi3 = strings.TrimSpace(d.Kardiologi3)
		m.Kardiologi4 = strings.TrimSpace(d.Kardiologi4)
		m.Kardiologi5 = strings.TrimSpace(d.Kardiologi5)
		m.Kardiologi6 = strings.TrimSpace(d.Kardiologi6)
		m.Pernapasan1 = strings.TrimSpace(d.Pernapasan1)
		m.Pernapasan2 = strings.TrimSpace(d.Pernapasan2)
		m.Pernapasan3 = strings.TrimSpace(d.Pernapasan3)
		m.Syaraf1 = strings.TrimSpace(d.Syaraf1)
		m.Syaraf2 = strings.TrimSpace(d.Syaraf2)
		m.Syaraf3 = strings.TrimSpace(d.Syaraf3)
		m.Syaraf4 = strings.TrimSpace(d.Syaraf4)
		m.Pencernaan1 = strings.TrimSpace(d.Pencernaan1)
		m.Pencernaan2 = strings.TrimSpace(d.Pencernaan2)
		m.Pencernaan3 = strings.TrimSpace(d.Pencernaan3)
		m.Pencernaan4 = strings.TrimSpace(d.Pencernaan4)
		m.Pembedahan1 = strings.TrimSpace(d.Pembedahan1)
		m.Pembedahan2 = strings.TrimSpace(d.Pembedahan2)
		m.Hematologi1 = strings.TrimSpace(d.Hematologi1)
		m.Hematologi2 = strings.TrimSpace(d.Hematologi2)
		m.Infeksi = strings.TrimSpace(d.Infeksi)
		m.Nik = support.Nullable(d.Nik)
		return nil
	},
	Refs: []Ref[model.ChecklistKriteriaMasukHcu]{
		{Column: "nik", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.ChecklistKriteriaMasukHcu) any { return StrPtr(m.Nik) }},
	},
}
