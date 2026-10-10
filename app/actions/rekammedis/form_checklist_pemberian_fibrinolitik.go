package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormChecklistPemberianFibrinolitik checklist pemberian fibrinolitik (RMChecklistPemberianFibrinolitik).
var FormChecklistPemberianFibrinolitik = &Form[model.ChecklistPemberianFibrinolitik, request.ChecklistPemberianFibrinolitikData]{
	Slug:  "checklist-pemberian-fibrinolitik",
	Label: "checklist pemberian fibrinolitik",
	Spec: repo.Spec{
		Table:  "checklist_pemberian_fibrinolitik",
		Keys:   []string{"no_rawat"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.ChecklistPemberianFibrinolitik, key repo.Key) error {
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		return nil
	},
	KeyOf: func(m *model.ChecklistPemberianFibrinolitik) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat}
	},
	Waktu:   func(m *model.ChecklistPemberianFibrinolitik) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.ChecklistPemberianFibrinolitik) []string { return []string{m.Nip} },
	Fill: func(m *model.ChecklistPemberianFibrinolitik, d request.ChecklistPemberianFibrinolitikData) error {
		vTanggal, err := support.ParseDateTime(d.Tanggal)
		if err != nil {
			return err
		}
		m.Tanggal = vTanggal
		m.Nip = strings.TrimSpace(d.Nip)
		m.KontraIndikasi1 = support.Nullable(d.KontraIndikasi1)
		m.KeteranganKontraIndikasi1 = support.Nullable(d.KeteranganKontraIndikasi1)
		m.KontraIndikasi2 = support.Nullable(d.KontraIndikasi2)
		m.KeteranganKontraIndikasi2 = support.Nullable(d.KeteranganKontraIndikasi2)
		m.KontraIndikasi3 = support.Nullable(d.KontraIndikasi3)
		m.KeteranganKontraIndikasi3 = support.Nullable(d.KeteranganKontraIndikasi3)
		m.KontraIndikasi4 = support.Nullable(d.KontraIndikasi4)
		m.KeteranganKontraIndikasi4 = support.Nullable(d.KeteranganKontraIndikasi4)
		m.KontraIndikasi5 = support.Nullable(d.KontraIndikasi5)
		m.KeteranganKontraIndikasi5 = support.Nullable(d.KeteranganKontraIndikasi5)
		m.KontraIndikasi6 = support.Nullable(d.KontraIndikasi6)
		m.KeteranganKontraIndikasi6 = support.Nullable(d.KeteranganKontraIndikasi6)
		m.KontraIndikasi7 = support.Nullable(d.KontraIndikasi7)
		m.KeteranganKontraIndikasi7 = support.Nullable(d.KeteranganKontraIndikasi7)
		m.KontraIndikasi8 = support.Nullable(d.KontraIndikasi8)
		m.KeteranganKontraIndikasi8 = support.Nullable(d.KeteranganKontraIndikasi8)
		m.KontraIndikasi9 = support.Nullable(d.KontraIndikasi9)
		m.KeteranganKontraIndikasi9 = support.Nullable(d.KeteranganKontraIndikasi9)
		m.KontraIndikasi10 = support.Nullable(d.KontraIndikasi10)
		m.KeteranganKontraIndikasi10 = support.Nullable(d.KeteranganKontraIndikasi10)
		m.RisikoTinggi1 = support.Nullable(d.RisikoTinggi1)
		m.KeteranganRisikoTinggi1 = support.Nullable(d.KeteranganRisikoTinggi1)
		m.RisikoTinggi2 = support.Nullable(d.RisikoTinggi2)
		m.KeteranganRisikoTinggi2 = support.Nullable(d.KeteranganRisikoTinggi2)
		m.RisikoTinggi3 = support.Nullable(d.RisikoTinggi3)
		m.KeteranganRisikoTinggi3 = support.Nullable(d.KeteranganRisikoTinggi3)
		m.RisikoTinggi4 = support.Nullable(d.RisikoTinggi4)
		m.KeteranganRisikoTinggi4 = support.Nullable(d.KeteranganRisikoTinggi4)
		m.RisikoTinggi5 = support.Nullable(d.RisikoTinggi5)
		m.KeteranganRisikoTinggi5 = support.Nullable(d.KeteranganRisikoTinggi5)
		m.Kesimpulan = support.Nullable(d.Kesimpulan)
		m.PersyaratanEkgPreStreptase = support.Nullable(d.PersyaratanEkgPreStreptase)
		m.PersyaratanEkgPostStreptase = support.Nullable(d.PersyaratanEkgPostStreptase)
		m.CekTroponin = support.Nullable(d.CekTroponin)
		return nil
	},
	Refs: []Ref[model.ChecklistPemberianFibrinolitik]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.ChecklistPemberianFibrinolitik) any { return Str(m.Nip) }},
	},
}
