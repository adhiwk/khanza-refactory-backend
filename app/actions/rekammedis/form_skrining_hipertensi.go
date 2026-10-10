package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningHipertensi skrining hipertensi (RMSkriningHipertensi).
var FormSkriningHipertensi = &Form[model.SkriningHipertensi, request.SkriningHipertensiData]{
	Slug:  "skrining-hipertensi",
	Label: "skrining hipertensi",
	Spec: repo.Spec{
		Table:  "skrining_hipertensi",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningHipertensi, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningHipertensi) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningHipertensi) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningHipertensi) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningHipertensi, d request.SkriningHipertensiData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Anamnesis1 = support.Nullable(d.Anamnesis1)
		m.Anamnesis2 = support.Nullable(d.Anamnesis2)
		m.Anamnesis3 = support.Nullable(d.Anamnesis3)
		m.Anamnesis4 = support.Nullable(d.Anamnesis4)
		m.Anamnesis5 = support.Nullable(d.Anamnesis5)
		m.Anamnesis6 = support.Nullable(d.Anamnesis6)
		m.Anamnesis7 = support.Nullable(d.Anamnesis7)
		m.Anamnesis8 = support.Nullable(d.Anamnesis8)
		m.Sistole = strings.TrimSpace(d.Sistole)
		m.Diastole = strings.TrimSpace(d.Diastole)
		m.KlasifikasiHipertensi = strings.TrimSpace(d.KlasifikasiHipertensi)
		m.HasilSkrining = strings.TrimSpace(d.HasilSkrining)
		m.Keterangan = support.Nullable(d.Keterangan)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningHipertensi]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningHipertensi) any { return Str(m.Nip) }},
	},
}
