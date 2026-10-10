package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningDiabetesMelitus skrining diabetes melitus (RMSkriningDiabetesMelitus).
var FormSkriningDiabetesMelitus = &Form[model.SkriningDiabetesMelitus, request.SkriningDiabetesMelitusData]{
	Slug:  "skrining-diabetes-melitus",
	Label: "skrining diabetes melitus",
	Spec: repo.Spec{
		Table:  "skrining_diabetes_melitus",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningDiabetesMelitus, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningDiabetesMelitus) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningDiabetesMelitus) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningDiabetesMelitus) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningDiabetesMelitus, d request.SkriningDiabetesMelitusData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.Anamnesis1 = support.Nullable(d.Anamnesis1)
		m.Anamnesis2 = support.Nullable(d.Anamnesis2)
		m.Anamnesis3 = support.Nullable(d.Anamnesis3)
		m.Anamnesis4 = support.Nullable(d.Anamnesis4)
		m.Anamnesis5 = support.Nullable(d.Anamnesis5)
		m.Anamnesis6 = support.Nullable(d.Anamnesis6)
		m.Anamnesis7 = support.Nullable(d.Anamnesis7)
		m.Anamnesis8 = support.Nullable(d.Anamnesis8)
		m.Anamnesis9 = support.Nullable(d.Anamnesis9)
		m.Anamnesis10 = support.Nullable(d.Anamnesis10)
		m.Anamnesis11 = support.Nullable(d.Anamnesis11)
		m.Anamnesis12 = support.Nullable(d.Anamnesis12)
		m.BeratBadan = support.Nullable(d.BeratBadan)
		m.TinggiBadan = support.Nullable(d.TinggiBadan)
		m.Imt = support.Nullable(d.Imt)
		m.KasifikasiImt = support.Nullable(d.KasifikasiImt)
		m.HasilGds = support.Nullable(d.HasilGds)
		m.KeteranganGds = support.Nullable(d.KeteranganGds)
		m.HasilGdp = support.Nullable(d.HasilGdp)
		m.KeteranganGdp = support.Nullable(d.KeteranganGdp)
		m.HasilSkrining = support.Nullable(d.HasilSkrining)
		m.KeteranganSkrining = support.Nullable(d.KeteranganSkrining)
		return nil
	},
	Refs: []Ref[model.SkriningDiabetesMelitus]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningDiabetesMelitus) any { return Str(m.Nip) }},
	},
}
