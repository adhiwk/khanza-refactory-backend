package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormChecklistKriteriaMasukIsolasi checklist kriteria masuk isolasi (RMChecklistKriteriaMasukIsolasi).
var FormChecklistKriteriaMasukIsolasi = &Form[model.ChecklistKriteriaMasukIsolasi, request.ChecklistKriteriaMasukIsolasiData]{
	Slug:  "checklist-kriteria-masuk-isolasi",
	Label: "checklist kriteria masuk isolasi",
	Spec: repo.Spec{
		Table:  "checklist_kriteria_masuk_isolasi",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nik"},
	},
	SetKey: func(m *model.ChecklistKriteriaMasukIsolasi, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.ChecklistKriteriaMasukIsolasi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.ChecklistKriteriaMasukIsolasi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.ChecklistKriteriaMasukIsolasi) []string { return []string{derefStr(m.Nik)} },
	Fill: func(m *model.ChecklistKriteriaMasukIsolasi, d request.ChecklistKriteriaMasukIsolasiData) error {
		m.AirborneTb = strings.TrimSpace(d.AirborneTb)
		m.AirborneCampak = strings.TrimSpace(d.AirborneCampak)
		m.AirborneVarisela = strings.TrimSpace(d.AirborneVarisela)
		m.AirborneZosterDiseminata = strings.TrimSpace(d.AirborneZosterDiseminata)
		m.AirborneLainnya = strings.TrimSpace(d.AirborneLainnya)
		m.DropletCovid19 = strings.TrimSpace(d.DropletCovid19)
		m.DropletInfluenza = strings.TrimSpace(d.DropletInfluenza)
		m.DropletDifteri = strings.TrimSpace(d.DropletDifteri)
		m.DropletPertusis = strings.TrimSpace(d.DropletPertusis)
		m.DropletMeningitis = strings.TrimSpace(d.DropletMeningitis)
		m.DropletLainnya = strings.TrimSpace(d.DropletLainnya)
		m.KontakMdro = strings.TrimSpace(d.KontakMdro)
		m.KontakClostridium = strings.TrimSpace(d.KontakClostridium)
		m.KontakScabies = strings.TrimSpace(d.KontakScabies)
		m.KontakLukaDrainase = strings.TrimSpace(d.KontakLukaDrainase)
		m.KontakDiareInfeksi = strings.TrimSpace(d.KontakDiareInfeksi)
		m.KontakLainnya = strings.TrimSpace(d.KontakLainnya)
		m.KontakErat = strings.TrimSpace(d.KontakErat)
		m.RiwayatPerjalananWabah = strings.TrimSpace(d.RiwayatPerjalananWabah)
		m.RiwayatMdro = strings.TrimSpace(d.RiwayatMdro)
		m.HasilLabRadiologiPositif = strings.TrimSpace(d.HasilLabRadiologiPositif)
		m.GejalaKlinisMenular = strings.TrimSpace(d.GejalaKlinisMenular)
		m.PasienImunokompromis = strings.TrimSpace(d.PasienImunokompromis)
		m.InstruksiDpjp = strings.TrimSpace(d.InstruksiDpjp)
		m.PersetujuanIsolasi = strings.TrimSpace(d.PersetujuanIsolasi)
		m.KelengkapanJaminan = strings.TrimSpace(d.KelengkapanJaminan)
		m.KetersediaanApd = strings.TrimSpace(d.KetersediaanApd)
		m.FasilitasCuciTangan = strings.TrimSpace(d.FasilitasCuciTangan)
		m.TekananNegatifBerfungsi = strings.TrimSpace(d.TekananNegatifBerfungsi)
		m.PintuOtomatisBerfungsi = strings.TrimSpace(d.PintuOtomatisBerfungsi)
		m.IndikasiIsolasi = strings.TrimSpace(d.IndikasiIsolasi)
		m.JenisIsolasi = strings.TrimSpace(d.JenisIsolasi)
		m.DiagnosaIsolasi = strings.TrimSpace(d.DiagnosaIsolasi)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nik = support.Nullable(d.Nik)
		return nil
	},
	Refs: []Ref[model.ChecklistKriteriaMasukIsolasi]{
		{Column: "nik", Table: "pegawai", RefCol: "nik", Label: "pegawai", Value: func(m *model.ChecklistKriteriaMasukIsolasi) any { return StrPtr(m.Nik) }},
	},
}
