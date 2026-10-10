package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningFrailtySyndrome skrining frailty syndrome (RMSkriningFrailtySyndrome).
var FormSkriningFrailtySyndrome = &Form[model.SkriningFrailtySyndrome, request.SkriningFrailtySyndromeData]{
	Slug:  "skrining-frailty-syndrome",
	Label: "skrining frailty syndrome",
	Spec: repo.Spec{
		Table:  "skrining_frailty_syndrome",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningFrailtySyndrome, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningFrailtySyndrome) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningFrailtySyndrome) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningFrailtySyndrome) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningFrailtySyndrome, d request.SkriningFrailtySyndromeData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Resistensi = support.Nullable(d.Resistensi)
		m.NilaiResistensi = d.NilaiResistensi
		m.Aktivitas = support.Nullable(d.Aktivitas)
		m.NilaiAktivitas = d.NilaiAktivitas
		m.PenyakitTidakPernah = support.Nullable(d.PenyakitTidakPernah)
		m.PenyakitKanker = support.Nullable(d.PenyakitKanker)
		m.PenyakitGagalJantung = support.Nullable(d.PenyakitGagalJantung)
		m.PenyakitGinjal = support.Nullable(d.PenyakitGinjal)
		m.PenyakitNyeriDada = support.Nullable(d.PenyakitNyeriDada)
		m.PenyakitSeranganJantung = support.Nullable(d.PenyakitSeranganJantung)
		m.PenyakitStroke = support.Nullable(d.PenyakitStroke)
		m.PenyakitAsma = support.Nullable(d.PenyakitAsma)
		m.PenyakitNyeriSendi = support.Nullable(d.PenyakitNyeriSendi)
		m.PenyakitParuKronis = support.Nullable(d.PenyakitParuKronis)
		m.PenyakitHipertensi = support.Nullable(d.PenyakitHipertensi)
		m.PenyakitDiabetes = support.Nullable(d.PenyakitDiabetes)
		m.NilaiPenyakit = d.NilaiPenyakit
		m.UsahaBerjalan = support.Nullable(d.UsahaBerjalan)
		m.NilaiUsahaBerjalan = d.NilaiUsahaBerjalan
		m.BeratBadan = support.Nullable(d.BeratBadan)
		m.NilaiBeratBadan = d.NilaiBeratBadan
		m.NilaiTotal = d.NilaiTotal
		m.HasilSkrining = support.Nullable(d.HasilSkrining)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningFrailtySyndrome]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningFrailtySyndrome) any { return Str(m.Nip) }},
	},
}
