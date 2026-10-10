package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningNutrisiDewasa skrining nutrisi dewasa (RMSkriningNutrisiDewasa).
var FormSkriningNutrisiDewasa = &Form[model.SkriningNutrisiDewasa, request.SkriningNutrisiDewasaData]{
	Slug:  "skrining-nutrisi-dewasa",
	Label: "skrining nutrisi dewasa",
	Spec: repo.Spec{
		Table:  "skrining_nutrisi_dewasa",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningNutrisiDewasa, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningNutrisiDewasa) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningNutrisiDewasa) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningNutrisiDewasa) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningNutrisiDewasa, d request.SkriningNutrisiDewasaData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		if err := Enum("sg1", d.Sg1, "Tidak", "Tidak Yakin", "Ya, 1-5 Kg", "Ya, 6-10 Kg", "Ya, 11-15 Kg", "Ya, >15 Kg"); err != nil {
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
		m.TotalHasil = d.TotalHasil
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.SkriningNutrisiDewasa]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningNutrisiDewasa) any { return Str(m.Nip) }},
	},
}
