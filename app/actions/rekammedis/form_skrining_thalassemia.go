package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormSkriningThalassemia skrining talasemia (RMSkriningTalasemia).
var FormSkriningThalassemia = &Form[model.SkriningThalassemia, request.SkriningThalassemiaData]{
	Slug:  "skrining-thalassemia",
	Label: "skrining talasemia",
	Spec: repo.Spec{
		Table:  "skrining_thalassemia",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.SkriningThalassemia, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.SkriningThalassemia) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.SkriningThalassemia) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.SkriningThalassemia) []string { return []string{m.Nip} },
	Fill: func(m *model.SkriningThalassemia, d request.SkriningThalassemiaData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.TransfusiDarah = support.Nullable(d.TransfusiDarah)
		m.RutinTransfusi = support.Nullable(d.RutinTransfusi)
		m.SaudaraThalassemia = support.Nullable(d.SaudaraThalassemia)
		m.TumbuhKembangTerlambat = support.Nullable(d.TumbuhKembangTerlambat)
		m.Anemia = support.Nullable(d.Anemia)
		m.Ikterus = support.Nullable(d.Ikterus)
		m.PerutBuncit = support.Nullable(d.PerutBuncit)
		m.GiziKurang = support.Nullable(d.GiziKurang)
		m.FaciesCooley = support.Nullable(d.FaciesCooley)
		m.PerawakanPendek = support.Nullable(d.PerawakanPendek)
		m.HiperpigmentasiKulit = support.Nullable(d.HiperpigmentasiKulit)
		m.Hemoglobin = support.Nullable(d.Hemoglobin)
		m.Mvc = support.Nullable(d.Mvc)
		m.Mchc = support.Nullable(d.Mchc)
		m.DarahTepi = support.Nullable(d.DarahTepi)
		m.TindakLanjut = support.Nullable(d.TindakLanjut)
		return nil
	},
	Refs: []Ref[model.SkriningThalassemia]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.SkriningThalassemia) any { return Str(m.Nip) }},
	},
}
