package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormChecklistKriteriaKeluarIsolasi checklist kriteria keluar isolasi (RMChecklistKriteriaKeluarIsolasi).
var FormChecklistKriteriaKeluarIsolasi = &Form[model.ChecklistKriteriaKeluarIsolasi, request.ChecklistKriteriaKeluarIsolasiData]{
	Slug:  "checklist-kriteria-keluar-isolasi",
	Label: "checklist kriteria keluar isolasi",
	Spec: repo.Spec{
		Table:  "checklist_kriteria_keluar_isolasi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nik"},
	},
	SetKey: func(m *model.ChecklistKriteriaKeluarIsolasi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.ChecklistKriteriaKeluarIsolasi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.ChecklistKriteriaKeluarIsolasi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.ChecklistKriteriaKeluarIsolasi) []string { return []string{derefStr(m.Nik)} },
	Fill: func(m *model.ChecklistKriteriaKeluarIsolasi, d request.ChecklistKriteriaKeluarIsolasiData) error {
		m.GejalaMembaik = strings.TrimSpace(d.GejalaMembaik)
		m.TidakAdaIndikasiTransmisi = strings.TrimSpace(d.TidakAdaIndikasiTransmisi)
		m.HasilPenunjangMemenuhi = strings.TrimSpace(d.HasilPenunjangMemenuhi)
		m.KriteriaPedomanTerpenuhi = strings.TrimSpace(d.KriteriaPedomanTerpenuhi)
		m.PersetujuanDpjp = strings.TrimSpace(d.PersetujuanDpjp)
		m.Keputusan = strings.TrimSpace(d.Keputusan)
		m.Alasan = support.Nullable(d.Alasan)
		m.Nik = support.Nullable(d.Nik)
		return nil
	},
	Refs: []Ref[model.ChecklistKriteriaKeluarIsolasi]{
		{Column: "nik", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.ChecklistKriteriaKeluarIsolasi) any { return StrPtr(m.Nik) }},
	},
}
