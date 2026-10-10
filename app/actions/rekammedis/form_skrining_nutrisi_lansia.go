package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningNutrisiLansia skrining nutrisi lansia (RMSkriningNutrisiLansia).
var FormSkriningNutrisiLansia = &Form[model.SkriningNutrisiLansia, request.SkriningNutrisiLansiaData]{
	Slug:  "skrining-nutrisi-lansia",
	Label: "skrining nutrisi lansia",
	Spec: repo.Spec{
		Table:  "skrining_nutrisi_lansia",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningNutrisiLansia, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningNutrisiLansia) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningNutrisiLansia) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningNutrisiLansia) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningNutrisiLansia, d request.SkriningNutrisiLansiaData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Td = strings.TrimSpace(d.Td)
		m.Hr = strings.TrimSpace(d.Hr)
		m.Rr = strings.TrimSpace(d.Rr)
		m.Suhu = strings.TrimSpace(d.Suhu)
		m.Bb = strings.TrimSpace(d.Bb)
		m.Tbpb = strings.TrimSpace(d.Tbpb)
		m.Spo2 = strings.TrimSpace(d.Spo2)
		m.Alergi = strings.TrimSpace(d.Alergi)
		m.Sg1 = strings.TrimSpace(d.Sg1)
		m.Nilai1 = strings.TrimSpace(d.Nilai1)
		m.Sg2 = strings.TrimSpace(d.Sg2)
		m.Nilai2 = strings.TrimSpace(d.Nilai2)
		m.Sg3 = strings.TrimSpace(d.Sg3)
		m.Nilai3 = strings.TrimSpace(d.Nilai3)
		m.Sg4 = strings.TrimSpace(d.Sg4)
		m.Nilai4 = strings.TrimSpace(d.Nilai4)
		m.Sg5 = strings.TrimSpace(d.Sg5)
		m.Nilai5 = strings.TrimSpace(d.Nilai5)
		m.Sg6 = strings.TrimSpace(d.Sg6)
		m.Nilai6 = strings.TrimSpace(d.Nilai6)
		m.TotalHasil = d.TotalHasil
		m.SkorNutrisi = support.Nullable(d.SkorNutrisi)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningNutrisiLansia]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningNutrisiLansia) any { return Str(m.Nip) }},
	},
}
