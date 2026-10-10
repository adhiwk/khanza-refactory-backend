package rekammedis

import (
	"strings"
	"time"

	request "goravel/app/http/requests/rekammedis"
	model "goravel/app/models/rekammedis"
	repo "goravel/app/repository/rekammedis"
	"goravel/app/support"
)

// FormPenilaianRisikoJatuhNeonatus penilaian risiko jatuh neonatus (RMPenilaianRisikoJatuhNeonatus).
var FormPenilaianRisikoJatuhNeonatus = &Form[model.PenilaianRisikoJatuhNeonatus, request.PenilaianRisikoJatuhNeonatusData]{
	Slug:  "penilaian-risiko-jatuh-neonatus",
	Label: "penilaian risiko jatuh neonatus",
	Spec: repo.Spec{
		Table:  "penilaian_risiko_jatuh_neonatus",
		Keys:   []string{"no_rawat", "tanggal"},
		Waktu:  "tanggal",
		Search: []string{"no_rawat", "nip"},
	},
	SetKey: func(m *model.PenilaianRisikoJatuhNeonatus, key repo.Key) error {
		var err error
		m.NoRawat = strings.TrimSpace(key["no_rawat"])
		if m.Tanggal, err = KeyDateTime(key["tanggal"]); err != nil {
			return err
		}
		return err
	},
	KeyOf: func(m *model.PenilaianRisikoJatuhNeonatus) repo.Key {
		return repo.Key{"no_rawat": m.NoRawat, "tanggal": FmtDateTime(m.Tanggal)}
	},
	Waktu:   func(m *model.PenilaianRisikoJatuhNeonatus) time.Time { return Tm(m.Tanggal) },
	Petugas: func(m *model.PenilaianRisikoJatuhNeonatus) []string { return []string{m.Nip} },
	Fill: func(m *model.PenilaianRisikoJatuhNeonatus, d request.PenilaianRisikoJatuhNeonatusData) error {
		m.Intervensi1 = support.Nullable(d.Intervensi1)
		m.Intervensi2 = support.Nullable(d.Intervensi2)
		m.Intervensi3 = support.Nullable(d.Intervensi3)
		m.Intervensi4 = support.Nullable(d.Intervensi4)
		m.Intervensi5 = support.Nullable(d.Intervensi5)
		m.Intervensi6 = support.Nullable(d.Intervensi6)
		m.Intervensi7 = support.Nullable(d.Intervensi7)
		m.Intervensi8 = support.Nullable(d.Intervensi8)
		m.Intervensi9 = support.Nullable(d.Intervensi9)
		m.Edukasi1 = support.Nullable(d.Edukasi1)
		m.Edukasi2 = support.Nullable(d.Edukasi2)
		m.Edukasi3 = support.Nullable(d.Edukasi3)
		m.Edukasi4 = support.Nullable(d.Edukasi4)
		m.Edukasi5 = support.Nullable(d.Edukasi5)
		m.Sasaran1 = support.Nullable(d.Sasaran1)
		m.Sasaran2 = support.Nullable(d.Sasaran2)
		m.Sasaran3 = support.Nullable(d.Sasaran3)
		m.Sasaran4 = support.Nullable(d.Sasaran4)
		m.Evaluasi1 = support.Nullable(d.Evaluasi1)
		m.Evaluasi2 = support.Nullable(d.Evaluasi2)
		m.Evaluasi3 = support.Nullable(d.Evaluasi3)
		m.Nip = strings.TrimSpace(d.Nip)
		return nil
	},
	Refs: []Ref[model.PenilaianRisikoJatuhNeonatus]{
		{Column: "nip", Table: "petugas", RefCol: "nip", Label: "petugas", Value: func(m *model.PenilaianRisikoJatuhNeonatus) any { return Str(m.Nip) }},
	},
}
